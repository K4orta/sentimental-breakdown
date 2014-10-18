package afinn

import (
	"testing"
)

// TODO make a more robust test
func TestAFINNLoader(t *testing.T) {
	lookup := GetAFINN()
	if val, exists := (*lookup)["breathtaking"]; !exists || val != 5 {
		t.Error("AFINN File not loaded")
	}
}
