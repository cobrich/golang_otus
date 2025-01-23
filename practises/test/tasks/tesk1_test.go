package tasks_test

import (
	"test/tasks"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTitleCaseEmpty(t *testing.T){
	const str, minor, want = "", "", ""
	got := tasks.TitleCase(str, minor)
	require.Equal(t, want, got)
}

func TestTitleCaseWithoutMinor(t *testing.T){
	const str, minor, want = "My name is Bekzat", "", "MY NAME IS BEKZAT"
	got := tasks.TitleCase(str, minor)
	require.Equal(t, want, got)
}

func TestTitleCaseWithMinorInFirst(t *testing.T){
	const str, minor, want = "My name is Bekzat", "my bekzat", "MY NAME IS bekzat"
	got := tasks.TitleCase(str, minor)
	require.Equal(t, want, got)
}


type Cases struct{
	desc string
	str string
	minor string
	want string
}

func TestTableTitleCaseAllCases(t *testing.T){
	cases := []Cases{
		{
			desc: "case1",
			str: "",
			minor: "",
			want: "",
		},
		{
			desc: "case2",
			str: "My name is Bekzat",
			minor: "",
			want: "MY NAME IS BEKZAT",
		},
		{
			desc: "case3",
			str: "My name is Bekzat",
			minor: "my bekzat",
			want: "MY NAME IS bekzat",
		},		
	}

	for i := range cases{
		tC := cases[i]

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			got := tasks.TitleCase(tC.str, tC.minor)
			require.Equal(t, tC.want, got)
		})
	}
}