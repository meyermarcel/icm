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

package main

import (
	"log"
	"math/rand"
	"path/filepath"
	"regexp"
	"unicode/utf8"
)

const ownersFileName = "owners.json"

const dateFormat = "2006-01-02"

type owner struct {
	Code    string `json:"code"`
	Company string `json:"company"`
	City    string `json:"city"`
	Country string `json:"country"`
}

var loadedOwners map[string]owner
var pathToOwners string

func initOwners(appDirPath string) {
	pathToOwners = initFile(filepath.Join(appDirPath, ownersFileName), []byte(ownersJSON))
	jsonUnmarshal(readFile(pathToOwners), &loadedOwners)
}

type ownerCode struct {
	value string
}

func (c ownerCode) Value() string {
	return c.value
}

func newOwnerCode(value string) ownerCode {

	if utf8.RuneCountInString(value) != 3 {
		log.Fatalf("'%s' is not three characters", value)
	}

	if !regexp.MustCompile(`[A-Z]{3}`).MatchString(value) {
		log.Fatalf("'%s' must be 3 letters", value)
	}
	return ownerCode{value}
}

func getOwner(code ownerCode) owner {
	return loadedOwners[code.Value()]
}

func getRandomOwnerCodes(count int) []ownerCode {
	var codes []ownerCode

	for k := range loadedOwners {
		if len(codes) >= count {
			break
		}
		codes = append(codes, newOwnerCode(k))
	}
	rand.Shuffle(len(codes), func(i, j int) {
		codes[i], codes[j] = codes[j], codes[i]
	})
	return codes
}

func updateOwners(newOwners map[string]owner) {
	for k, v := range newOwners {
		if _, exists := loadedOwners[k]; !exists {
			loadedOwners[k] = v
		}
	}
	writeFile(pathToOwners, jsonMarshal(loadedOwners))
}

func getRegexPartOwners() string {
	var regexString string
	for k := range loadedOwners {
		regexString += k + "|"
	}
	return regexString[:len(regexString)-1]
}

