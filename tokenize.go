package sentimental

import (
	"regexp"
	"strings"
)

func tokenize(text string) []string {
	nonAlpha, _ := regexp.Compile("[^a-zA-Z- ]+")
	extraSpace, _ := regexp.Compile(" {2,}")
	transformed := extraSpace.ReplaceAllString(nonAlpha.ReplaceAllString(text, ""), "")
	return strings.Split(strings.ToLower(transformed), " ")
}
