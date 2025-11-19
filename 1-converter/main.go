package main

import "fmt"

func main() {
	const USD_TO_EUR = 0.8531
	const USD_TO_RUB = 81.1276
	const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR

}

func convert(amount float64, scrCurrency string, dstCurrency string) {

}

func scan() (float64, string, string) {
	var amount float64
	var src, dst string

	fmt.Scan(&amount, &src, &dst)

	return amount, src, dst
}
