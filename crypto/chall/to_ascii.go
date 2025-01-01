package cryptoencoding

// cryptohack chall
// Converts numbers to their ascii equiv.
// [99, 114, 121, 112, 116, 111, 123, 65, 83, 67, 73, 73, 95, 112, 114, 49, 110, 116, 52, 98, 108, 51, 125]

type Table struct {
	Number   uint16
	Alphabet string
}

var AsciiTable map[int]Table

func init() {
	AsciiTable = make(map[int]Table)
	for i := 0; i < 128; i++ {
		AsciiTable[i] = Table{
			Number: uint16(i),
			//rune = code point
			Alphabet: string(rune(i)),
		}
	}
}

func ArrToAscii(arr []int) string {
	finalString := ""
	for _, num := range arr {
		if entry, exists := AsciiTable[num]; exists {
			finalString += entry.Alphabet
		} else {
			finalString += "?"
		}
	}

	return finalString

}
