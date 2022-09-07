package gmapstruct

import "reflect"

func scanStructVal(obj any, gtags *int, valmap map[string]any) {
	//todo: elem
	val := reflect.ValueOf(obj)
	//foreach fields
	// for gtags {
	for true {
		field := val.FieldByName("fieldname")
		//todo: if UnmarshalJson else set
		v := "val"
		field.Set(reflect.ValueOf(v))
		//todo: is struct recursive scanstruct
		//todo: map val
		vm := valmap["fieldname"]
		if v, ok := vm.(map[string]interface{}); ok {
			scanStructVal(field.Interface(), gtags, v)
		} else {
			panic("have no field valmap")
		}
	}

}
