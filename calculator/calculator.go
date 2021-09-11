package calculator

import "fmt"

type ICalculator interface {
	Add() int
	AddMulti() int
	Subtract() int
	Multiply() int
	MultiplyMulti() int
	Devide() int
}

type Calculator struct {
	A     int
	B     int
	Multi []int
}

func (c Calculator) Add() int {
	return c.A + c.B
}

func (c Calculator) AddMulti() int {
	i := 0
	for index := 0; index < len(c.Multi); index++ {
		i = i + c.Multi[index]
	}

	return i
}

func (c Calculator) Subtract() int {
	return c.A - c.B
}

func (c Calculator) Multiply() int {
	return c.A * c.B
}

func (c Calculator) MultiplyMulti() int {
	i := 1
	for index := 0; index < len(c.Multi); index++ {
		i = i * c.Multi[index]
	}

	return i
}

func (c Calculator) Devide() int {
	if c.B == 0 {
		panic("Cannot Devide by zero")
	}
	return c.A / c.B
}

func TryOutRace() {
	done := make(chan bool)
	m := make(map[string]string)
	m["name"] = "world"
	go func() {
		m["name"] = "data race"
		done <- true
	}()
	// comment below code to simulate race condition
	<-done

	fmt.Println("Hello,", m["name"])

	// uncomment below code to simulate race condition
	// <-done
}
