package back_end

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var (
	firstNameRegex = regexp.MustCompile(`[\p{L} '-]*\p{L}$`)
	lastNameRegex  = regexp.MustCompile(`[\p{L}'-]*\p{L}$`)
	suffixRegex    = regexp.MustCompile(`[\p{L}.]*$`)
)

type ValidationResponse struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func commonChecks(field, val string) error {
	val = strings.TrimSpace(val)
	if len(strings.TrimSpace(val)) < 1 {
		return &ValidationError{Field: field, Message: "Cannot be empty."}
	}
	if !unicode.IsLetter(rune(val[0])) {
		return &ValidationError{Field: field, Message: "Must begin with a letter."}
	}
	return nil
}

func ValidFirstName(first string) error {
	field := "First Name"
	if err := commonChecks(field, first); err != nil {
		return err
	}
	if !firstNameRegex.MatchString(first) {
		return &ValidationError{Field: field, Message: "Only letters, hyphens, apostrophes, and spaces are allowed."}
	}
	return nil
}

func ValidLastName(last string) error {
	field := "Last Name"
	if err := commonChecks(last, last); err != nil {
		return err
	}
	if !lastNameRegex.MatchString(last) {
		return &ValidationError{Field: field, Message: "Only letters, hyphens, and apostrophes are allowed."}
	}
	return nil
}

func ValidSuffix(suffix string) error {
	suffix = strings.TrimSpace(suffix)
	field := "Suffix"
	if len(suffix) < 1 {
		return nil
	}
	if !suffixRegex.MatchString(suffix) {
		return &ValidationError{Field: field, Message: "Only letters and periods are allowed."}
	}
	return nil
}
