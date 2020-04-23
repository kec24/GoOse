package goose

import (
	"github.com/pkg/errors"
)

// Goose is the main entry point of the program
type Goose struct {
	config Configuration
}

// New returns a new instance of the article extractor
func New(config Configuration) Goose {
	return Goose{
		config: config,
	}
}

// NewWithDefaults returns a new instance of the article extractor with default configuration settings
func NewWithDefaults() Goose {
	return Goose{
		config: GetDefaultConfiguration(),
	}
}

// ExtractFromURL follows the URL, fetches the HTML page and returns an article object
func (g Goose) ExtractFromURL(url string) (*Article, error) {
	HTMLRequester := NewHTMLRequester(g.config)
	html, err := HTMLRequester.fetchHTML(url)
	if err != nil {
		return nil, errors.Wrap(err, "could not get HTML from site")
	}
	cc := NewCrawler(g.config)
	return cc.Crawl(html, url)
}

// ExtractFromRawHTML returns an article object from the raw HTML content
func (g Goose) ExtractFromRawHTML(RawHTML string, url string) (*Article, error) {
	cc := NewCrawler(g.config)
	return cc.Crawl(RawHTML, url)
}
