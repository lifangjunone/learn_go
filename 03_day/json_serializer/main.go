package main

import (
	"encoding/json"
	"fmt"
)

// Go-json　序列化

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

/*
结构体序列化
*/

func StructMonsterSerializer() string {
	monster := Monster{
		Name:     "牛逼",
		Age:      200,
		Birthday: "2006-01-02",
		Sal:      800000.0,
		Skill:    "牛逼逸轩",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误：　%v\n", err)
		return ""
	}
	fmt.Printf("结果：%v\n", string(data))
	return string(data)
}

func StructMonsterUnSerializer() {
	str := StructMonsterSerializer()
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("反序列化%v\n", monster)
}

/*
Map 序列化
*/

func MapSerializer() string {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "牛田"
	a["age"] = 100
	a["address"] = "那你呢"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("map 序列化错误：　%v\n", err)
		return ""
	}
	fmt.Printf("map 序列化结果: %v \n", string(data))
	return string(data)
}

func MapUnSerializer() {
	str := MapSerializer()
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", a)
}

/*
切片序列化
*/

func SliceSerializer() string {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "牛逼"
	m1["age"] = "33"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "牛逼2"
	m2["age"] = "332"
	m2["address"] = "天津"
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("Slice 序列化err:%v", err)
		return ""
	}
	fmt.Printf("slice 序列化成功：　%v\n", string(data))
	return string(data)
}

func SliceUnSerializer() {
	str := SliceSerializer()
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", str)
}

/*
基本数据类型序列化
*/

func BaseDataTypeSerializer() string {
	var num1 float64 = 245.33
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(data))
	return string(data)
}

func BaseDataTypeUnSerializer() {
	str := BaseDataTypeSerializer()
	var num float64
	err := json.Unmarshal([]byte(str), &num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}

func main() {
	//StructMonsterSerializer()
	//MapSerializer()
	//SliceSerializer()
	//BaseDataTypeSerializer()
	//StructMonsterUnSerializer()
	//MapUnSerializer()
	//SliceUnSerializer()
	BaseDataTypeUnSerializer()
}
