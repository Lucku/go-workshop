package mocking

import (
	"io/ioutil"
	"os"
	"testing"
)

// FileReader takes a file name relative to the current directory as an input and returns the file's contents as a
// string
type FileReader interface {
	readFile(fileName string) string
}

// RealFileReader takes a relative filename as an input, reads the file using the standard library's os package and
// returns the read content as a string
type RealFileReader struct {
}

// Real implementation of readFile
func (r RealFileReader) readFile(fileName string) string {

	f, _ := os.Open(fileName)

	fileBytes, _ := ioutil.ReadAll(f)

	return string(fileBytes)
}

// MockFileReader is a mock version of FileReader, containing the string content returned by readFile as a fixed
// attribute, so that it can be used for testing
type MockFileReader struct {
	content string
}

func (m MockFileReader) readFile(fileName string) string {
	return m.content
}

// DoSomething reads a file and does something with the contents
func DoSomething(reader FileReader) {

	// reading file
	reader.readFile("text.txt")

	// do something with the file contents
	// ...
}

func TestDoSomething(t *testing.T) {

	testContent := "klasndklansdöklnaskl"

	mockFileReader := MockFileReader{content: testContent}

	DoSomething(mockFileReader)

	// assert(...)
}
