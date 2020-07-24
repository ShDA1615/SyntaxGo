package main

import (
	"fmt"
	"math/rand"
)

var queue1 [5]int //вместимость очереди
var pqueue = &queue1

func main() {
	fmt.Println("Сгенерируем очередь и 4 членов. (0 - пусто)")
	for i := 0; i < len(queue1)-1; i++ {
		queue1[i] = i + 1
	}

	fmt.Println("Очередь состоит из:", queue1)
	fmt.Printf("\n")
	fmt.Println("Один вышел из очереди")
	fmt.Println("Забрали", removeFromQueue(pqueue))
	fmt.Println(queue1)
	fmt.Printf("\n")
	fmt.Println("В очередь хотят ещё трое")
	for i := 1; i < 4; i++ {
		n := rand.Intn(100)
		out := addToQueue(pqueue, n)
		fmt.Println("Добавили :", n)
		if out != 0 {
			fmt.Println("Выбросили :", out)
		}
		fmt.Println("Очередь состоит из:", queue1)
		fmt.Printf("\n")

	}

}

func addToQueue(q *[5]int, n int) int {
	var out = 0

	out = q[len(q)-1]

	for i := len(q) - 1; i > 0; i-- {
		q[i] = q[i-1]
	}
	q[0] = n
	return out
}

func removeFromQueue(q *[5]int) int {
	var out = 0
	for i := len(q) - 1; i >= 0; i-- {
		if q[i] != 0 {
			out = q[i]
			q[i] = 0
			break
		}

	}

	return out
}
