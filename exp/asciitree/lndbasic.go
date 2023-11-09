package asciitree

import (
	"strings"

	"github.com/tpfeiffer67/paragraph"
	"github.com/tpfeiffer67/tree"
)

// LndBasic render text tree with a label (that can be multi-line)
// and its description (if available)
//
// The label and description are extracted from the values stored in the node
// using the keys KeyLabel and KeyDescription, respectively. The values are
// then split into lines using the Lines function from the lines package and returned.
//
// If the KeyLabel or KeyDescription keys are not present in the node's values,
// the corresponding label or description will be an empty slice.
func LndBasic(n *tree.Node) (label []string, description []string) {
	if lbl, ok := n.Values()[KEY_LABEL]; ok {
		label = paragraph.Paragraph(strings.Split(lbl.(string), "\n"))
	}
	if desc, ok := n.Values()[KEY_DESCRIPTION]; ok {
		description = paragraph.Paragraph(strings.Split(desc.(string), "\n"))
	}
	return
}
