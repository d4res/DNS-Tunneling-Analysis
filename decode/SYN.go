package decode

import (
	"encoding/binary"
	"encoding/hex"
)

const (
	OPT_NAME    = 0x0001
	OPT_COMMAND = 0x0020
)

type SYNBody struct {
	Seq     string
	Options string
	Name    string
}

func parseSYN(data []byte) SYNBody {
	seq := data[0:2]     // 2 bytes
	options := data[2:4] // 2bytes
	data = data[4:]
	name := ""
	if binary.BigEndian.Uint16(options)&OPT_NAME == OPT_NAME {
		for i, v := range data {
			if v == byte(0) {
				break
			} else {
				name = name + string(data[i])
			}
		}
	} else {
		name = "[unnamed]"
	}

	return SYNBody{hex.EncodeToString(seq), hex.EncodeToString(options), name}
}
