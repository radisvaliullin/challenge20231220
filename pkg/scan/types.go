package scan

import "github.com/radisvaliullin/challenge20231220/pkg/mysqlproto"

type Result struct {
	Ok          bool
	ErrorStatus error
	Info        mysqlproto.HandShakePayloadV10
}
