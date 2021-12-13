package main

import (
	"fmt"
	"syscall/js"

	"github.com/marcopolo/go-ipld-schema-playground/helper"
)

func main() {

	// var schemaContents string
	// document.getElementById("ipldSchema").value
	schemaContents := js.Global().Get("document").Call("getElementById", "ipldSchema").Get("value").String()

	// document.getElementById("ipldSchema").value
	ipldData := js.Global().Get("document").Call("getElementById", "ipldData").Get("value").String()

	typName, err := helper.Validate([]byte(schemaContents), []byte(ipldData))
	if err == nil {
		fmt.Println("Is valid")
		js.Global().Get("document").Call("getElementById", "validationResult").Set("textContent", "✅ Valid as "+typName)
		js.Global().Get("document").Call("getElementById", "validationResult").Get("classList").Call("add", "valid")
	} else {
		js.Global().Get("document").Call("getElementById", "validationResult").Get("classList").Call("remove", "valid")
		js.Global().Get("document").Call("getElementById", "validationResult").Set("textContent", "❌ invalid: "+err.Error())

	}
}
