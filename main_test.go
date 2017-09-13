package main

import "testing"

func TestSplitInLinesToReturn1LineIncomplete(t *testing.T) {
	res := splitInLines("one line", 80, 2)

	if len(res) != 1 {
		t.Error("Expected 1, got ", len(res))
		t.Log("Input: \"one line\". Output: ", res)
	}
}

func TestSplitInLinesToReturn1LineLongerThanWidth(t *testing.T) {
	res := splitInLines("abcdefghijklmno", 3, 2)

	if len(res) != 4 {
		t.Error("Expected 4, got ", len(res))
		t.Log("Input: \"abcdefghijklmno\". Output: ", res)
	}
}

func TestSplitInLinesToReturn2Lines(t *testing.T) {
	res := splitInLines("two lines", 3, 2)

	if len(res) != 2 {
		t.Error("Expected 2, got ", len(res))
		t.Log("Input: \"two lines\". Output: ", res)
	}
}
