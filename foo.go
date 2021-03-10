package main

type Person struct{}

func (p Person) say() {
	// 获取 student 指针
}

type Student struct {
	Person
}

func foo() {
	s := Student{}
	s.say()
}

//func main() {
//	s := Student{}
//
//	s.say()
//}
