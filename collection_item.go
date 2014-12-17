package geekdo

// CollectionItem is a row in an advanced search result.
type CollectionItem struct {
	ID           int
	Kind         string
	Rank         int
	Thumbnail    string
	Name         string
	URL          string
	Year         int
	BayesAverage float64
	Average      float64
	UsersRated   int
}
