package mysqlproto

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

func DecodeFramePacket(rdr io.Reader) (FramePacket, error) {

	out := FramePacket{}

	// read header bytes
	headerBytes := make([]byte, PacketHeaderSize)
	_, err := io.ReadFull(rdr, headerBytes)
	if err != nil {
		log.Fatalf("decode frame packet: read header error: %v", err)
		return out, err
	}

	// decode header
	// decode len field
	// extend to 4 bytes (little endian)
	headerFieldLenBytes := append(headerBytes[:PacketHeaderFieldLenSize], 0)
	lenReader := bytes.NewReader(headerFieldLenBytes)
	err = binary.Read(lenReader, binary.LittleEndian, &out.Len)
	if err != nil {
		log.Fatalf("decode frame packet: decode header field len error: %v", err)
		return out, err
	}
	// decode seq id, take last item
	out.SequenceId = headerBytes[PacketHeaderFieldLenSize]

	// read payload
	payloadBytes := make([]byte, out.Len)
	_, err = io.ReadFull(rdr, payloadBytes)
	if err != nil {
		log.Fatalf("decode frame packet: read payload error: %v", err)
		return out, err
	}
	out.Payload = payloadBytes

	return out, nil
}

func DecodeHandShakePayloadV10(payload []byte) (HandShakePayloadV10, error) {

	out := HandShakePayloadV10{}
	var err error

	//
	buff := bytes.NewBuffer(payload)

	// read protocol version
	out.ProtocolVer, err = buff.ReadByte()
	if err != nil {
		log.Fatalf("decode handshake payload: read protocol version error: %v", err)
		return out, err
	}
	// read server version
	srvVerBytes, err := buff.ReadBytes(0)
	if err != nil {
		log.Fatalf("decode handshake payload: read server version error: %v", err)
		return out, err
	}
	out.ServerVersion = string(srvVerBytes)

	// decode part 1
	err = binary.Read(buff, binary.LittleEndian, &out.HandShakePayloadV10Part1)
	if err != nil {
		log.Fatalf("decode handshake payload: read part 1 error: %v", err)
		return out, err
	}

	// decode auth plugin data tail
	// check client plugin auth flag, if false return
	if !IsCapabFlag2ClientPluginAuth(out.CapabilityFlags2) {
		return out, nil
	}
	// read auth plugin data
	authPluginTailLen := out.AuthPluginDataLenORConstant - 8
	out.AuthPluginDataPart2 = make([]byte, authPluginTailLen)
	_, err = buff.Read(out.AuthPluginDataPart2)
	if err != nil {
		log.Fatalf("decode handshake payload: read part 2 error: %v", err)
		return out, err
	}
	// read auth plugin name
	authPluginBytes, err := buff.ReadBytes(0)
	if err != nil {
		log.Fatalf("decode handshake payload: read auth plugin name error: %v", err)
		return out, err
	}
	out.AuthPluginName = string(authPluginBytes)

	return out, nil
}
