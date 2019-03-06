package taxjar

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"gopkg.in/resty.v1"
	"math"
	"sync"
)

// Response define root object from TaxJar rate api endpoint
type Response struct {
	Rate Rate `json:"rate,omitempty"`
}

// Rate define rate object from TaxJar rate api endpoint
type Rate struct {
	Zip    string  `json:"zip,omitempty"`
	State  string  `json:"state,omitempty"`
	City   string  `json:"city,omitempty"`
	County string  `json:"county,omitempty"`
	Rate   float32 `json:"combined_rate,string,omitempty"`
}

// NewClient creates new application client to load and check vat rates.
func NewClient(db *leveldb.DB, maxRps int) *Client {
	return &Client{
		db:       db,
		limiter:  ratelimit.New(maxRps),
		context:  context.Background(),
		wg:       sync.WaitGroup{},
		messages: make(chan *Rate),
	}
}

// Client is a base object to to load and check vat rates.
type Client struct {
	db       *leveldb.DB
	limiter  ratelimit.Limiter
	context  context.Context
	wg       sync.WaitGroup
	messages chan *Rate
}

// RequestRate query rates from TaxJax by zip code.
func (c *Client) RequestRate(r *Record) {
	defer c.wg.Done()

	resp, err := resty.R().Get(r.Zip)
	if err != nil {
		zap.L().Error("TaxJar request failed", zap.Error(err), zap.String("zip", r.Zip))
		return
	}

	rateObj := &Response{}
	err = json.Unmarshal(resp.Body(), rateObj)
	if err != nil {
		zap.L().Error(
			"TaxJar response unmarshal failed",
			zap.Error(err),
			zap.String("zip", r.Zip),
			zap.String("response", resp.String()),
		)
		return
	}

	c.messages <- &rateObj.Rate
}

// ProcessRates check the rate changes against local embedded database and upload them in external micro service.
func (c *Client) ProcessRates() {
	for {
		select {
		case r, ok := <-c.messages:
			if !ok {
				zap.L().Error("Process rate channel closed")
				continue
			}

			data, err := c.db.Get([]byte(r.Zip), nil)
			if err != nil && err != leveldb.ErrNotFound {
				zap.L().Error("Can`t fetch cache data for zip", zap.Error(err), zap.String("zip", r.Zip))
				continue
			}

			rateBytes := float64ToByte(r.Rate)
			if bytes.Equal(data, rateBytes) {
				continue
			}
			err = c.db.Put([]byte(r.Zip), rateBytes, nil)
			if err != nil {
				zap.L().Error("Can`t update cache data for zip", zap.Error(err), zap.String("zip", r.Zip))
			}

			// Write to other micro service should be here.
			fmt.Printf("Value %v was read.\n", r)
		case <-c.context.Done():
			return
		}
	}
}

// Run load zip codes from Simplemaps csv file and request rate for each zip code from TaxJar.
func (c *Client) Run(file string) error {
	codes, err := readZipCodeFile(file)
	if err != nil {
		return nil
	}

	go c.ProcessRates()

	c.wg.Add(len(codes))
	for _, r := range codes {
		c.limiter.Take()
		go c.RequestRate(r)
	}

	c.wg.Wait()

	return nil
}

func float64ToByte(f float32) []byte {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], math.Float32bits(f))
	return buf[:]
}
