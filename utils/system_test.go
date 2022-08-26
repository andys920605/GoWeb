package utils_test

import (
	"GoWeb/utils"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestConvertToDateTimeString(t *testing.T) {
	var now = time.Date(2022, time.June, 10, 13, 55, 30, 0, time.Local)
	var want = "2022-06-10 13:55:30"
	var got = utils.ConvertToDateTimeString(&now, "-")
	assert.Equal(t, want, got)
}

func TestConvertHoursToSeconds(t *testing.T) {
	hours := 3
	want := 3 * 3600
	got := utils.ConvertHoursToSeconds(hours)
	assert.Equal(t, want, got)
}
