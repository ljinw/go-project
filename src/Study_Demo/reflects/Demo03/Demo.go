package main

import (
	"fmt"
	"reflect"
)

type mystruct struct {
	Name string
	Sex  int
	Age  int `json:"age"`
}

func main() {
	typeofmystruct := reflect.TypeOf(mystruct{})
	fmt.Println(typeofmystruct.Elem().Name())
	fmt.Println(typeofmystruct.Elem().Kind())
}
