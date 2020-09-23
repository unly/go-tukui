package tukui

import (
	"encoding/json"
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
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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

func TestRetail_GetTukUI(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
			"ui": {
				"tukui",
			},
		}))
		fmt.Fprint(w,
			`{
				"name": "Tukui",
				"author": "Tukz",
				"url": "https://www.tukui.org/downloads/tukui-18.28.zip",
				"version": "18.28",
				"changelog": "https://www.tukui.org/ui/tukui/changelog",
				"ticket": "https://git.tukui.org/Tukz/Tukui/issues",
				"git": "https://git.tukui.org/Tukz/Tukui",
				"id": -1,
				"patch": "8.3",
				"lastupdate": "2020-07-07",
				"web_url": "https://www.tukui.org/download.php?ui=tukui",
				"lastdownload": "2020-09-21 11:10:00",
				"donate_url": "http://www.tukui.org/support.php",
				"small_desc": "Minimalistic and lightweight world of warcraft user interface",
				"screenshot_url": "https://www.tukui.org/images/screenshots/t4.jpg",
				"downloads": 2147483000,
				"category": "Full UI Replacements"
			}`,
		)
	})

	tukui, _, err := client.RetailAddons.GetTukUI()
	if err != nil {
		t.Errorf("RetailAddons.GetTukUI() returned error: %v", err)
	}

	want := Addon{
		Id:            String("-1"),
		Author:        String("Tukz"),
		URL:           String("https://www.tukui.org/downloads/tukui-18.28.zip"),
		Version:       String("18.28"),
		Name:          String("Tukui"),
		Patch:         String("8.3"),
		LastUpdate:    String("2020-07-07"),
		WebUrl:        String("https://www.tukui.org/download.php?ui=tukui"),
		LastDownload:  String("2020-09-21 11:10:00"),
		DonateUrl:     String("http://www.tukui.org/support.php"),
		SmallDesc:     String("Minimalistic and lightweight world of warcraft user interface"),
		ScreenshotUrl: String("https://www.tukui.org/images/screenshots/t4.jpg"),
		Downloads:     String("2147483000"),
		Category:      String("Full UI Replacements"),
	}

	if !cmp.Equal(tukui, want) {
		t.Errorf("RetailAddons.GetTukUI() returned %+v, want %+v", tukui, want)
	}
}

func TestRetail_GetTukUI_Invalid(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
			"ui": {
				"tukui",
			},
		}))
		fmt.Fprint(w,
			`{
				"patch": 1.2
			}`,
		)
	})

	_, _, err := client.RetailAddons.GetTukUI()
	if err == nil {
		t.Errorf("RetailAddons.GetTukUI() returned no error")
	}
}

func TestRetail_GetTukUI_NoContent(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
			"ui": {
				"tukui",
			},
		}))
	})

	_, _, err := client.RetailAddons.GetTukUI()
	want := "empty response"
	if !cmp.Equal(err.Error(), want) {
		t.Errorf("RetailAddons.GetTukUI() returned %+v, want %+v", err, want)
	}
}

func TestRetail_GetElvUI(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
			"ui": {
				"elvui",
			},
		}))
		fmt.Fprint(w,
			`{
				"name": "ElvUI",
				"author": "Elv/Blazeflack",
				"url": "https://www.tukui.org/downloads/elvui-11.52.zip",
				"version": "11.52",
				"changelog": "https://www.tukui.org/ui/elvui/changelog",
				"ticket": "https://git.tukui.org/elvui/elvui/issues",
				"git": "https://git.tukui.org/elvui/elvui",
				"id": -2,
				"patch": "8.3",
				"lastupdate": "2020-09-04",
				"web_url": "https://www.tukui.org/download.php?ui=elvui",
				"lastdownload": "2020-09-21 13:34:28",
				"donate_url": "http://www.tukui.org/support.php",
				"small_desc": "A user interface designed around user-friendliness with extra features that are not included in the standard ui",
				"screenshot_url": "https://www.tukui.org/images/screenshots/DarkTheme_ThickBorders_DPS.jpg",
				"downloads": 2147483000,
				"category": "Full UI Replacements"
			}`,
		)
	})

	elvui, _, err := client.RetailAddons.GetElvUI()
	if err != nil {
		t.Errorf("RetailAddons.GetElvUI() returned error: %v", err)
	}

	want := Addon{
		Id:            String("-2"),
		Author:        String("Elv/Blazeflack"),
		URL:           String("https://www.tukui.org/downloads/elvui-11.52.zip"),
		Version:       String("11.52"),
		Name:          String("ElvUI"),
		Patch:         String("8.3"),
		LastUpdate:    String("2020-09-04"),
		WebUrl:        String("https://www.tukui.org/download.php?ui=elvui"),
		LastDownload:  String("2020-09-21 13:34:28"),
		DonateUrl:     String("http://www.tukui.org/support.php"),
		SmallDesc:     String("A user interface designed around user-friendliness with extra features that are not included in the standard ui"),
		ScreenshotUrl: String("https://www.tukui.org/images/screenshots/DarkTheme_ThickBorders_DPS.jpg"),
		Downloads:     String("2147483000"),
		Category:      String("Full UI Replacements"),
	}

	if !cmp.Equal(elvui, want) {
		t.Errorf("RetailAddons.GetElvUI() returned %+v, want %+v", elvui, want)
	}
}

