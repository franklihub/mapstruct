package gmapstruct

import (
	"gtags"
	"reflect"
	"strings"
)

func TidyMapDefVal(stags *gtags.Structs, dmap map[string]any) map[string]any {
	return tidyStructDefVal(stags, dmap)
}

var DefValTag = "d"

func tidyStructDefVal(stags *gtags.Structs, dmap map[string]any) map[string]any {
	////
	//scan field, include anonstruct
	for _, f := range stags.FieldNames() {
		// field := stags.FieldByName(f)
		field := stags.Field(f)
		if v, ok := dmap[field.Alias()]; !ok {
			dmap[field.Alias()] = field.Tags().Get(DefValTag).Val()
		} else {
			if _, ok := v.(string); ok {
				dmap[field.Alias()] = strings.ToLower(v.(string))
			}
		}
	}
	///scan nested
	for _, f := range stags.NestedNames() {
		nested := stags.NestedByName(f)
		//todo: nested name
		structname := nested.Alise()
		////
		if v, ok := dmap[structname]; ok {
			d := tidyStructDefVal(nested, v.(map[string]any))
			dmap[structname] = d
		} else {
			d := map[string]any{}
			d = tidyStructDefVal(nested, d)
			dmap[structname] = d
		}
		//
	}
	return dmap
}

func tidyFieldDefVal(field *gtags.Field, dmap map[string]any) map[string]any {
	fieldalias := field.Alias()

	v, ok := dmap[fieldalias]
	if ok {
		if reflect.ValueOf(v).IsZero() {
			delete(dmap, fieldalias)
			ok = false
		} else {
			v, _ = convKind(field.Type().Kind(), v)
			dmap[fieldalias] = v
		}
	}
	if !ok {
		v = field.Tags().Get(DefValTag).Val()
		if v != "" {
			v, _ = convKind(field.Type().Kind(), v)
			dmap[fieldalias] = v
		}
	}

	return dmap
}
