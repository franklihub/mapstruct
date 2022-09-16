package gmapstruct

import (
	"gtags"
	"strings"
)

var DefValTag = "d"

func TidyMapDefVal(stags *gtags.Field, dmap map[string]any) map[string]any {
	sdmap := stags.DMap(DefValTag)
	gtags.MergerMap(dmap, sdmap)
	// tidyStructDefVal(stags, dmap)
	// fmt.Println("dmap:", dmap)
	return dmap
}

func tidyStructDefVal(stags *gtags.Field, dmap map[string]any) map[string]any {
	////
	//scan field, include anonstruct
	fields := []*gtags.Field{}
	fields = append(fields, stags.Fields()...)
	// fields = append(fields, stags.AnonFields()...)

	for _, field := range fields {
		if v, ok := dmap[field.Alias()]; !ok {
			d := field.Tags().Get(DefValTag).Val()
			o := field.Tags().Get(DefValTag).Opts()
			if len(o) > 0 {
				d = d + "," + strings.Join(o, ",")
			}
			dmap[field.Alias()] = d
		} else {
			if _, ok := v.(string); ok {
				dmap[field.Alias()] = strings.ToLower(v.(string))
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
				dmap[field.Alias()] = d
			} else {
				if _, ok := v.(string); ok {
					dmap[field.Alias()] = strings.ToLower(v.(string))
				}
			}
		}
	}

	///scan nested
	for _, nested := range stags.Nesteds() {
		// nested := stags.NestedByName(f)
		//todo: nested name
		structname := nested.Alias()
		////
		//todo: UnmarshalJSON(string)
		if v, ok := dmap[structname]; ok {
			d := tidyStructDefVal(nested, v.(map[string]any))
			dmap[structname] = d
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
