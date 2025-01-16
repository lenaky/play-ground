package module

import (
	"github.com/quic-go/quic-go"
)

//go:generate mockgen -package $GOPACKAGE -destination $SOURCE_ROOT/module/gen_mock_$GOFILE github.com/lenaky/play-ground/module Connection

type Connection interface {
	quic.Connection
}
