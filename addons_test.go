package tukui

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRetail_GetAddon(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"addon": {
				"3",
			},
		}))
		fmt.Fprint(w,
			`{
				"id": "3",
				"name": "AddOnSkins",
				"small_desc": "Skins for AddOns",
				"author": "Azilroka",
				"version": "3.53",
				"screenshot_url": "https://www.tukui.org/3",
				"url": "https://www.tukui.org/addons.php?download=3",
				"category": "Skins",
				"downloads": "46156",
				"lastupdate": "2017-09-09 07:09:10",
				"patch": "7.2.5",
				"last_download": "2017-09-14 15:20:22",
				"web_url": "https://www.tukui.org/addons.php?id=3",
				"donate_url": "https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=TENQSTDX5SEWE"
			}`,
		)
	})

	addon, _, err := client.RetailAddons.GetAddon(3)
	if err != nil {
		t.Errorf("RetailAddons.GetAddon() returned error: %v", err)
	}

	want := Addon{
		Id:            String("3"),
		Name:          String("AddOnSkins"),
		SmallDesc:     String("Skins for AddOns"),
		Author:        String("Azilroka"),
		Version:       String("3.53"),
		ScreenshotUrl: String("https://www.tukui.org/3"),
		URL:           String("https://www.tukui.org/addons.php?download=3"),
		Category:      String("Skins"),
		Downloads:     String("46156"),
		LastUpdate:    String("2017-09-09 07:09:10"),
		Patch:         String("7.2.5"),
		LastDownload:  String("2017-09-14 15:20:22"),
		WebUrl:        String("https://www.tukui.org/addons.php?id=3"),
		DonateUrl:     String("https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=TENQSTDX5SEWE"),
	}

	if !cmp.Equal(addon, want) {
		t.Errorf("RetailAddons.GetAddon() returned %+v, want %+v", addon, want)
	}
}

func TestRetail_GetAddon_NegativeNumber(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"addon": {
				"-3",
			},
		}))
	})

	_, _, err := client.RetailAddons.GetAddon(-3)
	want := "empty response"
	if !cmp.Equal(err.Error(), want) {
		t.Errorf("RetailAddons.GetAddon() returned %+v, want %+v", err, want)
	}
}

func TestRetail_GetAddon_InvalidResponse(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"addon": {
				"42",
			},
		}))
		fmt.Fprint(w,
			`[
				{
					"Id":42
				}
			]`,
		)
	})

	_, _, err := client.RetailAddons.GetAddon(42)
	if err == nil {
		t.Errorf("RetailAddons.GetAddons() returned no error")
	}
}

func TestClassic_GetAddon(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"classic-addon": {
				"3",
			},
		}))
		fmt.Fprint(w,
			`{
				"id": "3",
				"name": "AddOnSkins",
				"small_desc": "Skins for AddOns",
				"author": "Azilroka",
				"version": "3.53",
				"screenshot_url": "https://www.tukui.org/3",
				"url": "https://www.tukui.org/classic-addons.php?download=3",
				"category": "Skins",
				"downloads": "46156",
				"lastupdate": "2017-09-09 07:09:10",
				"patch": "7.2.5",
				"last_download": "2017-09-14 15:20:22",
				"web_url": "https://www.tukui.org/classic-addons.php?id=3",
				"donate_url": "https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=TENQSTDX5SEWE"
			}`,
		)
	})

	addon, _, err := client.ClassicAddons.GetAddon(3)
	if err != nil {
		t.Errorf("ClassicAddons.GetAddon() returned error: %v", err)
	}

	want := Addon{
		Id:            String("3"),
		Name:          String("AddOnSkins"),
		SmallDesc:     String("Skins for AddOns"),
		Author:        String("Azilroka"),
		Version:       String("3.53"),
		ScreenshotUrl: String("https://www.tukui.org/3"),
		URL:           String("https://www.tukui.org/classic-addons.php?download=3"),
		Category:      String("Skins"),
		Downloads:     String("46156"),
		LastUpdate:    String("2017-09-09 07:09:10"),
		Patch:         String("7.2.5"),
		LastDownload:  String("2017-09-14 15:20:22"),
		WebUrl:        String("https://www.tukui.org/classic-addons.php?id=3"),
		DonateUrl:     String("https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=TENQSTDX5SEWE"),
	}

	if !cmp.Equal(addon, want) {
		t.Errorf("ClassicAddons.GetAddon() returned %+v, want %+v", addon, want)
	}
}

func TestClassic_GetAddon_NegativeNumber(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"classic-addon": {
				"-3",
			},
		}))
	})

	_, _, err := client.ClassicAddons.GetAddon(-3)
	want := "empty response"
	if !cmp.Equal(err.Error(), want) {
		t.Errorf("ClassicAddons.GetAddon() returned %+v, want %+v", err, want)
	}
}

