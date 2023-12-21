package mysqlproto

const (
	PacketHeaderSize           = 4
	PacketHeaderFieldLenSize   = 3
	PacketHeaderFieldSeqIdSize = 1
)

// Frame Packet is protocol top level frame packet encapsulated other packets as payload
// see detail: https://dev.mysql.com/doc/dev/mysql-server/latest/PAGE_PROTOCOL.html
type FramePacket struct {
	FramePacketHeader
	Payload []byte
}

type FramePacketHeader struct {
	Len        uint32
	SequenceId uint8
}

type HandShakePayloadV10 struct {
	ProtocolVer uint8
	// string<null>, bytes of string terminated by 0x00
	ServerVersion string
	HandShakePayloadV10Part1
	// Rest of the plugin provided data, $len=MAX(13, length of auth-plugin-data - 8)
	AuthPluginDataPart2 []byte
	// string<null>, if capabilities & CLIENT_PLUGIN_AUTH
	AuthPluginName string
}

type HandShakePayloadV10Part1 struct {
	ThreadId uint32
	// string<8>, first part
	AuthPluginDataPart1 [8]byte
	// 0x00 byte, terminating the first part
	Filler uint8
	// The lower 2 bytes of the Capabilities Flags
	CapabilityFlags1 uint16
	// default server a_protocol_character_set, only the lower 8-bits
	CharacterSet uint8
	// SERVER_STATUS_flags_enum
	StatusFlags uint16
	// The upper 2 bytes of the Capabilities Flags
	CapabilityFlags2 uint16
	// if capabilities & CLIENT_PLUGIN_AUTH
	// then AuthPluginDataLen, length of the combined auth_plugin_data, if auth_plugin_data_len is > 0
	// else constant 0x00
	AuthPluginDataLenORConstant uint8
	// string<10>
	Reserved [10]byte
}
