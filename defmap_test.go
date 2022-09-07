package gmapstruct

import (
	"testing"

	"gotest.tools/assert"
)

func Test_DefMap(t *testing.T) {
	data := map[string]any{}
	req := &ReqParam{}
	err := Map2Struct(req, data)
	assert.Equal(t, err, nil)
	// assert.Equal(t, "", tags.Get("p").Key())
}
