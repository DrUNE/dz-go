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
	fmt.Println(eur_to_rub)

}

func readUserInput() (amount float64) {
	fmt.Scan(&amount)
	return
}

func convertCurrency(amount float64, currencyFrom string, currencyTo string) float64 {
	fmt.Printf("selected currencyes %s %s", currencyFrom, currencyTo)
	return amount
}
