package sentimental

import (
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
	if score == 0 {
		t.Error("Mixed tokens test failed")
	}
}

func TestFullMeasureNegative(t *testing.T) {
	score := FullMeasure("I'm real mad about the awful churro I just ate.")
	if score.Score == 0 {
		t.Error("Negative FullMeasure test failed")
	}
}

func TestFullMeasureNeutral(t *testing.T) {
	score := FullMeasure("I opened the .XFL file and was thrust into an incredible world of high octane hardcore football for all the real daredevils out there.")
	if score.Comparative != 0 {
		t.Error("Neutral phrase FullMeasure test failed")
	}
}
