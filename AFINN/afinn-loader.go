package afinn

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"
)

// Load the AFINN file once
var cachedAFINN *map[string]int

func init() {
	var once sync.Once
	once.Do(func() { cachedAFINN = loadAFINN("afinn/AFINN-111.txt") })
}

func GetAFINN() *map[string]int {
	return cachedAFINN
}

func loadAFINN(path string) *map[string]int {
	dict := map[string]int{}
	// Open file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Could not open AFINN file")
		return nil
	}

	// Split on newline
	lines := strings.Split(string(data), "\n")

	// Create dictionary
	for _, v := range lines {
		split := strings.Split(v, "\t")
		if score, err := strconv.Atoi(split[1]); err == nil {
			dict[split[0]] = score
		}
	}

	return &dict
}
