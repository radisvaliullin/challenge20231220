package mysqlproto

import (
	"fmt"
	"strings"
)

func BuildHandShakeLogInfo(handshake *HandShakePayloadV10) string {
	builder := strings.Builder{}
	builder.WriteString("ProtocolVersion: ")
	builder.WriteString(fmt.Sprintf("%v", handshake.ProtocolVer))
	builder.WriteString("\n")
	builder.WriteString("ServerVersion: ")
	builder.WriteString(fmt.Sprintf("%v", handshake.ServerVersion))
	builder.WriteString("\n")
	builder.WriteString("ThreadId: ")
	builder.WriteString(fmt.Sprintf("%v", handshake.ThreadId))
	builder.WriteString("\n")
	builder.WriteString("Filler: ")
	builder.WriteString(fmt.Sprintf("%v", handshake.Filler))
	builder.WriteString("\n")
	builder.WriteString("CapabilityFlags1: ")
	builder.WriteString(fmt.Sprintf("%016b", handshake.CapabilityFlags1))
	builder.WriteString("\n")
	builder.WriteString("CapabilityFlags2: ")
	builder.WriteString(fmt.Sprintf("%016b", handshake.CapabilityFlags2))
	builder.WriteString("\n")
	builder.WriteString("CharacterSet: ")
	builder.WriteString(fmt.Sprintf("%X", handshake.CharacterSet))
	builder.WriteString("\n")
	builder.WriteString("StatusFlags: ")
	builder.WriteString(fmt.Sprintf("%016b", handshake.StatusFlags))
	builder.WriteString("\n")

	if IsCapabFlag2ClientPluginAuth(handshake.CapabilityFlags2) {
		builder.WriteString("AuthPluginDataLen: ")
		builder.WriteString(fmt.Sprintf("%v", handshake.AuthPluginDataLenORConstant))
		builder.WriteString("\n")
		builder.WriteString("AuthPluginData: ")
		authPluginDataBytes := append(handshake.AuthPluginDataPart1[:], handshake.AuthPluginDataPart2...)
		builder.WriteString(fmt.Sprintf("[% X]", authPluginDataBytes))
		builder.WriteString("\n")
		builder.WriteString("AuthPluginName: ")
		builder.WriteString(fmt.Sprintf("%v", handshake.AuthPluginName))
		builder.WriteString("\n")
	} else {
		builder.WriteString("00: ")
		builder.WriteString(fmt.Sprintf("%v", handshake.AuthPluginDataLenORConstant))
		builder.WriteString("\n")
	}
	return builder.String()
}
