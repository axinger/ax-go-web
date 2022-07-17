package test

import (
	"fmt"
	"testing"
)

// 字符串拼接
func Test_string(t *testing.T) {
	/**
	  	+拼接方式
	    这种方式是我在写golang经常用的方式，
	  	go语言用+拼接，php使用.拼接，不过由于golang中的字符串是不可变的类型，
	  	因此用 + 连接会产生一个新的字符串对效率有影响。
	*/
	s1 := "hello"
	s2 := "word"
	s3 := s1 + s2
	fmt.Println(s3) //s3 = "helloword"

	s4 := fmt.Sprintf("%s%s", "1", "2") //s3 = "helloword"
	fmt.Println("s4 = ", s4)

	s5 := fmt.Sprintf("%v%v", 1, 2) //s3 = "helloword"
	fmt.Println("s5 = ", s5)

	s6 := fmt.Sprintf("%v%v", "1", 2) //s3 = "helloword"
	fmt.Println("s6 = ", s6)
}
