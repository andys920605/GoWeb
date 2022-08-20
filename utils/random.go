package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// region constant
// random kind enum
const (
	RAND_KIND_NUM   = iota // 0: 純數字
	RAND_KIND_LOWER        // 1: 小寫字母
	RAND_KIND_UPPER        // 2: 大寫字母
	RAND_KIND_ALL          // 3: 數字、大小寫字母
)
const alphabet = "abcdefghijklmnopqrstuvwxyz"

// endregion

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
} // Austin 20220616

// random number
// @param length 幾位數
func RandomNum(length int) string {
	result := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		num := rand.Intn(10)
		result = result + strconv.Itoa(num)
	}
	return result
} // Austin 20210727

// generates a random string of length n
// @param n length
// @result string
func RandomString(n int) string {
	var sb strings.Builder
	length := len(alphabet)
	for idx := 0; idx < n; idx++ {
		character := alphabet[rand.Intn(length)]
		sb.WriteByte(character)
	}
	return sb.String()
}

// generates a bool value
// @result bool value
func RandomBool() bool {
	return rand.Intn(2) == 1
}

// 隨機字串
// @param size 幾位數
// @param kind 種類enum
// @return string
func Rand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", Rand(10, RAND_KIND_ALL))
}

// generateRangeNum 生成一個區間範圍的隨機數
// @param min 最小值
// @param max 最大值
// @result 區間範圍的隨機數
func GenerateRangeNum(min, max int) int {
	return min + rand.Intn(max-min+1)
}
