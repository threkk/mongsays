package main

import "testing"
import "fmt"

func TestSplitInLinesToReturn1LineIncomplete(t *testing.T) {
	res := splitInLines("one line", 80, 2)

	if len(res) != 1 {
		t.Error("Expected 1, got ", len(res))
		t.Log("Input: \"one line\". Output: ", res)
	}

	for _, word := range res {
		if len(word) > 80 {
			t.Error(fmt.Sprintf("Expected %s to be < 80, got %d", word, len(word)))
		}
	}
}

func TestSplitInLinesToReturn1LineLongerThanWidth(t *testing.T) {
	res := splitInLines("abcdefghijklmno", 5, 2)

	if len(res) != 5 {
		t.Error("Expected 5, got ", len(res))
		t.Log("Input: \"abcdefghijklmno\". Output: ", res)
	}

	for _, word := range res {
		if len(word) > (5 - 2) {
			t.Error(fmt.Sprintf("Expected %s to be < 3, got %d", word, len(word)))
		}
	}
}

func TestSplitInLinesToReturn2Lines(t *testing.T) {
	res := splitInLines("two lines", 7, 2)

	if len(res) != 2 {
		t.Error("Expected 2, got ", len(res))
		t.Log("Input: \"two lines\". Output: ", res)
	}

	for _, word := range res {
		if len(word) > (7 - 2) {
			t.Error(fmt.Sprintf("Expected %s to be < 4, got %d", word, len(word)))
		}
	}
}

func TestSplitInLinesToReturnComplexCase(t *testing.T) {
	res := splitInLines("one line and thesecondweirdpart", 12, 2)

	if len(res) != 4 {
		t.Error("Expected 4, got ", len(res))
		t.Log("Input \"one line and thesecondweirdpart\". Output:", res)
	}

	for _, word := range res {
		if len(word) > (12 - 2) {
			t.Error(fmt.Sprintf("Expected %s to be < 12, got %d", word, len(word)))
		}
	}
}

func TestSplitInLinesToReturnSeveralLines(t *testing.T) {
	lorem := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse vitae erat id dolor pharetra maximus. Donec eget est euismod, sagittis urna at, interdum leo. Etiam lacinia felis in euismod malesuada."
	res := splitInLines(lorem, 40, 4)

	if len(res) != 6 {
		t.Error("Expected 6, got ", len(res))
		t.Log("Output:", res)
	}

	for _, word := range res {
		if len(word) > (40 - 4) {
			t.Error(fmt.Sprintf("Expected %s to be < 36, got %d", word, len(word)))
		}
	}
}
