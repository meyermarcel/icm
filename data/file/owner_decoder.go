package file

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	// Needed for package embed.
	_ "embed"

	"github.com/meyermarcel/icm/cont"
	"github.com/meyermarcel/icm/data"
)

const (
	csvSep             = ';'
	csvFieldsPerRecord = 4
)

//go:embed owner.csv
var ownerCSV []byte

type owner struct {
	Company string
	City    string
	Country string
}

type ownerDecoder struct {
	owners map[string]owner
}

// NewOwnerDecoder writes owner file to path if it not exists and
// returns a struct that uses this file as a data source.
func NewOwnerDecoder(remoteOwnersPath, customOwnersPath string) (data.OwnerDecoder, error) {
	decoder := &ownerDecoder{}

	if err := initFile(remoteOwnersPath, ownerCSV); err != nil {
		return nil, err
	}
	ownersFile, err := os.Open(remoteOwnersPath)
	if err != nil {
		return nil, err
	}

	ownersMap, err := readCSV(ownersFile)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", remoteOwnersPath, err)
	}

	if len(ownersMap) == 0 {
		return nil, fmt.Errorf("%v: no owners found", remoteOwnersPath)
	}

	if _, err := os.Stat(customOwnersPath); err == nil {

		customOwnersFile, err := os.Open(customOwnersPath)
		if err != nil {
			return nil, err
		}

		customOwnersMap, err := readCSV(customOwnersFile)
		if err != nil {
			return nil, fmt.Errorf("%v: %w", customOwnersPath, err)
		}

		for k, v := range customOwnersMap {
			ownersMap[k] = v
		}
	}

	decoder.owners = ownersMap

	return decoder, nil
}

func readCSV(r io.Reader) (map[string]owner, error) {
	csvReader := csv.NewReader(r)

	csvReader.Comma = csvSep
	csvReader.FieldsPerRecord = csvFieldsPerRecord

	ownersMap := make(map[string]owner)

	for {
		rec, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, err
		}

		ownerCode := rec[0]

		if err := cont.IsOwnerCode(ownerCode); err != nil {
			return nil, err
		}

		ownersMap[ownerCode] = owner{
			Company: rec[1],
			City:    rec[2],
			Country: rec[3],
		}
	}

	return ownersMap, nil
}

// Decode returns an owner for an owner code.
func (od *ownerDecoder) Decode(code string) (bool, cont.Owner) {
	if val, ok := od.owners[code]; ok {
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
func (od *ownerDecoder) GetAllOwnerCodes() []string {
	var codes []string
	for ownerCode := range od.owners {
		codes = append(codes, ownerCode)
	}
	return codes
}

func initFile(path string, content []byte) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.WriteFile(path, content, 0o644); err != nil {
			return err
		}
	}
	return nil
}
