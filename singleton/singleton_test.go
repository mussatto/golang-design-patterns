package singleton

import (
	"testing"
	"fmt"
)

func TestGetMyInstance(t *testing.T) {
	instance := GetMyInstance()
	fmt.Println(instance.SomeVariable)
	instance.SomeVariable = "changing value"
	fmt.Println(instance.SomeVariable)

	instance2 := GetMyInstance()

	fmt.Println(instance2.SomeVariable)

	if(instance.SomeVariable != instance2.SomeVariable){
		t.Error("Should be the same value")
	}
}
