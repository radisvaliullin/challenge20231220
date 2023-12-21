package scan

import (
	"fmt"
	"net"

	"github.com/radisvaliullin/challenge20231220/pkg/mysqlproto"
)

type Config struct {
	Host string
	Port int
}

type Scan struct {
	config Config
}

func New(config Config) *Scan {
	s := &Scan{
		config: config,
	}
	return s
}

func (s *Scan) Run() error {

	var result Result
	defer func() {
		fmt.Printf("%+v\n", result)
	}()

	addr := fmt.Sprintf("%s:%v", s.config.Host, s.config.Port)

	// connect to server
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		err := fmt.Errorf("scanner: dial to server error: %v", err)
		result = getErrorResult(err)
		return err
	}
	defer conn.Close()

	// accept handshake packet from server
	frame, err := mysqlproto.DecodeFramePacket(conn)
	if err != nil {
		err := fmt.Errorf("scanner: decode frame packet error: %v", err)
		result = getErrorResult(err)
		return err
	}

	handshakeV10, err := mysqlproto.DecodeHandShakePayloadV10(frame.Payload)
	if err != nil {
		err := fmt.Errorf("decode handshake payload error: %v", err)
		result = getErrorResult(err)
		return err
	}

	result = Result{
		Ok:   true,
		Info: handshakeV10,
	}
	return nil
}

func getErrorResult(err error) Result {
	return Result{
		Ok:          false,
		ErrorStatus: err,
	}
}
