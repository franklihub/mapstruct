package gmapstruct

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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
	assert.Equal(t, val.Log.Level[0], "panic")
	cfgmap["heart_beat"] = "off"
	val2 := &Config{}
	err = Map2Struct(val2, cfgmap)
	assert.Equal(t, err, nil)
	assert.Equal(t, val2.Heartbeat, HeartBeat(-1))

	////
}

func Test_SliceString(t *testing.T) {
	type slice struct {
		Str []string `json:"str" d:"a,b,c"`
	}
	dmap := map[string]any{}

	v := &slice{}
	err := Map2Struct(v, dmap)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(v.Str), 3)
	assert.Equal(t, v.Str[0], "a")
	//
	dmap = map[string]any{
		"str": []string{"aa", "bb"},
	}
	err = Map2Struct(v, dmap)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(v.Str), 2)
	assert.Equal(t, v.Str[0], "aa")
}
func Test_SliceInt(t *testing.T) {
	type slice struct {
		Data []int `json:"str" d:"1,2,3"`
	}
	dmap := map[string]any{}

	v := &slice{}
	err := Map2Struct(v, dmap)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(v.Data), 3)
	assert.Equal(t, v.Data[0], 1)
	//
	dmap = map[string]any{
		"str": []string{"11", "22"},
	}
	err = Map2Struct(v, dmap)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(v.Data), 2)
	assert.Equal(t, v.Data[0], 11)
}

func Test_Int(t *testing.T) {
	type slice struct {
		Data int `json:"data" d:"200"`
	}
	dmap := map[string]any{}

	v := &slice{}
	err := Map2Struct(v, dmap)
	assert.Equal(t, err, nil)
	assert.Equal(t, v.Data, 200)
	//
	dmap = map[string]any{
		"data": "200",
	}
	err = Map2Struct(v, dmap)
	assert.Equal(t, err, nil)
	assert.Equal(t, v.Data, 200)
	//
	dmap = map[string]any{
		"data": "300",
	}
	err = Map2Struct(v, dmap)
	assert.Equal(t, v.Data, 300)
}

func Test_SliceDev(t *testing.T) {
	type slice struct {
		Data []string `json:"data" d:"a,b,c,d"`
	}
	dmap := map[string]any{}

	v := &slice{}
	err := Map2Struct(v, dmap)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(v.Data), 4)
	//
	dmap = map[string]any{
		"data": []string{"e"},
	}
	err = Map2Struct(v, dmap)
	assert.Equal(t, err, nil)
	assert.Equal(t, v.Data[0], "e")
}

func Test_Address(t *testing.T) {
	type slice struct {
		Data common.Address `json:"data" d:"0x0000000000000000000000000000000000000000"`
	}
	dmap := map[string]any{}

	v := &slice{}
	err := Map2Struct(v, dmap)
	assert.Equal(t, err, nil)
	assert.Equal(t, v.Data, common.HexToAddress("0x0000000000000000000000000000000000000000"))
	//
	dmap = map[string]any{
		"data": "0x7aa7f0528d551096463d60380139844f6d4a6ac2",
	}
	err = Map2Struct(v, dmap)
	assert.Equal(t, err, nil)
	assert.Equal(t, v.Data, common.HexToAddress("0x7aa7f0528d551096463d60380139844f6d4a6ac2"))
	///
	//
	dmap = map[string]any{
		"data": "abcdedfd",
	}
	err = Map2Struct(v, dmap)
	assert.Equal(t, err == nil, false)
	es := strings.Split(err.Error(), ":")
	assert.Equal(t, strings.TrimSpace(es[1]), "cannot unmarshal hex string without 0x prefix into Go value of type common.Address")
}
func Test_NonPointer(t *testing.T) {
	type nonPointer struct {
		Data string
	}
	dmap := map[string]any{}

	v := nonPointer{}
	err := Map2Struct(v, dmap)
	assert.Equal(t, err.Error(), "non-pointer gmapstruct.nonPointer")
}

func Test_Vtype(t *testing.T) {
	type conf struct {
		TimeOut int      `json:"time_out" d:"100"`
		Type    []string `json:"type" p:"type" d:"a,b,c" v:"in:a,b,c,d"`
	}

	v := conf{}
	dmap := map[string]any{}
	err := Map2Struct(&v, dmap)
	assert.Equal(t, err != nil, true)
}
