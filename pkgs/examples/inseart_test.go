package examples

import "testing"

func TestEgInseartOneWithStruct(t *testing.T) {
	if err := EgInseartOneWithStruct(); err != nil {
		t.Error(err)
	}
}

func TestEgInseartOneWithBson(t *testing.T) {
	if err := EgInseartOneWithBson(); err != nil {
		t.Error(err)
	}
}
