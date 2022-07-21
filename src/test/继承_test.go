package test

import (
	"fmt"
	"testing"
)

type People struct {
	name string
	age  int
}

func (people People) speak() {
	fmt.Println("people = ", people) //main.People
}

type Student struct {
	//嵌入匿名的成员 继承
	People
	grade int
}

// 重写父类方法
func (student Student) speak() {
	fmt.Println("name = ", student.name, "age = ", student.age, "grade = ", student.grade) //main.People
}

/**
通过嵌入匿名的成员，不仅可以继承匿名成员的内部成员，还可以继承匿名成员的方法。
需要注意的是，所有继承来的方法的接收者参数依然是那个匿名成员本身，而不是当前变量。

在Go语言中，this就是实现该方法的类型的对象。
*/
func Test_继承1(t *testing.T) {
	// 继承直接赋值
	var stu1 Student
	stu1.name = "jim"
	stu1.age = 18
	stu1.grade = 2
	stu1.speak() //编译期变成stu1.People.speak()

	//继承来的成员变量不能这样赋值
	var stu2 = Student{
		//name: "tom",
		//age : 18,
		grade: 22,
	}
	stu2.speak()

	//继承来的成员变量,构造函数赋值
	var stu3 = Student{
		People: People{
			name: "jack",
			age:  10,
		},
		grade: 22,
	}
	stu3.speak()
	//子类调用父类方法
	stu3.People.speak()
}
