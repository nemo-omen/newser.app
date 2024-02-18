package util

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func MakeUrl(str string) (*url.URL, error) {
	if IsUrl(str) {
		u, err := url.ParseRequestURI(str)
		fmt.Println("already a URL")
		if err != nil {
			log.Printf("error creating url: %+v", err)
			return &url.URL{}, err
		}
		return u, nil
	}

	if !strings.Contains(str, ".") {
		return &url.URL{}, fmt.Errorf("no domain in %v, cannot create URL", str)
	}

	if !strings.Contains(str, "http://") && !strings.Contains(str, "https://") {
		log.Printf("attempting to prepend scheme to %v\n", str)
		// assuming https, because that's basically all anyone
		// can use, unless localhost
		str = "https://" + str
		log.Printf("new url string %v\n", str)
	}

	u, err := url.ParseRequestURI(str)

	if err != nil {
		return &url.URL{}, fmt.Errorf("unable to create URL from %v", str)
	}

	return u, nil
}
