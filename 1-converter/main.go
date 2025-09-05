package main

import (
	"fmt"
)

const (
	EUR      = "EUR"
	RUB      = "RUB"
	USD      = "USD"
	usdToEur = 0.8532
	// usdToRub = 81.56
	// eurToRub = 81.56 / usdToEur
)

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
		fmt.Println("Доступные валюты: EUR, USD")
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
	var availableCurrency string

	switch source {
	case USD:
		availableCurrency = EUR
	case EUR:
		availableCurrency = USD
	}

	for {
		fmt.Printf("Введите код валюты, в которую необходимо перевести деньги. Не вводите %s\n", source)
		fmt.Printf("Доступная валюта: %s\n", availableCurrency)
		fmt.Scan(&input)

		if input != availableCurrency {
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
	var res float64
	switch {
	case source == EUR && target == USD:
		res = n / usdToEur
	case source == USD && target == EUR:
		res = n * usdToEur
	}
	return res
}
