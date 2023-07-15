package tools

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	milli := time.UnixMilli(1678377600000)
	fmt.Println(milli.String())
	zeroTime := ZeroTime()
	unixMilli := time.UnixMilli(zeroTime)
	fmt.Println(unixMilli.String())
}

func TestContext(t *testing.T) {
	//上下文的使用  对追踪非常有用
	ctx := context.Background()
	ctx = context.WithValue(ctx, "traceId", "AAA")
	BB(ctx)
}

func BB(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	value := ctx.Value("traceId")
	ctx = context.WithValue(ctx, "traceId", fmt.Sprintf("%v", value)+"_BBB")
	CC(ctx)
}

func CC(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	value := ctx.Value("traceId")
	ctx = context.WithValue(ctx, "traceId", fmt.Sprintf("%v", value)+"_CCC")
	fmt.Println(ctx.Value("traceId"))
}
