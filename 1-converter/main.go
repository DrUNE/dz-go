package main

import (
	"fmt"
)

const (
	usd = "usd"
	eur = "eur"
	rub = "rub"
)

type Rate = map[string]float64
type CurrencyRate = map[string]Rate

var currencyFromRate = CurrencyRate{
	usd: Rate{
		usd: 1.0,
		eur: 0.878,
		rub: 73.76,
	},
	eur: Rate{
		usd: 1.14,
		eur: 1.0,
		rub: 87.65,
	},
	rub: Rate{
		usd: 0.013,
		eur: 0.011,
		rub: 1.0,
	},
}

func main() {
	amount, currencyFrom, currencyTo := readUserInput()
	convertedAmount := convertCurrency(amount, currencyFrom, currencyTo)
	fmt.Printf("Результат конвертации: %.2f %s = %.2f %s\n", amount, currencyFrom, convertedAmount, currencyTo)
}

func readUserInput() (amount float64, currencyFrom string, currencyTo string) {

	for {
		fmt.Printf("Введите валюту (usd/eur/rub): ")
		fmt.Scan(&currencyFrom)
		if currencyFrom == "usd" || currencyFrom == "eur" || currencyFrom == "rub" {
			break
		}
		fmt.Printf("Введеная валюта (%s) неверна, повторите ввод.\n", currencyFrom)
	}

	for {
		fmt.Printf("Введите число: ")
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("scaned: %f,  Неверное число, повторите ввод\n", amount)
		} else {
			if amount > 0 {
				break
			} else {
				fmt.Printf("Введите число больше 0\n")
			}
		}
	}

	for {
		fmt.Printf("Введите валюту в которую произойдет конвертация (usd/eur/rub): ")
		fmt.Scan(&currencyTo)
		if currencyTo == "usd" || currencyTo == "eur" || currencyTo == "rub" {
			break
		}
		fmt.Printf("Введеная валюта (%s) неверна, повторите ввод.\n", currencyTo)
	}

	return
}

func convertCurrency(amount float64, currencyFrom string, currencyTo string) float64 {
	return amount * getCurrencyRate(currencyFrom, currencyTo)
}

func getCurrencyRate(currencyFrom string, currencyTo string) float64 {
	return currencyFromRate[currencyFrom][currencyTo]
}
