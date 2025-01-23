package main

import (
	"encoding/json"
	"fmt"
)

var data = "{\"1\": \"value1\", \"3\": \"value3\", \"2\": \"value2\"}"

// op : []{data: val1}, {data: val2}

type OP struct {
	Data string `json:"data"`
}

type Params struct {
	Data string `json:"data"`
}

var (
	queue = make(chan int, 2)
)

type AStruct struct {
	Name string `json:"name" mapstructure:"name"`
	Age  int    `json:"age" mapstructure:"age"`
}

type BStruct struct {
	Name string `json:"name" mapstructure:"name"`
	Age  int    `json:"age" mapstructure:"age"`
}

type CStruct struct {
	Name string `json:"name" mapstructure:"name"`
	Age  int    `json:"age" mapstructure:"age"`
}

type MainStruct struct {
	A AStruct                `json:"a" mapstructure:"a"`
	B BStruct                `json:"b" mapstructure:"b"`
	C CStruct                `json:"c" mapstructure:"c"`
	D map[string]interface{} `json:"d" mapstructure:"d"`
}

func main() {
	//res := `{"a":{"name":"ajksdhaskd asd asda","age":11},"b":{"name":"ajksdhaskd asd 123 asda","age":12},"c":{"name":"ajksdhaskd asd 324 asda","age":13},"d":{"name":"ajksdhaskd asd 312224 asda","age":14}}`
	res := `{"a":{"name":["ajksdhaskd asd asda"],"age":11},"b":{"name":"ajksdhaskd asd 123 asda","age":12},"c":{"name":"ajksdhaskd asd 324 asda","age":13},"d":{"name":"ajksdhaskd asd 312224 asda","age":14}}`
	obj := MainStruct{}
	er := json.Unmarshal([]byte(res), &obj)
	if er != nil {
		fmt.Println(er)
		return
	}

	fmt.Println(obj)
}
