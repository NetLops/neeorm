package main

import (
	"fmt"
	"reflect"
)

type Account struct {
	Username string
	Password string
}

func main() {
	typ := reflect.Indirect(reflect.ValueOf(&Account{})).Type()
	fmt.Println(typ)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fmt.Println(field)
	}
}