func TestClassic_GetAddon_InvalidResponse(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"classic-addon": {
				"42",
			},
		}))
		fmt.Fprint(w,
			`[
				{
					"Id":42
				}
			]`,
		)
	})

	_, _, err := client.ClassicAddons.GetAddon(42)
	if err == nil {
		t.Errorf("ClassicAddons.GetAddons() returned no error")
	}
}

func TestRetail_GetAddons(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"addons": {
				"all",
			},
		}))
		fmt.Fprint(w,
			`[
				{
					"id": "3",
					"name": "AddOnSkins",
					"small_desc": "Skins for AddOns",
					"author": "Azilroka",
					"version": "3.53",
					"screenshot_url": "https://www.tukui.org/3",
					"url": "https://www.tukui.org/addons.php?download=3",
					"category": "Skins",
					"downloads": "46156",
					"lastupdate": "2017-09-09 07:09:10",
					"patch": "7.2.5",
					"last_download": "2017-09-14 15:20:22",
					"web_url": "https://www.tukui.org/addons.php?id=3",
					"donate_url": "https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=TENQSTDX5SEWE"
				},
				{
					"id": "6", 
					"name": "LocationPlus for ElvUI",
					"small_desc": "Adds player location, coords + 2 Datatexts and a tooltip with info based on player location/level.", 
					"author": "Benik",
					"version": "2.48",
					"screenshot_url": "https://www.tukui.org/6",
					"url": "https://www.tukui.org/addons.php?download=6",
					"category": "Plugins: ElvUI",
					"downloads": "360540",
					"lastupdate": "2020-07-08 21:42:32",
					"patch": "8.3",
					"last_download": "2020-09-04 09:32:17",
					"web_url": "https://www.tukui.org/addons.php?id=6", 
					"changelog": "https://www.tukui.org/addons.php?id=6&changelog"				}
			]`,
		)
	})

	addons, _, err := client.RetailAddons.GetAddons()
	if err != nil {
		t.Errorf("RetailAddons.GetAddons() returned error: %v", err)
	}

	want := []Addon{
		{
			Id:            String("3"),
			Name:          String("AddOnSkins"),
			SmallDesc:     String("Skins for AddOns"),
			Author:        String("Azilroka"),
			Version:       String("3.53"),
			ScreenshotUrl: String("https://www.tukui.org/3"),
			URL:           String("https://www.tukui.org/addons.php?download=3"),
			Category:      String("Skins"),
			Downloads:     String("46156"),
			LastUpdate:    String("2017-09-09 07:09:10"),
			Patch:         String("7.2.5"),
			LastDownload:  String("2017-09-14 15:20:22"),
			WebUrl:        String("https://www.tukui.org/addons.php?id=3"),
			DonateUrl:     String("https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=TENQSTDX5SEWE"),
		},
		{
			Id:            String("6"),
			Name:          String("LocationPlus for ElvUI"),
			SmallDesc:     String("Adds player location, coords + 2 Datatexts and a tooltip with info based on player location/level."),
			Author:        String("Benik"),
			Version:       String("2.48"),
			ScreenshotUrl: String("https://www.tukui.org/6"),
			URL:           String("https://www.tukui.org/addons.php?download=6"),
			Category:      String("Plugins: ElvUI"),
			Downloads:     String("360540"),
			LastUpdate:    String("2020-07-08 21:42:32"),
			Patch:         String("8.3"),
			LastDownload:  String("2020-09-04 09:32:17"),
			WebUrl:        String("https://www.tukui.org/addons.php?id=6"),
		},
	}

	if !cmp.Equal(addons, want) {
		t.Errorf("RetailAddons.GetAddons() returned %+v, want %+v", addons, want)
	}
}

func TestRetail_GetAddons_Invalid(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"addons": {
				"all",
			},
		}))
		fmt.Fprint(w,
			`[
				{
					"Id":0
				}
			]`,
		)
	})

	_, _, err := client.RetailAddons.GetAddons()
	if err == nil {
		t.Errorf("RetailAddons.GetAddons() returned no error")
	}
}

func TestRetail_GetAddons_Empty(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"addons": {
				"all",
			},
		}))
		fmt.Fprint(w,
			`[
			]`,
		)
	})

	addons, _, err := client.RetailAddons.GetAddons()
	if err != nil {
		t.Errorf("RetailAddons.GetAddons() returned error: %v", err)
	}

	want := []Addon{}

	if !cmp.Equal(addons, want) {
		t.Errorf("RetailAddons.GetAddons() returned %+v, want %+v", addons, want)
	}
}

func TestRetail_GetAddons_NoContent(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"addons": {
				"all",
			},
		}))
	})

	_, _, err := client.RetailAddons.GetAddons()
	want := "empty response"
	if !cmp.Equal(err.Error(), want) {
		t.Errorf("RetailAddons.GetAddons() returned %+v, want %+v", err, want)
	}
}

