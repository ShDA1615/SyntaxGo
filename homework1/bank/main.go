package main

import "fmt"

func main() {

	fmt.Println("Программа расчета процентов по вкладу")

	fmt.Println("Введите сумму вклада (руб.):")
	var ruble float64
	fmt.Scanln(&ruble)
	var percent float64
	fmt.Println("Введите желаемый процент:")
	fmt.Scanln(&percent)

	ruble = getSum(ruble, percent, 5)
	fmt.Println("Через 5 лет будет " + fmt.Sprintf("%.2f", ruble) + " рублей")

}

func getSum(ruble, percent float64, year int) float64 {
	var sum float64
	sum = ruble
	for i := 1; i <= year; i++ {
		sum = sum + (sum/100)*percent
	}
	return sum
}
