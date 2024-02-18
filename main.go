package main

import (
	"current/service"
	"fmt"
	"net/http"
	"os"

	"github.com/mmcdole/gofeed"
)

func main() {
	feeds := []*gofeed.Feed{}
	api := service.NewAPI(http.DefaultClient)
	links, err := api.FindFeedLinks("jeffcaldwell.is")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(links) == 0 {
		links, err = api.GuessFeedLinks("https://jeffcaldwell.is")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	for _, link := range links {
		fmt.Println("feed link: ", link)
		feed, err := api.GetFeed(link)
		if err == nil {
			feeds = append(feeds, feed)
		}
	}

	for _, feed := range feeds {
		fmt.Println("feed title: ", feed.Title)
	}
}
