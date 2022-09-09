package gmapstruct

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gogf/gf/util/gconv"
)

type HeartBeat int

func (a *HeartBeat) UnmarshalJSON(data []byte) error {
	d := strings.TrimSpace(string(data))
	d = strings.Trim(d, `"`)
	if d == "off" {
		*a = -1
	} else {
		*a = HeartBeat(gconv.Int(d))
	}
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
	"addr":       "localhost",
	"port":       8888,
	"pool":       10,
	"heart_beat": "10",
}

func Test_Constant(t *testing.T) {
	fmt.Println(cfgmap)
}
