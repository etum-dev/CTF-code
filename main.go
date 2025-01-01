package main

import (
	hex_to_b64 "etum-dev/challset_1/chall"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Hello, what can i do for you?")
	argPtr := flag.String("in", "sample", "input string")
	flag.Parse()
	//chall 1
	myhex := *argPtr
	decode := hex_to_b64.Hexb64(myhex)
	fmt.Println(decode)

	fmt.Println("Cryptohacks chall:")
	myAscii := [23]int{99, 114, 121, 112, 116, 111, 123, 65, 83, 67, 73, 73, 95, 112, 114, 49, 110, 116, 52, 98, 108, 51, 125}
	ascii := hex_to_b64.ArrToAscii(myAscii[:])
	fmt.Println(ascii)
}

// SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
