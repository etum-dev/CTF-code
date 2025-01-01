package cryptoencoding

import (
	"fmt"
	"math/big"
)

/*
message: HELLO
ascii bytes: [72, 69, 76, 76, 79]
hex bytes: [0x48, 0x45, 0x4c, 0x4c, 0x4f]
base-16: 0x48454c4c4f
base-10: 310400273487
*/
func ByteToInt(b []byte) *big.Int {
	result := new(big.Int)
	result.SetBytes(b)
	return result

}
func IntToByte(num *big.Int) []byte {
	return num.Bytes()
}

func ToBase10(message string) int {
	base10_int := 123
	return base10_int
}
func ToBase16(message string) *big.Int {
	big_num, ok := new(big.Int).SetString(message, 16)
	if !ok {
		fmt.Println("error in base16 conversion")
	}
	return big_num
}

/*func ToHexBytes(message string) []string {

}*/
