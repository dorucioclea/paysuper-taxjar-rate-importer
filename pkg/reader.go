package taxjar

import (
	"encoding/csv"
	"io"
	"os"
)

// Record defines one entry from https://simplemaps.com/data/us-zips file with us zip codes.
type Record struct {
	Zip    string
	City   string
	State  string
	County string
}

const (
	Zip    = 0
	City   = 3
	State  = 5
	County = 11
)

func readZipCodeFile(file string) ([]*Record, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var locations []*Record

	reader := csv.NewReader(f)
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return locations, err
		}

		if row[0] == "zip" {
			// skip header
			continue
		}

		locations = append(locations, &Record{
			Zip:    row[Zip],
			City:   row[City],
			State:  row[State],
			County: row[County],
		})
	}
}
