package greeter_test

import (
	"test/greeter"
	"testing"
)

func TestGreet(t *testing.T){
	res := greeter.Greet("Bekzat")

	if res != "Hello Bekzat"{
		t.Errorf("not equal")
	}

}