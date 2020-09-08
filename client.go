package tukui

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const baseUrl = "https://www.tukui.org/api.php"

// The Client is a simple http client to access the TukUI.org API.
// Addons can be accessed either by the RetailAddons or ClassicAddons field.
type Client struct {
	url           string
	httpClient    *http.Client
	RetailAddons  *addonClient
	ClassicAddons *addonClient
}

// NewClient creates a new Client struct and returns a pointer to it.
// Optionnaly you can pass a pointer to a http.Client to be used. Alternatively,
// http.DefaultClient is used.
func NewClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	c := Client{
		url:        baseUrl,
		httpClient: client,
	}
	c.RetailAddons = &addonClient{
		client:  &c,
		classic: false,
	}
	c.ClassicAddons = &addonClient{
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
