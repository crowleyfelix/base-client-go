package errors

import (
	"fmt"
	"reflect"
)

type Callout struct {
	*errorWrapper
}

func NewCallout(messages ...string) Error {
	return &Callout{
		new("Failed on requesting.", messages...),
	}
}

type Serializing struct {
	*errorWrapper
}

func NewSerializing(messages ...string) Error {
	return &Serializing{
		new("Failed on serializing data.", messages...),
	}
}

type InvalidType struct {
	*errorWrapper
}

func NewInvalidType(actual, expected reflect.Type) Error {
	return &InvalidType{
		new(fmt.Sprintf("Actual type %s. Expected type %s.", actual, expected)),
	}
}
