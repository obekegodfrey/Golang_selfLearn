package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Error("Add(2, 3) = %d; want 5", result)
	}
}
