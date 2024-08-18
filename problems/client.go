package problems

import (
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	Base string
}

func NewClient(base string) *Client {
	return &Client{
		Base: base,
	}
}

func NewKenkooooClient() *Client {
	return NewClient("https://kenkoooo.com/atcoder/resources/")
}

func (c *Client) FullURL(req string) string {
	return c.Base + req + ".json"
}

func (c *Client) Fetch(name string) ([]byte, error) {
	resp, err := http.Get(c.FullURL(name))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (c *Client) FetchAndUnmarshal(name string, v interface{}) error {
	bytes, err := c.Fetch(name)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, v)
}

