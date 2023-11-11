package plantuml

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

const testPlantUML1 = `@startsalt
{
{T
+0000
++1111
++++<i>one one oneone oneoneone one</i>
++++<i>one oneone one</i>
+++2222
+++++<b>two two two</b>
++3333
+++4444
++++5555
++6666
++++<i>six six sixsixsixsix</i>
++++<i>six sixsix six</i>
++++<i>sixsixsixsix</i>
+++7777
++++8888
++++9999
++++++<b>nine nine</b>
++++AAAA
++++++aa aaa aaa a a aaaaaaa aaaaaa
++++++a a aaaaaaa aaaaaa aaaa aaaaaa
++++++aaa aaaaaa aaaa aa aa aaa aaa a
++++++a
+++++BBBB
++++++CCCC
+++DDDD
++++EEEE
+++++FFFF
++++++GGGG
++HHHH
}
}
@endsalt
`

func TestBuildPlantUMLSaltTreeWidget(t *testing.T) {
	root, _ := indentedtext.Import(Tree1)
	s := BuildPlantUMLSaltTreeWidget(root).String()
	assert.Equal(t, testPlantUML1, s)
}
