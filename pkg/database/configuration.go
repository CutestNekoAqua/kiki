package database

import (
	"fmt"
	"strings"
)

// ConnectionDetails contains all the information about a Database Connection.
type ConnectionDetails struct {
	User     string
	Name     string
	Password string
	Host     string
	Port     int
}

// Configure bootstraps a Database Connection.
func Configure(details *ConnectionDetails) {
	var build = make([]string, 0)

	if details.User != "" {
		build = append(build, fmt.Sprintf("user=%s", details.User))
	}

	if details.Name != "" {
		build = append(build, fmt.Sprintf("dbname=%s", details.Name))
	}

	if details.Password != "" {
		build = append(build, fmt.Sprintf("password=%s", details.Password))
	}

	if details.Host != "" {
		build = append(build, fmt.Sprintf("host=%s", details.Host))
	}

	if details.Port > 0 {
		build = append(build, fmt.Sprintf("port=%d", details.Port))
	}

	connectionString = strings.Join(build, " ")
}
