package cont

import (
	"github.com/meyermarcel/iso6346/equip_cat"
	"github.com/meyermarcel/iso6346/owner"
	"testing"
)

func TestCalcCheckDigit(t *testing.T) {
	type args struct {
		ownerCode  owner.Code
		equipCatId equip_cat.Id
		serialNum  SerialNum
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	  {"Test ABC U 123456 0",
		args{owner.NewCode("ABC"), equip_cat.NewIdU(), NewSerialNum(123456)},
	  0},
		{"Test MSK U 266654 2",
			args{owner.NewCode("MSK"), equip_cat.NewIdU(), NewSerialNum(266654)},
			2},
		{"Test NYK U 008685 2",
			args{owner.NewCode("NYK"), equip_cat.NewIdU(), NewSerialNum(8685)},
			2},
		{"Test NYK U 000000 0",
			args{owner.NewCode("NYK"), equip_cat.NewIdU(), NewSerialNum(0)},
			0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcCheckDigit(tt.args.ownerCode, tt.args.equipCatId, tt.args.serialNum); got != tt.want {
				t.Errorf("CalcCheckDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
