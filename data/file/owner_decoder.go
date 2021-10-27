package file

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"

	// needed for package embed
	_ "embed"

	"github.com/meyermarcel/icm/cont"
	"github.com/meyermarcel/icm/data"
)

const ownerFileName = "owner.json"

//go:embed owner.json
var ownerJSON []byte

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
	if err := initFile(filePath, ownerJSON); err != nil {
		return nil, err
	}
	b, err := os.ReadFile(filePath)
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
	return os.WriteFile(filepath.Join(of.path, ownerFileName), b, 0644)
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
		if err := os.WriteFile(path, content, 0644); err != nil {
			return err
		}
	}
	return nil
}
