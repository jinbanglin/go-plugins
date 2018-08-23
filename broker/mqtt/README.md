# MQTT Broker

The MQTT broker is useful for IoT based applications

## Usage

Drop in import

```go
import _ "github.com/jinbanglin/go-plugins/broker/mqtt"
```

Flag on command line

```shell
go run main.go --broker=mqtt
```

Alternatively use directly

```go
import (
	"github.com/jinbanglin/go-micro"
	"github.com/jinbanglin/go-plugins/broker/mqtt"
)


func main() {
	service := micro.NewService(
		micro.Name("my.service"),
		micro.Broker(mqtt.NewBroker()),
	)
}
```

## Encoding

Because MQTT does not support message headers the plugin encodes messages using JSON. 
If you prefer to send and receive the mqtt payload uninterpreted use the `noop` codec.

Example

```go
import (
    "github.com/jinbanglin/broker"
    "github.com/jinbanglin/broker/codec/noop"
    "github.com/jinbanglin/go-plugins/broker/mqtt"
)

b := mqtt.NewBroker(
    broker.Codec(noop.NewCodec()),
)
```
