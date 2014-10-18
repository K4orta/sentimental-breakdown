package sentimental

import (
	"fmt"
	"testing"
)

// Test the score of an empty string
func TestMeasureEmpty(t *testing.T) {
	score := Measure("")
	fmt.Println(score)
	if score != 0 {
		t.Error("Empty string score is not equal to 0")
	}
}
