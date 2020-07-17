package main

import (
	"fmt"
	"math"
)

var (
	cathet1, cathet2, hypotenuse float64
)

func main() {
	cathet1 = 150
	cathet2 = 200
	fmt.Println(`Дан прямоугольный треугольник с катетами ` +
		fmt.Sprintf("%.0f", cathet1) + ` ` +
		fmt.Sprintf("%.0f", cathet2) + ` "попугаев"`)
	area := getTriangleArea(cathet1, cathet2)
	fmt.Println("Площадь данного треугольник составит " + fmt.Sprintf("%.0f", area) + ` квадратных "попугаев"`)

	hypotenuse = getHypotenuse(cathet1, cathet2)
	p := getTrianglePerimeter(cathet1, cathet2, hypotenuse)
	fmt.Println("Периметр данного треугольник составит " + fmt.Sprintf("%.0f", p) + `  "попугаев"`)
	fmt.Println("Гипотенуза данного треугольник составит " + fmt.Sprintf("%.0f", hypotenuse) + `  "попугаев"`)
}

func getTriangleArea(cathet1, cathet2 float64) float64 {
	return cathet1 * cathet2
}

func getHypotenuse(cathet1, cathet2 float64) float64 {
	return math.Sqrt(cathet1*cathet1 + cathet2*cathet2)
}

func getTrianglePerimeter(cathet1, cathet2, hypotenuse float64) float64 {
	return cathet1 + cathet2 + hypotenuse
}
