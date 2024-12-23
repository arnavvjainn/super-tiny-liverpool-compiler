package main

import (
	"reflect"
	"testing"
)

const testInput = "(add TRENT (subtract SALAH VIRGIL))"

var testTokens = []token{
	token{
		kind:  "paren",
		value: "(",
	},
	token{
		kind:  "name",
		value: "add",
	},
	token{
		kind:  "player",
		value: "TRENT",
	},
	token{
		kind:  "paren",
		value: "(",
	},
	token{
		kind:  "name",
		value: "subtract",
	},
	token{
		kind:  "player",
		value: "SALAH",
	},
	token{
		kind:  "player",
		value: "VIRGIL",
	},
	token{
		kind:  "paren",
		value: ")",
	},
	token{
		kind:  "paren",
		value: ")",
	},
}

func TestTokenizer(t *testing.T) {
	result := tokenizer(testInput)
	if !reflect.DeepEqual(result, testTokens) {
		t.Error("\nExpected:", testTokens, "\nGot:", result)
	}
}
