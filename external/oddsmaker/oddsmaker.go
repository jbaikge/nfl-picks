package oddsmaker

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type Feed struct {
	XMLName xml.Name   `xml:"rss"`
	Items   []FeedItem `xml:"channel>item"`
}

type FeedItem struct {
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	PubDate string   `xml:"pubdate"`
}

type Odds struct {
	Home     TeamOdds
	Away     TeamOdds
	GameTime time.Time
}

type TeamOdds struct {
	Name   string
	Spread float64
	Total  float64
	Money  float64
}

const DateLayout = "Jan 02, 2006 03:04 PM"

var (
	Endpoint         = "http://www.referincome.com/odds/rss2/football_nfl.xml"
	LineRegexp       = regexp.MustCompile(`^([\w. ]+(?:-[AN])?) (-?\d+\.\d) [OU] \((\d+\.\d)\) (-?\d+\.\d) \| ([\w. ]+(?:-[AN])?) (-?\d+\.\d) [OU] \((\d+\.\d)\) (-?\d+\.\d) \(([^)]+)\)$`)
	ErrNoRegexpMatch = fmt.Errorf("Regexp did not match")
)

func CurrentFeed() (feed *Feed, err error) {
	resp, err := http.Get(Endpoint)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dec := xml.NewDecoder(resp.Body)
	feed = new(Feed)
	if err = dec.Decode(feed); err != nil {
		return
	}
	return
}

func CurrentOdds() (odds []Odds, err error) {
	f, err := CurrentFeed()
	if err != nil {
		return
	}
	odds = make([]Odds, 0, len(f.Items))
	for _, item := range f.Items {
		o, err := item.Odds()
		if err != nil {
			continue
		}
		odds = append(odds, o)
	}
	return
}

func (item *FeedItem) Odds() (o Odds, err error) {
	matches := LineRegexp.FindStringSubmatch(item.Title)
	if matches == nil {
		err = ErrNoRegexpMatch
		return
	}
	if o.Away, err = buildTeam(matches[1], matches[2], matches[3], matches[4]); err != nil {
		return
	}
	if o.Home, err = buildTeam(matches[5], matches[6], matches[7], matches[8]); err != nil {
		return
	}
	loc, err := time.LoadLocation("America/New_York")
	if o.GameTime, err = time.ParseInLocation(DateLayout, matches[9], loc); err != nil {
		return
	}
	return
}

func buildTeam(name, spread, total, money string) (team TeamOdds, err error) {
	if team.Name, err = translateName(name); err != nil {
		return
	}
	if team.Spread, err = strconv.ParseFloat(spread, 64); err != nil {
		return
	}
	if team.Total, err = strconv.ParseFloat(total, 64); err != nil {
		return
	}
	if team.Money, err = strconv.ParseFloat(money, 64); err != nil {
		return
	}
	return
}

func translateName(longName string) (id string, err error) {
	table := map[string]string{
		"Dallas":       "DAL",
		"Miami":        "MIA",
		"Kansas City":  "KC",
		"Tennessee":    "TEN",
		"Pittsburgh":   "PIT",
		"Atlanta":      "ATL",
		"Denver":       "DEN",
		"Chicago":      "CHI",
		"Carolina":     "CAR",
		"Jacksonville": "JAC",
		"Detroit":      "DET",
		"Tampa Bay":    "TB",
		"St. Louis":    "STL",
		"San Fran.":    "SF",
		"Arizona":      "ARI",
		"New York-N":   "NYG",
		"New York-A":   "NYJ",
		"Buffalo":      "BUF",
		"New Orleans":  "NO",
		"Baltimore":    "BAL",
		"Oakland":      "OAK",
		"Seattle":      "SEA",
		"Green Bay":    "GB",
		"Phila.":       "PHI",
	}
	var ok bool
	if id, ok = table[longName]; !ok {
		err = fmt.Errorf("Unkown team: %s", longName)
	}
	return
}
