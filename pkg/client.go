package taxjar

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"gopkg.in/resty.v1"
	"sync"
)

// TaxJarRateResponse define root object from TaxJar rate api endpoint
type TaxJarRateResponse struct {
	Rate TaxJarRate `json:"rate,omitempty"`
}

// TaxJarRate define rate object from TaxJar rate api endpoint
type TaxJarRate struct {
	Zip    string  `json:"zip,omitempty"`
	State  string  `json:"state,omitempty"`
	City   string  `json:"city,omitempty"`
	County string  `json:"county,omitempty"`
	Rate   float32 `json:"combined_rate,string,omitempty"`
}

func NewClient(maxRps int) *Client {
	return &Client{
		limiter:  ratelimit.New(maxRps),
		context:  context.Background(),
		wg:       sync.WaitGroup{},
		messages: make(chan *TaxJarRate),
	}
}

type Client struct {
	limiter  ratelimit.Limiter
	context  context.Context
	wg       sync.WaitGroup
	messages chan *TaxJarRate
}

func (c *Client) RequestRate(r *Record) {
	defer c.wg.Done()

	resp, err := resty.R().Get(r.Zip)
	if err != nil {
		zap.L().Error("TaxJar request failed", zap.Error(err), zap.String("zip", r.Zip))
		return
	}

	rateObj := &TaxJarRateResponse{}
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

func (c *Client) ProcessRates() {
	for {
		select {
		case rate, ok := <-c.messages:
			if ok {
				fmt.Printf("Value %v was read.\n", rate)
			} else {
				fmt.Println("Channel closed!")
			}
		case <-c.context.Done():
			return
		}
	}
}

func (c *Client) Run(file string) error {
	codes, err := readZipCodeFile(file)
	if err != nil {
		return nil
	}

	c.wg.Add(len(codes))
	for _, r := range codes {
		c.limiter.Take()
		go c.RequestRate(r)
	}

	go c.ProcessRates()
	c.wg.Wait()

	return nil
}
