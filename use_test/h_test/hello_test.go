// Тест для hello
package h_test

import "testing"

func TestHello(t *testing.T) {
	if v := Hello(); v != "hello" {
		t.Errorf("Expected 'hello', but got '%s'", v)
	}
}