package shorturl

type DefaultService struct {
	shortUrls []UrlData
}

func (s *DefaultService) GetShortUrls() ([]UrlData, error) {
	return s.shortUrls, nil
}

func (s *DefaultService) PostShortUrl(d UrlData) error {
	s.shortUrls = append(s.shortUrls, d)
	return nil
}
