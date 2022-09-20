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
	dmap = TidyMapDefVal(stags, dmap)
	assert.Equal(t, dmap["name"], "cfgname")
	assert.Equal(t, dmap["heart_beat"], "off")
	assert.Equal(t, dmap["log"].(map[string]any)["stdout"], "true")
	assert.Equal(t, dmap["log"].(map[string]any)["level"].([]string)[0], "error")
	assert.Equal(t, dmap["addr"], "url")
	assert.Equal(t, dmap["port"], "777")
}
