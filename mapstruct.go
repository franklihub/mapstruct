package gmapstruct

import (
	"context"
	"errors"
	"gtags"
	"reflect"

	"github.com/gogf/gf/util/gvalid"
)

func Map2Struct(req any, dmap map[string]any) error {
	val := reflect.ValueOf(req)
	if val.Kind() != reflect.Pointer || val.IsNil() {
		return errors.New("non-pointer " + val.Type().String())
	}
	stags := gtags.ParseStructTags(req)
	dmap = TidyMapDefVal(stags, dmap)
	//
	gerr := gvalid.CheckStructWithData(context.Background(), req, dmap, nil)
	if gerr != nil {
		return gerr
	}
	val = indirectVal(req)
	return decodeStruct(stags, val, dmap)
}