func TestClassic_GetAddons(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"classic-addons": {
				"all",
			},
		}))
		fmt.Fprint(w,
			`[
				{
					"id": "3",
					"name": "AddOnSkins",
					"small_desc": "Skins for AddOns",
					"author": "Azilroka",
					"version": "3.53",
					"screenshot_url": "https://www.tukui.org/3",
					"url": "https://www.tukui.org/classic-addons.php?download=3",
					"category": "Skins",
					"downloads": "46156",
					"lastupdate": "2017-09-09 07:09:10",
					"patch": "7.2.5",
					"last_download": "2017-09-14 15:20:22",
					"web_url": "https://www.tukui.org/classic-addons.php?id=3",
					"donate_url": "https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=TENQSTDX5SEWE"
				},
				{
					"id": "6", 
					"name": "LocationPlus for ElvUI",
					"small_desc": "Adds player location, coords + 2 Datatexts and a tooltip with info based on player location/level.", 
					"author": "Benik",
					"version": "2.48",
					"screenshot_url": "https://www.tukui.org/6",
					"url": "https://www.tukui.org/classic-addons.php?download=6",
					"category": "Plugins: ElvUI",
					"downloads": "360540",
					"lastupdate": "2020-07-08 21:42:32",
					"patch": "8.3",
					"last_download": "2020-09-04 09:32:17",
					"web_url": "https://www.tukui.org/classic-addons.php?id=6", 
					"changelog": "https://www.tukui.org/classic-addons.php?id=6&changelog"				}
			]`,
		)
	})

	addons, _, err := client.ClassicAddons.GetAddons()
	if err != nil {
		t.Errorf("ClassicAddons.GetAddons() returned error: %v", err)
	}

	want := []Addon{
		{
			Id:            String("3"),
			Name:          String("AddOnSkins"),
			SmallDesc:     String("Skins for AddOns"),
			Author:        String("Azilroka"),
			Version:       String("3.53"),
			ScreenshotUrl: String("https://www.tukui.org/3"),
			URL:           String("https://www.tukui.org/classic-addons.php?download=3"),
			Category:      String("Skins"),
			Downloads:     String("46156"),
			LastUpdate:    String("2017-09-09 07:09:10"),
			Patch:         String("7.2.5"),
			LastDownload:  String("2017-09-14 15:20:22"),
			WebUrl:        String("https://www.tukui.org/classic-addons.php?id=3"),
			DonateUrl:     String("https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=TENQSTDX5SEWE"),
		},
		{
			Id:            String("6"),
			Name:          String("LocationPlus for ElvUI"),
			SmallDesc:     String("Adds player location, coords + 2 Datatexts and a tooltip with info based on player location/level."),
			Author:        String("Benik"),
			Version:       String("2.48"),
			ScreenshotUrl: String("https://www.tukui.org/6"),
			URL:           String("https://www.tukui.org/classic-addons.php?download=6"),
			Category:      String("Plugins: ElvUI"),
			Downloads:     String("360540"),
			LastUpdate:    String("2020-07-08 21:42:32"),
			Patch:         String("8.3"),
			LastDownload:  String("2020-09-04 09:32:17"),
			WebUrl:        String("https://www.tukui.org/classic-addons.php?id=6"),
		},
	}

	if !cmp.Equal(addons, want) {
		t.Errorf("ClassicAddons.GetAddons() returned %+v, want %+v", addons, want)
	}
}

func TestClassic_GetAddons_Invalid(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"classic-addons": {
				"all",
			},
		}))
		fmt.Fprint(w,
			`[
				{
					"Id":0
				}
			]`,
		)
	})

	_, _, err := client.ClassicAddons.GetAddons()
	if err == nil {
		t.Errorf("ClassicAddons.GetAddons() returned no error")
	}
}

func TestClassic_GetAddons_Empty(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"classic-addons": {
				"all",
			},
		}))
		fmt.Fprint(w,
			`[
			]`,
		)
	})

	addons, _, err := client.ClassicAddons.GetAddons()
	if err != nil {
		t.Errorf("ClassicAddons.GetAddons() returned error: %v", err)
	}

	want := []Addon{}

	if !cmp.Equal(addons, want) {
		t.Errorf("ClassicAddons.GetAddons() returned %+v, want %+v", addons, want)
	}
}

func TestClassic_GetAddons_NoContent(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, http.MethodGet)
		testHttpQuery(t, r, url.Values(map[string][]string{
			"classic-addons": {
				"all",
			},
		}))
	})

	_, _, err := client.ClassicAddons.GetAddons()
	want := "empty response"
	if !cmp.Equal(err.Error(), want) {
		t.Errorf("ClassicAddons.GetAddons() returned %+v, want %+v", err, want)
	}
}
