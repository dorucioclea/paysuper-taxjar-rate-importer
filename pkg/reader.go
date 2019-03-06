package taxjar

import (
	"encoding/csv"
	"go.uber.org/zap"
	"io"
	"os"
)

const (
	Zip    = 0
	City   = 3
	State  = 5
	County = 11
)

// Record defines one entry from https://simplemaps.com/data/us-zips file with us zip codes.
type Record struct {
	Zip    string
	City   string
	State  string
	County string
}

func readZipCodeFile(file string) ([]*Record, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			zap.L().Error("The csv file with zip code was closed with error", zap.Error(err))
		}
	}()

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
