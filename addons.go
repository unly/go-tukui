package tukui

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Addon struct {
	Id            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	SmallDesc     *string `json:"small_desc,omitempty"`
	Author        *string `json:"author,omitempty"`
	Version       *string `json:"version,omitempty"`
	ScreenshotUrl *string `json:"screenshot_url,omitempty"`
	Category      *string `json:"category,omitempty"`
	Downloads     *string `json:"downloads,omitempty"`
	LastUpdate    *string `json:"lastupdate,omitempty"`
	Patch         *string `json:"patch,omitempty"`
	WebUrl        *string `json:"web_url,omitempty"`
	LastDownload  *string `json:"last_download,omitempty"`
	DonateUrl     *string `json:"donate_url,omitempty"`
}

type AddonClient struct {
	client  *Client
	classic bool
}

func (c *AddonClient) getQueryParameter(s string) string {
	if c.classic {
		return "classic-" + s
	}
	return s
}

func (c *AddonClient) GetAddon(id int) (Addon, *http.Response, error) {
	var addon Addon

	req, err := http.NewRequest(http.MethodGet, c.client.url, nil)
	if err != nil {
		return addon, nil, err
	}

	query := req.URL.Query()
	query.Add(c.getQueryParameter("addon"), strconv.Itoa(id))
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.httpClient.Do(req)
	if err != nil {
		return addon, resp, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&addon)

	return addon, resp, err
}

func (c *AddonClient) GetAddons() ([]Addon, *http.Response, error) {
	var addons []Addon

	req, err := http.NewRequest(http.MethodGet, c.client.url, nil)
	if err != nil {
		return addons, nil, err
	}

	query := req.URL.Query()
	query.Add(c.getQueryParameter("addons"), "all")
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.httpClient.Do(req)
	if err != nil {
		return addons, resp, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return addons, resp, err
	}

	err = json.Unmarshal(body, &addons)

	return addons, resp, err
}
