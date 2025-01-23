package greeter_test

import (
	"test/greeter"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGreet(t *testing.T){
	res := greeter.Greet("Bekzatt")

	const expected  = "Hello Bekzat"

	require.Equal(t, expected, res)
}