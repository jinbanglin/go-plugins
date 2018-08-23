package grpc

import (
	"github.com/jinbanglin/go-micro/errors"
	"github.com/jinbanglin/grpc-go/status"
)

func microError(err error) error {
	// no error
	switch err {
	case nil:
		return nil
	}

	// micro error
	if v, ok := err.(*errors.Error); ok {
		return v
	}

	// grpc error
	if s, ok := status.FromError(err); ok {
		return errors.Parse(s.Message())
	}

	// do nothing
	return err
}
