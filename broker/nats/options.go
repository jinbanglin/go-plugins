package nats

import (
	"context"

	"github.com/jinbanglin/go-micro/broker"
	"github.com/nats-io/go-nats"
)

type optionsKey struct{}

// Options accepts nats.Options
func Options(opts nats.Options) broker.Option {
	return func(o *broker.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, optionsKey{}, opts)
	}
}
