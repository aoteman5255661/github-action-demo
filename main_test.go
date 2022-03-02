package main

import "testing"

// 会miao
func TestAbs(t *testing.T) {
	saying := Cat()
	if saying != "Miao~~~~" {
		t.Errorf("Cat say %s", saying)
	}
}
