package authenticated

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type Params struct {
	Values url.Values
}

func NewParams() *Params {
	return &Params{Values: url.Values{}}
}
func (p *Params) Add(k string, v *string) *Params {
	if v != nil {
		p.Values.Add(k, *v)
	}
	return p
}

func (p *Params) AddInt64(k string, v *int64) *Params {
	if v != nil {
		p.Values.Add(k, strconv.FormatInt(*v, 10))
	}
	return p

}

func (p *Params) AddUint64(k string, v *uint64) *Params {
	if v != nil {
		p.Values.Add(k, strconv.FormatUint(*v, 10))
	}
	return p

}
func (p *Params) AddUint32(k string, v *uint32) *Params {
	if v != nil {
		p.Values.Add(k, strconv.FormatUint(uint64(*v), 10))
	}
	return p

}
func (p *Params) AddBoolean(k string, v *bool) *Params {
	if v != nil {
		p.Values.Add(k, strconv.FormatBool(*v))
	}
	return p
}
func (p *Params) String() string {
	return p.Values.Encode()
}

func Struct2Map(s interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("只支持结构体")
	}

	t := v.Type()
	m := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// 跳过未导出的字段
		if !field.IsExported() {
			continue
		}

		// 检查字段是否为零值，如果是则忽略
		if isZeroValue(value) {
			continue
		}

		// 获取json标签
		tag := field.Tag.Get("json")
		if tag == "" {
			tag = field.Name
		} else if tag == "-" {
			continue // 跳过标记为"-"的字段
		}

		// 如果tag包含逗号，如"name,omitempty"，只取第一部分
		if idx := strings.Index(tag, ","); idx != -1 {
			tag = tag[:idx]
		}

		m[tag] = value.Interface()
	}

	return m, nil
}

// 检查值是否为零值
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.String:
		return v.String() == ""
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Slice, reflect.Map, reflect.Chan:
		return v.IsNil() || v.Len() == 0
	case reflect.Struct:
		// 对于结构体，我们可以认为只有当所有字段都是零值时它才是零值
		// 或者你可以实现自定义逻辑
		return false // 这里简化处理，认为结构体永远不是零值
	}
	return false
}
