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
			},
			want: []Number{
				newNum("CMA", "U", "169312", 0),
				newNum("CMA", "U", "163192", 0),
			},
		},
		{
			name: "Test RCB U 001130 0",
			args: args{
				ownerCode:  "RCB",
				equipCatID: "U",
				serialNum:  "001130",
			},
			want: []Number{
				newNum("RCB", "U", "010130", 0),
			},
		},
		{
			name: "Test WSL U 801743 (1)0",
			args: args{
				ownerCode:  "WSL",
				equipCatID: "U",
				serialNum:  "801743",
			},
			want: []Number{
				newNum("WSL", "U", "810743", 0),
				newNum("WSL", "U", "807143", 0),
				newNum("WSL", "U", "801740", 3),
			},
		},
		{
			name: "Test APL U 689473 (1)0",
			args: args{
				ownerCode:  "APL",
				equipCatID: "U",
				serialNum:  "689473",
			},
			want: []Number{
				newNum("APL", "U", "869473", 0),
				newNum("APL", "U", "698473", 0),
				newNum("APL", "U", "684973", 0),
				newNum("APL", "U", "689743", 0),
				newNum("APL", "U", "689437", 0),
				newNum("APL", "U", "689470", 3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckTransposition(tt.args.ownerCode, tt.args.equipCatID, tt.args.serialNum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckTransposition() = %v, want %v", got, tt.want)
			}
		})
	}
}
