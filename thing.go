package geekdo

import (
	"strconv"
	"time"
)

// Things are a collection of things.
type Things struct {
	TermsOfUse string  `xml:"termsofuse,attr"`
	Total      int     `xml:"total,attr"`
	Items      []Thing `xml:"item"`
}

// ThingName is the name of a thing and a type which indicates primary or not.
type ThingName struct {
	Type      string `xml:"type,attr"`
	SortIndex int    `xml:"sortindex,attr"`
	Value     string `xml:"value,attr"`
}

// ThingPoll is a poll for a thing, usually things like ideal number of players.
type ThingPoll struct {
	Name       string               `xml:"name,attr"`
	Title      string               `xml:"title,attr"`
	TotalVotes int                  `xml:"totalvotes,attr"`
	ResultSets []ThingPollResultSet `xml:"results"`
}

// ThingPollResultSet is a set of results for a poll.
type ThingPollResultSet struct {
	NumPlayers string            `xml:"numplayers,attr"`
	Results    []ThingPollResult `xml:"result"`
}

// ThingPollResult is a result in a result set for a poll.
type ThingPollResult struct {
	Value    string `xml:"value,attr"`
	NumVotes int    `xml:"numvotes,attr"`
}

// ThingLink is a link for a thing, such as to other things.
type ThingLink struct {
	Type  string `xml:"type,attr"`
	ID    int    `xml:"id,attr"`
	Value string `xml:"value,attr"`
}

// ThingVideos is a collection of videos for a thing.
type ThingVideos struct {
	Total  int          `xml:"total,attr"`
	Videos []ThingVideo `xml:"video"`
}

// ThingVideo is a video for a thing, usually on YouTube.
type ThingVideo struct {
	ID       int    `xml:"id,attr"`
	Title    string `xml:"title,attr"`
	Category string `xml:"category,attr"`
	Language string `xml:"language,attr"`
	Link     string `xml:"link,attr"`
	Username string `xml:"username,attr"`
	UserID   int    `xml:"userid,attr"`
	PostDate string `xml:"postdate,attr"`
}

// PostDateTime returns a time.Time object given the video post date.
func (bgv ThingVideo) PostDateTime() (time.Time, error) {
	return time.Parse(time.RFC3339, bgv.PostDate)
}

// ThingVersion is a version of a thing, such as a printing run for a board
// game.
type ThingVersion struct {
	Type          string      `xml:"type,attr"`
	ID            int         `xml:"id,attr"`
	Thumbnail     string      `xml:"thumbnail"`
	Image         string      `xml:"image"`
	Name          ThingName   `xml:"name"`
	YearPublished IntValue    `xml:"yearpublished"`
	Format        string      `xml:"format"`
	ReleaseDate   StringValue `xml:"releasedate"`
	ProductCode   StringValue `xml:"productcode"`
	SeriesCode    StringValue `xml:"seriescode"`
	PageCount     IntValue    `xml:"pagecount"`
	ISBN10        StringValue `xml:"isbn10"`
	ISBN13        StringValue `xml:"isbn13"`
	Width         FloatValue  `xml:"width"`
	Height        FloatValue  `xml:"height"`
	Length        FloatValue  `xml:"length"`
	Depth         FloatValue  `xml:"depth"`
	Weight        FloatValue  `xml:"weight"`
	Description   string      `xml:"description"`
	Links         []ThingLink `xml:"link"`
}

// ThingComments is a collection of comments for a thing, possibly including
// ratings.
type ThingComments struct {
	Page       int            `xml:"page,attr"`
	TotalItems int            `xml:"totalitems,attr"`
	Comments   []ThingComment `xml:"comment"`
}

// ThingComment is a comment for a thing from a user, possibly including a
// rating.
type ThingComment struct {
	Username string `xml:"username,attr"`
	Rating   string `xml:"rating,attr"`
	Value    string `xml:"value,attr"`
}

// IntRating returns the integer value for the rating, or -1 if the rating is
// N/A
func (bgc ThingComment) IntRating() int {
	i, err := strconv.Atoi(bgc.Rating)
	if err != nil {
		return -1
	}
	return i
}

// ThingStats is a collection of stats for a thing.
type ThingStats struct {
	Page    int                 `xml:"page,attr"`
	Ratings []ThingStatsRatings `xml:"ratings"`
}

