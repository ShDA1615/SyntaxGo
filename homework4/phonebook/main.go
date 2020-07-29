package main

import (
	"fmt"
	"sort"
)

type contact struct {
	Name   string
	Adress string
	Phone  []string
}

type phonebook []contact

func (p phonebook) Len() int {
	return len(p)
}
func (p phonebook) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}
func (p phonebook) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p phonebook) PrintData() {
	for rec, pp := range p {
		fmt.Println("Запись №", rec+1)
		fmt.Printf("\t Имя: %s\n", pp.Name)
		fmt.Printf("\t Телефоны: \n")
		for num, phone := range p[rec].Phone {
			fmt.Printf("\t\t%v) %s\n", num+1, phone)
		}
	}
}

func main() {
	ph := make(phonebook, 0)
	ph = append(ph, contact{"Игорь", "Москва", []string{"+79011111111"}})
	ph = append(ph, contact{"Андрей", "", []string{"+79022222222"}})
	ph[len(ph)-1].Phone = append(ph[len(ph)-1].Phone, "+79055555555")
	ph = append(ph, contact{"Пётр", "Самара", []string{"+79022222222"}})
	fmt.Println("Телефонный справочник")
	ph.PrintData()
	sort.Sort(ph)
	fmt.Println("Телефонный справочник после сортировки")
	ph.PrintData()
}
