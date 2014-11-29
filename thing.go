package geekdo

import (
	"strconv"
	"time"
)

type Things struct {
	TermsOfUse string  `xml:"termsofuse,attr"`
	Items      []Thing `xml:"item"`
}

type ThingName struct {
	Type      string `xml:"type,attr"`
	SortIndex int    `xml:"sortindex,attr"`
	Value     string `xml:"value,attr"`
}

type ThingPoll struct {
	Name       string               `xml:"name,attr"`
	Title      string               `xml:"title,attr"`
	TotalVotes int                  `xml:"totalvotes,attr"`
	ResultSets []ThingPollResultSet `xml:"results"`
}

type ThingPollResultSet struct {
	NumPlayers string            `xml:"numplayers,attr"`
	Results    []ThingPollResult `xml:"result"`
}

type ThingPollResult struct {
	Value    string `xml:"value,attr"`
	NumVotes int    `xml:"numvotes,attr"`
}

type ThingLink struct {
	Type  string `xml:"type,attr"`
	Id    int    `xml:"id,attr"`
	Value string `xml:"value,attr"`
}

type ThingVideos struct {
	Total  int          `xml:"total,attr"`
	Videos []ThingVideo `xml:"video"`
}

type ThingVideo struct {
	Id       int    `xml:"id,attr"`
	Title    string `xml:"title,attr"`
	Category string `xml:"category,attr"`
	Language string `xml:"language,attr"`
	Link     string `xml:"link,attr"`
	Username string `xml:"username,attr"`
	UserId   int    `xml:"userid,attr"`
	PostDate string `xml:"postdate,attr"`
}

func (bgv ThingVideo) PostDateTime() (time.Time, error) {
	return time.Parse(time.RFC3339, bgv.PostDate)
}

type ThingVersion struct {
	Type          string      `xml:"type,attr"`
	Id            int         `xml:"id,attr"`
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

type ThingComments struct {
	Page       int            `xml:"page,attr"`
	TotalItems int            `xml:"totalitems,attr"`
	Comments   []ThingComment `xml:"comment"`
}

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

type ThingStats struct {
	Page    int                 `xml:"page,attr"`
	Ratings []ThingStatsRatings `xml:"ratings"`
}

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

type ThingStatsRank struct {
	Type         string  `xml:"type,attr"`
	Id           int     `xml:"id,attr"`
	Name         string  `xml:"name,attr"`
	FriendlyName string  `xml:"friendlyname,attr"`
	Value        int     `xml:"value,attr"`
	BayesAverage float64 `xml:"bayesaverage,attr"`
}

type ThingMarketplaceListing struct {
	ListDate  StringValue                  `xml:"listdate"`
	Price     ThingMarketplaceListingPrice `xml:"price"`
	Condition StringValue                  `xml:"condition"`
	Notes     StringValue                  `xml:"notes"`
	Link      ThingMarketplaceListingLink  `xml:"link"`
}

func (bgml ThingMarketplaceListing) Time() (time.Time, error) {
	return time.Parse("Mon, 02 Jan 2006 15:04:05 +0700 ", bgml.ListDate.Value)
}

type ThingMarketplaceListingPrice struct {
	Currency string  `xml:"currency,attr"`
	Value    float64 `xml:"value,attr"`
}

type ThingMarketplaceListingLink struct {
	Href  string `xml:"href,attr"`
	Title string `xml:"title,attr"`
}

type IntValue struct {
	Value int `xml:"value,attr"`
}

type StringValue struct {
	Value string `xml:"value,attr"`
}

type FloatValue struct {
	Value float64 `xml:"value,attr"`
}

type Thing struct {
	Type                string                    `xml:"type,attr"`
	Id                  int                       `xml:"id,attr"`
	Thumbnail           string                    `xml:"thumbnail"`
	Image               string                    `xml:"image"`
	Name                []ThingName               `xml:"name"`
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
