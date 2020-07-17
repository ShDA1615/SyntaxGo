package main

import (
	"fmt"
)

func main() {

	const dollarRate = 71.23
	fmt.Println("Программа конвнртации рублей в доллары")
	fmt.Println("Введите сумму в рублях для конвертации")
	var ruble float64
	fmt.Scanln(&ruble)
	dollar := ruble / dollarRate
	fmt.Println("За " + fmt.Sprintf("%.2f", ruble) + " рублей можно получить " + fmt.Sprintf("%.2f", dollar) + " долларов.")
}
