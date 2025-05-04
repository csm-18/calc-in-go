package main

import (
	"calc/calc"
	"fmt"
)

const VERSION = "1.0.0"

func main() {
	fmt.Println("calc", VERSION)

	fmt.Println("\n ('q' for quit)")
	for {
		fmt.Print(">> ")
		//read input expression
		exp := readline()

		if exp == "q" {
			break
		} else if exp == "" {
			continue
		}

		calc.Calc(exp)

	}
}

func readline() string {
	var text = ""
	for {
		var temp rune
		fmt.Scanf("%c", &temp)
		if temp == '\n' {
			break
		}
		text += string(temp)
	}
	return text

}
