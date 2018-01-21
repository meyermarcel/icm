package owner

import (
	"iso6346/parser"
	"iso6346/validator"
)

type CodeOptEquipCatId struct {
	ownerCode  validator.Input
	equipCatId validator.Input
}

func (oc CodeOptEquipCatId) IsValid() bool {
	return oc.ownerCode.IsComplete() && oc.equipCatId.IsComplete()
}

func (oc CodeOptEquipCatId) OwnerCode() validator.Input {
	return oc.ownerCode
}

func (oc CodeOptEquipCatId) EquipCatId() validator.Input {
	return oc.equipCatId
}

func Validate(pi parser.Input) CodeOptEquipCatId {

	return CodeOptEquipCatId{
		validator.NewInput(pi.GetMatch(0, 3), 3),
		validator.NewInput(pi.GetMatchSingle( 3), 1)}
}
