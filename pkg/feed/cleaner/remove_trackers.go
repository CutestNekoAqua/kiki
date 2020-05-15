package cleaner

import (
	"regexp"
	"strings"
)

// RemoveTrackers removed utm_ trackers.
func RemoveTrackers(link string) string {
	utm := regexp.MustCompile("utm_[^&]+")
	link = utm.ReplaceAllString(link, "")

	repetingAmp := regexp.MustCompile("&+")
	link = repetingAmp.ReplaceAllString(link, "&")

	link = strings.Trim(link, "\n\r ?&")

	return link
}
