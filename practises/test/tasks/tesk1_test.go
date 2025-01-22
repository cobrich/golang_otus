package tasks_test

import (
	"test/tasks"
	"testing"
)

func TestTitleCaseEmpty(t *testing.T){
	const str, minor, want = "", "", ""
	got := tasks.TitleCase(str, minor)

	if got != want {
		t.Errorf("TitleCase(%v, %v) = %v, want %v", str, minor, got, want)
	}
}

func TestTitleCaseWithoutMinor(t *testing.T){
	const str, minor, want = "My name is Bekzat", "", "MY NAME IS BEKZAT"
	got := tasks.TitleCase(str, minor)

	if got != want {
		t.Errorf("TitleCase(%v, %v) = %v, want %v", str, minor, got, want)
	}
}

func TestTitleCaseWithMinorInFirst(t *testing.T){
	const str, minor, want = "My name is Bekzat", "my bekzat", "MY NAME IS bekzat"
	got := tasks.TitleCase(str, minor)

	if got != want {
		t.Errorf("TitleCase(%v, %v) = %v, want %v", str, minor, got, want)
	}
}