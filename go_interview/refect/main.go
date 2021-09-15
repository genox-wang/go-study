package main

import (
	"fmt"
	"reflect"
)

type J struct {
	a string
	B string `json:"B"`
	C string `json:"C"`
	D string `json:"D" other:"E"`
}

func main() {
	j := new(J)
	t := reflect.TypeOf(j).Elem()
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("结构体内第%d个字段，Name是%v json是%v, other是%v\n", i, t.Field(i).Name, t.Field(i).Tag.Get("json"), t.Field(i).Tag.Get("other"))
	}
}
