package goose

import (
	"net/http"
	"time"
)

const defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7"

// Configuration is a wrapper for various config options
type Configuration struct {
	LocalStoragePath        string //not used in this version
	ImagesMinBytes          int    //not used in this version
	TargetLanguage          string
	ImageMagickConvertPath  string //not used in this version
	ImageMagickIdentifyPath string //not used in this version
	BrowserUserAgent        string
	Debug                   bool
	ExtractPublishDate      bool
	AdditionalDataExtractor bool
	EnableImageFetching     bool
	UseMetaLanguage         bool

	//path to the stopwords folder
	StopWordsPath string
	StopWords     StopWords
	Parser        *Parser

	Timeout time.Duration

	Cookies []*http.Cookie
	Headers map[string]string
}

// GetDefaultConfiguration returns safe default configuration options
func GetDefaultConfiguration() Configuration {
	return Configuration{
		LocalStoragePath:        "",   //not used in this version
		ImagesMinBytes:          4500, //not used in this version
		EnableImageFetching:     true,
		UseMetaLanguage:         true,
		TargetLanguage:          "en",
		ImageMagickConvertPath:  "/usr/bin/convert",  //not used in this version
		ImageMagickIdentifyPath: "/usr/bin/identify", //not used in this version
		BrowserUserAgent:        defaultUserAgent,
		Debug:                   false,
		ExtractPublishDate:      true,
		AdditionalDataExtractor: false,
		StopWordsPath:           "resources/stopwords",
		StopWords:               NewStopwords(), //TODO with path
		Parser:                  NewParser(),
		Timeout:                 time.Duration(5 * time.Second),
		Cookies:                 []*http.Cookie{},
	}
}
