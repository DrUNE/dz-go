package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!$%*&#")

type Account struct {
	Login    string
	Password string
	Url      string
}

func (a *Account) generatePassword(length int) {
	a.Password = generatePassword(length)
}
func newAccount(login, urlString, password string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	account := &Account{
		Login:    login,
		Url:      urlString,
		Password: password,
	}
	if account.Password == "" {
		account.generatePassword(12)
	}
	return account, nil
}
func main() {
	login := promptData("Введите логин")
	urlString := promptData("Введите URL")
	password := promptData("Введите пароль")
	myAccount, err := newAccount(login, urlString, password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(myAccount)
}

func promptData(s string) string {
	var input string
	fmt.Print(s + ": ")
	fmt.Scanln(&input)
	return input
}

func generatePassword(length int) string {
	password := make([]rune, length)
	for i := range length {
		password[i] = runes[rand.Intn(len(runes))]
	}
	return string(password)
}
