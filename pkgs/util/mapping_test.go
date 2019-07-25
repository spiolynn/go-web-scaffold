package util

import (
	"testing"
)

type People struct {
	Name string `json:"name_tile"`
	Age  int    `json:"age_size"`
}

func TestJson_to_Struct(t *testing.T) {

	jsonStr := `
    {
        "name_tile": "liuXX",
        "age_size": 12,
		"age_size1": 13
    }
    `
	var a_people People
	Json_to_Struct(jsonStr, &a_people)

	t.Logf("%v", a_people)

}

func TestStruct_to_Json(t *testing.T) {

	var a_people People
	a_people = People{"tony", 11}

	json_str, _ := Struct_to_Json(a_people)

	t.Logf("%s", json_str)
}

// 测试struct list to json
func TestStruct_to_Json2(t *testing.T) {
	var people []*People
	people = append(people, &People{"tony", 1})
	people = append(people, &People{"tony", 2})
	json_str, _ := Struct_to_Json(people)
	t.Logf("%s", json_str)
}

func TestJson_to_Map(t *testing.T) {

	jsonStr := `
    {
        "name": "Liu XX",
        "age": 18
    }
    `

	amap, err := Json_to_Map(jsonStr)
	if err == nil {
		t.Logf("%v", amap)
	}

}

func TestMap_to_Json(t *testing.T) {

	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "Liu xx"
	mapInstance["Age"] = 18
	mapInstance["Address"] = "广东 深圳"

	jsonStr, err := Map_to_Json(mapInstance)
	if err == nil {
		t.Logf("%s", jsonStr)
	}

}

func TestMap_to_Struct(t *testing.T) {

	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "liang637210"
	mapInstance["Age"] = 28

	var a_people People

	Map_to_Struct(mapInstance, &a_people)

	t.Logf("%v", a_people)
	t.Logf("%v", a_people.Name)
}

// 测试不完全的map
func TestMap_to_Struct2(t *testing.T) {
	mapInstance := make(map[string]interface{})

	var a_people People

	Map_to_Struct(mapInstance, &a_people)
	mapInstance["Age"] = 28
	mapInstance["Age1"] = 28

	t.Logf("%v", a_people)
	t.Logf("%v", a_people.Name)
	t.Logf("%v", a_people.Age)
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestStruct_to_Map(t *testing.T) {

	user := User{5, "zhangsan", "password"}
	data := Struct_to_Map(user)
	t.Logf("%v", data)
}
