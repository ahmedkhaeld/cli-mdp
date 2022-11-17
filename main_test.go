package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	inputFile = "./testdata/test1.md"
	expFile   = "./testdata/test1.md.html" //exp
)

func TestParseContent(t *testing.T) {
	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result, err := parseContent(input, "")
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
}

//update TestRun to take filename dynamically

func TestRun(t *testing.T) {
	//mockStdout will hold the file name
	//generated from calling the run()
	var mockStdout bytes.Buffer
	if err := run(inputFile, "", &mockStdout, true); err != nil {
		t.Fatal(err)
	}
	//remove any space type character from the file name
	resultFile := strings.TrimSpace(mockStdout.String())

	got, err := ioutil.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := ioutil.ReadFile(expFile)
	if err != nil {
		t.Fatal(err)
	}

	if len(expected) != len(got) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", got)
		t.Error("Result content does not match golden file")

	}

	os.Remove(resultFile)
}
