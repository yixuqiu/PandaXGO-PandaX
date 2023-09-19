package tool

import "testing"

func TestToCamelCase(t *testing.T) {
	camelCase := ToCamelCase("hello_world")
	t.Log(camelCase)
}

func TestGenerateID(t *testing.T) {
	id := GenerateID()
	t.Log(id)
}
