package cryptoencoding

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

//  Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.

func Hexb64(sHex string) string {
	// decode the hex:
	fmt.Println(sHex)
	dHex, err := hex.DecodeString(sHex)
	if err != nil {
		log.Fatal(err)
	}
	//encode the string:
	encoded := base64.StdEncoding.EncodeToString([]byte(dHex))

	return encoded
}
