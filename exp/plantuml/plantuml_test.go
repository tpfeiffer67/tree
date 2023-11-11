package plantuml

import (
	"fmt"

	"github.com/tpfeiffer67/tree/imp/indentedtext"
)

const Tree1 = `0000
	1111
	 style="italic"
	 description='''one one oneone oneoneone one
	 one oneone one'''
		2222
		 style="bold"
		 description='''two two two'''
	3333
		4444
			5555
	6666
	 style="italic"
	 description='''six six sixsixsixsix
	 six sixsix six
	 sixsixsixsix'''
		7777
			8888
			9999
			 style="bold"
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

const Tree2 = `0000
	1111
	 description='''one one oneone oneoneone one
	 one oneone one'''
		2222
		 color="#FFBBCC"
		 style="bold"
		 description='''two two two'''
	3333
	 color="#CCCCFF"
		4444
		 label='''444
		 444
		 444'''
			5555
	6666
	 description='''six <b>six</b> sixsixsixsix
	 six sixsix six
	 sixsixsixsix'''
		7777
			8888
			9999
			 color="#FFBBCC"
			 description="nine nine"
			AAAA
			 aa aaa aaa a a <i>aaaaaaa</i> aaaaaa
			 a a <b>aaaaaaa aaaaaa</b> aaaa aaaaaa
			 aaa aaaaaa aaaa aa aa aaa aaa a
			 a
				BBBB
					CCCC
		DDDD
			EEEE
			 color="#CCFFCC"
				FFFF
					GGGG
	HHHH
`

func ExampleBuildPlantUMLSaltTreeWidget() {
	root, _ := indentedtext.Import(Tree1)
	fmt.Println(BuildPlantUMLSaltTreeWidget(root, true).String())
	fmt.Println(BuildPlantUMLSaltTreeWidget(root, false).String())
	// Output:
	// @startsalt
	// {
	// {T
	// +0000
	// ++1111
	// ++++<i>one one oneone oneoneone one</i>
	// ++++<i>one oneone one</i>
	// +++2222
	// +++++<b>two two two</b>
	// ++3333
	// +++4444
	// ++++5555
	// ++6666
	// ++++<i>six six sixsixsixsix</i>
	// ++++<i>six sixsix six</i>
	// ++++<i>sixsixsixsix</i>
	// +++7777
	// ++++8888
	// ++++9999
	// ++++++<b>nine nine</b>
	// ++++AAAA
	// ++++++aa aaa aaa a a aaaaaaa aaaaaa
	// ++++++a a aaaaaaa aaaaaa aaaa aaaaaa
	// ++++++aaa aaaaaa aaaa aa aa aaa aaa a
	// ++++++a
	// +++++BBBB
	// ++++++CCCC
	// +++DDDD
	// ++++EEEE
	// +++++FFFF
	// ++++++GGGG
	// ++HHHH
	// }
	// }
	// @endsalt
	//
	// @startsalt
	// {
	// {T
	// +0000
	// ++1111
	// +++2222
	// ++3333
	// +++4444
	// ++++5555
	// ++6666
	// +++7777
	// ++++8888
	// ++++9999
	// ++++AAAA
	// +++++BBBB
	// ++++++CCCC
	// +++DDDD
	// ++++EEEE
	// +++++FFFF
	// ++++++GGGG
	// ++HHHH
	// }
	// }
	// @endsalt
}

func ExampleBuildPlantUMLMindMap() {
	root, _ := indentedtext.Import(Tree2)
	fmt.Println(BuildPlantUMLMindMap(root, true).String())
	fmt.Println(BuildPlantUMLMindMap(root, false).String())
	// Output:
	// @startmindmap
	// + 0000
	// ++ 1111
	// +++_ one one oneone oneoneone one\none oneone one
	// +++[#FFBBCC] 2222
	// ++++_ two two two
	// ++[#CCCCFF] 3333
	// +++ 444\n444\n444
	// ++++ 5555
	// ++ 6666
	// +++_ six <b>six</b> sixsixsixsix\nsix sixsix six\nsixsixsixsix
	// +++ 7777
	// ++++ 8888
	// ++++[#FFBBCC] 9999
	// +++++_ nine nine
	// ++++ AAAA
	// +++++_ aa aaa aaa a a <i>aaaaaaa</i> aaaaaa\na a <b>aaaaaaa aaaaaa</b> aaaa aaaaaa\naaa aaaaaa aaaa aa aa aaa aaa a\na
	// +++++ BBBB
	// ++++++ CCCC
	// +++ DDDD
	// ++++[#CCFFCC] EEEE
	// +++++ FFFF
	// ++++++ GGGG
	// ++ HHHH
	// @endmindmap
	//
	// @startmindmap
	// + 0000
	// ++ 1111
	// +++[#FFBBCC] 2222
	// ++[#CCCCFF] 3333
	// +++ 444\n444\n444
	// ++++ 5555
	// ++ 6666
	// +++ 7777
	// ++++ 8888
	// ++++[#FFBBCC] 9999
	// ++++ AAAA
	// +++++ BBBB
	// ++++++ CCCC
	// +++ DDDD
	// ++++[#CCFFCC] EEEE
	// +++++ FFFF
	// ++++++ GGGG
	// ++ HHHH
	// @endmindmap
}

func ExampleBuildPlantUMLWorkBreakdownStructure() {
	root, _ := indentedtext.Import(Tree2)
	fmt.Println(BuildPlantUMLWorkBreakdownStructure(root, true).String())
	fmt.Println(BuildPlantUMLWorkBreakdownStructure(root, false).String())
	// Output:
	// @startwbs
	// + 0000
	// ++ 1111
	// +++_ one one oneone oneoneone one\none oneone one
	// +++[#FFBBCC] 2222
	// ++++_ two two two
	// ++[#CCCCFF] 3333
	// +++ 444\n444\n444
	// ++++ 5555
	// ++ 6666
	// +++_ six <b>six</b> sixsixsixsix\nsix sixsix six\nsixsixsixsix
	// +++ 7777
	// ++++ 8888
	// ++++[#FFBBCC] 9999
	// +++++_ nine nine
	// ++++ AAAA
	// +++++_ aa aaa aaa a a <i>aaaaaaa</i> aaaaaa\na a <b>aaaaaaa aaaaaa</b> aaaa aaaaaa\naaa aaaaaa aaaa aa aa aaa aaa a\na
	// +++++ BBBB
	// ++++++ CCCC
	// +++ DDDD
	// ++++[#CCFFCC] EEEE
	// +++++ FFFF
	// ++++++ GGGG
	// ++ HHHH
	// @endwbs
	//
	// @startwbs
	// + 0000
	// ++ 1111
	// +++[#FFBBCC] 2222
	// ++[#CCCCFF] 3333
	// +++ 444\n444\n444
	// ++++ 5555
	// ++ 6666
	// +++ 7777
	// ++++ 8888
	// ++++[#FFBBCC] 9999
	// ++++ AAAA
	// +++++ BBBB
	// ++++++ CCCC
	// +++ DDDD
	// ++++[#CCFFCC] EEEE
	// +++++ FFFF
	// ++++++ GGGG
	// ++ HHHH
	// @endwbs
}
