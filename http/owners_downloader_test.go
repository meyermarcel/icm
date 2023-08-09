package http

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
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Code:</span>
                <span>AAA</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Company:</span>
                <span>A Company</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">City:</span>
                <span>A City</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Country:</span>
                <span>A Country</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Details:</span>
                <a class="upperCase withArrow" href="/bic-codes/aaau">View</a>
            </td>
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
            <span class="hideOnDesktop tdHeading">Code:</span>
            <span>AAAU</span>
            <span class="hideOnDesktop tdHeading">Company:</span>
            <span>A Company</span>
            <span class="hideOnDesktop tdHeading">City:</span>
            <span>A City</span>
            <span class="hideOnDesktop tdHeading">Country:</span>
            <span>A Country</span>
            <span class="hideOnDesktop tdHeading">Details:</span>
            <a class="upperCase withArrow" href="/bic-codes/aaau">View</a>
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
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Code:</span>
                <span>AAAU</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Company:</span>
                <span>A Company</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">City:</span>
                <span>A City</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Country:</span>
                <span>A Country</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Details:</span>
                <a class="upperCase withArrow" href="/bic-codes/aaau">View</a>
            </td>
        </tr>
		<tr>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Code:</span>
                <span>BBBU</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Company:</span>
                <span></span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">City:</span>
                <span>B City</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Country:</span>
                <span>B Country</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Details:</span>
                <a class="upperCase withArrow" href="/bic-codes/aaau">View</a>
            </td>
        </tr>
		<tr>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Code:</span>
                <span>CCCU</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Company:</span>
                <span>C Company</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">City:</span>
                <span></span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Country:</span>
                <span>C Country</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Details:</span>
                <a class="upperCase withArrow" href="/bic-codes/aaau">View</a>
            </td>
        </tr>
<tr>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Code:</span>
                <span>DDDU</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Company:</span>
                <span>D Company</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">City:</span>
                <span>D City</span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Country:</span>
                <span></span>
            </td>
			<td class="flexMobile align-items-center">
                <span class="hideOnDesktop tdHeading">Details:</span>
                <a class="upperCase withArrow" href="/bic-codes/aaau">View</a>
            </td>
        </tr>
	</table>
</body>
</html>
`)
}
