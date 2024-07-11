package prices

import (
	"bufio"
	"fmt"
	"os"

	"example.com/m/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	Prices            []float64
	TaxIncludedPrices map[string][]float64
}

func NewTaxIncludedPriceJob(taxRate float64) TaxIncludedPriceJob {

	return TaxIncludedPriceJob{
		TaxRate: taxRate,
		// Prices:  prices,
	}
}

func (job TaxIncludedPriceJob) Process() {

	result := make(map[string]float64)

	for _, price := range job.Prices {
		priceStr := fmt.Sprintf("%.2f", price)
		result[priceStr] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, _err := os.Open("prices.txt")

	if _err != nil {
		fmt.Println("Could not open file!, ", _err)
		return
	}

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	_err = scanner.Err()
	if _err != nil {
		fmt.Println("Failed to read file!, ", _err)
		file.Close()
		return
	}

	for _, val := range lines {
		var currentPrice, _err = conversion.StringToFloat(val)
		// strconv.ParseFloat(val, 64)

		if _err != nil {
			fmt.Println("Coverting from string to float64 failed.")
			file.Close()
			return
		}
		job.Prices = append(job.Prices, currentPrice)
	}

	// strconv.ParseFloat()
}
