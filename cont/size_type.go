package cont

import "fmt"

// Width describes width of first code in the specified standard size code.
type Width string

// Height describes height of first code in the specified standard size code.
type Height string

// Length describes length of second code in the specified standard size code.
type Length string

// TypeInfo has information about the specified standard type.
type TypeInfo string

// GroupInfo has information about the specified type group.
type GroupInfo string

// IsLengthCode checks for correct format
func IsLengthCode(code string) error {
	return isOneUpperAlphanumericChar(code)
}

// IsHeightWidthCode checks if string is one upper case alphanumeric character.
func IsHeightWidthCode(code string) error {
	return isOneUpperAlphanumericChar(code)
}

// IsTypeCode checks if string is two upper case alphanumeric characters.
func IsTypeCode(code string) error {
	if len(code) != 2 {
		return NewErrContValidate(fmt.Sprintf("%s is not 2 characters long", code))
	}
	if !isUpperAlphanumeric(code) {
		return NewErrContValidate(
			fmt.Sprintf("%s is not 2 upper case alphanumeric characters", code))
	}
	return nil
}
