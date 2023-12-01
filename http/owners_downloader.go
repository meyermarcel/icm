package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/meyermarcel/icm/cont"
	"golang.org/x/net/html"
)

type ownersDownloader struct {
	ownerURL string
}

func NewOwnersDownloader(ownerURL string) OwnersDownloader {
	return &ownersDownloader{ownerURL: ownerURL}
}

type OwnersDownloader interface {
	Download() ([]cont.Owner, error)
}

func (od *ownersDownloader) Download() ([]cont.Owner, error) {
	resp, err := http.Get(od.ownerURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	owners, err := parseOwners(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	return owners, nil
}

func parseOwners(body io.Reader) ([]cont.Owner, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	owners := make([]cont.Owner, 0)

	var getOwnerNode func(*html.Node) error
	getOwnerNode = func(n *html.Node) error {
		codeWithU := parseHTMLtdData(n, "Code:")
		if codeWithU != "" {

			if len(codeWithU) < 4 {
				return fmt.Errorf("parsing HTML failed of owner code failed because '%s' is too short", codeWithU)
			}
			code := codeWithU[0:3]

			companyTdNode := nextTdSibling(n)
			company := parseHTMLtdData(companyTdNode, "Company:")
			cityNode := nextTdSibling(nextTdSibling(companyTdNode))
			city := parseHTMLtdData(cityNode, "City:")
			countryNode := nextTdSibling(nextTdSibling(cityNode))
			countryCode := parseHTMLtdData(countryNode, "Country:")
			country, ok := countryCodeMap[countryCode]
			if !ok {
				country = countryCode
			}

			owners = append(owners,
				cont.Owner{
					Code:    code,
					Company: company,
					City:    city,
					Country: country,
				},
			)

			// If valid td tag found continue with sibling
			// instead of parsing every child.
			err := getOwnerNode(n.NextSibling)
			if err != nil {
				return err
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			err := getOwnerNode(c)
			if err != nil {
				return err
			}
		}
		return nil
	}
	err = getOwnerNode(doc)
	if err != nil {
		return nil, err
	}
	if len(owners) == 0 {
		return nil, fmt.Errorf("parsing HTML failed because no owner was parsed")
	}
	return owners, nil
}

func parseHTMLtdData(td *html.Node, spanDescription string) string {
	if td.Type == html.ElementNode && td.Data == "td" {
		// Iterate through nodes because simple inner text (e.g. '\n') is also a node.
		for c := td.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && c.Data == "span" && c.FirstChild != nil && c.FirstChild.Data == spanDescription {
				for ns := c.NextSibling; ns != nil; ns = ns.NextSibling {
					if ns.Type == html.ElementNode && ns.Data == "span" {
						// Empty span:
						// <span></span>
						if ns.FirstChild == nil {
							return ""
						}
						return ns.FirstChild.Data
					}
				}
			}
		}
	}
	return ""
}

func nextTdSibling(td *html.Node) *html.Node {
	for ns := td.NextSibling; ns != nil; ns = ns.NextSibling {
		if ns.Type == html.ElementNode && ns.Data == "td" {
			return ns
		}
	}
	return nil
}

// Copied from Wikipedia "ISO 3166-1 alpha-2" and manually adjusted.
var countryCodeMap = map[string]string{
	"AD": "Andorra",
	"AE": "United Arab Emirates",
	"AF": "Afghanistan",
	"AG": "Antigua and Barbuda",
	"AI": "Anguilla",
	"AL": "Albania",
	"AM": "Armenia",
	"AO": "Angola",
	"AQ": "Antarctica",
	"AR": "Argentina",
	"AS": "American Samoa",
	"AT": "Austria",
	"AU": "Australia",
	"AW": "Aruba",
	"AX": "Åland Islands",
	"AZ": "Azerbaijan",
	"BA": "Bosnia and Herzegovina",
	"BB": "Barbados",
	"BD": "Bangladesh",
	"BE": "Belgium",
	"BF": "Burkina Faso",
	"BG": "Bulgaria",
	"BH": "Bahrain",
	"BI": "Burundi",
	"BJ": "Benin",
	"BL": "Saint Barthélemy",
	"BM": "Bermuda",
	"BN": "Brunei Darussalam",
	"BO": "Bolivia",
	"BQ": "Bonaire, Sint Eustatius and Saba",
	"BR": "Brazil",
	"BS": "Bahamas",
	"BT": "Bhutan",
	"BV": "Bouvet Island",
	"BW": "Botswana",
	"BY": "Belarus",
	"BZ": "Belize",
	"CA": "Canada",
	"CC": "Cocos (Keeling) Islands",
	"CD": "Democratic Republic of the Congo",
	"CF": "Central African Republic",
	"CG": "Congo",
	"CH": "Switzerland",
	"CI": "Côte d'Ivoire",
	"CK": "Cook Islands",
	"CL": "Chile",
	"CM": "Cameroon",
	"CN": "China",
	"CO": "Colombia",
	"CR": "Costa Rica",
	"CU": "Cuba",
	"CV": "Cabo Verde",
	"CW": "Curaçao",
	"CX": "Christmas Island",
	"CY": "Cyprus",
	"CZ": "Czechia",
	"DE": "Germany",
	"DJ": "Djibouti",
	"DK": "Denmark",
	"DM": "Dominica",
	"DO": "Dominican Republic",
	"DZ": "Algeria",
	"EC": "Ecuador",
	"EE": "Estonia",
	"EG": "Egypt",
	"EH": "Western Sahara",
	"ER": "Eritrea",
	"ES": "Spain",
	"ET": "Ethiopia",
	"FI": "Finland",
	"FJ": "Fiji",
	"FK": "Falkland Islands (Malvinas)",
	"FM": "Micronesia",
	"FO": "Faroe Islands",
	"FR": "France",
	"GA": "Gabon",
	"GB": "United Kingdom",
	"GD": "Grenada",
	"GE": "Georgia",
	"GF": "French Guiana",
	"GG": "Guernsey",
	"GH": "Ghana",
	"GI": "Gibraltar",
	"GL": "Greenland",
	"GM": "Gambia",
	"GN": "Guinea",
	"GP": "Guadeloupe",
	"GQ": "Equatorial Guinea",
	"GR": "Greece",
	"GS": "South Georgia and the South Sandwich Islands",
	"GT": "Guatemala",
	"GU": "Guam",
	"GW": "Guinea-Bissau",
	"GY": "Guyana",
	"HK": "Hong Kong",
	"HM": "Heard Island and McDonald Islands",
	"HN": "Honduras",
	"HR": "Croatia",
	"HT": "Haiti",
	"HU": "Hungary",
	"ID": "Indonesia",
	"IE": "Ireland",
	"IL": "Israel",
	"IM": "Isle of Man",
	"IN": "India",
	"IO": "British Indian Ocean Territory",
	"IQ": "Iraq",
	"IR": "Iran",
	"IS": "Iceland",
	"IT": "Italy",
	"JE": "Jersey",
	"JM": "Jamaica",
	"JO": "Jordan",
	"JP": "Japan",
	"KE": "Kenya",
	"KG": "Kyrgyzstan",
	"KH": "Cambodia",
	"KI": "Kiribati",
	"KM": "Comoros",
	"KN": "Saint Kitts and Nevis",
	"KP": "Korea (Democratic People's Republic of)",
	"KR": "Korea, Republic of",
	"KW": "Kuwait",
	"KY": "Cayman Islands",
	"KZ": "Kazakhstan",
	"LA": "Lao",
	"LB": "Lebanon",
	"LC": "Saint Lucia",
	"LI": "Liechtenstein",
	"LK": "Sri Lanka",
	"LR": "Liberia",
	"LS": "Lesotho",
	"LT": "Lithuania",
	"LU": "Luxembourg",
	"LV": "Latvia",
	"LY": "Libya",
	"MA": "Morocco",
	"MC": "Monaco",
	"MD": "Moldova",
	"ME": "Montenegro",
	"MF": "Saint Martin (French part)",
	"MG": "Madagascar",
	"MH": "Marshall Islands",
	"MK": "North Macedonia",
	"ML": "Mali",
	"MM": "Myanmar",
	"MN": "Mongolia",
	"MO": "Macao",
	"MP": "Northern Mariana Islands",
	"MQ": "Martinique",
	"MR": "Mauritania",
	"MS": "Montserrat",
	"MT": "Malta",
	"MU": "Mauritius",
	"MV": "Maldives",
	"MW": "Malawi",
	"MX": "Mexico",
	"MY": "Malaysia",
	"MZ": "Mozambique",
	"NA": "Namibia",
	"NC": "New Caledonia",
	"NE": "Niger",
	"NF": "Norfolk Island",
	"NG": "Nigeria",
	"NI": "Nicaragua",
	"NL": "Netherlands",
	"NO": "Norway",
	"NP": "Nepal",
	"NR": "Nauru",
	"NU": "Niue",
	"NZ": "New Zealand",
	"OM": "Oman",
	"PA": "Panama",
	"PE": "Peru",
	"PF": "French Polynesia",
	"PG": "Papua New Guinea",
	"PH": "Philippines",
	"PK": "Pakistan",
	"PL": "Poland",
	"PM": "Saint Pierre and Miquelon",
	"PN": "Pitcairn",
	"PR": "Puerto Rico",
	"PS": "Palestine",
	"PT": "Portugal",
	"PW": "Palau",
	"PY": "Paraguay",
	"QA": "Qatar",
	"RE": "Réunion",
	"RO": "Romania",
	"RS": "Serbia",
	"RU": "Russian Federation",
	"RW": "Rwanda",
	"SA": "Saudi Arabia",
	"SB": "Solomon Islands",
	"SC": "Seychelles",
	"SD": "Sudan",
	"SE": "Sweden",
	"SG": "Singapore",
	"SH": "Saint Helena, Ascension and Tristan da Cunha",
	"SI": "Slovenia",
	"SJ": "Svalbard and Jan Mayen",
	"SK": "Slovakia",
	"SL": "Sierra Leone",
	"SM": "San Marino",
	"SN": "Senegal",
	"SO": "Somalia",
	"SR": "Suriname",
	"SS": "South Sudan",
	"ST": "Sao Tome and Principe",
	"SV": "El Salvador",
	"SX": "Sint Maarten (Dutch part)",
	"SY": "Syrian Arab Republic",
	"SZ": "Eswatini",
	"TC": "Turks and Caicos Islands",
	"TD": "Chad",
	"TF": "French Southern Territories",
	"TG": "Togo",
	"TH": "Thailand",
	"TJ": "Tajikistan",
	"TK": "Tokelau",
	"TL": "Timor-Leste",
	"TM": "Turkmenistan",
	"TN": "Tunisia",
	"TO": "Tonga",
	"TR": "Türkiye",
	"TT": "Trinidad and Tobago",
	"TV": "Tuvalu",
	"TW": "Taiwan (Republic of China)",
	"TZ": "Tanzania",
	"UA": "Ukraine",
	"UG": "Uganda",
	"UM": "United States Minor Outlying Islands",
	"US": "United States of America",
	"UY": "Uruguay",
	"UZ": "Uzbekistan",
	"VA": "Holy See",
	"VC": "Saint Vincent and the Grenadines",
	"VE": "Venezuela",
	"VG": "Virgin Islands (British)",
	"VI": "Virgin Islands (U.S.)",
	"VN": "Viet Nam",
	"VU": "Vanuatu",
	"WF": "Wallis and Futuna",
	"WS": "Samoa",
	"YE": "Yemen",
	"YT": "Mayotte",
	"ZA": "South Africa",
	"ZM": "Zambia",
	"ZW": "Zimbabwe",
}
