package handler

import "fmt"

var (
	ErrorNotLoggedIn  = "You are not logged in"
	ErrorFeedNotFound = func(resource string) string { return fmt.Sprintf("Feed not found at %v", resource) }
)
