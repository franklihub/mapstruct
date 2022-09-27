package gmapstruct

import (
	"gtags"
	"reflect"
	"strings"
)

var DefValTag = "d"

func TidyMapDefVal(stags *gtags.Field, dmap map[string]any) map[string]any {
	sdmap := stags.DMap(DefValTag)
	gtags.MergerMap(dmap, sdmap)
	dmap = tidyStructDefVal(stags, dmap)
	return dmap
}

func tidyStructDefVal(stags *gtags.Field, dmap map[string]any) map[string]any {
	////
	fields := []*gtags.Field{}
	fields = append(fields, stags.Fields()...)

	for _, field := range fields {
		if v, ok := dmap[field.Alias()]; !ok {
			d := field.Tags().Get(DefValTag).Val()
			o := field.Tags().Get(DefValTag).Opts()
			if len(o) > 0 {
				d = d + "," + strings.Join(o, ",")
			}
			dmap[field.Alias()] = tidyVal(field, d)
		} else {
			if _, ok := v.(string); ok {
				dmap[field.Alias()] = tidyVal(field, v)
			}
		}
	}
	///scan anonnested
	for _, field := range stags.AnonNesteds() {
		if field.IsStruct() {
			tidyStructDefVal(field, dmap)
		} else {
			if v, ok := dmap[field.Alias()]; !ok {
				d := field.Tags().Get(DefValTag).Val()
				o := field.Tags().Get(DefValTag).Opts()
				if len(o) > 0 {
					d = d + "," + strings.Join(o, ",")
				}
				dmap[field.Alias()] = tidyVal(field, d)
			} else {
				if _, ok := v.(string); ok {
					dmap[field.Alias()] = tidyVal(field, v)
				}
			}
		}
	}

	///scan nested
	for _, nested := range stags.Nesteds() {
		structname := nested.Alias()
		////
		if v, ok := dmap[structname]; ok {
			if nested.HasUnmarshal() {
			} else {
				d := tidyStructDefVal(nested, v.(map[string]any))
				dmap[structname] = d
			}
		} else {
			d := nested.Tags().Get(DefValTag).Val()
			if d != "" {
				dmap[structname] = d
			} else {
				d := map[string]any{}
				d = tidyStructDefVal(nested, d)
				dmap[structname] = d
			}
		}
		//
	}
	return dmap
}

func tidyVal(filed *gtags.Field, dval any) any {
	switch filed.Type().Kind() {
	case reflect.Slice:
		if s, ok := dval.(string); ok {
			s = strings.ToLower(s)
			slice := strings.Split(s, ",")
			if slice[0] != "" {
				for i, s := range slice {
					slice[i] = strings.TrimSpace(s)
				}
				return slice
			} else {
				return []string{}
			}
		} else {
			panic("tidyVal is not string:" + filed.Type().Name())
		}
	default:
		if s, ok := dval.(string); ok {
			return strings.ToLower(s)
		} else {
			panic("tidyVal is not string:" + filed.Type().Name())
		}
	}
}
