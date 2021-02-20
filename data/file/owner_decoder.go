// Copyright © 2018 Marcel Meyer meyermarcel@posteo.de
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package file

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/meyermarcel/icm/data"

	"io/ioutil"

	"github.com/meyermarcel/icm/cont"
)

const ownerFileName = "owner.json"

type owner struct {
	Code    string `json:"code"`
	Company string `json:"company"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// NewOwnerDecoderUpdater writes owner file to path if it not exists and
// returns a struct that uses this file as a data source.
func NewOwnerDecoderUpdater(path string) (data.OwnerDecodeUpdater, error) {

	ownersFile := &ownerDecoderUpdater{path: path}
	filePath := filepath.Join(ownersFile.path, ownerFileName)
	if err := initFile(filePath, []byte(ownerJSON)); err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &ownersFile.owners); err != nil {
		return nil, err
	}
	for ownerCode := range ownersFile.owners {
		if err := cont.IsOwnerCode(ownerCode); err != nil {
			return nil, err
		}
	}
	return ownersFile, nil
}

type ownerDecoderUpdater struct {
	owners map[string]owner
	path   string
}

// Decode returns an owner for an owner code.
func (of *ownerDecoderUpdater) Decode(code string) (bool, cont.Owner) {
	if val, ok := of.owners[code]; ok {
		return true, cont.Owner{
			Code:    code,
			Company: val.Company,
			City:    val.City,
			Country: val.Country,
		}
	}
	return false, cont.Owner{}
}

// GetAllOwnerCodes returns a count of owner codes.
func (of *ownerDecoderUpdater) GetAllOwnerCodes() []string {
	var codes []string
	for _, owner := range of.owners {
		codes = append(codes, owner.Code)
	}
	return codes
}

// Update accepts a map of owner code to owner and replaces/adds entries in the local owner file.
// Cancelled owners still exist to prevent removal of custom owners created by the user.
func (of *ownerDecoderUpdater) Update(newOwners map[string]cont.Owner) error {
	for k, v := range newOwners {
		of.owners[k] = toSerializableOwner(v)
	}
	b, err := marshalNoHTMLEsc(of.owners)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(of.path, ownerFileName), b, 0644); err != nil {
		return err
	}
	return nil
}

func marshalNoHTMLEsc(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	if err != nil {
		return nil, err
	}
	var fmtJSON bytes.Buffer
	err = json.Indent(&fmtJSON, buffer.Bytes(), "", "  ")
	if err != nil {
		return nil, err
	}
	return fmtJSON.Bytes(), nil
}

func toSerializableOwner(ownerToConvert cont.Owner) owner {
	return owner{ownerToConvert.Code,
		ownerToConvert.Company,
		ownerToConvert.City,
		ownerToConvert.Country}
}

func initFile(path string, content []byte) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := ioutil.WriteFile(path, content, 0644); err != nil {
			return err
		}
	}
	return nil
}

const ownerJSON = `{
  "AAA": {
    "code": "AAA",
    "company": "ASIA CONTAINER LEASING CO LTD",
    "city": "WAN CHAI",
    "country": "HK"
  },
  "AAC": {
    "code": "AAC",
    "company": "ACE GLOBAL LINES DWC - LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "AAF": {
    "code": "AAF",
    "company": "BUNDESMINISTERIUM FUR LANDESVERTEIDIGUNG",
    "city": "WIEN",
    "country": "Austria"
  },
  "AAG": {
    "code": "AAG",
    "company": "MTU ONSITE ENERGY SYSTEMS GMBH",
    "city": "RUHSTORF",
    "country": "Germany"
  },
  "AAI": {
    "code": "AAI",
    "company": "A.A SHIPPING LTD",
    "city": "TEL AVIV",
    "country": "Israel"
  },
  "AAM": {
    "code": "AAM",
    "company": "ALBATROSS TANK-LEASING CO., LIMITED",
    "city": "HONG KONG",
    "country": "HK"
  },
  "AAT": {
    "code": "AAT",
    "company": "AMFICO AGENCIES PVT LTD",
    "city": "Mumbai",
    "country": "India"
  },
  "ABB": {
    "code": "ABB",
    "company": "ATLANTIC BULK CARRIERS MANAGMENT LTD",
    "city": "NEW YORK, NY 10019",
    "country": "United States"
  },
  "ABC": {
    "code": "ABC",
    "company": "ABC MOBILE STORAGE",
    "city": "NORFOLK, NE-69702",
    "country": "United States"
  },
  "ABE": {
    "code": "ABE",
    "company": "L & T B.V.",
    "city": "'s-Hertogenbosch",
    "country": "Netherlands"
  },
  "ABL": {
    "code": "ABL",
    "company": "ETAT-MAJOR DES FORCES ARMEES BELGES",
    "city": "BRUSSELS",
    "country": "Belgium"
  },
  "ABM": {
    "code": "ABM",
    "company": "RM RAIL ABAKANVAGONMASH",
    "city": "ABAKAN",
    "country": "Russian Federation"
  },
  "ABY": {
    "code": "ABY",
    "company": "ABBEY LOGISTICS GROUP",
    "city": "LIVERPOOL",
    "country": "United Kingdom"
  },
  "ACA": {
    "code": "ACA",
    "company": "ARCTIC CONTAINER OY",
    "city": "HELSINKI",
    "country": "Finland"
  },
  "ACB": {
    "code": "ACB",
    "company": "ACB CONTAINERS NV",
    "city": "Antwerpen",
    "country": "Belgium"
  },
  "ACC": {
    "code": "ACC",
    "company": "ATLANTIC COAST CONTAINER",
    "city": "GREENVILLE, SC-29609",
    "country": "United States"
  },
  "ACD": {
    "code": "ACD",
    "company": "DESOTEC NV",
    "city": "ROESELARE",
    "country": "Belgium"
  },
  "ACE": {
    "code": "ACE",
    "company": "ACE CONTAINER SERVICES LTD",
    "city": "LEEDS, LS10 1RT",
    "country": "United Kingdom"
  },
  "ACI": {
    "code": "ACI",
    "company": "ALLSTATE CONTAINER INC",
    "city": "VALLEY STREAM, NY 11581",
    "country": "United States"
  },
  "ACL": {
    "code": "ACL",
    "company": "ATLANTIC CONTAINER LINE",
    "city": "WESTFIELD, NJ 07090",
    "country": "United States"
  },
  "ACN": {
    "code": "ACN",
    "company": "NOURYON FUNCTIONAL CHEMICALS",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "ACP": {
    "code": "ACP",
    "company": "CARDIGAS",
    "city": "HEUSDEN-ZOLDER",
    "country": "Belgium"
  },
  "ACR": {
    "code": "ACR",
    "company": "ACCIPITER OILFIELD EQUIPMENT",
    "city": "DALIAN",
    "country": "China"
  },
  "ACS": {
    "code": "ACS",
    "company": "COMMONWEALTH INDEPENDENT",
    "city": "MONTREAL, QC",
    "country": "Canada"
  },
  "ACX": {
    "code": "ACX",
    "company": "G2 OCEAN AS",
    "city": "Bergen",
    "country": "Norway"
  },
  "ADB": {
    "code": "ADB",
    "company": "AGROCEL INDUSTRIES PVT LTD.",
    "city": "Bhuj",
    "country": "India"
  },
  "ADH": {
    "code": "ADH",
    "company": "ADHOC S.R.L",
    "city": "VENICE",
    "country": "Italy"
  },
  "ADL": {
    "code": "ADL",
    "company": "QINGDAO IRFAN INTERNATIONAL LOGISTICS CO LTD",
    "city": "Qingdao",
    "country": "China"
  },
  "ADM": {
    "code": "ADM",
    "company": "ADMIRAL CONTAINER LINES INC LIMITED",
    "city": "VALLETTA",
    "country": "Malta"
  },
  "ADR": {
    "code": "ADR",
    "company": "BARCELONESA DE DROGAS Y PRODUCTOS QUIMICOS SA",
    "city": "CORNELLA DE LLOBREGAT",
    "country": "Spain"
  },
  "ADT": {
    "code": "ADT",
    "company": "ADEN TRANSEXIM SRL",
    "city": "Chisinau",
    "country": "Moldova, Republic of"
  },
  "ADX": {
    "code": "ADX",
    "company": "APPLIED CRYO TECHNOLOGIES",
    "city": "HOUSTON, TX-77075",
    "country": "United States"
  },
  "ADY": {
    "code": "ADY",
    "company": "ADY CONTAINER LLC",
    "city": "Baku",
    "country": "Azerbaijan"
  },
  "AEG": {
    "code": "AEG",
    "company": "AEGEUS ENERGY SOLUTIONS LTD",
    "city": "Tadcaster",
    "country": "United Kingdom"
  },
  "AEL": {
    "code": "AEL",
    "company": "MITSUBISHI CHEMICAL CORPORATION",
    "city": "TOKYO",
    "country": "Japan"
  },
  "AES": {
    "code": "AES",
    "company": "AES ANDRES",
    "city": "Santo Domingo",
    "country": "Dominican Republic"
  },
  "AEV": {
    "code": "AEV",
    "company": "BLUEHORN SA",
    "city": "Petit-Lancy",
    "country": "Switzerland"
  },
  "AEX": {
    "code": "AEX",
    "company": "AFRICA EXPRESS LINE LTD",
    "city": "Kent  ME19 4UY",
    "country": "United Kingdom"
  },
  "AFB": {
    "code": "AFB",
    "company": "AFFILIPS NV",
    "city": "TIENEN",
    "country": "Belgium"
  },
  "AFI": {
    "code": "AFI",
    "company": "AGMARK CAPITAL HOLDINGS INC.",
    "city": "NASHVILLE, TN 37201",
    "country": "United States"
  },
  "AFL": {
    "code": "AFL",
    "company": "AIR FLOW",
    "city": "ROUSSET",
    "country": "France"
  },
  "AFM": {
    "code": "AFM",
    "company": "ALFIL LOGISTICS S.A.",
    "city": "EL PRAT DE LLOBREGAT",
    "country": "Spain"
  },
  "AFS": {
    "code": "AFS",
    "company": "L'AUTOFRIGO SUD SRL",
    "city": "MONOPOLI",
    "country": "Italy"
  },
  "AGA": {
    "code": "AGA",
    "company": "LINDE GAS BENELUX",
    "city": "SCHIEDAM",
    "country": "Netherlands"
  },
  "AGB": {
    "code": "AGB",
    "company": "ARCHEAN CHEMICAL INDUSTRIES PRIVATE LIMITED",
    "city": "CHENNAI , TAMILNADU",
    "country": "India"
  },
  "AGC": {
    "code": "AGC",
    "company": "GREIWING LOGISTICS FOR YOU GMBH",
    "city": "GREVEN",
    "country": "Germany"
  },
  "AGF": {
    "code": "AGF",
    "company": "THE ARABIAN INDUSTRIAL GASES CO LLC DUBAI BRANCH",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "AGG": {
    "code": "AGG",
    "company": "ALISTAIR JAMES COMPANY",
    "city": "Dar es Salaam",
    "country": "Tanzania, United Republic of"
  },
  "AGH": {
    "code": "AGH",
    "company": "AFRICA GLOBAL LOGISTICS",
    "city": "ROUBAIX",
    "country": "France"
  },
  "AGK": {
    "code": "AGK",
    "company": "AGGREKO MANUFACTURING",
    "city": "DUMBARTON G82 3RG",
    "country": "United Kingdom"
  },
  "AGL": {
    "code": "AGL",
    "company": "AGUNSA-AGENCIA UNIVERSALES S.A",
    "city": "SANTIAGO",
    "country": "Chile"
  },
  "AGM": {
    "code": "AGM",
    "company": "AGMARK CAPITAL HOLDINGS INC.",
    "city": "NASHVILLE, TN 37201",
    "country": "United States"
  },
  "AGN": {
    "code": "AGN",
    "company": "ANGA SP. Z.O.O.",
    "city": "Gdansk",
    "country": "Poland"
  },
  "AGO": {
    "code": "AGO",
    "company": "AIRGAS USA LLC.",
    "city": "Tulsa, OK 74120-1633",
    "country": "United States"
  },
  "AGP": {
    "code": "AGP",
    "company": "A-GAS (UK) LTD",
    "city": "BRISTOL, BS20 7XH",
    "country": "United Kingdom"
  },
  "AGR": {
    "code": "AGR",
    "company": "AGRICOOL SAS",
    "city": "La Courneuve",
    "country": "France"
  },
  "AGS": {
    "code": "AGS",
    "company": "AGA GAS AB",
    "city": "MITCHELDEAN",
    "country": "United Kingdom"
  },
  "AGV": {
    "code": "AGV",
    "company": "AGROTRUST GMBH",
    "city": "Visbek",
    "country": "Germany"
  },
  "AGX": {
    "code": "AGX",
    "company": "AGT KONTEYNER LTD STI",
    "city": "Istanbul (Bahçelievler)",
    "country": "Turkey"
  },
  "AHG": {
    "code": "AHG",
    "company": "AL GHAITH INDUSTRIES L.L.C",
    "city": "Abu Dhabi",
    "country": "United Arab Emirates"
  },
  "AHL": {
    "code": "AHL",
    "company": "ABNORMAL LOAD ENGINEERING LIMITED",
    "city": "HIXON",
    "country": "United Kingdom"
  },
  "AIC": {
    "code": "AIC",
    "company": "ALLIED CONTAINER LINE PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "AID": {
    "code": "AID",
    "company": "ALBEMARLE CORPORATION",
    "city": "BATON ROUGE, LA 70801",
    "country": "United States"
  },
  "AIG": {
    "code": "AIG",
    "company": "AIRGGAS USA, LLC",
    "city": "Tulsa, OK 74120-1633",
    "country": "United States"
  },
  "AIL": {
    "code": "AIL",
    "company": "ARMITA INDIA SHIPPING PVT LTD",
    "city": "Mumbai",
    "country": "India"
  },
  "AIM": {
    "code": "AIM",
    "company": "TRS CONTAINERS & CHASSIS",
    "city": "AVENEL, NJ 0700-0188",
    "country": "United States"
  },
  "AIP": {
    "code": "AIP",
    "company": "A.I.P.P INDUSTRIES LTD",
    "city": "Kfar qari",
    "country": "Israel"
  },
  "AIR": {
    "code": "AIR",
    "company": "AIR PRODUCTS",
    "city": "ALLENTOWN, PA 18195",
    "country": "United States"
  },
  "AIS": {
    "code": "AIS",
    "company": "AZIENDA SERVIZI TRASPORTI LOGISTICA S.R.L.",
    "city": "CASALMAGGIORE",
    "country": "Italy"
  },
  "AIT": {
    "code": "AIT",
    "company": "ARKEMA B.V. – LOCATION ROTTERDAM",
    "city": "Vondelingenplaat",
    "country": "Netherlands"
  },
  "AIY": {
    "code": "AIY",
    "company": "AIYER SHIPPING CONTAINER LINE",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "AKB": {
    "code": "AKB",
    "company": "ALLNEX RESINS NETHERLANDS BV",
    "city": "BERGEN OP ZOOM",
    "country": "Netherlands"
  },
  "AKK": {
    "code": "AKK",
    "company": "AKKON DENIZCILIK NAKLIYAT VE TICARET A.S.",
    "city": "İstanbul",
    "country": "Turkey"
  },
  "AKL": {
    "code": "AKL",
    "company": "KAWASAKI KISEN KAISHA LTD - K LINE",
    "city": "TOKYO",
    "country": "Japan"
  },
  "AKM": {
    "code": "AKM",
    "company": "ALTALOGOS LLC",
    "city": "EKATERINBURG",
    "country": "Russian Federation"
  },
  "AKO": {
    "code": "AKO",
    "company": "HONG KONG SPECIALITY GASES CO.LTD",
    "city": "CHEUNG SHAWAN, KLN",
    "country": "HK"
  },
  "AKV": {
    "code": "AKV",
    "company": "AKVAZASHITA LLC",
    "city": "KALININGRAD",
    "country": "Russian Federation"
  },
  "ALA": {
    "code": "ALA",
    "company": "JSC EKOPET",
    "city": "KALININGRAD",
    "country": "Russian Federation"
  },
  "ALB": {
    "code": "ALB",
    "company": "ALBLAS INT. TRANSPORT BV",
    "city": "'S GRAVENDEEL",
    "country": "Netherlands"
  },
  "ALC": {
    "code": "ALC",
    "company": "AARTI INDUSTRIES LIMITED",
    "city": "MULUND WEST MUMBAI",
    "country": "India"
  },
  "ALE": {
    "code": "ALE",
    "company": "AIR LIQUIDE ELECTRONICS U.S. LP",
    "city": "MORRISVILLE, PA 19067",
    "country": "United States"
  },
  "ALF": {
    "code": "ALF",
    "company": "AIR LIQUIDE FAR EASTERN LTD",
    "city": "WUCHI /TAICHUNG",
    "country": "Taiwan, China"
  },
  "ALG": {
    "code": "ALG",
    "company": "AIR LIQUIDE GLOBAL HELIUM FZE",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "ALH": {
    "code": "ALH",
    "company": "AIR LIQUIDE TRINIDAD & TOBAGO LIMITED",
    "city": "POINT LISAS",
    "country": "Trinidad and Tobago"
  },
  "ALI": {
    "code": "ALI",
    "company": "UNIVERSAL FOREST PRODUCTS, INC.",
    "city": "MEDLEY, FL 33178",
    "country": "United States"
  },
  "ALJ": {
    "code": "ALJ",
    "company": "AIR LIQUIDE JAPAN LTD",
    "city": "TOKYO",
    "country": "Japan"
  },
  "ALK": {
    "code": "ALK",
    "company": "ALTUN LOJISTIK A.S.",
    "city": "MERSIN",
    "country": "Turkey"
  },
  "ALL": {
    "code": "ALL",
    "company": "ALCONET BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "ALM": {
    "code": "ALM",
    "company": "ALMAR CONTAINER INVESTMENTS INC.",
    "city": "DURBAN",
    "country": "South Africa"
  },
  "ALN": {
    "code": "ALN",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "ALO": {
    "code": "ALO",
    "company": "AIR LIQUIDE OIL AND GASES SERVICES LIMITED",
    "city": "ABERDEEN",
    "country": "United Kingdom"
  },
  "ALP": {
    "code": "ALP",
    "company": "AIR LIQUIDE FRANCE INDUSTRIE",
    "city": "CHALON-SUR-SAÔNE CEDEX",
    "country": "France"
  },
  "ALR": {
    "code": "ALR",
    "company": "ALMAR CONTAINER INVESTMENTS INC.",
    "city": "DURBAN",
    "country": "South Africa"
  },
  "ALS": {
    "code": "ALS",
    "company": "AIR LIQUIDE GAS AB",
    "city": "MALMO",
    "country": "Sweden"
  },
  "ALT": {
    "code": "ALT",
    "company": "VTG TANKTAINER ASSETS GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "ALU": {
    "code": "ALU",
    "company": "NORDIC BULKERS AB",
    "city": "GOTHENBURG",
    "country": "Sweden"
  },
  "ALV": {
    "code": "ALV",
    "company": "AIR LIQUIDE ELECTRONICS MATERIALS (ZHANGJIAGANG) CO.,LTD",
    "city": "JIANG SU PROVINCE",
    "country": "China"
  },
  "ALW": {
    "code": "ALW",
    "company": "AIR LIQUIDE GLOBAL ELECTRONICS MATERIALS",
    "city": "Seoul",
    "country": "Korea, Republic of"
  },
  "AMA": {
    "code": "AMA",
    "company": "ALIANÇA EQUIQ. CONTÊINERES  MANAUS LTDA",
    "city": "MANAUS",
    "country": "Brazil"
  },
  "AMB": {
    "code": "AMB",
    "company": "AMBIKA SRL",
    "city": "CABA",
    "country": "Argentina"
  },
  "AMC": {
    "code": "AMC",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "AME": {
    "code": "AME",
    "company": "ACE CRANES & ENGINEERING FZ-LLC",
    "city": "Al Hamra Al Jazirah",
    "country": "United Arab Emirates"
  },
  "AMF": {
    "code": "AMF",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "AMG": {
    "code": "AMG",
    "company": "AMGAS SERVICES INC",
    "city": "Rocky View",
    "country": "Canada"
  },
  "AMI": {
    "code": "AMI",
    "company": "AMLF LOGISTICS PVT LTD",
    "city": "NAVI MUMBAI",
    "country": "India"
  },
  "AMM": {
    "code": "AMM",
    "company": "AMMANN ITALY SPA",
    "city": "BUSSOLENGO VERONA",
    "country": "Italy"
  },
  "AMO": {
    "code": "AMO",
    "company": "AMOMATIC OY",
    "city": "Paimio",
    "country": "Finland"
  },
  "AMS": {
    "code": "AMS",
    "company": "AMASIS SHIPPING COMPANY LIMITED",
    "city": "Ho Chi Minh",
    "country": "Viet Nam"
  },
  "AMT": {
    "code": "AMT",
    "company": "ASIAN MARINE TRANSPORT CORP",
    "city": "CEBU CITY",
    "country": "Philippines"
  },
  "AMX": {
    "code": "AMX",
    "company": "AUTAMAROCCHI S.P.A",
    "city": "TRIESTE",
    "country": "Italy"
  },
  "AMZ": {
    "code": "AMZ",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "ANE": {
    "code": "ANE",
    "company": "ANSALDO ENERGIA SPA",
    "city": "GENOVA",
    "country": "Italy"
  },
  "ANG": {
    "code": "ANG",
    "company": "ANGOLA LNG LIMITED",
    "city": "Talatona, Belas",
    "country": "Angola"
  },
  "ANH": {
    "code": "ANH",
    "company": "ANHALT LOGISTICS GMBH & CO KG",
    "city": "BARGEN",
    "country": "Germany"
  },
  "ANL": {
    "code": "ANL",
    "company": "ANALCO SA",
    "city": "MONTEVIDEO",
    "country": "Uruguay"
  },
  "ANN": {
    "code": "ANN",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "ANY": {
    "code": "ANY",
    "company": "ONE WAY LEASE INC.",
    "city": "OAKLAND, CA-94607",
    "country": "United States"
  },
  "AOC": {
    "code": "AOC",
    "company": "INNOSPEC LIMITED",
    "city": "CHESHIRE, CH65 4EY",
    "country": "United Kingdom"
  },
  "AOI": {
    "code": "AOI",
    "company": "PT. ALTIC ONE INDONESIA",
    "city": "JAKARTA",
    "country": "Indonesia"
  },
  "AON": {
    "code": "AON",
    "company": "NEC ENERGY SOLUTIONS",
    "city": "WESTBOROUGH, MA-01582",
    "country": "United States"
  },
  "AOR": {
    "code": "AOR",
    "company": "OEG ASIA PACIFIC PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "APA": {
    "code": "APA",
    "company": "APR ENERGY LLC",
    "city": "FL,32226  JACKSONVILLE",
    "country": "United States"
  },
  "APD": {
    "code": "APD",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "APE": {
    "code": "APE",
    "company": "ALSHIRAWI EQUIPMENT CO LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "APH": {
    "code": "APH",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "API": {
    "code": "API",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "APJ": {
    "code": "APJ",
    "company": "AGC PLIBRICO CO., LTD.",
    "city": "Tokyo",
    "country": "Japan"
  },
  "APL": {
    "code": "APL",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "APM": {
    "code": "APM",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "APN": {
    "code": "APN",
    "company": "AQUA PHARMA AS",
    "city": "LILLEHAMMER",
    "country": "Norway"
  },
  "APR": {
    "code": "APR",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "APS": {
    "code": "APS",
    "company": "ALTORFER POWER SYSTEMS, INC",
    "city": "BARTONVILLE, IL-61607",
    "country": "United States"
  },
  "APV": {
    "code": "APV",
    "company": "SUN FLUORO SYSTEM CO LTD",
    "city": "OSAKA",
    "country": "Japan"
  },
  "APX": {
    "code": "APX",
    "company": "APEX GLOBAL ENGINEERING SDN. BHD.",
    "city": "Petaling Jaya",
    "country": "Malaysia"
  },
  "APZ": {
    "code": "APZ",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "AQI": {
    "code": "AQI",
    "company": "AQUAM INSULA PTE LTD",
    "city": "Lautoka",
    "country": "Fiji"
  },
  "ARA": {
    "code": "ARA",
    "company": "ARA F&D",
    "city": "GYEONGSANG-NAMDO",
    "country": "Korea, Republic of"
  },
  "ARC": {
    "code": "ARC",
    "company": "SUNMARINE SHIPPING SERVICES LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "ARD": {
    "code": "ARD",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "ARF": {
    "code": "ARF",
    "company": "SPANISH MINISTRY OF DEFENSE- SPANISH NAV",
    "city": "Madrid",
    "country": "Spain"
  },
  "ARG": {
    "code": "ARG",
    "company": "MAXIMA AIR SEPARATION CENTER LTD",
    "city": "ASHDOD",
    "country": "Israel"
  },
  "ARK": {
    "code": "ARK",
    "company": "ARKAS SHIPPING AND TRANSPORT SA",
    "city": "Istanbul",
    "country": "Turkey"
  },
  "ARM": {
    "code": "ARM",
    "company": "ADAPTAINER LTD",
    "city": "IPSWICH SUFFOLK IP4 1JX",
    "country": "United Kingdom"
  },
  "ARP": {
    "code": "ARP",
    "company": "EQUIPOS MOVILES DE CAMPANA ARPA S.A.U.",
    "city": "ZARAGOZA",
    "country": "Spain"
  },
  "ARR": {
    "code": "ARR",
    "company": "NIPPON GASES OFFSHORE TANKS LIMITED",
    "city": "ABERDEEN AB12 3LG",
    "country": "United Kingdom"
  },
  "ART": {
    "code": "ART",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "ARV": {
    "code": "ARV",
    "company": "GMK-TRANZIT LIMITED COMPANY",
    "city": "NABEREZHNYE CHELNY",
    "country": "Russian Federation"
  },
  "ASC": {
    "code": "ASC",
    "company": "ABS- CONTAINERS BV",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "ASD": {
    "code": "ASD",
    "company": "PRIVATE ENTREPRENEUR VYISLANKO TATIANA PETROVNA",
    "city": "PETROPAVLOVSK KAMCHATSKLY",
    "country": "Russian Federation"
  },
  "ASG": {
    "code": "ASG",
    "company": "NORDIC BULKERS AB",
    "city": "GOTHENBURG",
    "country": "Sweden"
  },
  "ASK": {
    "code": "ASK",
    "company": "AS TRUCKING GROUP LLC",
    "city": "Laredo",
    "country": "United States"
  },
  "ASL": {
    "code": "ASL",
    "company": "ALLIGATOR SHIPPING COMPANY LLC",
    "city": "ABU DHABI",
    "country": "United Arab Emirates"
  },
  "ASM": {
    "code": "ASM",
    "company": "ALPAS SERVIS LTD",
    "city": "Pskov",
    "country": "Russian Federation"
  },
  "ASN": {
    "code": "ASN",
    "company": "ALISAN ULUSLARARASI TASIMACILIK VE TICARET AS",
    "city": "ATASEHIR,  ISTANBUL",
    "country": "Turkey"
  },
  "ASO": {
    "code": "ASO",
    "company": "ATLANTIC SHIPPING LINE LIMITED",
    "city": "BRITISH VIRGIN ISLAND",
    "country": "British Indian Ocean Territory"
  },
  "ASP": {
    "code": "ASP",
    "company": "TANK MANAGEMENT A/S",
    "city": "OSLO",
    "country": "Norway"
  },
  "ASR": {
    "code": "ASR",
    "company": "ARMED FORCES OF THE SLOVAK REPUBLIC",
    "city": "BRATISLAVA",
    "country": "Slovakia"
  },
  "ASS": {
    "code": "ASS",
    "company": "SOCIETA AUTOTRASPORTI SPECIALI S.P.A",
    "city": "CASTELMASSA (RO)",
    "country": "Italy"
  },
  "AST": {
    "code": "AST",
    "company": "ALMALY TK",
    "city": "Astana",
    "country": "Kazakhstan"
  },
  "ASU": {
    "code": "ASU",
    "company": "ALCOSUISSE",
    "city": "BERN 9",
    "country": "Switzerland"
  },
  "ATA": {
    "code": "ATA",
    "company": "AMATO TRANSPORTS AFFRETEMENT",
    "city": "FOS SUR MER",
    "country": "France"
  },
  "ATB": {
    "code": "ATB",
    "company": "A-TAINER & SERVICE",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "ATC": {
    "code": "ATC",
    "company": "ALLSEAS ENGINEERING BV c/o BLUE MARINE OFFSHORE YARD SERVICE B.V.",
    "city": "DELFT",
    "country": "Netherlands"
  },
  "ATI": {
    "code": "ATI",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "ATK": {
    "code": "ATK",
    "company": "ACTIVE TANK SDN. BHD.",
    "city": "Shah Alam",
    "country": "Malaysia"
  },
  "ATL": {
    "code": "ATL",
    "company": "ALBATROSS TANK LEASING B.V.",
    "city": "Moerdijk",
    "country": "Netherlands"
  },
  "ATN": {
    "code": "ATN",
    "company": "ANELTIA TRADING, S.L.U.",
    "city": "Zaragoza",
    "country": "Spain"
  },
  "ATO": {
    "code": "ATO",
    "company": "ARKEMA",
    "city": "COLOMBES",
    "country": "France"
  },
  "ATR": {
    "code": "ATR",
    "company": "ANEL MITORAJ ANDRZEJ I MITORAJ ELZBIETA SP. J.",
    "city": "Sklamierzyce",
    "country": "Poland"
  },
  "ATS": {
    "code": "ATS",
    "company": "ANSTO",
    "city": "Kirrawee DC",
    "country": "Australia"
  },
  "ATT": {
    "code": "ATT",
    "company": "ATI FREIGHT LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "ATV": {
    "code": "ATV",
    "company": "PEACOCK CONTAINER BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "ATX": {
    "code": "ATX",
    "company": "ANDERSON TRUCKING SERVICES",
    "city": "St. Cloud, MN 56301",
    "country": "United States"
  },
  "AUG": {
    "code": "AUG",
    "company": "QUEHENBERGER LOGISTICS GMBH",
    "city": "STRASSWALCHEN",
    "country": "Austria"
  },
  "AUS": {
    "code": "AUS",
    "company": "SPECTAINER PTY LTD",
    "city": "SYDNEY",
    "country": "Australia"
  },
  "AVC": {
    "code": "AVC",
    "company": "AMVAC CHEMICAL COPORATION",
    "city": "AXIS, AL-36505",
    "country": "United States"
  },
  "AVG": {
    "code": "AVG",
    "company": "AVRORA LOGISTICS",
    "city": "St.Petersburg",
    "country": "Russian Federation"
  },
  "AVH": {
    "code": "AVH",
    "company": "JSC ATOMSPETSTRANS",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "AVI": {
    "code": "AVI",
    "company": "HELISOTA",
    "city": "KAUNAS",
    "country": "Lithuania"
  },
  "AVN": {
    "code": "AVN",
    "company": "AVANA LOGISTEK LIMITED",
    "city": "Mumbai",
    "country": "India"
  },
  "AVR": {
    "code": "AVR",
    "company": "SIA ALPA VAGONS",
    "city": "RIGA",
    "country": "Latvia"
  },
  "AVS": {
    "code": "AVS",
    "company": "REC ADVANCED SILICON MATERIALS LLC",
    "city": "SILVER BOW, MT 59750",
    "country": "United States"
  },
  "AWI": {
    "code": "AWI",
    "company": "ALFRED-WEGENER-INSTITUT HELMHOLTZ-ZENTRUM FUR POLAR",
    "city": "BREMERHAVEN",
    "country": "Germany"
  },
  "AWO": {
    "code": "AWO",
    "company": "OY WOIKOSKI AB",
    "city": "VOIKOSKI",
    "country": "Finland"
  },
  "AWS": {
    "code": "AWS",
    "company": "A W SHIP MANAGEMENT LTD",
    "city": "LONDON E1 8DE",
    "country": "United Kingdom"
  },
  "AWX": {
    "code": "AWX",
    "company": "ALASKA WEST EXPRESS INC.",
    "city": "FAIRBANKS, AK 99701",
    "country": "United States"
  },
  "AXE": {
    "code": "AXE",
    "company": "ASEAN SEAS LINE CO LIMITED",
    "city": "Shanghai",
    "country": "China"
  },
  "AXI": {
    "code": "AXI",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "AYT": {
    "code": "AYT",
    "company": "OZTIRYAKILER SAVUNMA EKIPMANLARI SAN.VE.TIC.A.S.",
    "city": "Zeytinburnu-Istanbul",
    "country": "Turkey"
  },
  "AZE": {
    "code": "AZE",
    "company": "EURASIA CONTAINER",
    "city": "Baku",
    "country": "Azerbaijan"
  },
  "AZL": {
    "code": "AZL",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "AZN": {
    "code": "AZN",
    "company": "AMAZON.COM SERVICES, INC.",
    "city": "Atlanta",
    "country": "United States"
  },
  "AZT": {
    "code": "AZT",
    "company": "PROTAXON TRADING LIMITED",
    "city": "NICOSIA",
    "country": "Cyprus"
  },
  "BAE": {
    "code": "BAE",
    "company": "BAE SYSTEMS",
    "city": "FRIDLEY, MN-55421",
    "country": "United States"
  },
  "BAF": {
    "code": "BAF",
    "company": "BULKHAUL LTD",
    "city": "MIDDLESBROUGH CLEVELAND",
    "country": "United Kingdom"
  },
  "BAK": {
    "code": "BAK",
    "company": "BADRAKH ENERGY LLC",
    "city": "Ulaanbaatar",
    "country": "Mongolia"
  },
  "BAL": {
    "code": "BAL",
    "company": "SPEDITION BAUMLE GMBH",
    "city": "Murg",
    "country": "Germany"
  },
  "BAM": {
    "code": "BAM",
    "company": "SPEDITION BAUMLE GMBH",
    "city": "Murg",
    "country": "Germany"
  },
  "BAN": {
    "code": "BAN",
    "company": "SAFBON PARS COMPRESSOR CO.",
    "city": "Tehran",
    "country": "Iran, Islamic Republic of"
  },
  "BAR": {
    "code": "BAR",
    "company": "BARBADOS BOTTLING CO.LTD",
    "city": "BRIDGETOWN",
    "country": "Barbados"
  },
  "BAS": {
    "code": "BAS",
    "company": "BRITISH ANTARCTIC SURVEY",
    "city": "CAMBRIDGE",
    "country": "United Kingdom"
  },
  "BAT": {
    "code": "BAT",
    "company": "CONTAINER BOXES 4 SALE",
    "city": "ROTTEDAM-MAASVLAKTE",
    "country": "Netherlands"
  },
  "BAX": {
    "code": "BAX",
    "company": "SUNMARINE SHIPPING SERVICES LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "BAY": {
    "code": "BAY",
    "company": "CHEMION LOGISTIK GMBH",
    "city": "LEVERKUSEN",
    "country": "Germany"
  },
  "BBB": {
    "code": "BBB",
    "company": "BOSS EQUIPMENT INC.",
    "city": "Roselle Park, NJ 07204",
    "country": "United States"
  },
  "BBC": {
    "code": "BBC",
    "company": "CONTENEDORES Y EMBALAJES NORMALIZADOS S.A",
    "city": "Llanera-Asturias",
    "country": "Spain"
  },
  "BBR": {
    "code": "BBR",
    "company": "ITAKA TRANS",
    "city": "Saint-Petersburg",
    "country": "Russian Federation"
  },
  "BBS": {
    "code": "BBS",
    "company": "LLC COMPANY SPECTRANSCONTAINER",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "BBT": {
    "code": "BBT",
    "company": "BRESS MUNDIAL CONTAINER  SERVICE GMBH",
    "city": "LANGEN",
    "country": "Germany"
  },
  "BBU": {
    "code": "BBU",
    "company": "BSL CONTAINERS LIMITED",
    "city": "KOWLOON",
    "country": "HK"
  },
  "BBX": {
    "code": "BBX",
    "company": "BESTBOX- CONTAINERHANDEL e.K",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "BCC": {
    "code": "BCC",
    "company": "BERTSCHI AG",
    "city": "DUERRENAESCH",
    "country": "Switzerland"
  },
  "BCD": {
    "code": "BCD",
    "company": "BELARUSIAN RAILWAY",
    "city": "MINSK",
    "country": "Belarus"
  },
  "BCE": {
    "code": "BCE",
    "company": "BEAUDIN CONSULTING EQUIPMENT",
    "city": "MONTCLAIR, CA-91763",
    "country": "United States"
  },
  "BCG": {
    "code": "BCG",
    "company": "BLOEDORN CONTAINER GMBH",
    "city": "DORTMUND",
    "country": "Germany"
  },
  "BCH": {
    "code": "BCH",
    "company": "BRAUN CONTAINER HANDELS GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "BCI": {
    "code": "BCI",
    "company": "AL BAHA COMPANY FOR CAUSTIC-CHLORINE INDUSTRY",
    "city": "AL DULAIL",
    "country": "Jordan"
  },
  "BCL": {
    "code": "BCL",
    "company": "BERMUDA CONTAINER LINE LTD",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "BCN": {
    "code": "BCN",
    "company": "BNS CONTAINER AS",
    "city": "OSLO",
    "country": "Norway"
  },
  "BCO": {
    "code": "BCO",
    "company": "N.V. COBO CONSTRUCTION",
    "city": "PARAMARIBO",
    "country": "Suriname"
  },
  "BCX": {
    "code": "BCX",
    "company": "BARCELONA CONTAINER SERVICE DEPOT S.L.",
    "city": "BARCELONA",
    "country": "Spain"
  },
  "BDC": {
    "code": "BDC",
    "company": "BD CONTAINERS BV",
    "city": "AMSTERDAM",
    "country": "Netherlands"
  },
  "BDS": {
    "code": "BDS",
    "company": "BULK DISTRIBUTORS SDN BHD",
    "city": "KLANG",
    "country": "Malaysia"
  },
  "BEA": {
    "code": "BEA",
    "company": "BEACON INTERMODAL LEASING LLC",
    "city": "BOSTON, MA 02199",
    "country": "United States"
  },
  "BEC": {
    "code": "BEC",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "BEG": {
    "code": "BEG",
    "company": "BEC GMBH",
    "city": "Pfullingen",
    "country": "Germany"
  },
  "BEL": {
    "code": "BEL",
    "company": "BELL-MAR SHIPPING LTD",
    "city": "HAIFA",
    "country": "Israel"
  },
  "BEM": {
    "code": "BEM",
    "company": "BASF SE LUDWIGSHAFEN",
    "city": "Ludwigshafen am Rhein",
    "country": "Germany"
  },
  "BER": {
    "code": "BER",
    "company": "BERG MANUFACTURING",
    "city": "SPOKANE, WA 99212",
    "country": "United States"
  },
  "BES": {
    "code": "BES",
    "company": "BESTCHEM GROUP HOLDINGS CO.,LIMITED",
    "city": "WAN CHAI",
    "country": "HK"
  },
  "BEX": {
    "code": "BEX",
    "company": "BLUE EXPRESS INC",
    "city": "SAKAI CITY",
    "country": "Japan"
  },
  "BEZ": {
    "code": "BEZ",
    "company": "YUTEK-M LLC",
    "city": "Severodonetsk",
    "country": "Ukraine"
  },
  "BFC": {
    "code": "BFC",
    "company": "BENNINGER & FÖLL LIQUID LOGISTICS GMBH",
    "city": "ABSTATT",
    "country": "Germany"
  },
  "BFF": {
    "code": "BFF",
    "company": "BAHAMAS FERRIES LTD.",
    "city": "Nassau",
    "country": "Bahamas"
  },
  "BFS": {
    "code": "BFS",
    "company": "H20 INCORPORATED",
    "city": "LAFAYETTE, LA 70508",
    "country": "United States"
  },
  "BFU": {
    "code": "BFU",
    "company": "BEAVERFIT NORTH AMERICA LLC",
    "city": "Reno",
    "country": "United States"
  },
  "BGB": {
    "code": "BGB",
    "company": "BULKTAINER RENTAL AND FINANCE LTD",
    "city": "ST HELIER / JERSEY",
    "country": "United Kingdom"
  },
  "BGF": {
    "code": "BGF",
    "company": "B.G.FREIGHT LINE B.V",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "BGG": {
    "code": "BGG",
    "company": "BRUNO SRL",
    "city": "GROTTAMINARDA",
    "country": "Italy"
  },
  "BGH": {
    "code": "BGH",
    "company": "HRC SHIPPING SDN BHD",
    "city": "KLANG, SELANGOR",
    "country": "Malaysia"
  },
  "BGK": {
    "code": "BGK",
    "company": "BAYER AG",
    "city": "LEVERKUSEN",
    "country": "Germany"
  },
  "BGL": {
    "code": "BGL",
    "company": "POSTEN NORGE AS",
    "city": "OSLO",
    "country": "Norway"
  },
  "BGR": {
    "code": "BGR",
    "company": "BUNDESANSTALT FUR GEOWISSENSCHAFTEN UND ROHSTOFFE",
    "city": "HANNOVER",
    "country": "Germany"
  },
  "BGS": {
    "code": "BGS",
    "company": "BRITISH GEOLOGICAL SURVEY",
    "city": "KEYWORTH, NOTTINGHAM",
    "country": "United Kingdom"
  },
  "BHC": {
    "code": "BHC",
    "company": "BRIDGEHEAD CONTAINER SERVICES LTD",
    "city": "LIVERPOOL",
    "country": "United Kingdom"
  },
  "BHF": {
    "code": "BHF",
    "company": "BEFAR GROUP CO.,LTD.",
    "city": "BINZHOU",
    "country": "China"
  },
  "BHL": {
    "code": "BHL",
    "company": "BIG LIFT SHIPPING",
    "city": "AMSTERDAM",
    "country": "Netherlands"
  },
  "BHS": {
    "code": "BHS",
    "company": "BUZWAIR GASES",
    "city": "Doha",
    "country": "Qatar"
  },
  "BIB": {
    "code": "BIB",
    "company": "CORSYDE INTERNATIONAL GMBH & CO KG",
    "city": "BERLIN",
    "country": "Germany"
  },
  "BID": {
    "code": "BID",
    "company": "BERTSCHI AG",
    "city": "DUERRENAESCH",
    "country": "Switzerland"
  },
  "BIR": {
    "code": "BIR",
    "company": "BUKARA",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "BIS": {
    "code": "BIS",
    "company": "BERMUDA INTERNATIONAL SHIPPING LTD",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "BIT": {
    "code": "BIT",
    "company": "MEEBERG CONTAINER SERVICE BV",
    "city": "MOERDIJK",
    "country": "Netherlands"
  },
  "BJC": {
    "code": "BJC",
    "company": "BOXJOIN CORPORATION",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "BJG": {
    "code": "BJG",
    "company": "GUOGE(BEIJING)ENERGY TECHNOLOGY CO., LTD",
    "city": "BEIJING",
    "country": "China"
  },
  "BJH": {
    "code": "BJH",
    "company": "SHANGHAI BAIJIN CHEMICAL GROUP CO LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "BKB": {
    "code": "BKB",
    "company": "L & T B.V.",
    "city": "'s-Hertogenbosch",
    "country": "Netherlands"
  },
  "BKD": {
    "code": "BKD",
    "company": "JSC CTR OF OP OF SPACE GROUND BASED INF",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "BKG": {
    "code": "BKG",
    "company": "BUNDESAMT FUR KARTOGRAPHIE UND GEODASIE",
    "city": "FRANKFURT AM MAIN",
    "country": "Germany"
  },
  "BKM": {
    "code": "BKM",
    "company": "BARENTS CONTAINER MANAGEMENT LTD",
    "city": "ARKHANGELSK",
    "country": "Russian Federation"
  },
  "BKW": {
    "code": "BKW",
    "company": "ROYAL BOSKALIS WESTMINSTER N.V.",
    "city": "PAPENDRECHT",
    "country": "Netherlands"
  },
  "BLE": {
    "code": "BLE",
    "company": "BELINKA PERKEMIJA  D.O.O.",
    "city": "LJUBLJANA- CR/NUCE",
    "country": "Slovenia"
  },
  "BLI": {
    "code": "BLI",
    "company": "INTERNATIONAL CARGO TERMINALS AND RAIL INFRASTRUCTURE PRIVATE LIMITED",
    "city": "Mumbai",
    "country": "India"
  },
  "BLJ": {
    "code": "BLJ",
    "company": "AVANA GLOBAL FZCO",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "BLK": {
    "code": "BLK",
    "company": "BULKHAUL LTD",
    "city": "MIDDLESBROUGH CLEVELAND",
    "country": "United Kingdom"
  },
  "BLP": {
    "code": "BLP",
    "company": "BALTIC SHIPPING (PVT) LTD",
    "city": "KARACHI",
    "country": "Pakistan"
  },
  "BLR": {
    "code": "BLR",
    "company": "BELINTERTRANS",
    "city": "MINSK",
    "country": "Belarus"
  },
  "BLS": {
    "code": "BLS",
    "company": "ADMIRAL CONTAINER LINES INC LIMITED",
    "city": "VALLETTA",
    "country": "Malta"
  },
  "BLT": {
    "code": "BLT",
    "company": "BALTICON S.A.",
    "city": "GDYNIA",
    "country": "Poland"
  },
  "BLX": {
    "code": "BLX",
    "company": "GSLINES - TRANSPORTES MARITIMOS",
    "city": "FUNCHAL",
    "country": "Portugal"
  },
  "BLZ": {
    "code": "BLZ",
    "company": "BLPL SINGAPORE PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "BML": {
    "code": "BML",
    "company": "KING OCEAN SERVICES LTD",
    "city": "DORAL, FL 33172",
    "country": "United States"
  },
  "BMO": {
    "code": "BMO",
    "company": "BEACON INTERMODAL LEASING LLC",
    "city": "BOSTON, MA 02199",
    "country": "United States"
  },
  "BMP": {
    "code": "BMP",
    "company": "BISMARK MARITIME LTD",
    "city": "PORT MORESBY",
    "country": "Papua New Guinea"
  },
  "BMR": {
    "code": "BMR",
    "company": "BIZNES MASSHTAB LTD.",
    "city": "Saint-Petersburg",
    "country": "Russian Federation"
  },
  "BMU": {
    "code": "BMU",
    "company": "BAUER RESOURCES GMBH",
    "city": "SCHROBENHAUSEN",
    "country": "Germany"
  },
  "BNB": {
    "code": "BNB",
    "company": "TRANSOLVE GLOBAL PTY LTD",
    "city": "SYDNEY",
    "country": "Australia"
  },
  "BND": {
    "code": "BND",
    "company": "EUROBITUMEN LLC",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "BNG": {
    "code": "BNG",
    "company": "SUPERIOR INTERNATIONAL IMPORTS",
    "city": "PORT MOODY",
    "country": "Canada"
  },
  "BNS": {
    "code": "BNS",
    "company": "BNS CONTAINER AS",
    "city": "OSLO",
    "country": "Norway"
  },
  "BOB": {
    "code": "BOB",
    "company": "LANXESS Solutions UK Ltd",
    "city": "MANCHESTER M17 1WT",
    "country": "United Kingdom"
  },
  "BOC": {
    "code": "BOC",
    "company": "LINDE GAS NA LLC",
    "city": "Bridgewater, NJ 08807",
    "country": "United States"
  },
  "BOD": {
    "code": "BOD",
    "company": "SPEDITION BODE GMBH & CO. KG",
    "city": "Reinfeld",
    "country": "Germany"
  },
  "BOH": {
    "code": "BOH",
    "company": "BOH ENVIRONMENTAL , LLC",
    "city": "COVINGTON, LA 70433",
    "country": "United States"
  },
  "BOI": {
    "code": "BOI",
    "company": "ESCOT A/S",
    "city": "AALBORG",
    "country": "Denmark"
  },
  "BOL": {
    "code": "BOL",
    "company": "BOLUDA LINES S.A",
    "city": "VALENCIA",
    "country": "Spain"
  },
  "BON": {
    "code": "BON",
    "company": "BOND INTERNATIONAL GROUP LIMITED",
    "city": "Tortola",
    "country": "Virgin Islands, British"
  },
  "BOO": {
    "code": "BOO",
    "company": "BOBE SPEDITIONS GMBH",
    "city": "BAD SALZUFLEN",
    "country": "Germany"
  },
  "BOR": {
    "code": "BOR",
    "company": "BORCHARD LINES LTD",
    "city": "LONDON EC1Y 4XY",
    "country": "United Kingdom"
  },
  "BOT": {
    "code": "BOT",
    "company": "BULK OIL AND LIQUID TRANSPORT PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "BOX": {
    "code": "BOX",
    "company": "BULKHAUL LTD",
    "city": "MIDDLESBROUGH CLEVELAND",
    "country": "United Kingdom"
  },
  "BPI": {
    "code": "BPI",
    "company": "BULK PIONEER LOGISTICS CO LTD",
    "city": "SHENZHEN",
    "country": "China"
  },
  "BPM": {
    "code": "BPM",
    "company": "BP MC KEEFRY LTD",
    "city": "SWATRAGH",
    "country": "Ireland"
  },
  "BPT": {
    "code": "BPT",
    "company": "BRAID LOGISTICS (UK) LIMITED",
    "city": "RENFREW",
    "country": "United Kingdom"
  },
  "BRB": {
    "code": "BRB",
    "company": "SASOL GERMANY GMBH",
    "city": "BRUNSBUETTEL",
    "country": "Germany"
  },
  "BRC": {
    "code": "BRC",
    "company": "QINHUANGDAO QIN-IN FERRY CO.,LTD",
    "city": "QINHUANGDAO",
    "country": "China"
  },
  "BRD": {
    "code": "BRD",
    "company": "BREDENOORD",
    "city": "Apeldoorn",
    "country": "Netherlands"
  },
  "BRE": {
    "code": "BRE",
    "company": "S.A.T.I. SRL",
    "city": "GIANICO (BS)",
    "country": "Italy"
  },
  "BRF": {
    "code": "BRF",
    "company": "BULKHAUL LTD",
    "city": "MIDDLESBROUGH CLEVELAND",
    "country": "United Kingdom"
  },
  "BRG": {
    "code": "BRG",
    "company": "ABU DHABI POLYMERS COMPANY LTD",
    "city": "ABU DHABI",
    "country": "United Arab Emirates"
  },
  "BRH": {
    "code": "BRH",
    "company": "BRENNTAG NV",
    "city": "Deerlijk-West Vlaanderen",
    "country": "Belgium"
  },
  "BRK": {
    "code": "BRK",
    "company": "BULKHAUL LTD",
    "city": "MIDDLESBROUGH CLEVELAND",
    "country": "United Kingdom"
  },
  "BRN": {
    "code": "BRN",
    "company": "LOCATANK S.A.",
    "city": "EYBENS",
    "country": "France"
  },
  "BRO": {
    "code": "BRO",
    "company": "T.S.B SARL",
    "city": "CASTELJALOUX",
    "country": "France"
  },
  "BRS": {
    "code": "BRS",
    "company": "BASF SE LUDWIGSHAFEN",
    "city": "Ludwigshafen am Rhein",
    "country": "Germany"
  },
  "BRT": {
    "code": "BRT",
    "company": "GROUPE SAMAT UK LTD",
    "city": "Leeds West Yorkshire",
    "country": "United Kingdom"
  },
  "BRZ": {
    "code": "BRZ",
    "company": "CONTAINER EXPRESS CORPORATION",
    "city": "FT LAUDERLALE, FL 33309",
    "country": "United States"
  },
  "BSA": {
    "code": "BSA",
    "company": "BLUE SKY LOGITRADE LIMITED",
    "city": "HONG-KONG",
    "country": "HK"
  },
  "BSB": {
    "code": "BSB",
    "company": "BIGSTEEL BOX CORPORATION",
    "city": "KELOWA, BC",
    "country": "Canada"
  },
  "BSC": {
    "code": "BSC",
    "company": "BESTCONT LTD",
    "city": "LONDON",
    "country": "United Kingdom"
  },
  "BSG": {
    "code": "BSG",
    "company": "SHANGHAI BAOSTEEL GASES CO. LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "BSI": {
    "code": "BSI",
    "company": "BLUE SKY INTERMODAL (UK) LTD",
    "city": "Marlow (Buckinghamshir)",
    "country": "United Kingdom"
  },
  "BSK": {
    "code": "BSK",
    "company": "BODEX-UKRAINE LTD",
    "city": "Kremenchuk",
    "country": "Ukraine"
  },
  "BSL": {
    "code": "BSL",
    "company": "BSL CONTAINERS LIMITED",
    "city": "KOWLOON",
    "country": "HK"
  },
  "BSS": {
    "code": "BSS",
    "company": "BHAVANI SHIPPING SERVICES (I) PVT LTD",
    "city": "CBD BELAPUR",
    "country": "India"
  },
  "BST": {
    "code": "BST",
    "company": "BEST CONTAINERS",
    "city": "PURMEREND",
    "country": "Netherlands"
  },
  "BSX": {
    "code": "BSX",
    "company": "SIX CONSTRUCT LTD. CO.",
    "city": "SHARJAH",
    "country": "United Arab Emirates"
  },
  "BTE": {
    "code": "BTE",
    "company": "BRUHN TRANSPORT EQUIPMENT GMBH & CO",
    "city": "LUBECK",
    "country": "Germany"
  },
  "BTG": {
    "code": "BTG",
    "company": "BRENNTAG UK LIMITED",
    "city": "WIDNES, WA8 0SH",
    "country": "United Kingdom"
  },
  "BTL": {
    "code": "BTL",
    "company": "BULK TAINER LOGISTICS",
    "city": "North Yorkshire",
    "country": "United Kingdom"
  },
  "BTS": {
    "code": "BTS",
    "company": "BTS TANK SOLUTIONS NV",
    "city": "DOTTIGNIES",
    "country": "Belgium"
  },
  "BTT": {
    "code": "BTT",
    "company": "DB CARGO BTT GMBH",
    "city": "MAINZ",
    "country": "Germany"
  },
  "BUC": {
    "code": "BUC",
    "company": "GEORG BUCHNER INTERN. SPEDITION GMBH",
    "city": "FURTH",
    "country": "Germany"
  },
  "BUD": {
    "code": "BUD",
    "company": "CAMECO CORPORATION",
    "city": "Saskatoon",
    "country": "Canada"
  },
  "BUK": {
    "code": "BUK",
    "company": "BULKTAINER RENTAL AND FINANCE LTD",
    "city": "ST HELIER / JERSEY",
    "country": "United Kingdom"
  },
  "BUL": {
    "code": "BUL",
    "company": "BULKGLOBAL LOGISTICS LTD",
    "city": "LONDON NW10 7DY",
    "country": "United Kingdom"
  },
  "BUR": {
    "code": "BUR",
    "company": "BAL CONTAINER LINE CO LIMITED",
    "city": "TSIMSHATSUI",
    "country": "China"
  },
  "BUS": {
    "code": "BUS",
    "company": "H. BUTEFUHR & SOHN GMBH & CO KG",
    "city": "DUISBURG",
    "country": "Germany"
  },
  "BUT": {
    "code": "BUT",
    "company": "BROINTERMED LINES LIMITED",
    "city": "HARWICH, CO12 3HH",
    "country": "United Kingdom"
  },
  "BUZ": {
    "code": "BUZ",
    "company": "BUZZATTI TRASPORTI SRL",
    "city": "SEDICO  (BL)",
    "country": "Italy"
  },
  "BVH": {
    "code": "BVH",
    "company": "CRYOGENIC CONTAINER SOLUTIONS B.V.",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "BVI": {
    "code": "BVI",
    "company": "STOLT TANK CONTAINERS LEASING LTD",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "BVN": {
    "code": "BVN",
    "company": "BENEVENTI SRL",
    "city": "SASSUOLO (MO)",
    "country": "Italy"
  },
  "BWC": {
    "code": "BWC",
    "company": "BAY LOGISTIK GMBH & CO KG",
    "city": "WAIBLINGEN",
    "country": "Germany"
  },
  "BWL": {
    "code": "BWL",
    "company": "BLUE WATER LINES PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "BWP": {
    "code": "BWP",
    "company": "LOGISTIKZENTRUM DER BUNDESWEHR",
    "city": "WILHELMSHAVEN",
    "country": "Germany"
  },
  "BWT": {
    "code": "BWT",
    "company": "DAMEN GREEN SOLUTIONS B.V",
    "city": "GORINCHEM",
    "country": "Netherlands"
  },
  "BXD": {
    "code": "BXD",
    "company": "BOXDIRECT AG",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "BXL": {
    "code": "BXL",
    "company": "SICOM S.P.A",
    "city": "CHERASCO (CN)",
    "country": "Italy"
  },
  "BXN": {
    "code": "BXN",
    "company": "BOXMAN ALPHA LTD",
    "city": "Nelson",
    "country": "New Zealand"
  },
  "BXR": {
    "code": "BXR",
    "company": "TITAN CONTAINERS A/S",
    "city": "TAASTRUP",
    "country": "Denmark"
  },
  "BXS": {
    "code": "BXS",
    "company": "BLUE EXPRESS (SHANGHAI) INTL FREIGHT FORWARDING CO.LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "BYM": {
    "code": "BYM",
    "company": "PRIVATE LIMITED COMPANY OU RODELINDA",
    "city": "TALLINN",
    "country": "Estonia"
  },
  "BYR": {
    "code": "BYR",
    "company": "BYRNE EQUIPMENT RENTAL LLC",
    "city": "Dubai",
    "country": "United Arab Emirates"
  },
  "BZT": {
    "code": "BZT",
    "company": "BOX2TRADE B.V",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "CAA": {
    "code": "CAA",
    "company": "CAI INTERNATIONAL",
    "city": "SAN FRANCISCO, CA 94105",
    "country": "United States"
  },
  "CAB": {
    "code": "CAB",
    "company": "CHAKIAT AGENCIES PVT LTD",
    "city": "CHENNAI",
    "country": "India"
  },
  "CAD": {
    "code": "CAD",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "CAE": {
    "code": "CAE",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "CAF": {
    "code": "CAF",
    "company": "CISAF SRL",
    "city": "BERGAMO",
    "country": "Italy"
  },
  "CAI": {
    "code": "CAI",
    "company": "CAI INTERNATIONAL",
    "city": "SAN FRANCISCO, CA 94105",
    "country": "United States"
  },
  "CAK": {
    "code": "CAK",
    "company": "CAMELLIA LINE CO LTD",
    "city": "FUKUOKA CITY",
    "country": "Japan"
  },
  "CAM": {
    "code": "CAM",
    "company": "MGC PURE CHEMICALS AMERICA",
    "city": "MESA,ARIZONA 85212",
    "country": "United States"
  },
  "CAN": {
    "code": "CAN",
    "company": "CHIMICA SARDA SRL",
    "city": "Sassari",
    "country": "Italy"
  },
  "CAP": {
    "code": "CAP",
    "company": "CAPE CONTAINER LINE (UK) LTD",
    "city": "LONDON",
    "country": "United Kingdom"
  },
  "CAR": {
    "code": "CAR",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "CAS": {
    "code": "CAS",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "CAT": {
    "code": "CAT",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "CAU": {
    "code": "CAU",
    "company": "CAMPAS CONTAINERS",
    "city": "Viby J",
    "country": "Denmark"
  },
  "CAV": {
    "code": "CAV",
    "company": "AELER TECHNOLOGIES SA",
    "city": "Lausanne",
    "country": "Switzerland"
  },
  "CAW": {
    "code": "CAW",
    "company": "CONWAY CONTAINER SOLUTIONS SIA",
    "city": "Riga",
    "country": "Latvia"
  },
  "CAX": {
    "code": "CAX",
    "company": "CAI INTERNATIONAL",
    "city": "SAN FRANCISCO, CA 94105",
    "country": "United States"
  },
  "CAZ": {
    "code": "CAZ",
    "company": "CAI INTERNATIONAL",
    "city": "SAN FRANCISCO, CA 94105",
    "country": "United States"
  },
  "CBC": {
    "code": "CBC",
    "company": "CONTAINER BROKERAGE COMPANY LTD",
    "city": "CHALFONT ST PETER",
    "country": "United Kingdom"
  },
  "CBH": {
    "code": "CBH",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "CBK": {
    "code": "CBK",
    "company": "BELARUSKALI JSC",
    "city": "SOLIGORSK, MINSK REGION",
    "country": "Belarus"
  },
  "CBL": {
    "code": "CBL",
    "company": "CBOX CONTAINERS EQUIPMENT BV",
    "city": "Amsterdam",
    "country": "Netherlands"
  },
  "CBM": {
    "code": "CBM",
    "company": "CONTAINER BEST LTD",
    "city": "Saint Petersbourg",
    "country": "Russian Federation"
  },
  "CBN": {
    "code": "CBN",
    "company": "CUBNER SAS",
    "city": "PERIGUEUX",
    "country": "France"
  },
  "CBO": {
    "code": "CBO",
    "company": "N.V. COBO CONSTRUCTION",
    "city": "PARAMARIBO",
    "country": "Suriname"
  },
  "CBP": {
    "code": "CBP",
    "company": "CBSL TRANSPORTATION SERVICES INC",
    "city": "CHICAGO, IL-60638",
    "country": "United States"
  },
  "CBR": {
    "code": "CBR",
    "company": "CANBERRA CONTAINERS",
    "city": "KAMBAH",
    "country": "Australia"
  },
  "CBT": {
    "code": "CBT",
    "company": "CONT-ASPHALT LTD.",
    "city": "LA CROIX (LUTRY)",
    "country": "Switzerland"
  },
  "CBX": {
    "code": "CBX",
    "company": "CBOX CONTAINERS EQUIPMENT BV",
    "city": "Amsterdam",
    "country": "Netherlands"
  },
  "CCC": {
    "code": "CCC",
    "company": "CONRAIL CONTAINER GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "CCD": {
    "code": "CCD",
    "company": "CONRO CONTAINER GMBH",
    "city": "SEEVETAL",
    "country": "Germany"
  },
  "CCE": {
    "code": "CCE",
    "company": "JUNTAI CONTAINER COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "CCG": {
    "code": "CCG",
    "company": "CATU CONTAINERS SA",
    "city": "CAROUGE-GENEVE",
    "country": "Switzerland"
  },
  "CCH": {
    "code": "CCH",
    "company": "YARA TERTRE S.A",
    "city": "TERTRE",
    "country": "Belgium"
  },
  "CCK": {
    "code": "CCK",
    "company": "CENTRAL CONTAINER LTD",
    "city": "Martonvásár",
    "country": "Hungary"
  },
  "CCL": {
    "code": "CCL",
    "company": "COSCO SHIPPING DEVELOPMENT (ASIA) CO.,LTD",
    "city": "Kwai Chung, New Territories",
    "country": "HK"
  },
  "CCM": {
    "code": "CCM",
    "company": "MERCITALIA RAIL",
    "city": "Rome",
    "country": "Italy"
  },
  "CCO": {
    "code": "CCO",
    "company": "A-GAS AMERICAS (REFRIGERANTS & CARBON OFFSETS)",
    "city": "RHOME, TX-76078",
    "country": "United States"
  },
  "CCP": {
    "code": "CCP",
    "company": "YOU FIRST EXPRESS DBA CONTAINER CHASSIS AND PARTS",
    "city": "Los Angeles",
    "country": "United States"
  },
  "CCR": {
    "code": "CCR",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "CCS": {
    "code": "CCS",
    "company": "SERVIAL CC",
    "city": "HELSINGE",
    "country": "Denmark"
  },
  "CCT": {
    "code": "CCT",
    "company": "CRAIGS CONTAINER COMPANY",
    "city": "LONDON SW10 9UH",
    "country": "United Kingdom"
  },
  "CCU": {
    "code": "CCU",
    "company": "M/S HIKARU SHIPPING LINE DMCC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "CCX": {
    "code": "CCX",
    "company": "CCR CONTAINERS SAS",
    "city": "NEUILLY SUR SEINE CEDEX",
    "country": "France"
  },
  "CCY": {
    "code": "CCY",
    "company": "CHIMI ART (CYPRUS) LTD",
    "city": "Nicosia",
    "country": "Cyprus"
  },
  "CDA": {
    "code": "CDA",
    "company": "CDA LOJISTIK A.S.",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "CDD": {
    "code": "CDD",
    "company": "CONTAINER PROVIDERS INTL (BVI) LTD",
    "city": "TORTOLA",
    "country": "Virgin Islands, British"
  },
  "CDK": {
    "code": "CDK",
    "company": "CONTAINERLAND DANMARK",
    "city": "HVIDOVRE",
    "country": "Denmark"
  },
  "CDM": {
    "code": "CDM",
    "company": "CONTAINER DEPOT MUNCHEN GMBH",
    "city": "UNTERFOEHRING",
    "country": "Germany"
  },
  "CDN": {
    "code": "CDN",
    "company": "COMPAGNIE DES ILES DU NORD",
    "city": "SAINT BARTHELEMY",
    "country": "France"
  },
  "CDS": {
    "code": "CDS",
    "company": "CIRQUE DU SOLEIL INC.",
    "city": "Montreal, H1Z 4M6",
    "country": "Canada"
  },
  "CEA": {
    "code": "CEA",
    "company": "ORANO TN INTERNATIONAL",
    "city": "VALOGNES",
    "country": "France"
  },
  "CEC": {
    "code": "CEC",
    "company": "CCR CONTAINERS SAS",
    "city": "NEUILLY SUR SEINE CEDEX",
    "country": "France"
  },
  "CEF": {
    "code": "CEF",
    "company": "SOLAR ENERGY FACTORY CO LTD",
    "city": "MAIZURU, KYOTO",
    "country": "Japan"
  },
  "CEI": {
    "code": "CEI",
    "company": "CONTENEURS EXPERTS INC",
    "city": "VAUDREUIL-DORION",
    "country": "Canada"
  },
  "CEM": {
    "code": "CEM",
    "company": "CHEMION LOGISTIK GMBH",
    "city": "LEVERKUSEN",
    "country": "Germany"
  },
  "CEN": {
    "code": "CEN",
    "company": "CUADROS ELECTRICOS NAZARENOS, S.L.",
    "city": "Sevilla",
    "country": "Spain"
  },
  "CEO": {
    "code": "CEO",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LIMITED",
    "city": "SAN FRANCISCO, CA 94108",
    "country": "United States"
  },
  "CEP": {
    "code": "CEP",
    "company": "C.E.P S.R.L",
    "city": "CALATAFIMI SEGESTA (TP)",
    "country": "Italy"
  },
  "CER": {
    "code": "CER",
    "company": "RIVERSIDE RESOURCE RECOVERY LTD",
    "city": "DAI7 6JY",
    "country": "United Kingdom"
  },
  "CES": {
    "code": "CES",
    "company": "CAI INTERNATIONAL",
    "city": "SAN FRANCISCO, CA 94105",
    "country": "United States"
  },
  "CET": {
    "code": "CET",
    "company": "CAREL S.A.",
    "city": "PIRAEUS",
    "country": "Greece"
  },
  "CEU": {
    "code": "CEU",
    "company": "CHINA EASTERN CONTAINERS LIMITED",
    "city": "HONG KONG",
    "country": "HK"
  },
  "CEX": {
    "code": "CEX",
    "company": "CHEMICAL EXPRESS SRL",
    "city": "NAPOLI   NA",
    "country": "Italy"
  },
  "CFB": {
    "code": "CFB",
    "company": "CHEMISCHE FABRIK KARL BUCHER GMBH",
    "city": "Waldstetten",
    "country": "Germany"
  },
  "CFC": {
    "code": "CFC",
    "company": "DEPARTMENT OF NATIONAL DEFENCE (DND)",
    "city": "OTTAWA, ONTARIO  K1A OK2",
    "country": "Canada"
  },
  "CFH": {
    "code": "CFH",
    "company": "SHANDONG CONGLIN FRUEHAUF AUTOMOBILE CO.,LTD",
    "city": "Longkou",
    "country": "China"
  },
  "CFI": {
    "code": "CFI",
    "company": "CFI INTERMODAL SRL",
    "city": "roma",
    "country": "Italy"
  },
  "CFL": {
    "code": "CFL",
    "company": "CRYOFLEET PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "CFM": {
    "code": "CFM",
    "company": "STATE ENTERPRISE RAILWAYS OF MOLDOVA",
    "city": "KISHINEV",
    "country": "Moldova, Republic of"
  },
  "CFR": {
    "code": "CFR",
    "company": "CF&S RUSSIA LTD",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "CFX": {
    "code": "CFX",
    "company": "PRAXAIR COSTA RICA S.A.",
    "city": "SAN JOSE",
    "country": "Costa Rica"
  },
  "CGA": {
    "code": "CGA",
    "company": "CUSTOMISED GAS AUSTRALIA GROUP PTY LTD",
    "city": "MULGRAVE NSW",
    "country": "Australia"
  },
  "CGB": {
    "code": "CGB",
    "company": "COREGAS PTY LTD",
    "city": "YENNORA",
    "country": "Australia"
  },
  "CGC": {
    "code": "CGC",
    "company": "C.A.B. COOPERATIVA AUTOTRASPORTATORI SOC.COOP",
    "city": "PERUGIA",
    "country": "Italy"
  },
  "CGG": {
    "code": "CGG",
    "company": "SEABED GEOSOLUTIONS AS",
    "city": "LAKSEVAAG",
    "country": "Norway"
  },
  "CGH": {
    "code": "CGH",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "CGI": {
    "code": "CGI",
    "company": "CONGLOBAL INDUSTRIES",
    "city": "SAN RAMON , CA 94583",
    "country": "United States"
  },
  "CGK": {
    "code": "CGK",
    "company": "CGK INTERNATIONAL LIMITED",
    "city": "HONG KONG",
    "country": "HK"
  },
  "CGL": {
    "code": "CGL",
    "company": "MILLENNIUM INORGANIC CHEMICALS THANN SAS",
    "city": "THANN",
    "country": "France"
  },
  "CGM": {
    "code": "CGM",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "CGO": {
    "code": "CGO",
    "company": "AIR LIQUIDE CONGO",
    "city": "POINTE NOIRE",
    "country": "Congo"
  },
  "CGR": {
    "code": "CGR",
    "company": "CARGOR BIZKAIA S.L.",
    "city": "ZIERBENA",
    "country": "Spain"
  },
  "CGS": {
    "code": "CGS",
    "company": "SAPIO PRODUZIONE IDROGENO OSSIGENO SRL",
    "city": "MONZA (MB)",
    "country": "Italy"
  },
  "CGT": {
    "code": "CGT",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "CHC": {
    "code": "CHC",
    "company": "CHALLENGE CONTAINER LOGISTIC INC",
    "city": "GUAYAQUIL",
    "country": "Ecuador"
  },
  "CHD": {
    "code": "CHD",
    "company": "CHEMINOVA  A/S",
    "city": "Harboöre",
    "country": "Denmark"
  },
  "CHE": {
    "code": "CHE",
    "company": "LLC CHEMRESURS",
    "city": "Yaroslavl",
    "country": "Russian Federation"
  },
  "CHH": {
    "code": "CHH",
    "company": "CLH CONTAINER LOGISTIC HAMBURG GMBH & CO KG",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "CHI": {
    "code": "CHI",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "CHL": {
    "code": "CHL",
    "company": "MARGUISA SHIPPING LINES, S.L.U.",
    "city": "MADRID",
    "country": "Spain"
  },
  "CHM": {
    "code": "CHM",
    "company": "ALBEMARLE GERMANY GMBH",
    "city": "LANGELSHEIM",
    "country": "Germany"
  },
  "CHQ": {
    "code": "CHQ",
    "company": "CHEMBRO QUIMICA LTDA",
    "city": "VITORIA-ES",
    "country": "Brazil"
  },
  "CHR": {
    "code": "CHR",
    "company": "CHEMOURS COMPANY FC LLC",
    "city": "WILMINGTON, DE-19898",
    "country": "United States"
  },
  "CHS": {
    "code": "CHS",
    "company": "CHS CONTAINER HANDEL GMBH",
    "city": "BREMEN",
    "country": "Germany"
  },
  "CHT": {
    "code": "CHT",
    "company": "COHERENT BOX INTERNATIONAL CO., LTD",
    "city": "Hong-Kong",
    "country": "HK"
  },
  "CHV": {
    "code": "CHV",
    "company": "CHV CONTAINERHANDELS UND VERMIETUNGSGES MBH",
    "city": "WIEN",
    "country": "Austria"
  },
  "CHX": {
    "code": "CHX",
    "company": "CHEMCO AS",
    "city": "STRUSSHAMN",
    "country": "Norway"
  },
  "CIA": {
    "code": "CIA",
    "company": "GEODIS RT ITALIA SRL",
    "city": "CORMANO",
    "country": "Italy"
  },
  "CIB": {
    "code": "CIB",
    "company": "CONT-ASPHALT LTD.",
    "city": "LA CROIX (LUTRY)",
    "country": "Switzerland"
  },
  "CIC": {
    "code": "CIC",
    "company": "CIMC CONTAINERS HOLDING COMPANY LTD",
    "city": "GUANGDONG",
    "country": "China"
  },
  "CIG": {
    "code": "CIG",
    "company": "EDF",
    "city": "MONTEVRAIN",
    "country": "France"
  },
  "CIL": {
    "code": "CIL",
    "company": "CARGOSTORE WORLDWIDE TRADING LIMITED",
    "city": "LONDON SW19 7QD",
    "country": "United Kingdom"
  },
  "CIO": {
    "code": "CIO",
    "company": "CENTRO INTERNAZIONALE STUDI CONTAINERS (C.I.S.CO.)",
    "city": "GENOA",
    "country": "Italy"
  },
  "CIP": {
    "code": "CIP",
    "company": "CAI INTERNATIONAL",
    "city": "SAN FRANCISCO, CA 94105",
    "country": "United States"
  },
  "CIS": {
    "code": "CIS",
    "company": "CON.S.AR. SOC. COOP. CONS.",
    "city": "RAVENNA RA",
    "country": "Italy"
  },
  "CJK": {
    "code": "CJK",
    "company": "CJ KOREA EXPRESS",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "CJT": {
    "code": "CJT",
    "company": "L & T B.V.",
    "city": "'s-Hertogenbosch",
    "country": "Netherlands"
  },
  "CJY": {
    "code": "CJY",
    "company": "CHONGQING TRANSPORTATION HOLDING (GROUP) CO,LTD",
    "city": "CHONGQING",
    "country": "China"
  },
  "CKC": {
    "code": "CKC",
    "company": "SOLSINO LTD",
    "city": "HONG KONG",
    "country": "HK"
  },
  "CKE": {
    "code": "CKE",
    "company": "CELERITY TANK LOGISTICS (CHINA) LTD",
    "city": "HONG KONG",
    "country": "HK"
  },
  "CKF": {
    "code": "CKF",
    "company": "LIANYUNGANG C-K FERRY CO LTD",
    "city": "LIANYUNGANG CITY",
    "country": "China"
  },
  "CKH": {
    "code": "CKH",
    "company": "CLEANCOR ENERGY SOLUTIONS LLC",
    "city": "New York",
    "country": "United States"
  },
  "CKL": {
    "code": "CKL",
    "company": "CEEKAY SHIPPING AND MARINE SERVICES LTD",
    "city": "MUMBAI",
    "country": "India"
  },
  "CKN": {
    "code": "CKN",
    "company": "KNAUF GIPS KG",
    "city": "IPHOFEN",
    "country": "Germany"
  },
  "CKR": {
    "code": "CKR",
    "company": "WANGYU INVESTMENT CORPORATION",
    "city": "Naperville",
    "country": "United States"
  },
  "CKS": {
    "code": "CKS",
    "company": "CK LINE CO LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "CKT": {
    "code": "CKT",
    "company": "COOL WORLD NEDERLAND BV",
    "city": "WAALWIJK",
    "country": "Netherlands"
  },
  "CKX": {
    "code": "CKX",
    "company": "CAKEBOXX TECHNOLOGIES,LLC",
    "city": "MCLEAN, VA 22101",
    "country": "United States"
  },
  "CLC": {
    "code": "CLC",
    "company": "CARU PRAHA S.R.O",
    "city": "PRAHA",
    "country": "Czech Republic"
  },
  "CLD": {
    "code": "CLD",
    "company": "CLDN CARGO NV",
    "city": "ZEEBRUGGE",
    "country": "Belgium"
  },
  "CLE": {
    "code": "CLE",
    "company": "CLEAN HARBORS ENVIRONMENTAL SERVICES, INC",
    "city": "Norwell, MA 02061",
    "country": "United States"
  },
  "CLH": {
    "code": "CLH",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "CLI": {
    "code": "CLI",
    "company": "CLBT S.R.L",
    "city": "BENTIVOGLIO - BO",
    "country": "Italy"
  },
  "CLJ": {
    "code": "CLJ",
    "company": "CLK TRANSPORT AND TRADING FZE",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "CLK": {
    "code": "CLK",
    "company": "CITY AIRPORT, LOGIS&TRAVEL",
    "city": "Seoul",
    "country": "Korea, Republic of"
  },
  "CLL": {
    "code": "CLL",
    "company": "CASSILON LIMITED",
    "city": "ESSEX CO7 6RB",
    "country": "United Kingdom"
  },
  "CLO": {
    "code": "CLO",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "CLP": {
    "code": "CLP",
    "company": "CARGOSOL LOGISTICS PVT.LTD.",
    "city": "Mumbai",
    "country": "India"
  },
  "CLQ": {
    "code": "CLQ",
    "company": "CLIQ CONTAINER TRADING LIMITED",
    "city": "Edinburgh",
    "country": "United Kingdom"
  },
  "CLR": {
    "code": "CLR",
    "company": "C.L.T. SOC. COOP. A R.L.",
    "city": "Ravenna",
    "country": "Italy"
  },
  "CLS": {
    "code": "CLS",
    "company": "CARAVEL LOGISTICS PRIVATE LIMITED",
    "city": "CHENNAI",
    "country": "India"
  },
  "CLT": {
    "code": "CLT",
    "company": "SOGECO INTERNATIONAL SA",
    "city": "PARADISO - LUGANO",
    "country": "Switzerland"
  },
  "CLV": {
    "code": "CLV",
    "company": "CLEVELAND CONTAINERS",
    "city": "STOCKTON ON TEES",
    "country": "United Kingdom"
  },
  "CLW": {
    "code": "CLW",
    "company": "CONLOG OY",
    "city": "YLIKIIMINKI",
    "country": "Finland"
  },
  "CLX": {
    "code": "CLX",
    "company": "AMPLE HARVEST CONTAINER FUND SPC",
    "city": "GRAND CAYMAN",
    "country": "Cayman Islands"
  },
  "CMA": {
    "code": "CMA",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "CMC": {
    "code": "CMC",
    "company": "CROWLEY LINER SERVICES INC",
    "city": "JACKSONVILLE, FL 32225",
    "country": "United States"
  },
  "CME": {
    "code": "CME",
    "company": "SPARTANK SP. Z.O.O",
    "city": "WROCLAW",
    "country": "Poland"
  },
  "CMG": {
    "code": "CMG",
    "company": "CEMENGAL SA",
    "city": "MADRID",
    "country": "Spain"
  },
  "CMI": {
    "code": "CMI",
    "company": "INTERMOBILE CONTAINER, LLC",
    "city": "Las Vegas, NV 89101",
    "country": "United States"
  },
  "CML": {
    "code": "CML",
    "company": "JAEBUM CO.,LTD.",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "CMN": {
    "code": "CMN",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "CMO": {
    "code": "CMO",
    "company": "LLC CRYOR",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "CMR": {
    "code": "CMR",
    "company": "AIR LIQUIDE CAMEROUN",
    "city": "DOUALA",
    "country": "Cameroon"
  },
  "CMT": {
    "code": "CMT",
    "company": "IJ-CONTAINER APS",
    "city": "GENTOFTE",
    "country": "Denmark"
  },
  "CMU": {
    "code": "CMU",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "CMY": {
    "code": "CMY",
    "company": "CRYOTECH LOGISTICS SDN BHD",
    "city": "SHAH ALAM, SELANGOR",
    "country": "Malaysia"
  },
  "CMZ": {
    "code": "CMZ",
    "company": "CORNELDER DE MOZAMBIQUE S.A",
    "city": "PO BOX 236 - BEIRA",
    "country": "Mozambique"
  },
  "CNA": {
    "code": "CNA",
    "company": "TECHNIC-FRANCE",
    "city": "SAINT-DENIS LA PLAINE",
    "country": "France"
  },
  "CNC": {
    "code": "CNC",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "CNE": {
    "code": "CNE",
    "company": "CAI INTERNATIONAL",
    "city": "SAN FRANCISCO, CA 94105",
    "country": "United States"
  },
  "CNI": {
    "code": "CNI",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "CNS": {
    "code": "CNS",
    "company": "SPINNAKER LEASING CORPORATION",
    "city": "San Francisco, CA 94111-2602",
    "country": "United States"
  },
  "CNW": {
    "code": "CNW",
    "company": "BABRA-POL SZYMON RZEPECKI",
    "city": "Radzymin",
    "country": "Poland"
  },
  "CNX": {
    "code": "CNX",
    "company": "IFREMER",
    "city": "Issy les Moulineaux",
    "country": "France"
  },
  "COC": {
    "code": "COC",
    "company": "SAMSKIP MCL BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "COD": {
    "code": "COD",
    "company": "CROSSOVER ASSET MANAGEMENT (SINGAPORE) PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "COG": {
    "code": "COG",
    "company": "ORANO TN INTERNATIONAL",
    "city": "VALOGNES",
    "country": "France"
  },
  "COM": {
    "code": "COM",
    "company": "COMMONWEALTH STEEL COMPANY PTY LTD",
    "city": "WARATAH",
    "country": "Australia"
  },
  "CON": {
    "code": "CON",
    "company": "CONTAINER OF AMERICA LLC",
    "city": "FT.LAUDERDALE, FL 33308",
    "country": "United States"
  },
  "COO": {
    "code": "COO",
    "company": "CO-OPERATOR CONTAINER TRANSPORT & LOGISTIC GMBH",
    "city": "BREMEN",
    "country": "Germany"
  },
  "COP": {
    "code": "COP",
    "company": "MAXX SERVICES LIMITED",
    "city": "LIMASSOL",
    "country": "Cyprus"
  },
  "COQ": {
    "code": "COQ",
    "company": "CONCISA",
    "city": "MADRID",
    "country": "Spain"
  },
  "COR": {
    "code": "COR",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "COX": {
    "code": "COX",
    "company": "CEYLON OXYGEN LIMITED",
    "city": "Colombo",
    "country": "Sri Lanka"
  },
  "COY": {
    "code": "COY",
    "company": "CONTAINENTAL LTD",
    "city": "RICHMOND SURREY",
    "country": "United Kingdom"
  },
  "COZ": {
    "code": "COZ",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "CPA": {
    "code": "CPA",
    "company": "COMBIPASS SAS",
    "city": "AVIGNON CEDEX 09",
    "country": "France"
  },
  "CPB": {
    "code": "CPB",
    "company": "CANONS PARK EQUIPMENT CO LTD",
    "city": "TST",
    "country": "HK"
  },
  "CPD": {
    "code": "CPD",
    "company": "IES DOWNSTREAM LLC",
    "city": "KAPOLEI, HI 96707-1807",
    "country": "United States"
  },
  "CPF": {
    "code": "CPF",
    "company": "CONSERVA SPA",
    "city": "BITONTO",
    "country": "Italy"
  },
  "CPG": {
    "code": "CPG",
    "company": "CARRIER REEFERS & GENSETS BV DIVISION OF CARRIER TRANSICOLT",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "CPH": {
    "code": "CPH",
    "company": "CHR HANSEN A/S",
    "city": "HOERSHOLM",
    "country": "Denmark"
  },
  "CPI": {
    "code": "CPI",
    "company": "CONTAINER PROVIDERS OMCC LTD",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "CPL": {
    "code": "CPL",
    "company": "CONTAINERPLUS BV",
    "city": "RUCPHEN",
    "country": "Netherlands"
  },
  "CPP": {
    "code": "CPP",
    "company": "CANADIAN PACIFIC RAILWAY COMPANY",
    "city": "MISSISSAUGA, ONT L5C 4R3",
    "country": "Canada"
  },
  "CPS": {
    "code": "CPS",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "CPT": {
    "code": "CPT",
    "company": "TK SRT",
    "city": "Krasnoyarsk",
    "country": "Russian Federation"
  },
  "CPW": {
    "code": "CPW",
    "company": "CONTAINER PROVIDERS INTL (BVI) LTD",
    "city": "TORTOLA",
    "country": "Virgin Islands, British"
  },
  "CRA": {
    "code": "CRA",
    "company": "CARBOX CO-DOS DE COSTA RICA",
    "city": "ALAJUELA",
    "country": "Costa Rica"
  },
  "CRB": {
    "code": "CRB",
    "company": "CHINA RAILWAY TIELONG CONTAINER LOGISTIC CO LTD DALIAN LOGISTIC BRANCH",
    "city": "DALIAN",
    "country": "China"
  },
  "CRD": {
    "code": "CRD",
    "company": "CARDO FACILITY AND LOGISTIC SERVICES GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "CRF": {
    "code": "CRF",
    "company": "C.R. TECHNOLOGY SYSTEMS S.P.A.",
    "city": "Roma",
    "country": "Italy"
  },
  "CRG": {
    "code": "CRG",
    "company": "CARGOTAINER GMBH",
    "city": "MANNHEIM",
    "country": "Germany"
  },
  "CRH": {
    "code": "CRH",
    "company": "CROS CONSTRUCT SRL",
    "city": "Bucharest",
    "country": "Romania"
  },
  "CRJ": {
    "code": "CRJ",
    "company": "CHINA RAILWAY INTERNATIONAL MULTIMODAL TRANSPORT CO.LTD",
    "city": "BEIJING",
    "country": "China"
  },
  "CRK": {
    "code": "CRK",
    "company": "CURT RICHTER SE",
    "city": "KOELN",
    "country": "Germany"
  },
  "CRL": {
    "code": "CRL",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "CRM": {
    "code": "CRM",
    "company": "ACE ENGINEERING CO LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "CRN": {
    "code": "CRN",
    "company": "PT INDONESIA RUIPI NICKEL AND CHROME ALLOY",
    "city": "JAKARTA",
    "country": "Indonesia"
  },
  "CRO": {
    "code": "CRO",
    "company": "CRYOS CO., LTD.",
    "city": "Busan",
    "country": "Korea, Republic of"
  },
  "CRP": {
    "code": "CRP",
    "company": "CIMAR RENTING SA",
    "city": "VILA DO CONDE",
    "country": "Portugal"
  },
  "CRR": {
    "code": "CRR",
    "company": "CRIMSON SHIPPING CO. INC.",
    "city": "CHICKASAW, AL-36611",
    "country": "United States"
  },
  "CRS": {
    "code": "CRS",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "CRT": {
    "code": "CRT",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "CRU": {
    "code": "CRU",
    "company": "CR CONTAINER TRADING GMBH",
    "city": "Hamburg",
    "country": "Germany"
  },
  "CRX": {
    "code": "CRX",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "CRY": {
    "code": "CRY",
    "company": "CRYOTAINER BV",
    "city": "SCHIEDAM",
    "country": "Netherlands"
  },
  "CSA": {
    "code": "CSA",
    "company": "CORPORATIVO DE SERVICIOS AMBIENTALES S.A DE C.V",
    "city": "VILLAHERMOSA, TABASCO",
    "country": "Mexico"
  },
  "CSC": {
    "code": "CSC",
    "company": "CONTAINER SOLUTIONS CO LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "CSF": {
    "code": "CSF",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "CSH": {
    "code": "CSH",
    "company": "CONTAINER SALES & HIRE LTD",
    "city": "SUNDERLAND, SR4 6SJ",
    "country": "United Kingdom"
  },
  "CSI": {
    "code": "CSI",
    "company": "CSI GROUP LLC",
    "city": "QUINCY MA-02169",
    "country": "United States"
  },
  "CSL": {
    "code": "CSL",
    "company": "COSCO SHIPPING DEVELOPMENT (ASIA) CO.,LTD",
    "city": "Kwai Chung, New Territories",
    "country": "HK"
  },
  "CSN": {
    "code": "CSN",
    "company": "COSCO SHIPPING LINES CO LTD",
    "city": "Shanghai",
    "country": "China"
  },
  "CSO": {
    "code": "CSO",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "CSP": {
    "code": "CSP",
    "company": "CONTAINERSPOT A/S",
    "city": "RISSKOV",
    "country": "Denmark"
  },
  "CSQ": {
    "code": "CSQ",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "CSR": {
    "code": "CSR",
    "company": "CRYO SERVICE SRL",
    "city": "SOMAGLIA (LO)",
    "country": "Italy"
  },
  "CSS": {
    "code": "CSS",
    "company": "CONTAINER SALES & HIRE LTD",
    "city": "SUNDERLAND, SR4 6SJ",
    "country": "United Kingdom"
  },
  "CST": {
    "code": "CST",
    "company": "CONTAINER-SERVICE FRIEDRICH TIEMANN GMBH",
    "city": "BREMERHAVEN",
    "country": "Germany"
  },
  "CSU": {
    "code": "CSU",
    "company": "CONTAINER SPEDITIONS UND TRANSPORT GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "CSV": {
    "code": "CSV",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "CSW": {
    "code": "CSW",
    "company": "3M KOREA HIGH TECH",
    "city": "Naju",
    "country": "Korea, Republic of"
  },
  "CSZ": {
    "code": "CSZ",
    "company": "CONTAINER STORAGE",
    "city": "Reno, NV 89521",
    "country": "United States"
  },
  "CTA": {
    "code": "CTA",
    "company": "COOLTAINER NEW ZEALAND LIMITED",
    "city": "CHRISTCHURCH",
    "country": "New Zealand"
  },
  "CTB": {
    "code": "CTB",
    "company": "CONTAINER TRADER LTD.",
    "city": "VARNA",
    "country": "Bulgaria"
  },
  "CTC": {
    "code": "CTC",
    "company": "TERMCOTANK SA",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "CTF": {
    "code": "CTF",
    "company": "CON.A.P. SCRL",
    "city": "FIORENZUOLA D'ARDA PC",
    "country": "Italy"
  },
  "CTH": {
    "code": "CTH",
    "company": "CHEMICAL TRANSFER CO. INC.",
    "city": "STOCKTON, CA 95206",
    "country": "United States"
  },
  "CTI": {
    "code": "CTI",
    "company": "CHLORITECH INDUSTRIES",
    "city": "Vadodara",
    "country": "India"
  },
  "CTK": {
    "code": "CTK",
    "company": "CONTANK S.A.",
    "city": "CASTELLBISBAL (BARCELONA)",
    "country": "Spain"
  },
  "CTL": {
    "code": "CTL",
    "company": "CTL-ICS B.V.",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "CTM": {
    "code": "CTM",
    "company": "CETEM CONTAINERS N.V.",
    "city": "ANTWERP",
    "country": "Belgium"
  },
  "CTN": {
    "code": "CTN",
    "company": "CONTARGO GMBH & CO. KG",
    "city": "Mannheim",
    "country": "Germany"
  },
  "CTP": {
    "code": "CTP",
    "company": "PT PELAYARAN CARAKA TIRTA PERKASA",
    "city": "JAKARTA",
    "country": "Indonesia"
  },
  "CTQ": {
    "code": "CTQ",
    "company": "CTR INTERNATIONAL INC.",
    "city": "Louiseville",
    "country": "Canada"
  },
  "CTR": {
    "code": "CTR",
    "company": "H.ESSERS TRANSPORT COMPANY NEDERLAND BV",
    "city": "VALKENSWAARD",
    "country": "Netherlands"
  },
  "CTS": {
    "code": "CTS",
    "company": "CYCLIFE FRANCE SA",
    "city": "CODOLET",
    "country": "France"
  },
  "CTT": {
    "code": "CTT",
    "company": "CANTAS IC VE DIS TIC.SOG.SIS.SAN.A.S.",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "CTU": {
    "code": "CTU",
    "company": "CZECH TECHNICAL UNIVERSITY IN PRAGUE",
    "city": "Bustehrad",
    "country": "Czech Republic"
  },
  "CTW": {
    "code": "CTW",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "CTX": {
    "code": "CTX",
    "company": "CONTAINEX CONTAINER -HANDELSGESELLSCHAFT MBH",
    "city": "WIENER NEUDORF",
    "country": "Austria"
  },
  "CUB": {
    "code": "CUB",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "CUC": {
    "code": "CUC",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "CUN": {
    "code": "CUN",
    "company": "COMBIUNITS AS",
    "city": "EIDSVAGNESET",
    "country": "Norway"
  },
  "CVA": {
    "code": "CVA",
    "company": "TAYLOR-WHARTON AMERICA INC",
    "city": "BAYTOWN,TX",
    "country": "United States"
  },
  "CVV": {
    "code": "CVV",
    "company": "CV SUKSES MAJU BERSAMA",
    "city": "CIKARANG, BEKASI",
    "country": "Indonesia"
  },
  "CWF": {
    "code": "CWF",
    "company": "SITC COWIN SUPPLY CHAIN LIMITED",
    "city": "Causeway Bay",
    "country": "HK"
  },
  "CWL": {
    "code": "CWL",
    "company": "CONTAINERWORLD (PTY) LTD",
    "city": "DURBAN",
    "country": "South Africa"
  },
  "CWS": {
    "code": "CWS",
    "company": "GCATAINER BV",
    "city": "MOERDIJK",
    "country": "Netherlands"
  },
  "CWT": {
    "code": "CWT",
    "company": "CARGOSTORE WORLDWIDE TRADING LIMITED",
    "city": "LONDON SW19 7QD",
    "country": "United Kingdom"
  },
  "CWW": {
    "code": "CWW",
    "company": "CONTAINER WORLD-WIDE, INC",
    "city": "Rowland Heights",
    "country": "United States"
  },
  "CXC": {
    "code": "CXC",
    "company": "MATSON NAVIGATION COMPANY, INC",
    "city": "OAKLAND, CA 94610",
    "country": "United States"
  },
  "CXD": {
    "code": "CXD",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "CXH": {
    "code": "CXH",
    "company": "CEX (HKG) LIMITED",
    "city": "CAUSEWAY BAY, HONG KONG",
    "country": "HK"
  },
  "CXI": {
    "code": "CXI",
    "company": "TITAN CONTAINERS A/S",
    "city": "TAASTRUP",
    "country": "Denmark"
  },
  "CXL": {
    "code": "CXL",
    "company": "CANGZHOU SUNHEAT CHEMICALS CO LTD",
    "city": "CANGZHOU CITY, HEBEI PROVINCE",
    "country": "China"
  },
  "CXN": {
    "code": "CXN",
    "company": "CONTAINER CORPORATION OF INDIA LTD.",
    "city": "NEW DELHI",
    "country": "India"
  },
  "CXP": {
    "code": "CXP",
    "company": "ORANO CYCLE",
    "city": "PIERRELATTE CEDEX",
    "country": "France"
  },
  "CXR": {
    "code": "CXR",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "CXS": {
    "code": "CXS",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "CXT": {
    "code": "CXT",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "CXW": {
    "code": "CXW",
    "company": "CONEXWEST",
    "city": "San Francisco",
    "country": "United States"
  },
  "CYA": {
    "code": "CYA",
    "company": "CYANCO INTERNATIONAL,LLC",
    "city": "RENO, NV 89523",
    "country": "United States"
  },
  "CYC": {
    "code": "CYC",
    "company": "TRANSPLANNER SP ZOO",
    "city": "BANINO",
    "country": "Poland"
  },
  "CYD": {
    "code": "CYD",
    "company": "CRYODIRECT LIMITED",
    "city": "CENTRAL HONG KONG",
    "country": "HK"
  },
  "CYL": {
    "code": "CYL",
    "company": "CRYOLOR",
    "city": "ENNERY",
    "country": "France"
  },
  "CYZ": {
    "code": "CYZ",
    "company": "CRRC YANGTZE CO., LTD.",
    "city": "Wuhan",
    "country": "China"
  },
  "CZL": {
    "code": "CZL",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "CZP": {
    "code": "CZP",
    "company": "CONZEPT CONTAINER MODULBAU & HANDEL GMBH",
    "city": "SALZBURG",
    "country": "Austria"
  },
  "CZT": {
    "code": "CZT",
    "company": "NEW CENTRANS INTERNATIONAL MARINE SHIPPING CO., LIMITED",
    "city": "",
    "country": "HK"
  },
  "CZZ": {
    "code": "CZZ",
    "company": "CAI INTERNATIONAL",
    "city": "SAN FRANCISCO, CA 94105",
    "country": "United States"
  },
  "DAA": {
    "code": "DAA",
    "company": "DANISH MINISTRY OF DEFENCE ACQUISITION AND LOGISTICS ORGANISATION",
    "city": "BALLERUP",
    "country": "Denmark"
  },
  "DAC": {
    "code": "DAC",
    "company": "DANCONTAINER A/S",
    "city": "Nordhavn",
    "country": "Denmark"
  },
  "DAD": {
    "code": "DAD",
    "company": "TANKTRAILER NEDERLAND",
    "city": "NUMANSDORP",
    "country": "Netherlands"
  },
  "DAI": {
    "code": "DAI",
    "company": "DAIKIN INDUSTRIES, LTD",
    "city": "OSAKA",
    "country": "Japan"
  },
  "DAK": {
    "code": "DAK",
    "company": "DARKA FOR TRADING & SERVICES CO. LTD",
    "city": "PORT SUDAN",
    "country": "Sudan"
  },
  "DAM": {
    "code": "DAM",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "DAW": {
    "code": "DAW",
    "company": "GLENCORE AGRICULTURE UK LTD",
    "city": "THAME",
    "country": "United Kingdom"
  },
  "DAX": {
    "code": "DAX",
    "company": "NAVIERA DIRECT AFRICA LINE S.A",
    "city": "MADRID",
    "country": "Spain"
  },
  "DAY": {
    "code": "DAY",
    "company": "DAL DEUTSCHE AFRIKA-LINIEN GMBH & CO",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "DBC": {
    "code": "DBC",
    "company": "DRY BOX INC",
    "city": "CHEHALIS, WA-98532",
    "country": "United States"
  },
  "DBO": {
    "code": "DBO",
    "company": "DE BOER CONTAINER TRADING B.V.",
    "city": "SUSTEREN",
    "country": "Netherlands"
  },
  "DBR": {
    "code": "DBR",
    "company": "DRILL BASE OIL RESOURCES SDN BHD",
    "city": "MONT\\'KIARA",
    "country": "Malaysia"
  },
  "DBV": {
    "code": "DBV",
    "company": "DBV BAUMASCHINEN & BAUGERAETEVERTRIEBS GMBH",
    "city": "DEGGENDORF",
    "country": "Germany"
  },
  "DBX": {
    "code": "DBX",
    "company": "DB CARGO AG",
    "city": "MAINZ",
    "country": "Germany"
  },
  "DBZ": {
    "code": "DBZ",
    "company": "ISTALKO ALKOL TRANSIT TICARET LIMITED SIRKETI",
    "city": "KABATAS-BEYOGLU, ISTANBUL",
    "country": "Turkey"
  },
  "DCA": {
    "code": "DCA",
    "company": "CACI",
    "city": "RAMBOUILLET CEDEX",
    "country": "France"
  },
  "DCH": {
    "code": "DCH",
    "company": "DMITRIEVSKY CHEMICAL PLANT - PRODUCTION",
    "city": "Kineshma, Ivanovo region",
    "country": "Russian Federation"
  },
  "DCI": {
    "code": "DCI",
    "company": "DANA CONTAINER INC.",
    "city": "AVENEL, NJ 07001",
    "country": "United States"
  },
  "DCL": {
    "code": "DCL",
    "company": "DEUCON CHEMIELOGISTIK GMBH",
    "city": "Hanstedt",
    "country": "Germany"
  },
  "DCM": {
    "code": "DCM",
    "company": "MINISTERE DE LA DEFENSE (CSOA/SIMMT)",
    "city": "VILLACOUBLAY",
    "country": "France"
  },
  "DCN": {
    "code": "DCN",
    "company": "EDF DCN",
    "city": "ST DENIS CEDEX",
    "country": "France"
  },
  "DCO": {
    "code": "DCO",
    "company": "DOW SILICONES UK LIMITED",
    "city": "BARRY CF63 2YL",
    "country": "United Kingdom"
  },
  "DCS": {
    "code": "DCS",
    "company": "DANISH CONTAINER SUPPLY",
    "city": "NOERRESUNDBY",
    "country": "Denmark"
  },
  "DCT": {
    "code": "DCT",
    "company": "BERTSCHI AG",
    "city": "DUERRENAESCH",
    "country": "Switzerland"
  },
  "DCX": {
    "code": "DCX",
    "company": "NOURYON PPC AB",
    "city": "BOHUS",
    "country": "Sweden"
  },
  "DDA": {
    "code": "DDA",
    "company": "DOW DEUTSCHLAND ANLAGEN GMBH",
    "city": "WERK STADE",
    "country": "Germany"
  },
  "DDC": {
    "code": "DDC",
    "company": "DAL DEUTSCHE AFRIKA-LINIEN GMBH & CO",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "DDD": {
    "code": "DDD",
    "company": "WNG CONTAINER SERVICE CO LIMITED",
    "city": "Shanghai",
    "country": "China"
  },
  "DDE": {
    "code": "DDE",
    "company": "DE DECKER - VAN RIET BVBA",
    "city": "LONDERZEEL-MALDEREN",
    "country": "Belgium"
  },
  "DEA": {
    "code": "DEA",
    "company": "JSC ECOMET-S",
    "city": "ST PETERSBURG",
    "country": "Russian Federation"
  },
  "DEB": {
    "code": "DEB",
    "company": "NASACO INTERNATIONAL LIMITED",
    "city": "COSSONAY-VILLE",
    "country": "Switzerland"
  },
  "DEC": {
    "code": "DEC",
    "company": "DECCAN FINE CHEMICALS INDIA PRIVATE LIMITED",
    "city": "HYDERABAD",
    "country": "India"
  },
  "DEK": {
    "code": "DEK",
    "company": "FORVARD TRANS SERVICE LTD",
    "city": "SAINT PETERSBURG",
    "country": "Russian Federation"
  },
  "DEL": {
    "code": "DEL",
    "company": "AMARAL INDUSTRIES COMMON LAW",
    "city": "OAKLAND",
    "country": "United States"
  },
  "DEN": {
    "code": "DEN",
    "company": "HARDING CONTAINERS INTERNATIONAL INC.",
    "city": "LONG BEACH, CA 90810",
    "country": "United States"
  },
  "DFD": {
    "code": "DFD",
    "company": "DSV ROAD HOLDING A/S",
    "city": "HEDEHUSENE",
    "country": "Denmark"
  },
  "DFI": {
    "code": "DFI",
    "company": "VENTURA TRADING LTD",
    "city": "HAMILTON, HM",
    "country": "Bermuda"
  },
  "DFL": {
    "code": "DFL",
    "company": "LLC DFI-TT",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "DFO": {
    "code": "DFO",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "DFS": {
    "code": "DFS",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "DFV": {
    "code": "DFV",
    "company": "DESAI FRUITS VENTURE PRIVATE LTD",
    "city": "Navsari",
    "country": "India"
  },
  "DFW": {
    "code": "DFW",
    "company": "STEEL CONTAINERS NET",
    "city": "BURLESON, TX-76028",
    "country": "United States"
  },
  "DGF": {
    "code": "DGF",
    "company": "DANMAR LINES LTD",
    "city": "BASEL",
    "country": "Switzerland"
  },
  "DGG": {
    "code": "DGG",
    "company": "DAWSONGROUP GLOBAL",
    "city": "Milton Keynes",
    "country": "United Kingdom"
  },
  "DGI": {
    "code": "DGI",
    "company": "ISOCHEM LOGISTICS LLC",
    "city": "Houston, TX 77078",
    "country": "United States"
  },
  "DGR": {
    "code": "DGR",
    "company": "SHANGHAI DONG GUILIAN CONTAINER RENTAL CO.,LTD",
    "city": "Shanghai",
    "country": "China"
  },
  "DHB": {
    "code": "DHB",
    "company": "DEN HARTOGH DRY BULK LOGISTICS",
    "city": "Hull",
    "country": "United Kingdom"
  },
  "DHD": {
    "code": "DHD",
    "company": "DEN HARTOGH GLOBAL BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "DHG": {
    "code": "DHG",
    "company": "DEN HARTOGH GAS LOGISTICS",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "DHI": {
    "code": "DHI",
    "company": "DEN HARTOGH LIQUID LOGISTICS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "DHR": {
    "code": "DHR",
    "company": "DEN HARTOGH LIQUID LOGISTICS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "DHZ": {
    "code": "DHZ",
    "company": "DHL GLOBAL FORWARDING (ITALY) S.P.A",
    "city": "POZZUOLO MARTESANA (MI)",
    "country": "Italy"
  },
  "DIA": {
    "code": "DIA",
    "company": "DREDGING INTERNATIONAL ASIA PACIFIC PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "DID": {
    "code": "DID",
    "company": "BERTSCHI AG",
    "city": "DUERRENAESCH",
    "country": "Switzerland"
  },
  "DIG": {
    "code": "DIG",
    "company": "DAESUNG INDUSTRIAL GASES CO.LTD",
    "city": "GYEONGGI-DO",
    "country": "Korea, Republic of"
  },
  "DIL": {
    "code": "DIL",
    "company": "DOUGLAS INDUSTRIAL & LOGISTICS (SI) LTD",
    "city": "Honiara SB",
    "country": "Solomon Islands"
  },
  "DIN": {
    "code": "DIN",
    "company": "DINLANKA LOGISTICS (PVT) LTD",
    "city": "Colombo",
    "country": "Sri Lanka"
  },
  "DIS": {
    "code": "DIS",
    "company": "STAR",
    "city": "MAMOUDZOU",
    "country": "Mayotte"
  },
  "DJC": {
    "code": "DJC",
    "company": "DJ CONTAINERS BVBA SARL",
    "city": "ANTWERPEN",
    "country": "Belgium"
  },
  "DJI": {
    "code": "DJI",
    "company": "DJIBOUTI SHIPPING COMPANY FZE",
    "city": "",
    "country": "Djibouti"
  },
  "DJL": {
    "code": "DJL",
    "company": "DONGJIN SHIPPING CO,LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "DKL": {
    "code": "DKL",
    "company": "ROYAL NETHERLANDS AIR FORCE",
    "city": "BREDA",
    "country": "Netherlands"
  },
  "DKO": {
    "code": "DKO",
    "company": "DAELIM CORPORATION",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "DLA": {
    "code": "DLA",
    "company": "DENTOON BV",
    "city": "HELLEVOETSLUIS",
    "country": "Netherlands"
  },
  "DLF": {
    "code": "DLF",
    "company": "D'ALFONSO AUTOTRASPORTI SRL",
    "city": "CROTONE KR",
    "country": "Italy"
  },
  "DLH": {
    "code": "DLH",
    "company": "HONGKONG DALIANG MARINE LIMITED",
    "city": "HONGKONG",
    "country": "China"
  },
  "DLK": {
    "code": "DLK",
    "company": "STOLT TANK CONTAINERS LEASING LTD",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "DLM": {
    "code": "DLM",
    "company": "DLM DIENSTLEISTUNGEN & MANAGEMENT GMBH",
    "city": "BREMEN",
    "country": "Germany"
  },
  "DLS": {
    "code": "DLS",
    "company": "DELTA SERVICE LOCATION",
    "city": "CORBAS",
    "country": "France"
  },
  "DLT": {
    "code": "DLT",
    "company": "DALREFTRANS LTD",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "DLV": {
    "code": "DLV",
    "company": "DELVER AGENTS LLC",
    "city": "Seattle, WA 98119",
    "country": "United States"
  },
  "DMB": {
    "code": "DMB",
    "company": "DAMEN SHIPYARDS GORINCHEM",
    "city": "GORINCHEM",
    "country": "Netherlands"
  },
  "DME": {
    "code": "DME",
    "company": "DME AEROSOL LLC",
    "city": "Pervomaysky workers settlement",
    "country": "Russian Federation"
  },
  "DMI": {
    "code": "DMI",
    "company": "DESKTOP METAL",
    "city": "Burlington",
    "country": "United States"
  },
  "DML": {
    "code": "DML",
    "company": "DOLPHIN MOVERS LIMITED",
    "city": "Enfield",
    "country": "United Kingdom"
  },
  "DMZ": {
    "code": "DMZ",
    "company": "LLC TEHINVESTPOSTACH",
    "city": "KYIV",
    "country": "Ukraine"
  },
  "DNA": {
    "code": "DNA",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "DNC": {
    "code": "DNC",
    "company": "DANCONTAINER A/S",
    "city": "Nordhavn",
    "country": "Denmark"
  },
  "DNE": {
    "code": "DNE",
    "company": "DOERSAM + NICKEL TRANSPORT GMBH",
    "city": "MAINZ",
    "country": "Germany"
  },
  "DNR": {
    "code": "DNR",
    "company": "SPECTRANSGARANT LLC",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "DNS": {
    "code": "DNS",
    "company": "JSC NAC KAZATOMPROM",
    "city": "ASTANA",
    "country": "Kazakhstan"
  },
  "DNV": {
    "code": "DNV",
    "company": "CARGOSTORE WORLDWIDE TRADING LIMITED",
    "city": "LONDON SW19 7QD",
    "country": "United Kingdom"
  },
  "DOC": {
    "code": "DOC",
    "company": "USDOC NOAA / PMEL",
    "city": "SEATTLE, WA 98115",
    "country": "United States"
  },
  "DOD": {
    "code": "DOD",
    "company": "MILITARY SURFACE DEPLOYMENT AND DISTRIBUTION COMMAND",
    "city": "SCOTT AFB,IL 62225-5006",
    "country": "United States"
  },
  "DOG": {
    "code": "DOG",
    "company": "BIG DOG CONTAINERS INC",
    "city": "New Westminster",
    "country": "Canada"
  },
  "DOL": {
    "code": "DOL",
    "company": "DOLPHIN LINE SHIPPING LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "DON": {
    "code": "DON",
    "company": "COMMISSIONARIA S.R.L.",
    "city": "RAVENNA RA",
    "country": "Italy"
  },
  "DOV": {
    "code": "DOV",
    "company": "DONMEZ VARIL AS",
    "city": "IZMIR",
    "country": "Turkey"
  },
  "DOW": {
    "code": "DOW",
    "company": "PBBPOLISUR S.R.L",
    "city": "BUENOS AIRES",
    "country": "Argentina"
  },
  "DPC": {
    "code": "DPC",
    "company": "DUPONT",
    "city": "WILMINGTON, NC 19805",
    "country": "United States"
  },
  "DPI": {
    "code": "DPI",
    "company": "EDF",
    "city": "MONTEVRAIN",
    "country": "France"
  },
  "DPK": {
    "code": "DPK",
    "company": "UNITED NATIONS LOGISTICS BASE",
    "city": "BRINDISI",
    "country": "Italy"
  },
  "DPS": {
    "code": "DPS",
    "company": "DART PORTABLE STORAGE, INC",
    "city": "EAGAN, MN 55121",
    "country": "United States"
  },
  "DRA": {
    "code": "DRA",
    "company": "DRACONTAINERS CORP",
    "city": "MARBELLA,",
    "country": "Panama"
  },
  "DRS": {
    "code": "DRS",
    "company": "FILTERCARE BV",
    "city": "OSS",
    "country": "Netherlands"
  },
  "DRY": {
    "code": "DRY",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "DSL": {
    "code": "DSL",
    "company": "DANUBE SHIPPING SERVICE LIMITED",
    "city": "Hong Kong, Wan Chai",
    "country": "HK"
  },
  "DSN": {
    "code": "DSN",
    "company": "CEA",
    "city": "St Paul Lez Durance",
    "country": "France"
  },
  "DSR": {
    "code": "DSR",
    "company": "DISA RED DE SERVICIOS PETROLIFEROS, S.A.",
    "city": "Santa Cruz de Tenerife",
    "country": "Spain"
  },
  "DSS": {
    "code": "DSS",
    "company": "PT INDONESIA GUANG CHING NICKEL AND STAINLESS STEEL IND",
    "city": "SOUTH JAKARTA",
    "country": "Indonesia"
  },
  "DST": {
    "code": "DST",
    "company": "DMS S.R.O.",
    "city": "DUKOVANY",
    "country": "Czech Republic"
  },
  "DSZ": {
    "code": "DSZ",
    "company": "DRAGON GROUP",
    "city": "HONG KONG",
    "country": "HK"
  },
  "DTC": {
    "code": "DTC",
    "company": "DECCAN TRANSCON LEASING PVT LTD",
    "city": "Hyderabad",
    "country": "India"
  },
  "DTG": {
    "code": "DTG",
    "company": "NIPPON GASES OFFSHORE LIMITED",
    "city": "ABERDEEN",
    "country": "United Kingdom"
  },
  "DTL": {
    "code": "DTL",
    "company": "DANTECO INDUSTRIES BV",
    "city": "BERGSCHENHOEK",
    "country": "Netherlands"
  },
  "DTO": {
    "code": "DTO",
    "company": "TERMCOTANK SA",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "DTP": {
    "code": "DTP",
    "company": "VENTURA TRADING LTD",
    "city": "HAMILTON, HM",
    "country": "Bermuda"
  },
  "DTX": {
    "code": "DTX",
    "company": "HARDING CONTAINERS INTERNATIONAL INC.",
    "city": "LONG BEACH, CA 90810",
    "country": "United States"
  },
  "DUC": {
    "code": "DUC",
    "company": "DUTCH ANTILLEAN CONTAINER LEASING NV",
    "city": "LIMASSOL",
    "country": "Cyprus"
  },
  "DUT": {
    "code": "DUT",
    "company": "DUTCHTAINER",
    "city": "Barendrecht",
    "country": "Netherlands"
  },
  "DUV": {
    "code": "DUV",
    "company": "DUVEL MOORTGAT NV",
    "city": "Puurs",
    "country": "Belgium"
  },
  "DVK": {
    "code": "DVK",
    "company": "SELENA LTD",
    "city": "St. Petersburg",
    "country": "Russian Federation"
  },
  "DVR": {
    "code": "DVR",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "DWA": {
    "code": "DWA",
    "company": "EVONIK DEGUSSA ANTWERPEN N.V",
    "city": "ANTWERPEN  4",
    "country": "Belgium"
  },
  "DWB": {
    "code": "DWB",
    "company": "WARSTEINER BRAUEREI",
    "city": "WARSTEIN",
    "country": "Germany"
  },
  "DWF": {
    "code": "DWF",
    "company": "DONGWOO FINE CHEM CO, LTD",
    "city": "JEONBUK",
    "country": "Korea, Republic of"
  },
  "DXI": {
    "code": "DXI",
    "company": "PT. DEXIN STEEL INDONESIA",
    "city": "JAKARTA",
    "country": "Indonesia"
  },
  "DXT": {
    "code": "DXT",
    "company": "DIOXITEK SA",
    "city": "BUENOS AIRES",
    "country": "Argentina"
  },
  "DYL": {
    "code": "DYL",
    "company": "DONGYOUNG SHIPPING CO.,LTD",
    "city": "Jung-gu, Seoul",
    "country": "Korea, Republic of"
  },
  "DZL": {
    "code": "DZL",
    "company": "TORINO GRAZYNA NYCZ",
    "city": "BOLESLAWIEC",
    "country": "Poland"
  },
  "EAC": {
    "code": "EAC",
    "company": "EAG INTERNATIONAL CONTAINER LIMITED",
    "city": "HONG KONG",
    "country": "HK"
  },
  "EAD": {
    "code": "EAD",
    "company": "EADS SPACE TRANSPORTATION SAS",
    "city": "LES MUREAUX",
    "country": "France"
  },
  "EAF": {
    "code": "EAF",
    "company": "AGI",
    "city": "BH15 3SS, POOLE",
    "country": "United Kingdom"
  },
  "EAG": {
    "code": "EAG",
    "company": "ERNST AUTOTRANSPORT AG",
    "city": "ZURICH",
    "country": "Switzerland"
  },
  "EAI": {
    "code": "EAI",
    "company": "CSRU LOCACOES DE EQUIPAMENTOS, VEICULOS E TRANSPORTE S.A",
    "city": "MONTEVIDEO",
    "country": "Uruguay"
  },
  "EAL": {
    "code": "EAL",
    "company": "EUROASIA TOTAL LOGISTICS (M) SDN. BHD.",
    "city": "Penang",
    "country": "Malaysia"
  },
  "EAM": {
    "code": "EAM",
    "company": "ASSOCIATED ASPHALT TAMPA, LLC",
    "city": "TAMPA, FL 33619",
    "country": "United States"
  },
  "EAR": {
    "code": "EAR",
    "company": "EURO-ASIAN RUISHENG INDUSTRIAL CO,LIMITED",
    "city": "HONG KONG",
    "country": "HK"
  },
  "EAS": {
    "code": "EAS",
    "company": "WATERFRONT CONTAINER LEASING CO INC.",
    "city": "SAN FRANCISCO, CA 94109",
    "country": "United States"
  },
  "EAX": {
    "code": "EAX",
    "company": "EAS INTERNATIONAL SHIPPING CO LTD.",
    "city": "TIANJIN",
    "country": "China"
  },
  "EBC": {
    "code": "EBC",
    "company": "TWS TANKCONTAINER-LEASING GMBH & CO KG",
    "city": "Hamburg",
    "country": "Germany"
  },
  "ECA": {
    "code": "ECA",
    "company": "ENERCON GMBH",
    "city": "AURICH",
    "country": "Germany"
  },
  "ECB": {
    "code": "ECB",
    "company": "EUROPEAN CONTAINERS  NV",
    "city": "ZEEBRUGGE",
    "country": "Belgium"
  },
  "ECC": {
    "code": "ECC",
    "company": "ALOREM",
    "city": "Beynost",
    "country": "France"
  },
  "ECE": {
    "code": "ECE",
    "company": "HQEC",
    "city": "STRASBOURG CEDEX 1",
    "country": "France"
  },
  "ECI": {
    "code": "ECI",
    "company": "ECEM BV",
    "city": "AMSTERDAM",
    "country": "Netherlands"
  },
  "ECL": {
    "code": "ECL",
    "company": "TRANSPORT L'ECLIPSE",
    "city": "Corbas",
    "country": "France"
  },
  "ECM": {
    "code": "ECM",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "ECN": {
    "code": "ECN",
    "company": "ECONSHIP MARINE PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "ECO": {
    "code": "ECO",
    "company": "ECO.GE.M.M.A",
    "city": "ASSEMINI (CA)",
    "country": "Italy"
  },
  "ECP": {
    "code": "ECP",
    "company": "ECLIPS PTY LTD",
    "city": "Fyshwick",
    "country": "Australia"
  },
  "ECR": {
    "code": "ECR",
    "company": "ENERGYST GROUP SERVICES B.V.",
    "city": "BREDA",
    "country": "Netherlands"
  },
  "ECS": {
    "code": "ECS",
    "company": "BALTICON S.A.",
    "city": "GDYNIA",
    "country": "Poland"
  },
  "ECU": {
    "code": "ECU",
    "company": "ECOPRO-M LLC",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "EDC": {
    "code": "EDC",
    "company": "DAHER NUCLEAR TECHNOLOGIES SAS",
    "city": "Marignane Cedex",
    "country": "France"
  },
  "EDD": {
    "code": "EDD",
    "company": "TEHNOTRANS LTD",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "EDP": {
    "code": "EDP",
    "company": "EUROPEAN DATA PROJECT S.R.O",
    "city": "ROUSINOV",
    "country": "Czech Republic"
  },
  "EDR": {
    "code": "EDR",
    "company": "EDER ZOLTAN E.V.",
    "city": "ZALAEGERSZEG",
    "country": "Hungary"
  },
  "EDZ": {
    "code": "EDZ",
    "company": "E-COMMODITIES HOLDINGS LIMITED",
    "city": "Beiijng",
    "country": "China"
  },
  "EET": {
    "code": "EET",
    "company": "ENERGO-TRANS",
    "city": "Moscou",
    "country": "Russian Federation"
  },
  "EFC": {
    "code": "EFC",
    "company": "FINSTERWALDER CONTAINER GMBH",
    "city": "KAUFBEUREN",
    "country": "Germany"
  },
  "EFG": {
    "code": "EFG",
    "company": "ELECTRONIC FUOROCARBONS, LLC",
    "city": "HATFIELD, PA-19440",
    "country": "United States"
  },
  "EFK": {
    "code": "EFK",
    "company": "ELIT FASHION CLUB, PE.,",
    "city": "Kremenchuk, Poltava region",
    "country": "Ukraine"
  },
  "EGA": {
    "code": "EGA",
    "company": "ELGAS",
    "city": "Matraville",
    "country": "Australia"
  },
  "EGF": {
    "code": "EGF",
    "company": "GUIBERT FRERES SARL",
    "city": "SAINT-PIERRE",
    "country": "France"
  },
  "EGH": {
    "code": "EGH",
    "company": "EVERGREEN MARINE (HONG KONG) LTD",
    "city": "Wan Chai",
    "country": "HK"
  },
  "EGS": {
    "code": "EGS",
    "company": "EVERGREEN MARINE (SG) PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "EHP": {
    "code": "EHP",
    "company": "EURO-HEL  MAREK PIGŁAS",
    "city": "Odolanów",
    "country": "Poland"
  },
  "EHS": {
    "code": "EHS",
    "company": "EHS TANK.CO.LIMITED",
    "city": "CENTRAL",
    "country": "HK"
  },
  "EIA": {
    "code": "EIA",
    "company": "ISOTANK MANAGEMENT PTE.LTD.",
    "city": "Singapore",
    "country": "Singapore"
  },
  "EIM": {
    "code": "EIM",
    "company": "EIMSKIP ISLAND EHF",
    "city": "REYKJAVIK",
    "country": "Iceland"
  },
  "EIR": {
    "code": "EIR",
    "company": "ALLIANCEUROPE",
    "city": "BOLBEC",
    "country": "France"
  },
  "EIS": {
    "code": "EIS",
    "company": "EVERGREEN  INTERNATIONAL S.A.",
    "city": "PANAMA CITY",
    "country": "Panama"
  },
  "EIT": {
    "code": "EIT",
    "company": "GAINING ENTERPRISE S.A.",
    "city": "PANAMA CITY",
    "country": "Panama"
  },
  "EKB": {
    "code": "EKB",
    "company": "CONTMASTER",
    "city": "Ekaterinburg",
    "country": "Russian Federation"
  },
  "EKC": {
    "code": "EKC",
    "company": "VAN DEN BOSCH TRANSPORTE GMBH",
    "city": "GUNSKIRCHEN",
    "country": "Austria"
  },
  "EKL": {
    "code": "EKL",
    "company": "KAWASAKI KISEN KAISHA LTD - K LINE",
    "city": "TOKYO",
    "country": "Japan"
  },
  "EKO": {
    "code": "EKO",
    "company": "FORTUM WASTE SOLUTIONS OY",
    "city": "Fortum",
    "country": "Finland"
  },
  "EKT": {
    "code": "EKT",
    "company": "E KAY TECHNOLOGY CO LTD",
    "city": "TAOYUAN CITY",
    "country": "Taiwan, China"
  },
  "EKY": {
    "code": "EKY",
    "company": "E-KWAN ENTERPRISE CO., LTD",
    "city": "Kaohsiung",
    "country": "Taiwan, China"
  },
  "ELA": {
    "code": "ELA",
    "company": "ELA CONTAINER GMBH",
    "city": "HAREN (EMS)",
    "country": "Germany"
  },
  "ELB": {
    "code": "ELB",
    "company": "ELBURG GLOBAL BV",
    "city": "Huizen",
    "country": "Netherlands"
  },
  "ELC": {
    "code": "ELC",
    "company": "ELA CONTAINER OFFSHORE GMBH",
    "city": "HAREN (Ems)",
    "country": "Germany"
  },
  "ELE": {
    "code": "ELE",
    "company": "THALES TRANSPORTATION SYSTEMS GMBH",
    "city": "DITZINGEN",
    "country": "Germany"
  },
  "ELF": {
    "code": "ELF",
    "company": "ELFCON CONTAINER SERVICE AB",
    "city": "GOTHENBURG",
    "country": "Sweden"
  },
  "ELG": {
    "code": "ELG",
    "company": "THE EGYPTIAN OPERATING COMPANY FOR NATURAL GAS LIQUEFACTION PROJECTS PRIVATE",
    "city": "El Behiera",
    "country": "Egypt"
  },
  "ELK": {
    "code": "ELK",
    "company": "EURO NORDIC LOGISTICS B.V",
    "city": "RIDDERKERK",
    "country": "Netherlands"
  },
  "ELP": {
    "code": "ELP",
    "company": "ELEPHANT PLASTERBOARD NZ",
    "city": "HENDERSON - AUCKLAND",
    "country": "New Zealand"
  },
  "ELS": {
    "code": "ELS",
    "company": "ELECTROSEP TECHONOLOGIES INC.",
    "city": "Saint-Lambert",
    "country": "Canada"
  },
  "ELT": {
    "code": "ELT",
    "company": "EL-TRANS S.A.",
    "city": "RADZYMIN",
    "country": "Poland"
  },
  "EMA": {
    "code": "EMA",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "EMB": {
    "code": "EMB",
    "company": "NORAY CONTAINERS LOGISTICS, S.L.",
    "city": "Santa Perpètua de Mogoda",
    "country": "Spain"
  },
  "EMC": {
    "code": "EMC",
    "company": "EVERGREEN MARINE CORP (TAIWAN) LTD",
    "city": "TAOYUAN COUNTY",
    "country": "Taiwan, China"
  },
  "EMD": {
    "code": "EMD",
    "company": "EQUIPMENT MANAGEMENT SERVICES LLC",
    "city": "HOUSTON, TX 77049",
    "country": "United States"
  },
  "EMG": {
    "code": "EMG",
    "company": "EMIRATES GAS LLC",
    "city": "JEBEL ALI, DUBAI",
    "country": "United Arab Emirates"
  },
  "EMI": {
    "code": "EMI",
    "company": "ECOMARINE INTERNATIONAL SEATRADE LTD",
    "city": "LOME",
    "country": "Togo"
  },
  "EMJ": {
    "code": "EMJ",
    "company": "MJ CONLOG GMBH",
    "city": "Adendorf",
    "country": "Germany"
  },
  "EMK": {
    "code": "EMK",
    "company": "EMKAY LINES (PVT) LTD",
    "city": "PORT LOUIS",
    "country": "Mauritius"
  },
  "EML": {
    "code": "EML",
    "company": "LUGMAIR HANDELS U. TRANSPORT GMBH",
    "city": "ROITHAM",
    "country": "Austria"
  },
  "EMR": {
    "code": "EMR",
    "company": "EQUIPMENT MANAGEMENT SERVICES LLC",
    "city": "HOUSTON, TX 77049",
    "country": "United States"
  },
  "EMS": {
    "code": "EMS",
    "company": "EQUIPMENT MANAGEMENT SERVICES LLC",
    "city": "HOUSTON, TX 77049",
    "country": "United States"
  },
  "EMT": {
    "code": "EMT",
    "company": "EUROTRANSAC S.L.",
    "city": "VALENCIA",
    "country": "Spain"
  },
  "ENA": {
    "code": "ENA",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "ENB": {
    "code": "ENB",
    "company": "ECO MODAL TRANSPORT & SHIPPING BVBA",
    "city": "KALLO",
    "country": "Belgium"
  },
  "END": {
    "code": "END",
    "company": "ENDEL GDF SUEZ",
    "city": "KOUROU",
    "country": "French Guiana"
  },
  "ENG": {
    "code": "ENG",
    "company": "BASF NEDERLAND BV",
    "city": "DE MEERN",
    "country": "Netherlands"
  },
  "ENL": {
    "code": "ENL",
    "company": "EURO NORDIC LOGISTICS B.V",
    "city": "RIDDERKERK",
    "country": "Netherlands"
  },
  "ENM": {
    "code": "ENM",
    "company": "GSLINES - TRANSPORTES MARITIMOS",
    "city": "FUNCHAL",
    "country": "Portugal"
  },
  "ENR": {
    "code": "ENR",
    "company": "TRANSCOM LLP",
    "city": "ALMATY",
    "country": "Kazakhstan"
  },
  "ENS": {
    "code": "ENS",
    "company": "ENTRUST SHIPPING L.L.C",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "ENT": {
    "code": "ENT",
    "company": "ENTRANS AS",
    "city": "STABEKK",
    "country": "Norway"
  },
  "EOH": {
    "code": "EOH",
    "company": "ESPINA OBRAS HIDRAULICAS S.A",
    "city": "SANTIAGO DE COMPOSTELA",
    "country": "Spain"
  },
  "EOL": {
    "code": "EOL",
    "company": "EXIM CONTAINER SERVICES PTE. LTD.",
    "city": "Singapore",
    "country": "Singapore"
  },
  "EOS": {
    "code": "EOS",
    "company": "SHELL CATALYSTS & TECHNOLOGIES PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "EPC": {
    "code": "EPC",
    "company": "EPC UK",
    "city": "GREAT OAKLEY",
    "country": "United Kingdom"
  },
  "EPI": {
    "code": "EPI",
    "company": "EPIC-CONCEPTS LLC",
    "city": "GARDEN CITY, GA 31408",
    "country": "United States"
  },
  "EPL": {
    "code": "EPL",
    "company": "EPL INTERNATIONAL PTY LTD",
    "city": "MATRAVILLE, NSW 2036",
    "country": "Australia"
  },
  "EPP": {
    "code": "EPP",
    "company": "SCANDIC CONTAINER OY",
    "city": "ESPOO",
    "country": "Finland"
  },
  "EQP": {
    "code": "EQP",
    "company": "INTERNATIONAL EQUIPMENT MANAGEMENT USA,LLC",
    "city": "TORRANCE,CA 90501",
    "country": "United States"
  },
  "EQR": {
    "code": "EQR",
    "company": "TRITON INTERNATIONAL LTD",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "EQU": {
    "code": "EQU",
    "company": "EQUIPE CONTAINER SERVICES",
    "city": "SAN RAFAEL, CA 94903",
    "country": "United States"
  },
  "ERF": {
    "code": "ERF",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "ERI": {
    "code": "ERI",
    "company": "ERONTRANS SP. Z.O.O",
    "city": "PRUSZCZ GDANSKI",
    "country": "Poland"
  },
  "ERM": {
    "code": "ERM",
    "company": "ERMONT SAS",
    "city": "LORETTE",
    "country": "France"
  },
  "ERN": {
    "code": "ERN",
    "company": "WIHELM ERNST GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "ESC": {
    "code": "ESC",
    "company": "EXPRESSWAY CONTAINER LINE LLP",
    "city": "KOLKATA",
    "country": "India"
  },
  "ESD": {
    "code": "ESD",
    "company": "EMIRATES SHIPPING (HONG KONG) LTD",
    "city": "WONG CHUK HANG",
    "country": "HK"
  },
  "ESE": {
    "code": "ESE",
    "company": "ESSECO SRL",
    "city": "S.MARTINO TRECATE  NO",
    "country": "Italy"
  },
  "ESG": {
    "code": "ESG",
    "company": "INDUSTRIAL GASES NEW ZEALAND",
    "city": "Auckland",
    "country": "New Zealand"
  },
  "ESL": {
    "code": "ESL",
    "company": "ETHIOPIAN SHIPPING & LOGISTICS SERVICES ENTERPRISE",
    "city": "ADDIS ABABA",
    "country": "Ethiopia"
  },
  "ESP": {
    "code": "ESP",
    "company": "EMIRATES SHIPPING (HONG KONG) LTD",
    "city": "WONG CHUK HANG",
    "country": "HK"
  },
  "ESS": {
    "code": "ESS",
    "company": "KAWASAKI KISEN KAISHA LTD - K LINE",
    "city": "TOKYO",
    "country": "Japan"
  },
  "ETB": {
    "code": "ETB",
    "company": "ETI BAKIR A.S.",
    "city": "KASTAMONU",
    "country": "Turkey"
  },
  "ETC": {
    "code": "ETC",
    "company": "ACL UK",
    "city": "FELIXSTOWE IP11 7QG",
    "country": "United Kingdom"
  },
  "ETE": {
    "code": "ETE",
    "company": "ETEICOMPS S.L.",
    "city": "CARCAIXENT",
    "country": "Spain"
  },
  "ETN": {
    "code": "ETN",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "ETR": {
    "code": "ETR",
    "company": "CJSC ELF FILLING",
    "city": "Elektrougli",
    "country": "Russian Federation"
  },
  "ETT": {
    "code": "ETT",
    "company": "ELBTAINER TRADING GMBH",
    "city": "Barsbüttel",
    "country": "Germany"
  },
  "EUC": {
    "code": "EUC",
    "company": "EUCON SHIPPING AND TRANSPORT LTD",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "EUD": {
    "code": "EUD",
    "company": "EUROSPEDA, SRC",
    "city": "HUMENNE",
    "country": "Slovakia"
  },
  "EUF": {
    "code": "EUF",
    "company": "RAIL SERVICE SRL",
    "city": "ROUIGO",
    "country": "Italy"
  },
  "EUG": {
    "code": "EUG",
    "company": "CAVALIER CONTAINERS PTY LTD",
    "city": "BRISBANE QUEENSLAND",
    "country": "Australia"
  },
  "EUL": {
    "code": "EUL",
    "company": "EUROAFRICA SHIPPING LINES CYPRUS LIMITED",
    "city": "LIMASSOL",
    "country": "Cyprus"
  },
  "EUR": {
    "code": "EUR",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "EUV": {
    "code": "EUV",
    "company": "DEN HARTOGH LIQUID LOGISTICS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "EUX": {
    "code": "EUX",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "EVA": {
    "code": "EVA",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "EVK": {
    "code": "EVK",
    "company": "EVONIK CORPORATION",
    "city": "Parsippany, NJ 07054",
    "country": "United States"
  },
  "EVO": {
    "code": "EVO",
    "company": "EVOLUTION GERADORES LTDA",
    "city": "Itajai-Santa Catarina (SC)",
    "country": "Brazil"
  },
  "EVR": {
    "code": "EVR",
    "company": "OPERAIL",
    "city": "TALLINN",
    "country": "Estonia"
  },
  "EWA": {
    "code": "EWA",
    "company": "MOVE INTERMODAL NV",
    "city": "GENK",
    "country": "Netherlands"
  },
  "EWC": {
    "code": "EWC",
    "company": "EAST WEST CONTINENTAL CONTAINER LINE",
    "city": "VLADIMIR MKR VUREVETS",
    "country": "Russian Federation"
  },
  "EXF": {
    "code": "EXF",
    "company": "EXSIF WORLDWIDE",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "EXO": {
    "code": "EXO",
    "company": "EXODUS CHEMTANK PVT LIMITED",
    "city": "MUMBAI",
    "country": "India"
  },
  "EXP": {
    "code": "EXP",
    "company": "RAIL CARGO LOGISTICS - AUSTRIA GMBH",
    "city": "VIENNA",
    "country": "Austria"
  },
  "EXX": {
    "code": "EXX",
    "company": "EXSIF WORLDWIDE",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "EYT": {
    "code": "EYT",
    "company": "HARBIN EYT TECH&TRADE SHARE CO., LTD",
    "city": "HARBIN,",
    "country": "China"
  },
  "EZC": {
    "code": "EZC",
    "company": "EASY CONTAINER AG",
    "city": "Rotkreuz",
    "country": "Switzerland"
  },
  "FAA": {
    "code": "FAA",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "FAC": {
    "code": "FAC",
    "company": "TWS TANKCONTAINER-LEASING GMBH & CO KG",
    "city": "Hamburg",
    "country": "Germany"
  },
  "FAF": {
    "code": "FAF",
    "company": "CACI",
    "city": "RAMBOUILLET CEDEX",
    "country": "France"
  },
  "FAK": {
    "code": "FAK",
    "company": "FRIEDRICH A. KRUSE JUN / INTERNATIONALE SPEDITION",
    "city": "BRUNSBUTTEL",
    "country": "Germany"
  },
  "FAM": {
    "code": "FAM",
    "company": "FLEXBOX COLUMBIA SAS",
    "city": "BOGOTA",
    "country": "Colombia"
  },
  "FAN": {
    "code": "FAN",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "FAR": {
    "code": "FAR",
    "company": "INDIVIDUAL ENTREPRENEUR YEGOROV EVGENY",
    "city": "Magadan, Magadanskaya oblast",
    "country": "Russian Federation"
  },
  "FAT": {
    "code": "FAT",
    "company": "FERRARI ALDO TRASPORTI SRL",
    "city": "FIORENZUOLA D'ARDA (PC)",
    "country": "Italy"
  },
  "FAV": {
    "code": "FAV",
    "company": "JSC INFOTECH-BALTIKA-M",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "FAZ": {
    "code": "FAZ",
    "company": "STP SHIPPING AND TRADING INC",
    "city": "ASTORIA, NY 11105",
    "country": "United States"
  },
  "FBA": {
    "code": "FBA",
    "company": "AMUS",
    "city": "Chungcheongbuk-do",
    "country": "Korea, Republic of"
  },
  "FBI": {
    "code": "FBI",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "FBL": {
    "code": "FBL",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "FBR": {
    "code": "FBR",
    "company": "FRIGORIFICOS BRIGIDO LDA",
    "city": "VALVERDE",
    "country": "Portugal"
  },
  "FBT": {
    "code": "FBT",
    "company": "INTERGERMANIA TRANSPORT GMBH",
    "city": "JESTEBURG",
    "country": "Germany"
  },
  "FBX": {
    "code": "FBX",
    "company": "FLEXBOX COLUMBIA SAS",
    "city": "BOGOTA",
    "country": "Colombia"
  },
  "FCB": {
    "code": "FCB",
    "company": "FLEXBOX COLUMBIA SAS",
    "city": "BOGOTA",
    "country": "Colombia"
  },
  "FCC": {
    "code": "FCC",
    "company": "FAR EASTERN SHIPPING PLC",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "FCG": {
    "code": "FCG",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "FCI": {
    "code": "FCI",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "FCL": {
    "code": "FCL",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "FCN": {
    "code": "FCN",
    "company": "F.LLI CANIL SPA",
    "city": "BESSICA (TV)",
    "country": "Italy"
  },
  "FCR": {
    "code": "FCR",
    "company": "FINNLINES PLC",
    "city": "HELSINKI",
    "country": "Finland"
  },
  "FCX": {
    "code": "FCX",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "FDA": {
    "code": "FDA",
    "company": "FORMOSA DAIKIN ADVANCED CHEMICALS",
    "city": "TAIPEI",
    "country": "Taiwan, China"
  },
  "FDC": {
    "code": "FDC",
    "company": "SHANGHAI FNS LOGISTICS CO LTD",
    "city": "Shanghai",
    "country": "China"
  },
  "FDD": {
    "code": "FDD",
    "company": "SEALEASE N.V.",
    "city": "WILLEMSTAD, CURACAO,",
    "country": "Netherlands Antilles"
  },
  "FDK": {
    "code": "FDK",
    "company": "FLSMIDTH A/S",
    "city": "VALBY COPENHAGEN",
    "country": "Denmark"
  },
  "FDP": {
    "code": "FDP",
    "company": "DEL MONTE FRESH PRODUCE COMPANY",
    "city": "CORAL GABLES, FL 33134",
    "country": "United States"
  },
  "FDX": {
    "code": "FDX",
    "company": "FEDEX FREIGHT INC",
    "city": "HARRISON, AR 72601",
    "country": "United States"
  },
  "FEL": {
    "code": "FEL",
    "company": "FAR EAST LANDBRIDGE LTD",
    "city": "LIMASSOL",
    "country": "Cyprus"
  },
  "FEM": {
    "code": "FEM",
    "company": "FEDERAL EMERGENCY MANAGEMENT AGENCY",
    "city": "Washington",
    "country": "United States"
  },
  "FER": {
    "code": "FER",
    "company": "S.G.T. SRL",
    "city": "FIORENZUOLA D'ARDA PC",
    "country": "Italy"
  },
  "FES": {
    "code": "FES",
    "company": "FAR EASTERN SHIPPING PLC",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "FET": {
    "code": "FET",
    "company": "PEROXYCHEM SPAIN S.L",
    "city": "LA ZAIDA",
    "country": "Spain"
  },
  "FFA": {
    "code": "FFA",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "FFC": {
    "code": "FFC",
    "company": "LLC \" SPECIAL CONTAINER LINES\"",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "FFL": {
    "code": "FFL",
    "company": "SAURASHTRA FREIGHT PRIVATE LIMITED",
    "city": "Mumbai",
    "country": "India"
  },
  "FFM": {
    "code": "FFM",
    "company": "FUJIFILM ELECTRONIC MATERIALS USA, INC.",
    "city": "MESA, AZ 85212",
    "country": "United States"
  },
  "FFU": {
    "code": "FFU",
    "company": "FF LOCATION",
    "city": "LA MADELEINE",
    "country": "France"
  },
  "FGK": {
    "code": "FGK",
    "company": "FEDERAL FREIGHT COMPANY, JSC",
    "city": "Yekaterinburg",
    "country": "Russian Federation"
  },
  "FGN": {
    "code": "FGN",
    "company": "FLINT GROUP NETHERLANDS B.V.",
    "city": "´s-Gravenzande",
    "country": "Netherlands"
  },
  "FGR": {
    "code": "FGR",
    "company": "HOOVER FERGUSON UK LIMITED",
    "city": "ABERDEEN",
    "country": "United Kingdom"
  },
  "FGT": {
    "code": "FGT",
    "company": "F&G TRADING LIMITED",
    "city": "LONDON",
    "country": "United Kingdom"
  },
  "FGW": {
    "code": "FGW",
    "company": "CATERPILLAR (NI) LTD",
    "city": "LARNE",
    "country": "United Kingdom"
  },
  "FHF": {
    "code": "FHF",
    "company": "FHF FLURFOERDERGERAETE GMBH",
    "city": "BREMEN",
    "country": "Germany"
  },
  "FIA": {
    "code": "FIA",
    "company": "FRIO INDUSTRIAS ARGENTINAS SA",
    "city": "Villa Mercedes, San Luis",
    "country": "Argentina"
  },
  "FIB": {
    "code": "FIB",
    "company": "FLEXITANK INC.",
    "city": "PUERTO RICO, PR 00968",
    "country": "United States"
  },
  "FIC": {
    "code": "FIC",
    "company": "HANS FISCHER LOGISTIK AG",
    "city": "UNTERVAZ",
    "country": "Switzerland"
  },
  "FII": {
    "code": "FII",
    "company": "PT FREEPORT INDONESIA",
    "city": "PAPUA",
    "country": "Indonesia"
  },
  "FIO": {
    "code": "FIO",
    "company": "PIETRO FIORENTINI SPA",
    "city": "ARCUGNANO (VI)",
    "country": "Italy"
  },
  "FIX": {
    "code": "FIX",
    "company": "FINNCONTAINERS OY LTD",
    "city": "HELSINKI",
    "country": "Finland"
  },
  "FJA": {
    "code": "FJA",
    "company": "ALCOA FJARDAAL SF",
    "city": "Reydarfjordur",
    "country": "Iceland"
  },
  "FJK": {
    "code": "FJK",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "FJW": {
    "code": "FJW",
    "company": "FUSHUN JUNWEI LOGISTICS CO., LTD.",
    "city": "Fushun",
    "country": "China"
  },
  "FKK": {
    "code": "FKK",
    "company": "DALKOMHOLOD JSC",
    "city": "Vladivostok",
    "country": "Russian Federation"
  },
  "FLB": {
    "code": "FLB",
    "company": "FLOSIT",
    "city": "CASABLANCA",
    "country": "Morocco"
  },
  "FLC": {
    "code": "FLC",
    "company": "FINNLINES PLC",
    "city": "HELSINKI",
    "country": "Finland"
  },
  "FLD": {
    "code": "FLD",
    "company": "FELD ENTERTAINMENT,INC",
    "city": "Palmetto, FL 34221,",
    "country": "United States"
  },
  "FLF": {
    "code": "FLF",
    "company": "FINNLINES PLC",
    "city": "HELSINKI",
    "country": "Finland"
  },
  "FLO": {
    "code": "FLO",
    "company": "FLOWBOX S.A",
    "city": "BUENOS AIRES",
    "country": "Argentina"
  },
  "FLR": {
    "code": "FLR",
    "company": "FINNLINES PLC",
    "city": "HELSINKI",
    "country": "Finland"
  },
  "FLT": {
    "code": "FLT",
    "company": "EIMSKIP ISLAND EHF",
    "city": "REYKJAVIK",
    "country": "Iceland"
  },
  "FLU": {
    "code": "FLU",
    "company": "FLUORCHEMIE STULLN GMBH",
    "city": "STULLN",
    "country": "Germany"
  },
  "FMB": {
    "code": "FMB",
    "company": "P & O FERRYMASTERS LTD",
    "city": "ZEEBRUGGE",
    "country": "Belgium"
  },
  "FMD": {
    "code": "FMD",
    "company": "NEDCARGO TRANSPORT & DISTRIBUTIE BV",
    "city": "HAAFTEN",
    "country": "Netherlands"
  },
  "FME": {
    "code": "FME",
    "company": "FAMESA EXPLOSIVOS SAC",
    "city": "LIMA 33",
    "country": "Peru"
  },
  "FMI": {
    "code": "FMI",
    "company": "FONG-MING GASES INDUSTRIAL CO., LTD",
    "city": "Taipei",
    "country": "Taiwan, China"
  },
  "FML": {
    "code": "FML",
    "company": "FLEET SHIP MANAGEMENT PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "FMM": {
    "code": "FMM",
    "company": "MOERDIJK TANKCONTAINER TRADING BV",
    "city": "Moerdijk",
    "country": "Netherlands"
  },
  "FMR": {
    "code": "FMR",
    "company": "FYFFES INTERNATIONAL",
    "city": "DUNDALK, CO. LOUTH",
    "country": "Ireland"
  },
  "FMS": {
    "code": "FMS",
    "company": "SWEDISH DEFENCE FMV",
    "city": "STOCKHOLM",
    "country": "Sweden"
  },
  "FOD": {
    "code": "FOD",
    "company": "ALLIANCEUROPE",
    "city": "BOLBEC",
    "country": "France"
  },
  "FOG": {
    "code": "FOG",
    "company": "FIVE OCEANS LOGISTICS LTD",
    "city": "Hong Kong",
    "country": "HK"
  },
  "FOH": {
    "code": "FOH",
    "company": "NANTONG FOHSIN CONTAINER MANUFACTURING CO, LTD",
    "city": "JIANGSU",
    "country": "China"
  },
  "FOL": {
    "code": "FOL",
    "company": "COMPACT CONTAINER SYSTEMS, LLC",
    "city": "BOCA RATON, FL-33432",
    "country": "United States"
  },
  "FOS": {
    "code": "FOS",
    "company": "FUJIAN ORIENT SHIPPING CO,LTD",
    "city": "FUZHOU, FUJIAN PROVINCE",
    "country": "China"
  },
  "FOT": {
    "code": "FOT",
    "company": "H&S FOODTRANS BV",
    "city": "BARNEVELD",
    "country": "Netherlands"
  },
  "FPE": {
    "code": "FPE",
    "company": "AXENS",
    "city": "SALINDRES",
    "country": "France"
  },
  "FPM": {
    "code": "FPM",
    "company": "FORMOSA PLASTICS MARINE CORP",
    "city": "TAIPEI",
    "country": "Taiwan, China"
  },
  "FPO": {
    "code": "FPO",
    "company": "FREEPORT GASES LTD",
    "city": "GRAND BAHAMA",
    "country": "Bahamas"
  },
  "FPT": {
    "code": "FPT",
    "company": "RAFFLES LEASE PTE Ltd.",
    "city": "Singapore",
    "country": "Singapore"
  },
  "FRA": {
    "code": "FRA",
    "company": "CONTAINER OPERATORS SA",
    "city": "SAN ANTONIO",
    "country": "Chile"
  },
  "FRC": {
    "code": "FRC",
    "company": "GCATAINER BV",
    "city": "MOERDIJK",
    "country": "Netherlands"
  },
  "FRI": {
    "code": "FRI",
    "company": "FRICON REEFER SALES",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "FRL": {
    "code": "FRL",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "FSC": {
    "code": "FSC",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "FSR": {
    "code": "FSR",
    "company": "SPANISH MINISTRY OF DEFENCE-SPANISH AIRFORCE",
    "city": "MADRID",
    "country": "Spain"
  },
  "FSS": {
    "code": "FSS",
    "company": "FUGRO SUBSEA TECHNOLOGIES PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "FST": {
    "code": "FST",
    "company": "FSTA TRUCKING, INC.",
    "city": "DAVAO CITY",
    "country": "Philippines"
  },
  "FSZ": {
    "code": "FSZ",
    "company": "FUJIFILM ELECTRONIC MATERIALS (SUZHOU) CO.LTD",
    "city": "Suzhou",
    "country": "China"
  },
  "FTA": {
    "code": "FTA",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "FTC": {
    "code": "FTC",
    "company": "SIA FERRO TERMINAL",
    "city": "Liepaja",
    "country": "Latvia"
  },
  "FTG": {
    "code": "FTG",
    "company": "FINTRANS GL",
    "city": "St Petesburg",
    "country": "Russian Federation"
  },
  "FTI": {
    "code": "FTI",
    "company": "FIBA TECHNOLOGIES, INC",
    "city": "Littleton, MA 01460",
    "country": "United States"
  },
  "FTK": {
    "code": "FTK",
    "company": "FAIRTECK HOLDING PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "FTL": {
    "code": "FTL",
    "company": "FRONTIER LOGISTICS  LP",
    "city": "LA PORTE TX-77571",
    "country": "United States"
  },
  "FTT": {
    "code": "FTT",
    "company": "FISCHINGER GMBH TANKSPEDITION + LOG",
    "city": "BOEHRINGEN",
    "country": "Germany"
  },
  "FUD": {
    "code": "FUD",
    "company": "UFUDU SPECIALISED SPACE SOLUTIONS (PTY) LTD",
    "city": "FONTAINEBLEAU",
    "country": "South Africa"
  },
  "FUG": {
    "code": "FUG",
    "company": "FUGRO NETHERLANDS MARINE B.V.",
    "city": "Nootdorp",
    "country": "Netherlands"
  },
  "FUJ": {
    "code": "FUJ",
    "company": "FUJIFILM ELECTRONIC MATERIALS (EUROPE) N.V.",
    "city": "ZWIJNDRECHT",
    "country": "Belgium"
  },
  "FUK": {
    "code": "FUK",
    "company": "TRADECORP INTERNATIONAL PTY LTD.",
    "city": "BRISBANE",
    "country": "Australia"
  },
  "FUR": {
    "code": "FUR",
    "company": "ARIA FARIN JAAM INTL CO",
    "city": "TEHRAN",
    "country": "Iran, Islamic Republic of"
  },
  "FUW": {
    "code": "FUW",
    "company": "YUCHAI DONGTE SPECIAL PURPOSE AUTOMOBILE CO LTD",
    "city": "Suizhou City",
    "country": "China"
  },
  "FUX": {
    "code": "FUX",
    "company": "VELOTRANS SRL",
    "city": "NAPOLI",
    "country": "Italy"
  },
  "FVG": {
    "code": "FVG",
    "company": "FLEXI-VAN CORP",
    "city": "KENILWORTH, NJ 07033",
    "country": "United States"
  },
  "FVI": {
    "code": "FVI",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "FVL": {
    "code": "FVL",
    "company": "MOVETAINER",
    "city": "AZEITAO",
    "country": "Portugal"
  },
  "FWR": {
    "code": "FWR",
    "company": "FUWA MECHANICAL ENGINEERING (HK) CO., LTD.",
    "city": "FOSHAN",
    "country": "China"
  },
  "FWU": {
    "code": "FWU",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "FXL": {
    "code": "FXL",
    "company": "FLEXBOX COLUMBIA SAS",
    "city": "BOGOTA",
    "country": "Colombia"
  },
  "FZJ": {
    "code": "FZJ",
    "company": "FORSCHUNGSZENTRUM JUELICH GMBH",
    "city": "JUELICH",
    "country": "Germany"
  },
  "FZQ": {
    "code": "FZQ",
    "company": "TEIJIN ARAMID BV",
    "city": "ARNHEM",
    "country": "Netherlands"
  },
  "GAA": {
    "code": "GAA",
    "company": "GUJARAT ALKALIES AND CHEMICALS LIMITED",
    "city": "vadodara",
    "country": "India"
  },
  "GAC": {
    "code": "GAC",
    "company": "ABODI",
    "city": "MADRID",
    "country": "Spain"
  },
  "GAD": {
    "code": "GAD",
    "company": "MARINE ANTIPOLLUTION ENTERPRISE JSCO",
    "city": "VARNA",
    "country": "Bulgaria"
  },
  "GAE": {
    "code": "GAE",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "GAF": {
    "code": "GAF",
    "company": "MACS MARITIME CARRIER SHIPPING GMBH & CO",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "GAL": {
    "code": "GAL",
    "company": "GALCO NV",
    "city": "BRUXELLES",
    "country": "Belgium"
  },
  "GAM": {
    "code": "GAM",
    "company": "STOLPI GAMAR EHF",
    "city": "REYKJAVIK",
    "country": "Iceland"
  },
  "GAO": {
    "code": "GAO",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "GAR": {
    "code": "GAR",
    "company": "GARTNER KG - INTERNATIONALE TRANSPORTE",
    "city": "LAMBACH",
    "country": "Austria"
  },
  "GAS": {
    "code": "GAS",
    "company": "BULKHAUL LTD",
    "city": "MIDDLESBROUGH CLEVELAND",
    "country": "United Kingdom"
  },
  "GAT": {
    "code": "GAT",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "GAV": {
    "code": "GAV",
    "company": "PELLEGRINI TRASPORTI SRL",
    "city": "NOGAROLE ROCCA (VERONA)",
    "country": "Italy"
  },
  "GAY": {
    "code": "GAY",
    "company": "SA ELECTRICITY GAY",
    "city": "CHALONS EN CHAMPAGNE",
    "country": "France"
  },
  "GAZ": {
    "code": "GAZ",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "GBK": {
    "code": "GBK",
    "company": "TAM INTERNATIONAL LP",
    "city": "SASKATOON, SK",
    "country": "Canada"
  },
  "GBM": {
    "code": "GBM",
    "company": "BOURGEY MONTREUIL S.A",
    "city": "MERY",
    "country": "France"
  },
  "GBS": {
    "code": "GBS",
    "company": "MOMENTIVE PERFORMANCE MATERIALS GMBH",
    "city": "LEVERKUSEN",
    "country": "Germany"
  },
  "GBT": {
    "code": "GBT",
    "company": "GUIDO BERNARDINI SRL",
    "city": "TERNI TR",
    "country": "Italy"
  },
  "GCA": {
    "code": "GCA",
    "company": "GCA INTERMODAL",
    "city": "MONTELIMAR CEDEX",
    "country": "France"
  },
  "GCB": {
    "code": "GCB",
    "company": "TRANSPORT GHEYS NV",
    "city": "MOL",
    "country": "Belgium"
  },
  "GCC": {
    "code": "GCC",
    "company": "GREEN CONTAINERS AB",
    "city": "Læsø",
    "country": "Denmark"
  },
  "GCE": {
    "code": "GCE",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "GCG": {
    "code": "GCG",
    "company": "SICO - SOCIETA ITALIANA CARBURO OSSIGENO",
    "city": "MILANO  MI",
    "country": "Italy"
  },
  "GCM": {
    "code": "GCM",
    "company": "SHAANXI GREATROAD INDUSTRIAL CO LTD",
    "city": "SHAANXI",
    "country": "China"
  },
  "GCN": {
    "code": "GCN",
    "company": "GRIMALDI GROUP SPA",
    "city": "NAPOLI NA",
    "country": "Italy"
  },
  "GCR": {
    "code": "GCR",
    "company": "GARDNER CRYOGENICS",
    "city": "BETHLEHEM, PA 18017",
    "country": "United States"
  },
  "GCU": {
    "code": "GCU",
    "company": "NOBLE CONTAINER LEASING SINGAPORE PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "GCX": {
    "code": "GCX",
    "company": "GLOBAL CONTAINER INTERNATIONAL LLC",
    "city": "Hamilton",
    "country": "Bermuda"
  },
  "GDN": {
    "code": "GDN",
    "company": "GAZPROM DOBYCHA NOYABRSK",
    "city": "Noyabrsk",
    "country": "Russian Federation"
  },
  "GDS": {
    "code": "GDS",
    "company": "GOLDEN SHIP LLC",
    "city": "Ho Chi Minh",
    "country": "Viet Nam"
  },
  "GEC": {
    "code": "GEC",
    "company": "GREAT EXTEND CO,LTD",
    "city": "TAIPEI",
    "country": "Taiwan, China"
  },
  "GEK": {
    "code": "GEK",
    "company": "STELLAR FREIGHT LTD.",
    "city": "NEW-YORK, NY 10004",
    "country": "United States"
  },
  "GEL": {
    "code": "GEL",
    "company": "HOYER GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "GEN": {
    "code": "GEN",
    "company": "GCATAINER BV",
    "city": "MOERDIJK",
    "country": "Netherlands"
  },
  "GEQ": {
    "code": "GEQ",
    "company": "GREENWELL EQUIPMENT",
    "city": "ABERDEEN",
    "country": "United Kingdom"
  },
  "GER": {
    "code": "GER",
    "company": "GEORGIAN RAILWAY LLC",
    "city": "TBILISI",
    "country": "Georgia"
  },
  "GES": {
    "code": "GES",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "GET": {
    "code": "GET",
    "company": "TARROS SPA",
    "city": "LA SPEZIA  SP",
    "country": "Italy"
  },
  "GEU": {
    "code": "GEU",
    "company": "GESTION EUROPEA DE CARCAS SA",
    "city": "CAMPO DE CRIPTANA, CIUDAD REAL",
    "country": "Spain"
  },
  "GEW": {
    "code": "GEW",
    "company": "SUEZ WATER TECHNOLOGIES & SOLUTIONS (UK) LIMITED PARTNERSHIP",
    "city": "PETERBOROUGH",
    "country": "United Kingdom"
  },
  "GFC": {
    "code": "GFC",
    "company": "GULF FIRST CONTAINER LINE (UK) LTD",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "GFR": {
    "code": "GFR",
    "company": "COSIARMA SPA",
    "city": "GENOA",
    "country": "Italy"
  },
  "GFT": {
    "code": "GFT",
    "company": "GOTEBORGS FJARRTRANSPORTER AB",
    "city": "HISINGS KARRA",
    "country": "Sweden"
  },
  "GFX": {
    "code": "GFX",
    "company": "GULF FLUOR L.L.C",
    "city": "Abu Dhabi",
    "country": "United Arab Emirates"
  },
  "GFZ": {
    "code": "GFZ",
    "company": "GEOMAR",
    "city": "KIEL",
    "country": "Germany"
  },
  "GGA": {
    "code": "GGA",
    "company": "GLOBAL MARITIME ALGERIE",
    "city": "Batna",
    "country": "Algeria"
  },
  "GGC": {
    "code": "GGC",
    "company": "CARGOSHELL BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "GGE": {
    "code": "GGE",
    "company": "GESAN GRUPOS ELECTROGENOS",
    "city": "MUEL",
    "country": "Spain"
  },
  "GGG": {
    "code": "GGG",
    "company": "GLOBAL GASES",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "GGM": {
    "code": "GGM",
    "company": "JSC TAIMYR OIL COMPANY",
    "city": "KRASNOYARSK",
    "country": "Russian Federation"
  },
  "GGO": {
    "code": "GGO",
    "company": "G2 OCEAN AS",
    "city": "Bergen",
    "country": "Norway"
  },
  "GGS": {
    "code": "GGS",
    "company": "SEACOR ISLAND LINES, LLC",
    "city": "FT LAUDERDALE, FL 33316",
    "country": "United States"
  },
  "GHB": {
    "code": "GHB",
    "company": "GRIEPE CONTAINER GMBH",
    "city": "BREMEN",
    "country": "Germany"
  },
  "GHC": {
    "code": "GHC",
    "company": "GHC GERLING, HOLZ&CO.GmbH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "GHT": {
    "code": "GHT",
    "company": "GREEN REVOLUTION COOLING",
    "city": "Austin, TX 78758",
    "country": "United States"
  },
  "GIC": {
    "code": "GIC",
    "company": "EMPRESA DE GASES INDUSTRIALES",
    "city": "GUANABACOA",
    "country": "Cuba"
  },
  "GIN": {
    "code": "GIN",
    "company": "GAS INNOVATIONS",
    "city": "LA PORTE, TX 77571",
    "country": "United States"
  },
  "GIP": {
    "code": "GIP",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "GIT": {
    "code": "GIT",
    "company": "GIEZENDANNER TRANSPORT AG",
    "city": "ROTHRIST",
    "country": "Switzerland"
  },
  "GIV": {
    "code": "GIV",
    "company": "ROLLER CHEMICAL SRL",
    "city": "Fornovo s. giovanni",
    "country": "Italy"
  },
  "GJS": {
    "code": "GJS",
    "company": "GUANGDONG PLACTICS EXCHANGE CO,LTD",
    "city": "GUANGZHOU CITY",
    "country": "China"
  },
  "GKR": {
    "code": "GKR",
    "company": "GRUZOVAYA KORPORATSIYA",
    "city": "Ulyanovsk",
    "country": "Russian Federation"
  },
  "GLA": {
    "code": "GLA",
    "company": "GASPRO LATINOAMERICA S.A",
    "city": "PANAMA",
    "country": "Panama"
  },
  "GLB": {
    "code": "GLB",
    "company": "LANXESS SOLUTIONS US INC",
    "city": "EL DORADO, AR 71730",
    "country": "United States"
  },
  "GLD": {
    "code": "GLD",
    "company": "TOUAX",
    "city": "LA DEFENSE",
    "country": "France"
  },
  "GLK": {
    "code": "GLK",
    "company": "LANXESS SOLUTIONS US INC",
    "city": "EL DORADO, AR 71730",
    "country": "United States"
  },
  "GLL": {
    "code": "GLL",
    "company": "GLOBAL CONTAINER LOGISTICS,LLC",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "GLO": {
    "code": "GLO",
    "company": "GLOBAL OCEAN LINK LITHUANIA UAB",
    "city": "VILNIUS",
    "country": "Lithuania"
  },
  "GLP": {
    "code": "GLP",
    "company": "LLC PROLEASINGGROUP",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "GLQ": {
    "code": "GLQ",
    "company": "GLOBAL EQUIPMENT  LOGISTICS LTD",
    "city": "GRAYS, WEST THURROCK",
    "country": "United Kingdom"
  },
  "GLR": {
    "code": "GLR",
    "company": "SHOUZUN (SHANGHAI) TRADING CO LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "GLU": {
    "code": "GLU",
    "company": "GLOBELINK UNIMAR LOJISTIK A.S.",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "GMC": {
    "code": "GMC",
    "company": "GEM CONTAINERS LIMITED",
    "city": "London SW1Y 6LS",
    "country": "United Kingdom"
  },
  "GMD": {
    "code": "GMD",
    "company": "GEMADEPT CORPORATION",
    "city": "HO CHI MINH",
    "country": "Viet Nam"
  },
  "GMG": {
    "code": "GMG",
    "company": "GMT SHIPPING (HK) LTD",
    "city": "HONG-KONG",
    "country": "HK"
  },
  "GML": {
    "code": "GML",
    "company": "MARITIME GLOBAL LINE LTD",
    "city": "Limassol",
    "country": "Cyprus"
  },
  "GMM": {
    "code": "GMM",
    "company": "CALORIE FLUOR",
    "city": "BUC",
    "country": "France"
  },
  "GMO": {
    "code": "GMO",
    "company": "GOLD STAR LINE LTD",
    "city": "Kowloon, Hong Kong",
    "country": "HK"
  },
  "GMX": {
    "code": "GMX",
    "company": "QINGDAO GREATMICRO SUPPLY CHAIN CO.,LTD",
    "city": "QINGDAO",
    "country": "China"
  },
  "GNC": {
    "code": "GNC",
    "company": "NIPPON CONCEPT CORPORATION",
    "city": "CHIYODA-KU - TOKYO",
    "country": "Japan"
  },
  "GNF": {
    "code": "GNF",
    "company": "GLOBAL NUCLEAR FUELS",
    "city": "WILMINGTON, NC 28402",
    "country": "United States"
  },
  "GNI": {
    "code": "GNI",
    "company": "VALE NOUVELLE-CALEDONIE",
    "city": "NOUMEA CCEDEX",
    "country": "New Caledonia"
  },
  "GNL": {
    "code": "GNL",
    "company": "GANG ZONG TRADE (SHANGHAI) CO., LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "GNS": {
    "code": "GNS",
    "company": "SAMSKIP MCL BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "GOA": {
    "code": "GOA",
    "company": "PHOENIX INTERNATIONAL SRL",
    "city": "GENOVA",
    "country": "Italy"
  },
  "GOC": {
    "code": "GOC",
    "company": "3R ENVIRONMENTAL TECHNOLOGY CO.,LTD",
    "city": "Huizhou City",
    "country": "China"
  },
  "GOJ": {
    "code": "GOJ",
    "company": "GOLDFLEET MANAGEMENT LTD",
    "city": "MARLOW, BUCKS SL7 1DP",
    "country": "United Kingdom"
  },
  "GOK": {
    "code": "GOK",
    "company": "GOKBIL NAKLIYAT DEP.LOJ.SAN.VE DIS TIC.A.S.",
    "city": "KOZYATAGI /KADIKOY / ISTANBUL",
    "country": "Turkey"
  },
  "GOR": {
    "code": "GOR",
    "company": "TTES. J. GORGORI, S.L",
    "city": "TARRAGONA",
    "country": "Spain"
  },
  "GPF": {
    "code": "GPF",
    "company": "PRIVATE ENTERPRISE FIRMA GLORIA",
    "city": "ZAPOROHYE",
    "country": "Ukraine"
  },
  "GPL": {
    "code": "GPL",
    "company": "MANAGEMENT CONTROL & MAINTENANCE S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "GPV": {
    "code": "GPV",
    "company": "BALTPLUS-REGION",
    "city": "ST PETERSBURG",
    "country": "Russian Federation"
  },
  "GRB": {
    "code": "GRB",
    "company": "GOODRICH MARITIME PRIVATE LIMITED",
    "city": "GOVANDI, MUMBAI",
    "country": "India"
  },
  "GRC": {
    "code": "GRC",
    "company": "BORCHARD LINES LTD",
    "city": "LONDON EC1Y 4XY",
    "country": "United Kingdom"
  },
  "GRD": {
    "code": "GRD",
    "company": "TOUAX",
    "city": "LA DEFENSE",
    "country": "France"
  },
  "GRE": {
    "code": "GRE",
    "company": "M.A. GRENDI DAL 1828 S.P.A.",
    "city": "OPERA (MI)",
    "country": "Italy"
  },
  "GRF": {
    "code": "GRF",
    "company": "GATEWAY RAIL FREIGHT LIMITED",
    "city": "NEW DEHLI",
    "country": "India"
  },
  "GRI": {
    "code": "GRI",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "GRK": {
    "code": "GRK",
    "company": "JSC GRUDAKONTA",
    "city": "Klaipeda",
    "country": "Lithuania"
  },
  "GRL": {
    "code": "GRL",
    "company": "GRUBER GMBH  & CO.KG",
    "city": "LUDWIGSHAFEN/  RHEIN",
    "country": "Germany"
  },
  "GRM": {
    "code": "GRM",
    "company": "GOODRICH MARITIME PRIVATE LIMITED",
    "city": "GOVANDI, MUMBAI",
    "country": "India"
  },
  "GRP": {
    "code": "GRP",
    "company": "TRISTAR ENGINNERING CONSULTING LOGISTIC SA",
    "city": "CHIASSO",
    "country": "Switzerland"
  },
  "GRR": {
    "code": "GRR",
    "company": "GRANDEE PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "GRT": {
    "code": "GRT",
    "company": "HULDA MARITIME INTERNATIONAL CONTAINER",
    "city": "KOWLOON",
    "country": "HK"
  },
  "GRU": {
    "code": "GRU",
    "company": "GRUBER GMBH  & CO.KG",
    "city": "LUDWIGSHAFEN/  RHEIN",
    "country": "Germany"
  },
  "GRW": {
    "code": "GRW",
    "company": "GRW ENGINEERING (PTY) LTD",
    "city": "Worcester",
    "country": "South Africa"
  },
  "GRX": {
    "code": "GRX",
    "company": "GOODRICH MARITIME PRIVATE LIMITED",
    "city": "GOVANDI, MUMBAI",
    "country": "India"
  },
  "GSC": {
    "code": "GSC",
    "company": "NAVITRANS FRANCE",
    "city": "Marseille",
    "country": "France"
  },
  "GSH": {
    "code": "GSH",
    "company": "GSH OF ALABAMA, LLC",
    "city": "Huntsville, AL 35805",
    "country": "United States"
  },
  "GSI": {
    "code": "GSI",
    "company": "SPINELLI S.R.L",
    "city": "GENOVA",
    "country": "Italy"
  },
  "GSK": {
    "code": "GSK",
    "company": "GESTION DE SERVICIOS MARITIMOS AEREOS Y TERRESTRES, S.A",
    "city": "BARCELONA",
    "country": "Spain"
  },
  "GSL": {
    "code": "GSL",
    "company": "GOLD STAR LINE LTD",
    "city": "Kowloon, Hong Kong",
    "country": "HK"
  },
  "GSM": {
    "code": "GSM",
    "company": "AIR LIQUIDE HEALTHCARE ESPAÑA, S.L.",
    "city": "MADRID",
    "country": "Spain"
  },
  "GSO": {
    "code": "GSO",
    "company": "GLOBE OPUS SHIPPING LINE (UK) LTDTD",
    "city": "SURREY CRO 3PS",
    "country": "United Kingdom"
  },
  "GSP": {
    "code": "GSP",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "GSS": {
    "code": "GSS",
    "company": "G2 OCEAN AS",
    "city": "Bergen",
    "country": "Norway"
  },
  "GST": {
    "code": "GST",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "GTA": {
    "code": "GTA",
    "company": "GALLIKER TRANSPORT AG",
    "city": "ALTISHOFEN",
    "country": "Switzerland"
  },
  "GTC": {
    "code": "GTC",
    "company": "CTR INTERNATIONAL INC.",
    "city": "Louiseville",
    "country": "Canada"
  },
  "GTI": {
    "code": "GTI",
    "company": "MSC- MEDITERRANEAN SHIPPING COMPANY S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "GTL": {
    "code": "GTL",
    "company": "GTL AGENCIES (S) PTE.LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "GTM": {
    "code": "GTM",
    "company": "GTM MANUFACTURING LLC",
    "city": "AMARILLO, TX 79103",
    "country": "United States"
  },
  "GTP": {
    "code": "GTP",
    "company": "GTS SPA",
    "city": "GENOVA",
    "country": "Italy"
  },
  "GTR": {
    "code": "GTR",
    "company": "GENTENAAR TRANSPORT BV",
    "city": "KLUNDERT",
    "country": "Netherlands"
  },
  "GTT": {
    "code": "GTT",
    "company": "GETRAS SRL",
    "city": "Fiumicino",
    "country": "Italy"
  },
  "GUR": {
    "code": "GUR",
    "company": "CORAL TANKS PRIVATE LIMITED",
    "city": "Singapore",
    "country": "Singapore"
  },
  "GUT": {
    "code": "GUT",
    "company": "UES INTERNATIONAL (HK)  HOLDINGS LIMITED",
    "city": "SHANGHAI",
    "country": "China"
  },
  "GVC": {
    "code": "GVC",
    "company": "UES INTERNATIONAL (HK)  HOLDINGS LIMITED",
    "city": "SHANGHAI",
    "country": "China"
  },
  "GVD": {
    "code": "GVD",
    "company": "UES INTERNATIONAL (HK)  HOLDINGS LIMITED",
    "city": "SHANGHAI",
    "country": "China"
  },
  "GVS": {
    "code": "GVS",
    "company": "GULFVOSTOK LLC",
    "city": "Vladivostok",
    "country": "Russian Federation"
  },
  "GVT": {
    "code": "GVT",
    "company": "GRAND VIEW CONTAINER TRADING (HK) CO LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "GWD": {
    "code": "GWD",
    "company": "GRILLO-WERKE AG",
    "city": "DUISBURG",
    "country": "Germany"
  },
  "GWG": {
    "code": "GWG",
    "company": "GRIMM & WULFF ANLAGEN UND SYSTEMBAU GMBH",
    "city": "Seevetal",
    "country": "Germany"
  },
  "GWI": {
    "code": "GWI",
    "company": "GLORY WELL INDUSTRIES LIMITED",
    "city": "SAN PO KONG , KOWLOON",
    "country": "HK"
  },
  "GWL": {
    "code": "GWL",
    "company": "GREENWAY LOGISTICS LIMITED",
    "city": "HONG KONG",
    "country": "HK"
  },
  "GXL": {
    "code": "GXL",
    "company": "CLDN CARGO NV",
    "city": "ZEEBRUGGE",
    "country": "Belgium"
  },
  "GZE": {
    "code": "GZE",
    "company": "GAZECHIM FROID",
    "city": "BEZIERS",
    "country": "France"
  },
  "GZG": {
    "code": "GZG",
    "company": "INFRASTRUCTURE AND PRODUCTION LLC",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "GZP": {
    "code": "GZP",
    "company": "GAZPROM MARKETING & TRADING SWITZERLAND AG",
    "city": "ZUG",
    "country": "Switzerland"
  },
  "GZT": {
    "code": "GZT",
    "company": "GAZETEC INDUSTRIES S.A.L",
    "city": "JDEIDEH",
    "country": "Lebanon"
  },
  "HAA": {
    "code": "HAA",
    "company": "HAANPAA INTERNATIONAL AB",
    "city": "HELSINGBORG",
    "country": "Sweden"
  },
  "HAC": {
    "code": "HAC",
    "company": "HAIAN TRANSPORT & STEVEDORING JSC",
    "city": "Hanoi",
    "country": "Viet Nam"
  },
  "HAG": {
    "code": "HAG",
    "company": "HERRENKNECHT AG",
    "city": "Schwanau",
    "country": "Germany"
  },
  "HAH": {
    "code": "HAH",
    "company": "SHANGHAI HAI HUA SHIPPING CO LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "HAK": {
    "code": "HAK",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "HAL": {
    "code": "HAL",
    "company": "HEUNG-A SHIPPING CO LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "HAM": {
    "code": "HAM",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HAN": {
    "code": "HAN",
    "company": "HCT HANSA CONTAINER TRADING GMBH",
    "city": "Hamburg",
    "country": "Germany"
  },
  "HAP": {
    "code": "HAP",
    "company": "H.A.P. FOODS HOLLAND B.V.",
    "city": "H.I. AMBACHT",
    "country": "Netherlands"
  },
  "HAR": {
    "code": "HAR",
    "company": "HARBOUR RENTAL & SERVICES BV",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "HAS": {
    "code": "HAS",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "HAT": {
    "code": "HAT",
    "company": "HOST RAIL LTD",
    "city": "KALININGRAD",
    "country": "Russian Federation"
  },
  "HAU": {
    "code": "HAU",
    "company": "HAULDER S.A.",
    "city": "MONTEVIDEO",
    "country": "Uruguay"
  },
  "HBB": {
    "code": "HBB",
    "company": "HABIBA INTERNATIONAL GROUP LIMITED",
    "city": "Kowloon, Hong Kong",
    "country": "HK"
  },
  "HBC": {
    "code": "HBC",
    "company": "HOLDERCHEM SAL",
    "city": "BAABDA",
    "country": "Lebanon"
  },
  "HBG": {
    "code": "HBG",
    "company": "HABAS SINAI VE TIBBI GAZLAR",
    "city": "SOGANLIK -KARTAL-ISTAMBUL",
    "country": "Turkey"
  },
  "HBL": {
    "code": "HBL",
    "company": "SHANGHAI BAILIN INTERNATIONAL TRANSPORTATION CO.LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "HBN": {
    "code": "HBN",
    "company": "HALLIBURTON",
    "city": "Houston, TX 77032-3219",
    "country": "United States"
  },
  "HBS": {
    "code": "HBS",
    "company": "HANBAO CONTAIINER SHIPPING & TRADING GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HBT": {
    "code": "HBT",
    "company": "HANBAO CONTAIINER SHIPPING & TRADING GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HBW": {
    "code": "HBW",
    "company": "HAL CHEMICALS LIMITED",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "HCC": {
    "code": "HCC",
    "company": "HCCR HAMBURGER CONTAINER & CHASSIS REPARATUR GmbH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HCD": {
    "code": "HCD",
    "company": "CONTAINERDIENST HAEMMERLE GMBH",
    "city": "BLUDENZ",
    "country": "Austria"
  },
  "HCG": {
    "code": "HCG",
    "company": "HOUSTON CONTAINER CONNECTION",
    "city": "HOUSTON, TX-77013",
    "country": "United States"
  },
  "HCI": {
    "code": "HCI",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "HCO": {
    "code": "HCO",
    "company": "CONICAL CONTAINER INDUSTRIE CONSULTING",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HCS": {
    "code": "HCS",
    "company": "HACON CONTAINERS B.V",
    "city": "EUROPOORT",
    "country": "Netherlands"
  },
  "HCT": {
    "code": "HCT",
    "company": "HAMILTON CONTAINER TERMINAL",
    "city": "Hamilton",
    "country": "Canada"
  },
  "HCV": {
    "code": "HCV",
    "company": "HONG CHUN ELECTRIC & MACHINERY CO LTD",
    "city": "MIAO-LI",
    "country": "Taiwan, China"
  },
  "HCZ": {
    "code": "HCZ",
    "company": "HARDING CONTAINERS INTERNATIONAL INC.",
    "city": "LONG BEACH, CA 90810",
    "country": "United States"
  },
  "HDF": {
    "code": "HDF",
    "company": "HUADONG FERRY CO LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "HDG": {
    "code": "HDG",
    "company": "HYUNDAI GLOVIS",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "HDM": {
    "code": "HDM",
    "company": "HMM CO., LTD.",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "HDP": {
    "code": "HDP",
    "company": "HIMDILING PRO LTD",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "HDS": {
    "code": "HDS",
    "company": "KUEHNE + NAGEL DRINKS LOGISTICS LTD",
    "city": "LINFORD WOOD  MK14 6BW",
    "country": "United Kingdom"
  },
  "HDX": {
    "code": "HDX",
    "company": "HDS LINES",
    "city": "TEHRAN",
    "country": "Iran, Islamic Republic of"
  },
  "HEF": {
    "code": "HEF",
    "company": "PORTSAID STEVEDORING COMPANY",
    "city": "PORTSAID",
    "country": "Egypt"
  },
  "HEI": {
    "code": "HEI",
    "company": "HELIUM 24 LLC",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "HEM": {
    "code": "HEM",
    "company": "TEMPUS LINK LTD",
    "city": "GABROVO",
    "country": "Bulgaria"
  },
  "HER": {
    "code": "HER",
    "company": "HERC RENTALS",
    "city": "PACHECO, CA 94553",
    "country": "United States"
  },
  "HES": {
    "code": "HES",
    "company": "HUTCHINSON EQUIPMENT SERVICES",
    "city": "SAN FRANCISCO CA 94127",
    "country": "United States"
  },
  "HEX": {
    "code": "HEX",
    "company": "HAN EXPRESS",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "HEY": {
    "code": "HEY",
    "company": "MERCURIUS CONTAINER TRADING GMBH",
    "city": "Hamburg",
    "country": "Germany"
  },
  "HFA": {
    "code": "HFA",
    "company": "DERIVADOS DEL FLUOR SA",
    "city": "ONTON-CASTRO URDIALES",
    "country": "Spain"
  },
  "HFC": {
    "code": "HFC",
    "company": "SPEDITION RHEINLAND H.FREUND GMBH&CO. KG",
    "city": "FRECHEN",
    "country": "Germany"
  },
  "HFD": {
    "code": "HFD",
    "company": "DO-FLUORIDE CHEMICALS CO.,LTD",
    "city": "Jiaozuo",
    "country": "China"
  },
  "HFL": {
    "code": "HFL",
    "company": "PETRO-CANADA LUBRICANTS INC.",
    "city": "Mississauga",
    "country": "Canada"
  },
  "HFX": {
    "code": "HFX",
    "company": "MAGHREB CONTAINER INTERNATIONAL S.L",
    "city": "MADRID",
    "country": "Spain"
  },
  "HGA": {
    "code": "HGA",
    "company": "HYBAS INTERNATIONAL LLC",
    "city": "HOUSTON, TX 77084",
    "country": "United States"
  },
  "HGB": {
    "code": "HGB",
    "company": "HOYER GLOBAL TRANSPORT",
    "city": "BOTLEK-ROTTERDAM",
    "country": "Netherlands"
  },
  "HGC": {
    "code": "HGC",
    "company": "THE GAS COMPANY LLC D.B.A. HAWAII GAS",
    "city": "Honolulu, HI 96813",
    "country": "United States"
  },
  "HGD": {
    "code": "HGD",
    "company": "HANGZHOU GREENDA ELECTRONIC MATERIALS CO., LTD.",
    "city": "HANGZHOU, 310051",
    "country": "China"
  },
  "HGF": {
    "code": "HGF",
    "company": "HOYER GLOBAL TRANSPORT",
    "city": "BOTLEK-ROTTERDAM",
    "country": "Netherlands"
  },
  "HGH": {
    "code": "HGH",
    "company": "HOYER GLOBAL TRANSPORT",
    "city": "BOTLEK-ROTTERDAM",
    "country": "Netherlands"
  },
  "HGL": {
    "code": "HGL",
    "company": "JF HILLEBRAND LIMITED",
    "city": "DUBLIN 18",
    "country": "Ireland"
  },
  "HGT": {
    "code": "HGT",
    "company": "HOYER GLOBAL TRANSPORT",
    "city": "BOTLEK-ROTTERDAM",
    "country": "Netherlands"
  },
  "HHA": {
    "code": "HHA",
    "company": "HENRI HARSCH HH SA",
    "city": "CAROUGE GENEVE",
    "country": "Switzerland"
  },
  "HHM": {
    "code": "HHM",
    "company": "H.HENRIKSEN AS",
    "city": "TONSBERG",
    "country": "Norway"
  },
  "HHR": {
    "code": "HHR",
    "company": "HIMOINSA",
    "city": "SAN JAVIER, MURCIA",
    "country": "Spain"
  },
  "HHX": {
    "code": "HHX",
    "company": "HK HONORTRANS DEVELOPMENT LIMITED",
    "city": "QINGDAO",
    "country": "China"
  },
  "HIB": {
    "code": "HIB",
    "company": "HOYER GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HIC": {
    "code": "HIC",
    "company": "HAESAERTS INTERMODAL NV",
    "city": "ZWIJNDRECHT",
    "country": "Belgium"
  },
  "HIL": {
    "code": "HIL",
    "company": "HILLSTONE HOLDING B.V.",
    "city": "Galder",
    "country": "Netherlands"
  },
  "HIM": {
    "code": "HIM",
    "company": "BASELL POLIOLEFINE ITALIA  SRL",
    "city": "MILANO",
    "country": "Italy"
  },
  "HIT": {
    "code": "HIT",
    "company": "JSC HIMINVESTTRANS",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "HJL": {
    "code": "HJL",
    "company": "HONGJI (HK) CONTAINER DEVELLOPPEMENT LIMITED",
    "city": "WANCHAI",
    "country": "HK"
  },
  "HJM": {
    "code": "HJM",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "HKH": {
    "code": "HKH",
    "company": "HONG-KUANG HI TECH, CORP",
    "city": "ZHUBEI CITY",
    "country": "Taiwan, China"
  },
  "HKN": {
    "code": "HKN",
    "company": "HUBERT KLAESENER JR. FLUESSIGKEISTRANSPORTE GMBH & CO KG",
    "city": "MARL",
    "country": "Germany"
  },
  "HKT": {
    "code": "HKT",
    "company": "KLAESER INT FACHSPEDITION & FAHRZEUGBAU GMBH",
    "city": "HERTEN",
    "country": "Germany"
  },
  "HLB": {
    "code": "HLB",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HLC": {
    "code": "HLC",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HLE": {
    "code": "HLE",
    "company": "SHIHLIEN FINE CHEMICALS CO. LTD.",
    "city": "Taoyuan County",
    "country": "Taiwan, China"
  },
  "HLG": {
    "code": "HLG",
    "company": "HAEFELI AG",
    "city": "LENZBURG",
    "country": "Switzerland"
  },
  "HLM": {
    "code": "HLM",
    "company": "LANFER TRANSPORTE GMBH & CO. KG",
    "city": "Meppen",
    "country": "Germany"
  },
  "HLS": {
    "code": "HLS",
    "company": "HLS CONTAINER BREMEN",
    "city": "BREMEN",
    "country": "Germany"
  },
  "HLT": {
    "code": "HLT",
    "company": "TRANS DISTANCE LINE LTD",
    "city": "HONG KONG",
    "country": "HK"
  },
  "HLX": {
    "code": "HLX",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HMC": {
    "code": "HMC",
    "company": "EVERGREEN MARINE (UK) LIMITED",
    "city": "LONDON",
    "country": "United Kingdom"
  },
  "HMH": {
    "code": "HMH",
    "company": "HOOVER FERGUSON UK LIMITED",
    "city": "HOUSTON, TX 77077",
    "country": "United States"
  },
  "HMK": {
    "code": "HMK",
    "company": "WEC LINES BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "HMM": {
    "code": "HMM",
    "company": "HMM CO., LTD.",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "HMP": {
    "code": "HMP",
    "company": "HOYER GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HMR": {
    "code": "HMR",
    "company": "HYUNDAI MOTORSPORT GMBH",
    "city": "ALZENAU",
    "country": "Germany"
  },
  "HMS": {
    "code": "HMS",
    "company": "MED LOGISTICS (MALTA) LTD",
    "city": "SPB2807 St.Paul’s Bay",
    "country": "Malta"
  },
  "HNK": {
    "code": "HNK",
    "company": "HONEYTAK INTERMODAL LIMITED",
    "city": "Yantai",
    "country": "China"
  },
  "HNL": {
    "code": "HNL",
    "company": "VAN OORD SHIP MANAGEMENT",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "HNP": {
    "code": "HNP",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "HNS": {
    "code": "HNS",
    "company": "TRITON INTERNATIONAL LTD",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "HOB": {
    "code": "HOB",
    "company": "BURG SERVICE BV",
    "city": "ZEVENBERGEN",
    "country": "Netherlands"
  },
  "HOF": {
    "code": "HOF",
    "company": "HOFER TANKTRANSPORTE AG",
    "city": "ROTHRIST",
    "country": "Switzerland"
  },
  "HOL": {
    "code": "HOL",
    "company": "HOLCIM (SCHWEIZ) AG",
    "city": "ZURICH",
    "country": "Switzerland"
  },
  "HOR": {
    "code": "HOR",
    "company": "ALASHANKOU HORIZON PETROLEUM AND GAS INC",
    "city": "Alashankou",
    "country": "China"
  },
  "HOT": {
    "code": "HOT",
    "company": "HOYER GLOBAL TRANSPORT",
    "city": "BOTLEK-ROTTERDAM",
    "country": "Netherlands"
  },
  "HOU": {
    "code": "HOU",
    "company": "HARDING CONTAINERS INTERNATIONAL INC.",
    "city": "LONG BEACH, CA 90810",
    "country": "United States"
  },
  "HOY": {
    "code": "HOY",
    "company": "HOYER GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HPA": {
    "code": "HPA",
    "company": "HIDROJEN PEROKSIT AS",
    "city": "BANDIRMA / BALIKESIR",
    "country": "Turkey"
  },
  "HPC": {
    "code": "HPC",
    "company": "HUIZHOU SINGAMAS ENERGY EQUIPMENT CO,LTD.",
    "city": "XINXU",
    "country": "China"
  },
  "HPE": {
    "code": "HPE",
    "company": "BECKER MARINE SYSTEMS GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HPG": {
    "code": "HPG",
    "company": "TIANJIN SINO-PEAK INTERNATIONAL TRADE CO LTD",
    "city": "BINHAI NEW DISTRICT, TIANJIN 300460",
    "country": "China"
  },
  "HPJ": {
    "code": "HPJ",
    "company": "HUAI'AN PORT LOGISTICS GROUP CO.,LTD.",
    "city": "HUAI'AN",
    "country": "China"
  },
  "HRZ": {
    "code": "HRZ",
    "company": "MATSON NAVIGATION COMPANY, INC",
    "city": "OAKLAND, CA 94610",
    "country": "United States"
  },
  "HSC": {
    "code": "HSC",
    "company": "H&S CONTAINER GMBH",
    "city": "BREMEN",
    "country": "Germany"
  },
  "HSI": {
    "code": "HSI",
    "company": "HOYER GLOBAL TRANSPORT",
    "city": "BOTLEK-ROTTERDAM",
    "country": "Netherlands"
  },
  "HSL": {
    "code": "HSL",
    "company": "HUISMAN EQUIPMENT B.V.",
    "city": "SCHIEDAM",
    "country": "Netherlands"
  },
  "HST": {
    "code": "HST",
    "company": "MACS MARITIME CARRIER SHIPPING GMBH & CO",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HTC": {
    "code": "HTC",
    "company": "HALDOR TOPSOE A/S",
    "city": "KGS.LYNGBY",
    "country": "Denmark"
  },
  "HTG": {
    "code": "HTG",
    "company": "JIANGXI HUATE ELECTRONIC CHEMICALS CO., LTD",
    "city": "Jiujiang",
    "country": "China"
  },
  "HTI": {
    "code": "HTI",
    "company": "HUDSON TECHNOLOGIES INC.",
    "city": "PEARL RIVER, NY 10965",
    "country": "United States"
  },
  "HTS": {
    "code": "HTS",
    "company": "HTS LLC",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "HTU": {
    "code": "HTU",
    "company": "KIECA TRANS",
    "city": "Janowiec Wlkp",
    "country": "Poland"
  },
  "HTX": {
    "code": "HTX",
    "company": "JINGMEN HONGTU SPECIAL AIRCRAFT MANUFACTURING CO LTD",
    "city": "Jingmen",
    "country": "China"
  },
  "HUK": {
    "code": "HUK",
    "company": "HUKTRA NV",
    "city": "ZEEBRUGGE",
    "country": "Belgium"
  },
  "HUN": {
    "code": "HUN",
    "company": "UNICON INTERNATIONAL CORP.",
    "city": "Yixing City",
    "country": "China"
  },
  "HVL": {
    "code": "HVL",
    "company": "HARINERA VILAFRANQUINA S.A.",
    "city": "BARCELONA",
    "country": "Spain"
  },
  "HVR": {
    "code": "HVR",
    "company": "HYDRAUVISION RENTAL B.V.",
    "city": "SCHOONDIJKE",
    "country": "Netherlands"
  },
  "HWB": {
    "code": "HWB",
    "company": "HANS W. BARBE CHEMISCHE ERZEUGNISSE GMBH",
    "city": "WIESBADEN-SCHIERSTEIN",
    "country": "Germany"
  },
  "HWH": {
    "code": "HWH",
    "company": "HELLEMAN WAREHOUSING",
    "city": "NUMANSDORP",
    "country": "Netherlands"
  },
  "HWI": {
    "code": "HWI",
    "company": "HEALDWORKS INC",
    "city": "DEL NORTE, CO 81132",
    "country": "United States"
  },
  "HZF": {
    "code": "HZF",
    "company": "HANGZHOU FINE FLUOROTECH CO., LTD",
    "city": "Hangzhou",
    "country": "China"
  },
  "HZK": {
    "code": "HZK",
    "company": "HZ KONTEJNERY S.R.O.",
    "city": "Praha",
    "country": "Czech Republic"
  },
  "IAA": {
    "code": "IAA",
    "company": "INTERASIA LINES SINGAPORE PTE.LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "IAC": {
    "code": "IAC",
    "company": "I.A.C.C. INTERNATIONAL  ASSOCIATED",
    "city": "CAIRO",
    "country": "Egypt"
  },
  "IAG": {
    "code": "IAG",
    "company": "INDERMUEHLE LOGISTIK AG",
    "city": "REKINGEN",
    "country": "Switzerland"
  },
  "IAK": {
    "code": "IAK",
    "company": "INVEST AGRO CAPITAL LTD.",
    "city": "Kyiv",
    "country": "Ukraine"
  },
  "IAP": {
    "code": "IAP",
    "company": "HOLIDAY ON ICE PRODUCTIONS BV",
    "city": "UTRECHT",
    "country": "Netherlands"
  },
  "IBC": {
    "code": "IBC",
    "company": "DEN HARTOGH DRY BULK LOGISTICS LTD",
    "city": "Hull, HU3 4AE",
    "country": "United Kingdom"
  },
  "IBE": {
    "code": "IBE",
    "company": "IBERCISTER CISTERNAS IBERICAS",
    "city": "LOURES",
    "country": "Portugal"
  },
  "IBF": {
    "code": "IBF",
    "company": "DEN HARTOGH DRY BULK LOGISTICS LTD",
    "city": "Hull, HU3 4AE",
    "country": "United Kingdom"
  },
  "IBK": {
    "code": "IBK",
    "company": "INTERBULK,INC",
    "city": "JACKSON , MS 39236",
    "country": "United States"
  },
  "IBL": {
    "code": "IBL",
    "company": "INBULK TECHNOLOGIES LTD",
    "city": "STOCKTON ON TEES TS17 6PT",
    "country": "United Kingdom"
  },
  "IBM": {
    "code": "IBM",
    "company": "JSC INFOTECH-BALTIKA-M",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "IBT": {
    "code": "IBT",
    "company": "DEN HARTOGH GLOBAL BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "ICA": {
    "code": "ICA",
    "company": "ICON INTERNATIONAL CONTAINER SERVICE (AMERICA) LLC",
    "city": "FORT LAUDERDALE, FL-33309",
    "country": "United States"
  },
  "ICB": {
    "code": "ICB",
    "company": "INTERNATIONAL CONTAINER POOL PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "ICD": {
    "code": "ICD",
    "company": "ICON INTERNATIONAL CONTAINER SCE GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "ICH": {
    "code": "ICH",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "ICK": {
    "code": "ICK",
    "company": "INFINITY LOGISTICS & TRANSPORT SDN BHD",
    "city": "Klang",
    "country": "Malaysia"
  },
  "ICL": {
    "code": "ICL",
    "company": "EXSIF WORLDWIDE",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "ICN": {
    "code": "ICN",
    "company": "ICT - INTERNATIONAL CONTAINER TRANSPORT GMBH",
    "city": "NEUSS",
    "country": "Germany"
  },
  "ICO": {
    "code": "ICO",
    "company": "ICON INTERNATIONAL CONTAINER SCE GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "ICS": {
    "code": "ICS",
    "company": "TAL INTERNATIONAL",
    "city": "PURCHASE, NY 10577-2135",
    "country": "United States"
  },
  "ICT": {
    "code": "ICT",
    "company": "INTERCONTINENTAL TANKS LIMITED",
    "city": "",
    "country": "Gibraltar"
  },
  "ICU": {
    "code": "ICU",
    "company": "INDEPENDENT CONTAINER LINE LTD",
    "city": "GLEN ALLEN, VA 23060",
    "country": "United States"
  },
  "IDC": {
    "code": "IDC",
    "company": "PT INDOCEMENT TUNGGAL PRAKARSA TBK",
    "city": "BOGOR",
    "country": "Indonesia"
  },
  "IDG": {
    "code": "IDG",
    "company": "DINGES LOGISTICS",
    "city": "Grünstadt",
    "country": "Germany"
  },
  "IDS": {
    "code": "IDS",
    "company": "SWIRE OILFIELD SERVICES LTD",
    "city": "ABERDEEN AB123LF",
    "country": "United Kingdom"
  },
  "IDT": {
    "code": "IDT",
    "company": "ODYSSEY FOODTRANS",
    "city": "Irvine, CA 92614",
    "country": "United States"
  },
  "IEA": {
    "code": "IEA",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "IEC": {
    "code": "IEC",
    "company": "MARIANAS GAS CORPORATION DBA ISLAND EQUIPMENT CO",
    "city": "TAMUNING",
    "country": "Guam"
  },
  "IFF": {
    "code": "IFF",
    "company": "DEN HARTOGH DRY BULK LOGISTICS LTD",
    "city": "Hull, HU3 4AE",
    "country": "United Kingdom"
  },
  "IFL": {
    "code": "IFL",
    "company": "INTERFLOW (T.C.S.) LTD",
    "city": "LONDON, EC3A 7LP",
    "country": "United Kingdom"
  },
  "IFT": {
    "code": "IFT",
    "company": "LEIBNIZ INSTITUT FOR TROPOSPHERIC RESEARCH",
    "city": "LEIPZIG",
    "country": "Germany"
  },
  "IGE": {
    "code": "IGE",
    "company": "ALPSANNA LIMITED",
    "city": "Dublin",
    "country": "Ireland"
  },
  "IGL": {
    "code": "IGL",
    "company": "IGLU GROUP PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "IGS": {
    "code": "IGS",
    "company": "CDN CONTAINER DEPOT NUERNBERG GMBH",
    "city": "Nuernberg",
    "country": "Germany"
  },
  "IHO": {
    "code": "IHO",
    "company": "INTERMODAL TANK TRANSPORT INC.",
    "city": "HOUSTON,TX 77064",
    "country": "United States"
  },
  "IHT": {
    "code": "IHT",
    "company": "ITC HOLLAND TRANSPORT B.V",
    "city": "OSS",
    "country": "Netherlands"
  },
  "IIC": {
    "code": "IIC",
    "company": "IWATANI CORPORATION",
    "city": "OSAKA",
    "country": "Japan"
  },
  "IIS": {
    "code": "IIS",
    "company": "INTERTRADE INTERNATIONAL SERVICES",
    "city": "Massagno",
    "country": "Switzerland"
  },
  "IJC": {
    "code": "IJC",
    "company": "IJ-CONTAINER APS",
    "city": "GENTOFTE",
    "country": "Denmark"
  },
  "IJS": {
    "code": "IJS",
    "company": "IJSFABRIEK STROMBEEK NV",
    "city": "MEISE (EVERSEM)",
    "country": "Belgium"
  },
  "IKK": {
    "code": "IKK",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "IKM": {
    "code": "IKM",
    "company": "LLOYDS MARITIME & TRADING LIMITED",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "IKS": {
    "code": "IKS",
    "company": "TAL INTERNATIONAL",
    "city": "PURCHASE, NY 10577-2135",
    "country": "United States"
  },
  "ILB": {
    "code": "ILB",
    "company": "ILB INTERFREIGHT LOGISTICS BENELUX BV",
    "city": "Oss",
    "country": "Netherlands"
  },
  "ILC": {
    "code": "ILC",
    "company": "TRANSAFE SERVICES LIMITED",
    "city": "CALCUTTA",
    "country": "India"
  },
  "ILF": {
    "code": "ILF",
    "company": "ООО «IRKUTSK OIL COMPANY»",
    "city": "Irkutsk",
    "country": "Russian Federation"
  },
  "ILG": {
    "code": "ILG",
    "company": "ILGINNOVATIVE LOGISTICS GROUP GMBH",
    "city": "LINZ",
    "country": "Austria"
  },
  "ILK": {
    "code": "ILK",
    "company": "INNOVATIVE B2B LOGISTICS SOLUTIONS PRIVATE LIMITED",
    "city": "GURGAON, HARYANA",
    "country": "India"
  },
  "ILT": {
    "code": "ILT",
    "company": "SOLVOCHEM HOLLAND BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "ILU": {
    "code": "ILU",
    "company": "INSTITUTE OF INT CONTAINER LESSORS (IICL)",
    "city": "WASHIGTON, DC 20036",
    "country": "United States"
  },
  "IMC": {
    "code": "IMC",
    "company": "INDUSTRIAL MARITIME  CARRIERS INC",
    "city": "NEW ORLEANS, LA 70130",
    "country": "United States"
  },
  "IME": {
    "code": "IME",
    "company": "EJERCITO DE TIERRA - MINISTERO DE DEFENS",
    "city": "MADRID",
    "country": "Spain"
  },
  "IMI": {
    "code": "IMI",
    "company": "SULLAIR ARGENTINA S.A.",
    "city": "BUENOS AIRES",
    "country": "Argentina"
  },
  "IMM": {
    "code": "IMM",
    "company": "INDUSTRIAS MONFEL S.A. DE C.V.",
    "city": "SAN LUIS POTOSI, SLP",
    "country": "Mexico"
  },
  "IMP": {
    "code": "IMP",
    "company": "IMPERIAL CHEMICAL TRANSPORT GMBH",
    "city": "DUISBURG",
    "country": "Germany"
  },
  "IMS": {
    "code": "IMS",
    "company": "IMMOBILIARE STELLA SRL",
    "city": "NARNI SCALO (TR)",
    "country": "Italy"
  },
  "IMT": {
    "code": "IMT",
    "company": "ITALIA MARITTIMA S.P.A",
    "city": "TRIESTE TS",
    "country": "Italy"
  },
  "IMW": {
    "code": "IMW",
    "company": "IMWATER TREATMENT PLANTS S.L",
    "city": "GIJON ASTURIAS",
    "country": "Spain"
  },
  "INB": {
    "code": "INB",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "INC": {
    "code": "INC",
    "company": "INNOVACION CONTENEDOR PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "IND": {
    "code": "IND",
    "company": "INDOX ENERGY SYSTEMS SL",
    "city": "ANGLESOLA",
    "country": "Spain"
  },
  "INF": {
    "code": "INF",
    "company": "INNOFREIGHT SPEDITIONS GMBH",
    "city": "BRUCK AN DER MUR",
    "country": "Austria"
  },
  "ING": {
    "code": "ING",
    "company": "TRIFLEET LEASING HOLDING B.V.",
    "city": "DORDRECHT",
    "country": "Netherlands"
  },
  "INI": {
    "code": "INI",
    "company": "INNOCENTI DEPOSITI SPA",
    "city": "LIMITO DI PIOLTELLO (MI)",
    "country": "Italy"
  },
  "INK": {
    "code": "INK",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "INL": {
    "code": "INL",
    "company": "TRANSLINER PTE  LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "INN": {
    "code": "INN",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "INR": {
    "code": "INR",
    "company": "LOGITEX INVEST, LLC",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "INS": {
    "code": "INS",
    "company": "INSTANTA",
    "city": "ZORY",
    "country": "Poland"
  },
  "INT": {
    "code": "INT",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "INV": {
    "code": "INV",
    "company": "INGAS LLC",
    "city": "Mariupol",
    "country": "Ukraine"
  },
  "IOU": {
    "code": "IOU",
    "company": "TRANS DISTANCE LINE LTD",
    "city": "HONG KONG",
    "country": "HK"
  },
  "IPF": {
    "code": "IPF",
    "company": "INSTITUT POLAIRE FRANCAIS PAUL-EMILE VICTOR",
    "city": "PLOUZANE",
    "country": "France"
  },
  "IPG": {
    "code": "IPG",
    "company": "INSTANT PRODUCTS GROUP PTY LTD",
    "city": "LANDSDALE",
    "country": "Australia"
  },
  "IPI": {
    "code": "IPI",
    "company": "IRIS LOGISTICS, INC.",
    "city": "Olongapo",
    "country": "Philippines"
  },
  "IPM": {
    "code": "IPM",
    "company": "IPM TECHNOLOGIES",
    "city": "Le Passage",
    "country": "France"
  },
  "IPO": {
    "code": "IPO",
    "company": "INTERMODAL PARTNERSHIP OU",
    "city": "Tallinn",
    "country": "Estonia"
  },
  "IPS": {
    "code": "IPS",
    "company": "INTERMODAL SOLUTIONS GROUP PTY LTD",
    "city": "BELLA VISTA, NSW 2153",
    "country": "Australia"
  },
  "IPX": {
    "code": "IPX",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "IRI": {
    "code": "IRI",
    "company": "IRS INTERNATIONAL",
    "city": "Brooklyn",
    "country": "Australia"
  },
  "IRK": {
    "code": "IRK",
    "company": "LTD IRTEK",
    "city": "Shelehov, Irkutskaya oblast",
    "country": "Russian Federation"
  },
  "IRN": {
    "code": "IRN",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "IRS": {
    "code": "IRS",
    "company": "ISLAMIC REPUBLIC OF IRAN SHIPPING LINES (IRISL)",
    "city": "TEHRAN",
    "country": "Iran, Islamic Republic of"
  },
  "ISA": {
    "code": "ISA",
    "company": "NORDIC BULKERS AB",
    "city": "GOTHENBURG",
    "country": "Sweden"
  },
  "ISC": {
    "code": "ISC",
    "company": "STAR CONTAINER SPAIN S.A.",
    "city": "ALGECIRAS",
    "country": "Spain"
  },
  "ISD": {
    "code": "ISD",
    "company": "ISRO DEPARTMENT OF SPACE GOVT OF INDIA",
    "city": "BANGALORE",
    "country": "India"
  },
  "ISL": {
    "code": "ISL",
    "company": "INDUSTRIAL SOLVENTS & CHEMICALS PVT, LTD",
    "city": "Mumbai",
    "country": "India"
  },
  "ISR": {
    "code": "ISR",
    "company": "ISR INVESTMENTS LIMITED",
    "city": "NICOSIA",
    "country": "Cyprus"
  },
  "ISS": {
    "code": "ISS",
    "company": "AMERICA S TANK CONTAINER SERVICE INC. (ATCOS INC)",
    "city": "PANAMA CITY",
    "country": "Panama"
  },
  "ITA": {
    "code": "ITA",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "ITC": {
    "code": "ITC",
    "company": "INTERCONTAINER S.A.",
    "city": "ALFAFAR-VALENCIA",
    "country": "Spain"
  },
  "ITD": {
    "code": "ITD",
    "company": "INTERMODALTRASPORTI SRL",
    "city": "FERENTINO (FR)",
    "country": "Italy"
  },
  "ITE": {
    "code": "ITE",
    "company": "MOVE INTERMODAL NV",
    "city": "Genk",
    "country": "Belgium"
  },
  "ITG": {
    "code": "ITG",
    "company": "ITT ULUSLARARASIŞ TAŞ SAN.TIC LTD ŞTI.",
    "city": "KOCAELİ",
    "country": "Turkey"
  },
  "ITK": {
    "code": "ITK",
    "company": "INTERTANK LTDA",
    "city": "RIO DE JANEIRO",
    "country": "Brazil"
  },
  "ITL": {
    "code": "ITL",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "ITO": {
    "code": "ITO",
    "company": "ITOCHU CORPORATION",
    "city": "Tokyo",
    "country": "Japan"
  },
  "ITQ": {
    "code": "ITQ",
    "company": "ITAC SPA",
    "city": "Ponticino",
    "country": "Italy"
  },
  "ITS": {
    "code": "ITS",
    "company": "E-WAY ALLIANCE (INTERNATIONAL) PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "ITT": {
    "code": "ITT",
    "company": "INTERMODAL TANK TRANSPORT INC.",
    "city": "HOUSTON,TX 77064",
    "country": "United States"
  },
  "ITX": {
    "code": "ITX",
    "company": "INTERPORT MAINTENANCE CO INC",
    "city": "NEWARK, NJ 07105",
    "country": "United States"
  },
  "IVL": {
    "code": "IVL",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "IWG": {
    "code": "IWG",
    "company": "IWS INTERNATIONALE WEIN UND SPIRITUOSEN TRANSPORTE GMBH",
    "city": "PRISDORF",
    "country": "Germany"
  },
  "IWL": {
    "code": "IWL",
    "company": "IMPALA TERMINALS SWITZERLAND SARL",
    "city": "Geneva",
    "country": "Switzerland"
  },
  "IWR": {
    "code": "IWR",
    "company": "IWS RUSSIA GMBH",
    "city": "PRISDORF",
    "country": "Germany"
  },
  "IWT": {
    "code": "IWT",
    "company": "INTERNATIONAL WAREHOUSING AND TRANSPORT",
    "city": "ROTTERDAM-HOOGVLIET",
    "country": "Netherlands"
  },
  "IXM": {
    "code": "IXM",
    "company": "IXOM OPERATIONS PTY LTD",
    "city": "East Melbourne",
    "country": "Australia"
  },
  "IXW": {
    "code": "IXW",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "IXX": {
    "code": "IXX",
    "company": "INDAIA LOGISTIK GMBH",
    "city": "BAD CAMBERG",
    "country": "Germany"
  },
  "IZO": {
    "code": "IZO",
    "company": "IZOTOP ST PETERSBURG JSC",
    "city": "ST. PETERSBURG",
    "country": "Russian Federation"
  },
  "JAA": {
    "code": "JAA",
    "company": "JORDAN BROMINE COMPANY LIMITED",
    "city": "AMMAN",
    "country": "Jordan"
  },
  "JAC": {
    "code": "JAC",
    "company": "ALPHA SHIPPING AGENCY (PTY) LTD",
    "city": "DURBAN KZN",
    "country": "South Africa"
  },
  "JAH": {
    "code": "JAH",
    "company": "CONTAINER STORE SPA",
    "city": "Viña del Mar",
    "country": "Chile"
  },
  "JAM": {
    "code": "JAM",
    "company": "JAMM RENTING SL",
    "city": "PAMPLONA",
    "country": "Spain"
  },
  "JAY": {
    "code": "JAY",
    "company": "JAY CONTAINER SERVICES COMPANY PVT. LTD.",
    "city": "NEW BOMBAY",
    "country": "India"
  },
  "JBB": {
    "code": "JBB",
    "company": "SUEZ RV OSIS SUD EST",
    "city": "VAULX-EN-VELIN",
    "country": "France"
  },
  "JBK": {
    "code": "JBK",
    "company": "SPECIALTY TRAILER LEASING INC",
    "city": "AMARILLO,TX 79124",
    "country": "United States"
  },
  "JBT": {
    "code": "JBT",
    "company": "L & T B.V.",
    "city": "'s-Hertogenbosch",
    "country": "Netherlands"
  },
  "JCF": {
    "code": "JCF",
    "company": "JC FRAGOSO REPAROS ME",
    "city": "MACAE",
    "country": "Brazil"
  },
  "JCJ": {
    "code": "JCJ",
    "company": "CASHIHOR ENTERPRISE CO LTD",
    "city": "KAOHSIUNG CITY",
    "country": "Taiwan, China"
  },
  "JDC": {
    "code": "JDC",
    "company": "TRANSPORTS CAPELLE",
    "city": "VEZENOBRES",
    "country": "France"
  },
  "JDN": {
    "code": "JDN",
    "company": "JAN DE NUL NV",
    "city": "HOFSTADE",
    "country": "Belgium"
  },
  "JDS": {
    "code": "JDS",
    "company": "WEIHAI JIAODONG INTERNATIONAL CONTAINER SHIPPING CO ., LTD",
    "city": "Weihai",
    "country": "China"
  },
  "JDT": {
    "code": "JDT",
    "company": "INT. TRANSPORT COMPANY JAN DOHMEN BV",
    "city": "Herkenbosch",
    "country": "Netherlands"
  },
  "JDZ": {
    "code": "JDZ",
    "company": "JANSENS & DIEPERINK",
    "city": "ZAANDAM",
    "country": "Netherlands"
  },
  "JEN": {
    "code": "JEN",
    "company": "INNIO JENBACHER GMBH & CO OG",
    "city": "Jenbach",
    "country": "Austria"
  },
  "JFS": {
    "code": "JFS",
    "company": "ST JOHN FREIGHT SYSTEMS LTD",
    "city": "TUTICORIN",
    "country": "India"
  },
  "JGF": {
    "code": "JGF",
    "company": "JAPAN GROUND SELF DEFENSE FORCE",
    "city": "Yokohama",
    "country": "Japan"
  },
  "JGR": {
    "code": "JGR",
    "company": "JOHN G.RUSELL (TRANSPORT) LTD",
    "city": "GLASGOW G52 4XB",
    "country": "United Kingdom"
  },
  "JHC": {
    "code": "JHC",
    "company": "HAMMELMANN TRANSPORT GMBH",
    "city": "ENNIGERLOH",
    "country": "Germany"
  },
  "JHG": {
    "code": "JHG",
    "company": "JING HE SCIENCE CO, LTD",
    "city": "TAOYUAN COUNTY",
    "country": "Taiwan, China"
  },
  "JHT": {
    "code": "JHT",
    "company": "JUHUA TRADING (HONG KONG) LIMITED",
    "city": "HONG KONG",
    "country": "HK"
  },
  "JJA": {
    "code": "JJA",
    "company": "JINGJIANG ASIAN-PACIFIC LOGISTICS EQUIPMENT  CO.,LTD",
    "city": "Jingjiang",
    "country": "China"
  },
  "JJD": {
    "code": "JJD",
    "company": "ITEC ENGINEERING",
    "city": "CHAMBLY",
    "country": "France"
  },
  "JJF": {
    "code": "JJF",
    "company": "JJ FORWARDER SL",
    "city": "Molina de segura",
    "country": "Spain"
  },
  "JLG": {
    "code": "JLG",
    "company": "JLG CONTAINER SERVICES CO., LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "JLI": {
    "code": "JLI",
    "company": "JOSEF LINDBERG SANDARNE AB",
    "city": "SANDARNE",
    "country": "Sweden"
  },
  "JMK": {
    "code": "JMK",
    "company": "J.M. KOTHARI  & SONS",
    "city": "MUMBAI, MAHARASHTRA",
    "country": "India"
  },
  "JNI": {
    "code": "JNI",
    "company": "JNI CONSULTING LIMITED",
    "city": "KOWLOON",
    "country": "HK"
  },
  "JNJ": {
    "code": "JNJ",
    "company": "JINAN ENERGY CONSTRUCTION & DEVELOPMENT GROUP CO., LTD.",
    "city": "Jinan",
    "country": "China"
  },
  "JOB": {
    "code": "JOB",
    "company": "JOBACHEM GMBH",
    "city": "Dassel",
    "country": "Germany"
  },
  "JON": {
    "code": "JON",
    "company": "JONTRANS",
    "city": "Le chambon feugerolles",
    "country": "France"
  },
  "JOT": {
    "code": "JOT",
    "company": "JAPAN OIL TRANSPORTATION CO. LTD",
    "city": "TOKYO",
    "country": "Japan"
  },
  "JPR": {
    "code": "JPR",
    "company": "JP RYAN LTD",
    "city": "DUBLIN 3",
    "country": "Ireland"
  },
  "JPS": {
    "code": "JPS",
    "company": "JP CONTAINER LIMITED",
    "city": "PUDONG",
    "country": "China"
  },
  "JRD": {
    "code": "JRD",
    "company": "J.TRAD CO.,LTD.",
    "city": "Tokyo",
    "country": "Japan"
  },
  "JSD": {
    "code": "JSD",
    "company": "JUSDA ENERGY TECHNOLOGY (SHANGHAI) CO LTD",
    "city": "Shanghai",
    "country": "China"
  },
  "JSE": {
    "code": "JSE",
    "company": "SHIN ETSU CHEMICAL CO LTD",
    "city": "CHIYODA-KU TOKYO",
    "country": "Japan"
  },
  "JSK": {
    "code": "JSK",
    "company": "CRYOGENIC LOGISTIC EQUIPMENT & TRADING PTE LTD",
    "city": "",
    "country": "Singapore"
  },
  "JSL": {
    "code": "JSL",
    "company": "JIANGSU LINTEC ENGINEERING EQUIPMENT CO LTD",
    "city": "JIANGYIN",
    "country": "China"
  },
  "JSS": {
    "code": "JSS",
    "company": "THE CHINA NAVIGATION CO PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "JSV": {
    "code": "JSV",
    "company": "JSV LOGISTIC S.L.",
    "city": "Miranda de Ebro. Burgos.",
    "country": "Spain"
  },
  "JTA": {
    "code": "JTA",
    "company": "JUNTAI CONTAINER COMPANY",
    "city": "Hongkong",
    "country": "HK"
  },
  "JTD": {
    "code": "JTD",
    "company": "SHANGHAI JIYANG TECHNOLOGY & DEVELOPMENT CO., LTD.",
    "city": "SHANGHAI",
    "country": "China"
  },
  "JTL": {
    "code": "JTL",
    "company": "JIANGSU TERCEL LOGISTICS EQUIPMENT CO., LTD",
    "city": "Taixing City, Jiangsu",
    "country": "China"
  },
  "JTM": {
    "code": "JTM",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "JTS": {
    "code": "JTS",
    "company": "JOINT TANK SERVICES FZCO",
    "city": "Dubai, Jebel Ali",
    "country": "United Arab Emirates"
  },
  "JUH": {
    "code": "JUH",
    "company": "BSH HAUSGERATE GMBH",
    "city": "GIENGEN",
    "country": "Germany"
  },
  "JUR": {
    "code": "JUR",
    "company": "JURI KRAUS",
    "city": "LUDWIGSHAFEN",
    "country": "Germany"
  },
  "JXJ": {
    "code": "JXJ",
    "company": "GOLD STAR LINE LTD",
    "city": "Kowloon, Hong Kong",
    "country": "HK"
  },
  "JXL": {
    "code": "JXL",
    "company": "GOLD STAR LINE LTD",
    "city": "Kowloon, Hong Kong",
    "country": "HK"
  },
  "JYT": {
    "code": "JYT",
    "company": "GUANGXI JINGYITONG LOGISTICS LTD",
    "city": "Fang cheng Port City",
    "country": "China"
  },
  "JZN": {
    "code": "JZN",
    "company": "A.J JONGENEEL EN ZONEN TRANSPORT B.V",
    "city": "VALKENBURG (Z.H)",
    "country": "Netherlands"
  },
  "JZP": {
    "code": "JZP",
    "company": "JINZHOU PORT CO.,LTD.",
    "city": "Jinzhou",
    "country": "China"
  },
  "JZT": {
    "code": "JZT",
    "company": "TAIHE GASES (JINGZHOU) LTD",
    "city": "Jingzhou",
    "country": "China"
  },
  "KAL": {
    "code": "KAL",
    "company": "NKG KALA HAMBURG GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "KAS": {
    "code": "KAS",
    "company": "TAIYO NIPPON SANSO INDIA PVT. LTD.",
    "city": "PUNE, MAHARASHTRA",
    "country": "India"
  },
  "KAZ": {
    "code": "KAZ",
    "company": "JSC KUIBYSHEV AZOT",
    "city": "Togliatti",
    "country": "Russian Federation"
  },
  "KBS": {
    "code": "KBS",
    "company": "K-BOX SOLUTION",
    "city": "Phnom Penh",
    "country": "Cambodia"
  },
  "KCB": {
    "code": "KCB",
    "company": "IMOTO LINES, LTD.",
    "city": "Kobe",
    "country": "Japan"
  },
  "KCN": {
    "code": "KCN",
    "company": "K-CON GMBH",
    "city": "KREUZAU",
    "country": "Germany"
  },
  "KCO": {
    "code": "KCO",
    "company": "UMICORE FINLAND OY",
    "city": "KOKKOLA",
    "country": "Finland"
  },
  "KCP": {
    "code": "KCP",
    "company": "KOIKE MEDICAL CO LTD",
    "city": "TOKYO",
    "country": "Japan"
  },
  "KCS": {
    "code": "KCS",
    "company": "KANTO KAGAKU SINGAPORE PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "KCT": {
    "code": "KCT",
    "company": "KC TRADING BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "KDC": {
    "code": "KDC",
    "company": "KUKDONG MARITIME EQUIPMENT SERVICE",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "KDE": {
    "code": "KDE",
    "company": "ABANG LOGISTICS CO., LTD.",
    "city": "Gunpo",
    "country": "Korea, Republic of"
  },
  "KDY": {
    "code": "KDY",
    "company": "KDY LOGISTICS AND CONTAINER A.S.",
    "city": "Majuro",
    "country": "Marshall Island"
  },
  "KEM": {
    "code": "KEM",
    "company": "KEMITO BVBA",
    "city": "Brussels",
    "country": "Belgium"
  },
  "KEN": {
    "code": "KEN",
    "company": "KENT REMOVALS & STORAGE",
    "city": "Clayton",
    "country": "Australia"
  },
  "KER": {
    "code": "KER",
    "company": "TEMPRA AS",
    "city": "AALESUND",
    "country": "Norway"
  },
  "KEY": {
    "code": "KEY",
    "company": "FASTGROW / KEYUN GROUP",
    "city": "TIANJIN",
    "country": "China"
  },
  "KFR": {
    "code": "KFR",
    "company": "KILFROST LIMITED",
    "city": "HALTWHISTLE",
    "country": "United Kingdom"
  },
  "KFT": {
    "code": "KFT",
    "company": "KAERCHER FUTURETECH GMBH",
    "city": "SCHWAIKHEIM",
    "country": "Germany"
  },
  "KGP": {
    "code": "KGP",
    "company": "JSC  KAMCHATGAZPROM",
    "city": "PETROPAVLOVSK-KAMCHATSKY",
    "country": "Russian Federation"
  },
  "KGS": {
    "code": "KGS",
    "company": "KGS KELLER GERATE & SERVICE GMBH",
    "city": "Offenbach",
    "country": "Germany"
  },
  "KHJ": {
    "code": "KHJ",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "KHL": {
    "code": "KHL",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "KIG": {
    "code": "KIG",
    "company": "KEN INDUSTRIAL GASES PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "KIL": {
    "code": "KIL",
    "company": "KIMIA INTERNATIONAL PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "KIV": {
    "code": "KIV",
    "company": "KEES INTL VEEN TANKTRANSPORTEN BV",
    "city": "EUROPOORT-ROTTERDAM",
    "country": "Netherlands"
  },
  "KJC": {
    "code": "KJC",
    "company": "ITC REFRIGERATION PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "KKB": {
    "code": "KKB",
    "company": "KERNKRAFTWERK BRUNSBUTTEL GMBH & CO. OHG",
    "city": "Brunsbüttel",
    "country": "Germany"
  },
  "KKF": {
    "code": "KKF",
    "company": "KAWASAKI KISEN KAISHA LTD - K LINE",
    "city": "TOKYO",
    "country": "Japan"
  },
  "KKK": {
    "code": "KKK",
    "company": "OZTURK KONTEYNER SAN VE TIC LTD STI",
    "city": "Istanbul",
    "country": "Turkey"
  },
  "KKL": {
    "code": "KKL",
    "company": "KAWASAKI KISEN KAISHA LTD - K LINE",
    "city": "TOKYO",
    "country": "Japan"
  },
  "KKT": {
    "code": "KKT",
    "company": "KAWASAKI KISEN KAISHA LTD - K LINE",
    "city": "TOKYO",
    "country": "Japan"
  },
  "KLA": {
    "code": "KLA",
    "company": "MINISTERIE VAN DEFENSIE",
    "city": "UTRECHT",
    "country": "Netherlands"
  },
  "KLC": {
    "code": "KLC",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "KLD": {
    "code": "KLD",
    "company": "KALKEDON INTERNATIONAL TRANSPORT & TRADE LTD",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "KLF": {
    "code": "KLF",
    "company": "KAWASAKI KISEN KAISHA LTD - K LINE",
    "city": "TOKYO",
    "country": "Japan"
  },
  "KLI": {
    "code": "KLI",
    "company": "KLINGE CORPORATION",
    "city": "YORK, PA-17402",
    "country": "United States"
  },
  "KLK": {
    "code": "KLK",
    "company": "CONTAINER LEASING COMPANY",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "KLN": {
    "code": "KLN",
    "company": "KART (THAILAND) LIMITED",
    "city": "BANGKOK",
    "country": "Thailand"
  },
  "KLO": {
    "code": "KLO",
    "company": "KLOIBER GMBH",
    "city": "PETERSHAUSEN",
    "country": "Germany"
  },
  "KLT": {
    "code": "KLT",
    "company": "KAWASAKI KISEN KAISHA LTD - K LINE",
    "city": "TOKYO",
    "country": "Japan"
  },
  "KLY": {
    "code": "KLY",
    "company": "SAMERUS REEFER CONTAINER S.A.",
    "city": "Majuro",
    "country": "Marshall Island"
  },
  "KMA": {
    "code": "KMA",
    "company": "ROYAL NETHERLANDS NAVY",
    "city": "DEN HAAG",
    "country": "Netherlands"
  },
  "KMB": {
    "code": "KMB",
    "company": "KAMBARA KISEN CO LTD",
    "city": "HIROSHIMA-PREF",
    "country": "Japan"
  },
  "KMC": {
    "code": "KMC",
    "company": "KMARIN OCEAN LOGISTICS (KOL)",
    "city": "Seoul",
    "country": "Korea, Republic of"
  },
  "KMI": {
    "code": "KMI",
    "company": "KUANG MING ENTERPRISE CO. LTD",
    "city": "TAIPEI",
    "country": "Taiwan, China"
  },
  "KML": {
    "code": "KML",
    "company": "MAXWAY MARITIME S.A",
    "city": "DALIAN",
    "country": "China"
  },
  "KMO": {
    "code": "KMO",
    "company": "KAMYOO DEVELOPMENT HOLDING PTE. LTD.",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "KMP": {
    "code": "KMP",
    "company": "TANK ONE NV",
    "city": "Antwerpen",
    "country": "Belgium"
  },
  "KMT": {
    "code": "KMT",
    "company": "KOREA MARINE TRANSPORT CO / E. C. TEAM",
    "city": "SEOUL CITY",
    "country": "Korea, Republic of"
  },
  "KNH": {
    "code": "KNH",
    "company": "PROCHEMICAL GROUP, S.R.O.",
    "city": "Bratislava",
    "country": "Slovakia"
  },
  "KNL": {
    "code": "KNL",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "KNT": {
    "code": "KNT",
    "company": "KANTO CORPORATION INC.",
    "city": "PORTLAND, OR 97203",
    "country": "United States"
  },
  "KOC": {
    "code": "KOC",
    "company": "KOREA OCEAN BUSINESS CORPORATION",
    "city": "Busan",
    "country": "Korea, Republic of"
  },
  "KON": {
    "code": "KON",
    "company": "KONTIX SP. Z O.O.",
    "city": "Kobylnica",
    "country": "Poland"
  },
  "KOS": {
    "code": "KOS",
    "company": "KING OCEAN SERVICES LTD",
    "city": "DORAL, FL 33172",
    "country": "United States"
  },
  "KPC": {
    "code": "KPC",
    "company": "KIRLOSKAR PNEUMATIC CO. LTD. HADAPSAR, PUNE",
    "city": "PUNE 411013",
    "country": "India"
  },
  "KPK": {
    "code": "KPK",
    "company": "KRAJMAN PIOTR KRAJEWSKI",
    "city": "Warsaw",
    "country": "Poland"
  },
  "KPS": {
    "code": "KPS",
    "company": "KANTO-PPC (SHANGHAI) INC.",
    "city": "SHANGHAI",
    "country": "China"
  },
  "KQT": {
    "code": "KQT",
    "company": "KEIKU CO., LTD.",
    "city": "Taoyuan",
    "country": "Taiwan, China"
  },
  "KRC": {
    "code": "KRC",
    "company": "KIWIRAIL",
    "city": "Auckland",
    "country": "New Zealand"
  },
  "KRF": {
    "code": "KRF",
    "company": "KOREA RAILROAD RESEARCH INSTITUTE",
    "city": "Uiwang-si, Gyeonggi-do",
    "country": "Korea, Republic of"
  },
  "KRI": {
    "code": "KRI",
    "company": "KRICON SERVICES B.V.",
    "city": "Lepelstraat",
    "country": "Netherlands"
  },
  "KRL": {
    "code": "KRL",
    "company": "KRIBHCO INFRASTRUCTURE LTD",
    "city": "NOIDA",
    "country": "India"
  },
  "KRM": {
    "code": "KRM",
    "company": "HANS-GEORG KRAMER MOBELTRANSPORTE",
    "city": "Bielefeld",
    "country": "Germany"
  },
  "KRO": {
    "code": "KRO",
    "company": "KRONOCHEM LLC",
    "city": "Mogilev district",
    "country": "Belarus"
  },
  "KRP": {
    "code": "KRP",
    "company": "JSC «KRASNOYARSK RIVER PORT»",
    "city": "Krasnoyarsk-city",
    "country": "Russian Federation"
  },
  "KRT": {
    "code": "KRT",
    "company": "KRAMPITZ TRANKSYSTEM GMBH",
    "city": "DAHLENBURG",
    "country": "Germany"
  },
  "KSA": {
    "code": "KSA",
    "company": "KSAN SIA",
    "city": "RIGA",
    "country": "Latvia"
  },
  "KSF": {
    "code": "KSF",
    "company": "KING STAR FREIGHT PVT. LTD.",
    "city": "Mumbai",
    "country": "India"
  },
  "KSI": {
    "code": "KSI",
    "company": "LLC KISLOROD PREMIUM",
    "city": "Blagoveshchensk, Amurskaya oblast'",
    "country": "Russian Federation"
  },
  "KST": {
    "code": "KST",
    "company": "SCHILDECKER TRANSPORT GMBH",
    "city": "PISCHELSDORF",
    "country": "Austria"
  },
  "KSZ": {
    "code": "KSZ",
    "company": "SILVA BY FLLC",
    "city": "SMORGON",
    "country": "Belarus"
  },
  "KTC": {
    "code": "KTC",
    "company": "AEROSOL GAS COMPANY, INC",
    "city": "San Ramon, CA 94583",
    "country": "United States"
  },
  "KTD": {
    "code": "KTD",
    "company": "CONDACO & KTD-M GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "KTI": {
    "code": "KTI",
    "company": "KTI-PLERSCH KALTETECHNIK GMBH",
    "city": "Balzheim",
    "country": "Germany"
  },
  "KTK": {
    "code": "KTK",
    "company": "LLC KTK TRAIDING (INN 7805711190)",
    "city": "Saint Petersburg",
    "country": "Russian Federation"
  },
  "KTN": {
    "code": "KTN",
    "company": "K-TAINER TRADING BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "KTO": {
    "code": "KTO",
    "company": "KTO NV",
    "city": "IZEGEM",
    "country": "Belgium"
  },
  "KTS": {
    "code": "KTS",
    "company": "KTZ EXPRESS JSC",
    "city": "Nur-Sultan",
    "country": "Kazakhstan"
  },
  "KTT": {
    "code": "KTT",
    "company": "KOTTA CONTAINER COMPANY LTD",
    "city": "Saint Petersburg",
    "country": "Russian Federation"
  },
  "KTZ": {
    "code": "KTZ",
    "company": "KTZ EXPRESS JSC",
    "city": "Nur-Sultan",
    "country": "Kazakhstan"
  },
  "KUB": {
    "code": "KUB",
    "company": "KUBE & KUBENZ INT. SPEDITION-UND",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "KUD": {
    "code": "KUD",
    "company": "GATEWAY CONTAINER SALES & HIRE PTY LTD",
    "city": "HEMMANT, Queensland",
    "country": "Australia"
  },
  "KUU": {
    "code": "KUU",
    "company": "KAMCHATKA SHIPPING COMPANY",
    "city": "Petropavlovsk-Kamchatsky",
    "country": "Russian Federation"
  },
  "KVO": {
    "code": "KVO",
    "company": "RUSVINYL LLC",
    "city": "Kstovo",
    "country": "Russian Federation"
  },
  "KWC": {
    "code": "KWC",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "KWK": {
    "code": "KWK",
    "company": "KWIK EQUIPMENT SALES LLC",
    "city": "Pearland, TX-77581, Texas",
    "country": "United States"
  },
  "KWT": {
    "code": "KWT",
    "company": "W. KOBRUNNER HANDELS-U. GUTERTRANSPORTE GMBH",
    "city": "REGAU",
    "country": "Austria"
  },
  "KWU": {
    "code": "KWU",
    "company": "FRAMATOME GMBH",
    "city": "ERLANGEN",
    "country": "Germany"
  },
  "KXT": {
    "code": "KXT",
    "company": "KAWASAKI KISEN KAISHA LTD - K LINE",
    "city": "TOKYO",
    "country": "Japan"
  },
  "KYS": {
    "code": "KYS",
    "company": "KYOGOKU UNYU SHOJI CO LTD",
    "city": "TOKYO",
    "country": "Japan"
  },
  "KZE": {
    "code": "KZE",
    "company": "KTZ EXPRESS JSC",
    "city": "Nur-Sultan",
    "country": "Kazakhstan"
  },
  "LAB": {
    "code": "LAB",
    "company": "LAGERMAX AUTOTRANSPORT GMBH",
    "city": "STRASSWALCHEN",
    "country": "Austria"
  },
  "LAF": {
    "code": "LAF",
    "company": "L & T B.V.",
    "city": "'s-Hertogenbosch",
    "country": "Netherlands"
  },
  "LAH": {
    "code": "LAH",
    "company": "L & T B.V.",
    "city": "'s-Hertogenbosch",
    "country": "Netherlands"
  },
  "LAM": {
    "code": "LAM",
    "company": "LAM KARA HAVA VE DENIZ TASIMACILIK AS",
    "city": "Istanbul",
    "country": "Turkey"
  },
  "LAR": {
    "code": "LAR",
    "company": "LAROUTE SA",
    "city": "ZUG",
    "country": "Switzerland"
  },
  "LAS": {
    "code": "LAS",
    "company": "SOCOMAT",
    "city": "BOULOGNE",
    "country": "France"
  },
  "LAU": {
    "code": "LAU",
    "company": "LAUTERBACH SPEDITIONS-GMBH",
    "city": "BERG",
    "country": "Germany"
  },
  "LBC": {
    "code": "LBC",
    "company": "LIEBHERR-WERK BIBERACH GMBH",
    "city": "BIBERACH",
    "country": "Germany"
  },
  "LBG": {
    "code": "LBG",
    "company": "LUBIAO CONTAINER CO., LTD",
    "city": "Road Town Tortola",
    "country": "Virgin Islands, British"
  },
  "LBI": {
    "code": "LBI",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "LBO": {
    "code": "LBO",
    "company": "TERMCOTANK SA",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "LCE": {
    "code": "LCE",
    "company": "LIVENT",
    "city": "CHARLOTTE, NC-28208",
    "country": "United States"
  },
  "LCG": {
    "code": "LCG",
    "company": "LOTUS CONTAINERS GMBH",
    "city": "EGESTORF",
    "country": "Germany"
  },
  "LCI": {
    "code": "LCI",
    "company": "HEXAGON LINCOLN",
    "city": "LINCOLN, NE 68524",
    "country": "United States"
  },
  "LCK": {
    "code": "LCK",
    "company": "LOCATANK S.A.",
    "city": "EYBENS",
    "country": "France"
  },
  "LCL": {
    "code": "LCL",
    "company": "CONTAINER LEASING",
    "city": "SAINT PETERSBURG",
    "country": "Russian Federation"
  },
  "LCR": {
    "code": "LCR",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "LCT": {
    "code": "LCT",
    "company": "LEVADA CARGO LLC",
    "city": "Kiev",
    "country": "Ukraine"
  },
  "LCU": {
    "code": "LCU",
    "company": "LANCER CONTAINER LINES LTD",
    "city": "NAVI MUMBAI",
    "country": "India"
  },
  "LCX": {
    "code": "LCX",
    "company": "LOGIX CAPITAL LLC",
    "city": "DORAL",
    "country": "United States"
  },
  "LCY": {
    "code": "LCY",
    "company": "LCY CHEMICAL CORP.",
    "city": "Kaohsiung",
    "country": "Taiwan, China"
  },
  "LDR": {
    "code": "LDR",
    "company": "TURKEY RUSSIA LOGISTICS CORPORATION",
    "city": "SAMSUN",
    "country": "Turkey"
  },
  "LDV": {
    "code": "LDV",
    "company": "LINDE GAS VIETNAM LTD",
    "city": "Ba Ria Vung Tau Province",
    "country": "Viet Nam"
  },
  "LDZ": {
    "code": "LDZ",
    "company": "STATE JOINT STOCK COMPANY-LATVIJAS DZELZCELS",
    "city": "RIGA",
    "country": "Latvia"
  },
  "LEA": {
    "code": "LEA",
    "company": "SO.GE.SE SRL",
    "city": "LIVORNO",
    "country": "Italy"
  },
  "LEC": {
    "code": "LEC",
    "company": "INDUSTRIAL DYNAMICS",
    "city": "DELMAS",
    "country": "Haiti"
  },
  "LED": {
    "code": "LED",
    "company": "SPECTEKHKOMPLEKT LTD.CO",
    "city": "St Petersburg",
    "country": "Russian Federation"
  },
  "LEE": {
    "code": "LEE",
    "company": "TRANSPORTBEDRIJF H.A. LEEMANS B.V.",
    "city": "VRIEZENVEEN",
    "country": "Netherlands"
  },
  "LEG": {
    "code": "LEG",
    "company": "LEGEND LOGISTICS LIMITED",
    "city": "Singapore",
    "country": "Singapore"
  },
  "LEH": {
    "code": "LEH",
    "company": "IMPERIAL CHEMICAL TRANSPORT GMBH",
    "city": "DUISBURG",
    "country": "Germany"
  },
  "LEI": {
    "code": "LEI",
    "company": "SET LINNINGS INTERNATIONAL S.A",
    "city": "MOITA",
    "country": "Portugal"
  },
  "LEL": {
    "code": "LEL",
    "company": "LOGISTICS EXPEDITORS SDN BHD",
    "city": "KLANG,  SELANGOR",
    "country": "Malaysia"
  },
  "LEP": {
    "code": "LEP",
    "company": "LAVA ENGINEERING COMPANY",
    "city": "Erode",
    "country": "India"
  },
  "LER": {
    "code": "LER",
    "company": "BBC CHARTERING GMBH - AS AGENTS",
    "city": "LEER",
    "country": "Germany"
  },
  "LET": {
    "code": "LET",
    "company": "BLUE BALTIC SHIPPING & TRADING LIMITED",
    "city": "LA HABANA",
    "country": "Cuba"
  },
  "LFC": {
    "code": "LFC",
    "company": "HONGKONG SHANGHAI LIFENG LEASING COMPANY LIMITED",
    "city": "Shanghai",
    "country": "China"
  },
  "LFG": {
    "code": "LFG",
    "company": "P&R EQUIPMENT & FINANCE CORP.",
    "city": "Zug",
    "country": "Switzerland"
  },
  "LFI": {
    "code": "LFI",
    "company": "LOFTY STARS GLOBAL LIMITED",
    "city": "MAHE",
    "country": "Seychelles"
  },
  "LFT": {
    "code": "LFT",
    "company": "LFT GMBH INTERNATIONALE",
    "city": "MÜNCHSMÜNSTER",
    "country": "Germany"
  },
  "LGA": {
    "code": "LGA",
    "company": "LINGGAS (TIANJIN) LIMITED",
    "city": "DISTRICT TIANJIN",
    "country": "China"
  },
  "LGE": {
    "code": "LGE",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "LGI": {
    "code": "LGI",
    "company": "LINDE GAS ITALIA SRL",
    "city": "ARLUNO (MI)",
    "country": "Italy"
  },
  "LGK": {
    "code": "LGK",
    "company": "AB LG CARGO",
    "city": "VILNIUS",
    "country": "Lithuania"
  },
  "LGN": {
    "code": "LGN",
    "company": "LOG IN LOGISTICA INTERMODAL S.A.",
    "city": "RIO DE JANEIRO",
    "country": "Brazil"
  },
  "LGS": {
    "code": "LGS",
    "company": "LINDE GAS SINGAPORE PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "LGV": {
    "code": "LGV",
    "company": "ORANO DS - DEMENTELEMENT ET SERVICES",
    "city": "GIF SUR YVETTE",
    "country": "France"
  },
  "LGX": {
    "code": "LGX",
    "company": "LOGIX S.A.",
    "city": "BUENOS AIRES",
    "country": "Argentina"
  },
  "LHE": {
    "code": "LHE",
    "company": "BLUE SKY (HK) INTERNATIONAL TRADING LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "LHZ": {
    "code": "LHZ",
    "company": "LVIV CHEMICAL PLANT JSC",
    "city": "Lviv",
    "country": "Ukraine"
  },
  "LIC": {
    "code": "LIC",
    "company": "INTERMODAL CONTAINER SERVICE",
    "city": "VILNIUS",
    "country": "Lithuania"
  },
  "LIN": {
    "code": "LIN",
    "company": "LINDE GMBH, LINDE GAS DEUTSCHLAND",
    "city": "PULLACH",
    "country": "Germany"
  },
  "LIQ": {
    "code": "LIQ",
    "company": "LIQUIMET S.p.A.",
    "city": "Treviso",
    "country": "Italy"
  },
  "LIX": {
    "code": "LIX",
    "company": "CROSSTRADE SHIPPING LTD",
    "city": "ROAD TOWN, TORTOLA",
    "country": "Virgin Islands, British"
  },
  "LKC": {
    "code": "LKC",
    "company": "LAURITZEN KOSAN A/S",
    "city": "Hellerup",
    "country": "Denmark"
  },
  "LKD": {
    "code": "LKD",
    "company": "BULKUID",
    "city": "Houston, TX 77042",
    "country": "United States"
  },
  "LLN": {
    "code": "LLN",
    "company": "BDC INTERNATIONAL SA",
    "city": "LOUVAIN LA NEUVE",
    "country": "Belgium"
  },
  "LLT": {
    "code": "LLT",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "LMC": {
    "code": "LMC",
    "company": "IGNAZIO MESSINA & CO",
    "city": "GENOVA  GE",
    "country": "Italy"
  },
  "LME": {
    "code": "LME",
    "company": "LANTIA MARITIMA",
    "city": "Torrejon de Ardoz",
    "country": "Spain"
  },
  "LMR": {
    "code": "LMR",
    "company": "L.M.C. (LEMARECHAL CELESTIN)",
    "city": "VALOGNES",
    "country": "France"
  },
  "LND": {
    "code": "LND",
    "company": "LINDE GAS NA LLC",
    "city": "Bridgewater, NJ 08807",
    "country": "United States"
  },
  "LNG": {
    "code": "LNG",
    "company": "GEFEST LTD",
    "city": "Saint Petersburg",
    "country": "Russian Federation"
  },
  "LNN": {
    "code": "LNN",
    "company": "LLC LUKOIL-NIZHEGORODNEFTEORGSINTEZ",
    "city": "Nizhny Novgorod",
    "country": "Russian Federation"
  },
  "LNX": {
    "code": "LNX",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "LOC": {
    "code": "LOC",
    "company": "LOCALLTAINER LOCACOES DE CONTAINERS COM.E.SERV.LTDA",
    "city": "SANTOS",
    "country": "Brazil"
  },
  "LOE": {
    "code": "LOE",
    "company": "LOESCHE GMBH",
    "city": "NEUSS",
    "country": "Germany"
  },
  "LOG": {
    "code": "LOG",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "LOM": {
    "code": "LOM",
    "company": "LEVORATO MARCEVAGGI SRL",
    "city": "CAMPALTO VE",
    "country": "Italy"
  },
  "LON": {
    "code": "LON",
    "company": "LONZA LTD",
    "city": "BASEL",
    "country": "Switzerland"
  },
  "LOS": {
    "code": "LOS",
    "company": "EMERSON CLIMATE TECHNOLOGIES",
    "city": "HOJBJERG",
    "country": "Denmark"
  },
  "LOT": {
    "code": "LOT",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "LOU": {
    "code": "LOU",
    "company": "INTERNATIONAL SHIPPING CO. INC.",
    "city": "WILMINGTON, NC-28412",
    "country": "United States"
  },
  "LPC": {
    "code": "LPC",
    "company": "CATERPILLAR (NI) LTD",
    "city": "LARNE",
    "country": "United Kingdom"
  },
  "LPK": {
    "code": "LPK",
    "company": "SYARIKAT LOGISTIK PETIKEMAS SDN BHD",
    "city": "KUANTAN",
    "country": "Malaysia"
  },
  "LPL": {
    "code": "LPL",
    "company": "LAUDE SMART INTERMODAL S.A.",
    "city": "TORUN",
    "country": "Poland"
  },
  "LRS": {
    "code": "LRS",
    "company": "JSC LENA UNITED SHIPPING COMPANY",
    "city": "YAKUTSK",
    "country": "Russian Federation"
  },
  "LRT": {
    "code": "LRT",
    "company": "LAUDE LLC",
    "city": "Kaliningrad",
    "country": "Russian Federation"
  },
  "LSC": {
    "code": "LSC",
    "company": "LORENZO SHIPPING CORP.",
    "city": "MANILA",
    "country": "Philippines"
  },
  "LSE": {
    "code": "LSE",
    "company": "ONE WAY LEASE INC.",
    "city": "OAKLAND, CA-94607",
    "country": "United States"
  },
  "LSF": {
    "code": "LSF",
    "company": "LANGH GROUP OY AB",
    "city": "PIIKKIO",
    "country": "Finland"
  },
  "LSG": {
    "code": "LSG",
    "company": "LANDSEA GLOBAL TASIMACILIK VE TICARET A.",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "LSH": {
    "code": "LSH",
    "company": "LS HOLDING APS",
    "city": "Tilst",
    "country": "Denmark"
  },
  "LSI": {
    "code": "LSI",
    "company": "LS INTERTANK APS",
    "city": "FREDERICIA",
    "country": "Denmark"
  },
  "LSL": {
    "code": "LSL",
    "company": "SEA-LAND LOGISTICS ENTERPRISES PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "LSN": {
    "code": "LSN",
    "company": "LS-NIKKO COPPER",
    "city": "ULSAN CITY",
    "country": "Korea, Republic of"
  },
  "LSP": {
    "code": "LSP",
    "company": "KUNSHAN ASIA UNION ELECTRONIC CHEMICAL CO LTD",
    "city": "Kunshan",
    "country": "China"
  },
  "LSR": {
    "code": "LSR",
    "company": "LUOYANG SUNRUI SPECIAL EQUIPMENT CO.,LTD",
    "city": "Luoyang",
    "country": "China"
  },
  "LST": {
    "code": "LST",
    "company": "LS INTERTANK APS",
    "city": "FREDERICIA",
    "country": "Denmark"
  },
  "LTB": {
    "code": "LTB",
    "company": "IBERIA CRANES & CONTAINER, SL",
    "city": "ALGECIRAS",
    "country": "Spain"
  },
  "LTC": {
    "code": "LTC",
    "company": "LEASE TC BV",
    "city": "Bergen",
    "country": "Netherlands"
  },
  "LTE": {
    "code": "LTE",
    "company": "TEU CONSERVICES LTD",
    "city": "LIMASSOL",
    "country": "Cyprus"
  },
  "LTG": {
    "code": "LTG",
    "company": "INNER MONGOLIA LANTAI SODIUM INDUSTRY CO LTD",
    "city": "INNER MONGOLIA",
    "country": "China"
  },
  "LTI": {
    "code": "LTI",
    "company": "ITALIA MARITTIMA S.P.A",
    "city": "TRIESTE TS",
    "country": "Italy"
  },
  "LTS": {
    "code": "LTS",
    "company": "LOGICSHIPPINGTRADING APS",
    "city": "Aalborg Øst",
    "country": "Denmark"
  },
  "LTU": {
    "code": "LTU",
    "company": "LITHUANIAN ARMED FORCES LOGISTICS SUPPORT COMMAND",
    "city": "VILNIUS",
    "country": "Lithuania"
  },
  "LTX": {
    "code": "LTX",
    "company": "LONDON TANK MANAGEMENT LTD",
    "city": "SOUTHEND ON SEA",
    "country": "United Kingdom"
  },
  "LUN": {
    "code": "LUN",
    "company": "NUOVA LOGISTICA LUCIANU SRL",
    "city": "Olbia",
    "country": "Italy"
  },
  "LUP": {
    "code": "LUP",
    "company": "ARKEMA GMBH",
    "city": "GUNZBURG",
    "country": "Germany"
  },
  "LUV": {
    "code": "LUV",
    "company": "L&P LOGISTICS, INC",
    "city": "Punta Gorda, FL 33950",
    "country": "United States"
  },
  "LVI": {
    "code": "LVI",
    "company": "LINGGAS, LIMITED",
    "city": "Beijing, 100101",
    "country": "China"
  },
  "LVN": {
    "code": "LVN",
    "company": "FLEXBOX COLUMBIA SAS",
    "city": "BOGOTA",
    "country": "Colombia"
  },
  "LVS": {
    "code": "LVS",
    "company": "LCS CONTAINERS",
    "city": "BURTON-ON-TRENT",
    "country": "United Kingdom"
  },
  "LWC": {
    "code": "LWC",
    "company": "PEOPLE TECHNOLOGY SOLUTIONS LTD",
    "city": "LONDON",
    "country": "United Kingdom"
  },
  "LXA": {
    "code": "LXA",
    "company": "LUXEMBOURG ARMY",
    "city": "DIEKIRCH",
    "country": "Luxembourg"
  },
  "LXL": {
    "code": "LXL",
    "company": "KMT GAS",
    "city": "Sichuan Province,",
    "country": "China"
  },
  "LXS": {
    "code": "LXS",
    "company": "LANXESS  DEUTSCHLAND GMBH",
    "city": "LEVERKUSEN",
    "country": "Germany"
  },
  "LYG": {
    "code": "LYG",
    "company": "DONG FANG INTERNATIONAL CONTAINER (LIANYUNGANG) CO., LTD",
    "city": "LIANYUNGANG",
    "country": "China"
  },
  "LYK": {
    "code": "LYK",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "LYS": {
    "code": "LYS",
    "company": "DFDS A/S",
    "city": "COPENHAGEN",
    "country": "Denmark"
  },
  "MAC": {
    "code": "MAC",
    "company": "ADR TRASPORTI SRL",
    "city": "CATANIA",
    "country": "Italy"
  },
  "MAE": {
    "code": "MAE",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MAG": {
    "code": "MAG",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "MAK": {
    "code": "MAK",
    "company": "MAKIOS LOGISTICS",
    "city": "THESSALONIKI",
    "country": "Greece"
  },
  "MAL": {
    "code": "MAL",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MAM": {
    "code": "MAM",
    "company": "MAMMOET EUROPE BV",
    "city": "SCHIEDAM",
    "country": "Netherlands"
  },
  "MAN": {
    "code": "MAN",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "MAR": {
    "code": "MAR",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "MAS": {
    "code": "MAS",
    "company": "MAS SHIPPING & TRADING",
    "city": "ALEXANDRIA",
    "country": "Egypt"
  },
  "MAT": {
    "code": "MAT",
    "company": "MATSON NAVIGATION COMPANY, INC",
    "city": "OAKLAND, CA 94610",
    "country": "United States"
  },
  "MAU": {
    "code": "MAU",
    "company": "MAURITIUS OIL REFINERIED LTD",
    "city": "",
    "country": "Mauritius"
  },
  "MAX": {
    "code": "MAX",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "MBA": {
    "code": "MBA",
    "company": "MARTI INFRA AG",
    "city": "MOOSSEEDORF",
    "country": "Switzerland"
  },
  "MBB": {
    "code": "MBB",
    "company": "MULTIBOXX LIMITED",
    "city": "KOWLOON",
    "country": "HK"
  },
  "MBD": {
    "code": "MBD",
    "company": "MULTIBOXX LIMITED",
    "city": "KOWLOON",
    "country": "HK"
  },
  "MBF": {
    "code": "MBF",
    "company": "MBF CARPENTERS SHIPPING LTD",
    "city": "KUALA LUMPUR",
    "country": "Malaysia"
  },
  "MBG": {
    "code": "MBG",
    "company": "MULTIBOXX LIMITED",
    "city": "KOWLOON",
    "country": "HK"
  },
  "MBI": {
    "code": "MBI",
    "company": "MORBRIDGE INTERNATIONAL LTD",
    "city": "LIVERPOOL",
    "country": "United Kingdom"
  },
  "MBT": {
    "code": "MBT",
    "company": "MULTIBOXX LIMITED",
    "city": "KOWLOON",
    "country": "HK"
  },
  "MBW": {
    "code": "MBW",
    "company": "MAX BOEGL STIFTUNG & CO KG",
    "city": "SENGENTHAL",
    "country": "Germany"
  },
  "MBX": {
    "code": "MBX",
    "company": "MOBILBOX CONTAINER TRADING LTD",
    "city": "BUDAPEST",
    "country": "Hungary"
  },
  "MCA": {
    "code": "MCA",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MCB": {
    "code": "MCB",
    "company": "MOELLER CHEMIE GMBH & CO. KG",
    "city": "Steinfurt",
    "country": "Germany"
  },
  "MCC": {
    "code": "MCC",
    "company": "MAXON C&T CORPORATION LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "MCD": {
    "code": "MCD",
    "company": "4D BUILDING, INC.",
    "city": "Milford, MI 48381",
    "country": "United States"
  },
  "MCE": {
    "code": "MCE",
    "company": "MERCURY CONTAINER TRADING GMBH",
    "city": "Hamburg",
    "country": "Germany"
  },
  "MCF": {
    "code": "MCF",
    "company": "STATE ENTERPRISE RAILWAYS OF MOLDOVA",
    "city": "KISHINEV",
    "country": "Moldova, Republic of"
  },
  "MCG": {
    "code": "MCG",
    "company": "MITSUI CHEMICALS, INC.",
    "city": "Tokyo",
    "country": "Japan"
  },
  "MCH": {
    "code": "MCH",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MCI": {
    "code": "MCI",
    "company": "MAERSK CONTAINER INDUSTRY AS",
    "city": "TINGLEV",
    "country": "Denmark"
  },
  "MCK": {
    "code": "MCK",
    "company": "MOUNTAIN CONTAINER TRADING INC",
    "city": "Clinton",
    "country": "United States"
  },
  "MCL": {
    "code": "MCL",
    "company": "BORCHARD LINES LTD",
    "city": "LONDON EC1Y 4XY",
    "country": "United Kingdom"
  },
  "MCM": {
    "code": "MCM",
    "company": "MANAGEMENT CONTROL & MAINTENANCE S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "MCN": {
    "code": "MCN",
    "company": "MAROC CONTENEURS INTERNATIONAL",
    "city": "CASABLANCA",
    "country": "Morocco"
  },
  "MCO": {
    "code": "MCO",
    "company": "4SEAS SHIPPING AND FORWARDING BV",
    "city": "BERKEL EN RODENRIJD",
    "country": "Netherlands"
  },
  "MCP": {
    "code": "MCP",
    "company": "SEALAND MAERSK ASIA PTE LTD",
    "city": "SOUTHPOINT",
    "country": "Singapore"
  },
  "MCR": {
    "code": "MCR",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MCT": {
    "code": "MCT",
    "company": "MARIANA EXPRESS LINES PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "MCU": {
    "code": "MCU",
    "company": "MCR MOBILE CONTAINER REPAIR AB",
    "city": "Gothenburg",
    "country": "Sweden"
  },
  "MCZ": {
    "code": "MCZ",
    "company": "MINISTRY OF DEFENCE OF THE CZECH REPUBLIC",
    "city": "PRAGUE",
    "country": "Czech Republic"
  },
  "MDC": {
    "code": "MDC",
    "company": "MEDLOG (SHANGHAI) CO.,LTD.",
    "city": "Shanghai",
    "country": "China"
  },
  "MDE": {
    "code": "MDE",
    "company": "MINISTERE DE LA DEFENSE (CSOA/SIMMT)",
    "city": "VILLACOUBLAY",
    "country": "France"
  },
  "MDI": {
    "code": "MDI",
    "company": "M.D. SRL",
    "city": "Napoli",
    "country": "Italy"
  },
  "MDK": {
    "code": "MDK",
    "company": "RHENUS LOGISTICS AUSTRIA GMBH",
    "city": "KREMS",
    "country": "Austria"
  },
  "MDL": {
    "code": "MDL",
    "company": "MITSUBISHI CORPORATION",
    "city": "TOKYO",
    "country": "Japan"
  },
  "MDO": {
    "code": "MDO",
    "company": "QINGDAO MMD OFFSHORE  CO LTD",
    "city": "QINGDAO CITY,",
    "country": "China"
  },
  "MDP": {
    "code": "MDP",
    "company": "MODERN SHIPPING LINES CO LTD",
    "city": "Hong Kong",
    "country": "HK"
  },
  "MDS": {
    "code": "MDS",
    "company": "QINETIQ TARGET SYSTEMS LIMITED",
    "city": "ASHFORD, KENT",
    "country": "United Kingdom"
  },
  "MDU": {
    "code": "MDU",
    "company": "MDU MAIN-DONAU-UMSCHLAGS-TRANSPORT-GMBH",
    "city": "Nürnberg",
    "country": "Germany"
  },
  "MEB": {
    "code": "MEB",
    "company": "MEEBERG CONTAINER SERVICE BV",
    "city": "MOERDIJK",
    "country": "Netherlands"
  },
  "MEC": {
    "code": "MEC",
    "company": "MELFI MARINE CORP. S.A",
    "city": "CIUDAD DE LA HABANA",
    "country": "Cuba"
  },
  "MED": {
    "code": "MED",
    "company": "MSC- MEDITERRANEAN SHIPPING COMPANY S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "MEE": {
    "code": "MEE",
    "company": "MEPAVEX LOGISTICS BV",
    "city": "Bergen op Zoom",
    "country": "Netherlands"
  },
  "MEG": {
    "code": "MEG",
    "company": "SHANGHAI MEG IMPORT AND EXPORT CORPORATION",
    "city": "SHANGHAI",
    "country": "China"
  },
  "MEL": {
    "code": "MEL",
    "company": "JORN BOLDING A/S",
    "city": "Esbjerg V",
    "country": "Denmark"
  },
  "MEM": {
    "code": "MEM",
    "company": "MELKWEG / FRITOM BV",
    "city": "BOLSWARD",
    "country": "Netherlands"
  },
  "MEN": {
    "code": "MEN",
    "company": "MENADIONA, S.L.",
    "city": "PALAFOLLS",
    "country": "Spain"
  },
  "MER": {
    "code": "MER",
    "company": "MARIANA EXPRESS LINES PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "MES": {
    "code": "MES",
    "company": "NEST LTD INTERNATIONAL TRANSPORTS FLEXITANK SERVICES",
    "city": "GLYFADA",
    "country": "Greece"
  },
  "MET": {
    "code": "MET",
    "company": "METAMORFOZA CZ s.r.o.",
    "city": "Senov",
    "country": "Czech Republic"
  },
  "MEX": {
    "code": "MEX",
    "company": "MARIANA EXPRESS LINES PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "MFA": {
    "code": "MFA",
    "company": "MAINFREIGHT (AUSTRALIA)",
    "city": "Prestons",
    "country": "Australia"
  },
  "MFC": {
    "code": "MFC",
    "company": "FINSTERWALDER CONTAINER GMBH",
    "city": "KAUFBEUREN",
    "country": "Germany"
  },
  "MFR": {
    "code": "MFR",
    "company": "MARFRET S.A.",
    "city": "MARSEILLE",
    "country": "France"
  },
  "MFT": {
    "code": "MFT",
    "company": "MARFRET S.A.",
    "city": "MARSEILLE",
    "country": "France"
  },
  "MGE": {
    "code": "MGE",
    "company": "MONTER GLOBAL LOGISTICS (S) PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "MGG": {
    "code": "MGG",
    "company": "MESSER GROUP GMBH",
    "city": "BAD SODEN",
    "country": "Germany"
  },
  "MGI": {
    "code": "MGI",
    "company": "MEGA-INLINER SYSTEMS B.V.",
    "city": "Valkenswaard",
    "country": "Netherlands"
  },
  "MGL": {
    "code": "MGL",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "MGM": {
    "code": "MGM",
    "company": "MEGA CONTAINERS SALES, INC",
    "city": "WALNUT, CA 91789",
    "country": "United States"
  },
  "MGN": {
    "code": "MGN",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "MGS": {
    "code": "MGS",
    "company": "MGS OFFSHORE SDN BHD",
    "city": "Kuala Lumpur",
    "country": "Malaysia"
  },
  "MGT": {
    "code": "MGT",
    "company": "MAGITU GMBH",
    "city": "Berlin",
    "country": "Germany"
  },
  "MHC": {
    "code": "MHC",
    "company": "MAG CONTAINER LINES INC",
    "city": "MAHE",
    "country": "Seychelles"
  },
  "MHF": {
    "code": "MHF",
    "company": "MHF SERVICES",
    "city": "WEXFORD, PA 150190-9289",
    "country": "United States"
  },
  "MHH": {
    "code": "MHH",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MHI": {
    "code": "MHI",
    "company": "MITSUBISHI HEAVY INDUSTRIES, LTD.",
    "city": "Nagoya",
    "country": "Japan"
  },
  "MHV": {
    "code": "MHV",
    "company": "MAINPORT CONTAINER SERVICES B.V.",
    "city": "Rotterdam/Pernis",
    "country": "Netherlands"
  },
  "MIA": {
    "code": "MIA",
    "company": "MONTEBELLO AG",
    "city": "PONTRESINA",
    "country": "Switzerland"
  },
  "MIC": {
    "code": "MIC",
    "company": "PAN PACIFIC INTL TRADING GROUP (HK) LTD",
    "city": "HONG KONG",
    "country": "HK"
  },
  "MIE": {
    "code": "MIE",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MII": {
    "code": "MII",
    "company": "MIIL OÜ",
    "city": "Lagedi",
    "country": "Estonia"
  },
  "MIM": {
    "code": "MIM",
    "company": "MIMU BVBA",
    "city": "ANTWERP",
    "country": "Belgium"
  },
  "MIN": {
    "code": "MIN",
    "company": "SPEDITION MINOR GMBH",
    "city": "GELSENKIRCHEN",
    "country": "Germany"
  },
  "MIO": {
    "code": "MIO",
    "company": "MIOULIO SHIPPING COMPANY LIMITED",
    "city": "Limassol",
    "country": "Cyprus"
  },
  "MIP": {
    "code": "MIP",
    "company": "PT. INDONESIA MOROWALI INDUSTRIAL PARK",
    "city": "Jakarta Selatan",
    "country": "Indonesia"
  },
  "MIR": {
    "code": "MIR",
    "company": "MIRROR PALACE BVBA",
    "city": "Rijkevorsel",
    "country": "Belgium"
  },
  "MIS": {
    "code": "MIS",
    "company": "MIOULIO SHIPPING COMPANY LIMITED",
    "city": "Limassol",
    "country": "Cyprus"
  },
  "MIT": {
    "code": "MIT",
    "company": "PS SURVEY & CLAIM SERVICES APS",
    "city": "RISSKOV",
    "country": "Denmark"
  },
  "MIX": {
    "code": "MIX",
    "company": "LOGISTICA SUARDIAZ SL",
    "city": "MADRID",
    "country": "Spain"
  },
  "MJT": {
    "code": "MJT",
    "company": "ALE HEAVYLIFT B.V.",
    "city": "BREDA",
    "country": "Netherlands"
  },
  "MKA": {
    "code": "MKA",
    "company": "RWE NUCLEAR GMBH",
    "city": "MULHEIM-KAERLICH",
    "country": "Germany"
  },
  "MKL": {
    "code": "MKL",
    "company": "MEDKON LINES",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "MKT": {
    "code": "MKT",
    "company": "MK REFTRANS CO. LTD.",
    "city": "Nakhodka",
    "country": "Russian Federation"
  },
  "MKV": {
    "code": "MKV",
    "company": "TANK-YARD LLC",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "MKY": {
    "code": "MKY",
    "company": "MACKAY CONTAINERS PTY LTD",
    "city": "NORTH MACKAY",
    "country": "Australia"
  },
  "MLA": {
    "code": "MLA",
    "company": "MANUFACTURING MANAGEMENT, MALTA LTD",
    "city": "Kalkara",
    "country": "Malta"
  },
  "MLC": {
    "code": "MLC",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "MLG": {
    "code": "MLG",
    "company": "MOLGAS ENERGIA, SAU",
    "city": "SAN FERNANDO DE HENARES",
    "country": "Spain"
  },
  "MLI": {
    "code": "MLI",
    "company": "MITSUBISHI CHEMICAL LOGISTICS CORPORATION",
    "city": "MINATO-KU, TOKYO",
    "country": "Japan"
  },
  "MLK": {
    "code": "MLK",
    "company": "MATLACK LEASING LLC",
    "city": "BALA CYNWYD, PA 19004",
    "country": "United States"
  },
  "MLL": {
    "code": "MLL",
    "company": "METRO LOGISTICS INTERNATIONAL (PVT) LTD",
    "city": "Kelaniya",
    "country": "Sri Lanka"
  },
  "MLM": {
    "code": "MLM",
    "company": "MOOSE TANK PARTS",
    "city": "Westmaas",
    "country": "Netherlands"
  },
  "MLR": {
    "code": "MLR",
    "company": "LLC «REPETEK»",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "MLS": {
    "code": "MLS",
    "company": "MILLER INTERMODAL LOGISTICS SERVICES INC",
    "city": "RIDGELAND, MS-39152",
    "country": "United States"
  },
  "MLT": {
    "code": "MLT",
    "company": "MULCON SA",
    "city": "MONTEVIDEO",
    "country": "Uruguay"
  },
  "MLU": {
    "code": "MLU",
    "company": "MILITARY & LOGISTIC CONTAINER SHELTERS",
    "city": "Chicago",
    "country": "United States"
  },
  "MMA": {
    "code": "MMA",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MMC": {
    "code": "MMC",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "MMJ": {
    "code": "MMJ",
    "company": "MARITIME EQUIPMENT SUPPLY",
    "city": "MOKA",
    "country": "France"
  },
  "MML": {
    "code": "MML",
    "company": "CFL MULTIMODAL SA",
    "city": "DUDELANGE",
    "country": "Luxembourg"
  },
  "MMM": {
    "code": "MMM",
    "company": "MCCAUGHRIN MARITIME MARINE SYSTEMS INC.",
    "city": "WAYNE, MI 48184",
    "country": "United States"
  },
  "MMN": {
    "code": "MMN",
    "company": "MOBILE MINI UK LIMITED",
    "city": "STOCKTON ON TEES TS18 3TX",
    "country": "United Kingdom"
  },
  "MMR": {
    "code": "MMR",
    "company": "MAINSTREAM ENGINEERING CORPORATION",
    "city": "ROCKLEDGE, FL 32955",
    "country": "United States"
  },
  "MMX": {
    "code": "MMX",
    "company": "MARINOR ASSOCIATES",
    "city": "Houston, TX 77019",
    "country": "United States"
  },
  "MNB": {
    "code": "MNB",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MNO": {
    "code": "MNO",
    "company": "MAINCO",
    "city": "Tourlaville",
    "country": "France"
  },
  "MNT": {
    "code": "MNT",
    "company": "TRANSPORT MONTSANT S.A",
    "city": "BARCELONA",
    "country": "Spain"
  },
  "MOA": {
    "code": "MOA",
    "company": "MITSUI O.S.K. LINES LTD",
    "city": "MINATO-KU, TOKYO",
    "country": "Japan"
  },
  "MOB": {
    "code": "MOB",
    "company": "MOVE INTERMODAL NV",
    "city": "Genk",
    "country": "Belgium"
  },
  "MOC": {
    "code": "MOC",
    "company": "MACS MARITIME CARRIER SHIPPING GMBH & CO",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "MOD": {
    "code": "MOD",
    "company": "MODUL LTD",
    "city": "ST PETERSBURG",
    "country": "Russian Federation"
  },
  "MOE": {
    "code": "MOE",
    "company": "MITSUI O.S.K. LINES LTD",
    "city": "MINATO-KU, TOKYO",
    "country": "Japan"
  },
  "MOF": {
    "code": "MOF",
    "company": "MITSUI O.S.K. LINES LTD",
    "city": "MINATO-KU, TOKYO",
    "country": "Japan"
  },
  "MOG": {
    "code": "MOG",
    "company": "MITSUI O.S.K. LINES LTD",
    "city": "MINATO-KU, TOKYO",
    "country": "Japan"
  },
  "MOJ": {
    "code": "MOJ",
    "company": "MONTEIRO CONSTRUCOES E SERVICOS LTDA",
    "city": "ITAJAI SANTA CATARINA",
    "country": "Brazil"
  },
  "MOL": {
    "code": "MOL",
    "company": "MITSUI O.S.K. LINES LTD",
    "city": "MINATO-KU, TOKYO",
    "country": "Japan"
  },
  "MOM": {
    "code": "MOM",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "MON": {
    "code": "MON",
    "company": "MONSANTO COMPANY",
    "city": "SAINT LOUIS, MO-63167",
    "country": "United States"
  },
  "MOO": {
    "code": "MOO",
    "company": "MODEX ENERGY RENTALS SINGAPORE PTE. LTD.",
    "city": "Singapore",
    "country": "Singapore"
  },
  "MOR": {
    "code": "MOR",
    "company": "MITSUI O.S.K. LINES LTD",
    "city": "MINATO-KU, TOKYO",
    "country": "Japan"
  },
  "MOS": {
    "code": "MOS",
    "company": "MITSUI O.S.K. LINES LTD",
    "city": "MINATO-KU, TOKYO",
    "country": "Japan"
  },
  "MOT": {
    "code": "MOT",
    "company": "MITSUI O.S.K. LINES LTD",
    "city": "MINATO-KU, TOKYO",
    "country": "Japan"
  },
  "MPA": {
    "code": "MPA",
    "company": "MITSUBISHI POLYCRYSTALLINE SILICON AM CO",
    "city": "Theodore, AL 36582",
    "country": "United States"
  },
  "MPB": {
    "code": "MPB",
    "company": "LLC REFPEREVOZKY",
    "city": "SAINT PETERSBURG",
    "country": "Russian Federation"
  },
  "MPK": {
    "code": "MPK",
    "company": "POLYGON MPK, LLC",
    "city": "Mogilev",
    "country": "Belarus"
  },
  "MPM": {
    "code": "MPM",
    "company": "MOMENTIVE PERFORMANCE MATERIALS GMBH",
    "city": "LEVERKUSEN",
    "country": "Germany"
  },
  "MPR": {
    "code": "MPR",
    "company": "MONTEVIDEO PORT SERVICES S.A",
    "city": "MONTEVIDEO",
    "country": "Uruguay"
  },
  "MPS": {
    "code": "MPS",
    "company": "MODULAR PONTOON SYSTEMS BV",
    "city": "AALST",
    "country": "Netherlands"
  },
  "MRB": {
    "code": "MRB",
    "company": "GERMAN AEROSPACE CENTER",
    "city": "WESSLING",
    "country": "Germany"
  },
  "MRC": {
    "code": "MRC",
    "company": "LLC \"EMRCC\"",
    "city": "Saint-Petersburg",
    "country": "Russian Federation"
  },
  "MRG": {
    "code": "MRG",
    "company": "MARGUISA SHIPPING LINES, S.L.U.",
    "city": "MADRID",
    "country": "Spain"
  },
  "MRK": {
    "code": "MRK",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MRN": {
    "code": "MRN",
    "company": "MARINETEC LTD",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "MRS": {
    "code": "MRS",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MRT": {
    "code": "MRT",
    "company": "PT MERATUS LINE",
    "city": "SURABAYA",
    "country": "Indonesia"
  },
  "MRV": {
    "code": "MRV",
    "company": "MERVIELDE TRANSPORT NV",
    "city": "ERTVELDE-RIEME",
    "country": "Belgium"
  },
  "MRZ": {
    "code": "MRZ",
    "company": "MARENZANA SPA",
    "city": "NOVI LIGURE (AL)",
    "country": "Italy"
  },
  "MSA": {
    "code": "MSA",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MSC": {
    "code": "MSC",
    "company": "MSC- MEDITERRANEAN SHIPPING COMPANY S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "MSD": {
    "code": "MSD",
    "company": "MSC- MEDITERRANEAN SHIPPING COMPANY S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "MSF": {
    "code": "MSF",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MSH": {
    "code": "MSH",
    "company": "MALAYSIAN SHIPPING CORP.SENDIRIAN BERHAD",
    "city": "",
    "country": "Malaysia"
  },
  "MSK": {
    "code": "MSK",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MSM": {
    "code": "MSM",
    "company": "MSC- MEDITERRANEAN SHIPPING COMPANY S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "MSN": {
    "code": "MSN",
    "company": "MSC- MEDITERRANEAN SHIPPING COMPANY S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "MSO": {
    "code": "MSO",
    "company": "M/S CONTAINERS A/S",
    "city": "RISSKOV",
    "country": "Denmark"
  },
  "MSP": {
    "code": "MSP",
    "company": "MSC- MEDITERRANEAN SHIPPING COMPANY S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "MSR": {
    "code": "MSR",
    "company": "MICROSTAR LOGISTICS",
    "city": "DENVER, CO 80202",
    "country": "United States"
  },
  "MSS": {
    "code": "MSS",
    "company": "MSSA",
    "city": "SAINT MARCEL",
    "country": "France"
  },
  "MST": {
    "code": "MST",
    "company": "MSC- MEDITERRANEAN SHIPPING COMPANY S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "MSU": {
    "code": "MSU",
    "company": "M/S CONTAINERS A/S",
    "city": "RISSKOV",
    "country": "Denmark"
  },
  "MSV": {
    "code": "MSV",
    "company": "MSC- MEDITERRANEAN SHIPPING COMPANY S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "MSW": {
    "code": "MSW",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MSY": {
    "code": "MSY",
    "company": "MSC- MEDITERRANEAN SHIPPING COMPANY S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "MSZ": {
    "code": "MSZ",
    "company": "MSC- MEDITERRANEAN SHIPPING COMPANY S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "MTA": {
    "code": "MTA",
    "company": "FINEBRO 3 S.L.",
    "city": "MADRID",
    "country": "Spain"
  },
  "MTB": {
    "code": "MTB",
    "company": "MULTIBOXX LIMITED",
    "city": "KOWLOON",
    "country": "HK"
  },
  "MTC": {
    "code": "MTC",
    "company": "MERLION HOLDINGS PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "MTE": {
    "code": "MTE",
    "company": "BAYER AGRICULTURE BVBA",
    "city": "ANTWERPEN",
    "country": "Belgium"
  },
  "MTG": {
    "code": "MTG",
    "company": "MATHESON TRI-GAS INC",
    "city": "WARREN",
    "country": "United States"
  },
  "MTI": {
    "code": "MTI",
    "company": "MILLER INTERMODAL LOGISTICS SERVICES INC",
    "city": "RIDGELAND, MS-39152",
    "country": "United States"
  },
  "MTL": {
    "code": "MTL",
    "company": "DUTCH ANTILLEAN CONTAINER LEASING NV",
    "city": "LIMASSOL",
    "country": "Cyprus"
  },
  "MTM": {
    "code": "MTM",
    "company": "MTT SHIPPING SDN. BHD.",
    "city": "PORT KLANG",
    "country": "Malaysia"
  },
  "MTN": {
    "code": "MTN",
    "company": "MOUNTAIN HOLDING B.V.",
    "city": "Galder",
    "country": "Netherlands"
  },
  "MTO": {
    "code": "MTO",
    "company": "METANO  LTD",
    "city": "ROWLANDS GILL, TYNE & WEAR NE39 IEH",
    "country": "United Kingdom"
  },
  "MTP": {
    "code": "MTP",
    "company": "METRANS, a.s.",
    "city": "PRAGUE 10",
    "country": "Czech Republic"
  },
  "MTR": {
    "code": "MTR",
    "company": "TAL INTERNATIONAL",
    "city": "PURCHASE, NY 10577-2135",
    "country": "United States"
  },
  "MTS": {
    "code": "MTS",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "MTT": {
    "code": "MTT",
    "company": "TATIANA TSIBIKOVA",
    "city": "St Petersburg",
    "country": "Russian Federation"
  },
  "MTU": {
    "code": "MTU",
    "company": "MTU ONSITE ENERGY CORPORATION",
    "city": "MANKATO, MN-56001",
    "country": "United States"
  },
  "MTV": {
    "code": "MTV",
    "company": "CONTAINERS MTV",
    "city": "GENOVA (GE)",
    "country": "Italy"
  },
  "MTY": {
    "code": "MTY",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "MUC": {
    "code": "MUC",
    "company": "G.T.S SPA",
    "city": "BARI",
    "country": "Italy"
  },
  "MUK": {
    "code": "MUK",
    "company": "MUTO CO., LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "MUL": {
    "code": "MUL",
    "company": "«MUVLINK» LLC",
    "city": "Minsk",
    "country": "Belarus"
  },
  "MUM": {
    "code": "MUM",
    "company": "INTERRAIL EUROPE GMBH",
    "city": "FRANKFURT AM MAIN",
    "country": "Germany"
  },
  "MUO": {
    "code": "MUO",
    "company": "MARGUISA SHIPPING LINES, S.L.U.",
    "city": "MADRID",
    "country": "Spain"
  },
  "MUR": {
    "code": "MUR",
    "company": "FRAMATOME SAS",
    "city": "CHALON SUR SAONE",
    "country": "France"
  },
  "MUT": {
    "code": "MUT",
    "company": "MUTUALISTA ACOREANA TRANSP. MARITIMOS SA",
    "city": "LISBOA",
    "country": "Portugal"
  },
  "MVC": {
    "code": "MVC",
    "company": "OY MOONWAY AB",
    "city": "TURKU -",
    "country": "Finland"
  },
  "MVI": {
    "code": "MVI",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MVL": {
    "code": "MVL",
    "company": "MASTER VALLEY RESOURCES LTD",
    "city": "Saint-Petersburg",
    "country": "Russian Federation"
  },
  "MVW": {
    "code": "MVW",
    "company": "MORITA CHEMICAL INDUSTRIES CO, LTD",
    "city": "OSAKA",
    "country": "Japan"
  },
  "MWA": {
    "code": "MWA",
    "company": "MEDWAY ASSETS - GESTÃO DE ATIVOS S.A.",
    "city": "Lisboa",
    "country": "Portugal"
  },
  "MWC": {
    "code": "MWC",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MWI": {
    "code": "MWI",
    "company": "MASTERANK WAX INC",
    "city": "Pleasanton, CA 94588",
    "country": "United States"
  },
  "MWL": {
    "code": "MWL",
    "company": "MUWON LS LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "MWM": {
    "code": "MWM",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MWT": {
    "code": "MWT",
    "company": "SHANGHAI MILKYWAY CHEMICAL LOGISTICS CO, Ltd",
    "city": "SHANGHAI 201203",
    "country": "China"
  },
  "MWW": {
    "code": "MWW",
    "company": "MESSER LLC",
    "city": "BRIDGEWATER, NJ 08807",
    "country": "United States"
  },
  "MXC": {
    "code": "MXC",
    "company": "MAXICON CONTAINER LINE PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "MXF": {
    "code": "MXF",
    "company": "MEXICHEM FLUOR JAPAN LIMITED",
    "city": "Tokyo",
    "country": "Japan"
  },
  "MXM": {
    "code": "MXM",
    "company": "MODALIS",
    "city": "AIX-EN-PROVENCE",
    "country": "France"
  },
  "MXS": {
    "code": "MXS",
    "company": "MAXX SERVICES LIMITED",
    "city": "LIMASSOL",
    "country": "Cyprus"
  },
  "NAB": {
    "code": "NAB",
    "company": "ECST CONTAINER SERVICES & TRADING GMBH",
    "city": "SEEVETAL / BULLENHAUSEN",
    "country": "Germany"
  },
  "NAF": {
    "code": "NAF",
    "company": "NDMA",
    "city": "KJELLER",
    "country": "Norway"
  },
  "NAG": {
    "code": "NAG",
    "company": "TRANSPORT NAGELS N.V.",
    "city": "Antwerpen",
    "country": "Belgium"
  },
  "NAI": {
    "code": "NAI",
    "company": "SAGER ASSETS S.A.",
    "city": "HUNENBERG",
    "country": "Switzerland"
  },
  "NAK": {
    "code": "NAK",
    "company": "WEST OCEAN SHIPPING COMPANY LIMITED",
    "city": "CENTRAL",
    "country": "HK"
  },
  "NAR": {
    "code": "NAR",
    "company": "MT CONTAINER GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "NAT": {
    "code": "NAT",
    "company": "SHAPE",
    "city": "MONS",
    "country": "Belgium"
  },
  "NAX": {
    "code": "NAX",
    "company": "NIPPON AEROSIL LTD",
    "city": "3,MITA-CHO",
    "country": "Japan"
  },
  "NAZ": {
    "code": "NAZ",
    "company": "JSC NEVINNOMYSSKY AZOT",
    "city": "NEVINNOMYSSK, STAVROPOL TERRITORY",
    "country": "Russian Federation"
  },
  "NBC": {
    "code": "NBC",
    "company": "BUNNIK CREATIONS",
    "city": "Bleiswijk",
    "country": "Netherlands"
  },
  "NBG": {
    "code": "NBG",
    "company": "ROLANDE LNG B.V",
    "city": "JL GIESSEN",
    "country": "Netherlands"
  },
  "NBT": {
    "code": "NBT",
    "company": "NISSAN BUTSURYU CO LTD",
    "city": "TOKYO",
    "country": "Japan"
  },
  "NBY": {
    "code": "NBY",
    "company": "NINGBO OCEAN SHIPPING CO,LTD",
    "city": "NINGBO,ZHEJIANG PROV",
    "country": "China"
  },
  "NCC": {
    "code": "NCC",
    "company": "NIPPON CONCEPT CORPORATION",
    "city": "CHIYODA-KU - TOKYO",
    "country": "Japan"
  },
  "NCL": {
    "code": "NCL",
    "company": "CAREER (HK) CO., LIMITED",
    "city": "HONGKONG",
    "country": "HK"
  },
  "NCS": {
    "code": "NCS",
    "company": "DAHER NUCLEAR TECHNOLOGIES GMBH",
    "city": "HANAU",
    "country": "Germany"
  },
  "NCT": {
    "code": "NCT",
    "company": "ORCA CONTAINER ASSET MANAGEMENT",
    "city": "CAPE TOWN",
    "country": "South Africa"
  },
  "NCX": {
    "code": "NCX",
    "company": "NOR LINES INTERNATIONAL",
    "city": "STAVANGER",
    "country": "Norway"
  },
  "NDA": {
    "code": "NDA",
    "company": "INTERNATIONAL NUCLEAR SERVICES LTD",
    "city": "WARRINGTON WA36AS",
    "country": "United Kingdom"
  },
  "NDC": {
    "code": "NDC",
    "company": "NIDEC",
    "city": "ROCHE LA MOLIERE",
    "country": "France"
  },
  "NDS": {
    "code": "NDS",
    "company": "NILE DUTCH AFRICA LINE",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "NDX": {
    "code": "NDX",
    "company": "NORDEX ENERGY GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "NEC": {
    "code": "NEC",
    "company": "NAT SHIPPING BAGGING SERVICES LTD. & ASSOCIATES",
    "city": "HAROLD HILL RM3 8UF",
    "country": "United Kingdom"
  },
  "NEN": {
    "code": "NEN",
    "company": "BOLUDA LINES S.A",
    "city": "VALENCIA",
    "country": "Spain"
  },
  "NEO": {
    "code": "NEO",
    "company": "NEOMAT SRO",
    "city": "Galanta",
    "country": "Slovakia"
  },
  "NEP": {
    "code": "NEP",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "NES": {
    "code": "NES",
    "company": "BOREALIS AB",
    "city": "STENUNGSUND",
    "country": "Sweden"
  },
  "NET": {
    "code": "NET",
    "company": "SOLAR NETWORKS SP ZOO",
    "city": "Leszno",
    "country": "Poland"
  },
  "NEV": {
    "code": "NEV",
    "company": "CAI INTERNATIONAL",
    "city": "SAN FRANCISCO, CA 94105",
    "country": "United States"
  },
  "NEW": {
    "code": "NEW",
    "company": "CR CONTAINER TRADING GMBH",
    "city": "Hamburg",
    "country": "Germany"
  },
  "NFA": {
    "code": "NFA",
    "company": "MITSUBISHI GAS CHEMICAL COMPANY INC",
    "city": "TOKYO",
    "country": "Japan"
  },
  "NFC": {
    "code": "NFC",
    "company": "NATIONAL FACTORY FOR PROCESSING AND TREATING MINERALS",
    "city": "Dammam",
    "country": "Saudi Arabia"
  },
  "NFL": {
    "code": "NFL",
    "company": "DFDS A/S",
    "city": "COPENHAGEN",
    "country": "Denmark"
  },
  "NFP": {
    "code": "NFP",
    "company": "NAVIN FLUORINE INTERNATIONAL LTD",
    "city": "MUMBAI",
    "country": "India"
  },
  "NFR": {
    "code": "NFR",
    "company": "BOLUDA LINES S.A",
    "city": "VALENCIA",
    "country": "Spain"
  },
  "NFS": {
    "code": "NFS",
    "company": "KMG SINGAPORE PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "NFW": {
    "code": "NFW",
    "company": "BOLUDA LINES S.A",
    "city": "VALENCIA",
    "country": "Spain"
  },
  "NGC": {
    "code": "NGC",
    "company": "VALQUA NGC, INC.",
    "city": "HOUSTON, TX-77093",
    "country": "United States"
  },
  "NGE": {
    "code": "NGE",
    "company": "NIPPON GASES BELGIUM",
    "city": "Schoten",
    "country": "Belgium"
  },
  "NGT": {
    "code": "NGT",
    "company": "FTLL LIMITED",
    "city": "TORTOLA",
    "country": "Virgin Islands, British"
  },
  "NGZ": {
    "code": "NGZ",
    "company": "JOINT STOCK COMPANY \"AZOT\"",
    "city": "Novomoskovsk. Tula region",
    "country": "Russian Federation"
  },
  "NHC": {
    "code": "NHC",
    "company": "MITSUI & CO LTD",
    "city": "OSAKA",
    "country": "Japan"
  },
  "NHI": {
    "code": "NHI",
    "company": "NIPPON GASES NORGE AS",
    "city": "OSLO",
    "country": "Norway"
  },
  "NID": {
    "code": "NID",
    "company": "NILE DUTCH AFRICA LINE",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "NIK": {
    "code": "NIK",
    "company": "NORILSK FORWARDING COMPANY LTD",
    "city": "NORILSK",
    "country": "Russian Federation"
  },
  "NIM": {
    "code": "NIM",
    "company": "RRT RHEIN-RUHR TERMINAL GESELLSCHAFT FUR CONTAINER",
    "city": "DUISBURG",
    "country": "Germany"
  },
  "NIN": {
    "code": "NIN",
    "company": "REF TRANS 2013 LTD",
    "city": "TBILISI",
    "country": "Georgia"
  },
  "NIO": {
    "code": "NIO",
    "company": "ROYAL NIOZ",
    "city": "HORNTJE TEXEL",
    "country": "Netherlands"
  },
  "NIP": {
    "code": "NIP",
    "company": "NORTH INLAND PORT INTERNATIONAL LOGISTICS CO.,LTD",
    "city": "ULANQAB",
    "country": "China"
  },
  "NIQ": {
    "code": "NIQ",
    "company": "NUUK IMEQ",
    "city": "NUUK",
    "country": "Greenland"
  },
  "NIR": {
    "code": "NIR",
    "company": "NIRINT SHIPPING BV",
    "city": "Barendrecht",
    "country": "Netherlands"
  },
  "NIS": {
    "code": "NIS",
    "company": "TECHNOPYR SA",
    "city": "PIRAEUS",
    "country": "Greece"
  },
  "NJB": {
    "code": "NJB",
    "company": "CS EURASIA LEASING (SEA) PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "NJF": {
    "code": "NJF",
    "company": "NJ FROMENT AND COMPANY LIMITED",
    "city": "Stamford",
    "country": "United Kingdom"
  },
  "NKL": {
    "code": "NKL",
    "company": "NORTHERN LIGHTS SHIPPING",
    "city": "PIRAEUS",
    "country": "Greece"
  },
  "NKN": {
    "code": "NKN",
    "company": "PJSC NIZHNEKAMSKNEFTEKHIM",
    "city": "Tatarstan Republic",
    "country": "Russian Federation"
  },
  "NKT": {
    "code": "NKT",
    "company": "ALTUN LOJISTIK A.S.",
    "city": "MERSIN",
    "country": "Turkey"
  },
  "NLA": {
    "code": "NLA",
    "company": "NRS LOGISTICS INC.",
    "city": "White Plains, NY 10606",
    "country": "United States"
  },
  "NLB": {
    "code": "NLB",
    "company": "NIRMA LIMITED",
    "city": "BHAVNAGAR",
    "country": "India"
  },
  "NLC": {
    "code": "NLC",
    "company": "LVNF NUCLEAR & LOGISTIC CONSULTING",
    "city": "Saint-Paul-Trois-Châteaux",
    "country": "France"
  },
  "NLH": {
    "code": "NLH",
    "company": "NEPTUNE PACIFIC LINE PTE LTD",
    "city": "AUCKLAND AIRPORT",
    "country": "New Zealand"
  },
  "NLK": {
    "code": "NLK",
    "company": "KSH GMBH",
    "city": "Unterschleißheim",
    "country": "Germany"
  },
  "NLL": {
    "code": "NLL",
    "company": "NOBLE CONTAINER LEASING LIMITED",
    "city": "JORDAN, KOWLOON",
    "country": "HK"
  },
  "NLN": {
    "code": "NLN",
    "company": "NOR LINES INTERNATIONAL",
    "city": "STAVANGER",
    "country": "Norway"
  },
  "NLR": {
    "code": "NLR",
    "company": "LLC \"NEW LOGISTIC\"",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "NMA": {
    "code": "NMA",
    "company": "COMMANDER, NAVAL AIR SYSTEMS COMMAND AIR 6.7.6.2",
    "city": "PATUXENT RIVER, MD 20670",
    "country": "United States"
  },
  "NMC": {
    "code": "NMC",
    "company": "NEWCREST MINING LIMITED",
    "city": "Orange",
    "country": "Australia"
  },
  "NMF": {
    "code": "NMF",
    "company": "NATIONAL OCEANOGRAPHY CENTRE",
    "city": "SOUTHAMPTON",
    "country": "United Kingdom"
  },
  "NMH": {
    "code": "NMH",
    "company": "MASSY GAS PRODUCTS (TRINIDAD) LTD",
    "city": "POINT LISAS",
    "country": "Trinidad and Tobago"
  },
  "NMK": {
    "code": "NMK",
    "company": "PJSC MINING AND METALLURGICAL COMPANY NORILSK NICKEL",
    "city": "DUDINKA",
    "country": "Russian Federation"
  },
  "NMS": {
    "code": "NMS",
    "company": "TRANSNOMICS LTD",
    "city": "KHAN POR SEN CHEY, PHNOMPENH",
    "country": "Cambodia"
  },
  "NMT": {
    "code": "NMT",
    "company": "NMT KONTEYNER A.S",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "NNC": {
    "code": "NNC",
    "company": "CANADIAN ROYALTIES INC",
    "city": "MONTREAL ,QUEBEC",
    "country": "Canada"
  },
  "NOA": {
    "code": "NOA",
    "company": "NORWEGIAN DEFENCE MATERIAL AGENCY (NDMA)",
    "city": "Bergen",
    "country": "Norway"
  },
  "NOB": {
    "code": "NOB",
    "company": "NORDIC BULKERS AB",
    "city": "GOTHENBURG",
    "country": "Sweden"
  },
  "NOC": {
    "code": "NOC",
    "company": "NORA SPA",
    "city": "S STEFANO MAGRA (SP)",
    "country": "Italy"
  },
  "NOF": {
    "code": "NOF",
    "company": "NORTH OIL AND FATS GMBH",
    "city": "Hamburg",
    "country": "Germany"
  },
  "NOK": {
    "code": "NOK",
    "company": "NAUKA LINES PTE LTD",
    "city": "New Delhi",
    "country": "India"
  },
  "NOL": {
    "code": "NOL",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "NON": {
    "code": "NON",
    "company": "TSL COMPANIES, INC.",
    "city": "Omaha",
    "country": "United States"
  },
  "NOR": {
    "code": "NOR",
    "company": "TANK MANAGEMENT A/S",
    "city": "OSLO",
    "country": "Norway"
  },
  "NOS": {
    "code": "NOS",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "NOW": {
    "code": "NOW",
    "company": "NORWOOD SHIPPING SA",
    "city": "QUART DE POBLET- VALENCIA",
    "country": "Spain"
  },
  "NOX": {
    "code": "NOX",
    "company": "LEEDEN NATIONAL OXYGEN LTD",
    "city": "JURONG",
    "country": "Singapore"
  },
  "NPC": {
    "code": "NPC",
    "company": "NEWPORT TANK CONTAINERS INC",
    "city": "CLEVELAND, OH 44116",
    "country": "United States"
  },
  "NPF": {
    "code": "NPF",
    "company": "BOLUDA LINES S.A",
    "city": "VALENCIA",
    "country": "Spain"
  },
  "NPJ": {
    "code": "NPJ",
    "company": "BOLUDA LINES S.A",
    "city": "VALENCIA",
    "country": "Spain"
  },
  "NPL": {
    "code": "NPL",
    "company": "SAIGON NEWPORT ONE MEMBER LTD LIABILITY CORP",
    "city": "HO CHI MINH CITY",
    "country": "Viet Nam"
  },
  "NPO": {
    "code": "NPO",
    "company": "OOO NPO REASIB",
    "city": "Tomsk",
    "country": "Russian Federation"
  },
  "NPP": {
    "code": "NPP",
    "company": "NPP SPETSAVIA LLC",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "NPR": {
    "code": "NPR",
    "company": "SEA STAR LINE LLC",
    "city": "JACKSONVILLE, FL 32256",
    "country": "United States"
  },
  "NPT": {
    "code": "NPT",
    "company": "NEPTUNE GLOBAL LOGISTICS PTE LTD",
    "city": "",
    "country": "Singapore"
  },
  "NPW": {
    "code": "NPW",
    "company": "BOLUDA LINES S.A",
    "city": "VALENCIA",
    "country": "Spain"
  },
  "NPZ": {
    "code": "NPZ",
    "company": "JSC MOZYR OR",
    "city": "Mozyr-11",
    "country": "Belarus"
  },
  "NRA": {
    "code": "NRA",
    "company": "AEROGAS PROCESSORS LIMITED",
    "city": "COUVA",
    "country": "Trinidad and Tobago"
  },
  "NRC": {
    "code": "NRC",
    "company": "SHANGHAI NRS-NORINCO LOGISTICS CO.LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "NRI": {
    "code": "NRI",
    "company": "NATIONAL REFRIGERANTS INC.",
    "city": "ROSENHAYN, NJ 08352",
    "country": "United States"
  },
  "NRK": {
    "code": "NRK",
    "company": "NR COOLING SERVICES B.V.",
    "city": "Numansdorp",
    "country": "Netherlands"
  },
  "NRS": {
    "code": "NRS",
    "company": "NRS CORPORATION",
    "city": "TOKYO",
    "country": "Japan"
  },
  "NSA": {
    "code": "NSA",
    "company": "THE NATIONAL SHIPPING CO OF SAUDI ARABIA",
    "city": "JEDDAH",
    "country": "Saudi Arabia"
  },
  "NSC": {
    "code": "NSC",
    "company": "NORTHERN SHIPPING COMPANY",
    "city": "ARKHANGELSK",
    "country": "Russian Federation"
  },
  "NSE": {
    "code": "NSE",
    "company": "NORTH SEA EXPRESS",
    "city": "ZEEBRUGGE",
    "country": "Belgium"
  },
  "NSG": {
    "code": "NSG",
    "company": "NBC SPORTS",
    "city": "STAMFORD, CT-06902",
    "country": "United States"
  },
  "NSM": {
    "code": "NSM",
    "company": "GRU COMEDIL SRL",
    "city": "FONTANAFREDDA",
    "country": "Italy"
  },
  "NSO": {
    "code": "NSO",
    "company": "NORDIC SHELTER AS",
    "city": "GAMLE FREDRIKSTAD",
    "country": "Norway"
  },
  "NSR": {
    "code": "NSR",
    "company": "NAM SUNG SHIPPING CO LTD",
    "city": "SEOUL 100-760",
    "country": "Korea, Republic of"
  },
  "NSS": {
    "code": "NSS",
    "company": "NAM SUNG SHIPPING CO LTD",
    "city": "SEOUL 100-760",
    "country": "Korea, Republic of"
  },
  "NST": {
    "code": "NST",
    "company": "NEPAL SHIPPING AND MULTI MODEL TRANSPORT",
    "city": "KATHMANDU",
    "country": "Nepal"
  },
  "NSW": {
    "code": "NSW",
    "company": "NORDDEUTSCHE SEEKABELWERKE GMBH",
    "city": "Nordenham",
    "country": "Germany"
  },
  "NTA": {
    "code": "NTA",
    "company": "NT RENTAL APS",
    "city": "AALBORG",
    "country": "Denmark"
  },
  "NTO": {
    "code": "NTO",
    "company": "NARVATANK LTD",
    "city": "NARVA",
    "country": "Estonia"
  },
  "NTT": {
    "code": "NTT",
    "company": "NANTONG TANK CONTAINER CO, LTD",
    "city": "JIANGSU",
    "country": "China"
  },
  "NUS": {
    "code": "NUS",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "NVG": {
    "code": "NVG",
    "company": "NOVA GAS TECHNOLOGIES,INC",
    "city": "NORTH CHARLESTON,SC 29418",
    "country": "United States"
  },
  "NVO": {
    "code": "NVO",
    "company": "NAVIO SHIPPING LLC",
    "city": "Dubai",
    "country": "United Arab Emirates"
  },
  "NWB": {
    "code": "NWB",
    "company": "NIJHOF WASSINK B.V.",
    "city": "RIJSSEN",
    "country": "Netherlands"
  },
  "NWN": {
    "code": "NWN",
    "company": "NWM-EUROPE NV",
    "city": "EVERGEM",
    "country": "Belgium"
  },
  "NXS": {
    "code": "NXS",
    "company": "NEXTER SYSTEMS",
    "city": "VERSAILLES",
    "country": "France"
  },
  "NXT": {
    "code": "NXT",
    "company": "NEXTER MUNITIONS",
    "city": "VERSAILLES CEDEX",
    "country": "France"
  },
  "NYC": {
    "code": "NYC",
    "company": "NIYAC CORPORATION",
    "city": "TOKYO 135-0041",
    "country": "Japan"
  },
  "NYK": {
    "code": "NYK",
    "company": "NYK LINE",
    "city": "TOKYO",
    "country": "Japan"
  },
  "NZD": {
    "code": "NZD",
    "company": "NOVOZYMES A/S",
    "city": "BAGSVERD",
    "country": "Denmark"
  },
  "NZK": {
    "code": "NZK",
    "company": "NOVOZYMES A/S",
    "city": "BAGSVERD",
    "country": "Denmark"
  },
  "NZL": {
    "code": "NZL",
    "company": "CONTAINER SALES & LEASING LTD",
    "city": "AUCKLAND",
    "country": "New Zealand"
  },
  "OAC": {
    "code": "OAC",
    "company": "LINDE-HADJIKYRIAKOS GAS LTD.",
    "city": "STROVOLOS",
    "country": "Cyprus"
  },
  "OAK": {
    "code": "OAK",
    "company": "OAK RIDGE NATIONAL LABORATORY",
    "city": "Oak Ridge, TN-37831",
    "country": "United States"
  },
  "OAS": {
    "code": "OAS",
    "company": "OCEAN AXIS SHIPPING SERVICES LLC",
    "city": "Bur-DUBAI",
    "country": "United Arab Emirates"
  },
  "OAT": {
    "code": "OAT",
    "company": "OPAL ASIA INDIA PVT LTD",
    "city": "Mumbai",
    "country": "India"
  },
  "OBC": {
    "code": "OBC",
    "company": "N.V. COBO CONSTRUCTION",
    "city": "PARAMARIBO",
    "country": "Suriname"
  },
  "OBE": {
    "code": "OBE",
    "company": "FERROCARRIL ANTOFAGASTA A BOLIVIA",
    "city": "Antofagasta",
    "country": "Chile"
  },
  "OBT": {
    "code": "OBT",
    "company": "JIANGSU O-BEST NEW MATERIALS CO LTD",
    "city": "NANTONG, JIANGSU",
    "country": "China"
  },
  "OCA": {
    "code": "OCA",
    "company": "OCEANBOX CONTAINERS LTD",
    "city": "Limassol",
    "country": "Cyprus"
  },
  "OCB": {
    "code": "OCB",
    "company": "OCEAN BLUE FINANCE GROUP LIMITED",
    "city": "WANHAI",
    "country": "HK"
  },
  "OCC": {
    "code": "OCC",
    "company": "ORANGE CONTAINER LINE B.V.",
    "city": "RHOON",
    "country": "Netherlands"
  },
  "OCE": {
    "code": "OCE",
    "company": "OCEAN EXPRESS COMPANY",
    "city": "ALEXANDRIA",
    "country": "Egypt"
  },
  "OCG": {
    "code": "OCG",
    "company": "OCEAN CONTAINER LEASING SERVICES LIMITED",
    "city": "WANCHAI",
    "country": "HK"
  },
  "OCH": {
    "code": "OCH",
    "company": "OCEAN CHEMICAL DMCC",
    "city": "Dubai",
    "country": "United Arab Emirates"
  },
  "OCI": {
    "code": "OCI",
    "company": "ARKEMA INC.",
    "city": "KING OF PRUSSIA, PA 19406",
    "country": "United States"
  },
  "OCL": {
    "code": "OCL",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "OCO": {
    "code": "OCO",
    "company": "AONOCO LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "OCV": {
    "code": "OCV",
    "company": "O.C.C. OVERBEEK CONTAINER CONTROL BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "ODD": {
    "code": "ODD",
    "company": "WELFIT ODDY (PTY) LIMITED",
    "city": "PORT ELIZABETH",
    "country": "South Africa"
  },
  "OEG": {
    "code": "OEG",
    "company": "OEG OFFSHORE PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "OER": {
    "code": "OER",
    "company": "ORIENTAL EQUIPMENT SERVICES INC",
    "city": "Secaucus, NJ 07094",
    "country": "United States"
  },
  "OES": {
    "code": "OES",
    "company": "OSTU-STETTIN",
    "city": "Leoben",
    "country": "Austria"
  },
  "OEX": {
    "code": "OEX",
    "company": "MARGUISA SHIPPING LINES, S.L.U.",
    "city": "MADRID",
    "country": "Spain"
  },
  "OFF": {
    "code": "OFF",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "OFS": {
    "code": "OFS",
    "company": "OILFIELD HIRE AND SERVICES",
    "city": "Coleshill, Birmingham, B46 1JY",
    "country": "United Kingdom"
  },
  "OGC": {
    "code": "OGC",
    "company": "SPEEDGAS PTY LTD",
    "city": "BROOKVALE NSW 2100",
    "country": "Australia"
  },
  "OGP": {
    "code": "OGP",
    "company": "ROZA MIRA LLC",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "OGT": {
    "code": "OGT",
    "company": "ORPHAN GRAIN TRAIN INC.",
    "city": "NORFOLK, NE-68702",
    "country": "United States"
  },
  "OKW": {
    "code": "OKW",
    "company": "NDUNESE SEALAND SDN BHD",
    "city": "Johor Bahru",
    "country": "Malaysia"
  },
  "OLC": {
    "code": "OLC",
    "company": "OILCORE UAB",
    "city": "Mazeikiai",
    "country": "Lithuania"
  },
  "OLT": {
    "code": "OLT",
    "company": "ODYSSEY FOODTRANS",
    "city": "Irvine, CA 92614",
    "country": "United States"
  },
  "OMA": {
    "code": "OMA",
    "company": "OMNI TANKER PTY LTD",
    "city": "SMEATON GRANGE",
    "country": "Australia"
  },
  "OMC": {
    "code": "OMC",
    "company": "OMINA GROUP (PTY) LTD",
    "city": "Bryanston",
    "country": "South Africa"
  },
  "OMN": {
    "code": "OMN",
    "company": "AIR LIQUID ITALIA SERVICE SRL",
    "city": "MILANO MI",
    "country": "Italy"
  },
  "OMS": {
    "code": "OMS",
    "company": "ORICA MINING SERVICES PORTUGAL",
    "city": "Aljustrel",
    "country": "Portugal"
  },
  "ONE": {
    "code": "ONE",
    "company": "OCEAN NETWORK EXPRESS PTE. LTD.",
    "city": "Singapore",
    "country": "Singapore"
  },
  "ONL": {
    "code": "ONL",
    "company": "OMAN SHIPPING COMPANY S.A.O.C",
    "city": "MUSCAT",
    "country": "Oman"
  },
  "OOC": {
    "code": "OOC",
    "company": "ORIENT OVERSEAS CONTAINER LINE LTD.",
    "city": "WANCHAI",
    "country": "HK"
  },
  "OOL": {
    "code": "OOL",
    "company": "ORIENT OVERSEAS CONTAINER LINE LTD.",
    "city": "WANCHAI",
    "country": "HK"
  },
  "OPD": {
    "code": "OPD",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "OPL": {
    "code": "OPL",
    "company": "OPAL ASIA INDIA PVT LTD",
    "city": "Mumbai",
    "country": "India"
  },
  "OPO": {
    "code": "OPO",
    "company": "COMMONWEALTH OF AUSTRALIA",
    "city": "Rabat-Souissi",
    "country": "Morocco"
  },
  "OPT": {
    "code": "OPT",
    "company": "OPTIMODAL INC",
    "city": "WEST CHESTER, PA 19381",
    "country": "United States"
  },
  "ORA": {
    "code": "ORA",
    "company": "INTERMODAL LOGISTICS PLUS LLC",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "ORB": {
    "code": "ORB",
    "company": "ORBIT CHEMICAL INDUSTRIES LIMITED",
    "city": "NAIROBI",
    "country": "Kenya"
  },
  "ORC": {
    "code": "ORC",
    "company": "ORCA CONTAINER ASSET MANAGEMENT",
    "city": "CAPE TOWN",
    "country": "South Africa"
  },
  "ORE": {
    "code": "ORE",
    "company": "ORIENT EXPRESS LLC",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "ORG": {
    "code": "ORG",
    "company": "OREGON TEKNOLOJI HIZMETLERI AS",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "ORI": {
    "code": "ORI",
    "company": "VINCI & CAMPAGNA SPA",
    "city": "Cagliari",
    "country": "Italy"
  },
  "ORN": {
    "code": "ORN",
    "company": "ORANO CYCLE",
    "city": "PIERRELATTE CEDEX",
    "country": "France"
  },
  "ORO": {
    "code": "ORO",
    "company": "AUTOGASERVICE S.R.L.",
    "city": "Maniago",
    "country": "Italy"
  },
  "ORP": {
    "code": "ORP",
    "company": "OJC OCETROVSKIY RIVER PORT",
    "city": "UST-KUT",
    "country": "Russian Federation"
  },
  "OSC": {
    "code": "OSC",
    "company": "OCI COMPANY LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "OSH": {
    "code": "OSH",
    "company": "OMEGA SERVIS HOLDING",
    "city": "ZELATOVICE",
    "country": "Czech Republic"
  },
  "OSK": {
    "code": "OSK",
    "company": "OCEANIC STAR LINE LTD",
    "city": "FUJAIRAH",
    "country": "United Arab Emirates"
  },
  "OSL": {
    "code": "OSL",
    "company": "OSTERMANN TRANSPORTE GMBH",
    "city": "Höxter",
    "country": "Germany"
  },
  "OSM": {
    "code": "OSM",
    "company": "KOPPERS PERFORMANCE CHEMICALS NEW ZEALAND",
    "city": "AUCKLAND",
    "country": "New Zealand"
  },
  "OSR": {
    "code": "OSR",
    "company": "SIEMENS GAS AND POWER GMBH & CO. KG",
    "city": "NURNBERG",
    "country": "Germany"
  },
  "OSS": {
    "code": "OSS",
    "company": "DEN HARTOGH LIQUID LOGISTICS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "OST": {
    "code": "OST",
    "company": "SUNCONTRACT GMBH",
    "city": "Basel",
    "country": "Switzerland"
  },
  "OTA": {
    "code": "OTA",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "OTE": {
    "code": "OTE",
    "company": "CAI INTERNATIONAL",
    "city": "SAN FRANCISCO, CA 94105",
    "country": "United States"
  },
  "OTF": {
    "code": "OTF",
    "company": "OTND ONET TECHNOLOGIES",
    "city": "PIERRELATTE",
    "country": "France"
  },
  "OTL": {
    "code": "OTL",
    "company": "OVERLAND TOTAL LOGISTIC SERVICES (M) SDN BHD",
    "city": "Penang",
    "country": "Malaysia"
  },
  "OTP": {
    "code": "OTP",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "OTS": {
    "code": "OTS",
    "company": "OTSUKA CHEMICAL CO LTD",
    "city": "OSAKA",
    "country": "Japan"
  },
  "OTT": {
    "code": "OTT",
    "company": "OTTEVANGER MILLING ENGINEERS",
    "city": "MOERKAPELLE",
    "country": "Netherlands"
  },
  "OUL": {
    "code": "OUL",
    "company": "OULUN KONTTIVUOKRAUS OY",
    "city": "Oulu",
    "country": "Finland"
  },
  "OUT": {
    "code": "OUT",
    "company": "OUTOTEC (FINLAND) OY",
    "city": "ESPOO",
    "country": "Finland"
  },
  "OVL": {
    "code": "OVL",
    "company": "O.V. LAHTINEN OY",
    "city": "HELSINKI",
    "country": "Finland"
  },
  "OVZ": {
    "code": "OVZ",
    "company": "JOINT VENTURE 'OVZ-TRANS'LLC",
    "city": "MINSK",
    "country": "Belarus"
  },
  "OWE": {
    "code": "OWE",
    "company": "OWENS ENERGY",
    "city": "SHOW LOW, AZ-85901",
    "country": "United States"
  },
  "OWH": {
    "code": "OWH",
    "company": "OY HACKLIN LTD",
    "city": "PORI",
    "country": "Finland"
  },
  "OWL": {
    "code": "OWL",
    "company": "ONE WAY LEASE INC.",
    "city": "OAKLAND, CA-94607",
    "country": "United States"
  },
  "OWM": {
    "code": "OWM",
    "company": "OILWELL MIDDLE EAST FZE",
    "city": "SHARJAH",
    "country": "United Arab Emirates"
  },
  "OWN": {
    "code": "OWN",
    "company": "TARROS SPA",
    "city": "LA SPEZIA  SP",
    "country": "Italy"
  },
  "OWS": {
    "code": "OWS",
    "company": "ODIS IRRIGATION EQUIPMENT (2002) LTD",
    "city": "PETAH-TIKYA",
    "country": "Israel"
  },
  "OXN": {
    "code": "OXN",
    "company": "SIPCAM OXON S.P.A",
    "city": "MILANO MI",
    "country": "Italy"
  },
  "OXY": {
    "code": "OXY",
    "company": "OXYMONTAGE",
    "city": "SAINT DIVY",
    "country": "France"
  },
  "OZM": {
    "code": "OZM",
    "company": "GJSC OTEW",
    "city": "Osipovichi",
    "country": "Belarus"
  },
  "OZT": {
    "code": "OZT",
    "company": "OZTURK NAKLIYAT VE INSAAT SAN.TIC.LTD.STI",
    "city": "Atasehir, Istanbul",
    "country": "Turkey"
  },
  "PAC": {
    "code": "PAC",
    "company": "XPO LOGISTICS",
    "city": "DUBLIN, OH 43016",
    "country": "United States"
  },
  "PAG": {
    "code": "PAG",
    "company": "PAGANELLA SPA",
    "city": "MANTOVA MN",
    "country": "Italy"
  },
  "PAL": {
    "code": "PAL",
    "company": "PAN ASIA LOGISTICS INDIA PVT LTD",
    "city": "Mylapore Chennai",
    "country": "India"
  },
  "PAN": {
    "code": "PAN",
    "company": "PANGAS",
    "city": "DAGMERSELLEN",
    "country": "Switzerland"
  },
  "PAR": {
    "code": "PAR",
    "company": "PARKOLIO SHIPPING COMPANY LIMITED",
    "city": "LIMASSOL",
    "country": "Cyprus"
  },
  "PAS": {
    "code": "PAS",
    "company": "PACIFIC LINES",
    "city": "Ho Chi Minh",
    "country": "Viet Nam"
  },
  "PAT": {
    "code": "PAT",
    "company": "PACIFICO CONTAINER LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "PAV": {
    "code": "PAV",
    "company": "LOGIKA LTD",
    "city": "ATHENS",
    "country": "Greece"
  },
  "PBB": {
    "code": "PBB",
    "company": "PBBPOLISUR S.R.L",
    "city": "BUENOS AIRES",
    "country": "Argentina"
  },
  "PBF": {
    "code": "PBF",
    "company": "PACIFIC BULK FUEL LIMITED",
    "city": "AUCKLAND",
    "country": "New Zealand"
  },
  "PBI": {
    "code": "PBI",
    "company": "PRO BOX INC.",
    "city": "HOUSTON, TX 77229-4425",
    "country": "United States"
  },
  "PBK": {
    "code": "PBK",
    "company": "PAN BRIDGE SHIPPING CO LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "PBS": {
    "code": "PBS",
    "company": "P. LEASING CORPORATION LIMITED",
    "city": "BELIZE CITY",
    "country": "Belize"
  },
  "PBX": {
    "code": "PBX",
    "company": "PREMIER BOX PTY LTD",
    "city": "BURPENGARY EAST",
    "country": "Australia"
  },
  "PCA": {
    "code": "PCA",
    "company": "GENERAL SERVICING AND TRADING LTD",
    "city": "PORT VILLA",
    "country": "Vanuatu"
  },
  "PCC": {
    "code": "PCC",
    "company": "PROJEXTRADE",
    "city": "PARIS",
    "country": "France"
  },
  "PCF": {
    "code": "PCF",
    "company": "TRANSPORTES PORTUARIOS S.A",
    "city": "BARCELONA",
    "country": "Spain"
  },
  "PCH": {
    "code": "PCH",
    "company": "PIETERSE CONTAINERHANDEL & TRANSPORT B.V.",
    "city": "TER AAR",
    "country": "Netherlands"
  },
  "PCI": {
    "code": "PCI",
    "company": "PACIFIC INTERNATIONAL LINES (PTE) LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "PCL": {
    "code": "PCL",
    "company": "PAN CONTINENTAL SHIPPING CO LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "PCM": {
    "code": "PCM",
    "company": "UNITED INITIATORS GMBH",
    "city": "PULLACH",
    "country": "Germany"
  },
  "PCT": {
    "code": "PCT",
    "company": "PC LOGISTIC SP. Z O.O.",
    "city": "DEBICA",
    "country": "Poland"
  },
  "PCV": {
    "code": "PCV",
    "company": "PEACOCK CONTAINER BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "PDA": {
    "code": "PDA",
    "company": "DALIAN JIFA BOHAI RIM CONTAINER LINES CO LTD",
    "city": "Dalian",
    "country": "China"
  },
  "PDI": {
    "code": "PDI",
    "company": "PERIMETER SOLUTIONS LP",
    "city": "CLAYTON, MO 63105",
    "country": "United States"
  },
  "PDL": {
    "code": "PDL",
    "company": "PACIFIC DIRECT LINE",
    "city": "AUCKLAND",
    "country": "New Zealand"
  },
  "PDQ": {
    "code": "PDQ",
    "company": "KATOEN NATIE TANK OPERATIONS N.V.",
    "city": "KALLO (KIELDRECHT)",
    "country": "Belgium"
  },
  "PEF": {
    "code": "PEF",
    "company": "VERSALIS FRANCE SAS",
    "city": "MARDYCK",
    "country": "France"
  },
  "PEG": {
    "code": "PEG",
    "company": "PRODIT ENGINEERING S.R.L.",
    "city": "SANTENA",
    "country": "Italy"
  },
  "PEI": {
    "code": "PEI",
    "company": "VERSALIS S.P.A",
    "city": "SAN DONATO MILANESE (MI)",
    "country": "Italy"
  },
  "PEN": {
    "code": "PEN",
    "company": "PENSPEN LTD",
    "city": "RICHMOND",
    "country": "United Kingdom"
  },
  "PER": {
    "code": "PER",
    "company": "PERTHON GROUP",
    "city": "Gothenburg",
    "country": "Sweden"
  },
  "PEX": {
    "code": "PEX",
    "company": "PENEX CONTAINER LINES (THAILAND) CO,LTD",
    "city": "BANGKOK",
    "country": "Thailand"
  },
  "PFC": {
    "code": "PFC",
    "company": "PETIT FORESTIER CONTAINER",
    "city": "VILLEPINTE",
    "country": "France"
  },
  "PFL": {
    "code": "PFL",
    "company": "PACIFIC FORUM LINE (NZ) LTD",
    "city": "AUCKLAND",
    "country": "New Zealand"
  },
  "PGA": {
    "code": "PGA",
    "company": "PAN-GULF LOGISTICS LTD",
    "city": "NANNING LIANGQING DISTRICT",
    "country": "China"
  },
  "PGH": {
    "code": "PGH",
    "company": "THE PASHA GROUP",
    "city": "SAN RAFAEL, CA-94903",
    "country": "United States"
  },
  "PGI": {
    "code": "PGI",
    "company": "PAO GAN INDUSTRIAL CO LTD",
    "city": "SIANSI",
    "country": "Taiwan, China"
  },
  "PGO": {
    "code": "PGO",
    "company": "FIVE OCEANS LOGISTICS LTD",
    "city": "Hong Kong",
    "country": "HK"
  },
  "PGT": {
    "code": "PGT",
    "company": "TOUAX",
    "city": "LA DEFENSE",
    "country": "France"
  },
  "PGX": {
    "code": "PGX",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "PHE": {
    "code": "PHE",
    "company": "AIR LIQUIDE GLOBAL HELIUM FZE",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "PHI": {
    "code": "PHI",
    "company": "AIR LIQUIDE PIPELINE UTILITIES SERVICES INC",
    "city": "CALAMBA, LAGUNA",
    "country": "Philippines"
  },
  "PHL": {
    "code": "PHL",
    "company": "SPINNAKER LEASING CORPORATION",
    "city": "San Francisco, CA 94111-2602",
    "country": "United States"
  },
  "PHM": {
    "code": "PHM",
    "company": "POLOMA MANAGEMENT LTD",
    "city": "ST PETERSBURG",
    "country": "Russian Federation"
  },
  "PHR": {
    "code": "PHR",
    "company": "PYEONG HWA REEFER SERVICES INC",
    "city": "BUSAN",
    "country": "Korea, Republic of"
  },
  "PIA": {
    "code": "PIA",
    "company": "SAS TRANSPORTS PIALLA",
    "city": "Pierrelatte",
    "country": "France"
  },
  "PIC": {
    "code": "PIC",
    "company": "AUTOTRASPORTI PICCININI S.R.L.",
    "city": "PARMA PR",
    "country": "Italy"
  },
  "PIF": {
    "code": "PIF",
    "company": "PEDDIE INVESTMENT TRUST",
    "city": "CAPE TOWN",
    "country": "South Africa"
  },
  "PIL": {
    "code": "PIL",
    "company": "PACIFIC INTERNATIONAL LINES (PTE) LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "PIM": {
    "code": "PIM",
    "company": "PIMK LTD",
    "city": "PLOVDIV",
    "country": "Bulgaria"
  },
  "PIZ": {
    "code": "PIZ",
    "company": "PRAXAIR PERU S.R.L",
    "city": "LIMA",
    "country": "Peru"
  },
  "PKE": {
    "code": "PKE",
    "company": "PANTOS LOGISTICS CO.,LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "PKK": {
    "code": "PKK",
    "company": "EKC INTERNATIONAL FZE",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "PLC": {
    "code": "PLC",
    "company": "PLCS INTERNATIONAL",
    "city": "HOOGVLIET",
    "country": "Netherlands"
  },
  "PLG": {
    "code": "PLG",
    "company": "GULBRANDSEN COMPANIES",
    "city": "CLINTON, NJ 08809",
    "country": "United States"
  },
  "PLN": {
    "code": "PLN",
    "company": "POLL NUSSBAUMER GMBH",
    "city": "GMUNDEN",
    "country": "Austria"
  },
  "PLT": {
    "code": "PLT",
    "company": "PANEUROPA TRANSPORT GMBH",
    "city": "Bakum",
    "country": "Germany"
  },
  "PLW": {
    "code": "PLW",
    "company": "INSPEKTORAT WSPARCIA SZ U.L DWERNICKIEGO 1",
    "city": "BYDGOSZCZ",
    "country": "Poland"
  },
  "PMA": {
    "code": "PMA",
    "company": "PETRO-TAINO TRANSPORT CORPORATION",
    "city": "Penuelas (Porto Rico)",
    "country": "United States"
  },
  "PMB": {
    "code": "PMB",
    "company": "PBM GAS LTD",
    "city": "ODESSA",
    "country": "Ukraine"
  },
  "PMK": {
    "code": "PMK",
    "company": "MDB CHEMICALS (INDIA) PVT LTD",
    "city": "Nasik Road - Dist. Nasik (MH)",
    "country": "India"
  },
  "PML": {
    "code": "PML",
    "company": "PERMA SHIPPING LINE PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "PMM": {
    "code": "PMM",
    "company": "PFEIFER SEIL- UND HEBETECHNIK GMBH",
    "city": "Memmingen",
    "country": "Germany"
  },
  "PMT": {
    "code": "PMT",
    "company": "PUUMAAILM LTD",
    "city": "JURI, HARJUMAA",
    "country": "Estonia"
  },
  "PMU": {
    "code": "PMU",
    "company": "PANALON MULTIMODAL",
    "city": "VILLARROBLEDO",
    "country": "Spain"
  },
  "PMX": {
    "code": "PMX",
    "company": "PSI TIRE SUPPLY LLC",
    "city": "CHRISTIANSTED, VI-00820",
    "country": "United States"
  },
  "PNA": {
    "code": "PNA",
    "company": "SPECTRA CHEMICALS AND COMMODITIES (SHANGHAI)",
    "city": "SHANGHAI",
    "country": "China"
  },
  "PNB": {
    "code": "PNB",
    "company": "JORN BOLDING A/S",
    "city": "Esbjerg V",
    "country": "Denmark"
  },
  "PNC": {
    "code": "PNC",
    "company": "PACIFIC ENVIRONMENTAL CORP.",
    "city": "HONOLULU, HI 96817",
    "country": "United States"
  },
  "PNE": {
    "code": "PNE",
    "company": "PACIFIC NORTHWEST EQUIPMENT INC",
    "city": "KENT, WA 98032",
    "country": "United States"
  },
  "PNH": {
    "code": "PNH",
    "company": "KQ SPEEDIE LUB 2 LLC",
    "city": "LIHUE,  HI 96766",
    "country": "United States"
  },
  "PNI": {
    "code": "PNI",
    "company": "PT. PELAYARAN NASIONAL INDONESIA",
    "city": "JAKARTA",
    "country": "Indonesia"
  },
  "PNS": {
    "code": "PNS",
    "company": "PANASYSTEM HANDELS GMBH",
    "city": "Vienna",
    "country": "Austria"
  },
  "PNW": {
    "code": "PNW",
    "company": "PACIFIC NORTHWEST EQUIPMENT INC",
    "city": "KENT, WA 98032",
    "country": "United States"
  },
  "POA": {
    "code": "POA",
    "company": "PRIMY OCEAN AIR LOGISTICS CO LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "POC": {
    "code": "POC",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "POL": {
    "code": "POL",
    "company": "PANOCEAN",
    "city": "Seoul",
    "country": "Korea, Republic of"
  },
  "POM": {
    "code": "POM",
    "company": "FSUE PRODUCTION ASSOCIATION MAYAK",
    "city": "OZERSK",
    "country": "Russian Federation"
  },
  "PON": {
    "code": "PON",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "POR": {
    "code": "POR",
    "company": "PORTEK SYSTEMS & EQUIPMENT PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "PPC": {
    "code": "PPC",
    "company": "KANTO-PPC INC",
    "city": "TAOYAN CITY",
    "country": "Taiwan, China"
  },
  "PPG": {
    "code": "PPG",
    "company": "ALTIVIA SPECIALTY CHEMICALS",
    "city": "Houston, TX 77002",
    "country": "United States"
  },
  "PPL": {
    "code": "PPL",
    "company": "BERTSCHI GLOBAL AG",
    "city": "DURRENASCH",
    "country": "Switzerland"
  },
  "PPS": {
    "code": "PPS",
    "company": "NEPTUNE PACIFIC LINE PTE LTD",
    "city": "AUCKLAND AIRPORT",
    "country": "New Zealand"
  },
  "PPT": {
    "code": "PPT",
    "company": "TANKGUARD B.V.",
    "city": "BARENDRECHT",
    "country": "Netherlands"
  },
  "PRA": {
    "code": "PRA",
    "company": "COOPERATIVA PARATORI GENOVA A R.L.",
    "city": "GENOVA",
    "country": "Italy"
  },
  "PRB": {
    "code": "PRB",
    "company": "PRS PASSIVE REFRIGERATION SOLUTIONS S.A.",
    "city": "LUGANO",
    "country": "Switzerland"
  },
  "PRC": {
    "code": "PRC",
    "company": "PRIDE-CHEM INDUSTRIES PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "PRF": {
    "code": "PRF",
    "company": "PRIMEFUELS LIMITED",
    "city": "Nairobi",
    "country": "Kenya"
  },
  "PRG": {
    "code": "PRG",
    "company": "PROGECO",
    "city": "MARSEILLE",
    "country": "France"
  },
  "PRK": {
    "code": "PRK",
    "company": "PROTANK LOGISTICS CO LTD",
    "city": "TAIPEI CITY",
    "country": "Taiwan, China"
  },
  "PRM": {
    "code": "PRM",
    "company": "PARAMAR S.A.",
    "city": "Asuncion",
    "country": "Paraguay"
  },
  "PRO": {
    "code": "PRO",
    "company": "PRO-TRANS A/S",
    "city": "HINNERUP",
    "country": "Denmark"
  },
  "PRS": {
    "code": "PRS",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "PRT": {
    "code": "PRT",
    "company": "GS LINES - TRANSPORTES MARITIMOS LDA",
    "city": "FUNCHAL-MADEIRA",
    "country": "Portugal"
  },
  "PRX": {
    "code": "PRX",
    "company": "PRAXAIR HELIUM",
    "city": "SPRING, TX-77380",
    "country": "United States"
  },
  "PSC": {
    "code": "PSC",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "PSD": {
    "code": "PSD",
    "company": "PIGOTT SHAFT DRILLING LIMITED",
    "city": "PRESTON",
    "country": "United Kingdom"
  },
  "PSE": {
    "code": "PSE",
    "company": "PANSTAR CO, LTD",
    "city": "BUSAN",
    "country": "Korea, Republic of"
  },
  "PSG": {
    "code": "PSG",
    "company": "PSG CORPORATION",
    "city": "BUSAN",
    "country": "Korea, Republic of"
  },
  "PSH": {
    "code": "PSH",
    "company": "PALMDEAL SHIPPING INC",
    "city": "BUENOS AIRES",
    "country": "Argentina"
  },
  "PSM": {
    "code": "PSM",
    "company": "TRADING LOGISTIC SAC SRL",
    "city": "La Spezia",
    "country": "Italy"
  },
  "PSN": {
    "code": "PSN",
    "company": "PARSONS CONTAINERS LTD",
    "city": "Stockton-on-Tees",
    "country": "United Kingdom"
  },
  "PSO": {
    "code": "PSO",
    "company": "PENTALVER TRANSPORT LIMITED",
    "city": "SOUTHAMPTON",
    "country": "United Kingdom"
  },
  "PSR": {
    "code": "PSR",
    "company": "PARATORI S.R.L.",
    "city": "GENOVA",
    "country": "Italy"
  },
  "PSS": {
    "code": "PSS",
    "company": "PENTALVER TRANSPORT LIMITED",
    "city": "SOUTHAMPTON",
    "country": "United Kingdom"
  },
  "PST": {
    "code": "PST",
    "company": "TETRA PAK PACKAGING SOLUTIONS SPA",
    "city": "MODENA",
    "country": "Italy"
  },
  "PSX": {
    "code": "PSX",
    "company": "PSG CONTAINER CO LTD",
    "city": "WANCHAI",
    "country": "HK"
  },
  "PTB": {
    "code": "PTB",
    "company": "GCATAINER BV",
    "city": "MOERDIJK",
    "country": "Netherlands"
  },
  "PTC": {
    "code": "PTC",
    "company": "MAMMOET EUROPE BV",
    "city": "SCHIEDAM",
    "country": "Netherlands"
  },
  "PTK": {
    "code": "PTK",
    "company": "PALTANK LTD",
    "city": "SOUTHPORT PR9OER",
    "country": "United Kingdom"
  },
  "PTL": {
    "code": "PTL",
    "company": "PLC  LTD",
    "city": "SAINT-PETERSBURG",
    "country": "Russian Federation"
  },
  "PTR": {
    "code": "PTR",
    "company": "POSTEN NORGE AS",
    "city": "MO I RANA",
    "country": "Norway"
  },
  "PTT": {
    "code": "PTT",
    "company": "PARITET LTD",
    "city": "Irkutsk",
    "country": "Russian Federation"
  },
  "PTV": {
    "code": "PTV",
    "company": "PISTON TANK CORPORATION",
    "city": "S.T LOUIS, MO 63026",
    "country": "United States"
  },
  "PUL": {
    "code": "PUL",
    "company": "PORTPACK UK LIMITED",
    "city": "Hucknall Nottinghamshire",
    "country": "United Kingdom"
  },
  "PVC": {
    "code": "PVC",
    "company": "PAC-VAN INC",
    "city": "INDIANAPOLIS, IN-46216",
    "country": "United States"
  },
  "PVD": {
    "code": "PVD",
    "company": "UNIT45 BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "PWL": {
    "code": "PWL",
    "company": "INTERSUL REPAROS E MANUT. CONT. LTDA.",
    "city": "Rio Grande",
    "country": "Brazil"
  },
  "PWR": {
    "code": "PWR",
    "company": "POWERTECH INC",
    "city": "MONROVIA",
    "country": "Liberia"
  },
  "PXA": {
    "code": "PXA",
    "company": "PRAXAIR ARGENTINA SRL",
    "city": "RICARDO ROJAS.TIGRE",
    "country": "Argentina"
  },
  "PXC": {
    "code": "PXC",
    "company": "KAWASAKI KISEN KAISHA LTD - K LINE",
    "city": "TOKYO",
    "country": "Japan"
  },
  "PYA": {
    "code": "PYA",
    "company": "POLISH YACHTING ASSOCIATION",
    "city": "WARSAW",
    "country": "Poland"
  },
  "PYH": {
    "code": "PYH",
    "company": "PISCINAS Y HORMIGONES 7 ISLAS, SLU",
    "city": "EL SAUZAL",
    "country": "Spain"
  },
  "QBC": {
    "code": "QBC",
    "company": "CUBICFARM MANUFACTURING CORP",
    "city": "Langley",
    "country": "Canada"
  },
  "QBL": {
    "code": "QBL",
    "company": "QBEX LOGISTICS BV",
    "city": "Ridderkerk",
    "country": "Netherlands"
  },
  "QBX": {
    "code": "QBX",
    "company": "CHS CONTAINER HANDEL GMBH",
    "city": "BREMEN",
    "country": "Germany"
  },
  "QIB": {
    "code": "QIB",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "QIM": {
    "code": "QIM",
    "company": "TCS TRANS S.L.",
    "city": "BARCELONA",
    "country": "Spain"
  },
  "QIN": {
    "code": "QIN",
    "company": "QINGBOX GROUP CORP.",
    "city": "Pittsford",
    "country": "United States"
  },
  "QNL": {
    "code": "QNL",
    "company": "QATAR NAVIGATION (QSC)",
    "city": "DOHA",
    "country": "Qatar"
  },
  "QNN": {
    "code": "QNN",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "QPC": {
    "code": "QPC",
    "company": "QUIMPAC DE COLOMBIA",
    "city": "CALI",
    "country": "Colombia"
  },
  "QQI": {
    "code": "QQI",
    "company": "QINGDAO WSR CONTAINER SERVICE CO., LTD",
    "city": "Qingdao",
    "country": "China"
  },
  "QSS": {
    "code": "QSS",
    "company": "PT. INDONESIA TSINGSHAN STAINLESS STEEL",
    "city": "JAKARTA",
    "country": "Indonesia"
  },
  "QTM": {
    "code": "QTM",
    "company": "QUANTUM FUEL SYSTEM LLC",
    "city": "Lake Forest",
    "country": "United States"
  },
  "RAC": {
    "code": "RAC",
    "company": "TRASMEDITERRANEA CARGO S.A.U.",
    "city": "BARCELONA",
    "country": "Spain"
  },
  "RAE": {
    "code": "RAE",
    "company": "REACH AMERICA  ESG, LTD.",
    "city": "SAN JOSE, CA 95110",
    "country": "United States"
  },
  "RAF": {
    "code": "RAF",
    "company": "PDAD PRODUCTOS DEL AIRE DOMINICANA SA",
    "city": "SAN PEDRO DE MACORIS",
    "country": "Dominican Republic"
  },
  "RAI": {
    "code": "RAI",
    "company": "RAINBOW CONTAINERS GMBH",
    "city": "APENSEN",
    "country": "Germany"
  },
  "RAL": {
    "code": "RAL",
    "company": "ROYAL ARCTIC LINE A/S",
    "city": "NUUK",
    "country": "Greenland"
  },
  "RAS": {
    "code": "RAS",
    "company": "RASEEF CONTAINERS SERVICES LLC",
    "city": "Amman",
    "country": "Jordan"
  },
  "RAV": {
    "code": "RAV",
    "company": "FLEXBOX COLUMBIA SAS",
    "city": "BOGOTA",
    "country": "Colombia"
  },
  "RAW": {
    "code": "RAW",
    "company": "QUANZHOU RENJIAN ANTONG LOGISTICS CO LTD",
    "city": "QUANZHOU CITY",
    "country": "China"
  },
  "RAY": {
    "code": "RAY",
    "company": "KNOW HOW TRADING B.V.",
    "city": "Willemstad",
    "country": "Netherlands"
  },
  "RBB": {
    "code": "RBB",
    "company": "BESED LLC",
    "city": "STERLITAMAK",
    "country": "Russian Federation"
  },
  "RBC": {
    "code": "RBC",
    "company": "RB HOLDING B.V.",
    "city": "ELSLOO",
    "country": "Netherlands"
  },
  "RBE": {
    "code": "RBE",
    "company": "RAT LOGISTICS GMBH",
    "city": "MANNHEIM",
    "country": "Germany"
  },
  "RBO": {
    "code": "RBO",
    "company": "REEFERBOX CO",
    "city": "Alleroed",
    "country": "Denmark"
  },
  "RCB": {
    "code": "RCB",
    "company": "BIGGE CRANE AND RIGGING CO",
    "city": "SAN LEANDRO, CA 94577",
    "country": "United States"
  },
  "RCD": {
    "code": "RCD",
    "company": "UNIT45 BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "RCI": {
    "code": "RCI",
    "company": "BRENNTAG SPA",
    "city": "Assago",
    "country": "Italy"
  },
  "RCL": {
    "code": "RCL",
    "company": "R.M.I. CHEMICAL LOGISTICS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "RCO": {
    "code": "RCO",
    "company": "ROYAL CARGO, INC.",
    "city": "Paranaque City",
    "country": "Philippines"
  },
  "RCT": {
    "code": "RCT",
    "company": "DEFENCE CONTAINER MANAGEMENT SERVICE",
    "city": "BICESTER, OXFORDSHIRE OX26 6JP",
    "country": "United Kingdom"
  },
  "RCX": {
    "code": "RCX",
    "company": "R.C.C  CONTAINERS BV",
    "city": "ROTTERDAM-BOTLEK",
    "country": "Netherlands"
  },
  "RDC": {
    "code": "RDC",
    "company": "LIGNES MARITIMES CONGOLAISES (L.M.C)",
    "city": "KINSHASA",
    "country": "Congo"
  },
  "RDH": {
    "code": "RDH",
    "company": "ELECTRONIC MATERIALS",
    "city": "SEELZE",
    "country": "Germany"
  },
  "RDK": {
    "code": "RDK",
    "company": "YILMAZ COMPANY LTD",
    "city": "Nairobi",
    "country": "Kenya"
  },
  "RDL": {
    "code": "RDL",
    "company": "RODELLA TRASPORTI SRL",
    "city": "MEDOLE (MN)",
    "country": "Italy"
  },
  "RDM": {
    "code": "RDM",
    "company": "RODYMAR SHIPPING CO",
    "city": "PORT SAID",
    "country": "Egypt"
  },
  "RDR": {
    "code": "RDR",
    "company": "ROEDER CARTAGE CO INC",
    "city": "LIMA, OH 45801",
    "country": "United States"
  },
  "RDV": {
    "code": "RDV",
    "company": "REDAVIA GMBH",
    "city": "MUNICH",
    "country": "Germany"
  },
  "RDX": {
    "code": "RDX",
    "company": "RADIX CO., LTD.",
    "city": "Seoul",
    "country": "Korea, Republic of"
  },
  "REA": {
    "code": "REA",
    "company": "ANDRADE & SANTOS LOCACAO DE MODULOS E IMPORTAÇÃO LTDA",
    "city": "SANTOS",
    "country": "Brazil"
  },
  "RED": {
    "code": "RED",
    "company": "SITEBOX STORAGE",
    "city": "WICHITA, KS-67217",
    "country": "United States"
  },
  "REF": {
    "code": "REF",
    "company": "PETIT FORESTIER CONTAINER NEDERLAND BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "REG": {
    "code": "REG",
    "company": "RCL FEEDER PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "REL": {
    "code": "REL",
    "company": "RINCHEM EQUIPMENT LEASING",
    "city": "ALBUQUERQUE",
    "country": "United States"
  },
  "RET": {
    "code": "RET",
    "company": "REPARACIONES TECNICAS DE CONTENEDORES S.L.",
    "city": "Pinto",
    "country": "Spain"
  },
  "REV": {
    "code": "REV",
    "company": "REVISS SERVICES (UK) LTD",
    "city": "ABINGDON, OXON",
    "country": "United Kingdom"
  },
  "REW": {
    "code": "REW",
    "company": "RENTAL WORLD HOLDING BV",
    "city": "OUD-BEIJERLAND",
    "country": "Netherlands"
  },
  "REY": {
    "code": "REY",
    "company": "REYSAS TASIMACILIK ve LOJISTIK TIC.LTD.STI.",
    "city": "ÜSKÜDAR / ISTANBUL",
    "country": "Turkey"
  },
  "RFB": {
    "code": "RFB",
    "company": "BIMICON CONTAINER SERVICE GMBH",
    "city": "Hamburg",
    "country": "Germany"
  },
  "RFC": {
    "code": "RFC",
    "company": "RAFFLES LEASE PTE Ltd.",
    "city": "Singapore",
    "country": "Singapore"
  },
  "RFL": {
    "code": "RFL",
    "company": "R.M.I. CHEMICAL LOGISTICS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "RFS": {
    "code": "RFS",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LIMITED",
    "city": "SAN FRANCISCO, CA 94108",
    "country": "United States"
  },
  "RFT": {
    "code": "RFT",
    "company": "REFTERMINAL LTD",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "RFU": {
    "code": "RFU",
    "company": "CTC-LOGISTIKA",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "RGA": {
    "code": "RGA",
    "company": "RENEGADE GAS PTY LTD",
    "city": "Ingleburn",
    "country": "Australia"
  },
  "RGC": {
    "code": "RGC",
    "company": "RAVA GROUP CONTAINER SERVICES",
    "city": "Medley, FL-33178",
    "country": "United States"
  },
  "RGH": {
    "code": "RGH",
    "company": "RICH GLORY (HONG KONG) LIMITED",
    "city": "HONG KONG",
    "country": "HK"
  },
  "RGT": {
    "code": "RGT",
    "company": "RIGTANK GLOBAL S.A.",
    "city": "PANAMA CITY",
    "country": "Panama"
  },
  "RHA": {
    "code": "RHA",
    "company": "LLC RAILGO",
    "city": "CITY OF UFA (REPUBLIC OF BASHKORTOSTAN)",
    "country": "Russian Federation"
  },
  "RHN": {
    "code": "RHN",
    "company": "LIMITED LIABILITY COMPANY \"VENTEK\"",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "RHT": {
    "code": "RHT",
    "company": "TRANSPORT ROGER HEINEN",
    "city": "EUPEN",
    "country": "Belgium"
  },
  "RIB": {
    "code": "RIB",
    "company": "RIBARROJA STATION S.L.",
    "city": "Ribarroja del Turia (Valencia)",
    "country": "Spain"
  },
  "RIC": {
    "code": "RIC",
    "company": "RINNEN GMBH & CO KG",
    "city": "MOERS",
    "country": "Germany"
  },
  "RIE": {
    "code": "RIE",
    "company": "OOSTERHOUT CONTAINER TERMINAL",
    "city": "Oosterhout",
    "country": "Netherlands"
  },
  "RIG": {
    "code": "RIG",
    "company": "CARGOSTORE WORLDWIDE TRADING LIMITED",
    "city": "LONDON SW19 7QD",
    "country": "United Kingdom"
  },
  "RII": {
    "code": "RII",
    "company": "RINA SERVICES S.P.A.",
    "city": "GENOVA",
    "country": "Italy"
  },
  "RIL": {
    "code": "RIL",
    "company": "ZEAMARINE CARRIER GMBH",
    "city": "BREMEN",
    "country": "Germany"
  },
  "RIN": {
    "code": "RIN",
    "company": "HERMANN RINNEN JUNIOR",
    "city": "MOERS",
    "country": "Germany"
  },
  "RIV": {
    "code": "RIV",
    "company": "RIVOIRA OPERATIONS",
    "city": "MILANO",
    "country": "Italy"
  },
  "RJC": {
    "code": "RJC",
    "company": "FLEXBOX COLUMBIA SAS",
    "city": "BOGOTA",
    "country": "Colombia"
  },
  "RKH": {
    "code": "RKH",
    "company": "RIN KAGAKU KOGYO CO LTD",
    "city": "TOYAMA",
    "country": "Japan"
  },
  "RKN": {
    "code": "RKN",
    "company": "ROKONORD LTD",
    "city": "SAINT PETERSBOURG",
    "country": "Russian Federation"
  },
  "RKS": {
    "code": "RKS",
    "company": "LLC \"RUSKITSOYUZ LK\"",
    "city": "NOVOSIBIRSK",
    "country": "Russian Federation"
  },
  "RLC": {
    "code": "RLC",
    "company": "DEFENCE CONTAINER MANAGEMENT SERVICE",
    "city": "BICESTER, OXFORDSHIRE OX26 6JP",
    "country": "United Kingdom"
  },
  "RLS": {
    "code": "RLS",
    "company": "RR SHIPPING PRIVATE LIMITED",
    "city": "NAVI MUMBAI",
    "country": "India"
  },
  "RLT": {
    "code": "RLT",
    "company": "RAFFLES LEASE PTE Ltd.",
    "city": "Singapore",
    "country": "Singapore"
  },
  "RMC": {
    "code": "RMC",
    "company": "MULTISTAR LEASING LTD",
    "city": "CHESHIRE CW11 1BA",
    "country": "United Kingdom"
  },
  "RMG": {
    "code": "RMG",
    "company": "RESOLVE MARINE GROUP, INC.",
    "city": "Fort Lauderdale, FL 33316",
    "country": "United States"
  },
  "RMM": {
    "code": "RMM",
    "company": "EMERSON CLIMATE TECHNOLOGIES",
    "city": "HOJBJERG",
    "country": "Denmark"
  },
  "RMS": {
    "code": "RMS",
    "company": "TRANSPORTS F.RAMOS,S.L",
    "city": "REUS",
    "country": "Spain"
  },
  "RMT": {
    "code": "RMT",
    "company": "R.M.I. CHEMICAL LOGISTICS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "RNM": {
    "code": "RNM",
    "company": "RNM TRANSPORTE QUIMICOS LDA",
    "city": "VILA NOVA FAMALICAO",
    "country": "Portugal"
  },
  "ROA": {
    "code": "ROA",
    "company": "ROAM GLOBAL SOLUTIONS LLC",
    "city": "HOUSTON, TX-77035",
    "country": "United States"
  },
  "ROB": {
    "code": "ROB",
    "company": "ROBERTO BUCCI S.P.A",
    "city": "NAPLES",
    "country": "Italy"
  },
  "ROC": {
    "code": "ROC",
    "company": "ROLAND UMSCHLAGSGESELLSCHAFT",
    "city": "Bremen",
    "country": "Germany"
  },
  "ROD": {
    "code": "ROD",
    "company": "MAURITIUS SHIPPING CORPORATION LTD",
    "city": "PORT LOUIS",
    "country": "Mauritius"
  },
  "ROE": {
    "code": "ROE",
    "company": "CONTAINERHANDEL CARSTEN ROEHE GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "ROS": {
    "code": "ROS",
    "company": "ROOS SPEDITION GMBH",
    "city": "DURMERSHEIM",
    "country": "Germany"
  },
  "ROU": {
    "code": "ROU",
    "company": "ITA ROUTE",
    "city": "ARQUES",
    "country": "France"
  },
  "ROX": {
    "code": "ROX",
    "company": "EMCO MARINE LTD /  CONTAINERS DPT.",
    "city": "HAIFA",
    "country": "Israel"
  },
  "RPB": {
    "code": "RPB",
    "company": "GVT INTERMODAL FREIGHTMANAGEMENT B.V.",
    "city": "TILBURG",
    "country": "Netherlands"
  },
  "RPC": {
    "code": "RPC",
    "company": "TIC PETERSBURG LTD",
    "city": "ST PETERSBURG",
    "country": "Russian Federation"
  },
  "RPG": {
    "code": "RPG",
    "company": "GROWTECH INDUSTRIES LLC",
    "city": "BUFFALO",
    "country": "United States"
  },
  "RPU": {
    "code": "RPU",
    "company": "WINCH ENERGY",
    "city": "Sevenoaks",
    "country": "United Kingdom"
  },
  "RRI": {
    "code": "RRI",
    "company": "TRITON INTERNATIONAL LTD",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "RRM": {
    "code": "RRM",
    "company": "REMAIN GMBH CONTAINER DEPOT AND REPAIR",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "RRP": {
    "code": "RRP",
    "company": "RESPOL RESINAS S.A.",
    "city": "Pinheiros-Leiria",
    "country": "Portugal"
  },
  "RRR": {
    "code": "RRR",
    "company": "REBLOCK CONTAINER AS.",
    "city": "LILLESTROM",
    "country": "Norway"
  },
  "RRS": {
    "code": "RRS",
    "company": "ARCTIC REEFER SERVICE & REPAIR B.V",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "RRX": {
    "code": "RRX",
    "company": "PETIT FORESTIER CONTAINER NEDERLAND BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "RSB": {
    "code": "RSB",
    "company": "RSB LOGISTIC PROJEKSPEDITION GMBH",
    "city": "COLOGNE",
    "country": "Germany"
  },
  "RSC": {
    "code": "RSC",
    "company": "RED SEA NAVIGATION CO.",
    "city": "PORT SAID",
    "country": "Egypt"
  },
  "RSE": {
    "code": "RSE",
    "company": "REEFER SALES EUROPE BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "RSF": {
    "code": "RSF",
    "company": "RSL FREIGHT CO LLC",
    "city": "PO BOX 112076 - 009714 DUBAI",
    "country": "United Arab Emirates"
  },
  "RSG": {
    "code": "RSG",
    "company": "SIA REAL INVESTMENT",
    "city": "RIGA",
    "country": "Latvia"
  },
  "RSI": {
    "code": "RSI",
    "company": "RASA INDUSTRIES, LTD.",
    "city": "Tokyo",
    "country": "Japan"
  },
  "RSL": {
    "code": "RSL",
    "company": "MATSON SOUTH PACIFIC LTD",
    "city": "AUCKLAND",
    "country": "New Zealand"
  },
  "RSP": {
    "code": "RSP",
    "company": "REACH HOLDING GROUP (SHANGHAI) CO.,LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "RSS": {
    "code": "RSS",
    "company": "ROYAL WOLF TRADING AUSTRALIA PTY LTD",
    "city": "GORDON",
    "country": "Australia"
  },
  "RST": {
    "code": "RST",
    "company": "UNIT45 BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "RSV": {
    "code": "RSV",
    "company": "JSC REFSERVICE",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "RSX": {
    "code": "RSX",
    "company": "RS PIONEER NAVIGATION LTD.",
    "city": "APIA",
    "country": "Samoa"
  },
  "RTA": {
    "code": "RTA",
    "company": "ROBOTANKS",
    "city": "KYIV",
    "country": "Ukraine"
  },
  "RTB": {
    "code": "RTB",
    "company": "RTSB GMBH",
    "city": "Friedrichsdorf",
    "country": "Germany"
  },
  "RTG": {
    "code": "RTG",
    "company": "TEHGAZ LLC",
    "city": "ORENBURG",
    "country": "Russian Federation"
  },
  "RTH": {
    "code": "RTH",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "RTM": {
    "code": "RTM",
    "company": "ROTTERDAM TANKCONTAINERS LEASING",
    "city": "BERGENKOEK",
    "country": "Netherlands"
  },
  "RTS": {
    "code": "RTS",
    "company": "ROBERT KUKLA GMBH",
    "city": "MUNICH",
    "country": "Germany"
  },
  "RTT": {
    "code": "RTT",
    "company": "RASA TECHNOLOGY TAIWAN LTD",
    "city": "TAICHUNG COUNTY",
    "country": "Taiwan, China"
  },
  "RTV": {
    "code": "RTV",
    "company": "DAILY RENTAL SERVICES BV",
    "city": "VLAARDINGEN",
    "country": "Netherlands"
  },
  "RUB": {
    "code": "RUB",
    "company": "RUBINO RS S.R.L.",
    "city": "MONOPOLI BA",
    "country": "Italy"
  },
  "RUN": {
    "code": "RUN",
    "company": "GEOTRANS",
    "city": "LE PORT CEDEX",
    "country": "Reunion"
  },
  "RUR": {
    "code": "RUR",
    "company": "EESTI CHEM OU",
    "city": "TALLIN",
    "country": "Estonia"
  },
  "RUS": {
    "code": "RUS",
    "company": "BALTICA TRANS",
    "city": "ST PETERSBURG",
    "country": "Russian Federation"
  },
  "RVH": {
    "code": "RVH",
    "company": "RAMA VESSEL HANDLERS PVT LTD",
    "city": "GANDHIDHAM",
    "country": "India"
  },
  "RWA": {
    "code": "RWA",
    "company": "TRISTAR CONTAINER SERVICES (ASIA) PVT. LTD.",
    "city": "CHENNAI",
    "country": "India"
  },
  "RWL": {
    "code": "RWL",
    "company": "ROYAL WOLF TRADING AUSTRALIA PTY LTD",
    "city": "GORDON",
    "country": "Australia"
  },
  "RWT": {
    "code": "RWT",
    "company": "ROYAL WOLF TRADING AUSTRALIA PTY LTD",
    "city": "GORDON",
    "country": "Australia"
  },
  "RWY": {
    "code": "RWY",
    "company": "AO AK ZHELEZNYE DOROGI YAKUTII",
    "city": "Aldan",
    "country": "Russian Federation"
  },
  "RYC": {
    "code": "RYC",
    "company": "C.A.K DE RIJKE BV",
    "city": "SPIJKENISSE",
    "country": "Netherlands"
  },
  "RYD": {
    "code": "RYD",
    "company": "QINGDAO RUIYUAN DATONG CONTAINER CO., LTD",
    "city": "QINGDAO",
    "country": "China"
  },
  "RYL": {
    "code": "RYL",
    "company": "BEIJING RUI LI HENG YI LOGISTICS TECHNOLOGY PLC",
    "city": "Beijing",
    "country": "China"
  },
  "RZD": {
    "code": "RZD",
    "company": "PJSC TRANSCONTAINER",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "SAB": {
    "code": "SAB",
    "company": "SOLARIS CHEMTECH INDUSTRIES LTD.",
    "city": "RATADIYA",
    "country": "India"
  },
  "SAC": {
    "code": "SAC",
    "company": "WALLENIUS WILHELMSEN LOGISTICS AS",
    "city": "LYSAKER",
    "country": "Norway"
  },
  "SAE": {
    "code": "SAE",
    "company": "EDF",
    "city": "SAINT-DENIS",
    "country": "France"
  },
  "SAI": {
    "code": "SAI",
    "company": "STAR S.P.A.",
    "city": "LOCATE DI TRIULZI MI",
    "country": "Italy"
  },
  "SAJ": {
    "code": "SAJ",
    "company": "GOODRICH MARITIME LLC",
    "city": "BUR DUBAI, DUBAI",
    "country": "United Arab Emirates"
  },
  "SAL": {
    "code": "SAL",
    "company": "NTC TANK CONTAINER SERVICES BV",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "SAM": {
    "code": "SAM",
    "company": "SAMUDA CHEMICAL COMPLEX LIMITED",
    "city": "Dhaka",
    "country": "Bangladesh"
  },
  "SAN": {
    "code": "SAN",
    "company": "SAMSKIP HF",
    "city": "REYKJAVIK",
    "country": "Iceland"
  },
  "SAO": {
    "code": "SAO",
    "company": "ARMASUISSE",
    "city": "BERNE",
    "country": "Switzerland"
  },
  "SAT": {
    "code": "SAT",
    "company": "DSV SPA",
    "city": "PIOLTELLO, MI",
    "country": "Italy"
  },
  "SAV": {
    "code": "SAV",
    "company": "SAR TRASPORTI SCPA",
    "city": "RAVENNA",
    "country": "Italy"
  },
  "SAX": {
    "code": "SAX",
    "company": "NOBLE CONTAINER LEASING LIMITED",
    "city": "JORDAN, KOWLOON",
    "country": "HK"
  },
  "SAZ": {
    "code": "SAZ",
    "company": "SEVEN ASSET LTD",
    "city": "IPSWICH IP1 1XF",
    "country": "United Kingdom"
  },
  "SBA": {
    "code": "SBA",
    "company": "SEIBOW LOGISTICS LIMITED",
    "city": "KWUN TONG, HONG KONG",
    "country": "HK"
  },
  "SBG": {
    "code": "SBG",
    "company": "SEA STAR LINE LLC",
    "city": "JACKSONVILLE, FL 32256",
    "country": "United States"
  },
  "SBH": {
    "code": "SBH",
    "company": "C. STEINWEG OMAN LLC",
    "city": "LIWA",
    "country": "Oman"
  },
  "SBI": {
    "code": "SBI",
    "company": "SEA BOX, INC",
    "city": "CINNAMINSON, NJ 08077, New Jersey",
    "country": "United States"
  },
  "SBK": {
    "code": "SBK",
    "company": "SEABROOK TANK SERVICES LTD",
    "city": "BRIGHTON-LE-SANDS / LIVERPOOL",
    "country": "United Kingdom"
  },
  "SBO": {
    "code": "SBO",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "SBP": {
    "code": "SBP",
    "company": "SATYESH BRINECHEM PRIVATE LIMITED",
    "city": "Ahmedabad",
    "country": "India"
  },
  "SBS": {
    "code": "SBS",
    "company": "SBS CONTAINERS LIMITED",
    "city": "LIMASSOL",
    "country": "Cyprus"
  },
  "SBT": {
    "code": "SBT",
    "company": "ООО SIBTRANS",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "SBV": {
    "code": "SBV",
    "company": "SCHEIER BRENNSTOFFE UND BEGRUNUNGSTECHNIK GMBH",
    "city": "BURS",
    "country": "Austria"
  },
  "SCB": {
    "code": "SCB",
    "company": "SCANDI BULK AB",
    "city": "Gothenburg",
    "country": "Sweden"
  },
  "SCC": {
    "code": "SCC",
    "company": "SHARKCAGE AS",
    "city": "Fornebu",
    "country": "Norway"
  },
  "SCD": {
    "code": "SCD",
    "company": "SPECIALTY COATINGS LTD",
    "city": "LANCASHIRE, BB3 2EN",
    "country": "United Kingdom"
  },
  "SCE": {
    "code": "SCE",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "SCF": {
    "code": "SCF",
    "company": "SCF GROUP",
    "city": "Adelaide",
    "country": "Australia"
  },
  "SCH": {
    "code": "SCH",
    "company": "SYNTHESIS CHIMICA S.R.L.",
    "city": "CASTELLO D'AGOGNA PV",
    "country": "Italy"
  },
  "SCK": {
    "code": "SCK",
    "company": "SCK-CEN",
    "city": "MOL",
    "country": "Belgium"
  },
  "SCL": {
    "code": "SCL",
    "company": "SMART CONTAINER LINES PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "SCM": {
    "code": "SCM",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "SCN": {
    "code": "SCN",
    "company": "SEA- CARGO AS",
    "city": "NESTTUN",
    "country": "Norway"
  },
  "SCP": {
    "code": "SCP",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "SCQ": {
    "code": "SCQ",
    "company": "SEA CONNECT UAB",
    "city": "KLAIPEDA",
    "country": "Lithuania"
  },
  "SCS": {
    "code": "SCS",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "SCT": {
    "code": "SCT",
    "company": "SOUTHCOAST TRANSPORT",
    "city": "",
    "country": "Ireland"
  },
  "SCV": {
    "code": "SCV",
    "company": "SECORA CONTAINERS SP. Z O.O.",
    "city": "Wroclaw",
    "country": "Poland"
  },
  "SCX": {
    "code": "SCX",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "SCZ": {
    "code": "SCZ",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "SDD": {
    "code": "SDD",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "SDI": {
    "code": "SDI",
    "company": "TREDI",
    "city": "Saint-Vulbas",
    "country": "France"
  },
  "SDL": {
    "code": "SDL",
    "company": "SODIMAC S.A.",
    "city": "Santiago",
    "country": "Chile"
  },
  "SDM": {
    "code": "SDM",
    "company": "MARS COMPANY",
    "city": "takasaki  city gunnma",
    "country": "Japan"
  },
  "SDN": {
    "code": "SDN",
    "company": "SEDNA CONTAINERS B.V.",
    "city": "Landsmeer",
    "country": "Netherlands"
  },
  "SDO": {
    "code": "SDO",
    "company": "SPINNAKER EQUIPMENT SERVICES INC.",
    "city": "San Francisco, CA 94111",
    "country": "United States"
  },
  "SDU": {
    "code": "SDU",
    "company": "SDW SHIPPING BV",
    "city": "Den haag",
    "country": "Netherlands"
  },
  "SDW": {
    "code": "SDW",
    "company": "FRANS DE WIT INTERNATIONAL BV",
    "city": "MOERDIJK",
    "country": "Netherlands"
  },
  "SEA": {
    "code": "SEA",
    "company": "SEALAND",
    "city": "FLORHAM PARK, NJ 07932",
    "country": "United States"
  },
  "SEB": {
    "code": "SEB",
    "company": "ACTIVE CONTAINERS SERVICES BV",
    "city": "ROTTERDAM-VONDELINGPLAAT",
    "country": "Netherlands"
  },
  "SEC": {
    "code": "SEC",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "SED": {
    "code": "SED",
    "company": "SCHENKER (DEUTSCHLAND) AG",
    "city": "EUSKIRCHEN",
    "country": "Germany"
  },
  "SEE": {
    "code": "SEE",
    "company": "FUGRO GEOSERVICES",
    "city": "FALMOUTH, CORNWALL",
    "country": "United Kingdom"
  },
  "SEF": {
    "code": "SEF",
    "company": "CROWLEY CARIBBEAN SERVICES LLC",
    "city": "JACKSONVILLE, FL 32225",
    "country": "United States"
  },
  "SEG": {
    "code": "SEG",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "SEI": {
    "code": "SEI",
    "company": "TOV SOLEXIM",
    "city": "Kyiv",
    "country": "Ukraine"
  },
  "SEK": {
    "code": "SEK",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "SEL": {
    "code": "SEL",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "SEM": {
    "code": "SEM",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "SES": {
    "code": "SES",
    "company": "SPINNAKER EQUIPMENT SERVICES INC.",
    "city": "San Francisco, CA 94111",
    "country": "United States"
  },
  "SET": {
    "code": "SET",
    "company": "ORYS TRICASTIN",
    "city": "PIERRELATTE",
    "country": "France"
  },
  "SEU": {
    "code": "SEU",
    "company": "STONEHAVEN ENGINEERING LTD",
    "city": "Stonehaven",
    "country": "United Kingdom"
  },
  "SEV": {
    "code": "SEV",
    "company": "SEVEL SPA SOCIETA EUROPA",
    "city": "ATESSA-CH",
    "country": "Italy"
  },
  "SEX": {
    "code": "SEX",
    "company": "CONTAINER NORWAY AS",
    "city": "FJELLSTRAND",
    "country": "Norway"
  },
  "SEY": {
    "code": "SEY",
    "company": "SEYCHELLES PETROLEUM COMPANY LTD",
    "city": "VICTORIA, MAHE",
    "country": "Seychelles"
  },
  "SFB": {
    "code": "SFB",
    "company": "SOLVAY FLUOR GMBH",
    "city": "HANNOVER",
    "country": "Germany"
  },
  "SFE": {
    "code": "SFE",
    "company": "SF ENTERPRISES",
    "city": "OAKLAND, CA 94607",
    "country": "United States"
  },
  "SFF": {
    "code": "SFF",
    "company": "ACB AGENCIES",
    "city": "EKEREN",
    "country": "Belgium"
  },
  "SFI": {
    "code": "SFI",
    "company": "SFI TRANSIT",
    "city": "Pointe aux sables",
    "country": "Mauritius"
  },
  "SFR": {
    "code": "SFR",
    "company": "SCANFOR BVBA",
    "city": "WOMMELGEM",
    "country": "Belgium"
  },
  "SFS": {
    "code": "SFS",
    "company": "CJSC \"SEAFOOD-SERVICE\"",
    "city": "Minsk",
    "country": "Belarus"
  },
  "SFU": {
    "code": "SFU",
    "company": "SAN FU CHEMICAL CO, LTD",
    "city": "SHAN-HUA, TAINAN",
    "country": "Taiwan, China"
  },
  "SGA": {
    "code": "SGA",
    "company": "MINISTERE DE LA DEFENSE (DCSID)",
    "city": "VERSAILLES",
    "country": "France"
  },
  "SGB": {
    "code": "SGB",
    "company": "SOUTHERN INDUSTRIAL GAS SDN BHD",
    "city": "SENAI JOHOR",
    "country": "Malaysia"
  },
  "SGC": {
    "code": "SGC",
    "company": "SOGECO INTERNATIONAL SA",
    "city": "PARADISO - LUGANO",
    "country": "Switzerland"
  },
  "SGI": {
    "code": "SGI",
    "company": "SCOTLAND GAS NETWORKS",
    "city": "NEWBRIDGE, EH28 8TG - SCOTLAND",
    "country": "United Kingdom"
  },
  "SGK": {
    "code": "SGK",
    "company": "SYNERGY GASES (K) LTD",
    "city": "MOMBASA",
    "country": "Kenya"
  },
  "SGL": {
    "code": "SGL",
    "company": "SICGIL INDUSTRIAL GASES LTD",
    "city": "CHENNAI - TAMIL NADU",
    "country": "India"
  },
  "SGM": {
    "code": "SGM",
    "company": "SIGMA DENIZCILIK KONT.VE.LOJ.SANT.TIC.",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "SGN": {
    "code": "SGN",
    "company": "SOGIN SPA",
    "city": "CAORSO PC",
    "country": "Italy"
  },
  "SGP": {
    "code": "SGP",
    "company": "SSB CRYOGENIC EQUIPMENT PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "SGR": {
    "code": "SGR",
    "company": "SOGECO INTERNATIONAL SA",
    "city": "PARADISO - LUGANO",
    "country": "Switzerland"
  },
  "SGS": {
    "code": "SGS",
    "company": "SOUTHERN GAS SERVICES LIMITED",
    "city": "Rolleston West",
    "country": "New Zealand"
  },
  "SGT": {
    "code": "SGT",
    "company": "SERLUX",
    "city": "LUXEMBOURG",
    "country": "Luxembourg"
  },
  "SHC": {
    "code": "SHC",
    "company": "SHANGHAI STAR HOUSE CO . LTD",
    "city": "Huinan TOWN",
    "country": "China"
  },
  "SHE": {
    "code": "SHE",
    "company": "SHL OFFSHORE CONTRACTORS BV",
    "city": "ZOETERMEER",
    "country": "Netherlands"
  },
  "SHK": {
    "code": "SHK",
    "company": "SEA HAWK LINES PVT LIMITED",
    "city": "CHENNAI",
    "country": "India"
  },
  "SHL": {
    "code": "SHL",
    "company": "SMITH-HOLLAND BV",
    "city": "SPIJKENISSE",
    "country": "Netherlands"
  },
  "SHM": {
    "code": "SHM",
    "company": "SACHEM JAPAN GODO KAISHA",
    "city": "HIGASHIOSAKA",
    "country": "Japan"
  },
  "SHN": {
    "code": "SHN",
    "company": "KARL SCHMIDT SPEDITION GMBH & CO KG",
    "city": "HEILBRONN",
    "country": "Germany"
  },
  "SHO": {
    "code": "SHO",
    "company": "SHIPPER OWNED CONTAINER LLC",
    "city": "MCKINNEY, TX 77380",
    "country": "United States"
  },
  "SHS": {
    "code": "SHS",
    "company": "LLC \"TEKHGAZ\"",
    "city": "Blagoveshchensk (Amur region)",
    "country": "Russian Federation"
  },
  "SHT": {
    "code": "SHT",
    "company": "SHEGA TRANS SHA",
    "city": "Tirana",
    "country": "Albania"
  },
  "SIA": {
    "code": "SIA",
    "company": "SIAD SOCIETA ITALIANA ACETILENE & DERIVATI SPA",
    "city": "BERGAMO  BG",
    "country": "Italy"
  },
  "SIB": {
    "code": "SIB",
    "company": "GE SIMONS INTERNATIONAL TRANSPORT BV",
    "city": "HILVARENBEEK",
    "country": "Netherlands"
  },
  "SIC": {
    "code": "SIC",
    "company": "SUNRISE CONTAINER SERVICE CO. LTD",
    "city": "Mahe",
    "country": "Seychelles"
  },
  "SID": {
    "code": "SID",
    "company": "CHEMIKALIEN UND FLUESSIGKEITSTRANSPORTE A. SIEPMANN GMBH",
    "city": "DUISBURG",
    "country": "Germany"
  },
  "SIE": {
    "code": "SIE",
    "company": "SOMER S ISLE SHIPPING LTD",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "SIF": {
    "code": "SIF",
    "company": "NEWPORT EUROPE B.V.",
    "city": "Moerdijk",
    "country": "Netherlands"
  },
  "SII": {
    "code": "SII",
    "company": "THE SHIPPING CORPORATION OF INDIA LTD",
    "city": "MUMBAI",
    "country": "India"
  },
  "SIK": {
    "code": "SIK",
    "company": "SAMUDERA SHIPPING LINE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "SIL": {
    "code": "SIL",
    "company": "SUN INTERMODAL LIMITED",
    "city": "London",
    "country": "United Kingdom"
  },
  "SIM": {
    "code": "SIM",
    "company": "SHANGHAI SAFE-TRANSPORT CHEMICAL LOGISTICS CO.,LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "SIN": {
    "code": "SIN",
    "company": "BBC CHARTERING GMBH - AS AGENTS",
    "city": "LEER",
    "country": "Germany"
  },
  "SIP": {
    "code": "SIP",
    "company": "PT PERUSAHAAN PELAYARAN NUSANTARA PANURJWAN",
    "city": "JAKARTA",
    "country": "Indonesia"
  },
  "SIR": {
    "code": "SIR",
    "company": "OOO VED CONTACT",
    "city": "NOVOSIBIRSK,",
    "country": "Russian Federation"
  },
  "SIS": {
    "code": "SIS",
    "company": "ELBTAINER TRADING GMBH",
    "city": "Barsbüttel",
    "country": "Germany"
  },
  "SIT": {
    "code": "SIT",
    "company": "SITC CONTAINER LINES CO LTD",
    "city": "WANCHAI",
    "country": "HK"
  },
  "SJI": {
    "code": "SJI",
    "company": "PT. SINARJAYA INTIMPERKASA",
    "city": "SURABAYA",
    "country": "Indonesia"
  },
  "SJK": {
    "code": "SJK",
    "company": "SARJAK CONTAINER LINES PVT LTD",
    "city": "MUMBAI",
    "country": "India"
  },
  "SJR": {
    "code": "SJR",
    "company": "SUZHOU CRYSTAL CLEAR CHEMICAL CO,LTD",
    "city": "SUZHOU",
    "country": "China"
  },
  "SJU": {
    "code": "SJU",
    "company": "S JONES CONTAINERS LTD",
    "city": "WALSALL",
    "country": "United Kingdom"
  },
  "SKA": {
    "code": "SKA",
    "company": "SKAFF CRYOGENICS INC",
    "city": "BRENTWOOD, NH-03833",
    "country": "United States"
  },
  "SKB": {
    "code": "SKB",
    "company": "SWEDISH NUCLEAR FUEL AND WASTE MANAGEMENT",
    "city": "OSKARSHAMN",
    "country": "Sweden"
  },
  "SKC": {
    "code": "SKC",
    "company": "ALE HEAVYLIFT LIMITED",
    "city": "STAFFORDSHIRE",
    "country": "United Kingdom"
  },
  "SKH": {
    "code": "SKH",
    "company": "SINOKOR MERCHANT MARINE CORP",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "SKI": {
    "code": "SKI",
    "company": "BLUE SKY INTERMODAL (UK) LTD",
    "city": "Marlow (Buckinghamshir)",
    "country": "United Kingdom"
  },
  "SKL": {
    "code": "SKL",
    "company": "SINOKOR MERCHANT MARINE CORP",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "SKM": {
    "code": "SKM",
    "company": "MUTTI SRO",
    "city": "TRNAVA",
    "country": "Slovakia"
  },
  "SKN": {
    "code": "SKN",
    "company": "KUEHNE + NAGEL AS",
    "city": "Oslo",
    "country": "Norway"
  },
  "SKR": {
    "code": "SKR",
    "company": "SINOKOR MERCHANT MARINE CORP",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "SKY": {
    "code": "SKY",
    "company": "CAI INTERNATIONAL",
    "city": "SAN FRANCISCO, CA 94105",
    "country": "United States"
  },
  "SLB": {
    "code": "SLB",
    "company": "ETUDES ET PRODUCTION SCHLUMBERGER",
    "city": "CLAMART",
    "country": "France"
  },
  "SLC": {
    "code": "SLC",
    "company": "SPECTRA TANKTAINER LOGISTICS PTE. LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "SLE": {
    "code": "SLE",
    "company": "SPINNAKER EQUIPMENT SERVICES INC.",
    "city": "San Francisco, CA 94111",
    "country": "United States"
  },
  "SLH": {
    "code": "SLH",
    "company": "SEA LLOYD SHIPPING LINES PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "SLJ": {
    "code": "SLJ",
    "company": "SEALEASE N.V.",
    "city": "WILLEMSTAD, CURACAO,",
    "country": "Netherlands Antilles"
  },
  "SLP": {
    "code": "SLP",
    "company": "SEAGULL CONTAINER SERVICES PTE. LTD.",
    "city": "Singapore",
    "country": "Singapore"
  },
  "SLR": {
    "code": "SLR",
    "company": "SILVA LS SP. Z O.O.",
    "city": "Mielec",
    "country": "Poland"
  },
  "SLS": {
    "code": "SLS",
    "company": "SPINNAKER LEASING CORPORATION",
    "city": "San Francisco, CA 94111-2602",
    "country": "United States"
  },
  "SLT": {
    "code": "SLT",
    "company": "SCM LINES TRANSPORTES MARITIMOS",
    "city": "FUNCHAL-MADEIRA",
    "country": "Portugal"
  },
  "SLW": {
    "code": "SLW",
    "company": "SALTWORKS TECHOLOGIES INC",
    "city": "Richmond",
    "country": "Canada"
  },
  "SLZ": {
    "code": "SLZ",
    "company": "CS LEASING PTE LTD",
    "city": "",
    "country": "Singapore"
  },
  "SMA": {
    "code": "SMA",
    "company": "SMART CONTAINERS KAMIL BIESZK",
    "city": "BANINO",
    "country": "Poland"
  },
  "SMC": {
    "code": "SMC",
    "company": "SM CONTAINER LINES",
    "city": "Busan,",
    "country": "Korea, Republic of"
  },
  "SMD": {
    "code": "SMD",
    "company": "3M COMPANY",
    "city": "ST PAUL, MN 55144-1000",
    "country": "United States"
  },
  "SME": {
    "code": "SME",
    "company": "SMET BUILDING PRODUCTS LIMITED",
    "city": "Newry",
    "country": "United Kingdom"
  },
  "SMH": {
    "code": "SMH",
    "company": "STAALDUINEN LOGISTICS",
    "city": "MAASDIJK",
    "country": "Netherlands"
  },
  "SMI": {
    "code": "SMI",
    "company": "ANGOLANA DA NAVEGACAO",
    "city": "LUANDA",
    "country": "Angola"
  },
  "SML": {
    "code": "SML",
    "company": "SEABOARD MARINE LTD",
    "city": "MIAMI, FL 33166",
    "country": "United States"
  },
  "SMP": {
    "code": "SMP",
    "company": "SCS MULTIPORT B.V.",
    "city": "AMSTERDAM",
    "country": "Netherlands"
  },
  "SMQ": {
    "code": "SMQ",
    "company": "SMQ LTD",
    "city": "TAIPEI CITY,",
    "country": "Taiwan, China"
  },
  "SMR": {
    "code": "SMR",
    "company": "SOCIETE DE MAINTENANCE ET",
    "city": "SETE",
    "country": "France"
  },
  "SMS": {
    "code": "SMS",
    "company": "PT.SULAWESI MINING INVESTMENT",
    "city": "JAKARTA",
    "country": "Indonesia"
  },
  "SMT": {
    "code": "SMT",
    "company": "SEMIMATEX CO.,LTD",
    "city": "Longyan",
    "country": "China"
  },
  "SMU": {
    "code": "SMU",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "SNA": {
    "code": "SNA",
    "company": "SMARTLAGER NORGE AS",
    "city": "STAVERN",
    "country": "Norway"
  },
  "SNB": {
    "code": "SNB",
    "company": "SINOTRANS CONTAINER LINES CO.,LTD.",
    "city": "SHANGHAI",
    "country": "China"
  },
  "SNC": {
    "code": "SNC",
    "company": "INTERNATIONAL CONTAINER LEASE INC",
    "city": "BEIJING",
    "country": "China"
  },
  "SND": {
    "code": "SND",
    "company": "CANDO LOGISTICS LLC",
    "city": "Berlin, WI 54923",
    "country": "United States"
  },
  "SNF": {
    "code": "SNF",
    "company": "SNF S.A.S",
    "city": "ANDREZIEUX CEDEX",
    "country": "France"
  },
  "SNG": {
    "code": "SNG",
    "company": "SINOTRANS (GERMANY) GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "SNH": {
    "code": "SNH",
    "company": "SINOTRANS CONTAINER LINES CO.,LTD.",
    "city": "SHANGHAI",
    "country": "China"
  },
  "SNI": {
    "code": "SNI",
    "company": "STOLT TANK CONTAINERS LEASING LTD",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "SNS": {
    "code": "SNS",
    "company": "SLOMAN NEPTUNE SCHIFFAHRTS - AG",
    "city": "BREMEN",
    "country": "Germany"
  },
  "SNT": {
    "code": "SNT",
    "company": "STOLT TANK CONTAINERS LEASING LTD",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "SNU": {
    "code": "SNU",
    "company": "SITA TRANSPORT BV",
    "city": "ALPHEN AAN DEN RIJN",
    "country": "Netherlands"
  },
  "SOA": {
    "code": "SOA",
    "company": "SOL AVIATION SERVICES LIMITED",
    "city": "ST MICHAEL",
    "country": "Barbados"
  },
  "SOB": {
    "code": "SOB",
    "company": "SOBOLAK INTERNATIONAL GMBH",
    "city": "LEOBENDORF",
    "country": "Austria"
  },
  "SOC": {
    "code": "SOC",
    "company": "STAR SERVICE SRL",
    "city": "GENOVA  GE",
    "country": "Italy"
  },
  "SOE": {
    "code": "SOE",
    "company": "HANG ZHOU ZHE SHANG INDUSTRY TRADING",
    "city": "HANGZHOU",
    "country": "China"
  },
  "SOF": {
    "code": "SOF",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "SOI": {
    "code": "SOI",
    "company": "SHANDONG OCEAN INTERNATIONAL (HONG KONG) LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "SOK": {
    "code": "SOK",
    "company": "SICOM S.P.A",
    "city": "CHERASCO (CN)",
    "country": "Italy"
  },
  "SOL": {
    "code": "SOL",
    "company": "C.T.S  SRL",
    "city": "MONZA (MB)",
    "country": "Italy"
  },
  "SOM": {
    "code": "SOM",
    "company": "SOMAL",
    "city": "LE LAMENTIN CEDEX 2",
    "country": "France"
  },
  "SOT": {
    "code": "SOT",
    "company": "SOLETANCHE BACHY INTERNATIONAL",
    "city": "Montereau Fault Yonne",
    "country": "France"
  },
  "SOU": {
    "code": "SOU",
    "company": "SOUILLAC INDUSTRIES LTD.",
    "city": "Ebene",
    "country": "Mauritius"
  },
  "SOX": {
    "code": "SOX",
    "company": "AIR LIQUIDE SINGAPORE PRIVATE LIMITED",
    "city": "Singapore",
    "country": "Singapore"
  },
  "SPB": {
    "code": "SPB",
    "company": "SOUTHWEST CONTAINER SOLUTIONS",
    "city": "Scottsboro, AL 35769",
    "country": "United States"
  },
  "SPC": {
    "code": "SPC",
    "company": "SHANGHAI HAI HUA SHIPPING CO LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "SPD": {
    "code": "SPD",
    "company": "SPEDYCJA POLSKA  SPEDCONT SPOLKA ZO.O.",
    "city": "LODZ",
    "country": "Poland"
  },
  "SPE": {
    "code": "SPE",
    "company": "CIMC CONTAINERS HOLDING COMPANY LTD",
    "city": "GUANGDONG",
    "country": "China"
  },
  "SPH": {
    "code": "SPH",
    "company": "SPHB",
    "city": "ST PIERRE",
    "country": "Reunion"
  },
  "SPI": {
    "code": "SPI",
    "company": "SAS SPINDRIFT",
    "city": "Saint Philibert",
    "country": "France"
  },
  "SPK": {
    "code": "SPK",
    "company": "SPINNAKER LEASING CORPORATION",
    "city": "San Francisco, CA 94111-2602",
    "country": "United States"
  },
  "SPL": {
    "code": "SPL",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "SPN": {
    "code": "SPN",
    "company": "PT SALAM PACIFIC INDONESIA LINES",
    "city": "SURABAYA",
    "country": "Indonesia"
  },
  "SPP": {
    "code": "SPP",
    "company": "SHIVA PHARMACHEM LTD",
    "city": "GUJARAT",
    "country": "India"
  },
  "SPT": {
    "code": "SPT",
    "company": "SCHENK-PAPENDRECHT BV",
    "city": "PAPENDRECHT",
    "country": "Netherlands"
  },
  "SPW": {
    "code": "SPW",
    "company": "ICS TERMINALS (UK) LIMITED d/b/a SPACEWISE",
    "city": "RAINHAM, ESSEX RM13 8EU",
    "country": "United Kingdom"
  },
  "SRA": {
    "code": "SRA",
    "company": "ENDEL SRA",
    "city": "LILLE CEDEX",
    "country": "France"
  },
  "SRB": {
    "code": "SRB",
    "company": "LEIDOS",
    "city": "LONG BEACH, MS 39560",
    "country": "United States"
  },
  "SRC": {
    "code": "SRC",
    "company": "NORDIC BULKERS AB",
    "city": "GOTHENBURG",
    "country": "Sweden"
  },
  "SRE": {
    "code": "SRE",
    "company": "SARENS NV",
    "city": "WOLVERTEM",
    "country": "Belgium"
  },
  "SRG": {
    "code": "SRG",
    "company": "ROSI INVEST A/S",
    "city": "TANAGER",
    "country": "Norway"
  },
  "SRL": {
    "code": "SRL",
    "company": "SUPER RACK CO., LTD",
    "city": "Seoul",
    "country": "Korea, Republic of"
  },
  "SRP": {
    "code": "SRP",
    "company": "SARP INTERMODAL HIZMETLERI AS",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "SRS": {
    "code": "SRS",
    "company": "SKYROS MARITIME CORPORATION C/O ANDROS MARITIME AGENCIES LTD",
    "city": "LONDON, WC1A 1HB",
    "country": "United Kingdom"
  },
  "SRT": {
    "code": "SRT",
    "company": "SE.TRA.S. SRL",
    "city": "NARNI SCALO (TR)",
    "country": "Italy"
  },
  "SSB": {
    "code": "SSB",
    "company": "SPECIALIST SERVICES LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "SSC": {
    "code": "SSC",
    "company": "SHANGHAI SHOWA CHEMICALS CO.LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "SSD": {
    "code": "SSD",
    "company": "SASCO INTERNATIONAL SHIPPING CO,LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "SSF": {
    "code": "SSF",
    "company": "JSC SAKHALIN SHIPPING COMPANY LT",
    "city": "KHOLMSK",
    "country": "Russian Federation"
  },
  "SSG": {
    "code": "SSG",
    "company": "SURGUTNEFTEGAS PJSC",
    "city": "SURGUT",
    "country": "Russian Federation"
  },
  "SSM": {
    "code": "SSM",
    "company": "AVANA LOGISTEK LIMITED",
    "city": "Mumbai",
    "country": "India"
  },
  "SSR": {
    "code": "SSR",
    "company": "SSREEFER-SOLUCOES EM SERVICOS REEFER LTDA",
    "city": "ITAJAI-SANTA CATARINA",
    "country": "Brazil"
  },
  "SSS": {
    "code": "SSS",
    "company": "SEA SANDS SHIPPING L.L.C",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "SSW": {
    "code": "SSW",
    "company": "NANTONG CIMC SPECIAL LOGISTICS EQUIPMENT DEVELOPMENT CO., LTD.",
    "city": "GUIYANG CITY",
    "country": "China"
  },
  "SSX": {
    "code": "SSX",
    "company": "SAMSON SERVICE LLC",
    "city": "Saint-Petersburg",
    "country": "Russian Federation"
  },
  "STB": {
    "code": "STB",
    "company": "STOLT TANK CONTAINERS LEASING LTD",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "STC": {
    "code": "STC",
    "company": "FT MARINE SAS",
    "city": "FUVEAU",
    "country": "France"
  },
  "STF": {
    "code": "STF",
    "company": "\"STANDART-F\" LLC",
    "city": "Novorossiysk",
    "country": "Russian Federation"
  },
  "STG": {
    "code": "STG",
    "company": "SHOWA SPECIALTY GAS (TAIWAN) CO LTD",
    "city": "TAIPEI",
    "country": "Taiwan, China"
  },
  "STJ": {
    "code": "STJ",
    "company": "SCHENKER INC.",
    "city": "Freeport, NY 11520",
    "country": "United States"
  },
  "STK": {
    "code": "STK",
    "company": "SANTOKU CHEMICAL INDUSTRIES CO. LTD",
    "city": "SENDAI",
    "country": "Japan"
  },
  "STL": {
    "code": "STL",
    "company": "SHIPPING TRADING &LIGHTERAGE CO (STALCO)",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "STM": {
    "code": "STM",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "STN": {
    "code": "STN",
    "company": "SAMATAINERS",
    "city": "VIENNE CEDEX",
    "country": "France"
  },
  "STR": {
    "code": "STR",
    "company": "SEA STAR LINE LLC",
    "city": "JACKSONVILLE, FL 32256",
    "country": "United States"
  },
  "STU": {
    "code": "STU",
    "company": "STUBBE B.V.",
    "city": "GOUDA",
    "country": "Netherlands"
  },
  "STV": {
    "code": "STV",
    "company": "SUPERCONSORZIO TRASPORTI LUCANI",
    "city": "Viggiano",
    "country": "Italy"
  },
  "STW": {
    "code": "STW",
    "company": "SETTENTRIONALE TRASPORTI SPA",
    "city": "POSSAGNO (TV)",
    "country": "Italy"
  },
  "STX": {
    "code": "STX",
    "company": "STAXXON LLC",
    "city": "MONTCLAIR",
    "country": "United States"
  },
  "STZ": {
    "code": "STZ",
    "company": "SETOLAZAR ENERGIA Y MEDIOAMBIANTE",
    "city": "Humanes de Madrid",
    "country": "Spain"
  },
  "SUD": {
    "code": "SUD",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "SUG": {
    "code": "SUG",
    "company": "SUPAGAS PTY LTD",
    "city": "VICTORIA",
    "country": "Australia"
  },
  "SUI": {
    "code": "SUI",
    "company": "SWISSTERMINAL AG",
    "city": "FRENKENDORF",
    "country": "Switzerland"
  },
  "SUM": {
    "code": "SUM",
    "company": "SUMISHO GLOBAL LOGISTICS (CHINA) CO LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "SUN": {
    "code": "SUN",
    "company": "NICHIRIN GROUP CO",
    "city": "TOKYO",
    "country": "Japan"
  },
  "SUT": {
    "code": "SUT",
    "company": "SUTTONS INTERNATIONAL LTD",
    "city": "WIDNES, CHESHIRE",
    "country": "United Kingdom"
  },
  "SUZ": {
    "code": "SUZ",
    "company": "SUREZING CONTAINER INTERNATIONAL LTD",
    "city": "TORTOLA",
    "country": "Virgin Islands, British"
  },
  "SVD": {
    "code": "SVD",
    "company": "SAMSKIP MULTIMODAL BV",
    "city": "GENEMUIDEN",
    "country": "Netherlands"
  },
  "SVG": {
    "code": "SVG",
    "company": "SWAGEMAKERS INTERMODAAL TRANSPORT",
    "city": "Westdorpe",
    "country": "Netherlands"
  },
  "SVK": {
    "code": "SVK",
    "company": "SUOMEN VUOKRAKONTTI OY",
    "city": "RAJAMAKI",
    "country": "Finland"
  },
  "SVM": {
    "code": "SVM",
    "company": "SCHAVEMAKER LOGISTICS B.V",
    "city": "BEVERWIJK",
    "country": "Netherlands"
  },
  "SVN": {
    "code": "SVN",
    "company": "MINISTRY OF DEFENCE REPUBLIC OF SLOVENIA",
    "city": "LJUBLJANA",
    "country": "Slovenia"
  },
  "SVR": {
    "code": "SVR",
    "company": "SEVENR SIA",
    "city": "Riga",
    "country": "Latvia"
  },
  "SVS": {
    "code": "SVS",
    "company": "MLI- MEDITRANNEAN LOGISTICS INVESTIMENTS LTD",
    "city": "NICOSIA",
    "country": "Cyprus"
  },
  "SVW": {
    "code": "SVW",
    "company": "CHS CONTAINER HANDEL GMBH",
    "city": "BREMEN",
    "country": "Germany"
  },
  "SWC": {
    "code": "SWC",
    "company": "SWISSCONTAINER AG",
    "city": "GWATT",
    "country": "Switzerland"
  },
  "SWE": {
    "code": "SWE",
    "company": "DEN HARTOGH LIQUID LOGISTICS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "SWF": {
    "code": "SWF",
    "company": "SWIFT TRANSPORT INTERNATIONAL LOGISTICS PTE .LTD",
    "city": "TIANJIN",
    "country": "China"
  },
  "SWL": {
    "code": "SWL",
    "company": "SAMSKIP MCL BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "SWP": {
    "code": "SWP",
    "company": "SIEMENS WIND POWER A/S",
    "city": "BRANDE",
    "country": "Denmark"
  },
  "SWS": {
    "code": "SWS",
    "company": "SWISS WATER POWER INTERNATIONAL S.A.",
    "city": "GENEVE",
    "country": "Switzerland"
  },
  "SWT": {
    "code": "SWT",
    "company": "STOLT TANK CONTAINERS LEASING LTD",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "SXO": {
    "code": "SXO",
    "company": "STELLA EXPRESS (SINGAPORE) PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "SXR": {
    "code": "SXR",
    "company": "SOLAR EXPRESS LLC",
    "city": "Ulyanovsk",
    "country": "Russian Federation"
  },
  "SXS": {
    "code": "SXS",
    "company": "SOLVAY SPECIALTY POLYMERS ITALY SPA",
    "city": "BOLLATE (MI)",
    "country": "Italy"
  },
  "SXX": {
    "code": "SXX",
    "company": "KIRSEN GLOBAL SECURITY GMBH",
    "city": "Berlin",
    "country": "Germany"
  },
  "SYB": {
    "code": "SYB",
    "company": "JIANGSU XINYUE ASPHALT HI-TECH CO.,LTD",
    "city": "Zhenjiang",
    "country": "China"
  },
  "SYG": {
    "code": "SYG",
    "company": "SHENYANG SHIYUAN ENERGY CO., LTD.",
    "city": "Shenyang",
    "country": "China"
  },
  "SYP": {
    "code": "SYP",
    "company": "SAMYOUNG PURE CHEMICALS CO LTD",
    "city": "DONGNAM-GU, CHEONAN-SI",
    "country": "Korea, Republic of"
  },
  "SZA": {
    "code": "SZA",
    "company": "SAFRICAS/ SAFGAZ",
    "city": "KINSHASA",
    "country": "Congo"
  },
  "SZD": {
    "code": "SZD",
    "company": "VICONT TRADING GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "SZL": {
    "code": "SZL",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "SZR": {
    "code": "SZR",
    "company": "LLC SYNTEZRAIL",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "TAB": {
    "code": "TAB",
    "company": "STOLT TANK CONTAINERS LEASING LTD",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "TAC": {
    "code": "TAC",
    "company": "EAST CIRCLE INVESTMENT COPR",
    "city": "BP 0830-01580",
    "country": "Panama"
  },
  "TAF": {
    "code": "TAF",
    "company": "TERRES AUSTRALES ET ANTARCTIQUES FRANCAISES",
    "city": "SAINT PIERRE",
    "country": "France"
  },
  "TAH": {
    "code": "TAH",
    "company": "TA ASIA LOGISTICS PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "TAI": {
    "code": "TAI",
    "company": "STOLT TANK CONTAINERS LEASING LTD",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "TAL": {
    "code": "TAL",
    "company": "ALFRED TALKE GMBH & CO.KG",
    "city": "HURTH",
    "country": "Germany"
  },
  "TAN": {
    "code": "TAN",
    "company": "TACON GMBH TANK & CONTAINER SERVICES",
    "city": "BRAAK",
    "country": "Germany"
  },
  "TAP": {
    "code": "TAP",
    "company": "TOKUYAMA ELECTRONIC CHEMICALS PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "TAR": {
    "code": "TAR",
    "company": "TARROS SPA",
    "city": "LA SPEZIA  SP",
    "country": "Italy"
  },
  "TAS": {
    "code": "TAS",
    "company": "TANKSPAN LEASING LIMITED",
    "city": "SURREY GU8 6BQ",
    "country": "United Kingdom"
  },
  "TAT": {
    "code": "TAT",
    "company": "TRINIDAD DISTILLERS LTD",
    "city": "LAVENTILLE",
    "country": "Trinidad and Tobago"
  },
  "TAX": {
    "code": "TAX",
    "company": "TANK ASSET MANAGEMENT",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "TAY": {
    "code": "TAY",
    "company": "EUROPEA DE CONTENEDORES S.A. EUCONSA",
    "city": "EL PALMAR - MURCIA",
    "country": "Spain"
  },
  "TBB": {
    "code": "TBB",
    "company": "CHINA RAILWAY TIELONG CONTAINER LOG CORP LTD",
    "city": "BEIJING",
    "country": "China"
  },
  "TBE": {
    "code": "TBE",
    "company": "XINJIANG TBEA GROUP LOGISTICS CO.,LTD.",
    "city": "Changji",
    "country": "China"
  },
  "TBG": {
    "code": "TBG",
    "company": "CHINA RAILWAY TIELONG CONTAINER LOG CORP LTD",
    "city": "BEIJING",
    "country": "China"
  },
  "TBJ": {
    "code": "TBJ",
    "company": "CHINA RAILWAY CONTAINER TRANSPORT CORP",
    "city": "BEIJING",
    "country": "China"
  },
  "TBK": {
    "code": "TBK",
    "company": "TEIJIN LOGISTICS CO LTD",
    "city": "Matsuyama",
    "country": "Japan"
  },
  "TBL": {
    "code": "TBL",
    "company": "CHINA RAILWAY TIELONG CONTAINER LOGISTICS CO., LTD. BEIJING SPECIAL CONTAINER TECHNOLOGY DEVELOPMENT CENTER",
    "city": "DALIAN",
    "country": "China"
  },
  "TBP": {
    "code": "TBP",
    "company": "CHINA RAILWAY TIELONG CONTAINER LOG CORP LTD",
    "city": "BEIJING",
    "country": "China"
  },
  "TBS": {
    "code": "TBS",
    "company": "TIMBER BUILDING SYSTEMS PTY LTD",
    "city": "Dandenong",
    "country": "Australia"
  },
  "TBU": {
    "code": "TBU",
    "company": "TRUE BLUE CONTAINERS (2005) PTY LTD",
    "city": "MIDVALE",
    "country": "Australia"
  },
  "TBX": {
    "code": "TBX",
    "company": "BOMAG MARINI EQUIPAMENTOS LTDA",
    "city": "CACHOEIRINHA",
    "country": "Brazil"
  },
  "TCA": {
    "code": "TCA",
    "company": "BLUE BALTIC SHIPPING & TRADING LIMITED",
    "city": "LA HABANA",
    "country": "Cuba"
  },
  "TCC": {
    "code": "TCC",
    "company": "TANK CONTAINER CLEANING SERVICES S.R.L.",
    "city": "NAPOLI",
    "country": "Italy"
  },
  "TCD": {
    "code": "TCD",
    "company": "JIUJIANG TINCI MATERIALS TECHNOLOGY CO., LTD",
    "city": "JIUJIANG",
    "country": "China"
  },
  "TCE": {
    "code": "TCE",
    "company": "TASEK CORPORATION BHD",
    "city": "KUALA LUMPUR",
    "country": "Malaysia"
  },
  "TCG": {
    "code": "TCG",
    "company": "TRANSPORTS COTTARD-GLENAT CHIMIE",
    "city": "SALAISE-SUR-SANNE",
    "country": "France"
  },
  "TCH": {
    "code": "TCH",
    "company": "TORAN SOLUTION LTD C/O SANTOS CONTAINER LTDA",
    "city": "SAD VICENTE",
    "country": "Brazil"
  },
  "TCI": {
    "code": "TCI",
    "company": "TITAN CONTAINERS A/S",
    "city": "TAASTRUP",
    "country": "Denmark"
  },
  "TCK": {
    "code": "TCK",
    "company": "TRITON INTERNATIONAL LTD",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "TCL": {
    "code": "TCL",
    "company": "TAL INTERNATIONAL",
    "city": "PURCHASE, NY 10577-2135",
    "country": "United States"
  },
  "TCM": {
    "code": "TCM",
    "company": "TCS TRANS S.L.",
    "city": "BARCELONA",
    "country": "Spain"
  },
  "TCN": {
    "code": "TCN",
    "company": "TRITON INTERNATIONAL LTD",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "TCO": {
    "code": "TCO",
    "company": "JSC TECO",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "TCR": {
    "code": "TCR",
    "company": "TRANS-CHINA EXPRESS (HK) CO LIMITED",
    "city": "KOWLOON BAY",
    "country": "HK"
  },
  "TCS": {
    "code": "TCS",
    "company": "TANKCON B.V",
    "city": "SPIJKENISSE",
    "country": "Netherlands"
  },
  "TCT": {
    "code": "TCT",
    "company": "TI. CI. TI. SA",
    "city": "BALERNA",
    "country": "Switzerland"
  },
  "TCU": {
    "code": "TCU",
    "company": "TRIFLEET LEASING HOLDING B.V.",
    "city": "DORDRECHT",
    "country": "Netherlands"
  },
  "TCV": {
    "code": "TCV",
    "company": "TWS TANKCONTAINER-LEASING GMBH & CO KG",
    "city": "Hamburg",
    "country": "Germany"
  },
  "TCX": {
    "code": "TCX",
    "company": "TRANSMIT CONTAINERS LTD",
    "city": "ABERDEEN",
    "country": "United Kingdom"
  },
  "TDE": {
    "code": "TDE",
    "company": "TDS SHIPPING LLC",
    "city": "Dubai",
    "country": "United Arab Emirates"
  },
  "TDI": {
    "code": "TDI",
    "company": "TORANG DARYA SHIPPING CO LTD",
    "city": "TEHRAN",
    "country": "Iran, Islamic Republic of"
  },
  "TDL": {
    "code": "TDL",
    "company": "TRANS DISTANCE LINE LTD",
    "city": "HONG KONG",
    "country": "HK"
  },
  "TDR": {
    "code": "TDR",
    "company": "TAL INTERNATIONAL",
    "city": "PURCHASE, NY 10577-2135",
    "country": "United States"
  },
  "TDT": {
    "code": "TDT",
    "company": "TAL INTERNATIONAL",
    "city": "PURCHASE, NY 10577-2135",
    "country": "United States"
  },
  "TEA": {
    "code": "TEA",
    "company": "TEAKDECKING SYSTEMS INC",
    "city": "SARASOTA, FL 34243",
    "country": "United States"
  },
  "TEC": {
    "code": "TEC",
    "company": "TECHNOLOGY INDUSTRIAL GASES PRODUCTION CO",
    "city": "FAHEHEEL",
    "country": "Kuwait"
  },
  "TED": {
    "code": "TED",
    "company": "AGUNSA-AGENCIA UNIVERSALES S.A",
    "city": "SANTIAGO",
    "country": "Chile"
  },
  "TEE": {
    "code": "TEE",
    "company": "CRYO-TEC BVBA",
    "city": "MERKSEM",
    "country": "Belgium"
  },
  "TEG": {
    "code": "TEG",
    "company": "PT TEMAS TBK",
    "city": "JAKARTA UTARA",
    "country": "Indonesia"
  },
  "TEI": {
    "code": "TEI",
    "company": "BEIJING TRANS EURASIA INTERNATIONAL LOGISTICS LTD.",
    "city": "Beijing",
    "country": "China"
  },
  "TEL": {
    "code": "TEL",
    "company": "TRANS EATON INTERNATIONAL ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "TEM": {
    "code": "TEM",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "TEN": {
    "code": "TEN",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LIMITED",
    "city": "SAN FRANCISCO, CA 94108",
    "country": "United States"
  },
  "TEO": {
    "code": "TEO",
    "company": "TRANSGEO LLC (LIMITED LIABILITY COMPANY)",
    "city": "Saint-Petersburg",
    "country": "Russian Federation"
  },
  "TER": {
    "code": "TER",
    "company": "TECNIRUTA CONCISA S.A.",
    "city": "SANTURCE VIZCAYA",
    "country": "Spain"
  },
  "TET": {
    "code": "TET",
    "company": "TETRA4 (PTY) LTD",
    "city": "Johannesburg",
    "country": "South Africa"
  },
  "TEX": {
    "code": "TEX",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "TFC": {
    "code": "TFC",
    "company": "FAGIOLI SPA",
    "city": "REGGIO EMILIA",
    "country": "Italy"
  },
  "TFE": {
    "code": "TFE",
    "company": "TRANSFENNICA LTD",
    "city": "HELSINKI",
    "country": "Finland"
  },
  "TFG": {
    "code": "TFG",
    "company": "DB CARGO AG",
    "city": "MAINZ",
    "country": "Germany"
  },
  "TFH": {
    "code": "TFH",
    "company": "DHL FLEET SUPPORT ENGINEER",
    "city": "Widnes (Cheshire)",
    "country": "United Kingdom"
  },
  "TFM": {
    "code": "TFM",
    "company": "TFM",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "TFR": {
    "code": "TFR",
    "company": "CONTENEDORES TEIFER",
    "city": "LLEIDA",
    "country": "Spain"
  },
  "TFS": {
    "code": "TFS",
    "company": "INTERFERRY",
    "city": "Burlacha Balka, Chornomorsk",
    "country": "Ukraine"
  },
  "TFT": {
    "code": "TFT",
    "company": "TANKFORMATOR (S) PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "TGA": {
    "code": "TGA",
    "company": "TECHGLASS SP. ZOO.",
    "city": "KRAKOW",
    "country": "Poland"
  },
  "TGB": {
    "code": "TGB",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "TGC": {
    "code": "TGC",
    "company": "TOUAX",
    "city": "LA DEFENSE",
    "country": "France"
  },
  "TGE": {
    "code": "TGE",
    "company": "TAIWAN SPECIAL CHEMICALS CORPORATION",
    "city": "CHANGHUA COUNTY",
    "country": "Taiwan, China"
  },
  "TGG": {
    "code": "TGG",
    "company": "TIGER GAS (SHANGHAI) LIMITED",
    "city": "SHANGHAI",
    "country": "China"
  },
  "TGH": {
    "code": "TGH",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "TGI": {
    "code": "TGI",
    "company": "TAEWOONG LOGISTICS CO.LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "TGL": {
    "code": "TGL",
    "company": "TLS LOJISTIK A.S",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "TGP": {
    "code": "TGP",
    "company": "GENSETPOOL",
    "city": "Woodbridge",
    "country": "United States"
  },
  "TGR": {
    "code": "TGR",
    "company": "SPECTRANSGARANT LLC",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "TGS": {
    "code": "TGS",
    "company": "TOUAX CONTAINER LEASING PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "TGT": {
    "code": "TGT",
    "company": "NINGBO TRANSOCEAN GLOBAL TRANSPORTATION CORP.LTD",
    "city": "NINGBO",
    "country": "China"
  },
  "TGV": {
    "code": "TGV",
    "company": "LOGIX CAPITAL LLC",
    "city": "DORAL",
    "country": "United States"
  },
  "THA": {
    "code": "THA",
    "company": "CHONGQING TONGHUI GAS CO., LTD.",
    "city": "Chongqing",
    "country": "China"
  },
  "THB": {
    "code": "THB",
    "company": "CONT-ASPHALT LTD.",
    "city": "LA CROIX (LUTRY)",
    "country": "Switzerland"
  },
  "THE": {
    "code": "THE",
    "company": "THERMOCAR SRL  ITALIA",
    "city": "GENOVA  GE",
    "country": "Italy"
  },
  "THG": {
    "code": "THG",
    "company": "BEIJING TIANHAI CRYOGENIC EQUIPMENT CO. LTD.",
    "city": "Beijing",
    "country": "China"
  },
  "THI": {
    "code": "THI",
    "company": "THIELMANN FINANCIAL SOLUTIONS AG",
    "city": "Zug",
    "country": "Switzerland"
  },
  "THJ": {
    "code": "THJ",
    "company": "SHENZHEN TEHUAJIAN ADAMANTITE STRUCTURE MANUFACTURING CO., LTD.",
    "city": "Shenzhen",
    "country": "China"
  },
  "THK": {
    "code": "THK",
    "company": "TANK-CONTAINER PETROCHEMICAL COMPANY LTD",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "THP": {
    "code": "THP",
    "company": "ITALMATCH CHEMICALS S.P.A.",
    "city": "GENOVA (GE)",
    "country": "Italy"
  },
  "THQ": {
    "code": "THQ",
    "company": "THOR QUIMICOS DE MEXICO, SA DE CV",
    "city": "QUERETARO",
    "country": "Mexico"
  },
  "THS": {
    "code": "THS",
    "company": "TRADING HOUSE \"SINTEZ-OIL\", LTD",
    "city": "Dzerzhinsk",
    "country": "Russian Federation"
  },
  "THY": {
    "code": "THY",
    "company": "THIJSSEN DRILLING COMPANY BV",
    "city": "GEULLE",
    "country": "Netherlands"
  },
  "TIB": {
    "code": "TIB",
    "company": "TACON GMBH TANK & CONTAINER SERVICES",
    "city": "BRAAK",
    "country": "Germany"
  },
  "TIC": {
    "code": "TIC",
    "company": "LIVERANI GROUP SPA",
    "city": "SANTA MARIA DI ZEVIO (VR)",
    "country": "Italy"
  },
  "TIF": {
    "code": "TIF",
    "company": "TRIFLEET LEASING HOLDING B.V.",
    "city": "DORDRECHT",
    "country": "Netherlands"
  },
  "TIH": {
    "code": "TIH",
    "company": "TRADECORP INTERNATIONAL HONG KONG LIMITED",
    "city": "HONG KONG",
    "country": "HK"
  },
  "TII": {
    "code": "TII",
    "company": "TRITON INTERNATIONAL LTD",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "TIL": {
    "code": "TIL",
    "company": "TEC CONTAINER SOLUTIONS",
    "city": "WORTHING, WEST SUSSEX",
    "country": "United Kingdom"
  },
  "TIM": {
    "code": "TIM",
    "company": "ALPSANNA LIMITED",
    "city": "Dublin",
    "country": "Ireland"
  },
  "TIO": {
    "code": "TIO",
    "company": "ONET TECHNOLOGIES TI",
    "city": "MARSEILLE CEDEX 09",
    "country": "France"
  },
  "TIP": {
    "code": "TIP",
    "company": "TIP TRAILER SERVICES MANAGEMENT B.V.",
    "city": "Amsterdam",
    "country": "Netherlands"
  },
  "TIS": {
    "code": "TIS",
    "company": "TRANSPORT INDUSTRIAL SERVICE (TIS) LTD",
    "city": "ST PETERSBURG",
    "country": "Russian Federation"
  },
  "TIT": {
    "code": "TIT",
    "company": "TITAN CONTAINERS A/S",
    "city": "TAASTRUP",
    "country": "Denmark"
  },
  "TIU": {
    "code": "TIU",
    "company": "TACQUET INDUSTRIES",
    "city": "CARVIN",
    "country": "France"
  },
  "TIV": {
    "code": "TIV",
    "company": "TRANSINSULAR CABO VERDE",
    "city": "PRAIA",
    "country": "Cape Verde"
  },
  "TJZ": {
    "code": "TJZ",
    "company": "HANGZHOU TIEJI FREIGHT CO., LTD",
    "city": "HANGZHOU",
    "country": "China"
  },
  "TKC": {
    "code": "TKC",
    "company": "CONSERT S.R.L.",
    "city": "TRUCCAZZANO",
    "country": "Italy"
  },
  "TKI": {
    "code": "TKI",
    "company": "TOM KRAEMER, INC.",
    "city": "Cold Spring, MN 56320",
    "country": "United States"
  },
  "TKK": {
    "code": "TKK",
    "company": "TRUKAI INDUSTRIES LTD",
    "city": "Lae",
    "country": "Papua New Guinea"
  },
  "TKL": {
    "code": "TKL",
    "company": "SHANGHAI YUXIN CONTAINER LEASING CO LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "TKM": {
    "code": "TKM",
    "company": "TANKMASTER LTD",
    "city": "Novosaratovka",
    "country": "Russian Federation"
  },
  "TKN": {
    "code": "TKN",
    "company": "THERMO KING TRANSPORT KOELING B.V.",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "TKR": {
    "code": "TKR",
    "company": "PJSC TRANSCONTAINER",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "TKS": {
    "code": "TKS",
    "company": "TERMOCON LTD",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "TLC": {
    "code": "TLC",
    "company": "UNIFEEDER A/S",
    "city": "GDYNIA",
    "country": "Poland"
  },
  "TLE": {
    "code": "TLE",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "TLF": {
    "code": "TLF",
    "company": "NUROL TEKNOLOJI SAN. VE MAD. TIC. A.S.",
    "city": "ANKARA",
    "country": "Turkey"
  },
  "TLI": {
    "code": "TLI",
    "company": "TRANSPORT LOGISTICS INTERNATIONAL, INC",
    "city": "FULTON, MD 20759",
    "country": "United States"
  },
  "TLL": {
    "code": "TLL",
    "company": "TAL INTERNATIONAL",
    "city": "PURCHASE, NY 10577-2135",
    "country": "United States"
  },
  "TLN": {
    "code": "TLN",
    "company": "TRIDENT CONTAINER LEASING BV",
    "city": "AMSTERDAM",
    "country": "Netherlands"
  },
  "TLO": {
    "code": "TLO",
    "company": "TANKLOG GMBH & CO KG",
    "city": "MANNHEIM",
    "country": "Germany"
  },
  "TLS": {
    "code": "TLS",
    "company": "LLC \"TRANSPORT LOGISTICS SYSTEMS\"",
    "city": "Novosibirsk",
    "country": "Russian Federation"
  },
  "TLT": {
    "code": "TLT",
    "company": "TRADEMARK LEASING & TRADING BV",
    "city": "HOOGVLIET (RI)",
    "country": "Netherlands"
  },
  "TLV": {
    "code": "TLV",
    "company": "LLC «TIS-LOGISTIC»",
    "city": "Vladivostok",
    "country": "Russian Federation"
  },
  "TLX": {
    "code": "TLX",
    "company": "TRANS ASIAN SHIPPING SERVICES (P) LTD.",
    "city": "COCHIN",
    "country": "India"
  },
  "TMB": {
    "code": "TMB",
    "company": "TAMBOUR LTD.",
    "city": "AKKO",
    "country": "Israel"
  },
  "TMC": {
    "code": "TMC",
    "company": "TRANSMAR SHIPPING COMPANY",
    "city": "CAIRO",
    "country": "Egypt"
  },
  "TMD": {
    "code": "TMD",
    "company": "CARIBBEAN GAS CHEMICAL LIMITED (CGCL)",
    "city": "Port of Spain",
    "country": "Trinidad and Tobago"
  },
  "TME": {
    "code": "TME",
    "company": "TOMOESHOKAI CO; LTD",
    "city": "TOKYO",
    "country": "Japan"
  },
  "TMH": {
    "code": "TMH",
    "company": "TRIMODAL EUROPE B.V",
    "city": "HOOGVLIET / RT",
    "country": "Netherlands"
  },
  "TMI": {
    "code": "TMI",
    "company": "TRIFLEET LEASING HOLDING B.V.",
    "city": "DORDRECHT",
    "country": "Netherlands"
  },
  "TMK": {
    "code": "TMK",
    "company": "ES AT-ABRAY",
    "city": "Ashgabat",
    "country": "Turkmenistan"
  },
  "TML": {
    "code": "TML",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "TMM": {
    "code": "TMM",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "TMP": {
    "code": "TMP",
    "company": "T.M.P TRANS MODAL",
    "city": "NEUILLY S/SEINE",
    "country": "France"
  },
  "TMT": {
    "code": "TMT",
    "company": "GCATAINER BV",
    "city": "MOERDIJK",
    "country": "Netherlands"
  },
  "TMU": {
    "code": "TMU",
    "company": "FUH TOMASZ PROKORYM",
    "city": "Olsztyn",
    "country": "Poland"
  },
  "TMW": {
    "code": "TMW",
    "company": "VERBIO LOGISTIK GMBH",
    "city": "SCHWEDT",
    "country": "Germany"
  },
  "TMY": {
    "code": "TMY",
    "company": "TRANSINSULAR",
    "city": "LISBOA",
    "country": "Portugal"
  },
  "TNA": {
    "code": "TNA",
    "company": "TIONG NAM LOGISTICS SOLUTIONS SDN BHD",
    "city": "SHAH ALAM",
    "country": "Malaysia"
  },
  "TNC": {
    "code": "TNC",
    "company": "TAIYO NIPPON SANSO CORPORATION",
    "city": "TOKYO",
    "country": "Japan"
  },
  "TND": {
    "code": "TND",
    "company": "TCI SUPPLY CHAIN SOLUTIONS",
    "city": "GURGAON, HARYANA",
    "country": "India"
  },
  "TNF": {
    "code": "TNF",
    "company": "TANFAC INDUSTRIES LTD",
    "city": "CUDDALORE / TAMILNADU",
    "country": "India"
  },
  "TNI": {
    "code": "TNI",
    "company": "TN AMERICAS LLC",
    "city": "COLUMBIA, MD-21045",
    "country": "United States"
  },
  "TNK": {
    "code": "TNK",
    "company": "TANKONE B.V.",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "TNM": {
    "code": "TNM",
    "company": "TRANSPORTATION COMPANY TRANS MERIDIAN",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "TNO": {
    "code": "TNO",
    "company": "TANCO INTERNATIONAL (97) LTD.",
    "city": "HAIFA",
    "country": "Israel"
  },
  "TNP": {
    "code": "TNP",
    "company": "ORANO TN INTERNATIONAL",
    "city": "VALOGNES",
    "country": "France"
  },
  "TNR": {
    "code": "TNR",
    "company": "TECHNOAZOT R&P LTD",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "TNS": {
    "code": "TNS",
    "company": "UNICO LOGISTICS CO LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "TNZ": {
    "code": "TNZ",
    "company": "QUARTZ ENTERPRISES LTD",
    "city": "BELIZE CITY",
    "country": "Belize"
  },
  "TOA": {
    "code": "TOA",
    "company": "TOAGOSEI CO., LTD.",
    "city": "TOKYO",
    "country": "Japan"
  },
  "TOC": {
    "code": "TOC",
    "company": "TAIPEI OXYGEN & GAS CO.",
    "city": "New Taipei City",
    "country": "Taiwan, China"
  },
  "TOE": {
    "code": "TOE",
    "company": "TRANSOCEAN EQUIPMENT MANAGEMENT LLC",
    "city": "Little Rock",
    "country": "United States"
  },
  "TOI": {
    "code": "TOI",
    "company": "GRANDS TRAVAUX DE LOCEAN INDIEN (GTOI)",
    "city": "LE PORT - LA REUNION",
    "country": "France"
  },
  "TOL": {
    "code": "TOL",
    "company": "TAL INTERNATIONAL",
    "city": "PURCHASE, NY 10577-2135",
    "country": "United States"
  },
  "TON": {
    "code": "TON",
    "company": "TANK ONE NV",
    "city": "Antwerpen",
    "country": "Belgium"
  },
  "TOP": {
    "code": "TOP",
    "company": "TOPTAINER CONTAINERMANAGEMENT & SALES GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "TOR": {
    "code": "TOR",
    "company": "MAERSK A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "TOT": {
    "code": "TOT",
    "company": "SERVIZI LOGISTICI INTEGRATI S.R.L",
    "city": "SORA (FR)",
    "country": "Italy"
  },
  "TPB": {
    "code": "TPB",
    "company": "TIGER POWER",
    "city": "Heusden",
    "country": "Belgium"
  },
  "TPC": {
    "code": "TPC",
    "company": "WAN HAI LINES LTD",
    "city": "TAIPEI",
    "country": "Taiwan, China"
  },
  "TPF": {
    "code": "TPF",
    "company": "TPCO DEMOLITION",
    "city": "Saint-Maurice-la-Clouère",
    "country": "France"
  },
  "TPH": {
    "code": "TPH",
    "company": "TAL INTERNATIONAL",
    "city": "PURCHASE, NY 10577-2135",
    "country": "United States"
  },
  "TPK": {
    "code": "TPK",
    "company": "TRANSPEK INDUSTRY LIMITED",
    "city": "VADODARA",
    "country": "India"
  },
  "TPM": {
    "code": "TPM",
    "company": "EXSIF WORLDWIDE",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "TPO": {
    "code": "TPO",
    "company": "TRANS-PORTS S.A",
    "city": "VALENCIA",
    "country": "Spain"
  },
  "TPS": {
    "code": "TPS",
    "company": "BASF TAIWAN LIMITED",
    "city": "TAIPEI",
    "country": "Taiwan, China"
  },
  "TPT": {
    "code": "TPT",
    "company": "EXSIF WORLDWIDE",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "TPX": {
    "code": "TPX",
    "company": "TAL INTERNATIONAL",
    "city": "PURCHASE, NY 10577-2135",
    "country": "United States"
  },
  "TQM": {
    "code": "TQM",
    "company": "INTERMODAL TANK TRANSPORT INC.",
    "city": "HOUSTON,TX 77064",
    "country": "United States"
  },
  "TQR": {
    "code": "TQR",
    "company": "TRANSPORTES QUIMICOS RAMIREZ SL",
    "city": "SANLUCAR LA MAYOR-SEVILLA",
    "country": "Spain"
  },
  "TRD": {
    "code": "TRD",
    "company": "TAL INTERNATIONAL",
    "city": "PURCHASE, NY 10577-2135",
    "country": "United States"
  },
  "TRE": {
    "code": "TRE",
    "company": "TRIEST AG GROUP INC",
    "city": "TIFTON",
    "country": "United States"
  },
  "TRG": {
    "code": "TRG",
    "company": "TRIGAS EXPORT CA",
    "city": "CARACAS",
    "country": "Venezuela, Bolivarian Republic of"
  },
  "TRH": {
    "code": "TRH",
    "company": "TRITON INTERNATIONAL LTD",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "TRI": {
    "code": "TRI",
    "company": "TRITON INTERNATIONAL LTD",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "TRK": {
    "code": "TRK",
    "company": "TURKON CONTAINER TRANSPORTATION&SHIPPING",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "TRL": {
    "code": "TRL",
    "company": "TAL INTERNATIONAL",
    "city": "PURCHASE, NY 10577-2135",
    "country": "United States"
  },
  "TRM": {
    "code": "TRM",
    "company": "TARIMSAL KIMYA A.S.",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "TRN": {
    "code": "TRN",
    "company": "TRANSNET SA",
    "city": "PIRAEUS",
    "country": "Greece"
  },
  "TRO": {
    "code": "TRO",
    "company": "FELIX TROLL TRANSPORT GMBH",
    "city": "IMST",
    "country": "Austria"
  },
  "TRP": {
    "code": "TRP",
    "company": "TROPICAL TRAILER LEASING LLC",
    "city": "MIAMI, FL-33178",
    "country": "United States"
  },
  "TRR": {
    "code": "TRR",
    "company": "TRANSPORT RESOURCES INC.",
    "city": "MATAWAN, NJ 07747",
    "country": "United States"
  },
  "TRS": {
    "code": "TRS",
    "company": "FEDERTRASPORTI SPA",
    "city": "CASTELMAGGIORE (BO)",
    "country": "Italy"
  },
  "TRT": {
    "code": "TRT",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "TRU": {
    "code": "TRU",
    "company": "TRISTAR ENGINNERING CONSULTING LOGISTIC SA",
    "city": "CHIASSO",
    "country": "Switzerland"
  },
  "TRV": {
    "code": "TRV",
    "company": "TRITON INTERNATIONAL LTD",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "TSA": {
    "code": "TSA",
    "company": "TRISTAR CONTAINER SERVICES (ASIA) PVT. LTD.",
    "city": "CHENNAI",
    "country": "India"
  },
  "TSC": {
    "code": "TSC",
    "company": "TRANS-SERVICE-1 LLC",
    "city": "Lviv",
    "country": "Ukraine"
  },
  "TSE": {
    "code": "TSE",
    "company": "TOP SPEED ENERGY OVERSEAS CORP LIMITED",
    "city": "WANCHAI",
    "country": "HK"
  },
  "TSF": {
    "code": "TSF",
    "company": "SOULBRAIN CORP",
    "city": "GYEONGGI-DO",
    "country": "Korea, Republic of"
  },
  "TSG": {
    "code": "TSG",
    "company": "VENTURA TRADING LTD",
    "city": "HAMILTON, HM",
    "country": "Bermuda"
  },
  "TSH": {
    "code": "TSH",
    "company": "TSB Brasil Ltda – ME",
    "city": "Curitiba",
    "country": "Brazil"
  },
  "TSI": {
    "code": "TSI",
    "company": "STSI",
    "city": "GONESSE",
    "country": "France"
  },
  "TSK": {
    "code": "TSK",
    "company": "ESK SA",
    "city": "ALBUIXECH VALENCIA",
    "country": "Spain"
  },
  "TSL": {
    "code": "TSL",
    "company": "TCI SEAWAYS A DIVISION OF TRANSPORT CORPORATION OF INDIA LTD",
    "city": "CHENNAI",
    "country": "India"
  },
  "TSP": {
    "code": "TSP",
    "company": "TIPES S.P.A",
    "city": "OLGIATE MOLGORA",
    "country": "Italy"
  },
  "TSS": {
    "code": "TSS",
    "company": "T.S. LINES LTD",
    "city": "TAIPEI",
    "country": "Taiwan, China"
  },
  "TST": {
    "code": "TST",
    "company": "T.S. LINES LTD",
    "city": "TAIPEI",
    "country": "Taiwan, China"
  },
  "TSV": {
    "code": "TSV",
    "company": "SHIPPING COMPANY TRANSIT SV LLC",
    "city": "KRASNOYARSK",
    "country": "Russian Federation"
  },
  "TTA": {
    "code": "TTA",
    "company": "TRET SRL",
    "city": "Sesto Fiorentino",
    "country": "Italy"
  },
  "TTB": {
    "code": "TTB",
    "company": "EQUIPPED4U BV",
    "city": "Roosendaal",
    "country": "Netherlands"
  },
  "TTE": {
    "code": "TTE",
    "company": "TEAMTEC AS",
    "city": "TVEDESTRAND",
    "country": "Norway"
  },
  "TTI": {
    "code": "TTI",
    "company": "MASTERS GLOBAL LOGISTICS, INC.",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "TTK": {
    "code": "TTK",
    "company": "TRISTAR TRANSPORT LLC",
    "city": "Dubai",
    "country": "United Arab Emirates"
  },
  "TTM": {
    "code": "TTM",
    "company": "TITAN MED SARL",
    "city": "Monaco",
    "country": "Monaco"
  },
  "TTN": {
    "code": "TTN",
    "company": "TRITON INTERNATIONAL LTD",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "TTO": {
    "code": "TTO",
    "company": "TOTAL CONTAINERS PTY LTD",
    "city": "Wattleup",
    "country": "Australia"
  },
  "TTS": {
    "code": "TTS",
    "company": "LLC TRADETRANS",
    "city": "Novosibirsk",
    "country": "Russian Federation"
  },
  "TTT": {
    "code": "TTT",
    "company": "EUROTAINER SA",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "TUC": {
    "code": "TUC",
    "company": "TUCABI CONTAINER S.L.",
    "city": "ERANDIO",
    "country": "Spain"
  },
  "TUF": {
    "code": "TUF",
    "company": "TRANSUNIVERSE FORWARDING NV",
    "city": "Wondelgem",
    "country": "Belgium"
  },
  "TUK": {
    "code": "TUK",
    "company": "THALES UK",
    "city": "Templecombe, Somerset",
    "country": "United Kingdom"
  },
  "TUL": {
    "code": "TUL",
    "company": "CONTAINER SOLUTIONS S.A DE C.V",
    "city": "MEXICO",
    "country": "Mexico"
  },
  "TUT": {
    "code": "TUT",
    "company": "SHIPPING CONTAINER CONSULTANTS LTD",
    "city": "Hitcham Suffolk",
    "country": "United Kingdom"
  },
  "TVL": {
    "code": "TVL",
    "company": "JERMYN CORPORATION - TAVRIA LINE",
    "city": "DNIEPROPETROVSK",
    "country": "Ukraine"
  },
  "TVO": {
    "code": "TVO",
    "company": "TRANSPORT VAN OVERVELD BV",
    "city": "Etten Leur",
    "country": "Netherlands"
  },
  "TVR": {
    "code": "TVR",
    "company": "LLC TRANSSERVICE",
    "city": "Vologda",
    "country": "Russian Federation"
  },
  "TVT": {
    "code": "TVT",
    "company": "TESVOLT GMBH",
    "city": "Lutherstadt Wittenberg",
    "country": "Germany"
  },
  "TWA": {
    "code": "TWA",
    "company": "TAYLOR-WHARTON AMERICA INC",
    "city": "BAYTOWN,TX",
    "country": "United States"
  },
  "TWC": {
    "code": "TWC",
    "company": "THRU-WORLD COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "TWM": {
    "code": "TWM",
    "company": "4TESS  SP. Z.O.O",
    "city": "SOPOT",
    "country": "Poland"
  },
  "TWO": {
    "code": "TWO",
    "company": "2GO GROUP INC.",
    "city": "Pasay City",
    "country": "Philippines"
  },
  "TWP": {
    "code": "TWP",
    "company": "TROPICAL WESTERN PACIFIC PROGRAM OFFICE",
    "city": "LOS ALAMOS, NM 87545",
    "country": "United States"
  },
  "TWR": {
    "code": "TWR",
    "company": "TANKWAY BV",
    "city": "Zevenbergen",
    "country": "Netherlands"
  },
  "TWS": {
    "code": "TWS",
    "company": "TWS TANKCONTAINER-LEASING GMBH & CO KG",
    "city": "Hamburg",
    "country": "Germany"
  },
  "TWW": {
    "code": "TWW",
    "company": "TURKMEN AK YOL",
    "city": "Ashgabat",
    "country": "Turkmenistan"
  },
  "TXG": {
    "code": "TXG",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "TXW": {
    "code": "TXW",
    "company": "SHANGHAI ACE INVESTMENT & DEVELOPMENT CO., LTD.",
    "city": "Shanghai",
    "country": "China"
  },
  "TXX": {
    "code": "TXX",
    "company": "TRIDENT SEAFOODS CORP",
    "city": "SEATTLE, WA 98107-4000",
    "country": "United States"
  },
  "TYC": {
    "code": "TYC",
    "company": "TOKUYAMA CHEMICALS (ZHEJIANG) CO.,LTD",
    "city": "Jiaxing",
    "country": "China"
  },
  "TYR": {
    "code": "TYR",
    "company": "CRS REFRIGERATION LTD",
    "city": "MEATH",
    "country": "Ireland"
  },
  "TYS": {
    "code": "TYS",
    "company": "MITSUBISHI CHEMICAL TAIWAN Co., LTD",
    "city": "HSIN-CHU",
    "country": "Taiwan, China"
  },
  "TZK": {
    "code": "TZK",
    "company": "SUMISEI TAIWAN TECHNOLOGY CO LTD",
    "city": "CHANGHUA COUNTY",
    "country": "Taiwan, China"
  },
  "TZR": {
    "code": "TZR",
    "company": "TAZZETTI S.P.A.",
    "city": "Volpiano",
    "country": "Italy"
  },
  "UAA": {
    "code": "UAA",
    "company": "MOBIUS-GROUPPE EOOD",
    "city": "Sofia",
    "country": "Bulgaria"
  },
  "UAC": {
    "code": "UAC",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "UAE": {
    "code": "UAE",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "UAF": {
    "code": "UAF",
    "company": "UAFL UNITED AFRICA FEEDER LINE",
    "city": "MAPOU",
    "country": "Mauritius"
  },
  "UAK": {
    "code": "UAK",
    "company": "ALARA LOGISTICS",
    "city": "Humble",
    "country": "United States"
  },
  "UAS": {
    "code": "UAS",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "UBB": {
    "code": "UBB",
    "company": "DEN HARTOGH DRY BULK LOGISTICS LTD",
    "city": "Hull, HU3 4AE",
    "country": "United Kingdom"
  },
  "UBT": {
    "code": "UBT",
    "company": "INTERNATIONAL FREIGHT FORWARDING CENTRE OF ULAANBAATAR",
    "city": "Ulaanbaatar",
    "country": "Mongolia"
  },
  "UCC": {
    "code": "UCC",
    "company": "LLC UNIVERSAL CONTAINER COMPANY 1520",
    "city": "St. Petersburg",
    "country": "Russian Federation"
  },
  "UCL": {
    "code": "UCL",
    "company": "OCEAN AFRICA CONTAINER LINE (PTY) LTD",
    "city": "DURBAN",
    "country": "South Africa"
  },
  "UCS": {
    "code": "UCS",
    "company": "UNIVERSAL CONTAINER SERVICES LIMITED",
    "city": "IRLAM, MANCHESTER",
    "country": "United Kingdom"
  },
  "UEN": {
    "code": "UEN",
    "company": "U-EN CORPORATION",
    "city": "CHUO-KU, KOBE",
    "country": "Japan"
  },
  "UES": {
    "code": "UES",
    "company": "UES INTERNATIONAL (HK)  HOLDINGS LIMITED",
    "city": "SHANGHAI",
    "country": "China"
  },
  "UET": {
    "code": "UET",
    "company": "UES INTERNATIONAL (HK)  HOLDINGS LIMITED",
    "city": "SHANGHAI",
    "country": "China"
  },
  "UFC": {
    "code": "UFC",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "UGC": {
    "code": "UGC",
    "company": "UNIPER GLOBAL COMMODITIES SE",
    "city": "Düsseldorf",
    "country": "Germany"
  },
  "UGM": {
    "code": "UGM",
    "company": "GREENCOMPASS MARINE S.A.",
    "city": "",
    "country": "Panama"
  },
  "UGN": {
    "code": "UGN",
    "company": "UMDASCH GROUP NEWCON",
    "city": "Amstetten",
    "country": "Austria"
  },
  "UGP": {
    "code": "UGP",
    "company": "ROZA MIRA LLC",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "UHM": {
    "code": "UHM",
    "company": "HIMMAGISTRAL LTD",
    "city": "UGREN",
    "country": "Russian Federation"
  },
  "UKR": {
    "code": "UKR",
    "company": "SETTI",
    "city": "SYNYAVA",
    "country": "Ukraine"
  },
  "ULC": {
    "code": "ULC",
    "company": "UNICO LOGISTICS CO LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "ULT": {
    "code": "ULT",
    "company": "ACE ENGINEERING CO LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "UMJ": {
    "code": "UMJ",
    "company": "AS CONTAINERHANDEL & SERVICE GMBH",
    "city": "Berne",
    "country": "Germany"
  },
  "UML": {
    "code": "UML",
    "company": "SMART BULK TERMINAL LIMITED LIABILITY COMPANY",
    "city": "",
    "country": "Russian Federation"
  },
  "UMT": {
    "code": "UMT",
    "company": "ULTRAMAR LTD",
    "city": "RIGA",
    "country": "Latvia"
  },
  "UNC": {
    "code": "UNC",
    "company": "UNI-LOGISTICS SP. Z O.O.",
    "city": "GDYNIA",
    "country": "Poland"
  },
  "UND": {
    "code": "UND",
    "company": "UNITAINER TRADING GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "UNI": {
    "code": "UNI",
    "company": "UNITAS CONTAINERS LIMITED",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "UNL": {
    "code": "UNL",
    "company": "NAMMA SHIPPING LINES (NSL)",
    "city": "ALEXANDRIA",
    "country": "Egypt"
  },
  "UNO": {
    "code": "UNO",
    "company": "MALTHUS UNITEAM AS",
    "city": "SKEDSMOKORSET",
    "country": "Norway"
  },
  "UNX": {
    "code": "UNX",
    "company": "UNITAS CONTAINERS LIMITED",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "UPA": {
    "code": "UPA",
    "company": "KMG ULTRA PURE CHEMICALS LIMITED",
    "city": "ALFRETON DE55 4DA",
    "country": "United Kingdom"
  },
  "UPL": {
    "code": "UPL",
    "company": "UPLINE CZ, S.R.O.",
    "city": "JINOCANY",
    "country": "Czech Republic"
  },
  "URE": {
    "code": "URE",
    "company": "URENCO LTD",
    "city": "BAD BENTHEIN",
    "country": "Germany"
  },
  "URF": {
    "code": "URF",
    "company": "EUROTRANS LOGISTICS LLC",
    "city": "Yekaterinburg",
    "country": "Russian Federation"
  },
  "USA": {
    "code": "USA",
    "company": "UNITED STATES ARMY",
    "city": "SCOTT AFB, IL-62225-5006",
    "country": "United States"
  },
  "USF": {
    "code": "USF",
    "company": "UNITED STATES AIR FORCE",
    "city": "WRIGHT-PATTERSON AFB, OH-45433",
    "country": "United States"
  },
  "USG": {
    "code": "USG",
    "company": "MILITARY SDDC",
    "city": "Scott AFB",
    "country": "United States"
  },
  "USM": {
    "code": "USM",
    "company": "UNITED STATES MARINE CORPS",
    "city": "ALBANY, GA-31704",
    "country": "United States"
  },
  "USN": {
    "code": "USN",
    "company": "UNITED STATES NAVY",
    "city": "NORFOLK, VA-23511-3492",
    "country": "United States"
  },
  "USP": {
    "code": "USP",
    "company": "EXSIF WORLDWIDE",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "UTC": {
    "code": "UTC",
    "company": "STOLT TANK CONTAINERS LEASING LTD",
    "city": "HAMILTON",
    "country": "Bermuda"
  },
  "UTE": {
    "code": "UTE",
    "company": "UNIENERGY TECHNOLOGIES",
    "city": "MUKILTEO, WA-98275",
    "country": "United States"
  },
  "UTK": {
    "code": "UTK",
    "company": "UNITEK SOLVENT SERVICES, INC",
    "city": "KAPOLEI, HI 96707",
    "country": "United States"
  },
  "UTO": {
    "code": "UTO",
    "company": "EDF",
    "city": "MONTEVRAIN",
    "country": "France"
  },
  "UTT": {
    "code": "UTT",
    "company": "DEN HARTOGH GLOBAL BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "UUU": {
    "code": "UUU",
    "company": "ELBTAINER TRADING GMBH",
    "city": "Barsbüttel",
    "country": "Germany"
  },
  "UXO": {
    "code": "UXO",
    "company": "UXORIOUS IV",
    "city": "Oyster Bay",
    "country": "United States"
  },
  "UXX": {
    "code": "UXX",
    "company": "UES INTERNATIONAL (HK)  HOLDINGS LIMITED",
    "city": "SHANGHAI",
    "country": "China"
  },
  "UZU": {
    "code": "UZU",
    "company": "JSC UKRZALIZNYTSIA",
    "city": "KYIV",
    "country": "Ukraine"
  },
  "VAG": {
    "code": "VAG",
    "company": "VA GROUP LTD.",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "VAI": {
    "code": "VAI",
    "company": "VAISALA OYJ",
    "city": "Vantaa",
    "country": "Finland"
  },
  "VAN": {
    "code": "VAN",
    "company": "VANHUB EQUIPMENT SERVICES LTD",
    "city": "KOWLOON",
    "country": "HK"
  },
  "VAS": {
    "code": "VAS",
    "company": "VASI SHIPPING PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "VBK": {
    "code": "VBK",
    "company": "TRANSPORT VERBEKEN NV",
    "city": "DENDERMONDE",
    "country": "Belgium"
  },
  "VCC": {
    "code": "VCC",
    "company": "TC VANCONTAINER CO. LTD",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "VCY": {
    "code": "VCY",
    "company": "VÍTKOVICE CYLINDERS A.S.",
    "city": "Ostrava",
    "country": "Czech Republic"
  },
  "VDB": {
    "code": "VDB",
    "company": "VAN DEN BAN  AUTOBANDEN BV",
    "city": "HELLEVOETSLUIS",
    "country": "Netherlands"
  },
  "VDC": {
    "code": "VDC",
    "company": "VAN DOORN CONTAINER SALES AND RENTALS",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "VDE": {
    "code": "VDE",
    "company": "VAN DER ENT TOPMOVERS",
    "city": "SPIJKENISSE",
    "country": "Netherlands"
  },
  "VDL": {
    "code": "VDL",
    "company": "L.V.D. LEE EN ZONEN BV",
    "city": "DELFT",
    "country": "Netherlands"
  },
  "VDM": {
    "code": "VDM",
    "company": "SAMSKIP MULTIMODAL BV",
    "city": "GENEMUIDEN",
    "country": "Netherlands"
  },
  "VEG": {
    "code": "VEG",
    "company": "VERHOEK INTERNATIONAAL TRANSPORT BV",
    "city": "GENEMUIDEN",
    "country": "Netherlands"
  },
  "VER": {
    "code": "VER",
    "company": "VERSUM MATERIALS LLC",
    "city": "TEMPE",
    "country": "United States"
  },
  "VES": {
    "code": "VES",
    "company": "VEOLIA ENVIRONMENTAL SERVICES",
    "city": "Fermoy",
    "country": "Ireland"
  },
  "VEY": {
    "code": "VEY",
    "company": "ETS J VEYNAT SA",
    "city": "TRESSES",
    "country": "France"
  },
  "VEZ": {
    "code": "VEZ",
    "company": "BALTICA TRANS",
    "city": "ST PETERSBURG",
    "country": "Russian Federation"
  },
  "VGK": {
    "code": "VGK",
    "company": "JSC \"VAGON-K\"",
    "city": "Krasnoyarsk",
    "country": "Russian Federation"
  },
  "VGL": {
    "code": "VGL",
    "company": "VASCO GLOBAL MARITIME LLC",
    "city": "BUR DUBAI",
    "country": "United Arab Emirates"
  },
  "VGN": {
    "code": "VGN",
    "company": "VIRTUAL GAS NETWORK",
    "city": "Krugersdorp",
    "country": "South Africa"
  },
  "VHT": {
    "code": "VHT",
    "company": "L & T B.V.",
    "city": "'s-Hertogenbosch",
    "country": "Netherlands"
  },
  "VIA": {
    "code": "VIA",
    "company": "VIASEA SHIPPING AS",
    "city": "Moss",
    "country": "Norway"
  },
  "VIC": {
    "code": "VIC",
    "company": "CONTAINER SERVICES SOLENT LTD",
    "city": "SOUTHAMPTON SO30  2HE",
    "country": "United Kingdom"
  },
  "VIN": {
    "code": "VIN",
    "company": "VOERMAN INTERNATIONAL BV",
    "city": "DEN HAAG",
    "country": "Netherlands"
  },
  "VIR": {
    "code": "VIR",
    "company": "LOGISTICS TRADER LLC",
    "city": "Mobile, AL 36695",
    "country": "United States"
  },
  "VLG": {
    "code": "VLG",
    "company": "BROEKMAN LOGISTICS BELGIUM ANTWERP",
    "city": "Antwerpen",
    "country": "Belgium"
  },
  "VLL": {
    "code": "VLL",
    "company": "VL LOGISTIC JOINT STOCK COMPANY",
    "city": "PRIMORSKIY REGION",
    "country": "Russian Federation"
  },
  "VLN": {
    "code": "VLN",
    "company": "VELAN SAS",
    "city": "LYON CEDEX 07",
    "country": "France"
  },
  "VMG": {
    "code": "VMG",
    "company": "VAKARU MEDIENOS GRUPE",
    "city": "Klaipėda",
    "country": "Lithuania"
  },
  "VML": {
    "code": "VML",
    "company": "VASCO MARITIME PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "VMX": {
    "code": "VMX",
    "company": "GRUPO TEVIAN SA DE CV",
    "city": "GENERAL ESCOBEDO",
    "country": "Mexico"
  },
  "VOC": {
    "code": "VOC",
    "company": "BERTHOLD VOLLERS GMBH",
    "city": "BREMEN",
    "country": "Germany"
  },
  "VOD": {
    "code": "VOD",
    "company": "TANK MANAGEMENT A/S",
    "city": "OSLO",
    "country": "Norway"
  },
  "VOL": {
    "code": "VOL",
    "company": "VOLTA SHIPPING SERVICES L.L.C",
    "city": "Dubai",
    "country": "United Arab Emirates"
  },
  "VOR": {
    "code": "VOR",
    "company": "OCEAN RACE 1973 SLU",
    "city": "ALICANTE",
    "country": "Spain"
  },
  "VOT": {
    "code": "VOT",
    "company": "VOLTA SHIPPING SERVICES L.L.C",
    "city": "Dubai",
    "country": "United Arab Emirates"
  },
  "VPL": {
    "code": "VPL",
    "company": "VAN PLUS CORPORATION",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "VPT": {
    "code": "VPT",
    "company": "VPT KOMPRESSOREN GMBH",
    "city": "Remscheid",
    "country": "Germany"
  },
  "VRA": {
    "code": "VRA",
    "company": "NORDIC BULKERS AB",
    "city": "GOTHENBURG",
    "country": "Sweden"
  },
  "VRE": {
    "code": "VRE",
    "company": "VREDEVELD INTERMODAL",
    "city": "HOOGERSMILDE",
    "country": "Netherlands"
  },
  "VRS": {
    "code": "VRS",
    "company": "VLADREFTRANS CO, LTD",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "VRT": {
    "code": "VRT",
    "company": "SANOFI CHIMIE",
    "city": "ARRAS CEDEX 9",
    "country": "France"
  },
  "VSB": {
    "code": "VSB",
    "company": "V.S & B CONTAINERS LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "VSH": {
    "code": "VSH",
    "company": "VOLKER SCHILLING BAUSTOFFE U. TRANSPORTE",
    "city": "Harsefeld",
    "country": "Germany"
  },
  "VSL": {
    "code": "VSL",
    "company": "VSKY INT'L CORP",
    "city": "SHANGHAI",
    "country": "China"
  },
  "VST": {
    "code": "VST",
    "company": "V.S & B CONTAINERS LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "VTC": {
    "code": "VTC",
    "company": "VISBEEN B.V",
    "city": "NIEUWE TONGE",
    "country": "Netherlands"
  },
  "VTG": {
    "code": "VTG",
    "company": "VTG TANKTAINER GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "VTI": {
    "code": "VTI",
    "company": "MOVE INTERMODAL SRL",
    "city": "NOVARA",
    "country": "Italy"
  },
  "VTN": {
    "code": "VTN",
    "company": "LIMITED LIABILITY COMPANY «TORGUNION»",
    "city": "Vladivostok",
    "country": "Russian Federation"
  },
  "VTS": {
    "code": "VTS",
    "company": "L. VAN TIEL TRANSPORT B.V.",
    "city": "Schiedam",
    "country": "Netherlands"
  },
  "VTZ": {
    "code": "VTZ",
    "company": "RUBIS ENERGIE",
    "city": "PARIS LA DEFENCE CEDEX",
    "country": "France"
  },
  "VVI": {
    "code": "VVI",
    "company": "GASCONTAINER JSC",
    "city": "NOVOSIBIRSK",
    "country": "Russian Federation"
  },
  "VVK": {
    "code": "VVK",
    "company": "VERVAEKE GROUP",
    "city": "SPIERE-HELKYN",
    "country": "Belgium"
  },
  "VWS": {
    "code": "VWS",
    "company": "VESTAS NORTHERN EUROPE A/S",
    "city": "COPENHAGEN",
    "country": "Denmark"
  },
  "VZR": {
    "code": "VZR",
    "company": "TRASPORTI VECCHI ZIRONI SRL",
    "city": "REGGIO EMILIA",
    "country": "Italy"
  },
  "WAB": {
    "code": "WAB",
    "company": "HOYER GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "WAE": {
    "code": "WAE",
    "company": "WUHAN ASIA-EUROPE LOGISTICS CO.,LTD",
    "city": "Wuhan",
    "country": "China"
  },
  "WAF": {
    "code": "WAF",
    "company": "BOCS GMBH",
    "city": "Bremen",
    "country": "Germany"
  },
  "WAG": {
    "code": "WAG",
    "company": "WESTFALEN AG",
    "city": "MUENSTER",
    "country": "Germany"
  },
  "WAR": {
    "code": "WAR",
    "company": "WARTSILA FINLAND OY",
    "city": "VAASA",
    "country": "Finland"
  },
  "WAS": {
    "code": "WAS",
    "company": "WESTSTAR AVIATION SERVICES SDN BHD",
    "city": "AMPANG",
    "country": "Malaysia"
  },
  "WAU": {
    "code": "WAU",
    "company": "WAUTERS GLOBAL LOGISTICS NV",
    "city": "HAMME",
    "country": "Belgium"
  },
  "WAY": {
    "code": "WAY",
    "company": "SHANGHAI WAYSMOS FINE CHEMICAL CO.,LTD.",
    "city": "SHANGHAI",
    "country": "China"
  },
  "WBF": {
    "code": "WBF",
    "company": "WINE & BRANDY FACTORY ALLIANCE-1892",
    "city": "CHERNYAKHOVSK",
    "country": "Russian Federation"
  },
  "WBP": {
    "code": "WBP",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "WBR": {
    "code": "WBR",
    "company": "BRUHN TRANSPORT EQUIPMENT GMBH & CO",
    "city": "LUBECK",
    "country": "Germany"
  },
  "WBT": {
    "code": "WBT",
    "company": "WILBERT TRADING BV",
    "city": "MAASSLUIS",
    "country": "Netherlands"
  },
  "WBX": {
    "code": "WBX",
    "company": "WILLBOX LTD",
    "city": "Southampton Hants",
    "country": "United Kingdom"
  },
  "WCG": {
    "code": "WCG",
    "company": "WELDSHIP INDUSTRIES",
    "city": "BETHLEHEM, PA 18015",
    "country": "United States"
  },
  "WCH": {
    "code": "WCH",
    "company": "CW CONSULTING & TRADING E.U.",
    "city": "Krottendorf-Gaisfeld",
    "country": "Austria"
  },
  "WCI": {
    "code": "WCI",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "WCS": {
    "code": "WCS",
    "company": "WCS SANTANGELO GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "WCT": {
    "code": "WCT",
    "company": "WIENCONT CONTAINERTERMINAL GMBH",
    "city": "WIEN",
    "country": "Austria"
  },
  "WCX": {
    "code": "WCX",
    "company": "WEC LINES BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "WDF": {
    "code": "WDF",
    "company": "ALLNEX AUSTRIA GMBH",
    "city": "WERNDORF",
    "country": "Austria"
  },
  "WEB": {
    "code": "WEB",
    "company": "MOBILBOXX EUROPE GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "WEC": {
    "code": "WEC",
    "company": "WEC LINES BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "WED": {
    "code": "WED",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "WEK": {
    "code": "WEK",
    "company": "WEST-EXPRESS",
    "city": "Rokytne village",
    "country": "Ukraine"
  },
  "WEL": {
    "code": "WEL",
    "company": "WELCOR S.A.",
    "city": "Montevideo",
    "country": "Uruguay"
  },
  "WES": {
    "code": "WES",
    "company": "WESTERN GLOBAL HOLDINGS LIMITED",
    "city": "BRISTOL BS37 7LD",
    "country": "United Kingdom"
  },
  "WET": {
    "code": "WET",
    "company": "MOTOR CONTROLS INC.",
    "city": "Dallas",
    "country": "United States"
  },
  "WEW": {
    "code": "WEW",
    "company": "THIELMANN WEW GmbH",
    "city": "WEITEFELD",
    "country": "Germany"
  },
  "WEX": {
    "code": "WEX",
    "company": "WORLD EX LTD",
    "city": "Tbilisi",
    "country": "Georgia"
  },
  "WFH": {
    "code": "WFH",
    "company": "WATERFRONT CONTAINER LEASING CO INC.",
    "city": "SAN FRANCISCO, CA 94109",
    "country": "United States"
  },
  "WFT": {
    "code": "WFT",
    "company": "WIN FAST LINE LTD",
    "city": "Qingdao",
    "country": "China"
  },
  "WGC": {
    "code": "WGC",
    "company": "SPEDITION BAUMLE DORMAGEN GMBH",
    "city": "DORMAGEN",
    "country": "Germany"
  },
  "WGR": {
    "code": "WGR",
    "company": "WEIL GROUP RESOURCES, LLC",
    "city": "Richmond, VA 23220",
    "country": "United States"
  },
  "WGS": {
    "code": "WGS",
    "company": "WG SALARI BV",
    "city": "SITTARD",
    "country": "Netherlands"
  },
  "WHL": {
    "code": "WHL",
    "company": "WAN HAI LINES LTD",
    "city": "TAIPEI",
    "country": "Taiwan, China"
  },
  "WHS": {
    "code": "WHS",
    "company": "WAN HAI LINES LTD",
    "city": "TAIPEI",
    "country": "Taiwan, China"
  },
  "WIC": {
    "code": "WIC",
    "company": "WICAB CONTAINER AB",
    "city": "Gothenburg",
    "country": "Sweden"
  },
  "WID": {
    "code": "WID",
    "company": "WICKS B.V",
    "city": "BOLSWARD",
    "country": "Netherlands"
  },
  "WIK": {
    "code": "WIK",
    "company": "HONEYTAK INTERMODAL LIMITED",
    "city": "Yantai",
    "country": "China"
  },
  "WIP": {
    "code": "WIP",
    "company": "PT.INDONESIA WEDA BAY INDUSTRIAL PARK",
    "city": "JAKARTA SELATAN DKI JAKARTA",
    "country": "Indonesia"
  },
  "WIT": {
    "code": "WIT",
    "company": "LANXESS ORGANOMETALLICS GMBH",
    "city": "BERGKAMEN",
    "country": "Germany"
  },
  "WJS": {
    "code": "WJS",
    "company": "WOOJIN GLOBAL LOGISTICS",
    "city": "GANGSEO-GU,SEOUL",
    "country": "Korea, Republic of"
  },
  "WLG": {
    "code": "WLG",
    "company": "WAH LEE INDUSTRIAL CORP.",
    "city": "Kaohsiung",
    "country": "Taiwan, China"
  },
  "WLN": {
    "code": "WLN",
    "company": "WALLENIUS WILHELMSEN LOGISTICS AS",
    "city": "LYSAKER",
    "country": "Norway"
  },
  "WMG": {
    "code": "WMG",
    "company": "WAXING MOON GLOBAL LIMITED",
    "city": "TORTOLA",
    "country": "Virgin Islands, British"
  },
  "WML": {
    "code": "WML",
    "company": "WESTERMAN MULTIMODAL",
    "city": "Nieuwleusen",
    "country": "Netherlands"
  },
  "WND": {
    "code": "WND",
    "company": "WND (LIAO NING) HEAVY INDUSTRY CO LTD",
    "city": "Fushun",
    "country": "China"
  },
  "WNG": {
    "code": "WNG",
    "company": "WNG CONTAINER SERVICE CO LIMITED",
    "city": "Shanghai",
    "country": "China"
  },
  "WOL": {
    "code": "WOL",
    "company": "ANWOOD LOGISTICS SYSTEMS (SUZHOU) CO.,LTD",
    "city": "Suzhou China",
    "country": "China"
  },
  "WON": {
    "code": "WON",
    "company": "WONIK MATERIALS",
    "city": "CHEONGWON-GUN , CHUNGCHEONGBUK-DO",
    "country": "Korea, Republic of"
  },
  "WOS": {
    "code": "WOS",
    "company": "WEST OCEAN SHIPPING COMPANY LIMITED",
    "city": "CENTRAL",
    "country": "HK"
  },
  "WPS": {
    "code": "WPS",
    "company": "ALASKA STRUCTURES INC",
    "city": "LAS CRUCES",
    "country": "United States"
  },
  "WRX": {
    "code": "WRX",
    "company": "WEATHERHAVEN GLOBAL RESOURCES LTD",
    "city": "Coquitlum, BC V3K 6W5",
    "country": "Canada"
  },
  "WSC": {
    "code": "WSC",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "WSD": {
    "code": "WSD",
    "company": "WIDE SHINE DEVELOPMENT INC",
    "city": "HONG KONG",
    "country": "HK"
  },
  "WSG": {
    "code": "WSG",
    "company": "PNS UK LTD",
    "city": "Wakefield",
    "country": "United Kingdom"
  },
  "WSI": {
    "code": "WSI",
    "company": "INTRA DEFENSE TECHNOLOGIES",
    "city": "Riyadh",
    "country": "Saudi Arabia"
  },
  "WSL": {
    "code": "WSL",
    "company": "WESTWOOD SHIPPING LINES",
    "city": "PUYALLUP, WA 98374",
    "country": "United States"
  },
  "WSM": {
    "code": "WSM",
    "company": "WESTERN SALES & TESTING INC.",
    "city": "AMARILLO,  TX 79105",
    "country": "United States"
  },
  "WSN": {
    "code": "WSN",
    "company": "WILHELM STUHRENBERG GMBH & CO. KG",
    "city": "NORDENHAM",
    "country": "Germany"
  },
  "WSV": {
    "code": "WSV",
    "company": "NV TRUCK EN TANKCLEANING TACK NV",
    "city": "OOSTROZEBEKE",
    "country": "Belgium"
  },
  "WTH": {
    "code": "WTH",
    "company": "WICE LOGISTICS PCL",
    "city": "BANGKOK",
    "country": "Thailand"
  },
  "WTL": {
    "code": "WTL",
    "company": "WHITE LINE SHIPPING",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "WTS": {
    "code": "WTS",
    "company": "WEIGAND-GRUNDSTUCKS GMBH & CO. KG",
    "city": "LENGENBOSTEL",
    "country": "Germany"
  },
  "WWL": {
    "code": "WWL",
    "company": "WALLENIUS WILHELMSEN LOGISTICS AS",
    "city": "LYSAKER",
    "country": "Norway"
  },
  "WYY": {
    "code": "WYY",
    "company": "TEMC CO.,LTD.",
    "city": "Boeun-gun, Chungcheongbuk-do",
    "country": "Korea, Republic of"
  },
  "XAC": {
    "code": "XAC",
    "company": "ALPSCONTAINER",
    "city": "FREILASSING",
    "country": "Germany"
  },
  "XAG": {
    "code": "XAG",
    "company": "XI'AN INTERNATIONAL PORT MULTIMODAL TRANSPORTATION CO.,LTD",
    "city": "Xian",
    "country": "China"
  },
  "XAR": {
    "code": "XAR",
    "company": "OXYGEN & ARGON WORKS LTD",
    "city": "CAESAREA",
    "country": "Israel"
  },
  "XCT": {
    "code": "XCT",
    "company": "DIMC - DUBAI INTERNATIONAL MARINE CLUB",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "XFC": {
    "code": "XFC",
    "company": "HUBEI SINOPRHORUS ELECTRONICS MATERIALS CO LTD",
    "city": "HUBEI",
    "country": "China"
  },
  "XFV": {
    "code": "XFV",
    "company": "TRAWIND (NINGBO) MARINE LOGISTICS CO.,LTD.",
    "city": "Yingkou",
    "country": "China"
  },
  "XHC": {
    "code": "XHC",
    "company": "CXIC",
    "city": "CHANGZHOU JIANGSU",
    "country": "China"
  },
  "XIN": {
    "code": "XIN",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "XPA": {
    "code": "XPA",
    "company": "EXPANSION SAS",
    "city": "MERPINS",
    "country": "France"
  },
  "XPO": {
    "code": "XPO",
    "company": "XPO LOGISTICS",
    "city": "DUBLIN, OH 43016",
    "country": "United States"
  },
  "XRC": {
    "code": "XRC",
    "company": "XURONG CONTAINER SERVICE CO,LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "XTR": {
    "code": "XTR",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
  },
  "XXL": {
    "code": "XXL",
    "company": "2XL NV",
    "city": "ZEEBRUGGE",
    "country": "Belgium"
  },
  "XYL": {
    "code": "XYL",
    "company": "VIEWER DEVELOPMENT CO.LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "XYZ": {
    "code": "XYZ",
    "company": "\"AGRO CORN HOLDING\" LLC",
    "city": "Bilyaivka",
    "country": "Ukraine"
  },
  "YAK": {
    "code": "YAK",
    "company": "YAKUTSK CONTAINER AGENCY LLC",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "YDS": {
    "code": "YDS",
    "company": "HAINING INDUSAIR ELECTRONICS CO LTD",
    "city": "HAINING, ZHEJIANG",
    "country": "China"
  },
  "YDX": {
    "code": "YDX",
    "company": "PT INDONESIA GUANG CHING NICKEL AND STAINLESS STEEL IND",
    "city": "SOUTH JAKARTA",
    "country": "Indonesia"
  },
  "YEO": {
    "code": "YEO",
    "company": "YEOU FA CHEMICAL CO LTD",
    "city": "TAIPEI",
    "country": "Taiwan, China"
  },
  "YFC": {
    "code": "YFC",
    "company": "EVER FORTUNE LOGISTICS (HK) CO LTD",
    "city": "HUNG HOM-KOWLOON",
    "country": "HK"
  },
  "YGL": {
    "code": "YGL",
    "company": "YEMEN GULF LINE LTD",
    "city": "London",
    "country": "United Kingdom"
  },
  "YIM": {
    "code": "YIM",
    "company": "JIANGSU YIMA ROAD CONSTRUCTION MACHINERY TECHNOLOGY CO.,LTD.",
    "city": "jiangyin",
    "country": "China"
  },
  "YKT": {
    "code": "YKT",
    "company": "KITO INTERNATIONAL SHIPPING INCORPORATED",
    "city": "MAHE",
    "country": "Seychelles"
  },
  "YMG": {
    "code": "YMG",
    "company": "CENTRE DE RENOVATION MARCEL DAGORT",
    "city": "SAINT-PIERRE & MIQUELON",
    "country": "France"
  },
  "YML": {
    "code": "YML",
    "company": "YANG MING MARINE TRANSPORT CORP.",
    "city": "KEELUNG",
    "country": "Taiwan, China"
  },
  "YMM": {
    "code": "YMM",
    "company": "YANG MING MARINE TRANSPORT CORP.",
    "city": "KEELUNG",
    "country": "Taiwan, China"
  },
  "YNP": {
    "code": "YNP",
    "company": "JSC SAKHANEFTEGAZSBYT",
    "city": "Yakutsk, Republic of (Sakha) Yakutia",
    "country": "Russian Federation"
  },
  "YOI": {
    "code": "YOI",
    "company": "TRITON INTERNATIONAL LTD",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
  },
  "YOL": {
    "code": "YOL",
    "company": "ONIDA CONTAINER LEASING CO., LTD.",
    "city": "KOWLOON",
    "country": "HK"
  },
  "YPR": {
    "code": "YPR",
    "company": "SITRA",
    "city": "IEPER",
    "country": "Belgium"
  },
  "YSD": {
    "code": "YSD",
    "company": "YASUDA LOGISTICS CORPORATION",
    "city": "TOKYO",
    "country": "Japan"
  },
  "YTB": {
    "code": "YTB",
    "company": "STAR OIL MAURITANIE",
    "city": "NOUAKCHOTT",
    "country": "Mauritania"
  },
  "YXE": {
    "code": "YXE",
    "company": "YIWU TIMEX INDUSTRIAL INVESTMENT CO.,LTD",
    "city": "YIWU",
    "country": "China"
  },
  "YYC": {
    "code": "YYC",
    "company": "QUANZHOU YIYANG CONTAINER SERVICE CO.LTD",
    "city": "QUANZHOU CITY",
    "country": "China"
  },
  "ZAP": {
    "code": "ZAP",
    "company": "GRUPA ZAKLADY AZOTOWE PULAWY SA",
    "city": "PULAWY",
    "country": "Poland"
  },
  "ZCL": {
    "code": "ZCL",
    "company": "ZIM INTEGRATED SHIPPING SERVICES LTD",
    "city": "HAIFA",
    "country": "Israel"
  },
  "ZCS": {
    "code": "ZCS",
    "company": "ZIM INTEGRATED SHIPPING SERVICES LTD",
    "city": "HAIFA",
    "country": "Israel"
  },
  "ZDS": {
    "code": "ZDS",
    "company": "TELESPAZIO FRANCE",
    "city": "TOULOUSE",
    "country": "France"
  },
  "ZEE": {
    "code": "ZEE",
    "company": "ZEENNI S TRADING AGENCY",
    "city": "TRIPOLI",
    "country": "Lebanon"
  },
  "ZEN": {
    "code": "ZEN",
    "company": "ZEN BOX CO.,LTD",
    "city": "WALNUT",
    "country": "United States"
  },
  "ZET": {
    "code": "ZET",
    "company": "ZETA SYSTEM SPA",
    "city": "MATERA",
    "country": "Italy"
  },
  "ZGR": {
    "code": "ZGR",
    "company": "ZGROUP SAC",
    "city": "LIMA 36",
    "country": "Peru"
  },
  "ZHF": {
    "code": "ZHF",
    "company": "ZHONG HAN YANTAI FERRY CO LTD",
    "city": "YANTAI",
    "country": "China"
  },
  "ZHH": {
    "code": "ZHH",
    "company": "XI'AN ZHONGHUI LIANHUA INDUSTRY CO., LTD",
    "city": "Xi'an",
    "country": "China"
  },
  "ZHL": {
    "code": "ZHL",
    "company": "ZHENGHE LOGISTICS PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "ZHR": {
    "code": "ZHR",
    "company": "F.U.H TRANS-STAR S.C ZBIGNIEW, HALINA ROKOSZ",
    "city": "OCIEKA",
    "country": "Poland"
  },
  "ZHW": {
    "code": "ZHW",
    "company": "ZHAOHUA SUPPLY CHAIN MANAGEMENT GROUP CO., LTD.",
    "city": "TIANJIN",
    "country": "China"
  },
  "ZIM": {
    "code": "ZIM",
    "company": "ZIM INTEGRATED SHIPPING SERVICES LTD",
    "city": "HAIFA",
    "country": "Israel"
  },
  "ZIP": {
    "code": "ZIP",
    "company": "TANK SERVICE INC",
    "city": "HOUSTON, TX 77231 0417",
    "country": "United States"
  },
  "ZIT": {
    "code": "ZIT",
    "company": "SAFIM SRL",
    "city": "GENOVA",
    "country": "Italy"
  },
  "ZJW": {
    "code": "ZJW",
    "company": "XINJIANG TBEA GROUP LOGISTICS CO.,LTD.",
    "city": "Changji",
    "country": "China"
  },
  "ZMC": {
    "code": "ZMC",
    "company": "ZEAMARINE CARRIER GMBH",
    "city": "BREMEN",
    "country": "Germany"
  },
  "ZMN": {
    "code": "ZMN",
    "company": "ZHE JIANG MORITA NEW MATERIALS CO.,LTD",
    "city": "Jinhua City,Wuyi County",
    "country": "China"
  },
  "ZMO": {
    "code": "ZMO",
    "company": "ZIM INTEGRATED SHIPPING SERVICES LTD",
    "city": "HAIFA",
    "country": "Israel"
  },
  "ZMS": {
    "code": "ZMS",
    "company": "ZEPPELIN MOBILE SYSTEME GMBH",
    "city": "MECKENBEUREN",
    "country": "Germany"
  },
  "ZPL": {
    "code": "ZPL",
    "company": "KONRAD ZIPPEL SPEDITEUR GMBH & CO. KG",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "ZRG": {
    "code": "ZRG",
    "company": "CHINA GAS, HONGTONG  CORPORATION",
    "city": "Shenzhen",
    "country": "China"
  },
  "ZSE": {
    "code": "ZSE",
    "company": "ZHENG SUN ENGINEERING CO,LTD",
    "city": "KAOHSIUNG CITY",
    "country": "Taiwan, China"
  },
  "ZSJ": {
    "code": "ZSJ",
    "company": "CHINA CONTAINERIZED BULK LOGISTICS",
    "city": "Shanghai",
    "country": "China"
  },
  "ZSL": {
    "code": "ZSL",
    "company": "XINJIANG TBEA GROUP LOGISTICS CO.,LTD.",
    "city": "Changji",
    "country": "China"
  },
  "ZST": {
    "code": "ZST",
    "company": "JSC SIBUR-TRANS",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "ZTK": {
    "code": "ZTK",
    "company": "NIJMAN/ZEETANK INT TANKTRANSPORTEN BV",
    "city": "Spijkenisse",
    "country": "Netherlands"
  },
  "ZUZ": {
    "code": "ZUZ",
    "company": "CARGOSTORE WORLDWIDE TRADING LIMITED",
    "city": "LONDON SW19 7QD",
    "country": "United Kingdom"
  },
  "ZVL": {
    "code": "ZVL",
    "company": "OLEXIA GROUP S.A.",
    "city": "Luxembourg",
    "country": "Luxembourg"
  },
  "ZWF": {
    "code": "ZWF",
    "company": "ZIM WORLD FREIGHT PVT LTD",
    "city": "MUMBAI",
    "country": "India"
  },
  "ZWI": {
    "code": "ZWI",
    "company": "V+A ZWISSIG SA",
    "city": "SIERRE",
    "country": "Switzerland"
  },
  "ZYM": {
    "code": "ZYM",
    "company": "ZIJDERHAND MOERKAPELLE BV",
    "city": "MOERKAPELLE",
    "country": "Netherlands"
  }
}`
