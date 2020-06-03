package pojo

type Talker interface {
	SayHello(word, role string) (response string)
}
