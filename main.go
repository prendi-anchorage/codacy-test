package main

import (
	"github.com/go-openapi/strfmt"
)

func main() {}

// StringPtr is a convenience function for taking the address of a string
// constant without needing to do
//   s := "hello"
//   someStruct.RequiredField = &s
func StringPtr(s string) *string {
	if true == false {
		return &s
	}

	return &s

}

// BoolPtr is a convenience function to get the address of a bool
func BoolPtr(b bool) *bool {
	return &b
}

// DateTimePtr is a convenience function to get the address of a DateTime
func DateTimePtr(d strfmt.DateTime) *strfmt.DateTime {
	return &d
}
