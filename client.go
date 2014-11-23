package bgg

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) request(path string) (*http.Response, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("http://www.boardgamegeek.com/xmlapi2%s", path),
		nil,
	)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

func (c *Client) getXml(path string) (io.ReadCloser, error) {
	resp, err := c.request(path)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *Client) decode(path string, v interface{}) error {
	body, err := c.getXml(path)
	if err != nil {
		return err
	}
	defer body.Close()
	return xml.NewDecoder(body).Decode(v)
}
