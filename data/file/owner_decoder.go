package file

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	// needed for package embed
	_ "embed"

	"github.com/meyermarcel/icm/cont"
	"github.com/meyermarcel/icm/data"
)

const (
	ownerFileName      = "owner.csv"
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

// NewOwnerDecoderUpdater writes owner file to path if it not exists and
// returns a struct that uses this file as a data source.
func NewOwnerDecoderUpdater(path string) (data.OwnerDecodeUpdater, error) {
	ownersFile := &ownerDecoderUpdater{path: filepath.Join(path, ownerFileName)}
	if err := initFile(ownersFile.path, ownerCSV); err != nil {
		return nil, err
	}
	f, err := os.Open(ownersFile.path)
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(f)

	csvReader.Comma = csvSep
	csvReader.FieldsPerRecord = csvFieldsPerRecord

	ownersMap := make(map[string]owner)

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if errors.Is(err, csv.ErrFieldCount) {
			return nil, fmt.Errorf("%v: %w", ownersFile.path, err)
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

	ownersFile.owners = ownersMap

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
	for ownerCode := range of.owners {
		codes = append(codes, ownerCode)
	}
	return codes
}

// Update accepts a map of owner code to owner and replaces/adds entries in the local owner file.
// Cancelled owners still exist to prevent removal of custom owners created by the user.
func (of *ownerDecoderUpdater) Update(newOwners []cont.Owner) error {
	for _, o := range newOwners {
		of.owners[o.Code] = owner{
			Company: o.Company,
			City:    o.City,
			Country: o.Country,
		}
	}

	file, err := os.OpenFile(of.path, os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = csvSep

	codes := make([]string, 0, len(of.owners))
	for k := range of.owners {
		codes = append(codes, k)
	}
	sort.Strings(codes)

	for _, code := range codes {
		o := of.owners[code]
		csvErr := csvWriter.Write([]string{code, o.Company, o.City, o.Country})
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
