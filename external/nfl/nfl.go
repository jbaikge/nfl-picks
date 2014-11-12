package nfl

import (
	"encoding/xml"
	"fmt"
	"github.com/jbaikge/nfl-picks/picks"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type ScoreStrip struct {
	XMLName xml.Name `xml:"ss"`
	GameSet GameSet  `xml:"gms"`
}

type GameSet struct {
	XMLName xml.Name `xml:"gms"`
	Week    int      `xml:"w,attr"`
	Year    int      `xml:"y,attr"`
	Season  string   `xml:"t,attr"`
	Games   []Game   `xml:"g"`
}

type Game struct {
	XMLName   xml.Name `xml:"g"`
	GameId    int      `xml:"gsis,attr"`
	EventId   string   `xml:"eid,attr"`
	Clock     string   `xml:"k,attr"`
	Posession string   `xml:"p,attr"`
	Time      string   `xml:"t,attr"`
	Home      string   `xml:"h,attr"`
	HomeScore int      `xml:"hs,attr"`
	Away      string   `xml:"v,attr"`
	AwayScore int      `xml:"vs,attr"`
	Quarter   string   `xml:"q,attr"`
}

var (
	DataEndpoint = "http://www.nfl.com/ajax/scorestrip"
	LiveEndpoint = "http://www.nfl.com/liveupdate/scorestrip/ss.xml"
)

func CurrentGames() (year, week int, games []*picks.Game, err error) {
	return getGames(LiveEndpoint)
}

func GamesFor(year, week int) (games []*picks.Game, err error) {
	u, err := dataURL(year, "REG", week)
	if err != nil {
		return
	}
	_, _, games, err = getGames(u.String())
	return
}

func (g Game) Start() (t time.Time) {
	loc, _ := time.LoadLocation("America/New_York")
	var hour, minute int
	fmt.Sscanf(g.Time, "%d:%d", &hour, &minute)
	if hour != 9 && hour != 12 {
		hour += 12
	}
	year, _ := strconv.Atoi(g.EventId[0:4])
	month, _ := strconv.Atoi(g.EventId[4:6])
	day, _ := strconv.Atoi(g.EventId[6:8])
	t = time.Date(year, time.Month(month), day, hour, minute, 0, 0, loc)
	return
}

func (g Game) TimeLeft() (d time.Duration) {
	if g.Clock == "" {
		return
	}
	var min, sec time.Duration
	if _, err := fmt.Sscanf(g.Clock, "%d:%d", &min, &sec); err != nil {
		return
	}
	return min*time.Minute + sec*time.Second
}

func dataURL(seasonYear int, seasonType string, week int) (u *url.URL, err error) {
	if u, err = url.Parse(DataEndpoint); err != nil {
		return
	}
	values := u.Query()
	values.Set("season", fmt.Sprint(seasonYear))
	values.Set("seasonType", seasonType)
	values.Set("week", fmt.Sprint(week))
	u.RawQuery = values.Encode()
	return
}

func getGames(url string) (year, week int, games []*picks.Game, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dec := xml.NewDecoder(resp.Body)
	ss := new(ScoreStrip)
	if err = dec.Decode(ss); err != nil {
		return
	}

	week = ss.GameSet.Week
	year = ss.GameSet.Year
	nflGames := ss.GameSet.Games

	games = make([]*picks.Game, len(nflGames))
	for i, ng := range nflGames {
		games[i] = &picks.Game{
			Id:        picks.GameId(ng.Away, ng.Home, ng.Start()),
			Year:      year,
			Week:      week,
			Start:     ng.Start(),
			TimeLeft:  ng.TimeLeft(),
			Posession: ng.Posession,
			Home:      ng.Home,
			HomeScore: ng.HomeScore,
			Away:      ng.Away,
			AwayScore: ng.AwayScore,
			Quarter:   picks.Quarter(ng.Quarter),
		}
	}
	return
}