func TestRetail_GetElvUI_Invalid(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
			"ui": {
				"elvui",
			},
		}))
		fmt.Fprint(w,
			`{
				"patch": 1.2
			}`,
		)
	})

	_, _, err := client.RetailAddons.GetElvUI()
	if err == nil {
		t.Errorf("RetailAddons.GetElvUI() returned no error")
	}
}

func TestRetail_GetElvUI_NoContent(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
			"ui": {
				"elvui",
			},
		}))
	})

	_, _, err := client.RetailAddons.GetElvUI()
	want := "empty response"
	if !cmp.Equal(err.Error(), want) {
		t.Errorf("RetailAddons.GetElvUI() returned %+v, want %+v", err, want)
	}
}

func TestClassic_GetAddons(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
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

func TestClassic_GetTukUI(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
			"classic-addon": {
				"1",
			},
		}))
		fmt.Fprint(w,
			`{
				"id": "1",
				"name": "Tukui",
				"small_desc": "A clean, lightweight, minimalist and popular user interface among the warcraft community since 2007.",
				"author": "Tukz",
				"version": "1.38",
				"screenshot_url": "https://www.tukui.org/1",
				"url": "https://www.tukui.org/classic-addons.php?download=1",
				"category": "Interfaces",
				"downloads": "260226",
				"lastupdate": "2020-07-11 23:10:36",
				"patch": "1.13.4",
				"last_download": "2020-09-21 13:42:09",
				"web_url": "https://www.tukui.org/classic-addons.php?id=1",
				"changelog": "https://www.tukui.org/classic-addons.php?id=1&changelog",
				"donate_url": "https://www.tukui.org/support.php"
			}`,
		)
	})

	tukui, _, err := client.ClassicAddons.GetTukUI()
	if err != nil {
		t.Errorf("ClassicAddons.GetTukUI() returned error: %v", err)
	}

	want := Addon{
		Id:            String("1"),
		Author:        String("Tukz"),
		URL:           String("https://www.tukui.org/classic-addons.php?download=1"),
		Version:       String("1.38"),
		Name:          String("Tukui"),
		Patch:         String("1.13.4"),
		LastUpdate:    String("2020-07-11 23:10:36"),
		WebUrl:        String("https://www.tukui.org/classic-addons.php?id=1"),
		LastDownload:  String("2020-09-21 13:42:09"),
		DonateUrl:     String("https://www.tukui.org/support.php"),
		SmallDesc:     String("A clean, lightweight, minimalist and popular user interface among the warcraft community since 2007."),
		ScreenshotUrl: String("https://www.tukui.org/1"),
		Downloads:     String("260226"),
		Category:      String("Interfaces"),
	}

	if !cmp.Equal(tukui, want) {
		t.Errorf("ClassicAddons.GetTukUI() returned %+v, want %+v", tukui, want)
	}
}

func TestClassic_GetTukUI_Invalid(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
			"classic-addon": {
				"1",
			},
		}))
		fmt.Fprint(w,
			`{
				"id": 1,
				"name": "Tukui",
				"small_desc": "A clean, lightweight, minimalist and popular user interface among the warcraft community since 2007.",
				"author": "Tukz",
				"version": "1.38",
				"screenshot_url": "https://www.tukui.org/1",
				"url": "https://www.tukui.org/classic-addons.php?download=1",
				"category": "Interfaces",
				"downloads": "260226",
				"lastupdate": "2020-07-11 23:10:36",
				"patch": "1.13.4",
				"last_download": "2020-09-21 13:42:09",
				"web_url": "https://www.tukui.org/classic-addons.php?id=1",
				"changelog": "https://www.tukui.org/classic-addons.php?id=1&changelog",
				"donate_url": "https://www.tukui.org/support.php"
			}`,
		)
	})

	_, _, err := client.ClassicAddons.GetTukUI()
	if err == nil {
		t.Errorf("ClassicAddons.GetTukUI() returned no error")
	}
}

