package handset

import (
	"testing"
)

func TestParse(t *testing.T) {
	Parse(1320201)
	if len(sims) == 0 {
		t.Log("no data found.")
		t.Fail()
	}
}
