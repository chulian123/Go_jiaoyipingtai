package test

import (
	"fmt"
	"math/rand"

	"testing"
	"time"
)

//4位code函数生成测试

func TestGen4Code(t *testing.T) {
	n := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)
	fmt.Println(n)

}
