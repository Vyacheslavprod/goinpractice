package main

import "testing"

func TestName(t *testing.T) {
	name := getName()
	if name != "World" {
		t.Errorf("Respone from getName is unexpected value")
	}
}