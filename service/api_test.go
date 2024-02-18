package service

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAPI_FindFeedLinks(t *testing.T) {
	// httptest.Server will automatically select a
	// port, but we need to send absolute urls
	// so, we define a listener
	l, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		log.Fatal(err)
	}

	// and initialize the server without starting it
	ts := httptest.NewUnstartedServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("request path: ", r.URL.Path)

			if r.URL.Path == "/feedlinks" {
				w.Write([]byte(LinksDoc))
			}

			if r.URL.Path == "/nofeedlinks" {
				w.Write([]byte(NoLinksDoc))
			}
		},
	))

	// make sure default listener is closed
	ts.Listener.Close()
	// and assign our listener to testserver
	ts.Listener = l
	ts.Start()

	defer ts.Close()

	type args struct {
		siteUrl string
	}

	api := NewAPI(ts.Client())
	baseUrl := "http://127.0.0.1:6666"

	tests := []struct {
		name    string
		s       *API
		args    args
		want    []string
		wantErr bool
	}{
		{
			"Test nolinks",
			api,
			args{siteUrl: baseUrl + "/nofeedlinks"},
			[]string(nil),
			false,
		},
		{
			"Test multiple links",
			api,
			args{siteUrl: baseUrl + "/feedlinks"},
			[]string{
				"https://fake.com/feed/atom",
				"https://fake.com/feed/rss",
				"https://fake.com/feed/json",
			},
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewAPI(http.DefaultClient)
			got, err := s.FindFeedLinks(tt.args.siteUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.FindFeedLinks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("API.DiscoverFeedLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
