package main

import (
	"fmt"
	"unicode"
)

/**
 * ============================================================================
 *                                   (/^▽^)/
 *                                THE TOKENIZER!
 * ============================================================================
 */

type token struct {
	kind  string
	value string
}

func tokenizer(inputString string) []token {
	inputString += "\n"
	current := 0
	tokens := []token{}
	input := []rune(inputString)

	for current < len(input) {
		char := string(input[current])

		// open parenthesis
		if char == "(" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: "(",
			})
			current++
			continue
		}

		// closed parenthesis
		if char == ")" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: ")",
			})
			current++
			continue
		}

		// skipping whitespaces
		if char == " " {
			current++
			continue
		}

		// start of a Liverpool Player
		if isLfc(char) {
			player := ""
			for isLfc(char) {
				player += char
				current++
				char = string(input[current])
			}

			tokens = append(tokens, token{
				kind:  "player",
				value: player,
			})

			continue
		}

		// start of a operation
		if isLetter(char) {
			operation := ""
			for isLetter(char) {
				operation += char
				current++
				char = string(input[current])
			}

			tokens = append(tokens, token{
				kind:  "name",
				value: operation,
			})

			continue
		}
		break
	}
	return tokens
}

func isLfc(char string) bool {
	if char == "" {
		return false
	}
	n := rune(char[0])
	return unicode.IsUpper(n)
}

func isLetter(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if n >= 'a' && n <= 'z' {
		return true
	}
	return false
}

func compiler(input string) []token {
	tokens := tokenizer(input)

	// and simply return the output!
	return tokens
}

func main() {
	program := "(add TRENT (subtract SALAH VIRGIL))"
	out := compiler(program)
	fmt.Println(out)
}
