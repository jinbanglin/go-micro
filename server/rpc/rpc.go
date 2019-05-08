package rpc

import (
	"github.com/jinbanglin/go-micro/server"
)

func NewServer(opts ...server.Option) server.Server {
	return server.NewServer(opts...)
}
