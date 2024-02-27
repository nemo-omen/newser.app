package service

var LinksDoc string = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8"/>
		<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
		<link rel="alternate" href="/feed/atom" type="application/atom+xml" title="Atom Relative"/>
		<link rel="alternate" href="https://fake.com/feed/atom" type="application/atom+xml" title="Atom"/>
		<link rel="alternate" href="https://fake.com/feed/rss" type="application/rss+xml" title="RSS"/>
		<link rel="alternate" href="https://fake.com/feed/json" type="application/json" title="JSON"/>
		<link rel="stylesheet" href="/public/style/main.css"/>
		<title>Some Fake Title</title>
		</head>
		<body>
		<h1>Hello, links!</h1>
		</body>
		</html>
		`

var NoLinksDoc string = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8"/>
		<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
		<link rel="alternate" href="https://fake.com/feed/json" type="application/whatever" title="Whatever"/>
		<link rel="stylesheet" href="/public/style/main.css"/>
		<link rel="me" href="https://mastodon.com/@fakeperson"/>
		<title>Some Fake Title</title>
	</head>
	<body>
		<h1>Hello, links!</h1>
	</body>
	</html>`

var DummyRSS = `<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom" xmlns:rss="https://www.rssboard.org/rss-specification">
    <title>Test Title</title>
    <link>https://dummysite.com/feed/atom</link>
    <description>Test description</description>
    <pubDate>Wed, 14 Feb 2024 00:00:00 -0600</pubDate>
    <image>
        <url>https://dummysite.com/favicon.svg</url>
        <title>Test Image Title</title>
        <link>https://dummysite.com</link>
        <width>32</width>
        <height>32</height>
    </image>
    <item>
        <title>Test Item</title>
        <link>https://dummysite.com/blog/test-item</link>
        <description>This is an item description.</description>
        <content:encoded><![CDATA[<p>Test content</p>]]></content:encoded>
        <author>Some Author</author>
        <guid>https://dummysite.com/blog/test-item</guid>
        <pubDate>Sun, 11 Feb 2024 00:00:00 -0600</pubDate>
    </item>
</feed>`

var DummyAtom = `<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet href="/public/style/feed.xsl" type="text/xsl"?>
<feed xmlns="http://www.w3.org/2005/Atom">
    <title>Test Feed</title>
    <id>https://dummysite.com/feed/atom</id>
    <updated>2024-02-14T00:00:00-06:00</updated>
    <icon>https://dummysite.com/public/favicon.svg</icon>
    <logo>https://dummysite.com/public/favicon.svg</logo>
    <subtitle>Test description</subtitle>
    <link href="https://dummysite.com/feed/atom" rel="self"></link>
    <entry>
        <title>Test Item</title>
        <updated>2024-02-11T00:00:00-06:00</updated>
        <id>https://dummysite.com/blog/test-item</id>
        <content type="html">&lt;pTest content&lt;/p&gt;&#xA;</content>
        <link href="https://dummysite.com/blog/test-item" rel="alternate"></link>
        <link href="https://dummysite.com" rel="related" type="text/html"></link>
        <summary type="html">This is an item description.</summary>
        <author>
            <name>Some Author</name>
        </author>
    </entry>
</feed>`

var DummyJSON = `{
  "version": "https://jsonfeed.org/version/1.1",
  "title": "Test Feed",
  "home_page_url": "https://dummysite.com",
  "feed_url": "https://dummysite.com/feed/json",
  "description": "Test description",
  "favicon": "https://dummysite.com/public/favicon.svg",
  "items": [
    {
      "id": "https://dummysite.com/blog/test-item",
      "url": "https://dummysite.com/blog/test-item",
      "title": "Test Item",
      "content_html": "\u003cp\u003eTest content\u003c/p\u003e\n",
      "summary": "This is an item description.",
      "date_published": "2024-02-11T00:00:00-06:00",
      "date_modified": "2024-02-11T00:00:00-06:00",
      "author": {
        "name": "Some Author"
      },
      "authors": [
        {
          "name": "Some Author"
        }
      ]
    }
  ]
}`
