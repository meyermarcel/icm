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
			name: "Test ABC U 028373 (1)0",
			args: args{
				ownerCode:  "ABC",
				equipCatID: "U",
				serialNum:  "028373",
			},
			want: []Number{
				newNum("ABC", "U", "208373", 0),
				newNum("ABC", "U", "023873", 0),
				newNum("ABC", "U", "028337", 0),
				newNum("ABC", "U", "028370", 3),
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
