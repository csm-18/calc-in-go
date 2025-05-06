package calc

import (
	"fmt"
	"strings"
)

func Calc(exp string) {
	//remove leading and trailing spaces
	exp = strings.TrimSpace(exp)

	//check for invalid symbols in expression
	if !valid_symbols(exp) {
		fmt.Println("invalid expression!")
		return
	}

	//get tokens from expression
	err, tokens := lexer(exp)
	if err != nil {
		fmt.Println(err)
		return
	}

	//debug print tokens
	err, tokens = multiply_divide_remainder(tokens)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Printf("tokens: %v\n", tokens)
}

func valid_symbols(exp string) bool {
	var valid_symbols = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', '+', '-', '*', '/', '%', '(', ')'}

	for _, e := range exp {
		valid := false
		for _, vs := range valid_symbols {
			if e == vs {
				valid = true
				break
			}
		}

		if valid {
			continue
		} else {
			return false
		}
	}

	return true
}
