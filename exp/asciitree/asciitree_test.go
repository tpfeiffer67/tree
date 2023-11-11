package asciitree

import (
	"fmt"

	"github.com/tpfeiffer67/tree/imp/indentedtext"
)

const (
	Tree1 = `0000
	1111
	 style=1
	 description='''one one oneone oneoneone one
	 one oneone one'''
		2222
		 style=2
		 description="two two two"
	3333
		4444
			5555
	6666
	 style=1
	 description='''six six sixsixsixsix
	 six sixsix six
	 sixsixsixsix'''
		7777
			8888
			9999
			 style=2
			 description="nine nine"
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
	// │ │ one one oneone oneoneone one
	// │ │ one oneone one
	// │ └───2222
	// │     two two two
	// ├─┬─3333
	// │ └─┬─4444
	// │   └───5555
	// ├─┬─6666
	// │ │ six six sixsixsixsix
	// │ │ six sixsix six
	// │ │ sixsixsixsix
	// │ ├─┬─7777
	// │ │ ├───8888
	// │ │ ├───9999
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
