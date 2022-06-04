package decode

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func getStr(data []byte) (string, int) {
	ret := ""
	for i, v := range data {
		if v == 0 {
			return ret, i + 1
		} else {
			ret += string(data[i])
		}
	}
	return ret, 0
}

func TransDomain(s, domain string) []byte {
	s, _, _ = strings.Cut(s, domain)
	s = strings.Replace(s, ".", "", -1)
	fmt.Println(s)
	hex, _ := hex.DecodeString(s)
	return hex
}
