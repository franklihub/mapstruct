package gmapstruct

import (
	"gtags"
	"testing"

	"gotest.tools/assert"
)

func Test_DefMap(t *testing.T) {
	dmap := map[string]any{
		"name": "cfgname",
		"log":  map[string]any{},
		"addr": "url",
		"port": "777",
	}

	req := &Config{}
	stags := gtags.ParseStructTags(req)
	// err := Map2Struct(req, data)
	dmap = TidyMapDefVal(stags, dmap)
	assert.Equal(t, dmap["name"], "cfgname")
	assert.Equal(t, dmap["heart_beat"], "off")
	assert.Equal(t, dmap["log"].(map[string]any)["stdout"], "true")
	assert.Equal(t, dmap["log"].(map[string]any)["level"], "error")
	assert.Equal(t, dmap["addr"], "url")
	//todo: conv.int
	assert.Equal(t, dmap["port"], "777")
}
