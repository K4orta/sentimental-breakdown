package sentimental

import (
	"github.com/k4orta/sentimental/afinn"
)

func Measure(raw string) int {
	lookup := afinn.GetAFINN()
	score := 0
	for _, token := range tokenize(raw) {
		if points, exists := (*lookup)[token]; exists {
			score += points
		}
	}

	return score
}

type Stat struct {
	Score       int
	Comparative float32
	Tokens      []string
	WordCount   int
	Positive    []string
	Negative    []string
}

// A version of the measure function that returns more data
func FullMeasure(raw string) Stat {
	lookup := afinn.GetAFINN()
	result := Stat{}
	tokens := tokenize(raw)
	for _, token := range tokens {
		if points, exists := (*lookup)[token]; exists {
			if points > 0 {
				result.Positive = append(result.Positive, token)
			} else if points < 0 {
				result.Negative = append(result.Negative, token)
			}
			result.Score += points
			result.Tokens = append(result.Tokens, token)
		}
	}
	result.WordCount = len(tokens)

	// Make sure we don't divide by zero
	if len(result.Tokens) > 0 {
		result.Comparative = float32(result.Score) / float32(len(result.Tokens))
	}
	return result
}
