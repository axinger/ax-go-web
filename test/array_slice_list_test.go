package test

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_list_1(t *testing.T) {
	// 数组长度固定,一般常用切片
	/// 初始值长度 ...
	var list1 = [...]string{"1", "2"}
	fmt.Println(list1)
	// 不能再添加内容
}
func Test_list_2(t *testing.T) {
	fmt.Println("🐛🐛🐛🐛🐛 list遍历")
	list2 := [5]int{0: 1, 4: 5}
	fmt.Println(list2)
	for i := range list2 {
		fmt.Println("value = ", i)
	}

	for i, i2 := range list2 {
		fmt.Println("index = ", i, "value = ", i2)
	}
}
func Test_list_3(t *testing.T) {
	fmt.Println("======切片 slice========")
	// 切片 slice
	list3 := []int{1, 2}
	fmt.Println(list3)
	_ = append(list3, 3)

	fmt.Println(list3)

	/// 直接追加 多个,必须接收返回值,才能改变原切片内容
	list3 = append(list3, 3, 4)

	fmt.Println(list3)
}
func Test_list_6(t *testing.T) {

	// 切片不存值,底层数组存值
	/// list6 底层数组为 arr6
	arr6 := []int{1, 3, 5, 7, 9}
	fmt.Println("arr6 = ", arr6)

	arr6 = append(arr6, 4)

}
func Test_list_7(t *testing.T) {
	// Go语言内置容器list是一个双向链表(实际是一个环)。位于包list当中。
	l := list.New()
	l.PushBack(1)
	fmt.Printf("%T \n", l)

	fmt.Println("list = ", l)
	front := l.Front()
	fmt.Printf("Front = %T %v\n", front, front.Value)

}
