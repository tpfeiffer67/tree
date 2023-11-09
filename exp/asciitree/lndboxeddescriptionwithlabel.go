package asciitree

import (
	"strings"

	"github.com/tpfeiffer67/paragraph"
	"github.com/tpfeiffer67/tree"
)

func LndBoxedDescriptionWithTitle(n *tree.Node) (label []string, description []string) {
	if lbl, ok := n.Values()[KEY_LABEL]; ok {
		theLabel := lbl.(string)

		if cmts, ok := n.Values()[KEY_DESCRIPTION]; ok {
			lns := strings.Split(cmts.(string), "\n")
			label = paragraph.Paragraph(lns).AutoBox(paragraph.BoxSettings{Width: 40, TopLabel: theLabel, TopLabelAlign: paragraph.LabelAlignLeft, BottomLabel: "", BottomLabelAlign: paragraph.LabelAlignLeft}, paragraph.GetBoxPattern(paragraph.BoxStyleSingleLine))
		} else {
			label = paragraph.NewFromString(theLabel)
		}
	}
	return
}
