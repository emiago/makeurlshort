package shorturl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShortUrlService(t *testing.T) {
	SetShortUrlService(&DefaultService{})
	u1 := UrlData{Url: "http://localhost", OriginalUrl: "http://example.com"}
	err := PostShortUrl(u1)
	require.Nil(t, err, "Failed to insert url data")

	urls, err := GetShortUrls()
	require.Nil(t, err, "Failed to get url data")
	assert.Equal(t, 1, len(urls))
	assert.Equal(t, u1.Url, urls[0].Url)
	assert.Equal(t, u1.OriginalUrl, urls[0].OriginalUrl)
}
