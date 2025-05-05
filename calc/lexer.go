package calc

import (
	"fmt"
	"strings"
)

type TokenType int

const (
	NUM TokenType = iota
	LEFT_PAREN
	RIGHT_PAREN
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	REMAINDER
)

type Token struct {
	token_type TokenType
	value      string
}

func lexer(exp string) (error, []Token) {
	tokens := []Token{}

	x := 0
	for x < len(exp) {
		//check if it is a number
		if is_numeric(rune(exp[x])) {
			num := ""
			y := x
			for y < len(exp) && is_numeric(rune(exp[y])) {
				num += string(exp[y])
				y += 1
			}

			//check if it is a valid number
			if num[0] == '.' || num[len(num)-1] == '.' {
				return fmt.Errorf("invalid expression!"), nil
			}
			if strings.Contains(num, ".") {
				//check if it is a valid float
				if strings.Count(num, ".") > 1 {
					return fmt.Errorf("invalid expression!"), nil
				}
			}

			tokens = append(tokens, Token{NUM, num})
			x = y
			continue
		}

		switch exp[x] {
		case '(':
			tokens = append(tokens, Token{LEFT_PAREN, string(exp[x])})
		case ')':
			tokens = append(tokens, Token{RIGHT_PAREN, string(exp[x])})
		case '+':
			tokens = append(tokens, Token{PLUS, string(exp[x])})
		case '-':
			tokens = append(tokens, Token{MINUS, string(exp[x])})
		case '*':
			tokens = append(tokens, Token{MULTIPLY, string(exp[x])})
		case '/':
			tokens = append(tokens, Token{DIVIDE, string(exp[x])})
		case '%':
			tokens = append(tokens, Token{REMAINDER, string(exp[x])})

		}
		x += 1
	}

	return nil, tokens

}

func is_numeric(c rune) bool {
	numeric_symbols := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'}
	for _, ns := range numeric_symbols {
		if c == ns {
			return true
		}
	}
	return false
}
