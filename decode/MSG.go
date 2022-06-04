package decode

const (
	MESSAGE_TYPE_SYN  = 0x00
	MESSAGE_TYPE_MSG  = 0x01
	MESSAGE_TYPE_FIN  = 0x02
	MESSAGE_TYPE_PING = 0xFF
	MESSAGE_TYPE_ENC  = 0x03
)

// (uint16_t) packet_id
// (uint8_t) message_type [0x01]
// (uint16_t) session_id
// (uint16_t) seq
// (uint16_t) ack
// (byte[]) data
func ParseMsg(data []byte) (msg Msg) {
	msg.PacketId = data[0:2]
	msg.MsgType = data[2]
	msg.SessId = data[3:5]
	data = data[5:]
	switch msg.MsgType {
	case MESSAGE_TYPE_MSG:
		_ = data[0:2]
		_ = data[2:4]
		data = data[4:]
		if len(data) != 0 {
			cmd, _ := ParseCmd(data)
			msg.Payload = cmd
		}
		return msg
	case MESSAGE_TYPE_SYN:
		syn := parseSYN(data)
		msg.Payload = syn
		return msg
	default:
		return Msg{}
	}

}

type Msg struct {
	PacketId []byte
	MsgType  byte
	SessId   []byte
	Payload  any
}
