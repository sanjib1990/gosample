package main

import (
	"fmt"
	"sync"
)

//var pizzaNum = 0
//var pizzaName = ""
var wg sync.WaitGroup

func main() {
	//circle := Circle{2}
	//rect := Rectangle{1,1}
	//
	//fmt.Println(getArea(circle))
	//fmt.Println(getArea(rect))
	//
	//fmt.Println("Done!!")
	//
	//sampleString := "Hello World"
	//fmt.Println(strings.Contains(sampleString, "llo"))
	//fmt.Println(strings.Index(sampleString, "llo"))

	/**
	HTTP Server
	*/
	//http.HandleFunc("/", handler)
	//http.HandleFunc("/hi", handler2)
	//_ = http.ListenAndServe(":8080", nil)
	//stringChan := make(chan string)
	//for i:= 0; i < 3; i++ {
	//	go makeDough(stringChan)
	//	go addSause(stringChan)
	//	go addToppings(stringChan)
	//	time.Sleep(time.Millisecond * 5000)
	//}

	// go routine
	//wg.Add(1)
	//go say("hey")
	//wg.Add(1)
	//go say("there")
	//wg.Wait()

	// channels
	defer cleanUp()
	fooval := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooval, i)
	}
	wg.Wait()
	close(fooval)
	for item := range fooval {
		fmt.Println(item)
	}
}
func foo(c chan  int, val int)  {
	defer wg.Done()
	defer cleanUp()
	c <- val * 5
}

func cleanUp() {
	if r := recover(); r != nil {
		fmt.Println("Recovered: ", r)
	}
}

//func say(s string) {
//	defer wg.Done()
//	defer cleanUp()
//	for i := 0; i < 3; i++ {
//		time.Sleep(time.Millisecond*100)
//		fmt.Println(s, i)
//		if i == 2 {
//			panic("OMG!!")
//		}
//	}
//}
//
//func makeDough(stringChan chan string) {
//	pizzaNum++
//	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
//
//	fmt.Println("Make piz and send", pizzaName)
//
//	stringChan <- pizzaName
//	//time.Sleep(time.Millisecond * 10)
//}
//
//func addSause(stringChan chan string) {
//	pizza := <-stringChan
//	fmt.Println("Add Sauce and send ", pizza, "for toppings")
//
//	stringChan <- pizzaName
//	//time.Sleep(time.Millisecond * 10)
//}
//
//func addToppings(stringChan chan string) {
//	pizza := <-stringChan
//	fmt.Println("Add Topping send ", pizza, "for Delivery")
//
//	//time.Sleep(time.Millisecond * 10)
//}
//
//func count(id int) {
//	// go routine
//	for i := 0; i < 10; i++ {
//		fmt.Println(id, ":", i)
//		time.Sleep(time.Millisecond * 1000)
//	}
//}
//
//func handler(w http.ResponseWriter, r *http.Request) {
//	_, _ = fmt.Fprintf(w, "Hellp \n")
//}
//
//func handler2(w http.ResponseWriter, r *http.Request) {
//	_, _ = fmt.Fprintf(w, "HellThere whats up \n")
//}
//
//type shape interface {
//	area() float64
//}
//
//type Rectangle struct {
//	length  float64
//	breadth float64
//}
//
//func (r Rectangle) area() float64 {
//	return r.length * r.breadth
//}
//
//type Circle struct {
//	radius float64
//}
//
//func (c Circle) area() float64 {
//	return math.Pi * math.Pow(c.radius, 2)
//}
//
//func getArea(r shape) float64 {
//	return r.area()
//}
