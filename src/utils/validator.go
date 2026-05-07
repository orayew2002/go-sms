package utils

import "regexp"

func IsPhone(phone string) bool {
	var pattern = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	match := pattern.FindString(phone)

	return match != ""
}
