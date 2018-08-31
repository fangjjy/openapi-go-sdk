package util

import (
	"errors"
	"fmt"
	"reflect"
)

//StructToForm  格式化obj 为form表单数据
func StructToForm(obj interface{}) (map[string]interface{}, error) {
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	k := val.Kind()
	if k != reflect.Struct {
		return nil, errors.New("obj type only can struct")
	}
	fmt.Println(k)
	form := valuestructToForm(val)
	return form, nil
}

func valuestructToForm(val reflect.Value) map[string]interface{} {
	form := make(map[string]interface{}, 0)
	typ := val.Type()
	fileds := typ.NumField()
	for i := 0; i < fileds; i++ {
		v := val.Field(i)
		fieldname := typ.Field(i).Tag.Get("json")
		if fieldname == "" {
			fieldname = typ.Field(i).Name
		}
		switch v.Type().Kind() {
		case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.String,
			reflect.Float32, reflect.Float64:
			//form = append(form, fmt.Sprintf("%s=%v", fieldname, v))
			form[fieldname] = v
		case reflect.Struct:
			tmp := valuestructToForm(v)
			for k, v := range tmp {
				form[k] = v
			}
		case reflect.Slice:
			if v.Type().Elem().String() == "dto.Image" && v.Interface() != nil && v.Len() > 0 {
				t, _ := JsonEncode(v.Interface())
				form[fieldname] = t
			}
		}
	}
	return form
}
