package frame

import (
	"github.com/lonng/nano/internal/message"
)

const (
	Request  = message.Request
	Notify   = message.Notify
	Response = message.Response
	Push     = message.Push
)

type Type = message.Type
type Message = message.Message

type PacketCodec interface {
	// Encode
	Encode(*Message) ([]byte, error)
	// Decode
	Decode([]byte) ([]*Message, error)
}
