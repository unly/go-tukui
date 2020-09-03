package tukui

import "net/http"

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
