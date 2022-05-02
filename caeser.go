package main

import (
	"fmt"
	"strings"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

func encrypt(str string) string {
	outstr := ""

	for _, v := range str {
		pos := strings.Index(alphabet, fmt.Sprintf("%c", v))

		pos = (pos + 13) % 26

		outstr = fmt.Sprintf("%s%c", outstr, alphabet[pos])
	}

	return outstr
}

func main() {
	fmt.Println(encrypt("testout"))

	fmt.Println(encrypt("grfgbhg"))
}
