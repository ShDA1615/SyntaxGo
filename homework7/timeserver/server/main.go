package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	cancel := make(chan string)

	go func() {
		fmt.Println("Для завершения наберите exit")
		var s string
		fmt.Scanln(&s)
		fmt.Println("s", s)
		if s == "exit" {

			fmt.Println("<-cancel")
			cancel <- s
			fmt.Println(<-cancel)
		} else {
			fmt.Println("exit <>", s)
		}
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		select {
		case <-cancel:
			fmt.Println("выходим")
			return
		default:
			fmt.Println("go handleConn(conn)")
			go handleConn(conn)

		}
	}
	log.Println("Заваершаем работу")
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
