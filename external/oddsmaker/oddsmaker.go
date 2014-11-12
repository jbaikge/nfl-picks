package oddsmaker

import (
	"encoding/xml"
	"fmt"
	"github.com/jbaikge/nfl-picks/picks"
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

type Line struct {
	Home     TeamLine
	Away     TeamLine
	GameTime time.Time
	Updated  time.Time
}

type TeamLine struct {
	Name      string
	Spread    float64
	OverUnder float64
	Money     float64
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

func CurrentLines() (lines []*picks.Line, err error) {
	f, err := CurrentFeed()
	if err != nil {
		return
	}
	lines = make([]*picks.Line, 0, len(f.Items))
	for _, item := range f.Items {
		l, err := item.Line()
		if err != nil {
			continue
		}
		line := &picks.Line{
			Spread:    l.Home.Spread,
			OverUnder: l.Home.OverUnder,
			Updated:   l.Updated,
		}
		lines = append(lines, line)
	}
	return
}

func (item *FeedItem) Line() (l Line, err error) {
	matches := LineRegexp.FindStringSubmatch(item.Title)
	if matches == nil {
		err = ErrNoRegexpMatch
		return
	}
	if l.Away, err = newTeam(matches[1], matches[2], matches[3], matches[4]); err != nil {
		return
	}
	if l.Home, err = newTeam(matches[5], matches[6], matches[7], matches[8]); err != nil {
		return
	}
	loc, err := time.LoadLocation("America/New_York")
	if l.GameTime, err = time.ParseInLocation(DateLayout, matches[9], loc); err != nil {
		return
	}
	if l.Updated, err = time.Parse(time.RFC1123Z, item.PubDate); err != nil {
		return
	}
	return
}

func newTeam(name, spread, total, money string) (team TeamLine, err error) {
	if team.Name, err = translateName(name); err != nil {
		return
	}
	if team.Spread, err = strconv.ParseFloat(spread, 64); err != nil {
		return
	}
	if team.OverUnder, err = strconv.ParseFloat(total, 64); err != nil {
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
