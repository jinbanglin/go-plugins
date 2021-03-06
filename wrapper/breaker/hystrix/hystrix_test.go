package hystrix

import (
	"testing"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/jinbanglin/go-micro/client"
	"github.com/jinbanglin/go-micro/registry/mock"
	"github.com/jinbanglin/go-micro/selector"

	"context"
)

func TestBreaker(t *testing.T) {
	// setup
	r := mock.NewRegistry()
	s := selector.NewSelector(selector.Registry(r))

	c := client.NewClient(
		// set the selector
		client.Selector(s),
		// add the breaker wrapper
		client.Wrap(NewClientWrapper()),
	)

	req := c.NewRequest("test.service", "Test.Method", map[string]string{
		"foo": "bar",
	}, client.WithContentType("application/json"))

	var rsp map[string]interface{}

	// Force to point of trip
	for i := 0; i < (hystrix.DefaultVolumeThreshold * 3); i++ {
		c.Call(context.TODO(), req, rsp)
	}

	err := c.Call(context.TODO(), req, rsp)
	if err == nil {
		t.Error("Expecting tripped breaker, got nil error")
	}

	if err.Error() != "hystrix: circuit open" {
		t.Errorf("Expecting tripped breaker, got %v", err)
	}
}
