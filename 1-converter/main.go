package main

import "fmt"

const USD_TO_EUR = 0.8531
const USD_TO_RUB = 81.1276
const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR
const USD = "USD"
const RUB = "RUB"
const EUR = "EUR"

func main() {
	curr := scanCurrency()
	amount := scanAmount()
	fmt.Println("Input currency to convert:")
	dst := scanCurrency()
	fmt.Print(convert(amount, curr, dst))
}

func convert(amount float64, scrCurrency string, dstCurrency string) float64 {
	if scrCurrency == USD && dstCurrency == EUR {
		return amount * USD_TO_EUR
	} else if scrCurrency == USD && dstCurrency == RUB {
		return amount * USD_TO_RUB
	}
	return amount * EUR_TO_RUB
}

func scanCurrency() string {
	curr := ""
	for curr != RUB && curr != EUR && curr != USD {
		fmt.Println("Input currency: RUB, EUR, USD")
		fmt.Scan(&curr)
	}
	return curr
}

func scanAmount() float64 {
	amount := 0.0
	for amount <= 0 {
		fmt.Println("Input amount more than 0: ")
		fmt.Scan(&amount)
	}

	return amount
}
