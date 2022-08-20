package utils_test

import (
	"regexp"
	"testing"

	"GoWeb/utils"
)

func TestRandomInt(t *testing.T) {

}

// random umber test
func TestRandomNum(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "digital length is 1",
			want: 1,
		},
		{
			name: "digital length is 5",
			want: 5,
		},
		{
			name: "digital length is 10",
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := len([]rune(utils.RandomNum(tt.want)))
			if got != tt.want {
				t.Error("wrong result")
			}
		})
	}
}

func TestRandomString(t *testing.T) {

}

// random test
func TestRand(t *testing.T) {
	tests := []struct {
		name string
		kind int
		want int
	}{
		{
			name: "digital",
			kind: utils.RAND_KIND_NUM,
			want: 1,
		},
		{
			name: "digital length is 5",
			kind: utils.RAND_KIND_NUM,
			want: 5,
		},
		{
			name: "lower-case letters",
			kind: utils.RAND_KIND_LOWER,
			want: 1,
		},
		{
			name: "lower-case letters length is 10",
			kind: utils.RAND_KIND_LOWER,
			want: 10,
		},
		{
			name: "upper-case letters",
			kind: utils.RAND_KIND_UPPER,
			want: 1,
		},
		{
			name: "upper-case letters length is 15",
			kind: utils.RAND_KIND_UPPER,
			want: 15,
		},
		{
			name: "digital and alphabet",
			kind: utils.RAND_KIND_ALL,
			want: 1,
		},
		{
			name: "digital and alphabet length is 15",
			kind: utils.RAND_KIND_ALL,
			want: 20,
		},
	}
	var digitCheck = regexp.MustCompile(`^[0-9]+$`)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.kind {
			case utils.RAND_KIND_NUM:
				digitCheck = regexp.MustCompile(`^[0-9]+$`)
				got := len([]rune(utils.RandomNum(tt.want)))
				if true != digitCheck.MatchString(utils.Rand(tt.want, utils.RAND_KIND_NUM)) && got != tt.want {
					t.Error("wrong result")
				}
			case utils.RAND_KIND_LOWER:
				digitCheck = regexp.MustCompile(`^[a-z]+$`)
				got := len([]rune(utils.RandomNum(tt.want)))
				if true != digitCheck.MatchString(utils.Rand(tt.want, utils.RAND_KIND_LOWER)) && got != tt.want {
					t.Error("wrong result")
				}
			case utils.RAND_KIND_UPPER:
				digitCheck = regexp.MustCompile(`^[A-Z]+$`)
				got := len([]rune(utils.RandomNum(tt.want)))
				if true != digitCheck.MatchString(utils.Rand(tt.want, utils.RAND_KIND_UPPER)) && got != tt.want {
					t.Error("wrong result")
				}
			default:
				digitCheck = regexp.MustCompile(`\w`)
				got := len([]rune(utils.RandomNum(tt.want)))
				if true != digitCheck.MatchString(utils.Rand(tt.want, utils.RAND_KIND_ALL)) && got != tt.want {
					t.Error("wrong result")
				}
			}
		})
	}
}

func TestRandomMoney(t *testing.T) {

}

func TestRandomEmail(t *testing.T) {

}

// random range number test
func TestGenerateRangeNum(t *testing.T) {
	tests := []struct {
		name string
		min  int
		max  int
	}{
		{
			name: "test 1",
			min:  10,
			max:  30,
		},
		{
			name: "test 2",
			min:  10,
			max:  20,
		},
		{
			name: "test 3",
			min:  50,
			max:  100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.GenerateRangeNum(tt.min, tt.max)
			if got > tt.max || got < tt.min {
				t.Error("wrong result")
			}
		})
	}
}
