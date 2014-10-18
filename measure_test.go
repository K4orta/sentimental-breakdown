package sentimental

import (
	"fmt"
	"testing"
)

// Test the score of an empty string.
func TestMeasureEmpty(t *testing.T) {
	score := Measure("")
	if score != 0 {
		t.Error("Empty string score is not equal to 0")
	}
}

func TestMeasureSinglePositiveWord(t *testing.T) {
	score := Measure("fantastic")
	if score <= 0 {
		t.Error("\"Fantastic\" does not score positively")
	}
}

func TestMeasureSingleNegativeWord(t *testing.T) {
	score := Measure("obnoxious")
	if score >= 0 {
		t.Error("\"obnoxious\" does not score negatively")
	}
}

func TestMeasureMixed(t *testing.T) {
	score := Measure("Come out to the farm and eat my ass, idiot.")
	fmt.Println(score)
	if score == 0 {
		t.Error("\"obnoxious\" does not score negatively")
	}
}
