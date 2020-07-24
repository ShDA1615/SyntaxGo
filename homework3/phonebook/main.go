package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type contact struct {
	Adress string
	Phone  []string
}

func main() {
	phonebook := make(map[string]contact)
	phonebook["Василий"] = contact{"Москва", []string{"+79161111111"}}
	phonebook["Иван"] = contact{"Саранск", []string{"+79162222222", "+79163333333"}}
	fmt.Printf("Первая map-ка: \n\t %s\n", phonebook)

	//Тут попытка добавить ещё один номер.
	//Как бы это сделать без дополнительных переменных?
	fmt.Println("Добавляем Василию ещё один номер телефона")
	user := contact(phonebook["Василий"])
	user.Phone = append(user.Phone, "+79164444444")
	phonebook["Василий"] = user
	fmt.Printf("Первая map-ка: \n\t %s\n", phonebook)
	//-----------------------------------------------

	jsonStr, err := json.Marshal(phonebook)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("JSON, полученный из первой map-ки: \n\t %s\n", string(jsonStr))
	phonebook1 := make(map[string]contact)
	fmt.Printf("\n")
	fmt.Println(`Записываем JSON в файл "phonebook.json"`)
	fmt.Printf("\n")
	if err := ioutil.WriteFile("phonebook.json", jsonStr, 0777); err != nil {
		fmt.Println(err.Error())
	}
	data, _ := ioutil.ReadFile("phonebook.json")
	fmt.Printf("JSON, прочитанный из файла: \n\t %s\n", string(data))
	if err := json.Unmarshal(data, &phonebook1); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Вторая map-ка с загруженными данными: \n\t %s\n", phonebook1)
}
