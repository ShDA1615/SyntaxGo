package main

import (
	"fmt"
)

type Electrics interface {
	On()
	Off()
	Prepare()
}

type Lamp struct {
	Name       string
	Voltage    int
	IsPrepared bool
	IsRun      bool
}

func (l *Lamp) Prepare() {
	l.Voltage = 220
	l.IsPrepared = true
}

func (l *Lamp) isPrepared() bool {
	return l.Voltage == 220 && l.IsPrepared == true
}

func (l *Lamp) On() {
	if l.IsPrepared == true {
		l.IsRun = true
	}
}

func (l *Lamp) Off() {
	if l.IsRun == true {
		l.IsRun = false

	}
}

type Motor struct {
	Name       string
	Voltage    int
	IsPrepared bool
	IsRun      bool
}

func (m *Motor) Prepare() {
	m.Voltage = 380
	m.IsPrepared = true
}

func (m *Motor) On() {
	if m.IsPrepared == true {
		m.IsRun = true
	}
}

func (m *Motor) Off() {
	if m.IsRun == true {
		m.IsRun = false

	}
}

func Prepare(e ...Electrics) {
	for _, e := range e {
		e.Prepare()
	}
}

func SetOn(e ...Electrics) {
	for _, e := range e {
		e.On()
	}
}

func SetOff(e ...Electrics) {
	for _, e := range e {
		e.Off()
	}
}

func main() {
	l1 := &Lamp{
		Name:       "Лампа 1",
		Voltage:    0,
		IsPrepared: false,
		IsRun:      false,
	}

	m1 := &Motor{
		Name:       "Двигатель 1",
		Voltage:    0,
		IsPrepared: false,
		IsRun:      false,
	}
	fmt.Println(l1, m1)
	Prepare(l1, m1)
	SetOn(l1, m1)
	fmt.Println(l1, m1)
	SetOff(l1, m1)
	fmt.Println(l1, m1)
}
