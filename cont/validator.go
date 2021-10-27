package cont

import "fmt"

// ErrContValidate is an error for validation of a container number part.
type ErrContValidate struct {
	message string
}

// NewErrContValidate returns a new ErrContValidate.
func NewErrContValidate(message string) error {
	return &ErrContValidate{
		message: message,
	}
}

func (e *ErrContValidate) Error() string {
	return e.message
}

func isOneUpperAlphanumericChar(code string) error {
	if len(code) != 1 {
		return NewErrContValidate(fmt.Sprintf("%s is not 1 digit", code))
	}
	if !isUpperAlphanumeric(code) {
		return NewErrContValidate(
			fmt.Sprintf("%s is not 1 upper case alphanumeric character", code))
	}
	return nil

}

func isUpperAlphanumeric(s string) bool {
	for _, r := range s {
		if !isUpperLetter(string(r)) && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

func isUpperLetter(s string) bool {
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}
