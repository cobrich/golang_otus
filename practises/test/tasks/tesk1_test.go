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