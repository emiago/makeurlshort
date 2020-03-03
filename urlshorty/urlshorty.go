package urlshorty

import "net/url"

type UrlShortener interface {
	Parse(longurl string) (string, error)
}

var (
	parser UrlShortener
)

func init() {
	parser = &DefaultParser{
		Domain: "http://localhost:8080",
	}
}

func SetParser(p UrlShortener) {
	parser = p
}

func Parse(longurl string) (string, error) {
	_, err := url.ParseRequestURI(longurl)
	if err != nil {
		return "", err
	}
	return parser.Parse(longurl)
}
