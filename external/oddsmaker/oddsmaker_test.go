package oddsmaker

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	testingEndpoint = TestingEndpoint()
	liveEndpoint    = Endpoint
)

func TestingEndpoint() string {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("football_nfl.xml")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		io.Copy(w, f)
	})
	server := httptest.NewServer(handler)
	return server.URL
}

func TestFeed(t *testing.T) {
	Endpoint = testingEndpoint

	f, err := CurrentFeed()
	if err != nil {
		t.Fatalf("CurrentFeed: %s", err)
	}

	for _, item := range f.Items[:len(f.Items)-1] {
		o, err := item.Odds()
		if err != nil {
			t.Errorf("item.Odds: %s", err)
			t.Logf("item: %+v", item)
		}
		t.Logf("%+v", o)
	}
}
