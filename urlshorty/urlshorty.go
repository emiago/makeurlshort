package urlshorty

var (
	parser UrlShortener
)

func init() {
	parser = &DefaultParser{
		Domain: "localhost",
	}
}

func Parse(longurl string) (string, error) {
	return parser.Parse(longurl)
}