func TestClassic_GetTukUI_NoContent(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
			"classic-addon": {
				"1",
			},
		}))
	})

	_, _, err := client.ClassicAddons.GetTukUI()
	want := "empty response"
	if !cmp.Equal(err.Error(), want) {
		t.Errorf("ClassicAddons.GetTukUI() returned %+v, want %+v", err, want)
	}
}

func TestClassic_GetElvUI(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
			"classic-addon": {
				"2",
			},
		}))
		fmt.Fprint(w,
			`{
				"id": "2",
				"name": "ElvUI",
				"small_desc": "A USER INTERFACE DESIGNED AROUND USER-FRIENDLINESS WITH EXTRA FEATURES THAT ARE NOT INCLUDED IN THE STANDARD UI.\r\n",
				"author": "Elv",
				"version": "1.31",
				"screenshot_url": "https://www.tukui.org/2",
				"url": "https://www.tukui.org/classic-addons.php?download=2",
				"category": "Interfaces",
				"downloads": "1721530",
				"lastupdate": "2020-09-07 19:42:13",
				"patch": "1.13.5",
				"last_download": "2020-09-21 13:58:12",
				"web_url": "https://www.tukui.org/classic-addons.php?id=2",
				"changelog": "https://www.tukui.org/classic-addons.php?id=2&changelog"
			}`,
		)
	})

	elvui, _, err := client.ClassicAddons.GetElvUI()
	if err != nil {
		t.Errorf("ClassicAddons.GetElvUI() returned error: %v", err)
	}

	want := Addon{
		Id:            String("2"),
		Author:        String("Elv"),
		URL:           String("https://www.tukui.org/classic-addons.php?download=2"),
		Version:       String("1.31"),
		Name:          String("ElvUI"),
		Patch:         String("1.13.5"),
		LastUpdate:    String("2020-09-07 19:42:13"),
		WebUrl:        String("https://www.tukui.org/classic-addons.php?id=2"),
		LastDownload:  String("2020-09-21 13:58:12"),
		SmallDesc:     String("A USER INTERFACE DESIGNED AROUND USER-FRIENDLINESS WITH EXTRA FEATURES THAT ARE NOT INCLUDED IN THE STANDARD UI.\r\n"),
		ScreenshotUrl: String("https://www.tukui.org/2"),
		Downloads:     String("1721530"),
		Category:      String("Interfaces"),
	}

	if !cmp.Equal(elvui, want) {
		t.Errorf("ClassicAddons.GetElvUI() returned %+v, want %+v", elvui, want)
	}
}

func TestClassic_GetElvUI_Invalid(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
			"classic-addon": {
				"2",
			},
		}))
		fmt.Fprint(w,
			`{
				"id": 2,
				"name": "ElvUI",
				"small_desc": "A USER INTERFACE DESIGNED AROUND USER-FRIENDLINESS WITH EXTRA FEATURES THAT ARE NOT INCLUDED IN THE STANDARD UI.\r\n",
				"author": "Elv",
				"version": "1.31",
				"screenshot_url": "https://www.tukui.org/2",
				"url": "https://www.tukui.org/classic-addons.php?download=2",
				"category": "Interfaces",
				"downloads": "1721530",
				"lastupdate": "2020-09-07 19:42:13",
				"patch": "1.13.5",
				"last_download": "2020-09-21 13:58:12",
				"web_url": "https://www.tukui.org/classic-addons.php?id=2",
				"changelog": "https://www.tukui.org/classic-addons.php?id=2&changelog"
			}`,
		)
	})

	_, _, err := client.ClassicAddons.GetElvUI()
	if err == nil {
		t.Errorf("ClassicAddons.GetElvUI() returned no error")
	}
}

