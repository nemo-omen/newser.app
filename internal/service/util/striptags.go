package util

func StripTags(htmlString string) string {
	data := []rune{}
	outerHTML := false

	for _, c := range htmlString {
		if c == '<' {
			outerHTML = true
		}

		if c == '>' {
			outerHTML = false
		}

		if !outerHTML {
			data = append(data, c)
		}
	}
	return string(data)
}
