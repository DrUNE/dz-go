package main

import (
	"fmt"
)

const (
	usd_to_eur = 0.878
	usd_to_rub = 73.76
	eur_to_rub = usd_to_rub / usd_to_eur
)

func main() {
	amount, currencyFrom, currencyTo := readUserInput()
	convertedAmount := convertCurrency(amount, currencyFrom, currencyTo)
	fmt.Printf("Результат конвертации: %.2f %s = %.2f %s", amount, currencyFrom, convertedAmount, currencyTo)
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

func convertCurrency(amount float64, currencyFrom string, currencyTo string) (convertedAmount float64) {
	switch {
	case currencyFrom == "usd" && currencyTo == "eur":
		convertedAmount = amount * usd_to_eur
	case currencyTo == "usd" && currencyFrom == "eur":
		convertedAmount = amount / usd_to_eur
	case currencyFrom == "usd" && currencyTo == "rub":
		convertedAmount = amount * usd_to_rub
	case currencyTo == "usd" && currencyFrom == "rub":
		convertedAmount = amount / usd_to_rub
	case currencyFrom == "eur" && currencyTo == "rub":
		convertedAmount = amount * eur_to_rub
	case currencyTo == "eur" && currencyFrom == "rub":
		convertedAmount = amount / eur_to_rub
	default:
		convertedAmount = amount
	}
	return
}
