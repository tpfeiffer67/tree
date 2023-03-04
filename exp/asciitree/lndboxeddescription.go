package asciitree

import (
	"strings"

	"github.com/tpfeiffer67/multistrings"
	"github.com/tpfeiffer67/tree"
)

func LndBoxedDescription(n *tree.Node) (label []string, description []string) {
	if lbl, ok := n.Values()[KEY_LABEL]; ok {
		label = multistrings.MultiStrings(strings.Split(lbl.(string), "\n"))
	}
	if desc, ok := n.Values()[KEY_DESCRIPTION]; ok {
		description = multistrings.MultiStrings(strings.Split(desc.(string), "\n")).AutoBox(multistrings.BoxSettings{Width: 40, TopLabel: "", TopLabelAlign: multistrings.LabelAlignLeft, BottomLabel: "", BottomLabelAlign: multistrings.LabelAlignLeft}, multistrings.GetBoxPattern(multistrings.BoxStyleSingleLine))
	}
	return
}
