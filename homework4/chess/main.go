package main

import (
	"fmt"
	"strconv"
)

type Position struct {
	X, Y int
}
type Horse struct {
	CurrentPosition Position
	Letter          string
	DeltaPosition   [8]Position
	NextPosition    map[string]Position
}

func intToShessAdres(x, y int) string {
	if (x >= 97 && x <= 104) && (y >= 1 && y <= 8) {
		return string(x) + strconv.Itoa(y)
	} else {
		return ""
	}
}

func (h *Horse) setNewPosition(s string) {
	h.CurrentPosition.X = h.NextPosition[s].X
	h.CurrentPosition.Y = h.NextPosition[s].Y
	h.Letter = s
}

func (h *Horse) printPosition() {
	fmt.Printf("\nКонь на %s\n", h.Letter)
}
func (h *Horse) printNextPosition() {
	fmt.Println("Выберите ход:")
	for key, _ := range h.NextPosition {
		fmt.Printf("\t %s", key)
	}
	fmt.Printf("\n\n")
}

func (h *Horse) getNextPosition() {
	for key, _ := range h.NextPosition {
		delete(h.NextPosition, key)
	}
	for i := 0; i < 8; i++ {
		s := intToShessAdres(h.CurrentPosition.X+h.DeltaPosition[i].X, h.CurrentPosition.Y+h.DeltaPosition[i].Y)
		if s != "" {
			h.NextPosition[s] = Position{h.CurrentPosition.X + h.DeltaPosition[i].X, h.CurrentPosition.Y + h.DeltaPosition[i].Y}
		}
	}

}

func main() {

	horse := Horse{
		DeltaPosition: [8]Position{{X: 1, Y: 2}, {X: 2, Y: 1}, {X: 2, Y: -1}, {X: 2, Y: -2},
			{X: 1, Y: -2}, {X: -2, Y: -1}, {X: -2, Y: 1}, {X: -1, Y: 2}},
		CurrentPosition: Position{X: 98, Y: 1}, Letter: "b1",
		NextPosition: make(map[string]Position)}
	var n string

	for {
		horse.printPosition()
		horse.getNextPosition()
		horse.printNextPosition()

		fmt.Scanln(&n)
		if n == "exit" {
			break
		}
		horse.setNewPosition(n)
		horse.getNextPosition()

	}
}
