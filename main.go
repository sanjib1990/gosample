package main

import (
    "fmt"
    "sample/calculator"
)


func main() {
    fmt.Print("Hello World!! \n")

    cal := calculator.Calculator{A: 4, B: 5}
    fmt.Println(cal.Add())
    fmt.Println(cal.Subtract())
    fmt.Println(cal.Multiply())
    calculator.TryOutRace()
}