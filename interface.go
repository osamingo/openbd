package openbd

type (
	Fetcher interface {
		Get(isbn ...string) (map[string]BookInformer, error)
	}
	BookInformer interface {
		ISBN() string
		Title() string
		Author() string
		Cover() string
		PublishedAt() string
		Publisher() string
		Series() string
		Volume() string
		ShortDescription() string
		Description() string
		JBPADescription() string
		TableOfContents() string
	}
)
