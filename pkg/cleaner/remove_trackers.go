package cleaner

import (
	"regexp"
	"strings"
)

// RemoveTrackers removed utm_ trackers.
func RemoveTrackers(link string) string {
	utm := regexp.MustCompile("utm_[^&]+")
	link = utm.ReplaceAllString(link, "")

	repeatingAmp := regexp.MustCompile("&+")
	link = repeatingAmp.ReplaceAllString(link, "&")

	paramAmp := regexp.MustCompile(`\?&`)
	link = paramAmp.ReplaceAllString(link, "?")

	link = strings.Trim(link, "\n\r ?&")

	return link
}
