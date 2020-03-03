package shorturl

type UrlData struct {
	Url         string
	OriginalUrl string
}

var (
	service ShortUrlService
)

func SetShortUrlService(s ShortUrlService) {
	service = s
}

func GetShortUrls() ([]UrlData, error) {
	return service.GetShortUrls()
}

func PostShortUrl(s UrlData) error {
	return service.PostShortUrl(s)
