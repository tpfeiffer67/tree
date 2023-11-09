package asciitree

import (
	"strings"

	"github.com/tpfeiffer67/paragraph"
	"github.com/tpfeiffer67/tree"
)

// LndLabelOnly renders only the label. A label can be multi-line.
func LndLabelOnly(n *tree.Node) (label []string, description []string) {
	if lbl, ok := n.Values()[KEY_LABEL]; ok {
		label = paragraph.Paragraph(strings.Split(lbl.(string), "\n"))
	}
	return
}