const ownersJSON = `{
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
    "city": "AMSTERDAM",
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
    "company": "PJSC ABAKANVAGONMASH",
    "city": "ABAKAN",
    "country": "Russian Federation"
  },
  "ABN": {
    "code": "ABN",
    "company": "CJSC RTH LOGISTIC",
    "city": "MOSCOW",
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
    "company": "AKZO NOBEL POLYMER CHEMICALS BV",
    "city": "Rotterdam",
    "country": "Netherlands"
  },
  "ACP": {
    "code": "ACP",
    "company": "CARDIGAS",
    "city": "HEUSDEN-ZOLDER",
    "country": "Belgium"
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
  "ADE": {
    "code": "ADE",
    "company": "AMBROSIUS DEUTSCHLAND GMBH",
    "city": "Frankfurt am Main",
    "country": "Germany"
  },
  "ADH": {
    "code": "ADH",
    "company": "ADHOC S.R.L",
    "city": "VENICE",
    "country": "Italy"
  },
  "ADL": {
    "code": "ADL",
    "company": "AL ERFAN COMPANY LTD",
    "city": "YiWu, ZhengJian",
    "country": "China"
  },
  "ADM": {
    "code": "ADM",
    "company": "TETRA LEASING AND TRADING INC.",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "ADN": {
    "code": "ADN",
    "company": "ADEMN LLC",
    "city": "Ternopil",
    "country": "Ukraine"
  },
  "ADR": {
    "code": "ADR",
    "company": "BARCELONESA DE DROGAS Y PRODUCTOS QUIMICOS SA",
    "city": "CORNELLA DE LLOBREGAT",
    "country": "Spain"
  },
  "ADX": {
    "code": "ADX",
    "company": "APPLIED CRYO TECHNOLOGIES",
    "city": "HOUSTON, TX-77075",
    "country": "United States"
  },
  "AEC": {
    "code": "AEC",
    "company": "PELCHEM PTY (LTD)",
    "city": "PRETORIA",
    "country": "South Africa"
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
    "city": "Martigny",
    "country": "Switzerland"
  },
  "AEX": {
    "code": "AEX",
    "company": "AFRICA EXPRESS LINE LTD",
    "city": "Kent  ME19 4UY",
    "country": "United Kingdom"
  },
  "AFA": {
    "code": "AFA",
    "company": "AFAMSA",
    "city": "MOS-PONTEVEDRA",
    "country": "Spain"
  },
  "AFB": {
    "code": "AFB",
    "company": "AFFILIPS NV",
    "city": "TIENEN",
    "country": "Belgium"
  },
  "AFD": {
    "code": "AFD",
    "company": "T.C. BASBAKANLIK AFET VE ACIL DURUM YONETIMI BASKANLIGI AFAD",
    "city": "ANKARA",
    "country": "Turkey"
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
  "AHL": {
    "code": "AHL",
    "company": "ABNORMAL LOAD ENGINEERING LIMITED",
    "city": "HIXON",
    "country": "United Kingdom"
  },
  "AIB": {
    "code": "AIB",
    "company": "AIR LIQUIDE BRASIL LTDA",
    "city": "SAO PAULO",
    "country": "Brazil"
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
  "AIM": {
    "code": "AIM",
    "company": "TRS CONTAINERS & CHASSIS",
    "city": "AVENEL, NJ 0700-0188",
    "country": "United States"
  },
  "AIR": {
    "code": "AIR",
    "company": "AIR PRODUCTS",
    "city": "ALLENTOWN, PA 18195",
    "country": "United States"
  },
  "AIS": {
    "code": "AIS",
    "company": "TRASPORTI PESANTI DI STORTI TULLIO & C.S.R.L",
    "city": "CASALMAGGIORE",
    "country": "Italy"
  },
  "AIT": {
    "code": "AIT",
    "company": "ARKEMA B.V. – LOCATION ROTTERDAM",
    "city": "Vondelingenplaat",
    "country": "Netherlands"
  },
  "AKB": {
    "code": "AKB",
    "company": "ALLNEX RESINS NETHERLANDS BV",
    "city": "BERGEN OP ZOOM",
    "country": "Netherlands"
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
    "country": "Taiwan (China)"
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
    "company": "AIR LIQUIDE ELECTRONICS MATERIALS",
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
  "ALX": {
    "code": "ALX",
    "company": "AIR LIQUIDE ELECTRONICS GMBH",
    "city": "DUSSELDORF",
    "country": "Germany"
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
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "AND": {
    "code": "AND",
    "company": "ANDREX BVBA",
    "city": "Antwerpen",
    "country": "Belgium"
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
    "city": "LUANDA",
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
  "ANS": {
    "code": "ANS",
    "company": "AKZO NOBEL CHEMICALS (NINGBO) CO.,Ltd.",
    "city": "SHANGHAI, 200040",
    "country": "China"
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
  "AON": {
    "code": "AON",
    "company": "NEC ENERGY SOLUTIONS",
    "city": "WESTBOROUGH, MA-01582",
    "country": "United States"
  },
  "AOR": {
    "code": "AOR",
    "company": "OEG ASIA PACIFIC PTE LTD",
    "city": "SINGAPORE",
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
  "APL": {
    "code": "APL",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "APM": {
    "code": "APM",
    "company": "MAERSK LINE A/S",
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
  "APZ": {
    "code": "APZ",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "AQA": {
    "code": "AQA",
    "company": "OPERATOR LLC",
    "city": "Kazan",
    "country": "Russian Federation"
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
    "company": "ARGON ISOTANK LTD",
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
  "ATO": {
    "code": "ATO",
    "company": "ARKEMA",
    "city": "COLOMBES",
    "country": "France"
  },
  "ATR": {
    "code": "ATR",
    "company": "ANEL MITORAJ ANDRZEJ I MITORAJ ELZBIETA",
    "city": "Nowe Skalmierzyce",
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
    "company": "ALTERMIJ TANKVERHUUR BV",
    "city": "ROTTERDAM-BOTLEK",
    "country": "Netherlands"
  },
  "ATX": {
    "code": "ATX",
    "company": "ANDERSON TRUCKING SERVICES",
    "city": "St. Cloud",
    "country": "United States"
  },
  "AUG": {
    "code": "AUG",
    "company": "QUEHENBERGER LOGISTICS GMBH",
    "city": "STRASSWALCHEN",
    "country": "Austria"
  },
  "AUN": {
    "code": "AUN",
    "company": "ADVANCE INTERNATIONAL TRANSPORT INC",
    "city": "ETILER - ISTANBUL",
    "country": "Turkey"
  },
  "AUS": {
    "code": "AUS",
    "company": "CEC SYSTEMS PTY LTD",
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
  "AVL": {
    "code": "AVL",
    "company": "ALLIED VANTAGE LEASING LIMITED",
    "city": "Hong Kong",
    "country": "HK"
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
  "AWV": {
    "code": "AWV",
    "company": "AIR WATER VIETNAM CO LTD",
    "city": "BA RIA - VUNG TAU PROVINCE",
    "country": "Viet Nam"
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
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "AZK": {
    "code": "AZK",
    "company": "SOCIETE DES MINES D'AZELIK SA",
    "city": "NIAMEY",
    "country": "Niger"
  },
  "AZL": {
    "code": "AZL",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
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
    "city": "Roselle Park",
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
  "BCS": {
    "code": "BCS",
    "company": "BALTIC CONTAINER SERVICE LTD",
    "city": "SAINT-PETERSBURG",
    "country": "Russian Federation"
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
    "city": "Velsen",
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
  "BEF": {
    "code": "BEF",
    "company": "TRANSPORT SYSTEMS LLC",
    "city": "MOSCOW",
    "country": "Russian Federation"
  },
  "BEL": {
    "code": "BEL",
    "company": "BELL-MAR SHIPPING LTD",
    "city": "HAIFA",
    "country": "Israel"
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
    "company": "BEZEV SP Z O.O.",
    "city": "Warsaw",
    "country": "Poland"
  },
  "BFF": {
    "code": "BFF",
    "company": "BAHAMAS FERRIES LTD.",
    "city": "Nassau",
    "country": "Bahamas"
  },
  "BFR": {
    "code": "BFR",
    "company": "MAYBURY TRADE LLP",
    "city": "LONDON EC3V 0EH",
    "country": "United Kingdom"
  },
  "BFS": {
    "code": "BFS",
    "company": "H20 INCORPORATED",
    "city": "LAFAYETTE, LA 70508",
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
    "company": "HRC SHIPPING SDN.BHD",
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
  "BIO": {
    "code": "BIO",
    "company": "CHEMPLUS LTD",
    "city": "BALZAN",
    "country": "Malta"
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
  "BJH": {
    "code": "BJH",
    "company": "SHANGHAI BAIJIN CHEMICAL GROUP CO LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "BKB": {
    "code": "BKB",
    "company": "BROEKEMA BULK BV",
    "city": "WOUDENBERG",
    "country": "Netherlands"
  },
  "BKD": {
    "code": "BKD",
    "company": "FSUE CTR OF OP OF SPACE GROUND BASED INF",
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
    "city": "MUMBAI",
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
  "BLN": {
    "code": "BLN",
    "company": "BRIDGE LOGISTICS LLC",
    "city": "MOSCOW",
    "country": "Russian Federation"
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
    "company": "TETRA LEASING AND TRADING INC.",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "BLT": {
    "code": "BLT",
    "company": "BALTICON S.A.",
    "city": "GDYNIA",
    "country": "Poland"
  },
  "BLX": {
    "code": "BLX",
    "company": "EMPRESA NAVEGACAO MADEIRENSE LDA",
    "city": "FUNCHAL",
    "country": "Portugal"
  },
  "BLZ": {
    "code": "BLZ",
    "company": "BLPL SINGAPORE PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "BMC": {
    "code": "BMC",
    "company": "BIOMASS CONTROLS LLC",
    "city": "Putnam",
    "country": "United States"
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
    "company": "LINDE LLC",
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
    "city": "CHANTILLY, VA 20151",
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
    "company": "BULK OIL AND LIQUID TRANSPORT PVT LTD",
    "city": "SINGAPORE",
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
  "BQC": {
    "code": "BQC",
    "company": "BECHTEL OIL GAS & CHEMICALS INC",
    "city": "HOUSTON TX-77056-6503",
    "country": "United States"
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
  "BSB": {
    "code": "BSB",
    "company": "BIGSTEEL BOX CORPORATION",
    "city": "KELOWA, BC",
    "country": "Canada"
  },
  "BSC": {
    "code": "BSC",
    "company": "BGOOD SHIPPING CONTAINERS FZ LLC",
    "city": "RAS AL KHAIMAH",
    "country": "United Arab Emirates"
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
    "city": "BUCKINGHAMSHIRE, SL7 2NL",
    "country": "United Kingdom"
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
  "BSV": {
    "code": "BSV",
    "company": "BSV SHIPPING AGENCIES (S) PTE.LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "BSX": {
    "code": "BSX",
    "company": "NV BESIX SA",
    "city": "Brussels",
    "country": "Belgium"
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
  "BTR": {
    "code": "BTR",
    "company": "LLC BOTRANS",
    "city": "MOSCOW",
    "country": "Russian Federation"
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
  "BUN": {
    "code": "BUN",
    "company": "BUNDI SYSTEMS INTERNATIONAL LTD",
    "city": "AUCKLAND",
    "country": "New Zealand"
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
    "city": "SINGAPORE",
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
    "company": "PAN ADRIATIC HOLDINGS S.A.",
    "city": "LUGANO",
    "country": "Switzerland"
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
    "company": "HAMBURG SUED",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "CAE": {
    "code": "CAE",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
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
  "CBB": {
    "code": "CBB",
    "company": "IPES INDUNSTRIA DE PRODUTOS E EQUIPAMENTOS DE SOLDA LTDA",
    "city": "MANAUS",
    "country": "Brazil"
  },
  "CBC": {
    "code": "CBC",
    "company": "CONTAINER BROKERAGE COMPANY LTD",
    "city": "CHALFONT ST PETER",
    "country": "United Kingdom"
  },
  "CBF": {
    "code": "CBF",
    "company": "CHIQUITA BRANDS INC",
    "city": "CINCINNATI, OH 45202",
    "country": "United States"
  },
  "CBG": {
    "code": "CBG",
    "company": "CARBOGAZ S.A.",
    "city": "Croix-des-Bouquets",
    "country": "Haiti"
  },
  "CBH": {
    "code": "CBH",
    "company": "FLORENS CONTAINER SERVICES CO LTD",
    "city": "Taipa",
    "country": "Macao"
  },
  "CBI": {
    "code": "CBI",
    "company": "CHICAGO BRIDGE & IRON COMPANY",
    "city": "THE WOODLANDS,TX 77380",
    "country": "United States"
  },
  "CBK": {
    "code": "CBK",
    "company": "BELARUSKALI JSC",
    "city": "SOLIGORSK, MINSK REGION",
    "country": "Belarus"
  },
  "CBL": {
    "code": "CBL",
    "company": "CBOX CONTAINERS",
    "city": "Amsterdam",
    "country": "Netherlands"
  },
  "CBM": {
    "code": "CBM",
    "company": "CONTLEASE",
    "city": "Saint Petersbourg",
    "country": "Russian Federation"
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
    "company": "CBOX CONTAINERS",
    "city": "Amsterdam",
    "country": "Netherlands"
  },
  "CCA": {
    "code": "CCA",
    "company": "CEMENT CARTAGE COMPANY LIMITED",
    "city": "HAVELOCK N.B",
    "country": "Canada"
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
    "company": "CONTAINER PROVIDERS INT APS",
    "city": "COPENHAGEN K",
    "country": "Denmark"
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
  "CDP": {
    "code": "CDP",
    "company": "CONTAINERS RIO DE LA PLATA",
    "city": "BUENOS AIRES",
    "country": "Argentina"
  },
  "CEA": {
    "code": "CEA",
    "company": "AREVA TN INTERNATIONAL",
    "city": "MONTIGNY-LE-BRETONNEUX",
    "country": "France"
  },
  "CEF": {
    "code": "CEF",
    "company": "SOLAR ENERGY FACTORY CO LTD",
    "city": "MAIZURU, KYOTO",
    "country": "Japan"
  },
  "CEG": {
    "code": "CEG",
    "company": "CORBAN ENERGY GROUP",
    "city": "Elmwood Park",
    "country": "United States"
  },
  "CEI": {
    "code": "CEI",
    "company": "CONTENEURS EXPERTS S.D.INC",
    "city": "VAUDREUIL-DORION",
    "country": "Canada"
  },
  "CEM": {
    "code": "CEM",
    "company": "CHEMION LOGISTIK GMBH",
    "city": "LEVERKUSEN",
    "country": "Germany"
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
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
    "company": "BOURGEY MONTREUIL ITALIA SRL",
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
  "CII": {
    "code": "CII",
    "company": "CONTAINER-IT LLC",
    "city": "Portland, OR 97212",
    "country": "United States"
  },
  "CIL": {
    "code": "CIL",
    "company": "CARGOSTORE WORLDWIDE TRADING LIMITED",
    "city": "LONDON SW19 7QD",
    "country": "United Kingdom"
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
    "city": "AMSTERDAM",
    "country": "Netherlands"
  },
  "CJY": {
    "code": "CJY",
    "company": "CHONGQING TRANSPORTATION HOLDING (GROUP) CO,LTD",
    "city": "CHONGQING",
    "country": "China"
  },
  "CKB": {
    "code": "CKB",
    "company": "CEDARKNIGHT CAPITAL LIMITED",
    "city": "London",
    "country": "United Kingdom"
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
  "CKP": {
    "code": "CKP",
    "company": "CONTAINER K PLUS GMBH",
    "city": "HAMBURG",
    "country": "Germany"
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
    "city": "ZAPY",
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
    "city": "Norwell",
    "country": "United States"
  },
  "CLH": {
    "code": "CLH",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "CLI": {
    "code": "CLI",
    "company": "CLBT S.R.L",
    "city": "BENTIVOGLIO - BO",
    "country": "Italy"
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
    "city": "Middlesbrough",
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
    "company": "CONTAINER LEASING UK LTD",
    "city": "MONMOUTHSHIRE, NP26 5PT",
    "country": "United Kingdom"
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
    "company": "HAMBURG SUED",
    "city": "HAMBURG",
    "country": "Germany"
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
  "COB": {
    "code": "COB",
    "company": "SONANGOL PESQUISA & PRODUCAO SA",
    "city": "LUANDA",
    "country": "Angola"
  },
  "COC": {
    "code": "COC",
    "company": "SAMSKIP MCL BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "COG": {
    "code": "COG",
    "company": "AREVA TN INTERNATIONAL",
    "city": "MONTIGNY-LE-BRETONNEUX",
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
    "city": "FT.LAUDERDALE Florida",
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
  "COW": {
    "code": "COW",
    "company": "INGURAN LLC DBA SEXING TECHNOLOGIES",
    "city": "NAVASOTA, TX 77868",
    "country": "United States"
  },
  "COX": {
    "code": "COX",
    "company": "CEYLON OXYGEN LIMITED",
    "city": "Colombo",
    "country": "Sri Lanka"
  },
  "COZ": {
    "code": "COZ",
    "company": "MAERSK LINE A/S",
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
    "company": "FUTURA ENTERPRISE SRL",
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
    "company": "CONTAINER PROVIDERS INTERNATIONAL",
    "city": "COPENHAGEN K",
    "country": "Denmark"
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
    "company": "CONTAINER PROVIDERS INT APS",
    "city": "COPENHAGEN K",
    "country": "Denmark"
  },
  "CPZ": {
    "code": "CPZ",
    "company": "CARGO PLAN INTERNATIONAL (PVT) LIMITED",
    "city": "KARACHI",
    "country": "Pakistan"
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
  "CRG": {
    "code": "CRG",
    "company": "CARGOTAINER GMBH",
    "city": "MANNHEIM",
    "country": "Germany"
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
    "company": "ACE ENGINEERING & CO LTD",
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
  "CSB": {
    "code": "CSB",
    "company": "CEBU SEA CHARTERERS INC.",
    "city": "CEBU",
    "country": "Philippines"
  },
  "CSC": {
    "code": "CSC",
    "company": "CONTAINER SOLUTIONS CO LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "CSF": {
    "code": "CSF",
    "company": "CONTAINERSHIPS PLC (CONTAINERSHIPS OYJ)",
    "city": "Espoo",
    "country": "Finland"
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
    "company": "CONTAINERSHIPS PLC (CONTAINERSHIPS OYJ)",
    "city": "Espoo",
    "country": "Finland"
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
  "CSZ": {
    "code": "CSZ",
    "company": "CONTAINER STORAGE",
    "city": "Reno",
    "country": "United States"
  },
  "CTA": {
    "code": "CTA",
    "company": "COOLTAINER NEW ZEALAND LIMITED",
    "city": "CHRISTCHURCH",
    "country": "New Zealand"
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
    "city": "Trois-Rivieres",
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
    "company": "SOCODEI",
    "city": "BAGNOLS SUR CEZE",
    "country": "France"
  },
  "CTT": {
    "code": "CTT",
    "company": "CANTAS IC VE DIS TIC.SOG.SIS.SAN.A.S.",
    "city": "ISTANBUL",
    "country": "Turkey"
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
  "CUL": {
    "code": "CUL",
    "company": "CHINA UNITED LINES LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "CUN": {
    "code": "CUN",
    "company": "COMBIUNITS AS",
    "city": "EIDSVAGNESET",
    "country": "Norway"
  },
  "CVA": {
    "code": "CVA",
    "company": "CRYOGENIC VESSEL ALTERNATIVES INC",
    "city": "BAYTOWN,TX 77523",
    "country": "United States"
  },
  "CVV": {
    "code": "CVV",
    "company": "CV SUKSES MAJU BERSAMA",
    "city": "CIKARANG, BEKASI",
    "country": "Indonesia"
  },
  "CVX": {
    "code": "CVX",
    "company": "CHEVRON AUSTRALIA PTY LTD",
    "city": "PERTH WA 6000",
    "country": "Australia"
  },
  "CWF": {
    "code": "CWF",
    "company": "GUANGZHOU CO WIN INTERNATIONAL FREIGHT FORWARDING AGENCY CO",
    "city": "GUANGZHOU",
    "country": "China"
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
    "company": "CONTAINERSHIPS PLC (CONTAINERSHIPS OYJ)",
    "city": "Espoo",
    "country": "Finland"
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
    "company": "DANISH DEFENCE ACQUISITION AND LOGISTICS",
    "city": "BALLERUP",
    "country": "Denmark"
  },
  "DAC": {
    "code": "DAC",
    "company": "DANCONTAINER A/S",
    "city": "COPENHAGEN",
    "country": "Denmark"
  },
  "DAD": {
    "code": "DAD",
    "company": "TANKTRAILER NEDERLAND",
    "city": "NUMANSDORP",
    "country": "Netherlands"
  },
  "DAH": {
    "code": "DAH",
    "company": "DAHLSON INDUSTRIES LTD",
    "city": "ROCKY VIEW / ALBERTA, TIX 0K3",
    "country": "Canada"
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
  "DBG": {
    "code": "DBG",
    "company": "DREAMBOX GLOBAL",
    "city": "GARDENA, CA-90248",
    "country": "United States"
  },
  "DBO": {
    "code": "DBO",
    "company": "DE BOER CONTAINER TRADING B.V.",
    "city": "ECHT",
    "country": "Netherlands"
  },
  "DBR": {
    "code": "DBR",
    "company": "DRILL BASE OIL RESOURCES SDN BHD",
    "city": "MONT\\'KIARA",
    "country": "Malaysia"
  },
  "DBS": {
    "code": "DBS",
    "company": "DIONYSOS",
    "city": "SOFIA",
    "country": "Bulgaria"
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
    "company": "MINISTERE DE LA DEFENSE",
    "city": "VILLACOUBLAY",
    "country": "France"
  },
  "DCN": {
    "code": "DCN",
    "company": "EDF SA",
    "city": "ST DENIS",
    "country": "France"
  },
  "DCO": {
    "code": "DCO",
    "company": "DOW CORNING LIMITED",
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
    "company": "AKZO NOBEL PULP AND PERFORMANCE CHEMICALS AB",
    "city": "MALMO",
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
    "company": "SGS ESPANOLA DE CONTROL S.A.",
    "city": "MADRID",
    "country": "Spain"
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
    "city": "Hayward",
    "country": "United States"
  },
  "DEN": {
    "code": "DEN",
    "company": "HARDING CONTAINERS INTERNATIONAL INC.",
    "city": "LONG BEACH",
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
    "company": "FLORENS CONTAINER SERVICES CO LTD",
    "city": "Taipa",
    "country": "Macao"
  },
  "DFS": {
    "code": "DFS",
    "company": "FLORENS CONTAINER SERVICES CO LTD",
    "city": "Taipa",
    "country": "Macao"
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
  "DGI": {
    "code": "DGI",
    "company": "ISOCHEM LOGISTICS LLC",
    "city": "Houston",
    "country": "United States"
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
  "DHE": {
    "code": "DHE",
    "company": "D&H EQUIPMENT,LTD",
    "city": "BLANCO, TX 78606",
    "country": "United States"
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
    "city": "",
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
  "DIT": {
    "code": "DIT",
    "company": "DRY ICE TECH., LTD",
    "city": "Chia-Yi",
    "country": "Taiwan (China)"
  },
  "DJC": {
    "code": "DJC",
    "company": "DJ CONTAINERS GVC",
    "city": "ANTWERPEN",
    "country": "Belgium"
  },
  "DJI": {
    "code": "DJI",
    "company": "DJIBOUTI NATIONAL SHIPPING COMPANY",
    "city": "Djibouti",
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
  "DLT": {
    "code": "DLT",
    "company": "DALREFTRANS LTD",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "DLV": {
    "code": "DLV",
    "company": "DELVER AGENTS LLC",
    "city": "Seattle",
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
    "city": "COPENHAGEN",
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
    "city": "PORT COQUITLAM, BRITICH COLUMBIA",
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
  "DPD": {
    "code": "DPD",
    "company": "CHEMOURS",
    "city": "DORDRECHT",
    "country": "Netherlands"
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
    "city": "VERACRUZ",
    "country": "Mexico"
  },
  "DRY": {
    "code": "DRY",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
  },
  "DSN": {
    "code": "DSN",
    "company": "CEA",
    "city": "St Paul Lez Durance",
    "country": "France"
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
  "DTB": {
    "code": "DTB",
    "company": "BITUMINA REFINED OIL PRODUCTS TRADING LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "DTC": {
    "code": "DTC",
    "company": "DECCAN TRANSCON LEASING PVT LTD",
    "city": "Hyderabad",
    "country": "India"
  },
  "DTG": {
    "code": "DTG",
    "company": "DOMINION TECHNOLOGY GASES LTD",
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
    "city": "LONG BEACH",
    "country": "United States"
  },
  "DUC": {
    "code": "DUC",
    "company": "DUTCH ANTILLEAN CONTAINER LEASING NV",
    "city": "LIMASSOL",
    "country": "Cyprus"
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
  "DYK": {
    "code": "DYK",
    "company": "SPK",
    "city": "Ulyanovsk",
    "country": "Russian Federation"
  },
  "DYL": {
    "code": "DYL",
    "company": "DONGYOUNG SHIPPING CO.,LTD",
    "city": "Jung-gu, Seoul",
    "country": "Korea, Republic of"
  },
  "DYM": {
    "code": "DYM",
    "company": "DANMAR LOGISTICA CRYO SL",
    "city": "TARRAGONA",
    "country": "Spain"
  },
  "DYO": {
    "code": "DYO",
    "company": "DYLO INC",
    "city": "Mcallen",
    "country": "United States"
  },
  "DZD": {
    "code": "DZD",
    "company": "DZT LOGISTIC",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "DZL": {
    "code": "DZL",
    "company": "TORINO GRAZYNA NYCZ",
    "city": "BOLESLAWIEC",
    "country": "Poland"
  },
  "DZT": {
    "code": "DZT",
    "company": "JSC DALZAVOD TERMINAL",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
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
  "EAM": {
    "code": "EAM",
    "company": "MARIANI ASPHALT COMPANY",
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
  "ECT": {
    "code": "ECT",
    "company": "TRANS EUROPEAN CONTAINER UNIT T.E.C.U SRL",
    "city": "PADOVA PD",
    "country": "Italy"
  },
  "ECW": {
    "code": "ECW",
    "company": "ECCO WOOD SP. Z Z.O.",
    "city": "Katowice",
    "country": "Poland"
  },
  "EDC": {
    "code": "EDC",
    "company": "DAHER NUCLEAR TECHNOLOGIES",
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
  "EEE": {
    "code": "EEE",
    "company": "E2E SUPPLY CHAIN MANAGEMENT LIMITED",
    "city": "KARACHI",
    "country": "Pakistan"
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
  "EFT": {
    "code": "EFT",
    "company": "EUROPEAN FOOD TRANSPORT N.V.",
    "city": "WILLEBROEK",
    "country": "Belgium"
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
    "city": "Taoyuan City",
    "country": "Taiwan (China)"
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
  "EIC": {
    "code": "EIC",
    "company": "EDLOW INTERNATIONAL",
    "city": "Washington",
    "country": "United States"
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
    "country": "Taiwan (China)"
  },
  "EKY": {
    "code": "EKY",
    "company": "E-KWAN ENTERPRISE CO., LTD",
    "city": "Kaohsiung",
    "country": "Taiwan (China)"
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
    "city": "ELBURG",
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
  "ELO": {
    "code": "ELO",
    "company": "MACS MARITIME CARRIER SHIPPING GMBH & CO",
    "city": "HAMBURG",
    "country": "Germany"
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
    "company": "PHU EL-TRANS ELZBIETA SIEROCKA",
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
    "company": "EMBALEX S.L.",
    "city": "Palau Solita i Plegamans (Barcelona)",
    "country": "Spain"
  },
  "EMC": {
    "code": "EMC",
    "company": "EVERGREEN MARINE CORP (TAIWAN) LTD",
    "city": "TAOYUAN COUNTY",
    "country": "Taiwan (China)"
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
    "city": "Luneburg",
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
    "company": "HAMBURG SUED",
    "city": "HAMBURG",
    "country": "Germany"
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
  "ENE": {
    "code": "ENE",
    "company": "ADS DEMENTELEMENT & ASSAINISSMENT",
    "city": "Suresnes",
    "country": "France"
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
    "company": "EMPRESA NAVEGACAO MADEIRENSE LDA.",
    "city": "FUNCHAL",
    "country": "Portugal"
  },
  "ENR": {
    "code": "ENR",
    "company": "TRANSCOM LLP",
    "city": "ALMATY",
    "country": "Kazakhstan"
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
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "EOS": {
    "code": "EOS",
    "company": "CRI/CRITERION MARKETING ASIA PACIFIC PTE LTD",
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
  "ERN": {
    "code": "ERN",
    "company": "WIHELM ERNST GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "ESD": {
    "code": "ESD",
    "company": "EMIRATES SHIPPING (HONG KONG) LTD",
    "city": "QUARRY BAY",
    "country": "HK"
  },
  "ESE": {
    "code": "ESE",
    "company": "ESSECO SRL",
    "city": "S.MARTINO TRECATE  NO",
    "country": "Italy"
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
    "city": "QUARRY BAY",
    "country": "HK"
  },
  "ESR": {
    "code": "ESR",
    "company": "ECO SPRING RESOURCES PTE. LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "ESS": {
    "code": "ESS",
    "company": "KAWASAKI KISEN KAISHA LTD - K LINE",
    "city": "TOKYO",
    "country": "Japan"
  },
  "EST": {
    "code": "EST",
    "company": "VUDINVEST GAS",
    "city": "Viimsi",
    "country": "Estonia"
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
  "ETL": {
    "code": "ETL",
    "company": "ELEPHANT PETROL LLC",
    "city": "Tbilisi",
    "country": "Georgia"
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
    "company": "ELBTAINER TRADING GmbH",
    "city": "HAMBURG",
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
  "EVG": {
    "code": "EVG",
    "company": "EUROAZIA TRANS LLC",
    "city": "Saint-Petersburg",
    "country": "Russian Federation"
  },
  "EVK": {
    "code": "EVK",
    "company": "EVONIK CORPORATION",
    "city": "Parsippany",
    "country": "United States"
  },
  "EVO": {
    "code": "EVO",
    "company": "EVOLUTION GERADORES LTDA",
    "city": "ITAJAI",
    "country": "Brazil"
  },
  "EVR": {
    "code": "EVR",
    "company": "EVR CARGO",
    "city": "TALLINN",
    "country": "Estonia"
  },
  "EWA": {
    "code": "EWA",
    "company": "EWALS CARGO CARE BV",
    "city": "TEGELEN",
    "country": "Netherlands"
  },
  "EWC": {
    "code": "EWC",
    "company": "EAST WEST CONTINENTAL CONTAINER LINE",
    "city": "VLADIMIR MKR VUREVETS",
    "country": "Russian Federation"
  },
  "EWN": {
    "code": "EWN",
    "company": "JEN GMBH",
    "city": "Juelich",
    "country": "Germany"
  },
  "EWS": {
    "code": "EWS",
    "company": "DB CARGO (UK) LIMITED",
    "city": "DONCASTER DN4 5PN",
    "country": "United Kingdom"
  },
  "EXF": {
    "code": "EXF",
    "company": "EXSIF WORLDWIDE",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
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
  "FAA": {
    "code": "FAA",
    "company": "MAERSK LINE A/S",
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
    "company": "FLEX BOX COLUMBIA SAS",
    "city": "BOGOTA",
    "country": "Colombia"
  },
  "FAN": {
    "code": "FAN",
    "company": "HAPAG LLOYD A.G",
    "city": "HAMBURG",
    "country": "Germany"
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
    "city": "Gyeonggi-do",
    "country": "Korea, Republic of"
  },
  "FBI": {
    "code": "FBI",
    "company": "FLORENS CONTAINER SERVICES CO LTD",
    "city": "Taipa",
    "country": "Macao"
  },
  "FBL": {
    "code": "FBL",
    "company": "FLORENS CONTAINER SERVICES CO LTD",
    "city": "Taipa",
    "country": "Macao"
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
    "company": "FLEX BOX COLUMBIA SAS",
    "city": "BOGOTA",
    "country": "Colombia"
  },
  "FCB": {
    "code": "FCB",
    "company": "FLEX BOX COLUMBIA SAS",
    "city": "BOGOTA",
    "country": "Colombia"
  },
  "FCC": {
    "code": "FCC",
    "company": "FAR EASTERN SHIPPING PLC",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "FCD": {
    "code": "FCD",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "FCG": {
    "code": "FCG",
    "company": "FLORENS ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "FCI": {
    "code": "FCI",
    "company": "FLORENS CONTAINER SERVICES CO LTD",
    "city": "Taipa",
    "country": "Macao"
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
  "FFL": {
    "code": "FFL",
    "company": "SAURASHTRA FREIGHT PRIVATE LIMITED",
    "city": "Mumbai",
    "country": "India"
  },
  "FFM": {
    "code": "FFM",
    "company": "FUJIFILM ELECTRONIC MATERIALS USA, INC.",
    "city": "MESA",
    "country": "United States"
  },
  "FGN": {
    "code": "FGN",
    "company": "FLINT GROUP NETHERLANDS B.V.",
    "city": "´s-Gravenzande",
    "country": "Netherlands"
  },
  "FGR": {
    "code": "FGR",
    "company": "FERGUSON SEACABS LTD",
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
    "city": "CHUR",
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
  "FJK": {
    "code": "FJK",
    "company": "FLORENS CONTAINER SERVICES CO LTD",
    "city": "Taipa",
    "country": "Macao"
  },
  "FKI": {
    "code": "FKI",
    "company": "FINE LINK HOLDINGS LIMITED",
    "city": "MONGKOK, KOWLOON",
    "country": "HK"
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
    "city": "PALMETTO,FL 34221",
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
  "FLZ": {
    "code": "FLZ",
    "company": "FINE LINK HOLDINGS LIMITED",
    "city": "Yuen Long",
    "country": "HK"
  },
  "FMB": {
    "code": "FMB",
    "company": "P & O FERRYMASTERS LTD",
    "city": "ZEEBRUGGE",
    "country": "Belgium"
  },
  "FMC": {
    "code": "FMC",
    "company": "FMC CORP.",
    "city": "CHARLOTTE, NC-28208",
    "country": "United States"
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
    "city": "LIMA 32",
    "country": "Peru"
  },
  "FMI": {
    "code": "FMI",
    "company": "FONG-MING GASES INDUSTRIAL CO., LTD",
    "city": "Taipei",
    "country": "Taiwan (China)"
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
  "FNG": {
    "code": "FNG",
    "company": "FERUS",
    "city": "CALGARY, ALBERTA",
    "country": "Canada"
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
    "city": "OSS",
    "country": "Netherlands"
  },
  "FPE": {
    "code": "FPE",
    "company": "AXENS",
    "city": "SALINDRES",
    "country": "France"
  },
  "FPG": {
    "code": "FPG",
    "company": "FPG RAFFLES PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "FPM": {
    "code": "FPM",
    "company": "FORMOSA PLASTICS MARINE CORP",
    "city": "TAIPEI",
    "country": "Taiwan (China)"
  },
  "FPO": {
    "code": "FPO",
    "company": "FREEPORT GASES LTD",
    "city": "GRAND BAHAMA",
    "country": "Bahamas"
  },
  "FPR": {
    "code": "FPR",
    "company": "FPG RAFFLES PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "FPS": {
    "code": "FPS",
    "company": "COMPANY SERVICES SWISS SA",
    "city": "LUGANO",
    "country": "Switzerland"
  },
  "FPT": {
    "code": "FPT",
    "company": "FPG RAFFLES PTE LTD",
    "city": "SINGAPORE",
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
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "FRO": {
    "code": "FRO",
    "company": "FRIMO SA",
    "city": "Asuncion",
    "country": "Paraguay"
  },
  "FSC": {
    "code": "FSC",
    "company": "FLORENS CONTAINER SERVICES CO LTD",
    "city": "Taipa",
    "country": "Macao"
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
  "FTG": {
    "code": "FTG",
    "company": "FINTRANS GL",
    "city": "St Petesburg",
    "country": "Russian Federation"
  },
  "FTH": {
    "code": "FTH",
    "company": "LESEA GLOBAL FEED THE HUNGRY",
    "city": "South Bend",
    "country": "United States"
  },
  "FTI": {
    "code": "FTI",
    "company": "FIBA TECHNOLOGIES, INC",
    "city": "Littleton, MA 01460",
    "country": "United States"
  },
  "FTL": {
    "code": "FTL",
    "company": "FRONTIER LOGISTICS  LP",
    "city": "LA PORTE TX-77571",
    "country": "United States"
  },
  "FTP": {
    "code": "FTP",
    "company": "FAHAD THANAYYAN AL THANAYYAN & PARTNERS",
    "city": "RIYADH",
    "country": "Saudi Arabia"
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
    "city": "Brisloane",
    "country": "Australia"
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
    "company": "GUANGDONG FUWA EQUIPMENT MANUFACTURING CO., LTD.",
    "city": "FOSHAN",
    "country": "China"
  },
  "FWU": {
    "code": "FWU",
    "company": "TAYLOR MINSTER LEASING BV",
    "city": "SPYKENISSE",
    "country": "Netherlands"
  },
  "FXL": {
    "code": "FXL",
    "company": "FLEX BOX COLUMBIA SAS",
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
  "GAE": {
    "code": "GAE",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "GBK": {
    "code": "GBK",
    "company": "TAM INTERNATIONAL INC",
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
  "GCL": {
    "code": "GCL",
    "company": "CHELSEA SHIPPING PTY LTD",
    "city": "Sydney",
    "country": "Australia"
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
    "city": "SINGAPORE",
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
  "GEB": {
    "code": "GEB",
    "company": "GR LOGISTICS AND TERMINALS LLC",
    "city": "TBILISI",
    "country": "Georgia"
  },
  "GEC": {
    "code": "GEC",
    "company": "GREAT EXTEND CO,LTD",
    "city": "TAIPEI",
    "country": "Taiwan (China)"
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
    "company": "GULF FLUOR",
    "city": "ABU DHABI",
    "country": "United Arab Emirates"
  },
  "GFZ": {
    "code": "GFZ",
    "company": "GEOMAR",
    "city": "KIEL",
    "country": "Germany"
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
  "GGI": {
    "code": "GGI",
    "company": "INTERNATIONAL GOLDEN GROUP PJSC",
    "city": "ABU DHABI",
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
    "city": "Austin",
    "country": "United States"
  },
  "GIC": {
    "code": "GIC",
    "company": "EMPRESA DE GASES INDUSTRIALES",
    "city": "GUANABACOA",
    "country": "Cuba"
  },
  "GIL": {
    "code": "GIL",
    "company": "DOR CHEMICALS LTD",
    "city": "HAIFA",
    "country": "Israel"
  },
  "GIN": {
    "code": "GIN",
    "company": "GAS INNOVATIONS",
    "city": "LA PORTE",
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
    "company": "ROLLER CHEMICAL",
    "city": "Fornovo s. giovanni",
    "country": "Italy"
  },
  "GJG": {
    "code": "GJG",
    "company": "S.M.M.I  (SOCIETE DE MANUTENTION DE MATERIAUX INDUSTRIELS )",
    "city": "ABYMES",
    "country": "France"
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
  "GLG": {
    "code": "GLG",
    "company": "GLOBAL GAS SERVCES LLC",
    "city": "MUSCAT, PC 111",
    "country": "Oman"
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
  "GNR": {
    "code": "GNR",
    "company": "GENERON IGS",
    "city": "Houston",
    "country": "United States"
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
    "company": "HUIZHOU 3R ENVIROMENTAL CHEMICAL CO.,LTD",
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
  "GOL": {
    "code": "GOL",
    "company": "COMPOSITE ADVANCED TECHNOLOGIES CNG, LLC (CAT-CNG)",
    "city": "HOUSTON, TX-77075",
    "country": "United States"
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
  "GPI": {
    "code": "GPI",
    "company": "GULF PROCURERS FOR INDUSTRIAL SUPPLY W.L.L.",
    "city": "AL-KHOBAR",
    "country": "Saudi Arabia"
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
  "GRA": {
    "code": "GRA",
    "company": "GRANADA-2000",
    "city": "Moscow",
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
    "company": "HAMBURG SUED",
    "city": "HAMBURG",
    "country": "Germany"
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
  "GRN": {
    "code": "GRN",
    "company": "GREEN BOX CO. LTD",
    "city": "Seoul",
    "country": "Korea, Republic of"
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
    "city": "SINGAPORE",
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
  "GRV": {
    "code": "GRV",
    "company": "CRECEFULL S.A.",
    "city": "MONTEVIDEO",
    "country": "Uruguay"
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
    "company": "GRUPO GASMEDI S.L.U",
    "city": "MADRID",
    "country": "Spain"
  },
  "GSO": {
    "code": "GSO",
    "company": "GULF ORIENT SHIPPING SERVICE LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
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
    "city": "SINGAPORE",
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
  "GVI": {
    "code": "GVI",
    "company": "GVIDON LTD",
    "city": "Saint-Petersburg",
    "country": "Russian Federation"
  },
  "GVS": {
    "code": "GVS",
    "company": "GULFVOSTOK LTD",
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
    "city": "Haiphong",
    "country": "Viet Nam"
  },
  "HAE": {
    "code": "HAE",
    "company": "HAEWOO GLS CO LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "HAF": {
    "code": "HAF",
    "company": "HAFNARBAKKI HF",
    "city": "HAFNARFJORDUR",
    "country": "Iceland"
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
    "city": "Vlaardingen",
    "country": "Netherlands"
  },
  "HAS": {
    "code": "HAS",
    "company": "HAMBURG SUED",
    "city": "HAMBURG",
    "country": "Germany"
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
  "HAW": {
    "code": "HAW",
    "company": "HITACHI HIGH-TECH AW CRYO, INC.",
    "city": "Vancouver",
    "country": "Canada"
  },
  "HBC": {
    "code": "HBC",
    "company": "HOLDERCHEM BUILDING CHEMICALS S.A.L",
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
    "city": "Houston",
    "country": "United States"
  },
  "HBR": {
    "code": "HBR",
    "company": "HERNANDEZ BELLO, S.L",
    "city": "LA OROTAVA",
    "country": "Spain"
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
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "HCO": {
    "code": "HCO",
    "company": "CONICAL CONTAINER INDUSTRIE CONSULTING",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HCR": {
    "code": "HCR",
    "company": "SAF",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "HCS": {
    "code": "HCS",
    "company": "HACON EQUIPMENT B.V",
    "city": "EUROPOORT",
    "country": "Netherlands"
  },
  "HCV": {
    "code": "HCV",
    "company": "HONG CHUN ELECTRIC & MACHINERY CO LTD",
    "city": "MIAO-LI",
    "country": "Taiwan (China)"
  },
  "HCZ": {
    "code": "HCZ",
    "company": "HARDING CONTAINERS INTERNATIONAL INC.",
    "city": "LONG BEACH",
    "country": "United States"
  },
  "HDF": {
    "code": "HDF",
    "company": "HUADONG FERRY CO LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "HDM": {
    "code": "HDM",
    "company": "HYUNDAI MERCHANT MARINE CO LTD",
    "city": "SEOUL",
    "country": "Korea, Republic of"
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
  "HEM": {
    "code": "HEM",
    "company": "TEMPUS LINK LTD",
    "city": "GABROVO",
    "country": "Bulgaria"
  },
  "HER": {
    "code": "HER",
    "company": "HERC RENTALS",
    "city": "PACHECO",
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
    "company": "HOYER GLOBAL TRANSPORT BV",
    "city": "SPIJKENISSE",
    "country": "Netherlands"
  },
  "HGC": {
    "code": "HGC",
    "company": "THE GAS COMPANY LLC D.B.A. HAWAII GAS",
    "city": "Honolulu",
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
    "company": "HOYER GLOBAL TRANSPORT BV",
    "city": "SPIJKENISSE",
    "country": "Netherlands"
  },
  "HGH": {
    "code": "HGH",
    "company": "HOYER GLOBAL TRANSPORT BV",
    "city": "SPIJKENISSE",
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
    "company": "HOYER GLOBAL TRANSPORT BV",
    "city": "SPIJKENISSE",
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
    "city": "BREENDONK",
    "country": "Belgium"
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
    "country": "Taiwan (China)"
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
    "country": "Taiwan (China)"
  },
  "HLG": {
    "code": "HLG",
    "company": "HAEFELI AG",
    "city": "LENZBURG",
    "country": "Switzerland"
  },
  "HLS": {
    "code": "HLS",
    "company": "HLS CONTAINER BREMEN",
    "city": "BREMEN",
    "country": "Germany"
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
  "HME": {
    "code": "HME",
    "company": "HERALD MARINE AND ENERGY LIMITED",
    "city": "APAPA, LAGOS",
    "country": "Nigeria"
  },
  "HMH": {
    "code": "HMH",
    "company": "HOOVER MATERIALS HANDLING GROUP, INC",
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
    "company": "HYUNDAI MERCHANT MARINE CO LTD",
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
    "company": "HONEYTAK CONTAINER LIMITED",
    "city": "Yantai",
    "country": "China"
  },
  "HNL": {
    "code": "HNL",
    "company": "VAN OORD DREDGING AND MARINE CONTRACTORS BV",
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
    "company": "HOYER GLOBAL TRANSPORT BV",
    "city": "SPIJKENISSE",
    "country": "Netherlands"
  },
  "HOU": {
    "code": "HOU",
    "company": "HARDING CONTAINERS INTERNATIONAL INC.",
    "city": "LONG BEACH",
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
    "company": "HUIZHOU PACIFIC CONTAINER CO.LTD",
    "city": "XINXU",
    "country": "China"
  },
  "HPE": {
    "code": "HPE",
    "company": "BECKER MARINE SYSTEMS GMBH & CO KG",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "HPG": {
    "code": "HPG",
    "company": "TIANJIN SINO-PEAK INTERNATIONAL TRADE CO LTD",
    "city": "BINHAI NEW DISTRICT, TIANJIN 300460",
    "country": "China"
  },
  "HRT": {
    "code": "HRT",
    "company": "LOXAM ALQUILER S.A",
    "city": "Coslada - MADRID",
    "country": "Spain"
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
  "HSF": {
    "code": "HSF",
    "company": "TIANJIN LVHAI AGRICULTURAL TRADE CO LTD",
    "city": "TIANJIN CITY",
    "country": "China"
  },
  "HSI": {
    "code": "HSI",
    "company": "HOYER GLOBAL TRANSPORT BV",
    "city": "SPIJKENISSE",
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
    "company": "HUGO STINNES SCHIFFAHRT GMBH",
    "city": "ROSTOCK",
    "country": "Germany"
  },
  "HTC": {
    "code": "HTC",
    "company": "HALDOR TOPSOE A/S",
    "city": "KGS.LYNGBY",
    "country": "Denmark"
  },
  "HTI": {
    "code": "HTI",
    "company": "HUDSON TECHNOLOGIES INC.",
    "city": "PEARL RIVER, NY 10965",
    "country": "United States"
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
  "IAA": {
    "code": "IAA",
    "company": "INTERASIA LINES SINGAPORE PTE.LTD",
    "city": "SINGAPORE",
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
  "IAI": {
    "code": "IAI",
    "company": "INTERTRONIC SOLUTIONS INC",
    "city": "ST LAZARE, QUEBEC",
    "country": "Canada"
  },
  "IAL": {
    "code": "IAL",
    "company": "IAL CONTAINER LINE (UK) LIMITED",
    "city": "LONDON",
    "country": "United Kingdom"
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
  "ICG": {
    "code": "ICG",
    "company": "INDUSTRIAL CHEMICALS LTD",
    "city": "GRAYS, ESSEX",
    "country": "United Kingdom"
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
    "company": "OL&T FOODTRANS LLC",
    "city": "Irvine",
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
    "company": "INTERFLOW (T.C.S.) LTD, JAPAN BRANCH OFFICE",
    "city": "Tokyo",
    "country": "Japan"
  },
  "IFT": {
    "code": "IFT",
    "company": "LEIBNIZ INSTITUT FOR TROPOSPHERIC RESEARCH",
    "city": "LEIPZIG",
    "country": "Germany"
  },
  "IGE": {
    "code": "IGE",
    "company": "INTERGERMANIA LOGISTICS PTY,LTD",
    "city": "FAIRLAND",
    "country": "South Africa"
  },
  "IGL": {
    "code": "IGL",
    "company": "IGLU GROUP PTE LTD",
    "city": "SINGAPORE",
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
  "ILG": {
    "code": "ILG",
    "company": "ILGINNOVATIVE LOGISTICS GROUP GMBH",
    "city": "LINZ",
    "country": "Austria"
  },
  "ILI": {
    "code": "ILI",
    "company": "TANK CONTAINERS INC.",
    "city": "ROSWELL, GA 30075",
    "country": "United States"
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
    "city": "WASHIGTON, DC 20036-3946",
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
    "company": "INDUSTRIAS  METALICAS  IMETAL  S.A",
    "city": "GIJON",
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
    "city": "SINGAPORE",
    "country": "Singapore"
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
    "city": "MILANO",
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
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "INN": {
    "code": "INN",
    "company": "SEACUBE CONTAINERS LLC",
    "city": "WOODCLIFF LAKE, N.J. 07677",
    "country": "United States"
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
    "company": "INGAS AE",
    "city": "Mariupol",
    "country": "Ukraine"
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
  "IRG": {
    "code": "IRG",
    "company": "REACH HOLDING GROUP (SHANGHAI) CO.,LTD",
    "city": "SHANGHAI",
    "country": "China"
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
  "ISK": {
    "code": "ISK",
    "company": "NIMSCHU ISKUDOW INC.",
    "city": "Montreal",
    "country": "Canada"
  },
  "ISL": {
    "code": "ISL",
    "company": "INDUSTRIAL SOLVENTS & CHEMICALS PVT, LTD",
    "city": "Mumbai",
    "country": "India"
  },
  "ISP": {
    "code": "ISP",
    "company": "INDRA SISTEMAS PORTUGAL",
    "city": "AMADORA",
    "country": "Portugal"
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
  "IXS": {
    "code": "IXS",
    "company": "SUZHOU INTEROX SEM CO LTD",
    "city": "Suzhou",
    "country": "China"
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
    "city": "AMSTERDAM",
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
    "country": "Taiwan (China)"
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
    "company": "JAN DOHMEN BV",
    "city": "HERKENBOSCH",
    "country": "Netherlands"
  },
  "JDZ": {
    "code": "JDZ",
    "company": "JANSENS & DIEPERINK",
    "city": "ZAANDAM",
    "country": "Netherlands"
  },
  "JFS": {
    "code": "JFS",
    "company": "ST JOHN FREIGHT SYSTEMS LTD",
    "city": "TUTICORIN",
    "country": "India"
  },
  "JGA": {
    "code": "JGA",
    "company": "JGA TANKERS SL",
    "city": "Pozoblanco",
    "country": "Spain"
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
    "country": "Taiwan (China)"
  },
  "JHT": {
    "code": "JHT",
    "company": "JUHUA TRADING (HONG KONG) LIMITED",
    "city": "HONG KONG",
    "country": "HK"
  },
  "JJD": {
    "code": "JJD",
    "company": "ITEC ENGINEERING",
    "city": "CHAMBLY",
    "country": "France"
  },
  "JKC": {
    "code": "JKC",
    "company": "KUNSHAN CONTAINER SPECIAL EQUIPMENT CO.LTD",
    "city": "QIANDENG TOWN",
    "country": "China"
  },
  "JKT": {
    "code": "JKT",
    "company": "ICOR INTERNATIONAL INC",
    "city": "INDIANAPOLIS, IN 46236",
    "country": "United States"
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
  "JSE": {
    "code": "JSE",
    "company": "SHIN ETSU CHEMICAL CO LTD",
    "city": "CHIYODA-KU TOKYO",
    "country": "Japan"
  },
  "JSK": {
    "code": "JSK",
    "company": "CRYOGENIC LOGISTIC EQUIPMENT & TRADING PTE LTD",
    "city": "Singapore",
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
    "city": "Tai Xing",
    "country": "China"
  },
  "JTM": {
    "code": "JTM",
    "company": "FLORENS CONTAINER SERVICES CO LTD",
    "city": "Taipa",
    "country": "Macao"
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
  "JXN": {
    "code": "JXN",
    "company": "SHANDONG CONNERGY CO LTD",
    "city": "Ji'nan",
    "country": "China"
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
  "KAL": {
    "code": "KAL",
    "company": "NKG KALA HAMBURG GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "KAM": {
    "code": "KAM",
    "company": "WOLFGANG ULRICH ISCONT LINES LTD. C/O EIMSKIP TRANSPORT GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "KAR": {
    "code": "KAR",
    "company": "TOO RKTFG ALAN LTD",
    "city": "ALMATY",
    "country": "Kazakhstan"
  },
  "KAS": {
    "code": "KAS",
    "company": "TAIYO NIPPON SANSO K-AIR INDIA PVT. LTD",
    "city": "PUNE, MAHARASHTRA",
    "country": "India"
  },
  "KAZ": {
    "code": "KAZ",
    "company": "JSC KUIBYSHEV AZOT",
    "city": "Togliatti",
    "country": "Russian Federation"
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
    "company": "OMG KOKKOLA CHEMICALS OY",
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
    "city": "SINGAPORE",
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
  "KEM": {
    "code": "KEM",
    "company": "KEMAT N.V.",
    "city": "Brussels",
    "country": "Belgium"
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
    "company": "HAMBURG SUED",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "KHL": {
    "code": "KHL",
    "company": "HAMBURG SUED",
    "city": "HAMBURG",
    "country": "Germany"
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
    "city": "SINGAPORE",
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
    "company": "CONTAINERSHIPS PLC (CONTAINERSHIPS OYJ)",
    "city": "Espoo",
    "country": "Finland"
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
  "KLS": {
    "code": "KLS",
    "company": "MARINI MAKINA",
    "city": "ANKARA",
    "country": "Turkey"
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
  "KMI": {
    "code": "KMI",
    "company": "KUANG MING ENTERPRISE CO. LTD",
    "city": "TAIPEI",
    "country": "Taiwan (China)"
  },
  "KML": {
    "code": "KML",
    "company": "MAXWAY MARITIME S.A",
    "city": "DALIAN",
    "country": "China"
  },
  "KMP": {
    "code": "KMP",
    "company": "TANK ONE NV",
    "city": "Kallo",
    "country": "Belgium"
  },
  "KMR": {
    "code": "KMR",
    "company": "KAMER GERI DONUSUM MUH SAN DIS TIC LTD",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "KMS": {
    "code": "KMS",
    "company": "YANG MING MARINE TRANSPORT CORP.",
    "city": "KEELUNG",
    "country": "Taiwan (China)"
  },
  "KMT": {
    "code": "KMT",
    "company": "KOREA MARINE TRANSPORT CO / E. C. TEAM",
    "city": "SEOUL CITY",
    "country": "Korea, Republic of"
  },
  "KNL": {
    "code": "KNL",
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "KNO": {
    "code": "KNO",
    "company": "BAHAMAS BULK BITUMEN LTD",
    "city": "Nassau",
    "country": "Bahamas"
  },
  "KNT": {
    "code": "KNT",
    "company": "KANTO CORPORATION INC.",
    "city": "PORTLAND, OR 97203",
    "country": "United States"
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
  "KPH": {
    "code": "KPH",
    "company": "KAZPHOSPHAT LLC",
    "city": "ALMATY",
    "country": "Kazakhstan"
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
    "country": "Taiwan (China)"
  },
  "KRC": {
    "code": "KRC",
    "company": "KIWIRAIL",
    "city": "Auckland",
    "country": "New Zealand"
  },
  "KRI": {
    "code": "KRI",
    "company": "KRICON SERVICES B.V.",
    "city": "OUD VOSSERNEER",
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
  "KSM": {
    "code": "KSM",
    "company": "KEPPEL SMIT TOWAGE PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "KST": {
    "code": "KST",
    "company": "SCHILDECKER TRANSPORT GMBH",
    "city": "PISCHELSDORF",
    "country": "Austria"
  },
  "KSZ": {
    "code": "KSZ",
    "company": "KRONOSPAN SZCZECINEK SP. Z O.O.",
    "city": "Szczecinek",
    "country": "Poland"
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
    "company": "KTK TRADING",
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
    "company": "KTZ EXPRESS OPERATOR",
    "city": "ASTANA",
    "country": "Kazakhstan"
  },
  "KTZ": {
    "code": "KTZ",
    "company": "KTZ EXPRESS OPERATOR",
    "city": "ASTANA",
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
    "city": "",
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
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
    "company": "JSC KTZ EXPRESS",
    "city": "ASTANA",
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
    "city": "AMSTERDAM",
    "country": "Netherlands"
  },
  "LAH": {
    "code": "LAH",
    "company": "L & T B.V.",
    "city": "AMSTERDAM",
    "country": "Netherlands"
  },
  "LAI": {
    "code": "LAI",
    "company": "IAL CONTAINER LINE (UK) LIMITED",
    "city": "LONDON",
    "country": "United Kingdom"
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
  "LAT": {
    "code": "LAT",
    "company": "TRANSATLANTIC LINES LLC",
    "city": "GREENWICH, CT 06830",
    "country": "United States"
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
  "LBP": {
    "code": "LBP",
    "company": "LANDBRIDGE RIZHAO PORT CO LTD",
    "city": "Rizhao",
    "country": "China"
  },
  "LCE": {
    "code": "LCE",
    "company": "FMC CORP.",
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
  "LCU": {
    "code": "LCU",
    "company": "LANCER CONTAINER LINES LTD",
    "city": "NAVI MUMBAI",
    "country": "India"
  },
  "LCY": {
    "code": "LCY",
    "company": "LCY CHEMICAL CORP.",
    "city": "Kaohsiung",
    "country": "Taiwan (China)"
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
    "company": "LEHIGH EQUIPMENT COMPANY INC",
    "city": "BURLINGTON, KY-41005",
    "country": "United States"
  },
  "LEE": {
    "code": "LEE",
    "company": "TRANSPORTBEDRIJF H.A. LEEMANS B.V.",
    "city": "VRIEZENVEEN",
    "country": "Netherlands"
  },
  "LEG": {
    "code": "LEG",
    "company": "LEGEND SHIPPING PTE LTD",
    "city": "singapore",
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
    "city": "ROAD TOWN, TORTOLA",
    "country": "Virgin Islands, British"
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
  "LGC": {
    "code": "LGC",
    "company": "LOGICO-LOGISTYKA KONTENEROWA SP Z.O.O",
    "city": "SOPOT",
    "country": "Poland"
  },
  "LGE": {
    "code": "LGE",
    "company": "CARU CONTAINERS BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "LGG": {
    "code": "LGG",
    "company": "LG CHEM.",
    "city": "Chungbuk",
    "country": "Korea, Republic of"
  },
  "LGH": {
    "code": "LGH",
    "company": "LEGAL & GENERAL HOMES LTD",
    "city": "LEEDS",
    "country": "United Kingdom"
  },
  "LGI": {
    "code": "LGI",
    "company": "LINDE GAS ITALIA SRL",
    "city": "ARLUNO (MI)",
    "country": "Italy"
  },
  "LGK": {
    "code": "LGK",
    "company": "JOINT STOCK COMPANY  LITHUANIAN RAILWAYS",
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
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "LGT": {
    "code": "LGT",
    "company": "LABGAZ (THAILAND) CO LTD",
    "city": "BANGKOK",
    "country": "Thailand"
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
  "LIC": {
    "code": "LIC",
    "company": "INTERMODAL CONTAINER SERVICE",
    "city": "VILNIUS",
    "country": "Lithuania"
  },
  "LIF": {
    "code": "LIF",
    "company": "NANTONG LIANFU INDUSTRIAL CO LTD",
    "city": "QIDONG",
    "country": "China"
  },
  "LIN": {
    "code": "LIN",
    "company": "LINDE AG, LINDE GAS DEUTSCHLAND",
    "city": "PULLACH",
    "country": "Germany"
  },
  "LIQ": {
    "code": "LIQ",
    "company": "LIQUIMET S.p.A.",
    "city": "Treviso",
    "country": "Italy"
  },
  "LIR": {
    "code": "LIR",
    "company": "ROLAN LTD",
    "city": "HAIFA",
    "country": "Israel"
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
    "city": "Houston",
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
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "LMC": {
    "code": "LMC",
    "company": "IGNAZIO MESSINA & CO",
    "city": "GENOVA  GE",
    "country": "Italy"
  },
  "LMR": {
    "code": "LMR",
    "company": "L.M.C. (LEMARECHAL CELESTIN)",
    "city": "VALOGNES",
    "country": "France"
  },
  "LND": {
    "code": "LND",
    "company": "LINDE LLC",
    "city": "Bridgewater, NJ 08807",
    "country": "United States"
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
    "company": "MAERSK LINE A/S",
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
  "LPP": {
    "code": "LPP",
    "company": "LONGVIEW POWER LLC",
    "city": "Maidsville",
    "country": "United States"
  },
  "LPZ": {
    "code": "LPZ",
    "company": "LIFTING POINT",
    "city": "PENRITH NSW",
    "country": "Australia"
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
    "company": "OY LANGH SHIP AB",
    "city": "PIIKKIO",
    "country": "Finland"
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
  "LST": {
    "code": "LST",
    "company": "LS INTERTANK APS",
    "city": "FREDERICIA",
    "country": "Denmark"
  },
  "LSX": {
    "code": "LSX",
    "company": "LIGHTSAIL ENERGY INC.",
    "city": "Berkeley",
    "country": "United States"
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
  "LTU": {
    "code": "LTU",
    "company": "MINISTRY OF DEFENSE OF LITHUANIA",
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
    "city": "Punta Gorda",
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
    "company": "FLEX BOX COLUMBIA SAS",
    "city": "BOGOTA",
    "country": "Colombia"
  },
  "LVS": {
    "code": "LVS",
    "company": "LEAVESLEY CONTAINER  SERVICES",
    "city": "BURTON-ON-TRENT",
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
    "company": "CAN LINK INTERNATIONAL LEASE LTD",
    "city": "RICHMOND",
    "country": "Canada"
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
  "MAB": {
    "code": "MAB",
    "company": "MOHAMED ABDULRAHMAN AL-BAHAR METAL CABINETS AND ENCLOSURES MANUFACTURING LLC",
    "city": "Dubai",
    "country": "United Arab Emirates"
  },
  "MAE": {
    "code": "MAE",
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MAG": {
    "code": "MAG",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
    "company": "MAERSK LINE A/S",
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
    "company": "MAS LOGISTICS",
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
  "MAW": {
    "code": "MAW",
    "company": "MAKEWAY PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "MAX": {
    "code": "MAX",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
  "MBJ": {
    "code": "MBJ",
    "company": "MBJV LEASING LTD",
    "city": "CHESHIRE",
    "country": "United Kingdom"
  },
  "MBS": {
    "code": "MBS",
    "company": "M.S.D LEVANT SHIPPING S.L.",
    "city": "GRAO CASTELLON",
    "country": "Spain"
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
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
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
    "city": "Milford",
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
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MCI": {
    "code": "MCI",
    "company": "MAERSK CONTAINER INDUSTRY AS",
    "city": "TINGLEV",
    "country": "Denmark"
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
    "company": "MCC TRANSPORT SINGAPORE PTE LTD",
    "city": "SOUTHPOINT",
    "country": "Singapore"
  },
  "MCR": {
    "code": "MCR",
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MCS": {
    "code": "MCS",
    "company": "STAR CONTAINER SERVICES BV",
    "city": "MAASVLAKTE ROTTERDAM",
    "country": "Netherlands"
  },
  "MCT": {
    "code": "MCT",
    "company": "MARIANA EXPRESS LINES PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "MCX": {
    "code": "MCX",
    "company": "CARGO LEVANT SCHIFFAHRTSGESCHAFT MBH",
    "city": "BREMEN",
    "country": "Germany"
  },
  "MCZ": {
    "code": "MCZ",
    "company": "MINISTRY OF DEFENCE OF THE CZECH REPUBLIC",
    "city": "PRAGUE",
    "country": "Czech Republic"
  },
  "MDE": {
    "code": "MDE",
    "company": "MINISTERE DE LA DEFENSE",
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
    "company": "MMD OFFSHORE ENGINEERING LTD",
    "city": "JIADING",
    "country": "China"
  },
  "MDS": {
    "code": "MDS",
    "company": "QINETIQ TARGET SYSTEMS LIMITED",
    "city": "ASHFORD, KENT",
    "country": "United Kingdom"
  },
  "MDT": {
    "code": "MDT",
    "company": "COTRANSA TRANSPORTE MULTIMODAL LDA",
    "city": "LISBOA",
    "country": "Portugal"
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
    "company": "MELKWEG / FRITOM",
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
    "city": "SINGAPORE",
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
    "city": "SINGAPORE",
    "country": "Singapore"
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
  "MFZ": {
    "code": "MFZ",
    "company": "MAB OILFIELD ENGINEERING & SOLUTIONS FZC",
    "city": "HAMRIYAH FREE ZONE PHASE II",
    "country": "United Arab Emirates"
  },
  "MGG": {
    "code": "MGG",
    "company": "MESSER GROUP GMBH",
    "city": "BAD SODEN",
    "country": "Germany"
  },
  "MGI": {
    "code": "MGI",
    "company": "CO2-REDUCE BV",
    "city": "Valkenswaard",
    "country": "Netherlands"
  },
  "MGL": {
    "code": "MGL",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
  "MGX": {
    "code": "MGX",
    "company": "MCC TIANGONG EQUIPMENT LTD",
    "city": "TIANJIN",
    "country": "China"
  },
  "MHB": {
    "code": "MHB",
    "company": "BATIMPEX KERESKEDELMI KFT",
    "city": "GODOLLO",
    "country": "Hungary"
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
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MHI": {
    "code": "MHI",
    "company": "MITSUBISHI HEAVY INDUSTRIES, LTD.",
    "city": "Nagoya",
    "country": "Japan"
  },
  "MIA": {
    "code": "MIA",
    "company": "MONTEBELLO AG",
    "city": "PONTRESINA",
    "country": "Switzerland"
  },
  "MID": {
    "code": "MID",
    "company": "MIDAMI LIMITED",
    "city": "HONG-KONG",
    "country": "HK"
  },
  "MIE": {
    "code": "MIE",
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MIG": {
    "code": "MIG",
    "company": "MARITIME INTL GROUP LIMITED",
    "city": "MONGKOK, KOWLOON",
    "country": "HK"
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
  "MIT": {
    "code": "MIT",
    "company": "PS SURVEY & CLAIM SERVICES APS",
    "city": "RISSKOV",
    "country": "Denmark"
  },
  "MIX": {
    "code": "MIX",
    "company": "VAPORES SUARDIAZ S.A",
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
    "company": "RWE POWER AG - KRAFTWERK MÜLHEIM-KÄRLICH",
    "city": "MULHEIM-KAERLICH",
    "country": "Germany"
  },
  "MKL": {
    "code": "MKL",
    "company": "MEDKON LINES",
    "city": "İstanbul",
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
  "MLC": {
    "code": "MLC",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
    "company": "MAS LINE MANAGEMENT LTD",
    "city": "London",
    "country": "United Kingdom"
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
  "MMA": {
    "code": "MMA",
    "company": "MAERSK LINE A/S",
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
    "company": "ETS MAX GIRARDIN",
    "city": "ST PIERRE ET MIQUELON",
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
    "city": "ROCKLEDGE",
    "country": "United States"
  },
  "MMX": {
    "code": "MMX",
    "company": "MARINOR ASSOCIATES",
    "city": "Houston",
    "country": "United States"
  },
  "MNB": {
    "code": "MNB",
    "company": "MAERSK LINE A/S",
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
    "city": "KWAI CHUNG, N.T",
    "country": "HK"
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
    "city": "KWAI CHUNG, N.T",
    "country": "HK"
  },
  "MOF": {
    "code": "MOF",
    "company": "MITSUI O.S.K. LINES LTD",
    "city": "KWAI CHUNG, N.T",
    "country": "HK"
  },
  "MOG": {
    "code": "MOG",
    "company": "MITSUI O.S.K. LINES LTD",
    "city": "KWAI CHUNG, N.T",
    "country": "HK"
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
    "city": "KWAI CHUNG, N.T",
    "country": "HK"
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
    "city": "KWAI CHUNG, N.T",
    "country": "HK"
  },
  "MOS": {
    "code": "MOS",
    "company": "MITSUI O.S.K. LINES LTD",
    "city": "KWAI CHUNG, N.T",
    "country": "HK"
  },
  "MOT": {
    "code": "MOT",
    "company": "MITSUI O.S.K. LINES LTD",
    "city": "KWAI CHUNG, N.T",
    "country": "HK"
  },
  "MPA": {
    "code": "MPA",
    "company": "MITSUBISHI POLYCRYSTALLINE SILICON AM CO",
    "city": "Theodore",
    "country": "United States"
  },
  "MPB": {
    "code": "MPB",
    "company": "LLC REFPEREVOZKY",
    "city": "SAINT PETERSBURG",
    "country": "Russian Federation"
  },
  "MPC": {
    "code": "MPC",
    "company": "PT MASAJI PRAYASA CARGO",
    "city": "JAKARTA",
    "country": "Indonesia"
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
  "MRK": {
    "code": "MRK",
    "company": "MAERSK LINE A/S",
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
    "company": "MAERSK LINE A/S",
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
    "company": "MAERSK LINE A/S",
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
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MSG": {
    "code": "MSG",
    "company": "MIN SHENG LINES PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "MSH": {
    "code": "MSH",
    "company": "MALAYSIAN SHIPPING CORP.SENDIRIAN BERHAD",
    "city": "",
    "country": "Malaysia"
  },
  "MSK": {
    "code": "MSK",
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MSL": {
    "code": "MSL",
    "company": "MIN SHENG LINES SHANGHAI LIMITED",
    "city": "SHANGHAI",
    "country": "China"
  },
  "MSM": {
    "code": "MSM",
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
    "city": "DENVER",
    "country": "United States"
  },
  "MSS": {
    "code": "MSS",
    "company": "MSSA",
    "city": "SAINT MARCEL",
    "country": "France"
  },
  "MSU": {
    "code": "MSU",
    "company": "M/S CONTAINERS A/S",
    "city": "RISSKOV",
    "country": "Denmark"
  },
  "MSW": {
    "code": "MSW",
    "company": "MAERSK LINE A/S",
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
  "MTB": {
    "code": "MTB",
    "company": "MULTIBOXX LIMITED",
    "city": "KOWLOON",
    "country": "HK"
  },
  "MTC": {
    "code": "MTC",
    "company": "MERLION HOLDINGS PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "MTE": {
    "code": "MTE",
    "company": "MONSANTO EUROPE N.V.",
    "city": "ANTWERPEN",
    "country": "Belgium"
  },
  "MTG": {
    "code": "MTG",
    "company": "MATHESON TRI-GAS INC",
    "city": "BASKIING RIDGE, NJ 07920",
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
    "company": "METRANS JOINT STOCK CO",
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
    "company": "LARGUS UKRAINE",
    "city": "Kramatorsk, Donetsk region",
    "country": "Ukraine"
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
  "MUS": {
    "code": "MUS",
    "company": "MUSCAT GASES CO SAOG",
    "city": "RYSAYL PC 124",
    "country": "Oman"
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
    "company": "MAERSK LINE A/S",
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
  "MWC": {
    "code": "MWC",
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MWI": {
    "code": "MWI",
    "company": "MASTERANK WAX INC",
    "city": "Pleasanton",
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
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "MWT": {
    "code": "MWT",
    "company": "MILKYWAY CHEMICAL SUPPLY CHAIN SERVICE CO.,LTD",
    "city": "SHANGHAI 201203",
    "country": "China"
  },
  "MXC": {
    "code": "MXC",
    "company": "MAXICON CONTAINER LINE PTE LTD",
    "city": "SINGAPORE",
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
  "MZS": {
    "code": "MZS",
    "company": "MACEDONIAN RAILWAYS-TRASPORT JSC SKOPJE",
    "city": "SKOPJE",
    "country": "Macedonia, the former Yugoslav Republic of"
  },
  "NAB": {
    "code": "NAB",
    "company": "ECST CONTAINER SERVICES & TRADING GMBH",
    "city": "SEEVETAL / BULLENHAUSEN",
    "country": "Germany"
  },
  "NAC": {
    "code": "NAC",
    "company": "NASCO CHEMSOL INTERNATIONAL FZE",
    "city": "Sharjah",
    "country": "United Arab Emirates"
  },
  "NAF": {
    "code": "NAF",
    "company": "NORWEGIAN DEFENCE LOGISTIC ORG - AIR",
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
  "NAV": {
    "code": "NAV",
    "company": "NAVEMAR DE MEXICO",
    "city": "Mexico City",
    "country": "Mexico"
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
  "NBB": {
    "code": "NBB",
    "company": "NEO BEIT",
    "city": "PARIS",
    "country": "France"
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
  "NCS": {
    "code": "NCS",
    "company": "NUCLEAR CARGO + SERVICE GMBH",
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
    "company": "NOR LINES AS",
    "city": "STAVANGER",
    "country": "Norway"
  },
  "NDA": {
    "code": "NDA",
    "company": "INTERNATIONAL NUCLEAR SERVICES LTD",
    "city": "WARRINGTON WA36AS",
    "country": "United Kingdom"
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
    "company": "1701 TIDWELL LLC",
    "city": "HOUSTON, TX-77093",
    "country": "United States"
  },
  "NGF": {
    "code": "NGF",
    "company": "LNGTAINER LTD",
    "city": "HELSINKI",
    "country": "Finland"
  },
  "NGN": {
    "code": "NGN",
    "company": "COOL CARRIERS AB",
    "city": "STOCKHOLM",
    "country": "Sweden"
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
    "company": "YARA PRAXAIR",
    "city": "OSLO",
    "country": "Norway"
  },
  "NHJ": {
    "code": "NHJ",
    "company": "NOVO HORIZONTE JACAREPAGUA IMPORTACAO E EXPORTACAO LTDA",
    "city": "RIO DE JANEIRO",
    "country": "Brazil"
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
  "NIQ": {
    "code": "NIQ",
    "company": "NUUK IMEQ",
    "city": "NUUK",
    "country": "Greenland"
  },
  "NIR": {
    "code": "NIR",
    "company": "NIRINT SHIPPING BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "NIS": {
    "code": "NIS",
    "company": "TECHNOPYR SA",
    "city": "PIRAEUS",
    "country": "Greece"
  },
  "NKB": {
    "code": "NKB",
    "company": "NK CO.LTD",
    "city": "BUSAN",
    "country": "Korea, Republic of"
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
    "city": "White Plains",
    "country": "United States"
  },
  "NLB": {
    "code": "NLB",
    "company": "NIRMA LIMITED",
    "city": "BHAVNAGAR",
    "country": "India"
  },
  "NLL": {
    "code": "NLL",
    "company": "NOBLE CONTAINER LEASING LIMITED",
    "city": "JORDAN, KOWLOON",
    "country": "HK"
  },
  "NLN": {
    "code": "NLN",
    "company": "NOR LINES AS",
    "city": "STAVANGER",
    "country": "Norway"
  },
  "NMA": {
    "code": "NMA",
    "company": "COMMANDER, NAVAL AIR SYSTEMS COMMAND AIR 6.7.6.2",
    "city": "PATUXENT RIVER, MD 20670",
    "country": "United States"
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
  "NMX": {
    "code": "NMX",
    "company": "INNER MONGOLIA NEW ENERGY TRADING CENTER CO., LTD",
    "city": "Erdos",
    "country": "China"
  },
  "NNC": {
    "code": "NNC",
    "company": "CANADIAN ROYALTIES INC",
    "city": "MONTREAL ,QUEBEC",
    "country": "Canada"
  },
  "NOA": {
    "code": "NOA",
    "company": "NORWEGIAN ARMED FORCES",
    "city": "Oslo",
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
  "NOD": {
    "code": "NOD",
    "company": "WECO RORO",
    "city": "RUNGSTED KYST",
    "country": "Denmark"
  },
  "NOI": {
    "code": "NOI",
    "company": "NOLIS SPA",
    "city": "ALGER",
    "country": "Algeria"
  },
  "NOL": {
    "code": "NOL",
    "company": "CMA-CGM",
    "city": "MARSEILLE CEDEX 02",
    "country": "France"
  },
  "NOM": {
    "code": "NOM",
    "company": "SRT ROUND THE WORLD LTD",
    "city": "RAMSEY, ISLE OF MAN, IM8 2LQ",
    "country": "United Kingdom"
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
  "NPR": {
    "code": "NPR",
    "company": "SEA STAR LINE LLC",
    "city": "JACKSONVILLE, FL 32256",
    "country": "United States"
  },
  "NPT": {
    "code": "NPT",
    "company": "NEPTUNE INTERNATIONAL MULTIMODAL TRANSPORT (HK) COMPANY LTD",
    "city": "WANCHAI",
    "country": "HK"
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
    "company": "NRS LOGISTICS SHANGHAI CO LTD",
    "city": "SHANGHAI",
    "country": "China"
  },
  "NRI": {
    "code": "NRI",
    "company": "NATIONAL REFRIGERANTS INC.",
    "city": "PHILADELPHIA, PA 19154",
    "country": "United States"
  },
  "NRK": {
    "code": "NRK",
    "company": "NR COOLING SERVICES B.V.",
    "city": "Numansdorp",
    "country": "Netherlands"
  },
  "NRL": {
    "code": "NRL",
    "company": "FPG RAFFLES PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
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
  "OAL": {
    "code": "OAL",
    "company": "OCEAN AFRICA CONTAINER LINE (PTY) LTD",
    "city": "DURBAN",
    "country": "South Africa"
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
  "OBT": {
    "code": "OBT",
    "company": "JIANGSU O-BEST NEW MATERIALS CO LTD",
    "city": "NANTONG, JIANGSU",
    "country": "China"
  },
  "OCB": {
    "code": "OCB",
    "company": "OCEAN BLUE FINANCE GROUP LIMITED",
    "city": "SHAU KEI WAN",
    "country": "HK"
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
  "OCI": {
    "code": "OCI",
    "company": "ARKEMA INC.",
    "city": "KING OF PRUSSIA, PA 19406",
    "country": "United States"
  },
  "OCL": {
    "code": "OCL",
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
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
    "city": "Mailbox No. 5088",
    "country": "Singapore"
  },
  "OER": {
    "code": "OER",
    "company": "ORIENTAL EQUIPMENT SERVICES INC",
    "city": "Secaucus",
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
    "company": "JSC BEREZKAGAS OB",
    "city": "628011",
    "country": "Russian Federation"
  },
  "OGT": {
    "code": "OGT",
    "company": "ORPHAN GRAIN TRAIN INC.",
    "city": "NORFOLK, NE-68702",
    "country": "United States"
  },
  "OKC": {
    "code": "OKC",
    "company": "CAHAYA SAMUDERA SHIPPING PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "OLT": {
    "code": "OLT",
    "company": "OL&T FOODTRANS LLC",
    "city": "Irvine",
    "country": "United States"
  },
  "OLV": {
    "code": "OLV",
    "company": "OBORONLOGISTICS LLS",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "OMN": {
    "code": "OMN",
    "company": "AIR LIQUID ITALIA SERVICE SRL",
    "city": "MILANO MI",
    "country": "Italy"
  },
  "ONE": {
    "code": "ONE",
    "company": "OCEAN NETWORK EXPRESS PTE. LTD.",
    "city": "Singapore",
    "country": "Singapore"
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
  "OPA": {
    "code": "OPA",
    "company": "OPATRANS",
    "city": "RADZANOWO",
    "country": "Poland"
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
  "OPS": {
    "code": "OPS",
    "company": "MAXAM LIMITED",
    "city": "NAIROBI",
    "country": "Kenya"
  },
  "OPT": {
    "code": "OPT",
    "company": "OPTIMODAL INC",
    "city": "WEST CHESTER, PA 19381",
    "country": "United States"
  },
  "ORA": {
    "code": "ORA",
    "company": "INTERMODAL LOGISTICS LTD",
    "city": "",
    "country": "HK"
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
  "OSM": {
    "code": "OSM",
    "company": "KOPPERS PERFORMANCE CHEMICALS NEW ZEALAND",
    "city": "AUCKLAND",
    "country": "New Zealand"
  },
  "OSR": {
    "code": "OSR",
    "company": "SIEMENS AG",
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
    "company": "OTND ONET TECHNOLOGIES NUCLEAR DECOMMISSIONING",
    "city": "MARSEILLE",
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
    "company": "OXON ITALIA S.P.A.",
    "city": "PERO MI",
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
  "PAK": {
    "code": "PAK",
    "company": "SHARIF OXYGEN (PVT) LTD",
    "city": "LAHORE",
    "country": "Pakistan"
  },
  "PAL": {
    "code": "PAL",
    "company": "PAN ASIA LOGISTICS INDIA PVT LTD",
    "city": "ROYAPETTAH, CHENNAI",
    "country": "India"
  },
  "PAN": {
    "code": "PAN",
    "company": "PANGAS",
    "city": "DAGMERSELLEN",
    "country": "Switzerland"
  },
  "PAP": {
    "code": "PAP",
    "company": "PASTRELLO AUTOTRASPORTI SRL",
    "city": "MARGHERA VE",
    "country": "Italy"
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
    "city": "CALOUNDRA",
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
    "city": "SINGAPORE",
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
    "company": "ICL",
    "city": "CREVE COEUR, MO 63141",
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
  "PEK": {
    "code": "PEK",
    "company": "PEKH CORP",
    "city": "COSTA DEL ESTE",
    "country": "Panama"
  },
  "PEN": {
    "code": "PEN",
    "company": "PENSPEN LTD",
    "city": "RICHMOND",
    "country": "United Kingdom"
  },
  "PEQ": {
    "code": "PEQ",
    "company": "MY MINI CASA",
    "city": "Salt Lake City, UT-84065",
    "country": "United States"
  },
  "PER": {
    "code": "PER",
    "company": "PERTHON GROUP",
    "city": "Gothenburg",
    "country": "Sweden"
  },
  "PET": {
    "code": "PET",
    "company": "PETRONARS",
    "city": "Dubai",
    "country": "United Arab Emirates"
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
  "PGB": {
    "code": "PGB",
    "company": "PERGAN GMBH",
    "city": "BOCHOLT",
    "country": "Germany"
  },
  "PGF": {
    "code": "PGF",
    "company": "PACIFIC GATES CONTAINER LINES LTD",
    "city": "Hong Kong",
    "country": "HK"
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
    "country": "Taiwan (China)"
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
    "company": "FLORENS CONTAINER SERVICES CO LTD",
    "city": "Taipa",
    "country": "Macao"
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
    "company": "POLOMA HOLDING MANAGEMENT LTD",
    "city": "ST PETERSBURG",
    "country": "Russian Federation"
  },
  "PHO": {
    "code": "PHO",
    "company": "PHOENIX GLOBAL FREIGHT GROUP S.A.",
    "city": "PIRAEUS",
    "country": "Greece"
  },
  "PHR": {
    "code": "PHR",
    "company": "PYEONG HWA REEFER SERVICES INC",
    "city": "BUSAN",
    "country": "Korea, Republic of"
  },
  "PHX": {
    "code": "PHX",
    "company": "HUA XING LTD",
    "city": "POM",
    "country": "Papua New Guinea"
  },
  "PIA": {
    "code": "PIA",
    "company": "SAS TRANSPORTS PIALLA",
    "city": "Pierrelatte",
    "country": "France"
  },
  "PIC": {
    "code": "PIC",
    "company": "AUTOTRASPORTI PICCININI S.A.S.",
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
    "city": "SINGAPORE",
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
    "city": "CALLAO 2, LIMA",
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
  "PLS": {
    "code": "PLS",
    "company": "ALTO SHIPPING PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "PLW": {
    "code": "PLW",
    "company": "INSPEKTORAT WSPARCIA SZ U.L DWERNICKIEGO 1",
    "city": "BYDGOSZCZ",
    "country": "Poland"
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
    "company": "PERMA SHIPPING LINE PTE LTD   C/O. MERLION SHIPPING LLC.",
    "city": "Singapore",
    "country": "Singapore"
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
  "PNM": {
    "code": "PNM",
    "company": "PANAMA CARGO LINES LTD INC",
    "city": "PANAMA CITY",
    "country": "Panama"
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
    "company": "MAERSK LINE A/S",
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
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "POR": {
    "code": "POR",
    "company": "PORTEK SYSTEMS & EQUIPMENT PTE LTD",
    "city": "Singapore",
    "country": "Singapore"
  },
  "POU": {
    "code": "POU",
    "company": "PROCONTAINER B.V",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "PPC": {
    "code": "PPC",
    "company": "KANTO-PPC INC",
    "city": "TAOYAN CITY",
    "country": "Taiwan (China)"
  },
  "PPG": {
    "code": "PPG",
    "company": "ALTIVIA SPECIALTY CHEMICALS",
    "city": "Houston",
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
    "city": "AUCKLAND",
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
    "city": "SINGAPORE",
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
    "country": "Taiwan (China)"
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
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "PRT": {
    "code": "PRT",
    "company": "PORTUSLINE CONTAINERS INTERNATIONAL, S.A.",
    "city": "LISBOA",
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
    "company": "RAVENSCROFT HOLDINGS INC",
    "city": "CORA GABLES, FL 33134-7",
    "country": "United States"
  },
  "PSM": {
    "code": "PSM",
    "company": "TRADING LOGISTIC SAC SRL",
    "city": "La Spezia",
    "country": "Italy"
  },
  "PSO": {
    "code": "PSO",
    "company": "PENTALVER TRANSPORT LIMITED",
    "city": "SOUTHAMPTON",
    "country": "United Kingdom"
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
  "PTM": {
    "code": "PTM",
    "company": "PT MATIC ENVIROMENTAL SERVICES LTD",
    "city": "MRIEHEL",
    "country": "Malta"
  },
  "PTR": {
    "code": "PTR",
    "company": "POSTEN NORGE AS",
    "city": "MO I RANA",
    "country": "Norway"
  },
  "PTS": {
    "code": "PTS",
    "company": "XPO LOGISTICS",
    "city": "DUBLIN, OH 43016",
    "country": "United States"
  },
  "PTV": {
    "code": "PTV",
    "company": "PISTON TANK CORPORATION",
    "city": "S.T LOUIS, MO 63026",
    "country": "United States"
  },
  "PTY": {
    "code": "PTY",
    "company": "ACETI OXIGENO S.A",
    "city": "BOCA LA CAJA",
    "country": "Panama"
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
  "PWR": {
    "code": "PWR",
    "company": "POWERTECH INC",
    "city": "MONROVIA",
    "country": "Liberia"
  },
  "PWW": {
    "code": "PWW",
    "company": "PELICAN WORLDWIDE B.V",
    "city": "HEINENOORD",
    "country": "Netherlands"
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
  "QBX": {
    "code": "QBX",
    "company": "CHS CONTAINER HANDEL GMBH",
    "city": "BREMEN",
    "country": "Germany"
  },
  "QCG": {
    "code": "QCG",
    "company": "QUEST CAPITAL GROUP LLC",
    "city": "NORTH KANSAS CITY, MO 64116",
    "country": "United States"
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
  "QSS": {
    "code": "QSS",
    "company": "PT. INDONESIA TSINGSHAN STAINLESS STEEL",
    "city": "JAKARTA",
    "country": "Indonesia"
  },
  "QUB": {
    "code": "QUB",
    "company": "QUBE PORTS PTY LTD",
    "city": "SYDNEY",
    "country": "Australia"
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
    "city": "SAN JOSE",
    "country": "United States"
  },
  "RAF": {
    "code": "RAF",
    "company": "OXITEC S.R.L.",
    "city": "SAN PEDRO DE MACORIS",
    "country": "Dominican Republic"
  },
  "RAH": {
    "code": "RAH",
    "company": "IGL LIMITED",
    "city": "Kingston 11",
    "country": "Jamaica"
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
    "company": "FLEX BOX COLUMBIA SAS",
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
    "company": "BESED LTD",
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
    "city": "SAN LEANDRO",
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
  "RCJ": {
    "code": "RCJ",
    "company": "THE RUM COMPANY (JAMAICA) LTD.",
    "city": "KINGSTON 11",
    "country": "Jamaica"
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
    "company": "REFTRADE  BV",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "REG": {
    "code": "REG",
    "company": "RCL FEEDER PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "REP": {
    "code": "REP",
    "company": "SENVION GMBH",
    "city": "HAMBURG",
    "country": "Germany"
  },
  "RES": {
    "code": "RES",
    "company": "RED EAGLE SHIPPING AGENCIES PVT LTD",
    "city": "NEW DELHI",
    "country": "India"
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
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "RFC": {
    "code": "RFC",
    "company": "FPG RAFFLES PTE LTD",
    "city": "SINGAPORE",
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
    "company": "FPG RAFFLES PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "RFT": {
    "code": "RFT",
    "company": "REFTERMINAL LTD",
    "city": "VLADIVOSTOK",
    "country": "Russian Federation"
  },
  "RGA": {
    "code": "RGA",
    "company": "RENEGADE GAS PTY LTD",
    "city": "Ingleburn",
    "country": "Australia"
  },
  "RGG": {
    "code": "RGG",
    "company": "SPECGASTRANS LLC",
    "city": "Moscow region, Odintsovo district",
    "country": "Russian Federation"
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
    "company": "LLC ISR TRANS",
    "city": "Moscow",
    "country": "Russian Federation"
  },
  "RHH": {
    "code": "RHH",
    "company": "GREAT WALL RUN HENG FINANCIAL LEASING CO., LTD",
    "city": "Shanghai",
    "country": "China"
  },
  "RHL": {
    "code": "RHL",
    "company": "RAWHIDE LEASING COMPANY",
    "city": "Napa",
    "country": "United States"
  },
  "RHN": {
    "code": "RHN",
    "company": "TIS-LOGISTIC LLC",
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
    "company": "RINA INTERMODAL S.R.L",
    "city": "GENOVA",
    "country": "Italy"
  },
  "RIL": {
    "code": "RIL",
    "company": "RICKMERS-LINE GMBH & Co KG",
    "city": "HAMBURG",
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
  "RIX": {
    "code": "RIX",
    "company": "TRITICANTHA LOGISTIC LIMITED",
    "city": "LIMASSOL",
    "country": "Cyprus"
  },
  "RJC": {
    "code": "RJC",
    "company": "FLEX BOX COLUMBIA SAS",
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
    "company": "FPG RAFFLES PTE LTD",
    "city": "SINGAPORE",
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
    "city": "Fort Lauderdale",
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
  "RNC": {
    "code": "RNC",
    "company": "RABIA NOUFI LTD.",
    "city": "Haifa",
    "country": "Israel"
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
  "ROR": {
    "code": "ROR",
    "company": "LION CONTAINERS LTD",
    "city": "WALSALL",
    "country": "United Kingdom"
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
  "ROV": {
    "code": "ROV",
    "company": "ROVA MARITIEM B.V.",
    "city": "ROTTERDAM",
    "country": "Netherlands"
  },
  "ROX": {
    "code": "ROX",
    "company": "EMCO MARINE LTD /  CONTAINERS DPT.",
    "city": "HAIFA",
    "country": "Israel"
  },
  "RPB": {
    "code": "RPB",
    "company": "GVT INTERMODAL B.V.",
    "city": "TILBURG",
    "country": "Netherlands"
  },
  "RPC": {
    "code": "RPC",
    "company": "TIC PETERSBURG LTD",
    "city": "ST PETERSBURG",
    "country": "Russian Federation"
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
    "company": "R-CONTAINER IMPORT AS",
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
    "company": "REFTRADE  BV",
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
    "company": "RS CONTAINER GROUP",
    "city": "RIGA",
    "country": "Latvia"
  },
  "RSL": {
    "code": "RSL",
    "company": "MATSON SOUTH PACIFIC LTD",
    "city": "AUCKLAND",
    "country": "New Zealand"
  },
  "RSS": {
    "code": "RSS",
    "company": "ROYAL WOLF TRADING AUSTRALIA PTY LTD",
    "city": "HORNSBY",
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
    "city": "BERGEN (NH)",
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
    "country": "Taiwan (China)"
  },
  "RTV": {
    "code": "RTV",
    "company": "ROOK TRADING BV",
    "city": "VLAARDINGEN",
    "country": "Netherlands"
  },
  "RUB": {
    "code": "RUB",
    "company": "RUBINO SANTE",
    "city": "FASANO",
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
    "city": "HORNSBY",
    "country": "Australia"
  },
  "RWT": {
    "code": "RWT",
    "company": "ROYAL WOLF TRADING AUSTRALIA PTY LTD",
    "city": "HORNSBY",
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
  "SAH": {
    "code": "SAH",
    "company": "MUE KHABAROVSK \"SPETSAVTOHOZYAYSTVO",
    "city": "KHABAROVSK",
    "country": "Russian Federation"
  },
  "SAI": {
    "code": "SAI",
    "company": "STAR CHEMICAL LOGISTIC SPA",
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
  "SBZ": {
    "code": "SBZ",
    "company": "SOL (BARBADOS) LTD",
    "city": "ST.MICHAEL",
    "country": "Barbados"
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
    "company": "SPECIALTY COATINGS (DARWEN) LTD",
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
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "SCM": {
    "code": "SCM",
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "SCN": {
    "code": "SCN",
    "company": "SEA- CARGO AS",
    "city": "NESTTUN",
    "country": "Norway"
  },
  "SCO": {
    "code": "SCO",
    "company": "STAR CONTAINER SERVICES BV",
    "city": "MAASVLAKTE ROTTERDAM",
    "country": "Netherlands"
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
  "SCR": {
    "code": "SCR",
    "company": "STAR CONTAINER SERVICES BV",
    "city": "MAASVLAKTE ROTTERDAM",
    "country": "Netherlands"
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
  "SDI": {
    "code": "SDI",
    "company": "TREDI",
    "city": "Saint-Vulbas",
    "country": "France"
  },
  "SDM": {
    "code": "SDM",
    "company": "SIGMA SHIPPING & CONTAINER LOGISTICS CO.",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "SDN": {
    "code": "SDN",
    "company": "SEDNA CONTAINERS B.V.",
    "city": "Landsmeer",
    "country": "Netherlands"
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
    "city": "JACKSONVILLE",
    "country": "United States"
  },
  "SEG": {
    "code": "SEG",
    "company": "SEACO SRL",
    "city": "BRIDGETOWN BB11144",
    "country": "Barbados"
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
  "SFC": {
    "code": "SFC",
    "company": "SFCONTAINERS LLC",
    "city": "OAKLAND, CA-94621",
    "country": "United States"
  },
  "SFE": {
    "code": "SFE",
    "company": "SF ENTERPRISES",
    "city": "OAKLAND",
    "country": "United States"
  },
  "SFF": {
    "code": "SFF",
    "company": "ACB AGENCIES",
    "city": "ANTWERPEN",
    "country": "Belgium"
  },
  "SFL": {
    "code": "SFL",
    "company": "SEAFAST LOGISTICS",
    "city": "FELIXSTOWE, SUFFOLK IP11 7SS",
    "country": "United Kingdom"
  },
  "SFN": {
    "code": "SFN",
    "company": "JSC SOVFRAHT- NN",
    "city": "NIZHNI NOVGOROD",
    "country": "Russian Federation"
  },
  "SFR": {
    "code": "SFR",
    "company": "SCANFOR bvba",
    "city": "WOMMELGEM",
    "country": "Belgium"
  },
  "SFU": {
    "code": "SFU",
    "company": "SAN FU CHEMICAL CO, LTD",
    "city": "SHAN-HUA, TAINAN",
    "country": "Taiwan (China)"
  },
  "SGA": {
    "code": "SGA",
    "company": "MINISTERE DE LA DEFENSE",
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
  "SGL": {
    "code": "SGL",
    "company": "SICGIL INDUSTRIAL GASES LTD",
    "city": "CHENNAI - TAMIL NADU",
    "country": "India"
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
    "city": "SINGAPORE",
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
    "city": "CHRISTCHURCH",
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
  "SHD": {
    "code": "SHD",
    "company": "SHHD PRECISION LIMITED",
    "city": "BAOSHAN, SHANGHAI",
    "country": "China"
  },
  "SHE": {
    "code": "SHE",
    "company": "SHL OFFSHORE CONTRACTORS BV",
    "city": "ZOETERMEER",
    "country": "Netherlands"
  },
  "SHH": {
    "code": "SHH",
    "company": "SLOMAN NEPTUNE SCHIFFAHRTS - AG",
    "city": "BREMEN",
    "country": "Germany"
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
    "company": "SINOCHEM INTERNATIONAL CORPORATION",
    "city": "SHANGHAI",
    "country": "China"
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
    "city": "SINGAPORE",
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
    "company": "ELBTAINER TRADING GmbH",
    "city": "HAMBURG",
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
    "city": "BUCKINGHAMSHIRE, SL7 2NL",
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
  "SLJ": {
    "code": "SLJ",
    "company": "SEALEASE N.V.",
    "city": "WILLEMSTAD, CURACAO,",
    "country": "Netherlands Antilles"
  },
  "SLP": {
    "code": "SLP",
    "company": "SEAGULL CONTAINER SERVICES PTE. LTD.",
    "city": "SINGAPORE",
    "country": "Singapore"
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
    "city": "VANCOUVER",
    "country": "Canada"
  },
  "SLZ": {
    "code": "SLZ",
    "company": "CARU SPECIALIZED LEASING",
    "city": "SAN FRANCISCO, CA-94104",
    "country": "United States"
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
    "country": "Taiwan (China)"
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
  "SMU": {
    "code": "SMU",
    "company": "CONTAINERSHIPS PLC (CONTAINERSHIPS OYJ)",
    "city": "Espoo",
    "country": "Finland"
  },
  "SNA": {
    "code": "SNA",
    "company": "SMARTLAGER NORGE AS",
    "city": "LARVIK",
    "country": "Norway"
  },
  "SNB": {
    "code": "SNB",
    "company": "SINOTRANS CONTAINER LINES CO LTD",
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
    "city": "Berlin",
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
    "company": "SINOTRANS CONTAINER LINES CO LTD",
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
    "company": "SOFRANA ANL (NZ) LIMITED AS AGENT FOR SOFRANA ANL PTE LIMITED",
    "city": "AUCKLAND",
    "country": "New Zealand"
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
  "SOV": {
    "code": "SOV",
    "company": "SOR IBERICA S.A",
    "city": "ALZIRA",
    "country": "Spain"
  },
  "SOX": {
    "code": "SOX",
    "company": "AIR LIQUIDE SINGAPORE PRIVATE LIMITED",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "SPB": {
    "code": "SPB",
    "company": "SOUTHWEST CONTAINER SOLUTIONS",
    "city": "Scottsboro",
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
    "company": "SPEEDIC EQUIPMENT SERVICES CORP.",
    "city": "TORTOLA",
    "country": "Virgin Islands, British"
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
  "SPS": {
    "code": "SPS",
    "company": "CONTAINERS COMPLETE LIMITED T/A SPACEWISE",
    "city": "WELLINGTON",
    "country": "New Zealand"
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
    "company": "SRA SAVAC",
    "city": "BOLLENE",
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
    "company": "OJSC SURGUTNEFTEGAS",
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
  "SST": {
    "code": "SST",
    "company": "SATLINES SHIPPING LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "SSU": {
    "code": "SSU",
    "company": "SERVICELINE LTD",
    "city": "Vladivostok",
    "country": "Russian Federation"
  },
  "SSW": {
    "code": "SSW",
    "company": "GUIZHOU T.M.S LOGISTIC CO.LTD",
    "city": "GUIYANG CITY",
    "country": "China"
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
  "STG": {
    "code": "STG",
    "company": "SHOWA SPECIALTY GAS (TAIWAN) CO LTD",
    "city": "TAIPEI",
    "country": "Taiwan (China)"
  },
  "STJ": {
    "code": "STJ",
    "company": "SCHENKER INC.",
    "city": "Freeport",
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
  "STW": {
    "code": "STW",
    "company": "SETTENTRIONALE TRASPORTI SPA",
    "city": "POSSAGNO (TV)",
    "country": "Italy"
  },
  "STX": {
    "code": "STX",
    "company": "STAXXON LLC",
    "city": "MIDDLETOWN, OH-45044",
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
    "company": "HAMBURG SUED",
    "city": "HAMBURG",
    "country": "Germany"
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
  "SUP": {
    "code": "SUP",
    "company": "HB RENTALS",
    "city": "Broussard",
    "country": "United States"
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
    "company": "SAMSKIP VAN DIEREN MULTIMODAL BV",
    "city": "GENEMUIDEN",
    "country": "Netherlands"
  },
  "SVK": {
    "code": "SVK",
    "company": "SUOMEN VUOKRAKONTTI OY",
    "city": "RAJAMAKI",
    "country": "Finland"
  },
  "SVL": {
    "code": "SVL",
    "company": "SAVEL ELECTRONIC",
    "city": "GOLCUK-KOCAELI",
    "country": "Turkey"
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
    "company": "BULKCON TRANSPORT AB",
    "city": "GOTHENBURG",
    "country": "Sweden"
  },
  "SWF": {
    "code": "SWF",
    "company": "SWIFT TRANSPORT INTERNATIONAL LOGISTICS PTE .LTD",
    "city": "TIANJIN",
    "country": "China"
  },
  "SWG": {
    "code": "SWG",
    "company": "SUNSTONE WATER GROUP EUROPE APS",
    "city": "Vejle",
    "country": "Denmark"
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
    "city": "SINGAPORE",
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
    "city": "",
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
  "TBU": {
    "code": "TBU",
    "company": "TRUE BLUE CONTAINERS (2005) PTY LTD",
    "city": "MIDVACE",
    "country": "Australia"
  },
  "TCA": {
    "code": "TCA",
    "company": "BLUE BALTIC SHIPPING & TRADING LIMITED",
    "city": "ROAD TOWN, TORTOLA",
    "country": "Virgin Islands, British"
  },
  "TCC": {
    "code": "TCC",
    "company": "TANK CONTAINER CLEANING SERVICES S.R.L.",
    "city": "NAPOLI",
    "country": "Italy"
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
  "TCP": {
    "code": "TCP",
    "company": "TCOM L.P.",
    "city": "Columbia",
    "country": "United States"
  },
  "TCQ": {
    "code": "TCQ",
    "company": "TRANSPORT CANADA",
    "city": "WINNIPEG, MB",
    "country": "Canada"
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
  "TDI": {
    "code": "TDI",
    "company": "TORANG DARYA SHIPPING CO LTD",
    "city": "TEHRAN",
    "country": "Iran, Islamic Republic of"
  },
  "TDR": {
    "code": "TDR",
    "company": "TAL INTERNATIONAL",
    "city": "PURCHASE, NY 10577-2135",
    "country": "United States"
  },
  "TDS": {
    "code": "TDS",
    "company": "THEATRE DU SOLEIL",
    "city": "PARIS",
    "country": "France"
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
    "company": "PT PELAYARAN TEMPURAN EMAS,TBK",
    "city": "JAKARTA UTARA",
    "country": "Indonesia"
  },
  "TEL": {
    "code": "TEL",
    "company": "TRANS EATON INTERNATIONAL ASSET MANAGEMENT COMPANY LIMITED",
    "city": "Hong Kong",
    "country": "HK"
  },
  "TEM": {
    "code": "TEM",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
  "TEX": {
    "code": "TEX",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
    "city": "WARRINGTON",
    "country": "United Kingdom"
  },
  "TFL": {
    "code": "TFL",
    "company": "TOP FLEET",
    "city": "Katy",
    "country": "United States"
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
    "city": "SINGAPORE",
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
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
    "country": "Taiwan (China)"
  },
  "TGH": {
    "code": "TGH",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
    "company": "TRANSGULF LLC",
    "city": "CLEWISTON",
    "country": "United States"
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
  "THI": {
    "code": "THI",
    "company": "THIELMANN FINANCIAL SOLUTIONS AG",
    "city": "Zug",
    "country": "Switzerland"
  },
  "THL": {
    "code": "THL",
    "company": "PAN-ASIA CONTAINER SERVICE CO.LTD",
    "city": "Hong Kong",
    "country": "HK"
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
    "company": "ARMECH MECHANICAL LTD (TECTAINER/TEC)",
    "city": "BURGESS HILL, WEST SUSSEX",
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
  "TIX": {
    "code": "TIX",
    "company": "T.I.S CONTAINER LTD",
    "city": "Hong Kong",
    "country": "HK"
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
    "city": "Cold Spring",
    "country": "United States"
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
  "TLT": {
    "code": "TLT",
    "company": "TRADEMARK LEASING & TRADING BV",
    "city": "HOOGVLIET (RI)",
    "country": "Netherlands"
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
  "TML": {
    "code": "TML",
    "company": "TAYLOR MINSTER LEASING BV",
    "city": "SPYKENISSE",
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
    "company": "TRANSTANK PTY LTD",
    "city": "KEILOR PARK",
    "country": "Australia"
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
    "company": "AREVA TN INTERNATIONAL",
    "city": "MONTIGNY-LE-BRETONNEUX",
    "country": "France"
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
  "TOC": {
    "code": "TOC",
    "company": "TAIPEI OXYGEN & GAS CO.",
    "city": "New Taipei City",
    "country": "Taiwan (China)"
  },
  "TOD": {
    "code": "TOD",
    "company": "TTS LLC",
    "city": "FRISCO, TX 75033",
    "country": "United States"
  },
  "TOE": {
    "code": "TOE",
    "company": "TRANSOCEAN EQUIPMENT MANAGEMENT LLC",
    "city": "Shallotte NC 28459",
    "country": "United States"
  },
  "TOG": {
    "code": "TOG",
    "company": "TOL GASES LIMITED",
    "city": "DAR ES SALAAM",
    "country": "Tanzania, United Republic of"
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
    "city": "Kallo",
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
    "company": "MAERSK LINE A/S",
    "city": "Copenhagen",
    "country": "Denmark"
  },
  "TOS": {
    "code": "TOS",
    "company": "TOSS AS",
    "city": "DRAMMEN",
    "country": "Norway"
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
    "country": "Taiwan (China)"
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
    "country": "Taiwan (China)"
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
    "city": "TIFTON, GA-31793",
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
    "country": "Taiwan (China)"
  },
  "TST": {
    "code": "TST",
    "company": "T.S. LINES LTD",
    "city": "TAIPEI",
    "country": "Taiwan (China)"
  },
  "TSU": {
    "code": "TSU",
    "company": "TOMAS SANCHEZ TRANSPORTES CISTERNAS SL",
    "city": "SANTA MARTA",
    "country": "Spain"
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
    "company": "EQUIPPED4U",
    "city": "MOERDIJK",
    "country": "Netherlands"
  },
  "TTC": {
    "code": "TTC",
    "company": "ATLANTIC TRANS CONTAINERS",
    "city": "BANNALEC",
    "country": "France"
  },
  "TTE": {
    "code": "TTE",
    "company": "TEAMTEC AS",
    "city": "TVEDESTRAND",
    "country": "Norway"
  },
  "TTK": {
    "code": "TTK",
    "company": "TRISTAR TRANSPORT LLC",
    "city": "Dubai",
    "country": "United Arab Emirates"
  },
  "TTL": {
    "code": "TTL",
    "company": "TSINLIEN TRANSPORTATION CO. LTD.",
    "city": "Hong Kong",
    "country": "HK"
  },
  "TTN": {
    "code": "TTN",
    "company": "TRITON INTERNATIONAL LTD",
    "city": "PURCHASE, NY 10577",
    "country": "United States"
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
  "TTU": {
    "code": "TTU",
    "company": "TRANSTEK LTD",
    "city": "LYTKARINO TOWN-MOSCOW REGION",
    "country": "Russian Federation"
  },
  "TUC": {
    "code": "TUC",
    "company": "TUCABI CONTAINER S.L.",
    "city": "ERANDIO",
    "country": "Spain"
  },
  "TUF": {
    "code": "TUF",
    "company": "Transuniverse Forwarding NV",
    "city": "Wondelgem",
    "country": "Belgium"
  },
  "TUL": {
    "code": "TUL",
    "company": "CONTAINER SOLUTIONS S.A DE C.V",
    "city": "MEXICO D.F",
    "country": "Mexico"
  },
  "TVL": {
    "code": "TVL",
    "company": "JERMYN CORPORATION - TAVRIA LINE",
    "city": "DNIEPROPETROVSK",
    "country": "Ukraine"
  },
  "TVS": {
    "code": "TVS",
    "company": "TRANSVISION SHIPPING PTE LTD",
    "city": "NAVI MUMBAI",
    "country": "India"
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
    "company": "GEHOLD BV",
    "city": "BREDA",
    "country": "Netherlands"
  },
  "TWS": {
    "code": "TWS",
    "company": "TWS TANKCONTAINER-LEASING GMBH & CO KG",
    "city": "Hamburg",
    "country": "Germany"
  },
  "TXG": {
    "code": "TXG",
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
    "city": "HAMILTON, HM HX",
    "country": "Bermuda"
  },
  "TXX": {
    "code": "TXX",
    "company": "TRIDENT SEAFOODS CORP",
    "city": "SEATTLE, WA 98107-4000",
    "country": "United States"
  },
  "TYR": {
    "code": "TYR",
    "company": "CRS REFRIGERATION LTD",
    "city": "MEATH",
    "country": "Ireland"
  },
  "TYS": {
    "code": "TYS",
    "company": "TAI YOUNG HIGH TECH CO LTD",
    "city": "HSIN-CHU",
    "country": "Taiwan (China)"
  },
  "TZK": {
    "code": "TZK",
    "company": "SUMISEI TAIWAN TECHNOLOGY CO LTD",
    "city": "CHANGHUA COUNTY",
    "country": "Taiwan (China)"
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
  "UCO": {
    "code": "UCO",
    "company": "UCO KAMPEN",
    "city": "Kampen",
    "country": "Netherlands"
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
  "UGM": {
    "code": "UGM",
    "company": "GREENCOMPASS MARINE S.A.",
    "city": "",
    "country": "Panama"
  },
  "UGP": {
    "code": "UGP",
    "company": "JSC BEREZKAGAS UGRA",
    "city": "Khanty-Mansiysk, KhMAO-Yugra",
    "country": "Russian Federation"
  },
  "UHM": {
    "code": "UHM",
    "company": "HIMMAGISTRAL LTD",
    "city": "UGREN",
    "country": "Russian Federation"
  },
  "UII": {
    "code": "UII",
    "company": "LLC UKRAINIAN IMPULS INDUSTRIES",
    "city": "KIEV",
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
    "company": "ACE ENGINEERING & CO LTD",
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
    "company": "UNITEAM AS",
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
  "USE": {
    "code": "USE",
    "company": "U.S. EMBASSY - DEPARTMENT OF STATE",
    "city": "Baghdad",
    "country": "Iraq"
  },
  "USF": {
    "code": "USF",
    "company": "UNITED STATES AIR FORCE",
    "city": "WRIGHT-PATTERSON AFB, OH-45433",
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
  "UUJ": {
    "code": "UUJ",
    "company": "TNO",
    "city": "Delft",
    "country": "Netherlands"
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
  "VAC": {
    "code": "VAC",
    "company": "JSC VOSTOK AVIATION COMPANY",
    "city": "KHABAROVSK",
    "country": "Russian Federation"
  },
  "VAG": {
    "code": "VAG",
    "company": "VA GROUP LTD.",
    "city": "MOSCOW",
    "country": "Russian Federation"
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
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "VAT": {
    "code": "VAT",
    "company": "NEELE VAT LOGISTICS B.V",
    "city": "ROTTERDAM",
    "country": "Netherlands"
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
  "VCI": {
    "code": "VCI",
    "company": "VAN HATTUM EN BLANKEVOORT",
    "city": "Vianen",
    "country": "Netherlands"
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
    "company": "SAMSKIP VAN DIEREN MULTIMODAL BV",
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
    "company": "VERSUM MATERIALS",
    "city": "Allentown",
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
  "VGL": {
    "code": "VGL",
    "company": "VASCO GLOBAL MARITIME LLC",
    "city": "BUR DUBAI",
    "country": "United Arab Emirates"
  },
  "VHT": {
    "code": "VHT",
    "company": "L & T B.V.",
    "city": "AMSTERDAM",
    "country": "Netherlands"
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
    "city": "Mobile",
    "country": "United States"
  },
  "VIT": {
    "code": "VIT",
    "company": "VERSTEIJNEN'S INTERNATIONAAL TRANSPORTBEDRIJF BV",
    "city": "TILBURG",
    "country": "Netherlands"
  },
  "VLG": {
    "code": "VLG",
    "company": "VLS-GROUP",
    "city": "Gent",
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
  "VLR": {
    "code": "VLR",
    "company": "TERMINAL VLRP LTD",
    "city": "LENSK",
    "country": "Russian Federation"
  },
  "VML": {
    "code": "VML",
    "company": "VASCO MARITIME PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
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
    "company": "VOLVO OCEAN RACE SLU",
    "city": "ALICANTE",
    "country": "Spain"
  },
  "VPL": {
    "code": "VPL",
    "company": "VAN PLUS CORPORATION",
    "city": "SEOUL",
    "country": "Korea, Republic of"
  },
  "VPN": {
    "code": "VPN",
    "company": "VICKERY CONTAINER SALES LTD",
    "city": "SOUTHAMPTON S030 ZHE",
    "country": "United Kingdom"
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
  "VRK": {
    "code": "VRK",
    "company": "ANDRADE & SANTOS LOCACAO DE MODULOS E IMPORTAÇÃO LTDA",
    "city": "SANTOS",
    "country": "Brazil"
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
    "city": "VERTOLAYE",
    "country": "France"
  },
  "VSB": {
    "code": "VSB",
    "company": "V.S & B CONTAINERS LLC",
    "city": "DUBAI",
    "country": "United Arab Emirates"
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
  "WAC": {
    "code": "WAC",
    "company": "AIR LIQUIDE NIGERIA PLC",
    "city": "LAGOS",
    "country": "Nigeria"
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
  "WAP": {
    "code": "WAP",
    "company": "TOTAL STORAGE LTD",
    "city": "PARK RIDGE, NJ 07656",
    "country": "United States"
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
    "city": "SOUTHAMPTON",
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
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
    "city": "SEEVETAL",
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
  "WEI": {
    "code": "WEI",
    "company": "WEIR POWER & INDUSTRIAL FRANCE",
    "city": "SAINT VICTORET",
    "country": "France"
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
  "WEW": {
    "code": "WEW",
    "company": "WEW CONTAINER SYSTEMS GMBH",
    "city": "WEITEFELD",
    "country": "Germany"
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
    "country": "Taiwan (China)"
  },
  "WHS": {
    "code": "WHS",
    "company": "WAN HAI LINES LTD",
    "city": "TAIPEI",
    "country": "Taiwan (China)"
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
    "company": "HONEYTAK CONTAINER LIMITED",
    "city": "Yantai",
    "country": "China"
  },
  "WIN": {
    "code": "WIN",
    "company": "WINCHESTER ASSET MANAGEMENT PTY LTD",
    "city": "DURBAN",
    "country": "South Africa"
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
    "country": "Taiwan (China)"
  },
  "WLN": {
    "code": "WLN",
    "company": "WALLENIUS WILHELMSEN LOGISTICS AS",
    "city": "LYSAKER",
    "country": "Norway"
  },
  "WMB": {
    "code": "WMB",
    "company": "WHITE MARTINS GASES INDUSTRIALS LTD",
    "city": "SAO PAULO",
    "country": "Brazil"
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
    "city": "HONGKONG",
    "country": "HK"
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
  "WTL": {
    "code": "WTL",
    "company": "WHITE LINE SHIPPING",
    "city": "DUBAI",
    "country": "United Arab Emirates"
  },
  "WTR": {
    "code": "WTR",
    "company": "WORTHINGTON ARITAS BASINGLI KAPLAR SAN.A.S",
    "city": "ISTANBUL",
    "country": "Turkey"
  },
  "WTS": {
    "code": "WTS",
    "company": "WEIGAND-GRUNDSTUCKS GMBH & CO. KG",
    "city": "LENGENBOSTEL",
    "country": "Germany"
  },
  "WVP": {
    "code": "WVP",
    "company": "ALNAMAA CAPITAL HOLDINGS LIMITED",
    "city": "RIYADH",
    "country": "Saudi Arabia"
  },
  "WWL": {
    "code": "WWL",
    "company": "WALLENIUS WILHELMSEN LOGISTICS AS",
    "city": "LYSAKER",
    "country": "Norway"
  },
  "WWW": {
    "code": "WWW",
    "company": "WIND CONTAINER SERVICES LIMITED",
    "city": "IPSWICH, IP8 3IY",
    "country": "United Kingdom"
  },
  "XAC": {
    "code": "XAC",
    "company": "ALPS CONTAINER",
    "city": "FREILASSING",
    "country": "Germany"
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
    "company": "TEXTAINER EQUIPMENT MANAGEMENT LTD-",
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
  "XSL": {
    "code": "XSL",
    "company": "NSR ASSET LIMITED",
    "city": "WANCHAI, HONG KONG",
    "country": "HK"
  },
  "XTD": {
    "code": "XTD",
    "company": "YIJINHUOLUO XINGTAIDA LOGISTICS CO. LTD",
    "city": "Ordos",
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
  "YBF": {
    "code": "YBF",
    "company": "SPECIALTY MOVERS LLC",
    "city": "Ann Arbor",
    "country": "United States"
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
    "country": "Taiwan (China)"
  },
  "YFC": {
    "code": "YFC",
    "company": "EVER FORTUNE LOGISTICS (HK) CO LTD",
    "city": "HUNG HOM-KOWLOON",
    "country": "HK"
  },
  "YGM": {
    "code": "YGM",
    "company": "SHANGHAI YUZHOU LINE LTD",
    "city": "PUDONG",
    "country": "China"
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
  "YLU": {
    "code": "YLU",
    "company": "YILIAN LOGISTICS (SHANGHAI) CO LTD",
    "city": "Shanghai",
    "country": "China"
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
    "country": "Taiwan (China)"
  },
  "YMM": {
    "code": "YMM",
    "company": "YANG MING MARINE TRANSPORT CORP.",
    "city": "KEELUNG",
    "country": "Taiwan (China)"
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
  "YYC": {
    "code": "YYC",
    "company": "QUANZHOU YIYANG CONTAINER SERVICE CO.LTD",
    "city": "QUANZHOU CITY",
    "country": "China"
  },
  "ZAM": {
    "code": "ZAM",
    "company": "ZAMBELLI RIB-ROOF GMBH & CO KG",
    "city": "Stephansposching",
    "country": "Germany"
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
  "ZGX": {
    "code": "ZGX",
    "company": "FPG RAFFLES PTE LTD",
    "city": "SINGAPORE",
    "country": "Singapore"
  },
  "ZHF": {
    "code": "ZHF",
    "company": "YANTAI ZHONG HAN FERRY CO LTD",
    "city": "YANTAI",
    "country": "China"
  },
  "ZHR": {
    "code": "ZHR",
    "company": "F.U.H TRANS-STAR S.C ZBIGNIEW, HALINA ROKOSZ",
    "city": "OCIEKA",
    "country": "Poland"
  },
  "ZHW": {
    "code": "ZHW",
    "company": "TIANJIN ZHAOHUA PIONEER CO,LTD",
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
  "ZSE": {
    "code": "ZSE",
    "company": "ZHENG SUN ENGINEERING CO,LTD",
    "city": "KAOHSIUNG CITY",
    "country": "Taiwan (China)"
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
}
`
