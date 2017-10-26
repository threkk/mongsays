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
	res := splitInLines("abcdefghijklmno", 5, 2)

	if len(res) != 5 {
		t.Error("Expected 5, got ", len(res))
		t.Log("Input: \"abcdefghijklmno\". Output: ", res)
	}
}

func TestSplitInLinesToReturn2Lines(t *testing.T) {
	res := splitInLines("two lines", 6, 2)

	if len(res) != 2 {
		t.Error("Expected 2, got ", len(res))
		t.Log("Input: \"two lines\". Output: ", res)
	}
}

func TestSplitInLinesToReturnComplexCase(t *testing.T) {
	res := splitInLines("one line and thesecondweirdpart", 12, 2)
	t.Log(res)
	if len(res) != 3 {
		t.Error("Expected 3, got ", len(res))
		t.Log("Input \"one line and thesecondweirdpart\". Output:", res)
	}
}
