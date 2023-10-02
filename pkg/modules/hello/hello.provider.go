package hello

import "fmt"

type IHelloProvider interface {
	SayHelloToName(string) string
	SayHello() string
}

type helloProvider struct{}

func HelloProvider() IHelloProvider {
	return &helloProvider{}
}

func (p *helloProvider) SayHelloToName(name string) string {

	return fmt.Sprintf("Hello %s", name)
}

func (p *helloProvider) SayHello() string {

	return "Hello world"
}
