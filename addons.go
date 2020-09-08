package tukui

import (
	"net/http"
	"strconv"
)

// An Addon is the basic return of the HTTP call. It contains the fields specified by
// the API. Pointers to strings can be nil.
type Addon struct {
	// ID number of addon
	Id *string `json:"id,omitempty"`
	// the title of addon
	Name *string `json:"name,omitempty"`
	// a condensed description of addon for small area
	SmallDesc *string `json:"small_desc,omitempty"`
	// the author username
	Author *string `json:"author,omitempty"`
	// latest version of the addon uploaded on our network
	Version *string `json:"version,omitempty"`
	// a screenshot url of the addon
	ScreenshotUrl *string `json:"screenshot_url,omitempty"`
	// URL to download the .zip file
	URL *string `json:"url,omitempty"`
	// the main category where the addon is located
	Category *string `json:"category,omitempty"`
	// the total number of downloads of this addon
	Downloads *string `json:"downloads,omitempty"`
	// the last time the addon was updated
	LastUpdate *string `json:"lastupdate,omitempty"`
	// which World of Warcraft patch this addon is compatible with
	Patch *string `json:"patch,omitempty"`
	// url of the addon if an user want to visit his official download web page on our website
	WebUrl *string `json:"web_url,omitempty"`
	// when the addon was downloaded for the last time
	LastDownload *string `json:"last_download,omitempty"`
	// a donate url if the addon author accept donations
	DonateUrl *string `json:"donate_url,omitempty"`
}

type addonClient struct {
	client  *Client
	classic bool
}

func (c *addonClient) getQueryParameter(s string) string {
	if c.classic {
		return "classic-" + s
	}
	return s
}

// GetAddon returns the Addon for the given ID. The ID is a positive number.
// For non existing IDs the function will return an error.
func (c *addonClient) GetAddon(id int) (Addon, *http.Response, error) {
	var addon Addon

	req, err := http.NewRequest(http.MethodGet, c.client.url, nil)
	if err != nil {
		return addon, nil, err
	}

	query := req.URL.Query()
	query.Add(c.getQueryParameter("addon"), strconv.Itoa(id))
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.request(req, &addon)

	return addon, resp, err
}

// GetAddons returns a slice of all Addons available.
func (c *addonClient) GetAddons() ([]Addon, *http.Response, error) {
	var addons []Addon

	req, err := http.NewRequest(http.MethodGet, c.client.url, nil)
	if err != nil {
		return addons, nil, err
	}

	query := req.URL.Query()
	query.Add(c.getQueryParameter("addons"), "all")
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.request(req, &addons)

	return addons, resp, err
}
