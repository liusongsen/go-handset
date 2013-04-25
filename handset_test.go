package handset

import (
	"testing"
)

func TestEven(t *testing.T) {
	if !Even(2) {
		t.Log("2 should be even!")
		t.FailNow()
	}
}

func TestOdd(t *testing.T) {
	if !Even(3) {
		t.Log("3 should be even!")
		t.FailNow()
	}
}
