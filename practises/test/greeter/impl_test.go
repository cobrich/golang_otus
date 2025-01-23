package greeter_test

import (
	"test/greeter"
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestGreet(t *testing.T){
// 	res := greeter.Greet("Bekzat")

// 	const expected  = "Hello Bekzat"

// 	require.Equal(t, expected, res)
// }

func TestTableGreet(t *testing.T){
	cases := []struct{
		desc string
		input string
		expect string
	}{
		{
			desc: "case1",
			input: "",
			expect: "Hello world",
		},
		{
			desc: "case2",
			input: "Bekzat",
			expect: "Hello Bekzat",
		},
	}

	for i := range cases{
		tC := cases[i]

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			res := greeter.Greet(tC.input)
	
			require.Equal(t, tC.expect, res)
		})
	}
}