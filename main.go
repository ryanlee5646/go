package main

import (
	"fmt"
	"strings"
)


// 함수 리턴값 지정
func lenAndUpper (name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	fmt.Println("1")
	return 
}

// 함수 다중 인자값
func repeatMe(words ...string)([]string){
	return words
}

// for문
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

//if 문

func canIDrink(age int) bool {
	if koreanAge := age + 2 ; koreanAge < 18 {
		return false
	} 
	return true
	
}

//Swich 문

func canIDrink2(age int) bool {
	switch age {
	case 10 :
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	result := superAdd(1,2,3,4,5,6)
	fmt.Println(result)
	fmt.Println(canIDrink(16))
	fmt.Println(canIDrink2(18))
	// a,b := lenAndUpper("hello world")
	// fmt.Println(a,b)
	// array := repeatMe("nico","lynn", "dal", "marl")
}