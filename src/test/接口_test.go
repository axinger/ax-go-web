package test

import (
	"fmt"
	"testing"
)

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

type animal interface {
	eat() string
}
type person struct {
	name string
}

// 调用接口的方法,就表示实现了接口
func (person *person) eat() string {
	return fmt.Sprintf("person = %v", person)
}

// 花
type flowers struct {
	name string
}

// 参数是接口
func isAnimal(animal animal) {

	fmt.Println("animal = ", animal)
}

func Test_接口1(t *testing.T) {

	p := person{}
	p.name = "jim"
	eat := p.eat()
	fmt.Println("eat = ", eat)

	isAnimal(&p)

	f := flowers{}
	fmt.Println("f = ", f)

	// 可以看出,flowers 不符合参数是接口
	//isAnimal(&f)
}
