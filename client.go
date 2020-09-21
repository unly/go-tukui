package tukui

import "net/http"

const baseURL = "https://www.tukui.org/api.php"

// The Client is a simple http client to access the TukUI.org API.
// Addons can be accessed either by the RetailAddons or ClassicAddons field.
type Client struct {
	url           string
	httpClient    *http.Client
	RetailAddons  AddonClient
	ClassicAddons AddonClient
}

// NewClient creates a new Client struct and returns a pointer to it.
// Optionnaly you can pass a pointer to a http.Client to be used. Alternatively,
// http.DefaultClient is used.
func NewClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	c := Client{
		url:        baseURL,
		httpClient: client,
	}
	c.RetailAddons = newRetailClient(&c)
	c.ClassicAddons = newClassicClient(&c)

	return &c
}
