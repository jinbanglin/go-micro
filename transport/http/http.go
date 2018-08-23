package http

import (
	"github.com/jinbanglin/go-micro/transport"
)

func NewTransport(opts ...transport.Option) transport.Transport {
	return transport.NewTransport(opts...)
}
