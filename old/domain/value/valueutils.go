package value

import "regexp"

var AllowedCharsRX = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

func IsAllowedChars(s string) bool {
	return AllowedCharsRX.MatchString(s)
}

func IsLongerThanMinChars(s string, min int) bool {
	return len(s) >= min
}

func IsShorterThanMaxChars(s string, max int) bool {
	return len(s) <= max
}

func IsValidLength(s string, min, max int) bool {
	return len(s) >= min && len(s) <= max
}
