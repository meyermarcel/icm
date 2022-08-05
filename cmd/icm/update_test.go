package main

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/meyermarcel/icm/cont"
)

func Test_parseOwners(t *testing.T) {
	tests := []struct {
		name    string
		body    io.Reader
		want    []cont.Owner
		wantErr bool
	}{
		{
			"Parsing valid HTML body returns owners map",
			validBody(),
			[]cont.Owner{
				{
					Code:    "AAA",
					Company: "A Company",
					City:    "A City",
					Country: "A Country",
				},
				{
					Code:    "BBB",
					Company: "",
					City:    "B City",
					Country: "B Country",
				},
				{
					Code:    "CCC",
					Company: "C Company",
					City:    "",
					Country: "C Country",
				},
				{
					Code:    "DDD",
					Company: "D Company",
					City:    "D City",
					Country: "",
				},
			},
			false,
		},
		{
			"Parsing invalid HTML body with invalid length of owner code returns error",
			codeInvalidLength(),
			nil,
			true,
		},
		{
			"Parsing invalid HTML body with some missing <td> tags returns error",
			missingTdTags(),
			nil,
			true,
		},
		{
			"Parsing invalid HTML body with no owner returns error",
			noOwner(),
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseOwners(tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOwners() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOwners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func codeInvalidLength() io.Reader {
	return strings.NewReader(`<!DOCTYPE html>
<body>
	<table>
		<tr>
			<td data-label="Code">AA</td>
			<td data-label="Company">A Company</td>
			<td data-label="City" class="hidden-xs">A City</td>
			<td data-label="Country" class="hidden-xs">A Country</td>
			<td data-label="Details"><a href="https://link">View</a></td>
		</tr>
	</table>
</body>
</html>
`)
}

func missingTdTags() io.Reader {
	return strings.NewReader(`<!DOCTYPE html>
<body>
	<table>
		<tr>
			<td data-label="Code">AAAU</td>
			<td data-label="Company">A Company</td>
			<td data-label="City" class="hidden-xs">A City</td>
		</tr>
	</table>
</body>
</html>
`)
}

func noOwner() io.Reader {
	return strings.NewReader(`<!DOCTYPE html>
	<body>
		<table>
			<tr>
			</tr>
		</table>
	</body>
	</html>
`)
}

func validBody() io.Reader {
	return strings.NewReader(`<!DOCTYPE html>
<body>
	<table>
		<tr>
			<td data-label="Code">AAAU</td>
			<td data-label="Company">A Company</td>
			<td data-label="City" class="hidden-xs">A City</td>
			<td data-label="Country" class="hidden-xs">A Country</td>
			<td data-label="Details"><a href="https://link">View</a></td>
		</tr>
		<tr>
			<td data-label="Code">BBBU</td>
			<td data-label="Company"></td>
			<td data-label="City" class="hidden-xs">B City</td>
			<td data-label="Country" class="hidden-xs">B Country</td>
			<td data-label="Details"><a href="https://link">View</a></td>
		</tr>
		<tr>
			<td data-label="Code">CCCU</td>
			<td data-label="Company">C Company</td>
			<td data-label="City" class="hidden-xs"></td>
			<td data-label="Country" class="hidden-xs">C Country</td>
			<td data-label="Details"><a href="https://link">View</a></td>
		</tr>
		<tr>
			<td data-label="Code">DDDU</td>
			<td data-label="Company">D Company</td>
			<td data-label="City" class="hidden-xs">D City</td>
			<td data-label="Country" class="hidden-xs"></td>
			<td data-label="Details"><a href="https://link">View</a></td>
		</tr>
	</table>
</body>
</html>
`)
}
