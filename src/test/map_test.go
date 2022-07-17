package test

import (
	"fmt"
	"testing"
)

func Test_map_test1(t *testing.T) {
	m1 := make(map[string]any)
	m1["name"] = "tom"
	m1["age"] = "30"
	fmt.Printf("m1 type:%T value %v \n", m1, m1)
	fmt.Println("make map = ", m1)

	// 不要用new 创建map
	//m2 := new(map[string]string)
	//fmt.Printf("m2 use new create type:%T value %v \n", m2, m2)
	//////直接赋值会报 panic: assignment to entry in nil map
	//m2 = &m1
	//(*m2)["name"] = "fan"
	//////对m2的修改也会影响到m1
	////fmt.Println("new map = ", m1)

	map1 := map[any]any{
		"name": "jim",
		"age":  10,
		10:     "ten",
	}
	fmt.Println(map1)
	fmt.Println("map取值 = ", map1["name"])
	fmt.Println("map取值 = ", map1["name1"])
	if map1["name1"] == nil {
		fmt.Println("map取值 nil")
	}
	fmt.Println(map1[10])
}
