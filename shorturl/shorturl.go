package shorturl

type UrlData struct {
	Url         string
	OriginalUrl string
}

type ShortUrlService interface {
	GetShortUrls() ([]UrlData, error)
	PostShortUrl(data UrlData) error
}

var (
	service ShortUrlService
)

func init() {
	service = &DefaultService{}
}

func SetShortUrlService(s ShortUrlService) {
	service = s
}

func GetShortUrls() ([]UrlData, error) {
	return service.GetShortUrls()
}

func PostShortUrl(s UrlData) error {
	return service.PostShortUrl(s)
}
