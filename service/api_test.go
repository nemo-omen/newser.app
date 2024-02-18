package service

import (
	"current/common"
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

func TestAPI_GuessFeedLink(t *testing.T) {
	// TODO: Come up with better tests
	l, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		log.Fatal(err)
	}

	// and initialize the server without starting it
	ts := httptest.NewUnstartedServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/feed/atom" && r.URL.Path != "/feed/rss" && r.URL.Path != "/feed/json" {
				w.WriteHeader(404)
				w.Write([]byte("Not found"))
			}
			fmt.Println("request path: ", r.URL.Path)

			if r.URL.Path == "/feed/atom" {
				fmt.Println(r.URL.Path)
				w.Header().Set("Content-Type", string(common.XMLDoc))
				w.Write([]byte(DummyAtom))
			}

			if r.URL.Path == "/feed/rss" {
				fmt.Println(r.URL.Path)
				w.Header().Set("Content-Type", string(common.XMLDoc))
				w.Write([]byte(DummyRSS))
			}

			if r.URL.Path == "/feed/json" {
				fmt.Println(r.URL.Path)
				w.Header().Set("Content-Type", string(common.JSONDoc))
				w.Write([]byte(DummyJSON))
			}
		},
	))

	// make sure default listener is closed
	ts.Listener.Close()
	// and assign our listener to testserver
	ts.Listener = l
	ts.Start()

	defer ts.Close()

	api := NewAPI(ts.Client())
	baseUrl := "http://127.0.0.1:6666"

	type args struct {
		siteUrl string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Test link guessing",
			args{siteUrl: baseUrl},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.GuessFeedLinks(tt.args.siteUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.GuessFeedLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < 1 {
				t.Errorf("expected more than 0 links, got %v", len(got))
			}
		})
	}
}

func TestAPI_GetFeed(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		log.Fatal(err)
	}

	// and initialize the server without starting it
	ts := httptest.NewUnstartedServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/feed/atom" && r.URL.Path != "/feed/rss" && r.URL.Path != "/feed/json" {
				w.WriteHeader(404)
				w.Write([]byte("Not found"))
			}
			fmt.Println("request path: ", r.URL.Path)

			if r.URL.Path == "/feed/atom" {
				fmt.Println(r.URL.Path)
				w.Header().Set("Content-Type", string(common.XMLDoc))
				w.Write([]byte(DummyAtom))
			}

			if r.URL.Path == "/feed/rss" {
				fmt.Println(r.URL.Path)
				w.Header().Set("Content-Type", string(common.XMLDoc))
				w.Write([]byte(DummyRSS))
			}

			if r.URL.Path == "/feed/json" {
				fmt.Println(r.URL.Path)
				w.Header().Set("Content-Type", string(common.JSONDoc))
				w.Write([]byte(DummyJSON))
			}
		},
	))

	// make sure default listener is closed
	ts.Listener.Close()
	// and assign our listener to testserver
	ts.Listener = l
	ts.Start()

	defer ts.Close()

	api := NewAPI(ts.Client())
	baseUrl := "http://127.0.0.1:6666"

	type args struct {
		feedUrl string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test direct atom feed", args{feedUrl: baseUrl + "/feed/atom"}, "*gofeed.Feed", false},
		{"Test direct rss feed", args{feedUrl: baseUrl + "/feed/rss"}, "*gofeed.Feed", false},
		{"Test direct json feed", args{feedUrl: baseUrl + "/feed/json"}, "*gofeed.Feed", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			feed, err := api.GetFeed(tt.args.feedUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.GetFeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := reflect.TypeOf(feed).String()

			if got != tt.want {
				t.Errorf("expected %v, actual %v", tt.want, got)
			}

		})
	}
}
