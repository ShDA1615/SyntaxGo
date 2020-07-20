package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Введите натуральное число больше единицы")
	fmt.Scanln(&n)

	if n > 1 {
		//применение switch
		switch isDivBy2(n) {
		case true:
			fmt.Println("Число ", n, "чётное")
		case false:
			fmt.Println("Число ", n, "не чётное")
		}

		//Применение if
		if isDivBy3(n) == true {
			fmt.Println("Число", n, "делится на 3 без остатка.")
		} else {
			fmt.Println("Число", n, "не делится на 3 без остатка.")
		}
	} else {
		fmt.Println("Вы ввели значение не соответствующее условию")
	}

	fmt.Println("")
	fmt.Println("")

	fmt.Println("Вывод на экран  N первых чисел Фибоначчи, начиная от 0. По умолчанию N = 100")
	fmt.Println("Введите значение N.")

	fmt.Scanln(&n)
	if n < 3 {
		fmt.Println("Вы ввели слишком маленькое значение. Присваивается значение по умолчанию")
		n = 100
	}
	fmt.Println(fibNumbr(n))

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Нахождение простых чисел в диаппазоне от 2 до N. По умолчанию N = 100")
	fmt.Println("Введите значение N.")
	fmt.Scanln(&n)
	if n < 3 {
		fmt.Println("Вы ввели слишком маленькое значение. Присваивается значение по умолчанию")
		n = 100
	}
	arr := sieveOfEratosthenes(n)
	fmt.Println("В диаппазоне от 2 до", n, " ", len(arr), "простых чисел")
	fmt.Println(arr)
	fmt.Println("")
	fmt.Println("")

	fmt.Println("Введите целое положительное число:")
	fmt.Scanln(&n)
	if isSieve(n) == true {
		fmt.Println("Ваше число относится к простым")
	} else {
		fmt.Println("Ваше число не относится к простым")
	}
	fmt.Println("")
	fmt.Println("")

	fmt.Println("Задача № 3*. Заполнить массив из 100 элементов различными простыми числами.")
	fmt.Println("Нажмите ввод.")
	fmt.Scanln(&n)
	arrFib100 := make([]int, 100)
	var countFib = 0
	var i = 2
	for countFib < 100 {
		if isSieve(i) == true {
			arrFib100[countFib] = i
			countFib++
		}
		i++
	}
	fmt.Println(arrFib100)
}
