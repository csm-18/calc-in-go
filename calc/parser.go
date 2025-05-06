package calc

import (
	"errors"
	"math"
	"strconv"
)

func multiply_divide_remainder(tokens []Token) (error, []Token) {
	x := 0
	for x < len(tokens) {
		if len(tokens) == 1 && tokens[x].token_type == NUM {
			return nil, tokens
		} else if x != 0 && x+1 < len(tokens) && (tokens[x].token_type == MULTIPLY || tokens[x].token_type == DIVIDE || tokens[x].token_type == REMAINDER) {
			left_operand, err := strconv.ParseFloat(tokens[x-1].value, 64)
			if err != nil {
				return err, nil
			}
			right_operand, err := strconv.ParseFloat(tokens[x+1].value, 64)
			if err != nil {
				return err, nil
			}

			result := 0.0
			if tokens[x].token_type == MULTIPLY {
				result = left_operand * right_operand

			} else if tokens[x].token_type == DIVIDE {
				result = left_operand / right_operand
			} else if tokens[x].token_type == REMAINDER {
				result = math.Mod(left_operand, right_operand)
			}

			result_token := Token{NUM, strconv.FormatFloat(result, 'f', -1, 64)}
			temp_tokens := make([]Token, 0, len(tokens))
			temp_tokens = append(temp_tokens, tokens[:x-1]...)
			temp_tokens = append(temp_tokens, result_token)
			if x+2 < len(tokens) {
				temp_tokens = append(temp_tokens, tokens[x+2:]...)
			}
			tokens = temp_tokens
			x -= 1
			continue
		} else if tokens[x].token_type == NUM {
			x += 1
			continue
		} else {
			return errors.New("invalid expression!"), nil
		}

	}

	if len(tokens) == 1 && tokens[x].token_type == NUM {
		return nil, tokens
	} else {
		return errors.New("invalid expression!"), nil
	}

}
