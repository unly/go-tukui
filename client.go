package tukui

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const baseUrl = "https://www.tukui.org/api.php"

type Client struct {
	url           string
	httpClient    *http.Client
	RetailAddons  *AddonClient
	ClassicAddons *AddonClient
}

func NewClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	c := Client{
		url:        baseUrl,
		httpClient: client,
	}
	c.RetailAddons = &AddonClient{
		client:  &c,
		classic: false,
	}
	c.ClassicAddons = &AddonClient{
		client:  &c,
		classic: true,
	}

	return &c
}

func (c *Client) request(req *http.Request, data interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	if resp.ContentLength == 0 {
		return resp, errors.New("empty response")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}

	return resp, json.Unmarshal(body, data)
}
