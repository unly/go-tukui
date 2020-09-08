[![Go Report Card](https://goreportcard.com/badge/github.com/unly/go-tukui)](https://goreportcard.com/report/github.com/unly/go-tukui)
[![License](https://img.shields.io/badge/license-MIT-green)](https://github.com/unly/go-tukui/blob/master/LICENSE)
[![CI Status](https://github.com/unly/go-tukui/workflows/CI/badge.svg)](https://github.com/unly/go-tukui/actions?query=workflow%3ACI)
[![Test Coverage](https://codecov.io/gh/unly/go-tukui/branch/master/graph/badge.svg)](https://codecov.io/gh/unly/go-tukui)


# Golang Client for [tukui.org](https://tukui.org)

Simple Golang client for the TukUI [API](https://www.tukui.org/api.php) to fetch the hosted addons.
Supports both classic and retail addons.

## Install

```
go get github.com/unly/go-tukui
```

## How to Use

Create a new client for the TukUI API.
Optionally, you can pass a pointer to a [http.Client](https://golang.org/pkg/net/http/#Client).
Otherwise the default client is used.
```
client := tukui.NewClient(nil)
```

Query for a specific addon using its ID, e.g. 3.
```
addon, resp, err := client.ClassicAddons.GetAddon(3)
```

Or query all available addons.
```
addons, resp, err := client.ClassicAddons.GetAddons()
```

## License

Licensed under the [MIT](https://github.com/unly/go-tukui/blob/master/LICENSE) license.
