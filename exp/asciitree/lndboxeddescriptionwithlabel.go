package asciitree

import (
	"strings"

	"github.com/tpfeiffer67/multistrings"
	"github.com/tpfeiffer67/tree"
)

func LndBoxedDescriptionWithTitle(n *tree.Node) (label []string, description []string) {
	if lbl, ok := n.Values()[KEY_LABEL]; ok {
		theLabel := lbl.(string)

		if cmts, ok := n.Values()[KEY_DESCRIPTION]; ok {
			lns := strings.Split(cmts.(string), "\n")
			label = multistrings.MultiStrings(lns).AutoBox(multistrings.BoxSettings{Width: 40, TopLabel: theLabel, TopLabelAlign: multistrings.LabelAlignLeft, BottomLabel: "", BottomLabelAlign: multistrings.LabelAlignLeft}, multistrings.GetBoxPattern(multistrings.BoxStyleSingleLine))
		} else {
			label = multistrings.NewFromString(theLabel)
		}
	}
	return
}
