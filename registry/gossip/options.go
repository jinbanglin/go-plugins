package gossip

import (
	"context"

	"github.com/jinbanglin/go-micro/registry"
)

type contextSecretKey struct{}

func SecretKey(k []byte) registry.Option {
	return func(o *registry.Options) {
		o.Context = context.WithValue(o.Context, contextSecretKey{}, k)
	}
}
