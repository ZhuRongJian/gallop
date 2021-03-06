package config

import "fmt"

type Demo struct {
	Name string
}
type Configuration struct {
	Name string `value:"name"`
	A int `value:"a"`
	B *struct{
		C string `value:"c"`
	} `value:"b"`

}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

func (c *Configuration) NewDemo() *Demo {
	fmt.Println(c.B.C,"======")
	return &Demo{c.Name}
}

func (c *Configuration) NewDemo1() (string,*Demo) {
	return "demo",&Demo{Name: "larry"}
}

