package decode

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

const (
	OPT_NAME    = 0x0001
	OPT_COMMAND = 0x0020
)

type SYNBody struct {
	seq     []byte
	options []byte
	name    string
}

func (body SYNBody) String() string {
	return fmt.Sprintf("seq: %s, options: %s, name: %s", hex.EncodeToString(body.seq), hex.EncodeToString(body.options), body.name)
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

	return SYNBody{seq, options, name}
}
