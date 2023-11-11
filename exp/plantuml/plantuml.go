package plantuml

import (
	"strings"

	"github.com/tpfeiffer67/paragraph"
	"github.com/tpfeiffer67/tree"
	"github.com/tpfeiffer67/tree/exp/asciitree"
)

func LndPlantUMLSaltTree(n *tree.Node) ([]string, []string) {
	var label []string

	if lbl, ok := n.Values()[asciitree.KEY_LABEL]; ok {
		label = paragraph.Paragraph(strings.Split(lbl.(string), "\n")).MustacheNoErr(n.Values())
	}

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
	return label, []string{}
}

func BuildPlantUMLSaltTreeWidget(root *tree.Node) (plantuml paragraph.Paragraph) {
	const (
		startsalt = `@startsalt
{
{T`
		endsalt = `}
}
@endsalt`
	)

	plantuml = append(plantuml, startsalt)
	s := asciitree.Build(root, asciitree.Branchs{"+", "+", "", "++", "", "", "+", "+", "+", "+", "", ""}, LndPlantUMLSaltTree)
	plantuml = append(plantuml, s...)
	plantuml = append(plantuml, endsalt)
	return plantuml
}

/*
func LndPlantUMLMindMap(n *tree.Node) (label []string, description []string) {
	label = getAsOneLine(n, asciitree.KeyLabel, "%s", `\n`)
	description = getAsOneLine(n, asciitree.KeyDescription, "%s<i>", `</i>\n<i>`)
	return
}

func getAsOneLine(n *tree.Node, keyName string, left string, linesSep string) (result paragraph.Paragraph) {
	if s, ok := n.Values[keyName]; ok {
		result = paragraph.New(1)
		result[0] = left + strings.Replace(s.(string), "\n", linesSep, -1)
		result, _ = result.Mustache(n.Values)
	}
	return
}


type Branchs struct {
	AttachRoot      string // "▣ "
	AttachLabel     string // "="
	AttachLabelRepl string // "-" for multi-line label
	AttachComments  string // "~"
	VertTop        string // "┬"
	VertTopRepl    string // "─"
	VertWithBranch string // "├"
	Vert           string // "│"
	VertEnd        string // "└"
	VertSpace      string // "⦙"
	Horiz      string // "──"
	HorizSpace string // "."
}*/

/*

func BuildPlantUMLMindMap(root *tree.Node) (plantuml textview.StringsSlice) {
	const (
		startmindmap = "@startmindmap"
		stopmindmap  = "@endmindmap"
	)
	plantuml = append(plantuml, startmindmap)
	treetextview := textview.NewTextView().SetDefaultStyle("* {{label}}", "* {{label}}", "* {{label}}", "*", "*", "*", "*", "**_ %s", "**_ %s").BuildWithComments(root)
	for _, renderedline := range treetextview {
		plantuml = append(plantuml, renderedline)
	}
	plantuml = append(plantuml, stopmindmap)
	return plantuml
}

func BuildPlantUMLWorkBreakdownStructure(root *tree.Node) (plantuml textview.StringsSlice) {
	const (
		startwbs = "@startwbs"
		stopwbs  = "@endwbs"
	)
	plantuml = append(plantuml, startwbs)
	treetextview := textview.NewTextView().SetDefaultStyle("* <b>{{label}}</b>", "* <b>{{label}}</b>", "* <b>{{label}}</b>", "*", "*", "*", "*", "**_ %s", "**_ %s").BuildWithComments(root)
	for _, renderedline := range treetextview {
		plantuml = append(plantuml, renderedline)
	}
	plantuml = append(plantuml, stopwbs)
	return plantuml
}
*/
