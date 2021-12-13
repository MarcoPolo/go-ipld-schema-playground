package helper

import (
	"fmt"
	"testing"
)

// simple test
func TestValidate(t *testing.T) {
	data := `{ "baz": 1.23 }`
	schemaContents := `type Foo struct {
	baz Float
}`

	typeName, err := Validate([]byte(schemaContents), []byte(data))
	if err != nil {
		t.Error(err)
	}

	fmt.Println(typeName)
}
