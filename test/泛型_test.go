package test

import (
	"fmt"
	"testing"
)

// 指定类型的泛型
type Num interface {
	int | int32 | int64
}

func add[T any](a T, b T) string {

	return fmt.Sprintf("%v%v", a, b)
}

//func add2[T any](a T) {
//	fmt.Println("========= %v \n", a)
//
//	fmt.Println("查看值==========", a, "直接逗号拼接")
//	fmt.Printf("查看值v %v \n", a)
//}

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

}
