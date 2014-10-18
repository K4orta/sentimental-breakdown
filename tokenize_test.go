package sentimental

import (
	"testing"
)

func TestConvertToLower(t *testing.T) {
	tokens := tokenize("AAA bbb CCC")
	if tokens[0] != "aaa" {
		t.Error("Did not covert tokens to lowercase")
	}
}

func TestRemoveNoAlpha(t *testing.T) {
	tokens := tokenize(",,\t23421abc\n&&&3 32;:;..")
	if tokens[0] != "abc" {
		t.Log(tokens[0])
		t.Error("Did not remove non-alphabet characters")
	}
}
