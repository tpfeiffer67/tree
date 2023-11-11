package plantuml

import (
	"fmt"
	"strings"

	"github.com/tpfeiffer67/paragraph"
	"github.com/tpfeiffer67/tree"
	"github.com/tpfeiffer67/tree/exp/asciitree"
)

func BuildPlantUMLSaltTreeWidget(root *tree.Node, withDescription bool) paragraph.Paragraph {
	const (
		startsalt = `@startsalt
{
{T`
		endsalt = `}
}
@endsalt`
	)

	LndPlantUMLSaltTree := func(n *tree.Node) ([]string, []string) {
		var label []string

		if lbl, ok := n.Values()[asciitree.KEY_LABEL]; ok {
			label = paragraph.Paragraph(strings.Split(lbl.(string), "\n")).MustacheNoErr(n.Values())
		}

		if withDescription {
			if desc, ok := n.Values()[asciitree.KEY_DESCRIPTION]; ok {
				description := paragraph.Paragraph(strings.Split(desc.(string), "\n")).MustacheNoErr(n.Values())
				if style, ok := n.Values()["style"]; ok {
					switch style.(string) {
					case "italic":
						return label, description.Surround("<i>", "</i>")
					case "bold":
						return label, description.Surround("<b>", "</b>")
					}
				}
				return label, description
			}
		}
		return label, []string{}
	}
	return compose(startsalt, asciitree.Build(root, asciitree.Branchs{"+", "+", "", "++", "", "", "+", "+", "+", "+", "", ""}, LndPlantUMLSaltTree), endsalt)
}

func BuildPlantUMLmindmap(root *tree.Node, withDescription bool) paragraph.Paragraph {
	LndPlantUMLMindMap := func(n *tree.Node) ([]string, []string) {
		var label []string

		if lbl, ok := n.Values()[asciitree.KEY_LABEL]; ok {
			l := " " + strings.Replace(lbl.(string), "\n", `\n`, -1)

			if color, ok := n.Values()["color"]; ok {
				l = fmt.Sprintf("[%s]%s", color, l)
			}

			label = paragraph.Paragraph([]string{l}).MustacheNoErr(n.Values())
		}

		if withDescription {
			if desc, ok := n.Values()[asciitree.KEY_DESCRIPTION]; ok {
				l := strings.Replace(desc.(string), "\n", `\n`, -1)
				description := paragraph.Paragraph([]string{l}).MustacheNoErr(n.Values())
				return label, description
			}
		}
		return label, []string{}
	}
	return asciitree.Build(root, asciitree.Branchs{"+", "+", "", "+_ ", "", "", "+", "+", "+", "+", "", ""}, LndPlantUMLMindMap)
}

func BuildPlantUMLMindMap(root *tree.Node, withDescription bool) paragraph.Paragraph {
	return compose(`@startmindmap`, BuildPlantUMLmindmap(root, withDescription), `@endmindmap`)
}

func BuildPlantUMLWorkBreakdownStructure(root *tree.Node, withDescription bool) paragraph.Paragraph {
	return compose(`@startwbs`, BuildPlantUMLmindmap(root, withDescription), `@endwbs`)
}

func compose(s1 string, strslice []string, s2 string) []string {
	result := []string{s1}
	result = append(result, strslice...)
	result = append(result, s2)
	return result
}
