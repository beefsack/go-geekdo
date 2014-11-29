package bgg

import (
	"strconv"
	"time"
)

type BoardGames struct {
	TermsOfUse string      `xml:"termsofuse,attr"`
	BoardGames []BoardGame `xml:"item"`
}

type BoardGameName struct {
	Type      string `xml:"type,attr"`
	SortIndex int    `xml:"sortindex,attr"`
	Value     string `xml:"value,attr"`
}

type BoardGamePoll struct {
	Name       string                   `xml:"name,attr"`
	Title      string                   `xml:"title,attr"`
	TotalVotes int                      `xml:"totalvotes,attr"`
	ResultSets []BoardGamePollResultSet `xml:"results"`
}

type BoardGamePollResultSet struct {
	NumPlayers string                `xml:"numplayers,attr"`
	Results    []BoardGamePollResult `xml:"result"`
}

type BoardGamePollResult struct {
	Value    string `xml:"value,attr"`
	NumVotes int    `xml:"numvotes,attr"`
}

type BoardGameLink struct {
	Type  string `xml:"type,attr"`
	Id    int    `xml:"id,attr"`
	Value string `xml:"value,attr"`
}

type BoardGameVideos struct {
	Total  int              `xml:"total,attr"`
	Videos []BoardGameVideo `xml:"video"`
}

type BoardGameVideo struct {
	Id       int       `xml:"id,attr"`
	Title    string    `xml:"title,attr"`
	Category string    `xml:"category,attr"`
	Language string    `xml:"language,attr"`
	Link     string    `xml:"link,attr"`
	Username string    `xml:"username,attr"`
	UserId   int       `xml:"userid,attr"`
	PostDate time.Time `xml:"postdate,attr"`
}

type BoardGameVersion struct {
	Type          string          `xml:"type,attr"`
	Id            int             `xml:"id,attr"`
	Thumbnail     string          `xml:"thumbnail"`
	Image         string          `xml:"image"`
	Name          BoardGameName   `xml:"name"`
	YearPublished IntValue        `xml:"yearpublished"`
	ProductCode   StringValue     `xml:"productcode"`
	Width         FloatValue      `xml:"width"`
	Length        FloatValue      `xml:"length"`
	Depth         FloatValue      `xml:"depth"`
	Weight        FloatValue      `xml:"weight"`
	Links         []BoardGameLink `xml:"link"`
}

type BoardGameComments struct {
	Page       int                `xml:"page,attr"`
	TotalItems int                `xml:"totalitems,attr"`
	Comments   []BoardGameComment `xml:"comment"`
}

type BoardGameComment struct {
	Username string `xml:"username,attr"`
	Rating   string `xml:"rating,attr"`
	Value    string `xml:"value,attr"`
}

// IntRating returns the integer value for the rating, or -1 if the rating is
// N/A
func (bgc BoardGameComment) IntRating() int {
	i, err := strconv.Atoi(bgc.Rating)
	if err != nil {
		return -1
	}
	return i
}

type BoardGameStats struct {
	Page    int                     `xml:"page,attr"`
	Ratings []BoardGameStatsRatings `xml:"ratings"`
}

type BoardGameStatsRatings struct {
	Date          string               `xml:"date,attr"`
	UsersRated    IntValue             `xml:"usersrated"`
	Average       FloatValue           `xml:"average"`
	BayesAverage  FloatValue           `xml:"bayesaverage"`
	StdDev        FloatValue           `xml:"stddev"`
	Median        FloatValue           `xml:"median"`
	Owned         IntValue             `xml:"owned"`
	Trading       IntValue             `xml:"trading"`
	Wanting       IntValue             `xml:"wanting"`
	Wishing       IntValue             `xml:"wishing"`
	NumComments   IntValue             `xml:"numcomments"`
	NumWeights    IntValue             `xml:"numweights"`
	AverageWeight FloatValue           `xml:"averageweight"`
	Ranks         []BoardGameStatsRank `xml:"ranks>rank"`
}

// Time returns the ratings date as a time.Time object.
func (bgsr BoardGameStatsRatings) Time() (time.Time, error) {
	return time.Parse("20060102", bgsr.Date)
}

type BoardGameStatsRank struct {
	Type         string  `xml:"type,attr"`
	Id           int     `xml:"id,attr"`
	Name         string  `xml:"name,attr"`
	FriendlyName string  `xml:"friendlyname,attr"`
	Value        int     `xml:"value,attr"`
	BayesAverage float64 `xml:"bayesaverage,attr"`
}

type BoardGameMarketplaceListing struct {
	ListDate  StringValue                      `xml:"listdate"`
	Price     BoardGameMarketplaceListingPrice `xml:"price"`
	Condition StringValue                      `xml:"condition"`
	Notes     StringValue                      `xml:"notes"`
	Link      BoardGameMarketplaceListingLink  `xml:"link"`
}

func (bgml BoardGameMarketplaceListing) Time() (time.Time, error) {
	return time.Parse("Mon, 02 Jan 2006 15:04:05 +0700 ", bgml.ListDate.Value)
}

type BoardGameMarketplaceListingPrice struct {
	Currency string  `xml:"currency,attr"`
	Value    float64 `xml:"value,attr"`
}

type BoardGameMarketplaceListingLink struct {
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

type BoardGame struct {
	Type                string                        `xml:"type,attr"`
	Id                  int                           `xml:"id,attr"`
	Thumbnail           string                        `xml:"thumbnail"`
	Image               string                        `xml:"image"`
	Name                []BoardGameName               `xml:"name"`
	Description         string                        `xml:"description"`
	YearPublished       IntValue                      `xml:"yearpublished"`
	MinPlayers          IntValue                      `xml:"minplayers"`
	MaxPlayers          IntValue                      `xml:"maxplayers"`
	Polls               []BoardGamePoll               `xml:"poll"`
	PlayingTime         IntValue                      `xml:"playingtime"`
	MinAge              IntValue                      `xml:"minage"`
	Links               []BoardGameLink               `xml:"link"`
	Videos              BoardGameVideos               `xml:"videos"`
	Versions            []BoardGameVersion            `xml:"versions>item"`
	Comments            BoardGameComments             `xml:"comments"`
	Statistics          BoardGameStats                `xml:"statistics"`
	MarketplaceListings []BoardGameMarketplaceListing `xml:"marketplacelistings>listing"`
}
