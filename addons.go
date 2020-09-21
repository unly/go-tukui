package tukui

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

type uiAddon struct {
	Addon

	// id is an integer not a string for UIs
	Id *json.Number `json:"id,omitempty"`
	// downloads is an integer not a string for UIs
	Downloads *json.Number `json:"downloads,omitempty"`
	// last download has a different name for UIs
	LastDownload *string `json:"lastdownload,omitempty"`
}

// AddonClient is a set of functions that can be queried from the TukUI.org API
type AddonClient interface {
	// GetAddon returns the Addon for the given ID. The ID is a positive number.
	// For non existing IDs the function will return an error.
	GetAddon(id int) (Addon, *http.Response, error)
	// GetAddons returns a slice of all Addons available.
	GetAddons() ([]Addon, *http.Response, error)
	// GetTukUI returns the Addon for the main TukUI
	GetTukUI() (Addon, *http.Response, error)
	// GetElvUI returns the Addon for the main ElvUI
	GetElvUI() (Addon, *http.Response, error)
}

type retailClient struct {
	apiClient
}

type classicClient struct {
	apiClient
}

type apiClient struct {
	client *Client
}

func newRetailClient(client *Client) *retailClient {
	return &retailClient{
		apiClient: apiClient{
			client: client,
		},
	}
}

func newClassicClient(client *Client) *classicClient {
	return &classicClient{
		apiClient: apiClient{
			client: client,
		},
	}
}

func (r *retailClient) GetTukUI() (Addon, *http.Response, error) {
	var tukui uiAddon

	resp, err := r.queryAPI("ui", "tukui", &tukui)

	return convertAddon(tukui), resp, err
}

func (r *retailClient) GetElvUI() (Addon, *http.Response, error) {
	var elvui uiAddon

	resp, err := r.queryAPI("ui", "elvui", &elvui)

	return convertAddon(elvui), resp, err
}

func (r *retailClient) GetAddon(id int) (Addon, *http.Response, error) {
	var addon Addon

	resp, err := r.queryAPI("addon", strconv.Itoa(id), &addon)

	return addon, resp, err
}

func (r *retailClient) GetAddons() ([]Addon, *http.Response, error) {
	var addons []Addon

	resp, err := r.queryAPI("addons", "all", &addons)

	return addons, resp, err
}

func (c *classicClient) GetTukUI() (Addon, *http.Response, error) {
	return c.GetAddon(1)
}

func (c *classicClient) GetElvUI() (Addon, *http.Response, error) {
	return c.GetAddon(2)
}

func (c *classicClient) GetAddon(id int) (Addon, *http.Response, error) {
	var addon Addon

	resp, err := c.queryAPI("classic-addon", strconv.Itoa(id), &addon)

	return addon, resp, err
}

func (c *classicClient) GetAddons() ([]Addon, *http.Response, error) {
	var addons []Addon

	resp, err := c.queryAPI("classic-addons", "all", &addons)

	return addons, resp, err
}

func (a *apiClient) queryAPI(key, value string, data interface{}) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, a.client.url, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add(key, value)
	req.URL.RawQuery = query.Encode()

	resp, err := a.client.httpClient.Do(req)
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

func convertAddon(ui uiAddon) Addon {
	addon := Addon{
		Name:          ui.Name,
		SmallDesc:     ui.SmallDesc,
		Author:        ui.Author,
		Version:       ui.Version,
		ScreenshotUrl: ui.ScreenshotUrl,
		URL:           ui.URL,
		Category:      ui.Category,
		LastUpdate:    ui.LastUpdate,
		Patch:         ui.Patch,
		WebUrl:        ui.WebUrl,
		LastDownload:  ui.LastDownload,
		DonateUrl:     ui.DonateUrl,
	}

	if ui.Id != nil {
		id := ui.Id.String()
		addon.Id = &id
	}

	if ui.Downloads != nil {
		downloads := ui.Downloads.String()
		addon.Downloads = &downloads
	}

	return addon
}
