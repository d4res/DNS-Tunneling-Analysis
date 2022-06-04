package decode

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

type execPacket struct {
	name string
	cmd  string
}

type shellPacket struct {
	name string
}

type downPacket struct {
	filename string
}

type uploadPacket struct {
	filename string
	data     []byte
}
