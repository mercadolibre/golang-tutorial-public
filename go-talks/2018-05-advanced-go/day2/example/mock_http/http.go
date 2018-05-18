package mock_http

import (
	"net/http"
)

const ISO = "http://releases.ubuntu.com/18.04/ubuntu-18.04-desktop-amd64.iso"

func FetchISO() (*http.Response, error) {
	return http.Get(ISO)
}
