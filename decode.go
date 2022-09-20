package gmapstruct

import (
	"bytes"
	"encoding/json"
	"errors"
	"gtags"
	"reflect"
	"strconv"

	"github.com/gogf/gf/util/gconv"
)

func tryUnmarshal(stags *gtags.Field, val reflect.Value, data any) (bool, error) {
	field := val.FieldByIndex(stags.Index())
	if ok := gtags.TypMethod(field.Type(), "UnmarshalJSON"); ok {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode(data)
		///
		err := json.Unmarshal(buf.Bytes(), field.Addr().Interface())
		if err != nil {
			return true, err
		} else {
			return true, nil
		}
	}
	return false, nil
}
func decodeStruct(stags *gtags.Field, val reflect.Value, amap any) error {
	if ok, err := tryUnmarshal(stags, val, amap); ok {
		return err
	}
	///
	dmap := amap.(map[string]any)
	////
	for _, f := range stags.Fields() {
		err := decodeField(f, val, dmap[f.Alias()])
		if err != nil {
			return err
		}
	}
	//
	for _, anon := range stags.AnonNesteds() {
		err := decodeField(anon, val, dmap)
		if err != nil {
			return err
		}
	}
	///
	for _, nested := range stags.Nesteds() {
		err := decodeField(nested, val, dmap[nested.Alias()])
		if err != nil {
			return err
		}
	}

	return nil
}

func decodeField(field *gtags.Field, val reflect.Value, v any) error {
	if v == nil {
		return nil
	}
	if ok, err := tryUnmarshal(field, val, v); ok {
		return err
	}
	switch field.Type().Kind() {
	case reflect.Int:
		if _, ok := v.(string); ok {
			i, err := strconv.ParseInt(v.(string), 10, 0)
			if err != nil {
				return err
			} else {
				val.FieldByIndex(field.Index()).SetInt(i)
			}
		}
	case reflect.Int64:
		if _, ok := v.(string); ok {
			i, err := strconv.ParseInt(v.(string), 10, 0)
			if err != nil {
				return err
			} else {
				val.FieldByIndex(field.Index()).SetInt(i)
			}
		}
	case reflect.Uint, reflect.Uint64:
		if _, ok := v.(string); ok {
			i, err := strconv.ParseUint(v.(string), 10, 0)
			if err != nil {
				return err
			} else {
				val.FieldByIndex(field.Index()).SetUint(i)
			}
		}
	case reflect.Float32, reflect.Float64:
		i := gconv.Float64(v)
		val.FieldByIndex(field.Index()).SetFloat(i)
	case reflect.Bool:
		i := gconv.Bool(v)
		val.FieldByIndex(field.Index()).SetBool(i)
	case reflect.String:
		if _, ok := v.(string); ok {
			val.FieldByIndex(field.Index()).SetString(v.(string))
		} else {
			panic("is not string")
		}
	case reflect.Array:
		_, err := tryUnmarshal(field, val, v)
		return err
	case reflect.Slice:
		if field.Type().Elem().Kind() == reflect.String {
			s := []string{}
			for _, i := range v.([]string) {
				s = append(s, i)
			}
			val.FieldByIndex(field.Index()).Set(reflect.ValueOf(s))
		} else if field.Type().Elem().Kind() == reflect.Int {
			s := []int{}
			for _, i := range v.([]string) {
				s = append(s, gconv.Int(i))
			}
			val.FieldByIndex(field.Index()).Set(reflect.ValueOf(s))
		} else if field.Type().Elem().Kind() == reflect.Struct {
			return decodeStruct(field, val, v.(map[string]any))
		} else {
			return errors.New("notsupply slice " + field.Type().Elem().Name())
		}
	case reflect.Struct:
		return decodeStruct(field, val, v)
	default:
	}
	return nil
}
