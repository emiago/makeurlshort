package urlshorty

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type UrlShortener interface {
	Parse(longurl string) (string, error)
}

type DefaultParser struct {
	Domain string
}

func (p *DefaultParser) Parse(longurl string) (string, error) {
	uid := uuid.NewV4()
	return fmt.Sprintf("http://%s/%s", p.Domain, uid.String()), nil
}
