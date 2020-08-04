package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("Test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	for {
		record, e := reader.Read()
		if e != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}
