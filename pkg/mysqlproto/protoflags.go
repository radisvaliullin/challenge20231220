package mysqlproto

const (
	// Capability Flag 1
	CapabFlag1ClientProto41 = 512
	// Capability Flag 2
	CapabFlag2ClientPluginAuth = uint16(1) << 3
)

func IsCapabFlag1ClientProto41(capabFlag1 uint16) bool {
	return (capabFlag1 & CapabFlag1ClientProto41) != 0
}

func IsCapabFlag2ClientPluginAuth(capabFlag2 uint16) bool {
	return (capabFlag2 & CapabFlag2ClientPluginAuth) != 0
}
