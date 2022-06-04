package decode

import (
	"encoding/binary"
)

// parsing the command packet(which is usually inside a MSG packet)
// (uint32_t) length (of the rest of the message)
// (uint16_t) packed_id
// (uint16_t) command_id
// (variable...) command fields
func ParseCmd(data []byte) (any, int) {
	// length := data[0:4]
	// packetId := data[4:6]
	cmdId := binary.BigEndian.Uint16((data[6:8]))
	data = data[8:]
	switch cmdId {
	// server->client only
	case COMMAND_EXEC:
		name, i := getStr(data)
		data = data[i:]
		cmd, _ := getStr(data)
		return execPacket{name, cmd}, COMMAND_EXEC
	// server->client only
	case COMMAND_SHELL:
		name, _ := getStr(data)
		return shellPacket{name}, COMMAND_SHELL
	// server->client only
	case COMMAND_UPLOAD:
		filename, i := getStr(data)
		data = data[i:]
		return uploadPacket{filename, data}, COMMAND_UPLOAD
	// server->client only
	case COMMAND_DOWNLOAD:
		filename, _ := getStr(data)
		return downPacket{filename}, COMMAND_DOWNLOAD
	default:
		return nil, COMMAND_ERROR
	}
}
