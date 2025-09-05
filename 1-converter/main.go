package main

import "fmt"

func main() {
	const usdToEur = 0.8532
	const usdToRub = 81.56
	const eurToRub = 81.56 / usdToEur
}

func getUserInput() string {
	var input string
	fmt.Scan(&input)
	return input
}

func covertCurrensies(n float64, currency1 string, currency2 string) float64 {

}
