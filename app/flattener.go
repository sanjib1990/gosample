package app

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sanjib1990/goflat"
)

type Address struct {
	HouseNo    string `json:"house_no"`
	StreetName string `json:"street_name"`
	AreaName   string `json:"area_name"`
	City       string `json:"city"`
	State      string `json:"state"`
	Pincode    string `json:"pincode"`
}

func FlatenUser(user *User) map[string]interface{} {
	var in = map[string]interface{}{}

	bytarr, _ := json.Marshal(user)

	_ = json.Unmarshal(bytarr, &in)

	out, err := goflat.Flatten(in, nil)

	if err != nil {
		log.Fatal(err)
	}

	return out
}

func Unflatten(val map[string]interface{}) *User {
	op, err := goflat.Unflatten(val, nil)
	if err != nil {
		log.Fatal(err)
	}

	var user User
	bytarr, _ := json.Marshal(op)
	_ = json.Unmarshal(bytarr, &user)

	return &user
}

func flatten(input map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range input {
		switch v := v.(type) {
		case map[string]interface{}:
			for kk, vv := range v {
				result[k+"."+kk] = vv
			}
		case string:
			result[k] = v
		default:
			result[k] = fmt.Sprintf("%v", v)
		}
	}
	return result
}

func main() {
	input := map[string]interface{}{"a": "b", "c": map[string]interface{}{"d": "e"}}
	output := flatten(input)
	fmt.Println(output)
}
