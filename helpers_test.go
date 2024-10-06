package main

import (
	"testing"
)

func TestIsSubset(t *testing.T) {
	a := []string{"c++"};
	b := []string{"c++", "java", "asm"};
	t.Log(a, b);
	if !isSubset(a, b) {
		t.Error("a doesn contain all of b -> ", a, b);
	}

	a = []string{"c++", "6502"};
	b = []string{"c++", "java", "asm"};
	if isSubset(a, b) {
		t.Error("a doesn't contain all of b -> ", a, b);
	}
}

