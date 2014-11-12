package nfl

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	server     = *httptest.Server
	timeLayout = "2006-01-02 15:04"
)

var samplePreExpect = []struct {
	Time string
	Home string
	Away string
}{
	{"2014-11-06 20:25", "CIN", "CLE"},
	{"2014-11-09 13:00", "BUF", "KC"},
	{"2014-11-09 13:00", "DET", "MIA"},
	{"2014-11-09 13:00", "JAC", "DAL"},
	{"2014-11-09 13:00", "NO", "SF"},
	{"2014-11-09 13:00", "BAL", "TEN"},
	{"2014-11-09 13:00", "NYJ", "PIT"},
	{"2014-11-09 13:00", "TB", "ATL"},
	{"2014-11-09 16:05", "OAK", "DEN"},
	{"2014-11-09 16:25", "ARI", "STL"},
	{"2014-11-09 16:25", "SEA", "NYG"},
	{"2014-11-09 20:30", "GB", "CHI"},
	{"2014-11-10 20:30", "PHI", "CAR"},
}

func TestingEndpoint() string {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filename = r.RequestURI[1:]
		f, err := os.Open(filename)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		io.Copy(w, f)
	})
	server := httptest.NewServer(handler)
	return server.URL
}

func TestSamplePreGame(t *testing.T) {
	ss := new(ScoreStrip)
	if err := xml.Unmarshal([]byte(samplePreGame), ss); err != nil {
		t.Fatalf("xml.Unmarshal: %s", err)
	}
	if v := ss.GameSet.Year; v != 2014 {
		t.Fatalf("Year invalid: %d", v)
	}
	if v := len(ss.GameSet.Games); v != 13 {
		t.Fatalf("len(games): %d; expected 13", v)
	}
	for i, g := range ss.GameSet.Games {
		exp := samplePreGameExpect[i]
		if v := g.Start().Format(timeLayout); exp.Time != v {
			t.Errorf("[%d] Time Exp: %s; Got: %s", i, exp.Time, v)
		}
		if v := g.Home; exp.Home != v {
			t.Errorf("[%d] Home Exp: %s; Got: %s", i, exp.Home, v)
		}
		if v := g.Away; exp.Away != v {
			t.Errorf("[%d] Away Exp: %s; Got: %s", i, exp.Away, v)
		}
	}
}

func TestSampleActiveGame(t *testing.T) {
	ss := new(ScoreStrip)
	if err := xml.Unmarshal([]byte(sampleActiveGame), ss); err != nil {
		t.Fatalf("xml.Unmarshal: %s", err)
	}
	g := ss.GameSet.Games[0]
	if v := g.Quarter; v != "2" {
		t.Fatalf("Quarter Exp %s Got %s", "2", v)
	}
	if v := g.TimeLeft(); v != time.Minute+3*time.Second {
		t.Fatalf("TimeLeft Got %s", v)
	}
	if v := g.Posession; v != "CIN" {
		t.Fatalf("Posession Got %s", v)
	}
}
