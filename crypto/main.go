package main

import (
	cryptoencoding "etum-dev/chall/chall"
	"fmt"
	"math/big"
)

func main() {
	bytes_and_big_integers := "11515195063862318899931685488813747395775516287289682636499965282714637259206269"
	n, ok := new(big.Int).SetString(bytes_and_big_integers, 0)
	if !ok {
		fmt.Println("asd")
	}
	hex := cryptoencoding.IntToByte(n)
	fmt.Println(hex)
}
