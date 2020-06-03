package pojo

import "fmt"

//这部分代码可以通过mockgen生成用于测试的Mock对象

type Person struct {
	name string
	role string
}

func NewPerson(name, role string) *Person {
	return &Person{
		name: name,
		role: role,
	}
}

// 实现了SayHello方法 == 继承了Talker接口
func (p *Person) SayHello(name, role string) (word string) {
	return fmt.Sprintf("Hello %s(身份：%s), welcome to GoLand IDE. My name is %s", name, role, p.name)
}
