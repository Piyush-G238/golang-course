package main

import (
	"fmt"

	"example.com/m/filemanager"
	"example.com/m/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChan := make([]chan bool, len(taxRates))

	errChan := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// err := priceJob.Process()

		doneChan[index] = make(chan bool)
		errChan[index] = make(chan error)

		go priceJob.Process(doneChan[index], errChan[index])

		/*
		 select statement is used to choose which of a set of possible send or receive operations will proceed.
		 It allows you to wait on multiple communication operations simultaneously and proceed when one or more of them completes.
		*/

		select {
		case err := <-errChan[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChan[index]:
			fmt.Println("done")
		}

		// <-doneChan[index]

		/*
			if err != nil {
				fmt.Println("Could not process job")
				fmt.Println(err)
			}
		*/
	}
}
