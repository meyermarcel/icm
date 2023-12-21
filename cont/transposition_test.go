package cont

import (
	"reflect"
	"testing"
)

func TestCheckTransposition(t *testing.T) {
	type args struct {
		ownerCode  string
		equipCatID string
		serialNum  string
		checkDigit int
	}
	tests := []struct {
		name string
		args args
		want []Number
	}{
		{
			name: "Test CMA U 163912 (1)0",
			args: args{
				ownerCode:  "CMA",
				equipCatID: "U",
				serialNum:  "163912",
				checkDigit: 10,
			},
			want: []Number{
				{"CMA", "U", "169312", 0},
				{"CMA", "U", "163192", 0},
			},
		},
		{
			name: "Test RCB U 001130 0",
			args: args{
				ownerCode:  "RCB",
				equipCatID: "U",
				serialNum:  "001130",
				checkDigit: 0,
			},
			want: []Number{
				{"RCB", "U", "010130", 0},
			},
		},
		{
			name: "Test WSL U 801743 (1)0",
			args: args{
				ownerCode:  "WSL",
				equipCatID: "U",
				serialNum:  "801743",
				checkDigit: 10,
			},
			want: []Number{
				{"WSL", "U", "810743", 0},
				{"WSL", "U", "807143", 0},
				{"WSL", "U", "801740", 3},
			},
		},
		{
			name: "Test APL U 689473 (1)0",
			args: args{
				ownerCode:  "APL",
				equipCatID: "U",
				serialNum:  "689473",
				checkDigit: 10,
			},
			want: []Number{
				{"APL", "U", "869473", 0},
				{"APL", "U", "698473", 0},
				{"APL", "U", "684973", 0},
				{"APL", "U", "689743", 0},
				{"APL", "U", "689437", 0},
				{"APL", "U", "689470", 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckTransposition(tt.args.ownerCode, tt.args.equipCatID, tt.args.serialNum, tt.args.checkDigit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckTransposition() = %v, want %v", got, tt.want)
			}
		})
	}
}
