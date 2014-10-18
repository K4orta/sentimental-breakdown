package sentimental

import (
	"github.com/k4orta/sentimental/afinn"
)

func Measure(raw string) int {
	var lookup *map[string]int = afinn.GetAFINN()
	tokens := tokenize(raw)
	score := 0
	for _, token := range tokens {
		if points, exists := (*lookup)[token]; exists {
			score += points
		}
	}

	return score
}
