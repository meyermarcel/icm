package file

import (
	"encoding/json"
	"os"
	"path/filepath"

	// needed for package embed
	_ "embed"

	"github.com/meyermarcel/icm/cont"
	"github.com/meyermarcel/icm/data"
)

const equipCatIDsFileName = "equipment-category-id.json"

//go:embed equipment-category-id.json
var equipCatIDsJSON []byte

// NewEquipCatDecoder writes equipment category ID file to path if it not exists and
// returns a struct that uses this file as a data source.
func NewEquipCatDecoder(path string) (data.EquipCatDecoder, error) {
	equipCat := &equipCatDecoder{}
	pathToEquipCat := filepath.Join(path, equipCatIDsFileName)
	if err := initFile(pathToEquipCat, equipCatIDsJSON); err != nil {
		return nil, err
	}
	b, err := os.ReadFile(pathToEquipCat)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &equipCat.categories); err != nil {
		return nil, err
	}
	for ID := range equipCat.categories {
		if err := cont.IsEquipCatID(ID); err != nil {
			return nil, err
		}
	}
	return equipCat, err
}

type equipCatDecoder struct {
	categories map[string]string
}

// Decode decodes ID to equipment category ID.
func (ec *equipCatDecoder) Decode(ID string) (bool, cont.EquipCat) {
	if val, ok := ec.categories[ID]; ok {
		return true, cont.NewEquipCatID(ID, val)
	}
	return false, cont.EquipCat{}
}

// AllCatIDs returns all equipment category IDs.
func (ec *equipCatDecoder) AllCatIDs() []string {
	keys := make([]string, 0, len(ec.categories))
	for k := range ec.categories {
		keys = append(keys, k)
	}
	return keys
}
