package main

import "testing"

func TestCountWords(t *testing.T) {
	data := []struct{
		input string
		expectedCount int
	}{
		{"hello world", 2},
		{"The Go Programming Language", 4},
	}

	for _, d := range data {
		counter := new(WordCounter)
		_, err := counter.Write([]byte(d.input))
		if err != nil {
			t.Fatal(err)
		}
		if count := int(*counter); count != d.expectedCount {
			t.Errorf("expected output wasn't matched: got %d, want %d\n", count, d.expectedCount)
		}
	}

	t.Run("Count", func(t *testing.T){
		for _, d := range data {
			counter := new(WordCounter)
			err := count(counter, d.input)
			if err != nil {
				t.Fatal(err)
			}
			if count := int(*counter); count != d.expectedCount {
				t.Errorf("expected output wasn't matched: got %d, want %d\n", count, d.expectedCount)
			}
		}
	})
}

func TestCountLines(t *testing.T) {
	data := []struct{
		input string
		expectedCount int
	}{
		{"hello\nworld\n", 2},
		{"The\nGo\nProgramming\nLanguage\n", 4},
	}

	for _, d := range data {
		counter := new(LineCounter)
		_, err := counter.Write([]byte(d.input))
		if err != nil {
			t.Fatal(err)
		}
		if count := int(*counter); count != d.expectedCount {
			t.Errorf("expected output wasn't matched: got %d, want %d\n", count, d.expectedCount)
		}
	}

	t.Run("Count", func(t *testing.T){
		for _, d := range data {
			counter := new(LineCounter)
			err := count(counter, d.input)
			if err != nil {
				t.Fatal(err)
			}
			if count := int(*counter); count != d.expectedCount {
				t.Errorf("expected output wasn't matched: got %d, want %d\n", count, d.expectedCount)
			}
		}
	})
}