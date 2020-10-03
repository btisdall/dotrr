package config

import "testing"

func TestNewItem(t *testing.T) {
	item := NewItem("this", "that")
	if !(item.Key == "this" && item.Value == "that") {
		t.Errorf("NewItem did not contstruct the expected item")
	}
}

func TestMarshal(t *testing.T) {
	c := Map{
		"This":  "That",
		"Apple": "Orange",
		"hello": "goodbye",
	}

	expected := `Apple="Orange"
This="That"
hello="goodbye"`

	marshalled := Marshal(c)
	if marshalled != expected {
		t.Errorf("Marshalled did not equal expected:\nGot:\n%v\nWant:\n%v\n", marshalled, expected)
	}
}
