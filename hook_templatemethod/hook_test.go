package hook

import (
	"testing"
	"fmt"
)

func TestHook(T *testing.T){
	fmt.Println("t=TestHook")
	myDefaultMethods := MyDefaultMethods{}

	hookA := &MyHookA{}

	hookB := &MyHookB{}

	myDefaultMethods.TheDefaultMethod(hookA)

	myDefaultMethods.TheDefaultMethod(hookB)

	fmt.Println("END")
}
