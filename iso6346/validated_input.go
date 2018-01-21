package iso6346

type ValidatedInput struct {
	OwnerCode            Input
	EquipmentCategoryId  Input
	SerialNumber         Input
	CheckDigit           Input
	IsValidCheckDigit    bool
	CalculatedCheckDigit int
}

type Input struct {
	Value      string
	IsComplete bool
}

func NewValidatedInput(ownerCodeInput string,
	equipmentCategoryIdInput string,
	serialNumberInput string,
	checkDigitInput string) ValidatedInput {

	vi := ValidatedInput{OwnerCode: Input{ownerCodeInput, false},
		EquipmentCategoryId: Input{equipmentCategoryIdInput, false},
		SerialNumber: Input{serialNumberInput, false},
		CheckDigit: Input{checkDigitInput, false}}

	ownerCode, err := NewOwnerCode(vi.OwnerCode.Value)

	if err == nil {
		vi.OwnerCode.IsComplete = true
	}

	equipmentCategoryId, err := NewEquipmentCategoryId(vi.EquipmentCategoryId.Value)

	if err == nil {
		vi.EquipmentCategoryId.IsComplete = true
	}

	serialNumber, err := NewSerialNumber(vi.SerialNumber.Value)

	if err == nil {
		vi.SerialNumber.IsComplete = true
	}

	checkDigit, err := NewCheckDigit(vi.CheckDigit.Value)

	if err == nil {
		vi.CheckDigit.IsComplete = true
	}

	if !vi.IsCheckDigitCalculable() {
		return vi
	}

	vi.CalculatedCheckDigit = CalculateCheckDigit(ownerCode,
		equipmentCategoryId,
		serialNumber)

	if err != nil {
		return vi
	}

	vi.IsValidCheckDigit = vi.CalculatedCheckDigit == checkDigit.value

	return vi
}

func (vi ValidatedInput) IsCheckDigitCalculable() bool {
	return vi.OwnerCode.IsComplete && vi.EquipmentCategoryId.IsComplete && vi.SerialNumber.IsComplete
}
