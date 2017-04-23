package helpers

import (
	"testing"
)

var arrayTestCase = []string{"a", "b", "c"}

func TestSizeString(t *testing.T) {
	str := RandomString(2)
	if len(str) != 2 {
		t.Errorf("Esperado %s", 2)
		t.Errorf("Recebido %s", len(str))
	}
}

func TestContainItemInArray(t *testing.T) {
	item := "a"
	if !Contains(arrayTestCase, item) {
		t.Errorf("Esperado %s", true)
		t.Errorf("Recebido %s", false)
	}
}

func TestDoNotContainItemInArray(t *testing.T) {
	item := "d"
	if Contains(arrayTestCase, item) {
		t.Errorf("Esperado %s", false)
		t.Errorf("Recebido %s", true)
	}
}
