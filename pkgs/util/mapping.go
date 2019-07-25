package util

import (
	"encoding/json"
	"github.com/goinggo/mapstructure"
	"reflect"
)

/*
封装数据结构的映射
*/

// json to struct
func Json_to_Struct(json_str string, struct_ interface{}) error {
	err := json.Unmarshal([]byte(json_str), struct_)
	return err
}

// struct to json
func Struct_to_Json(struct_ interface{}) (string, error) {

	jsonBytes, err := json.Marshal(struct_)
	if err != nil {
		return "", err
	} else {
		return string(jsonBytes), nil
	}

}

// Json to Map

func Json_to_Map(json_str string) (map[string]interface{}, error) {

	var mapResult map[string]interface{}

	if err := json.Unmarshal([]byte(json_str), &mapResult); err != nil {
		return nil, err
	} else {
		return mapResult, nil
	}
}

// Map to Json

func Map_to_Json(amap map[string]interface{}) (string, error) {

	jsonStr, err := json.Marshal(amap)
	if err != nil {
		return "", err
	} else {
		return string(jsonStr), nil
	}
}

// Map to Struct
func Map_to_Struct(amap map[string]interface{}, struct_ interface{}) error {

	if err := mapstructure.Decode(amap, struct_); err != nil {
		return err
	} else {
		return nil
	}
}

// Struct to Map
func Struct_to_Map(obj interface{}) map[string]interface{} {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}

	return data
}

// 888888888888888888888

// *[]struct to json

//// struct to json
//func Structs_to_Json(structs_ []interface{}) (string, error) {
//
//	for _,obj := range structs_{
//
//
//	}
//
//	jsonBytes, err := json.Marshal(struct_)
//	if err != nil{
//		return "",err
//	}else{
//		return string(jsonBytes), nil
//	}
//
//}
