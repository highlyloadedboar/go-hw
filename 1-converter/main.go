package main

import "fmt"

const USD = "USD"
const RUB = "RUB"
const EUR = "EUR"

var exchangeRates = map[string]float64{
	USD: 1.0,
	EUR: 0.8531,
	RUB: 81.1276,
}

func main() {
	curr := scanCurrency()
	amount := scanAmount()
	fmt.Println("Input currency to convert:")
	dst := scanCurrency()
	result := convert(amount, curr, dst)
	fmt.Printf("%.2f %s = %.2f %s\n", amount, curr, result, dst)
}

func convert(amount float64, srcCurrency string, dstCurrency string) float64 {
	amountInUSD := amount / exchangeRates[srcCurrency]
	return amountInUSD * exchangeRates[dstCurrency]
}

func scanCurrency() string {
	curr := ""
	for !isValidCurrency(curr) {
		fmt.Println("Input currency: RUB, EUR, USD")
		fmt.Scan(&curr)
	}
	return curr
}

func isValidCurrency(curr string) bool {
	_, exists := exchangeRates[curr]
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
