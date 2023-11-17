package dir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	_, err := NewDirTree("..\\dir")
	assert.NoError(t, err)

	_, err = NewDirTree("..\\..\\foobar")
	assert.Equal(t, err.Error()[:35], "The specified file does not exist: ")

	_, err = NewDirTree("..\\dir\\dir.go") // this file exists but the given path must be a directiry
	assert.EqualError(t, err, "Given path is not an existing directory")
}

// TODO Implémenter ce test
/*
func ExampleDir() {
	fmt.Print(NewDirTree("..\\tree"))
	// Output:
	// digraph {
}
*/
