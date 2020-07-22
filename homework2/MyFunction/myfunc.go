package main

func isDivBy2(n int) bool {
	return n%2 == 0
}

func isDivBy3(n int) bool {
	return n%3 == 0
}

func fibNumbr(n int) []uint64 {
	arr := make([]uint64, n)
	for i := 0; i < n; i++ {
		switch i {
		case 0:
			arr[i] = 0
		case 1:
			arr[i] = 1
		default:
			arr[i] = arr[i-2] + arr[i-1]
		}
	}
	return arr
}

func sieveOfEratosthenes(n int) []int {
	arr := make([]int, n-1)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 2
	}
	var multiplier int //множитель "скачка"
	var countSieve int //Количество простых чисел в заданном диаппазоне
	for i := 0; i < len(arr); i++ {
		for (arr[i] == 0) && (i < len(arr)-1) {
			if i < len(arr) {
				i++
			}
		}
		var start = true // стартовый элемент, что бы исключить из проверки
		if arr[i] != 0 {
			multiplier = arr[i]
			for j := i; j < len(arr); j = (j + multiplier) {
				if start != true {
					arr[j] = 0

				} else {
					start = false
					countSieve++
				}
			}
		}
	}
	arr1 := make([]int, countSieve)
	var j int
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr1[j] = arr[i]
			j++
		}
	}
	return arr1
}

func isSieve(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