func TestClassic_GetElvUI_NoContent(t *testing.T) {
	client, mux, teardown := setupTestEnv()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		testHTTPQuery(t, r, url.Values(map[string][]string{
			"classic-addon": {
				"2",
			},
		}))
	})

	_, _, err := client.ClassicAddons.GetElvUI()
	want := "empty response"
	if !cmp.Equal(err.Error(), want) {
		t.Errorf("ClassicAddons.GetElvUI() returned %+v, want %+v", err, want)
	}
}

func TestConvertAddon_NoID_NoDownloads_NoLastDownload(t *testing.T) {
	ui := uiAddon{
		Addon: Addon{
			Author:        String("Elv"),
			URL:           String("https://www.tukui.org/classic-addons.php?download=2"),
			Version:       String("1.31"),
			Name:          String("ElvUI"),
			Patch:         String("1.13.5"),
			LastUpdate:    String("2020-09-07 19:42:13"),
			WebUrl:        String("https://www.tukui.org/classic-addons.php?id=2"),
			SmallDesc:     String("A USER INTERFACE DESIGNED AROUND USER-FRIENDLINESS WITH EXTRA FEATURES THAT ARE NOT INCLUDED IN THE STANDARD UI.\r\n"),
			ScreenshotUrl: String("https://www.tukui.org/2"),
			Category:      String("Interfaces"),
		},
		Id:           nil,
		Downloads:    nil,
		LastDownload: nil,
	}

	got := convertAddon(ui)

	want := Addon{
		Author:        String("Elv"),
		URL:           String("https://www.tukui.org/classic-addons.php?download=2"),
		Version:       String("1.31"),
		Name:          String("ElvUI"),
		Patch:         String("1.13.5"),
		LastUpdate:    String("2020-09-07 19:42:13"),
		WebUrl:        String("https://www.tukui.org/classic-addons.php?id=2"),
		SmallDesc:     String("A USER INTERFACE DESIGNED AROUND USER-FRIENDLINESS WITH EXTRA FEATURES THAT ARE NOT INCLUDED IN THE STANDARD UI.\r\n"),
		ScreenshotUrl: String("https://www.tukui.org/2"),
		Category:      String("Interfaces"),
	}

	if !cmp.Equal(got, want) {
		t.Errorf("convertAddon() returned %+v, want %+v", got, want)
	}
}

func TestConvertAddon_EmptyID(t *testing.T) {
	ui := uiAddon{
		Addon: Addon{
			Author:        String("Elv"),
			URL:           String("https://www.tukui.org/classic-addons.php?download=2"),
			Version:       String("1.31"),
			Name:          String("ElvUI"),
			Patch:         String("1.13.5"),
			LastUpdate:    String("2020-09-07 19:42:13"),
			WebUrl:        String("https://www.tukui.org/classic-addons.php?id=2"),
			SmallDesc:     String("A USER INTERFACE DESIGNED AROUND USER-FRIENDLINESS WITH EXTRA FEATURES THAT ARE NOT INCLUDED IN THE STANDARD UI.\r\n"),
			ScreenshotUrl: String("https://www.tukui.org/2"),
			Category:      String("Interfaces"),
		},
		Id:           (*json.Number)(String("")),
		Downloads:    (*json.Number)(String("123")),
		LastDownload: String("2020-09-21 13:58:12"),
	}

	got := convertAddon(ui)

	want := Addon{
		Author:        String("Elv"),
		URL:           String("https://www.tukui.org/classic-addons.php?download=2"),
		Version:       String("1.31"),
		Name:          String("ElvUI"),
		Patch:         String("1.13.5"),
		LastUpdate:    String("2020-09-07 19:42:13"),
		WebUrl:        String("https://www.tukui.org/classic-addons.php?id=2"),
		SmallDesc:     String("A USER INTERFACE DESIGNED AROUND USER-FRIENDLINESS WITH EXTRA FEATURES THAT ARE NOT INCLUDED IN THE STANDARD UI.\r\n"),
		ScreenshotUrl: String("https://www.tukui.org/2"),
		Category:      String("Interfaces"),
		Id:            String(""),
		Downloads:     String("123"),
		LastDownload:  String("2020-09-21 13:58:12"),
	}

	if !cmp.Equal(got, want) {
		t.Errorf("convertAddon() returned %+v, want %+v", got, want)
	}
}

func TestConvertAddon_EmptyAddon(t *testing.T) {
	ui := uiAddon{}

	got := convertAddon(ui)

	want := Addon{}

	if !cmp.Equal(got, want) {
		t.Errorf("convertAddon() returned %+v, want %+v", got, want)
	}
}
