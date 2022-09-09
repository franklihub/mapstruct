package gmapstruct

import (
	"testing"

	"gotest.tools/assert"
)

func Test_ParserMap(t *testing.T) {
	val := &Config{}
	err := Map2Struct(val, cfgmap)
	assert.Equal(t, err, nil)
	assert.Equal(t, val.Name, "cfgname")
	assert.Equal(t, val.Addr, "localhost")
	assert.Equal(t, val.Port, 8888)
	assert.Equal(t, val.Heartbeat, HeartBeat(10))
	assert.Equal(t, val.Log.Stdout, false)
	// assert.Equal(t, val.Log.Level, []{"panic","error"})
	cfgmap["heart_beat"] = "off"
	val2 := &Config{}
	err = Map2Struct(val2, cfgmap)
	assert.Equal(t, err, nil)
	assert.Equal(t, val.Heartbeat, HeartBeat(-1))

	////
}
