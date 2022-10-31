package main

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	inputFile  = "./testdata/test1.md"
	resultFile = "test1.md.html"
	expFile    = "./testdata/test1.md.html" //exp
)

func TestParseContent(t *testing.T) {
	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result := parseContent(input)

	expected, err := ioutil.ReadFile(expFile)
	if err != nil {
		t.Fatal(err)
	}

	if len(expected) != len(result) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("Result content does not match golden file")
	}
}

func TestRun(t *testing.T) {
	if err := run(inputFile); err != nil {
		t.Fatal(err)
	}
	result, err := ioutil.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := ioutil.ReadFile(expFile)
	if err != nil {
		t.Fatal(err)
	}

	if len(expected) != len(result) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("Result content does not match golden file")

	}

	os.Remove(resultFile)
}
