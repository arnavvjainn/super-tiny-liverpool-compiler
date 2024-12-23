package main

import (
	"reflect"
	"testing"
)

const testInput = "(add TRENT (subtract SALAH VIRGIL))"

var testTokens = []token{
	{
		kind:  "paren",
		value: "(",
	},
	{
		kind:  "name",
		value: "add",
	},
	{
		kind:  "player",
		value: "TRENT",
	},
	{
		kind:  "paren",
		value: "(",
	},
	{
		kind:  "name",
		value: "subtract",
	},
	{
		kind:  "player",
		value: "SALAH",
	},
	{
		kind:  "player",
		value: "VIRGIL",
	},
	{
		kind:  "paren",
		value: ")",
	},
	{
		kind:  "paren",
		value: ")",
	},
}

var testAst = ast{
	kind: "Program",
	body: []node{
		{
			kind: "CallExpression",
			name: "add",
			params: []node{
				{
					kind:  "PlayerLiteral",
					value: "TRENT",
				},
				{
					kind: "CallExpression",
					name: "subtract",
					params: []node{
						{
							kind:  "PlayerLiteral",
							value: "SALAH",
						},
						{
							kind:  "PlayerLiteral",
							value: "VIRGIL",
						},
					},
				},
			},
		},
	},
}

func TestTokenizer(t *testing.T) {
	result := tokenizer(testInput)
	if !reflect.DeepEqual(result, testTokens) {
		t.Error("\nExpected:", testTokens, "\nGot:", result)
	}
}

func TestParser(t *testing.T) {
	result := parser(testTokens)
	if !reflect.DeepEqual(result, testAst) {
		t.Error("\nExpected:", testAst, "\nGot:", result)
	}
}
