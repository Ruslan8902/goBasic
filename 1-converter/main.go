package main

import (
	"fmt"
)

const (
	EUR = "EUR"
	RUB = "RUB"
	USD = "USD"
)

var currencyRates = map[string]float64{
	"USDtoEUR": 0.8532,
	"USDtoRUB": 81.56,
	"EURtoRUB": 95.48,
	"EURtoUSD": 1.17,
	"RUBtoUSD": 0.012262,
	"RUBtoEUR": 0.010473,
}

func main() {
	source := getSourceCurrency()
	target := getTargetCurrency(source)
	q := getMoneyQuantity()
	result := covertCurrensies(q, source, target)
	fmt.Printf("%.2f", result)

}

func getSourceCurrency() string {
	var input string
	for {
		fmt.Println("Введите код валюты, из которой необходимо перевести деньги.")
		fmt.Println("Доступные валюты: EUR, USD, RUB")
		fmt.Scan(&input)

		if input != EUR && input != USD && input != RUB {
			continue
		} else {
			break
		}
	}
	return input

}

func getTargetCurrency(source string) string {
	var input string
	var availableCurrency1 string
	var availableCurrency2 string

	switch source {
	case USD:
		availableCurrency1 = EUR
		availableCurrency2 = RUB
	case EUR:
		availableCurrency1 = USD
		availableCurrency2 = RUB
	case RUB:
		availableCurrency1 = USD
		availableCurrency2 = EUR
	}

	for {
		fmt.Printf("Введите код валюты, в которую необходимо перевести деньги. Не вводите %s\n", source)
		fmt.Printf("Доступные валюты: %s, %s\n", availableCurrency1, availableCurrency2)
		fmt.Scan(&input)

		if input != availableCurrency2 && input != availableCurrency1 {
			continue
		} else {
			break
		}
	}
	return input

}

func getMoneyQuantity() float64 {
	var input float64
	for {
		fmt.Println("Сколько хотите перевести? Введите число.")
		_, err := fmt.Scan(&input)

		if err != nil || input <= 0 {
			continue
		} else {
			break
		}
	}
	return input
}

func covertCurrensies(n float64, source string, target string) float64 {
	var currencyRatesKey = source + "to" + target
	return n * currencyRates[currencyRatesKey]
}
