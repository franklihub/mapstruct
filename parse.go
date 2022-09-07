package gmapstruct

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"

	"github.com/franklihub/gtags"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

func conv(kind reflect.Kind, v any) any {

	switch kind {
	case reflect.Int, reflect.Int64:
		if _, ok := v.(string); ok {
			i, err := strconv.ParseInt(v.(string), 10, 0)
			if err != nil {
			} else {
				v = i
			}
		}
	case reflect.Uint, reflect.Uint64:
		if _, ok := v.(string); ok {
			i, err := strconv.ParseUint(v.(string), 10, 0)
			if err != nil {
			} else {
				v = i
			}
		}
	case reflect.Float32:
		v = gconv.Float32(v)
	case reflect.Float64:
		v = gconv.Float64(v)
	case reflect.Bool:
		v = gconv.Bool(v)
	case reflect.String:
	default:
	}
	return v
}

func Map2Struct(req any, m map[string]any) error {

	stags := gtags.ParseStructTags(req)

	m = TidyMapDefVal(stags, m)
	// m = defValTagToName(stags, m)
	gerr := gvalid.CheckStructWithData(context.Background(), req, m, nil)
	if gerr != nil {
		return gerr
	}
	///
	bm, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return decodeByteToStruct(req, bm)
	///
	// return decodeMapToStruct(req, m)
}

///
func decodeByteToStruct(obj any, bm []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(bm))
	// n := reflect.New(field.Type())
	err := dec.Decode(obj)
	return err
	///
}

////

func indirectType(obj any) reflect.Type {
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ
}
func indirectVal(obj any) reflect.Value {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
func decodeMapToStruct(obj any, m map[string]any) error {
	stag := gtags.ParseStructTags(obj)
	//need ptr
	val := indirectVal(obj)
	typ := indirectType(obj)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldtyp := typ.Field(i)
		if fieldtyp.Anonymous == false && fieldtyp.Type.Kind() == reflect.Struct {
			return errors.New("not supply nested struct")
		}
		if fieldtyp.Type.Kind() == reflect.Struct {
			err := decodeMapToStruct(field.Addr().Interface(), m)
			if err != nil {
				return err
			}
		}
		name := fieldtyp.Name
		tagname := stag.FieldByName(name).Tags().Get("p").Val()
		if v, ok := m[tagname]; ok {
			if v == "" {
				continue
			}
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)

			err := enc.Encode(v)
			if err != nil {
				return err
			}

			dec := json.NewDecoder(&buf)
			n := reflect.New(field.Type())
			err = dec.Decode(n.Interface())
			if err != nil {
				return err
			}
			field.Set(n.Elem())
		}
	}
	return nil
	///
}
func defValTagToName(stags *gtags.Structs, m map[string]any) map[string]any {
	return m
}
func mergeDefVal(stags *gtags.Structs, m map[string]any) map[string]any {
	names := stags.FieldNames()

	for _, n := range names {
		name := stags.FieldByName(n).Tags().Get("p").Val()
		if _, ok := m[name]; !ok {
			val := stags.FieldByName(n).Tags().Get("d").Val()
			m[name] = val
			m[n] = val
		}
	}
	///
	return m
}
