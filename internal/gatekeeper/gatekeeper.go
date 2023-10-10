package gatekeeper

import "reflect"

// function to validates the request and returns error if failed validation
type Validator func(req any) error

// function to sanitize the request inputs and returns error if any failed sanitization
type Sanitizer func(req any) error

var (
	validators map[reflect.Type]Validator
	sanitizers map[reflect.Type]Sanitizer
)

func init() {
	validators = make(map[reflect.Type]Validator)
	sanitizers = make(map[reflect.Type]Sanitizer)
}

// use provided validator to validate the request
func Validate(requestType any, validator Validator) {
	validators[reflect.TypeOf(requestType)] = validator
}

// use provided sanitizer to sanitize the request
func Sanitize(requestType any, sanitizer Sanitizer) {
	sanitizers[reflect.TypeOf(requestType)] = sanitizer
}
