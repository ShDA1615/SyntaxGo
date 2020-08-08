package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	responses := make(chan string, 3)
	var duration time.Duration
	var name string
	go func() {
		name, duration = request("https://golang.org/")
		responses <- name
	}()
	go func() {

		name, duration = request("https://www.google.com/")
		responses <- name
	}()
	go func() {

		name, duration = request("https://yandex.ru/")
		responses <- name
	}()
	log.Println(<-responses, duration) // возврат самого быстрого ответа
	log.Println(<-responses, duration)
	log.Println(<-responses, duration)

}

func request(hostname string) (string, time.Duration) {
	t0 := time.Now()
	_, err := http.Get(hostname)
	if err != nil {
		log.Fatalln(err)
	}
	t1 := time.Now()

	return hostname, t1.Sub(t0)
}
