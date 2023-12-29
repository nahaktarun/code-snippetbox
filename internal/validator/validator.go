package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

// Define type of validator struct
type Validator struct {
	FieldErrors map[string]string
}

// valid() returns true if the fieldErrors map doesn't contains any entries
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

// AddFieldErrors adds an error message to the FieldErrors
func (v *Validator) AddFieldErrors(key, message string) {

	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}
	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// CheckField() adds an error message to the fieldErrors map
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldErrors(key, message)
	}
}

// NotBlank() returns true if a value a not an empty string
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// maxChars() returns true if a value contains no more than n characters.
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// PermittedValue() returns true if a value is in list of specific permitted values
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}
