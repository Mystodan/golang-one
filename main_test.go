package main_test

import (
	"fmt"
	assignment1 "golang1/pckg"
	"testing"
)

func TestTestValidity(t *testing.T) {
	var rTests = []struct {
		inn     string
		compare bool
	}{
		{"1-Hello-2-World", true},
		{"-Hello-2-World", false},
		{"", false},
	}
	for _, tString := range rTests {
		testN := fmt.Sprintf("%s", tString.inn)
		t.Run(testN, func(t *testing.T) {
			solution := assignment1.TestValidity(tString.inn)
			if solution != tString.compare {
				t.Fatal("comparing failed!")
			}
		})
	}
}
func TestTestAverage(t *testing.T) {
	var rTests = []struct {
		inn     string
		compare uint
	}{
		{"1-Hello-2-World", uint((1 + 2) / 2)},
		{"-Hello-2-World", uint(0)},
		{"", uint(0)},
	}
	for _, tString := range rTests {
		testN := fmt.Sprintf("%s", tString.inn)
		t.Run(testN, func(t *testing.T) {
			solution := assignment1.AverageNumber(tString.inn)
			if solution != tString.compare {
				t.Fatal("comparing failed!")
			}
		})
	}
}
func TestTestWholeStory(t *testing.T) {
	var rTests = []struct {
		inn     string
		compare string
	}{
		{"1-Hello-2-World", "Hello World"},
		{"-Hello-2-World", ""},
		{"", ""},
	}
	for _, tString := range rTests {
		testN := fmt.Sprintf("%s", tString.inn)
		t.Run(testN, func(t *testing.T) {
			solution := assignment1.WholeStory(tString.inn)
			if solution != tString.compare {
				t.Fatal("comparing failed!")
			}
		})
	}
}

func TestTestStoryStats(t *testing.T) {
	var rTests = []struct {
		inn string
		//compare
		shortest string
		longest  string
		average  float64
		list     []string
	}{
		{"1-Hello-2-World", "Hello", "Hello", 5.0, []string{"Hello", "World"}},
	}
	for _, tString := range rTests {
		testN := fmt.Sprintf("%s", tString.inn)
		t.Run(testN, func(t *testing.T) {
			shortest, longest, average, list := assignment1.StoryStats(tString.inn)
			if shortest != tString.shortest || longest != tString.longest || average != tString.average || !compare(list, tString.list) {
				t.Fatal("comparing failed!")
			}
		})
	}
}

func compare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
