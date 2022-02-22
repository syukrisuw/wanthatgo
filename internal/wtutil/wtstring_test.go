package wtutil

import "testing"

func TestStringFunction(t *testing.T) {

	res := StringFunction()
	if !res {
		t.Fatalf(`Testing Failed`)
	}
}
