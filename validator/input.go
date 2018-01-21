package validator

type Input struct {
	value       string
	validLength int
}

func (in Input) IsComplete() bool {
	return len(in.value) == in.validLength
}

func (in Input) Value() string {
	return in.value
}

func (in Input) ValidLength() int {
	return in.validLength
}

func NewInput(value string, validLength int) Input {
	return Input{value: value, validLength: validLength}
}
