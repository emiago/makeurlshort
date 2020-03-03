
package shorturl 

type ShortUrlService interface{} {
	GetShortUrls() ([]UrlData, error)
	PostShortUrl(s UrlData) (error)
}

type ShortUrlsInMemory struct {
	shortUrls []ShortUrl
}

func (s *ShortUrlsInMemory) 