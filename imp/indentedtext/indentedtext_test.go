package indentedtext

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpfeiffer67/tree"
)

const (
	IndentedText1 = `0000
	0001
	 this is a
	 multiline
	 comment
		0002
	0003
`

	IndentedText2 = `0000
	AAAA
	 comments='''this is a
	 multiline
	 comment'''
	 style=1	
		BBBB
		 custom=1024
		 text="hello"	   
	CCCC
	 style=8
	DDDD
	 default multi-line
	 if not toml
`
)

const (
	test1 = `Node 0
	Node 1
		Node 11
		Node 12
			Node 121
			Node 122
			Node 123
	Node 2
Node x (error)
	Node 3
		Node 31
`
)

const (
	test2 = `	Node 0 (error)
	Node 1
		Node 11
		Node 12
			Node 121
			Node 122
			Node 123
	Node 2
	Node 3
		Node 31
`
)

const (
	test3 = `Node 0
	Node 1
			Node 11 (error)
		Node 12
			Node 121
			Node 122
			Node 123
	Node 2
	Node 3
		Node 31
`
)

func TestImport(t *testing.T) {
	assert := assert.New(t)
	n, _ := Import(IndentedText1)
	assert.Equal("0000", n.Values()["label"])
	n1 := n.Children()[0]
	n2 := n.Children()[0].Children()[0]
	n3 := n.Children()[1]
	assert.Equal("this is a\nmultiline\ncomment", n1.Values()["description"])
	assert.Equal("0002", n2.Values()["label"])
	assert.Equal("0003", n3.Values()["label"])
}

func TestImportToml(t *testing.T) {
	assert := assert.New(t)
	n, _ := Import(IndentedText2)

	n1 := n.Children()[0]
	n2 := n.Children()[0].Children()[0]
	n3 := n.Children()[1]
	n4 := n.Children()[2]

	assert.Equal("AAAA", n1.Values()["label"])
	assert.Equal("this is a\nmultiline\ncomment", n1.Values()["comments"])
	assert.Equal(int64(1), n1.Values()["style"])

	assert.Equal("BBBB", n2.Values()["label"])
	assert.Equal(int64(1024), n2.Values()["custom"])
	assert.Equal("hello", n2.Values()["text"])

	assert.Equal("CCCC", n3.Values()["label"])
	v, ok := n3.Values()["style"].(int64)
	assert.Equal(true, ok)
	assert.Equal(8, int(v))

	assert.Equal("DDDD", n4.Values()["label"])
	assert.Equal("default multi-line\nif not toml", n4.Values()["description"])
}

func TestImportErrors(t *testing.T) {
	assert := assert.New(t)

	_, err := Import(test1)
	assert.EqualError(err, "Import: Tree can have only one root")

	_, err = Import(test2)
	assert.EqualError(err, "Import: First line must be the root node")

	_, err = Import(test3)
	assert.EqualError(err, "Import: Deepness can only increment by one")
}

func TestImportFromFile(t *testing.T) {
	assert := assert.New(t)

	root, err := ImportFromFile("doesnotexist.txt")
	assert.Nil(root)
	assert.EqualError(err, "Import from file: Read string from file: open doesnotexist.txt: The system cannot find the file specified.")
	_, err = ImportFromFile("test/sample1.txt")
	assert.Nil(err)
}

func concatNodesLabelAndDescription(nodes []*tree.Node) string {
	s := ""
	for _, n := range nodes {
		s = s + n.ValueString("label", "") + "\n" + n.ValueString("description", "")
	}
	return s
}

func ExampleImportFromFile() {
	root, _ := ImportFromFile("test/sample1.txt")
	nodes := root.PreOrderIter()
	s := concatNodesLabelAndDescription(nodes)
	fmt.Print(s)
	// Output:0000
	// 1111
	// one one_one_one one one
	// one one
	// one_one one oneDEUX
	// two two two3333
	// 4444
	// this is a
	// multiline
	// comment5555
	// 6666
	// six six six_six_six
	// six six_six six
	// six six_six_six7777
	// This is a
	// multi-line label. It replaces
	// the "8888"
	// 9999
	// nine nineAAAA
	// aa aaa aaa a a aaaaaaa aaaaaa
	// a a aaaaaaa aaaaaa aaaa aaaaaa
	// a
	// aaa aaaaaa aaaa aa aa aaa aaa a
	// a aa aaBBBB
	// CCCC
	// DDDD
	// EEEE
	// FFFF
	// GGGG
	// HHHH
}
