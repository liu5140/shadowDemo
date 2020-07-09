package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/shopspring/decimal"
)

func Map2String(m map[string]interface{}) (result string) {
	list := make([]string, 0)
	for k, v := range m {
		t1 := fmt.Sprintf("%s=%s", k, fmt.Sprint(v))
		list = append(list, t1)
	}
	result = strings.Join(list, "&")
	return
}

func Map2Json(m map[string]interface{}) (result string) {
	b, _ := json.MarshalIndent(m, "", "    ")
	result = string(b)
	return
}

//string map排序
func SortMap(m map[string]string) (result string) {
	data := []string{}
	list := sort.StringSlice{}
	for k, _ := range m {
		list = append(list, k)
	}
	sort.Sort(list)
	for _, k := range list {
		if v, ok := m[k]; ok {
			data = append(data, k+"="+v)
		}
	}
	return strings.Join(data, "&")
}

//结构体转string Map
func StructToMap(data interface{}) (m map[string]string) {
	m = make(map[string]string)
	tt := reflect.TypeOf(data)
	vv := reflect.ValueOf(data)
	for i := 0; i < tt.NumField(); i++ {
		t := tt.Field(i).Type.String()
		key := tt.Field(i).Tag.Get("form")
		value := vv.Field(i)
		if t == "string" {
			m[key] = value.String()
		} else if t == "int64" {
			m[key] = fmt.Sprintf("%d", value.Int())
		} else if t == "decimal.Decimal" {
			m[key] = value.Interface().(decimal.Decimal).String()
		}
	}
	return
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
