package main

import "fmt"

func main() {
	var raw string = `아리랑\n아리랑\n아라리요`
	// raw := `아리랑\n아리랑\n아라리요`
	inter := "아리랑아리랑\n아리리요"

	fmt.Println(raw)
	fmt.Println()
	fmt.Println(inter + "1")

	fmt.Println("----------")
	println(raw)
	fmt.Println()
	println(inter + "1")
}
