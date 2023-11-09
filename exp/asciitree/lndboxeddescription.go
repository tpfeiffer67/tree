package asciitree

import (
	"strings"

	"github.com/tpfeiffer67/paragraph"
	"github.com/tpfeiffer67/tree"
)

func LndBoxedDescription(n *tree.Node) (label []string, description []string) {
	if lbl, ok := n.Values()[KEY_LABEL]; ok {
		label = paragraph.Paragraph(strings.Split(lbl.(string), "\n"))
	}
	if desc, ok := n.Values()[KEY_DESCRIPTION]; ok {
		description = paragraph.Paragraph(strings.Split(desc.(string), "\n")).AutoBox(paragraph.BoxSettings{Width: 40, TopLabel: "", TopLabelAlign: paragraph.LabelAlignLeft, BottomLabel: "", BottomLabelAlign: paragraph.LabelAlignLeft}, paragraph.GetBoxPattern(paragraph.BoxStyleSingleLine))
	}
	return
}
