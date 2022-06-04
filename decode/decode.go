package decode

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func DecodePlain(data string, domain string) error {
	data, _, _ = strings.Cut(data, "."+domain)

	bytes, err := hex.DecodeString(data)
	if err != nil {
		return err
	}

	for _, v := range bytes {
		fmt.Println(v)
	}
	return nil
}