// ThingStatsRatings are some aggregations of stats relating to a thing.
type ThingStatsRatings struct {
	Date          string           `xml:"date,attr"`
	UsersRated    IntValue         `xml:"usersrated"`
	Average       FloatValue       `xml:"average"`
	BayesAverage  FloatValue       `xml:"bayesaverage"`
	StdDev        FloatValue       `xml:"stddev"`
	Median        FloatValue       `xml:"median"`
	Owned         IntValue         `xml:"owned"`
	Trading       IntValue         `xml:"trading"`
	Wanting       IntValue         `xml:"wanting"`
	Wishing       IntValue         `xml:"wishing"`
	NumComments   IntValue         `xml:"numcomments"`
	NumWeights    IntValue         `xml:"numweights"`
	AverageWeight FloatValue       `xml:"averageweight"`
	Ranks         []ThingStatsRank `xml:"ranks>rank"`
}

// Time returns the ratings date as a time.Time object.
func (bgsr ThingStatsRatings) Time() (time.Time, error) {
	return time.Parse("20060102", bgsr.Date)
}

// ThingStatsRank is a rank for a thing given a type, such as a family.
type ThingStatsRank struct {
	Type         string  `xml:"type,attr"`
	ID           int     `xml:"id,attr"`
	Name         string  `xml:"name,attr"`
	FriendlyName string  `xml:"friendlyname,attr"`
	Value        int     `xml:"value,attr"`
	BayesAverage float64 `xml:"bayesaverage,attr"`
}

// ThingMarketplaceListing is a listing for a thing on the marketplace.
type ThingMarketplaceListing struct {
	ListDate  StringValue                  `xml:"listdate"`
	Price     ThingMarketplaceListingPrice `xml:"price"`
	Condition StringValue                  `xml:"condition"`
	Notes     StringValue                  `xml:"notes"`
	Link      ThingMarketplaceListingLink  `xml:"link"`
}

// Time returns the time the thing was listed on the marketplace.
func (bgml ThingMarketplaceListing) Time() (time.Time, error) {
	return time.Parse("Mon, 02 Jan 2006 15:04:05 +0700 ", bgml.ListDate.Value)
}

// ThingMarketplaceListingPrice is the price for a marketplace listing,
// including currency.
type ThingMarketplaceListingPrice struct {
	Currency string  `xml:"currency,attr"`
	Value    float64 `xml:"value,attr"`
}

// ThingMarketplaceListingLink is a link to a listing for a thing.
type ThingMarketplaceListingLink struct {
	Href  string `xml:"href,attr"`
	Title string `xml:"title,attr"`
}

// IntValue is a node wrapping an integer value.
type IntValue struct {
	Value int `xml:"value,attr"`
}

// StringValue is a node wrapping a string value.
type StringValue struct {
	Value string `xml:"value,attr"`
}

// FloatValue is a node wrapping a floating point value.
type FloatValue struct {
	Value float64 `xml:"value,attr"`
}

// Thing is a thing in Geekdo, such as a board game, a video game, an RPG etc.
type Thing struct {
	Type                string                    `xml:"type,attr"`
	ID                  int                       `xml:"id,attr"`
	Thumbnail           string                    `xml:"thumbnail"`
	Image               string                    `xml:"image"`
	Names               []ThingName               `xml:"name"`
	Description         string                    `xml:"description"`
	YearPublished       IntValue                  `xml:"yearpublished"`
	ReleaseDate         StringValue               `xml:"releasedate"`
	ProductCode         StringValue               `xml:"productcode"`
	MinPlayers          IntValue                  `xml:"minplayers"`
	MaxPlayers          IntValue                  `xml:"maxplayers"`
	Polls               []ThingPoll               `xml:"poll"`
	PlayingTime         IntValue                  `xml:"playingtime"`
	MinAge              IntValue                  `xml:"minage"`
	Links               []ThingLink               `xml:"link"`
	Videos              ThingVideos               `xml:"videos"`
	Versions            []ThingVersion            `xml:"versions>item"`
	Comments            ThingComments             `xml:"comments"`
	Statistics          ThingStats                `xml:"statistics"`
	MarketplaceListings []ThingMarketplaceListing `xml:"marketplacelistings>listing"`
}
