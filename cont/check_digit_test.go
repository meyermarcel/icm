package cont

import (
	"testing"
)

func TestCalcCheckDigit(t *testing.T) {
	type args struct {
		ownerCode  string
		equipCatID string
		serialNum  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test ABC U 123456 0",
			args{"ABC", "U", "123456"},
			0,
		},
		{
			"Test NYK U 008685 2",
			args{"NYK", "U", "008685"},
			2,
		},
		{
			"Test NYK U 000000 0",
			args{"NYK", "U", "000000"},
			10,
		},
		{
			"Test CMA U 163912 (1)0",
			args{"CMA", "U", "163912"},
			10,
		},
		{
			"Test CMA U 169312 0",
			args{"CMA", "U", "169312"},
			0,
		},
		{
			"Test CSQ U 305438 3",
			args{"CSQ", "U", "305438"},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcCheckDigit(tt.args.ownerCode, tt.args.equipCatID, tt.args.serialNum); got != tt.want {
				t.Errorf("calcCheckDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
