package hook

import "fmt"

type MyHookInterface interface {
	TheHookMethod() string
}

type MyDefaultMethods struct {

}

func (defaultMethods MyDefaultMethods) TheDefaultMethod(myHook MyHookInterface){
	fmt.Println("I am writing default stuff in the console")

	fmt.Println("then i am calling the hook")

	fmt.Println(myHook.TheHookMethod())

}

type MyHookA struct {

}

type MyHookB struct {

}

func (self *MyHookA) TheHookMethod() string{
	return "This is the HOOK A"
}

func (self *MyHookB) TheHookMethod() string{
	return "This is the HOOK B"
}