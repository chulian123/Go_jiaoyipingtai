package tools

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	milli := time.UnixMilli(1687968000000)
	fmt.Println(milli.String())
	zeroTime := ZeroTime()
	unixMilli := time.UnixMilli(zeroTime)
	fmt.Println(unixMilli.String())
}
