package main

import (
	"fmt"
	"reflect"
)

type myStruct struct {
	F float64 //capital
	B bool
	I int
}

func main() {
	var ms = &myStruct{0.1, false, 1}
	val := reflect.ValueOf(ms).Elem()
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		switch f.Interface().(type) {
		case int:
			fmt.Println("case int")
		case float64:
			fmt.Println("case float64")
		case bool:
			fmt.Println("case bool")
		}

	}
}
