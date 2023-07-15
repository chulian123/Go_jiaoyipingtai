package tools

import (
	"fmt"
	"math/rand"
)

func Rand4Num() string {
	intn := rand.Intn(9999)
	if intn < 1000 {
		intn = intn + 1000
	}
	return fmt.Sprintf("%d", intn)
}
