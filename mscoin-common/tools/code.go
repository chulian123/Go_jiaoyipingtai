package tools

import (
	"math/rand"
	"strconv"
	"time"
)

// Rand4Code 生辰验证码函数
func Rand4Code() string {
	rand.NewSource(time.Now().UnixNano())
	cod := rand.Intn(9999) + 1000
	code := strconv.Itoa(cod)
	return code
}
