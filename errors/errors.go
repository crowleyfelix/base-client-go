package errors

import (
	"fmt"
	"reflect"
)

type Callout struct {
	*errorWrapper
}

func NewCallout(messages ...string) *Callout {
	return &Callout{
		new("Failed on requesting.", messages...),
	}
}

type Serializing struct {
	*errorWrapper
}

func NewSerializing(messages ...string) *Serializing {
	return &Serializing{
		new("Failed on serializing data.", messages...),
	}
}

type InvalidType struct {
	*errorWrapper
}

func NewInvalidType(actual, expected reflect.Type) *InvalidType {
	return &InvalidType{
		new(fmt.Sprintf("Actual type %s. Expected type %s.", actual, expected)),
	}
}
