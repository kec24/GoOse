package goose

import (
	resty "github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// HTMLRequester can fetch the target HTML page
type HTMLRequester interface {
	fetchHTML(string) (string, error)
}

type htmlrequester struct {
	config Configuration
}

// NewHTMLRequester returns a crawler object initialised with the URL and the [optional] raw HTML body
func NewHTMLRequester(config Configuration) HTMLRequester {
	return htmlrequester{
		config: config,
	}
}

func (hr htmlrequester) fetchHTML(url string) (string, error) {
	client := resty.New()
	client.SetTimeout(hr.config.Timeout)
	resp, err := client.
		SetCookies(hr.config.Cookies).
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", hr.config.BrowserUserAgent).
		Get(url)

	if err != nil {
		return "", errors.Wrap(err, "could not perform request on "+url)
	}
	if resp.IsError() {
		return "", &badRequest{Message: "could not perform request with " + url + " status code " + string(resp.StatusCode())}
	}
	return resp.String(), nil
}

type badRequest struct {
	Message string `json:"message,omitempty"`
}

func (BadRequest *badRequest) Error() string {
	return "Required request fields are not filled"
}
