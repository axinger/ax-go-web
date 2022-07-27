package test

import (
	"fmt"
	"testing"
)

// 指定类型的泛型
type Num interface {
	// ~int 包含int 和MyInt
	~int | int32 | int64
}

type MyInt int

func add[T any](a T, b T) string {

	return fmt.Sprintf("%v%v", a, b)
}

//func add2[T any](a T) {
//	fmt.Println("========= %v \n", a)
//
//	fmt.Println("查看值==========", a, "直接逗号拼接")
//	fmt.Printf("查看值v %v \n", a)
//}

// map
type TMap[K string | int, V string | int] map[K]V

func Test_test1(t *testing.T) {

	sun := add[int](1, 2)
	fmt.Println(sun)

	//add[string]("foo", "bar")
	//
	////add(1,2)
	////add("foo","bar")
	//
	//add2[string]("a")

	ma1 := map[string]any{
		"name": "jim",
		"age":  10,
	}
	fmt.Println(ma1)

	ma2 := TMap[string, int]{}

	fmt.Println(ma2)

	s := make([]int, 10)
	fmt.Println(s)
}
