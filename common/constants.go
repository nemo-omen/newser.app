package common

type ContentType string

const (
	Atom ContentType = "application/atom+xml"
	RSS  ContentType = "application/rss+xml"
	JSON             = "application/json"
)

var ValidContentTypes = []ContentType{
	Atom,
	RSS,
	JSON,
}

type DocContentType string

const (
	XMLDoc  = "application/xml; charset=UTF-8"
	JSONDoc = "application/json; charset=UTF-8"
)

var ValidDocContentTypes = []DocContentType{
	XMLDoc,
	JSONDoc,
}

var CommonFeedPaths = []string{
	"/feed",
	"/rss",
	"/rss.xml",
	"/feed.xml",
	"/feed/",
	"/rss/",
	"/rss.xml/",
	"/feed.xml/",
	"/atom",
	"/feed/atom",
	"/feed/rss",
	"/feed/json",
}

var CommonFeedExtensions = []string{
	".rss",
	".xml",
	".json",
	".rss/",
	".xml/",
	".json/",
}
