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

func decodeToStruct(stags *gtags.Field, obj any, dmap map[string]any) error {
	//todo: need ptr
	val := indirectVal(obj)
	field := val.FieldByIndex(stags.Index())
	//
	if ok := gtags.TypMethod(field.Type(), "UnmarshalJSON"); ok {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode(dmap)
		///
		err := json.Unmarshal(buf.Bytes(), field.Addr().Interface())
		return err
		//todo: error
	}
	///
	fields := []*gtags.Field{}
	fields = append(fields, stags.Fields()...)
	fields = append(fields, stags.AnonFields()...)
	for _, sf := range fields {
		idx := sf.Index()
		field := val.FieldByIndex(idx)
		alias := sf.Alias()
		v := dmap[alias]
		///
		ok := gtags.TypMethod(field.Type(), "UnmarshalJSON")
		if ok {
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			enc.Encode(v)
			///

			err := json.Unmarshal(buf.Bytes(), field.Addr().Interface())
			if err != nil {
				return err
			}

		} else {
			cval, _ := convKind(field.Kind(), v)
			//todo: cannotset
			field.Set(reflect.ValueOf(cval))
		}
	}
	///
	for _, nested := range stags.Nesteds() {
		// ss := stags.NestedByName(sname)
		if nested.IsAnon() {
			decodeToStruct(nested, obj, dmap)
		} else {
			if v, ok := dmap[nested.Alias()]; ok {
				if ok := tryDecode(obj, nested, v); !ok {
					smap := v.(map[string]any)
					decodeToStruct(nested, obj, smap)
				}
			}
		}
	}

	///
	return nil
}

func tryDecode(obj any, stags *gtags.Field, dmap any) bool {

	val := indirectVal(obj)
	field := val.FieldByIndex(stags.Index())
	if ok := gtags.TypMethod(field.Type(), "UnmarshalJSON"); ok {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode(dmap)
		///
		err := json.Unmarshal(buf.Bytes(), field.Addr().Interface())
		//todo: err
		if err != nil {
			panic(err)

		}
		return true
	}
	return false
}
