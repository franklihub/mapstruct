package gmapstruct

import (
	"bytes"
	"encoding/json"
	"gtags"
	"reflect"
	"strconv"

	"github.com/gogf/gf/util/gconv"
)

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

func conv(field *gtags.Field, v any) (any, error) {
	if field.HasUnmarshal() {
		b, _ := json.Marshal(v)
		val := reflect.New(field.Type())
		dec := json.NewDecoder(bytes.NewBuffer(b))
		err := dec.Decode(val)
		return val, err
	} else {
		return convKind(field.Type().Kind(), v)
	}

}

func convKind(kind reflect.Kind, v any) (any, error) {

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
	return v, nil
}
