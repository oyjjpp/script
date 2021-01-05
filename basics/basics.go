package Basics

type User interface {
	Writer()
	Read()
}

type Student struct {
}

func (u Student) Writer() {}
func (u Student) Read()   {}

// 检查一个结构体是否实现了某个接口
var _ User = &Student{}
