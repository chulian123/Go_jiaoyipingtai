package base

import (
	"fmt"
	"testing"
)

func TestBacode(t *testing.T) {
	var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	encode := Base58Encode(b58Alphabet)
	fmt.Println(string(encode))
}
