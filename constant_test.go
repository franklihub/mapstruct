package gmapstruct

import (
	"fmt"
	"testing"
)

type HeartBeat int

func (a *HeartBeat) UnmarshalJson(data []byte) error {
	return nil
}

type Log struct {
	Stdout bool `json:"stdout" d:"true"`
	//todo: slice
	Level string `json:"level" d:"error,info,debug"`
}

type URL struct {
	Addr string `json:"addr" v:"required"`
	Port int    `json:"port" d:"123" v:"required"`
}
type Db struct {
	URL
	Pool int `json:"pool" d:"10"`
}
type Config struct {
	Name      string `json:"name" v:"required"`
	Db        `json:"db"`
	Log       Log       `json:"log" v:"required"`
	Heartbeat HeartBeat `json:"heart_beat" d:"off"`
}

var cfgmap map[string]any = map[string]any{
	"name": "cfgname",
	"log": map[string]any{
		"stdout": false,
		"level":  "panic,error",
	},
	"db": map[string]any{
		"addr": "localhost",
		"port": 8888,
		"pool": 10,
	},
	"heart_beat": "10",
}

func Test_Constant(t *testing.T) {
	fmt.Println(cfgmap)
}
