package main

import (
	"fmt"
	"time"
)

//марку авто, год выпуска, объем багажника/кузова, запущен ли двигатель, открыты ли окна, насколько заполнен объем багажника
type tCar struct {
	brand              string  //марка
	year               int     //год выпуска на 90000 часов наза от текущей даты
	trunkVolume        float32 //обьем багажника
	isRun              bool    //заведен
	wind               [5]bool //окна закрыты или нет
	trunkVolumePercent float32 //% заполнения багажника
}

func main() {
	var car tCar
	car.brand = "Волга"
	car.year = time.Now().Add(-90000 * time.Hour).Year() //год выпуска на 90000 часов наза от текущей даты
	car.trunkVolume = 100.00
	car.trunkVolumePercent = 45.5
	car.wind = [5]bool{true, true, true, true, true}
	fmt.Println(car)
	fmt.Printf("\n\n")
	fmt.Println("Открываем перевые два окна")
	car.wind[0] = false
	car.wind[1] = false
	fmt.Println(car)
	fmt.Printf("\n\n")
	fmt.Println("Заводим")
	car.isRun = true
	fmt.Println(car)
	fmt.Printf("\n\n")

}
