package urlshorty

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	SetParser(&DefaultParser{
		Domain: "http://localhost:8080/r",
	})

	_, err := Parse("bad url")
	assert.NotNil(t, err, "Parsing should fail")

	_, err = Parse("http://bad url")
	assert.NotNil(t, err, "Parsing should fail")

	_, err = Parse("http://example.com")
	assert.Nil(t, err, "Valid url does not work")

	surl, err := Parse("http://example.com?myquery=aaa")
	assert.Nil(t, err, "Valid url does not work")

	if !strings.HasPrefix(surl, "http://localhost:8080/r") {
		t.Error("Domain is not correct")
	}
}
