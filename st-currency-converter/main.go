package main

import (
	"fmt"
	"strings"
)

var rates = map[string]float64{
	"USD": 1.0,   // Доллар США
	"EUR": 0.92,  // Евро
	"RUB": 90.0,  // Российский рубль
	"JPY": 157.0, // Японская иена
	"CNY": 7.25,  // Китайский юань
	"GBP": 0.78,  // Британский фунт
	"KZT": 460.0, // Казахстанский тенге
	"TRY": 32.5,  // Турецкая лира
	"INR": 83.0,  // Индийская рупия
	"BRL": 5.12,  // Бразильский реал
	"AUD": 1.50,  // Австралийский доллар
	"CAD": 1.36,  // Канадский доллар
	"CHF": 0.89,  // Швейцарский франк
	"SEK": 10.8,  // Шведская крона
	"NOK": 10.5,  // Норвежская крона
}

var currency = make([]string, 0, len(rates))

func init() {
	for key := range rates {
		currency = append(currency, key)
	}
}

func main() {
	fmt.Println("Добро пожаловать в конвертер валют!")

	for {
		fmt.Println("Доступные валюты для конвертации:")
		for i, currency := range currency {
			fmt.Printf("%d. %s\n", i+1, currency)
		}

		var amount float64
		fmt.Println("Введите сумму в USD: ")
		_, err := fmt.Scan(&amount)
		if err != nil || amount <= 0 {
			fmt.Println("Ошибка. Введите сумму больше 0.")
			continue
		}

		var currencyIdx int
		fmt.Println("Введите номер валюты для конвертации из вышеперечисленных: ")
		_, err = fmt.Scan(&currencyIdx)
		if err != nil || currencyIdx < 1 || currencyIdx > len(currency) {
			fmt.Println("Ошибка. Введите корректный номер валюты для конвертации из вышеперечисленных")
			continue
		}

		selectedCurrency := rates[currency[currencyIdx-1]]

		converted := amount * selectedCurrency
		fmt.Printf("%.2f USD = %.2f %s\n", amount, converted, currency[currencyIdx-1])

		var continuation string
		fmt.Println("Хотите продолжить? (y/n)")
		fmt.Scan(&continuation)
		if strings.ToLower(continuation) != "y" {
			break
		}
	}
}
