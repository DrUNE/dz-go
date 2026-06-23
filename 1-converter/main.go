package main

import "fmt"

const (
	usd_to_eur = 0.878
	usd_to_rub = 73.76
	eur_to_rub = usd_to_rub / usd_to_eur
)

func main() {
	fmt.Println(eur_to_rub)
}
