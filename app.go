// A tool that calculates "Tax included" prices for given list of prices and tax rates
package main

import "example.com/m/prices"

func main() {
	// givenPrices := []float64{10, 20, 30}
	givenTaxRates := []float64{0, 0.8, 0.1, 0.12}

	// result := make(map[float64][]float64)

	for _, taxRate := range givenTaxRates {

		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.LoadData()
		priceJob.Process()
	}

}
