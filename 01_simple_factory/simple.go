package simplefactory

import "fmt"

//接口
type API interface {
	Say(name string) string
}

//根据type返回不同的结构体
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

//hiAPI
type hiAPI struct{}


func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

//HelloAPI
type helloAPI struct{}


func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
