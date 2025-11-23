package main

import "fmt"

const USD = "USD"
const RUB = "RUB"
const EUR = "EUR"

func main() {

	var exchangeRates = map[string]float64{
		USD: 1.0,
		EUR: 0.8531,
		RUB: 81.1276,
	}

	curr := scanCurrency(&exchangeRates)
	amount := scanAmount()
	fmt.Println("Input currency to convert:")
	dst := scanCurrency(&exchangeRates)
	result := convert(&exchangeRates, amount, curr, dst)
	fmt.Printf("%.2f %s = %.2f %s\n", amount, curr, result, dst)
}

func convert(exchangeRates *map[string]float64, amount float64, srcCurrency string, dstCurrency string) float64 {
	amountInUSD := amount / (*exchangeRates)[srcCurrency]
	return amountInUSD * (*exchangeRates)[dstCurrency]
}

func scanCurrency(exchangeRates *map[string]float64) string {
	curr := ""
	for !isValidCurrency(exchangeRates, curr) {
		fmt.Println("Input currency: RUB, EUR, USD")
		fmt.Scan(&curr)
	}
	return curr
}

func isValidCurrency(exchangeRates *map[string]float64, curr string) bool {
	_, exists := (*exchangeRates)[curr]
	return exists
}

func scanAmount() float64 {
	amount := 0.0
	for amount <= 0 {
		fmt.Println("Input amount more than 0: ")
		fmt.Scan(&amount)
	}
	return amount
}
