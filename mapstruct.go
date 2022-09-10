package gmapstruct

import (
	"bytes"
	"context"
	"encoding/json"
	"gtags"
	"reflect"

	"github.com/gogf/gf/util/gvalid"
)

func Map2Struct(req any, dmap map[string]any) error {

	stags := gtags.ParseStructTags(req)

	dmap = TidyMapDefVal(stags, dmap)
	gerr := gvalid.CheckStructWithData(context.Background(), req, dmap, nil)
	if gerr != nil {
		return gerr
	}
	return decodeToStruct(stags, req, dmap)
}

func decodeToStruct(stags *gtags.Structs, obj any, dmap map[string]any) error {
	//todo: need ptr
	val := indirectVal(obj)
	// typ := indirectType(obj)
	// val := reflect.ValueOf(obj)
	///
	for _, fname := range stags.FieldNames() {
		// sf := stags.FieldByName(fname)
		sf := stags.Field(fname)
		idx := sf.Index()
		field := val.FieldByIndex(idx)
		field.Type().Name()
		alias := sf.Alias()
		v := dmap[alias]
		///
		nval := reflect.New(field.Type())
		if method := nval.MethodByName("UnmarshalJSON"); method.IsValid() {
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			enc.Encode(v)
			///

			err := json.Unmarshal(buf.Bytes(), field.Addr().Interface())
			if err != nil {
				//todo: error
			}

		} else {
			cval, _ := convKind(field.Kind(), v)
			if field.CanSet() {
				field.Set(reflect.ValueOf(cval))
			} else {
				//todo: cannotset
			}
		}
	}
	///
	for _, sname := range stags.NestedNames() {
		ss := stags.NestedByName(sname)
		if ss.IsAnon() {
			decodeToStruct(ss, obj, dmap)
		} else {
			smap := dmap[ss.Alise()].(map[string]any)
			decodeToStruct(ss, obj, smap)
		}
	}

	///
	return nil
}
