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
	</html>
`

var FakeRSSFeed = ``

var FakeAtomFeed = ``

var FakeJSONFeed = ``
