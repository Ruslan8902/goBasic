package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	var currencies = []string{"EUR", "RUB", "USD"}

	var currencyRates = map[string]float64{
		"USDtoEUR": 0.8532,
		"USDtoRUB": 81.56,
		"EURtoRUB": 95.48,
		"EURtoUSD": 1.17,
		"RUBtoUSD": 0.012262,
		"RUBtoEUR": 0.010473,
	}
	source := getSourceCurrency(&currencies)
	target := getTargetCurrency(source, &currencies)
	q := getMoneyQuantity()
	result := covertCurrensies(q, source, target, &currencyRates)
	fmt.Printf("%.2f", result)

}

func getSourceCurrency(currencies *[]string) string {
	var input string
	for {
		fmt.Println("Введите код валюты, из которой необходимо перевести деньги.")
		fmt.Println("Доступные валюты: EUR, USD, RUB")
		fmt.Scan(&input)

		if !slices.Contains(*currencies, input) {
			continue
		} else {
			break
		}
	}
	return input

}

func getTargetCurrency(source string, currencies *[]string) string {
	var input string
	var availableCurrencies []string

	for _, curr := range *currencies {
		if curr != source {
			availableCurrencies = append(availableCurrencies, curr)
		}
	}

	for {
		fmt.Printf("Введите код валюты, в которую необходимо перевести деньги. Не вводите %s\n", source)
		fmt.Printf("Доступные валюты: %s\n", strings.Join(availableCurrencies, ", "))
		fmt.Scan(&input)

		if !slices.Contains(availableCurrencies, input) {
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

func covertCurrensies(n float64, source string, target string, currencyRates *map[string]float64) float64 {
	var currencyRatesKey = source + "to" + target
	return n * (*currencyRates)[currencyRatesKey]
}
