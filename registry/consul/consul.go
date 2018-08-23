package consul

import (
	"github.com/jinbanglin/go-micro/registry"
)

func NewRegistry(opts ...registry.Option) registry.Registry {
	return registry.NewRegistry(opts...)
}
