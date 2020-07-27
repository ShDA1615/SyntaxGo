package main

import (
	"fmt"

	"./calculator"
)

func printHelp() {

}
func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}
		//input = "sqrt(4)"
		if input == "exit" {
			break
		}

		if input == "help" {
			calculator.Help()
			continue
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
