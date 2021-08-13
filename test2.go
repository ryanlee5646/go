package main

import (
	"fmt"
	"reflect"
)


func main() {
	k := 10
	p := &k
	fmt.Println(reflect.TypeOf(p))
	// fmt.Println(p)

	if k := 1; k == 1 {
		fmt.Println(1)
	}else {
		fmt.Println(2)
	}

}