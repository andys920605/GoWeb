package utils

import (
	"fmt"
	"time"
)

var ConfigPath string = "local" // default win-env

// Get config path for local or docker
func GetConfigPath() string {
	if ConfigPath == "docker" {
		return "configs/config-docker"
	} else if ConfigPath == "local" {
		return "configs/config" // localhost
	} else if ConfigPath == "example" {
		return "../../configs/config"
	}
	return "config-local" // release - win-env
}

// convert to date time string
// @param time
// @param separator
// @result date time string
func ConvertToDateTimeString(now *time.Time, separator string) string {
	if now == nil {
		return ""
	}
	format := fmt.Sprintf("2006%s01%s02 15:04:05", separator, separator)
	return now.Format(format)
}

// convert hours to seconds
// @hours hours
// @result seconds
func ConvertHoursToSeconds(hours int) int {
	return int((time.Duration(hours) * time.Hour).Seconds())
}
