package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ghana7989/toll-calculator/types"
)

type Client struct {
	EndPoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		EndPoint: endpoint,
	}
}

func (c *Client) AggregateInvoice(d types.Distance) error {
	httpc := http.Client{}
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.EndPoint, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	res, err := httpc.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("the service responded with status code: %d", res.StatusCode)
	}
	return nil
}
