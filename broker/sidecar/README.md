# Sidecar Broker

This is a broker plugin for the micro [sidecar](https://github.com/jinbanglin/micro/tree/master/car)

## Usage

Here's a simple usage guide

### Run Sidecar

```
go get github.com/jinbanglin/micro
```

```
micro sidecar
```

### Import and Flag plugin

```
import _ "github.com/jinbanglin/go-plugins/broker/sidecar"
```

```
go run main.go --broker=sidecar
```
