package asciitree

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tpfeiffer67/paragraph"
	"github.com/tpfeiffer67/runestr"
	"github.com/tpfeiffer67/tree"
)

type Branchs struct {
	AttachRoot        string // "▣ "
	AttachLabel       string // " "
	AttachLabelRepl   string // for multi-line label
	AttachDescription string // " "

	VertTop        string // "┬"
	VertTopRepl    string // "─"
	VertWithBranch string // "├"
	Vert           string // "│"
	VertEnd        string // "└"
	VertRepl       string // " "

	Horiz     string // "─"
	HorizRepl string // " "
}

func NewTreePatternFromString(strPattern string, sep string) (Branchs, error) {
	items := strings.Split(strPattern, sep)
	if len(items) != 12 {
		return Branchs{"", "", "", "", "", "", "", "", "", "", "", ""}, errors.New("Wrong number of items in tree pattern")
	}
	var br Branchs
	br.AttachRoot = items[0]
	br.AttachLabel = items[1]
	br.AttachLabelRepl = items[2]
	br.AttachDescription = items[3]
	br.VertTop = items[4]
	br.VertTopRepl = items[5]
	br.VertWithBranch = items[6]
	br.Vert = items[7]
	br.VertEnd = items[8]
	br.VertRepl = items[9]
	br.Horiz = items[10]
	br.HorizRepl = items[11]
	return br, nil
}

func (br Branchs) Modify(hideVertTop bool, attachLabelLength int, attachDescriptionLength int, horizLength int, childBranchOffset int) Branchs {

	switch {
	case attachLabelLength == 0:
		br.AttachLabel = ""
		br.AttachLabelRepl = ""
	case attachLabelLength > 1:
		br.AttachLabel = strings.Repeat(br.AttachLabel, attachLabelLength)
		br.AttachLabelRepl = strings.Repeat(br.AttachLabelRepl, attachLabelLength)
	}

	// If we suppress the vertical top branch, then the child branch is attached directly to the label
	//
	//  ──┬LABEL      (hideVertTop = false)
	//    └─child
	//
	//  ──LABEL       (hideVertTop = true)
	//    └─child
	//
	if hideVertTop {
		br.VertTop = ""
		br.VertTopRepl = ""
		// The attachLabelRepl must be adjusted for the multi-lines labels
		l := runestr.Length(br.AttachLabel)
		if l > 0 {
			br.AttachLabelRepl = runestr.Left(br.AttachLabelRepl, l-1)
		}
	}

	switch {
	case attachDescriptionLength == 0:
		br.AttachDescription = ""
	case attachDescriptionLength > 1:
		br.AttachDescription = strings.Repeat(br.AttachDescription, attachDescriptionLength)
	}

	horizSpace := br.HorizRepl
	switch {
	case horizLength == 0:
		br.Horiz = ""
		br.HorizRepl = ""
	case horizLength > 1:
		br.Horiz = strings.Repeat(br.Horiz, horizLength)
		br.HorizRepl = strings.Repeat(br.HorizRepl, horizLength)
	}

	if childBranchOffset > 0 {
		br.HorizRepl = br.HorizRepl + strings.Repeat(horizSpace, childBranchOffset)
	}

	return br
}

func (br Branchs) String() string {
	return fmt.Sprintf(`Branchs{"%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s"}`,
		br.AttachRoot,
		br.AttachLabel,
		br.AttachLabelRepl,
		br.AttachDescription,
		br.VertTop,
		br.VertTopRepl,
		br.VertWithBranch,
		br.Vert,
		br.VertEnd,
		br.VertRepl,
		br.Horiz,
		br.HorizRepl)
}

func (br Branchs) Build(root *tree.Node, renderer func(*tree.Node) ([]string, []string)) (result paragraph.Paragraph) {
	prcssStr := func(n *tree.Node, s string) {
		result = append(result, s)
	}
	build(root, br, 0, "", false, renderer, prcssStr)
	return
}

func Build(root *tree.Node, br Branchs, renderer func(*tree.Node) ([]string, []string)) (result paragraph.Paragraph) {
	prcssStr := func(n *tree.Node, s string) {
		result = append(result, s)
	}
	build(root, br, 0, "", false, renderer, prcssStr)
	return
}

func BuildToValues(root *tree.Node, br Branchs, renderer func(*tree.Node) ([]string, []string), valueName string) {
	prcssStr := func(n *tree.Node, s string) {
		v := n.Values()[valueName]
		if v != nil {
			n.Values()[valueName] = v.(string) + "\n" + s
		} else {
			n.Values()[valueName] = s
		}
	}
	build(root, br, 0, "", false, renderer, prcssStr)
}

func build(n *tree.Node, br Branchs, deep int, indent string, isLastChild bool, renderer func(*tree.Node) ([]string, []string), processString func(*tree.Node, string)) (result paragraph.Paragraph) {
	var s string

	childrenCount := len(n.Children())
	label, description := renderer(n)

	// TODO Tester si %s dans la chaine, sinon l'ajouter à gauche
	// Dessin de l'attache selon le type de ligne
	if n.IsRoot() {
		s = br.AttachRoot
		// TODO vérifier plantage si pas de label
		processString(n, s+label[0])
	} else {
		var i int
		linesCount := len(label)
		linesBeforeAttach := (linesCount - 1) / 2
		linesAfterAttach := (linesCount-1)/2 + (linesCount-1)%2

		if linesBeforeAttach > 0 {
			s = indent + br.Vert + br.HorizRepl + br.VertRepl + br.AttachLabelRepl
			// Add lines before attach
			for j := 0; j < linesBeforeAttach; j++ {
				processString(n, s+label[i])
				i++
			}
		}

		if isLastChild {
			s = indent + br.VertEnd + br.Horiz
			indent = indent + br.VertRepl + br.HorizRepl
		} else {
			s = indent + br.VertWithBranch + br.Horiz
			indent = indent + br.Vert + br.HorizRepl
		}

		if childrenCount > 0 {
			s = s + br.VertTop
		} else {
			s = s + br.VertTopRepl
		}
		s = s + br.AttachLabel

		// attach
		processString(n, s+label[i])
		i++

		// lines after attach
		if linesAfterAttach > 0 {
			if childrenCount > 0 {
				s = indent + br.Vert + br.AttachLabelRepl
			} else {
				s = indent + br.VertRepl + br.AttachLabelRepl
			}
			for j := 0; j < linesAfterAttach; j++ {
				processString(n, s+label[i])
				i++
			}
		}
	}

	if len(description) > 0 {
		var descriptionAttach string
		if childrenCount > 0 {
			descriptionAttach = br.Vert + br.AttachDescription
		} else {
			descriptionAttach = br.VertRepl + br.AttachDescription
		}
		s = indent + descriptionAttach

		for _, line := range description {
			processString(n, s+line)
		}
	}

	if childrenCount > 0 {
		for i, child := range n.Children() {
			slist := build(child, br, deep+1, indent, (i == childrenCount-1), renderer, processString)
			for _, line := range slist {
				processString(child, line)
				result = append(result, line)
			}
		}
	}
	return result
}
