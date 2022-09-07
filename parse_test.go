package gmapstruct

import (
	"strings"
	"testing"

	"github.com/gogf/gf/test/gtest"
	"github.com/gogf/gf/util/gconv"
	"gotest.tools/assert"
)

type Conn int

func (a *Conn) UnmarshalJSON(data []byte) error {
	d := strings.TrimSpace(string(data))
	d = strings.Trim(d, `"`)
	if d == "no" {
		*a = -1
	} else if d == "yes" {
		*a = 1
	} else {
		i := gconv.Int(data)
		*a = Conn(i)
	}

	return nil
}

type Conf struct {
	TimeOut int `json:"time_out" d:"100"`
	// Conn    rpc.BlockNumber `json:"conn" d:"latest"`
	Conn Conn     `json:"conn" d:"latest"`
	Type []string `json:"type" d:"a,b,c" v:"in:a,b,c,d"`
}
type ReqParam struct {
	Url  string `json:"url"`
	User string `json:"user" d:"admin"`
	Pass string `json:"pass"`
	Conf
}

func Test_AnonDefVal(t *testing.T) {
	data := map[string]any{}
	req := &ReqParam{}
	err := Map2Struct(req, data)
	assert.Equal(t, err, nil)
	// assert.Equal(t, "", tags.Get("p").Key())
}

func Test_ParserMap(t *testing.T) {
	///
	// gtags.AliaseTag = "json"
	//valparse
	gtest.C(t, func(t *gtest.T) {
		data := map[string]any{
			"url":      "http://test.com",
			"user":     "root",
			"pass":     "passwd",
			"time_out": "99",
		}
		req := &ReqParam{}
		err := Map2Struct(req, data)
		t.AssertEQ(err, nil)
		t.AssertEQ(req.Url, "http://test.com")
		t.AssertEQ(req.User, "root")
		t.AssertEQ(req.Pass, "passwd")
		t.AssertEQ(req.TimeOut, 99)
		////
	})
	//defval
	gtest.C(t, func(t *gtest.T) {
		data := map[string]any{
			"user": "assing",
		}
		req := &ReqParam{}
		err := Map2Struct(req, data)
		t.AssertEQ(err, nil)
		t.AssertEQ(req.User, "assing")
		t.AssertEQ(req.TimeOut, 100)
		////
	})
	//marshal
	gtest.C(t, func(t *gtest.T) {
		data := map[string]any{
			"conn": "latest",
		}
		req := &ReqParam{}
		err := Map2Struct(req, data)
		t.AssertEQ(err, nil)
		// t.AssertEQ(req.Conn, rpc.BlockNumber(-1))
		////
	})
	gtest.C(t, func(t *gtest.T) {
		data := map[string]any{
			"conn": "earliest",
		}
		req := &ReqParam{}
		err := Map2Struct(req, data)
		t.AssertEQ(err, nil)
		// t.AssertEQ(req.Conn, rpc.BlockNumber(0))
		////
	})
	gtest.C(t, func(t *gtest.T) {
		data := map[string]any{
			"conn": "0x99",
		}
		req := &ReqParam{}
		err := Map2Struct(req, data)
		t.AssertEQ(err, nil)
		// t.AssertEQ(req.Conn, rpc.BlockNumber(0x99))
		////
	})
}
