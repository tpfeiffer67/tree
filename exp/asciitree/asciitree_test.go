package asciitree

import (
	"fmt"

	"testing"

	"github.com/tpfeiffer67/tree/imp/indentedtext"
)

const (
	Tree1 = `0000
	1111
	 style=1
	 one one oneone oneoneone one
	 one oneone one
		2222
		 style=2
		 two two two
	3333
		4444
			5555
	6666
	 style=1
	 six six sixsixsixsix
	 six sixsix six
	 sixsixsixsix
		7777
			8888
			9999
			 style=2
			 nine nine
			AAAA
			 aa aaa aaa a a aaaaaaa aaaaaa
			 a a aaaaaaa aaaaaa aaaa aaaaaa
			 aaa aaaaaa aaaa aa aa aaa aaa a
			 a
				BBBB
					CCCC
		DDDD
			EEEE
				FFFF
					GGGG
	HHHH
`
)

func ExampleExpText1() {
	root, _ := indentedtext.Import(Tree1)
	s := Build(root, GetTreePattern(TreeStyleSingleLine), LndLabelOnly)
	fmt.Print(s)
	// Output:
	// ■0000
	// ├─┬─1111
	// │ └───2222
	// ├─┬─3333
	// │ └─┬─4444
	// │   └───5555
	// ├─┬─6666
	// │ ├─┬─7777
	// │ │ ├───8888
	// │ │ ├───9999
	// │ │ └─┬─AAAA
	// │ │   └─┬─BBBB
	// │ │     └───CCCC
	// │ └─┬─DDDD
	// │   └─┬─EEEE
	// │     └─┬─FFFF
	// │       └───GGGG
	// └───HHHH
}

func ExampleExpText2() {
	root, _ := indentedtext.Import(Tree1)
	s := Build(root, GetTreePattern(TreeStyleSingleLine), LndBasic) // s = asciitree.Build(root, asciitree.GetTreePattern(asciitree.TreeStyleSingleLine), asciitree.LndBasic)
	fmt.Print(s)
	// Output:
	// ■0000
	// ├─┬─1111
	// │ │ style=1
	// │ │ one one oneone oneoneone one
	// │ │ one oneone one
	// │ └───2222
	// │     style=2
	// │     two two two
	// ├─┬─3333
	// │ └─┬─4444
	// │   └───5555
	// ├─┬─6666
	// │ │ style=1
	// │ │ six six sixsixsixsix
	// │ │ six sixsix six
	// │ │ sixsixsixsix
	// │ ├─┬─7777
	// │ │ ├───8888
	// │ │ ├───9999
	// │ │ │   style=2
	// │ │ │   nine nine
	// │ │ └─┬─AAAA
	// │ │   │ aa aaa aaa a a aaaaaaa aaaaaa
	// │ │   │ a a aaaaaaa aaaaaa aaaa aaaaaa
	// │ │   │ aaa aaaaaa aaaa aa aa aaa aaa a
	// │ │   │ a
	// │ │   └─┬─BBBB
	// │ │     └───CCCC
	// │ └─┬─DDDD
	// │   └─┬─EEEE
	// │     └─┬─FFFF
	// │       └───GGGG
	// └───HHHH
}

func TestBuild(t *testing.T) {
	/*
	   var root *tree.Node
	   //var s string
	   //var treetextview []string

	   root = imp.BuildTreeFromIndentedTextAndTomlData(samples.TreeSample1, "label", "description")

	   debugTree := asciitree.Build(root, asciitree.GetTreePattern(asciitree.TreeStyleSingleLineRounded), LndBasic)
	   fmt.Println(debugTree)

	   tview := NewTextView()
	   tview.SetDefaultStyle("{{label}}", "┬ {{label}}", "─ {{label}}", "├─", "└─", "  ", "│ ", "  - %s", "│ - %s")

	   treetextview = tview.BuildWithComments(root)
	   s = strslc.LinesToString(treetextview)
	   // TODO remplacer par exemple
	   assert.Equal(t, testTree1, s)

	   treetextview = tview.Build(root)
	   s = strslc.LinesToString(treetextview)
	   // TODO remplacer par exemple
	   assert.Equal(t, testTree2, s)

	   root.Adopt(NewLabelNode("IIII", 0, ""))
	   root.Adopt(NewLabelNode("JJJJ", 0, "jjj j jjj jj"))
	   treetextview = tview.BuildWithComments(root)
	   s = strslc.LinesToString(treetextview)
	   assert.Equal(t, testTree3, s)
	*/
}
