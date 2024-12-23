package main

import (
	"fmt"
	"log"
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

/**
* ============================================================================
*                                 ヽ/❀o ل͜ o\ﾉ
*                                THE PARSER!!!
* ============================================================================
 */

type node struct {
	kind       string
	value      string
	name       string
	callee     *node
	expression *node
	body       []node
	params     []node
	arguments  *[]node
	context    *[]node
}

type ast node

var pc int

var pt []token

func parser(tokens []token) ast {
	pc = 0
	pt = tokens

	ast := ast{
		kind: "Program",
		body: []node{},
	}

	for pc < len(pt) {
		ast.body = append(ast.body, walk())
	}

	return ast
}

// (add TRENT (subtract SALAH VIRGIL))
func walk() node {

	token := pt[pc]

	// token is player YNWA
	if token.kind == "player" {
		pc++
		return node{
			kind:  "PlayerLiteral",
			value: token.value,
		}
	}

	// start of expression
	if token.kind == "paren" && token.value == "(" {
		pc++
		token = pt[pc]

		n := node{
			kind:   "CallExpression",
			name:   token.value,
			params: []node{},
		}

		pc++
		token = pt[pc]

		// recursive call to walk() closing parenthesis
		for token.kind != "paren" || (token.kind == "paren" && token.value != ")") {
			n.params = append(n.params, walk())
			token = pt[pc]
		}
		pc++
		return n
	}

	log.Fatal(token.kind)
	return node{}
}

func compiler(input string) ast {
	tokens := tokenizer(input)
	ast := parser(tokens)
	return ast
}

func main() {
	program := "(add TRENT (subtract SALAH VIRGIL))"
	out := compiler(program)
	fmt.Println(out)
}
