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

const sizeFileName = "size.json"

//go:embed size.json
var lengthHeightWidthJSON []byte

type size struct {
	Length      map[string]string      `json:"length"`
	HeightWidth map[string]heightWidth `json:"heightWidth"`
}

type heightWidth struct {
	Width  string `json:"height"`
	Height string `json:"width"`
}

// NewSizeDecoder writes last update lengths, height and width file to path if it not exists and
// returns a struct that uses this file as a data source.
func NewSizeDecoder(path string) (data.LengthDecoder, data.HeightWidthDecoder, error) {
	pathToSizes := filepath.Join(path, sizeFileName)
	if err := initFile(pathToSizes, lengthHeightWidthJSON); err != nil {
		return nil, nil, err
	}
	b, err := os.ReadFile(pathToSizes)
	if err != nil {
		return nil, nil, err
	}

	var size size
	if err := json.Unmarshal(b, &size); err != nil {
		return nil, nil, err
	}
	for lengthCode := range size.Length {
		if err := cont.IsLengthCode(lengthCode); err != nil {
			return nil, nil, err
		}
	}
	for heightWidthCode := range size.HeightWidth {
		if err := cont.IsHeightWidthCode(heightWidthCode); err != nil {
			return nil, nil, err
		}
	}
	return &lengthDecoder{size.Length}, &heightWidthDecoder{size.HeightWidth}, nil
}

type lengthDecoder struct {
	lengths map[string]string
}

// Decode returns length for a given length code.
func (l *lengthDecoder) Decode(code string) (bool, cont.Length) {
	if val, ok := l.lengths[code]; ok {
		return true, cont.Length(val)
	}
	return false, ""
}

type heightWidthDecoder struct {
	heightWidths map[string]heightWidth
}

// Decode returns height and width for given height and width code.
func (hw *heightWidthDecoder) Decode(code string) (bool, cont.Height, cont.Width) {
	if val, ok := hw.heightWidths[code]; ok {
		return true, cont.Height(val.Height), cont.Width(val.Width)
	}
	return false, "", ""
}
