package geekdo

// CollectionItems is the root node of a collection request.
type CollectionItems struct {
	TotalItems int              `xml:"totalitems,attr"`
	TermsOfUse string           `xml:"termsofuse,attr"`
	PubDate    string           `xml:"pubdate,attr"`
	Items      []CollectionItem `xml:"item"`
}

// CollectionItem is an item inside a collection, such as a game.
type CollectionItem struct {
	ObjectType    string               `xml:"objecttype,attr"`
	ObjectID      int                  `xml:"objectid,attr"`
	SubType       string               `xml:"subtype,attr"`
	CollID        int                  `xml:"collid,attr"`
	Name          CollectionItemName   `xml:"name"`
	YearPublished int                  `xml:"yearpublished"`
	Image         string               `xml:"image"`
	Thumbnail     string               `xml:"thumbnail"`
	Stats         CollectionItemStats  `xml:"stats"`
	Status        CollectionItemStatus `xml:"status"`
	NumPlays      int                  `xml:"numplays"`
}

// CollectionItemName is a name and sort index for a collection item.
type CollectionItemName struct {
	SortIndex int    `xml:"sortindex,attr"`
	Value     string `xml:",chardata"`
}

// CollectionItemStatus describes the status of the game for the subject user.
type CollectionItemStatus struct {
	Own          bool   `xml:"own,attr"`
	PrevOwned    bool   `xml:"prevowned,attr"`
	ForTrade     bool   `xml:"fortrade,attr"`
	Want         bool   `xml:"want,attr"`
	WantToPlay   bool   `xml:"wanttoplay,attr"`
	WantToBuy    bool   `xml:"wanttobuy,attr"`
	Wishlist     bool   `xml:"wishlist,attr"`
	Preordered   bool   `xml:"preordered,attr"`
	LastModified string `xml:"lastmodified,attr"`
}

// CollectionItemStats are stats related to the item.
type CollectionItemStats struct {
	MinPlayers  int                       `xml:"minplayers,attr"`
	MaxPlayers  int                       `xml:"maxplayers,attr"`
	PlayingTime int                       `xml:"playingtime,attr"`
	NumOwned    int                       `xml:"numowned,attr"`
	Rating      CollectionItemStatsRating `xml:"rating"`
}

// CollectionItemStatsRating is a rating and stats regarding ratings.
type CollectionItemStatsRating struct {
	Value        int              `xml:"value,attr"`
	UsersRated   IntValue         `xml:"usersrated"`
	Average      FloatValue       `xml:"average"`
	BayesAverage FloatValue       `xml:"bayesaverage"`
	StdDev       FloatValue       `xml:"stddev"`
	Median       IntValue         `xml:"median"`
	Ranks        []ThingStatsRank `xml:"ranks>rank"`
}
