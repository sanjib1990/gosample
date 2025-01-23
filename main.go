package main

import (
	"encoding/json"
	"fmt"
)

type Flattener struct{}

func (f *Flattener) FlattenMap(input map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})

	for key, value := range input {
		if subMap, ok := value.(map[string]interface{}); ok {
			subOutput := f.FlattenMap(subMap)
			for subKey, subValue := range subOutput {
				newKey := fmt.Sprintf("%s.%s", key, subKey)
				output[newKey] = subValue
			}
		} else if subSlice, ok := value.([]interface{}); ok {
			for i, subValue := range subSlice {
				newKey := fmt.Sprintf("%s.%d", key, i)
				output[newKey] = subValue
			}
		} else {
			output[key] = value
		}
	}

	return output
}

func (f *Flattener) FlattenStruct(input interface{}) map[string]interface{} {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}

	var inputMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &inputMap)
	if err != nil {
		panic(err)
	}

	return f.FlattenMap(inputMap)
}

func amain() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	chunksize := 3
	totalLength := len(a)
	totalChunks := totalLength / chunksize
	for i := 0; i <= totalChunks; i++ {
		start := chunksize * i
		end := start + chunksize
		if end > totalLength {
			end = totalLength
		}

		if start == end {
			break
		}
		b := a[start:end]
		fmt.Println(b)
	}
}
