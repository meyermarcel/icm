package file

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	// Needed for package embed.
	_ "embed"

	"github.com/meyermarcel/icm/cont"
	"github.com/meyermarcel/icm/data"
)

const (
	remoteOwnersFileName = "owner.csv"
	customOwnersFileName = "custom-owner.csv"
	csvSep               = ';'
	csvFieldsPerRecord   = 4
)

//go:embed owner.csv
var ownerCSV []byte

type owner struct {
	Company string
	City    string
	Country string
}

type ownerDecoderUpdater struct {
	owners           map[string]owner
	remoteOwnersPath string
}

// NewOwnerDecoderUpdater writes owner file to path if it not exists and
// returns a struct that uses this file as a data source.
func NewOwnerDecoderUpdater(path string) (data.OwnerDecodeUpdater, error) {
	decoderUpdater := &ownerDecoderUpdater{
		remoteOwnersPath: filepath.Join(path, remoteOwnersFileName),
	}

	if err := initFile(decoderUpdater.remoteOwnersPath, ownerCSV); err != nil {
		return nil, err
	}
	ownersFile, err := os.Open(decoderUpdater.remoteOwnersPath)
	if err != nil {
		return nil, err
	}

	ownersMap, err := readCSV(ownersFile)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", decoderUpdater.remoteOwnersPath, err)
	}

	customOwnersPath := filepath.Join(path, customOwnersFileName)

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

	decoderUpdater.owners = ownersMap

	return decoderUpdater, nil
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
	for ownerCode := range of.owners {
		codes = append(codes, ownerCode)
	}
	return codes
}

// Update accepts a slice of owners and overwrites local owner.csv file.
func (of *ownerDecoderUpdater) Update(newOwners []cont.Owner) error {
	file, err := os.OpenFile(of.remoteOwnersPath, os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = csvSep

	for _, o := range newOwners {
		csvErr := csvWriter.Write([]string{o.Code, o.Company, o.City, o.Country})
		if csvErr != nil {
			return csvErr
		}
	}

	csvWriter.Flush()
	if csvWriter.Error() != nil {
		return err
	}

	return nil
}

func initFile(path string, content []byte) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.WriteFile(path, content, 0o644); err != nil {
			return err
		}
	}
	return nil
}
