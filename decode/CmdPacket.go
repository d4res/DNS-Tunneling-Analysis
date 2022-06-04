package decode

import (
	"encoding/binary"
	"encoding/hex"
)

const (
	COMMAND_PING     = 0x0000
	COMMAND_SHELL    = 0x0001
	COMMAND_EXEC     = 0x0002
	COMMAND_DOWNLOAD = 0x0003
	COMMAND_UPLOAD   = 0x0004
	COMMAND_SHUTDOWN = 0x0005
	COMMAND_DELAY    = 0x0006
	COMMAND_ERROR    = 0xFFFF
)

type ExecPacket struct {
	Name string
	Cmd  string
}

type ShellPacket struct {
	Name string
}

type DownPacket struct {
	Filename string
}

type UploadPacket struct {
	Filename string
	Data     string
}

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
		return ExecPacket{name, cmd}, COMMAND_EXEC
	// server->client only
	case COMMAND_SHELL:
		name, _ := getStr(data)
		return ShellPacket{name}, COMMAND_SHELL
	// server->client only
	case COMMAND_UPLOAD:
		filename, i := getStr(data)
		data = data[i:]
		return UploadPacket{filename, hex.EncodeToString(data)}, COMMAND_UPLOAD
	// server->client only
	case COMMAND_DOWNLOAD:
		filename, _ := getStr(data)
		return DownPacket{filename}, COMMAND_DOWNLOAD
	default:
		return nil, COMMAND_ERROR
	}
}
