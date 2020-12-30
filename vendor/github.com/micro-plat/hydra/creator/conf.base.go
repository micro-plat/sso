package creator

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ISUB interface {
	Sub(name string, s ...interface{}) ISUB
}

type BaseBuilder map[string]interface{}

//Sub 子配置
func (b BaseBuilder) Sub(name string, s ...interface{}) ISUB {
	if len(s) == 0 {
		panic(fmt.Sprintf("配置：%s值不能为空", name))
	}
	tp := reflect.TypeOf(s[0])
	val := reflect.ValueOf(s[0])
	if tp.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	switch tp.Kind() {
	case reflect.String:
		b[name] = json.RawMessage([]byte(val.Interface().(string)))
	case reflect.Struct, reflect.Ptr, reflect.Map:
		b[name] = val.Interface()
	default:
		panic(fmt.Sprintf("配置：%s值类型不支持", name))
	}
	return b
}
func (b BaseBuilder) Map() map[string]interface{} {
	return b
}
func (b BaseBuilder) Load() {
}
