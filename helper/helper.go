package helper

import (
	"bytes"
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/node/bindnode"

	"github.com/hashicorp/go-multierror"
)

func Validate(schemaContents []byte, marshalledData []byte) (typName string, err error) {
	schema, err := ipld.LoadSchemaBytes(schemaContents)
	if err != nil {
		return "", err
	}

	var errs *multierror.Error
	for _, t := range schema.GetTypes() {
		fmt.Println("validating with type", t.Name())
		builder := bindnode.Prototype(nil, t).NewBuilder()
		err = dagjson.Decode(builder, bytes.NewReader(marshalledData))
		if err != nil {
			fmt.Println("err", t.Name(), err)
			errs = multierror.Append(errs, fmt.Errorf("validating data as %s: %w", t.Name(), err))
			continue
		} else {
			return t.Name(), nil
		}
	}

	return "", errs
}
