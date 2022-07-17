package test

import (
	"fmt"
	"github.com/axinger/ax-go-web/src/main/go/model"
	"testing"
)

// test 必须以Test开头
func TestStudent(t *testing.T) {
	student := model.Student{
		Age:  10,
		Name: "jim",
	}
	fmt.Println("student = ", student)

	m := new(model.Student)
	fmt.Println("new对象 返回地址值,有&符号 = ", m)
	m.Name = "tom"
	fmt.Println("new对象 2 = ", m)
	// new对象 返回地址值,
	changeName(*m)
	fmt.Println("new对象 2 = ", m)
	// new对象 返回地址值,
	changeName2(m)
	fmt.Println("new对象 3 = ", m)

	m1 := model.Student{}
	m1.Name = "jim"
	fmt.Println("m1 = ", m1)
	changeName(m1)
	fmt.Println("m1 changeName = ", m1)

	m2 := model.Student{}
	m2.Name = "jim"
	fmt.Println("m2 = ", m2)
	changeName2(&m2)
	fmt.Println("m2 changeName2 = ", m2)

	//初始化对象,返回地址值
	m3 := &model.Student{}
	m3.Name = "jim"
	changeName(*m3)
	fmt.Println("m3 = ", m3)
	changeName2(m3)
	fmt.Println("m3 changeName2 = ", m3)
}

// 值传值
func changeName(s model.Student) {
	s.Name = s.Name + "1"
}

// 指针传值
func changeName2(s *model.Student) {
	s.Name = s.Name + "2"
}

// 初始化对象
func newStudent() model.Student {
	s := model.Student{}
	s.Name = "jim"
	return s
}

func newStudent2() *model.Student {
	s := &model.Student{}
	s.Name = "tom"
	return s
}
func newStudent3() model.Student {
	s := new(model.Student)
	s.Name = "jack"
	return *s
}

func newStudent4() *model.Student {
	s := new(model.Student)
	s.Name = "lili"
	return s
}

func Test_newStudent(t *testing.T) {

	student := newStudent()
	fmt.Println("student = ", student)
	student.Name = student.Name + "1"
	fmt.Println("student = ", student)

	student2 := newStudent2()
	fmt.Println("student2 = ", student2)
	student2.Name = student2.Name + "1"
	fmt.Println("student2 = ", student2)

	student3 := newStudent3()
	fmt.Println("student3 = ", student3)
	student3.Name = student3.Name + "1"
	fmt.Println("student3 = ", student3)

	student4 := newStudent4()
	fmt.Println("student4 = ", student4)
	student4.Name = student4.Name + "1"
	fmt.Println("student4 = ", student4)
}

func Test_student_dog(t *testing.T) {

	student := model.Student{}
	fmt.Println("st = ", student)

	student.Dog.Name = "dog1"
	fmt.Println("st = ", student)
}
