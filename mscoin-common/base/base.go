package base

import (
	"bytes"
	"math/big"
)

//如果我们要进行BTC钱包的充值提现等，首先必须有一个钱包地址，类似于：`1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa`，
//这是BTC的第一个地址，我们可以通过`https://www.blockchain.com/explorer/addresses/btc/1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa` 进行查看。
//> 所谓的地址，其实就是一个公钥，也就是说地址是通过一系列的加密算法得到的，所以我们就可以通过这样的规则得到一个属于我们的BTC地址
//### Base58加密
//Base58是用于比特币中使用的一种独特的编码方式，主要用于产生比特币的钱包地址。相比的Base64，Base58不使用数字 “0”，字母大写 “O”，字母大写 “I”，和字母小写 “L”，以及 “+” 和 “/” 符号。
//
//设计Base58主要的目的是：避免混淆。在某些字体下，数字0和字母大写O，以及字母大写我和字母小写升会非常相似。 不使用 “+” 和 “/” 的原因是非字母或数字的字符串作为帐号较难被接受。
//但是这个base58的计算量比BASE64的计算量多了很多。因为58不是2的整数倍，需要不断用除法去计算。而且长度也比的base64稍微多了一点。

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// ReverseBytes 字节数组反转
func ReverseBytes(data []byte) {

	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {

		data[i], data[j] = data[j], data[i]
	}
}

func Base58Encode(input []byte) []byte {
	var result []byte
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}
	ReverseBytes(result)
	for _, b := range input {
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}
	return result
}

// Base58Decode Base58转字节数组，解密
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0
	for _, b := range input {
		if b == 0x00 {
			zeroBytes++
		}
	}
	payload := input[zeroBytes:]
	for _, b := range payload {
		charIndex := bytes.IndexByte(b58Alphabet, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}
	decoded := result.Bytes()
	//decoded...表示将decoded所有字节追加
	//decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)
	return decoded
}
