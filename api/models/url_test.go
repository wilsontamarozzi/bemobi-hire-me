package models

import (
	"testing"
)

var URL_TEST = URL{
	Address: "http://example.com.br",
	Alias:   "wilson",
}

func TestURLVazia(t *testing.T) {
	url := URL{}

	if !url.IsEmpty() {
		t.Errorf("Esperado %t", true)
		t.Errorf("Recebido %t", false)
	}
}

func TestAliasVazio(t *testing.T) {
	url := URL_TEST
	url.Alias = ""

	if !url.AliasIsEmpty() {
		t.Errorf("Esperado %t", true)
		t.Errorf("Recebido %t", false)
	}
}

func TestValidateEnderecoVazio(t *testing.T) {
	url := URL_TEST
	url.Address = ""

	if err := url.Validate(); len(err) <= 0 {
		t.Errorf("Esperado %d", 1)
		t.Errorf("Recebido %d", len(err))
	}
}
