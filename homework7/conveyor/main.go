package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		fmt.Println("Канал генерации закрыт")
		close(naturals)
		
	}()

	
	go func() {
		for x:=range naturals{
			squares <- x*x
		}
		fmt.Println("Канал возведения в квадрат закрыт")
		close(squares)
	}()

	
	for x:= range squares{
		fmt.Println(x)
	}
}
