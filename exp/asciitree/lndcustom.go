package asciitree

import (
	"errors"
	"fmt"
	"image/color"
	"strings"

	"github.com/tpfeiffer67/paragraph"
	"github.com/tpfeiffer67/tree"
)

type LabelRenderOptions struct {
	Name             string // default is "label"
	MaxWidth         int    // max width for the label
	FormatLeft       string
	FormatRight      string
	ExtraLinesBefore int
	ExtraLinesAfter  int
}

type DescriptionRenderOptions struct {
	Name             string // default is "description"
	Show             bool
	MaxWidth         int // max width for the description text
	FormatLeft       string
	FormatRight      string
	ExtraLinesBefore int
	ExtraLinesAfter  int
}

type ValuesRenderOptions struct {
	Show          bool
	ExclusionList []string
	Format        string
}

type BoxRenderOptions struct {
	Settings paragraph.BoxSettings
	Pattern  paragraph.BoxPattern
}

type LndRenderOptions struct {
	Label       LabelRenderOptions
	Description DescriptionRenderOptions
	Box         BoxRenderOptions
	Values      ValuesRenderOptions
	Mustache    bool
}

func GetLndCustom(ropt *LndRenderOptions) func(*tree.Node) ([]string, []string) {
	return func(n *tree.Node) ([]string, []string) {
		var label paragraph.Paragraph
		var description paragraph.Paragraph

		label = getLinesAndFormat(n, ropt.Label.Name, ropt.Label.FormatLeft, ropt.Label.FormatRight, ropt.Mustache, ropt.Label.ExtraLinesBefore, ropt.Label.ExtraLinesAfter)

		if ropt.Description.Show {
			description = getLinesAndFormat(n, ropt.Description.Name, ropt.Description.FormatLeft, ropt.Description.FormatRight, ropt.Mustache, ropt.Description.ExtraLinesBefore, ropt.Description.ExtraLinesAfter)
		}

		if ropt.Values.Show {
			values := paragraph.New(4)
			for key, value := range n.Values() {
				if !stringInSlice(key, ropt.Values.ExclusionList) {
					if len(ropt.Values.Format) > 0 {
						values = append(values, fmt.Sprintf(ropt.Values.Format, key, value))
					} else {
						values = append(values, fmt.Sprintf("%s=%v", key, value))
					}
				}
			}
			description = description.Append(values.Sort())
		}
		return label, description
	}
}

func getLinesAndFormat(n *tree.Node, valueName string, formatLeft, formatRight string, mustache bool, extraLinesBefore int, extraLinesAfter int) paragraph.Paragraph {
	if value, ok := n.Values()[valueName]; ok {
		lns := paragraph.Paragraph(strings.Split(value.(string), "\n"))
		if len(formatLeft) > 0 || len(formatRight) > 0 {
			lns = lns.Surround(formatLeft, formatRight)
		}
		if mustache {
			lns, _ = lns.Mustache(n.Values())
		}
		maxWidth, cut, box, accolades := getAttributes(n, valueName)
		return formatLines(lns, maxWidth, cut, box, accolades, extraLinesBefore, extraLinesAfter)
	}
	return paragraph.New(0)
}

func getAttributes(n *tree.Node, prefix string) (maxWidth int, cut bool, box paragraph.BoxStyle, accolades paragraph.AccoladesStyle) {
	if w, ok := n.Values()[prefix+"Max"]; ok {
		maxWidth = w.(int)
	} else {
		maxWidth = 1000
	}

	if b, ok := n.Values()[prefix+"Box"]; ok {
		box, _ = paragraph.BoxStyleFromString(b.(string))
	} else {
		box = paragraph.BoxStyleNone
	}

	if a, ok := n.Values()[prefix+"Accolades"]; ok {
		accolades, _ = paragraph.AccoladesStyleFromString(a.(string))
	} else {
		accolades = paragraph.AccoladesStyleNone
	}

	// TODO implement
	cut = false

	return
}

func formatLines(linesIn paragraph.Paragraph, maxWidth int, cut bool, box paragraph.BoxStyle, accolades paragraph.AccoladesStyle, extraLinesBefore int, extraLinesAfter int) paragraph.Paragraph {
	if maxWidth > 0 {
		if cut {
			linesIn = linesIn.Cut(maxWidth)
		} else {
			linesIn = linesIn.Limit(maxWidth)
		}
	}

	if box != paragraph.BoxStyleNone {
		linesIn = linesIn.AutoBox(paragraph.BoxSettings{
			Width:            maxWidth,
			TopLabel:         "",
			TopLabelAlign:    paragraph.LabelAlignLeft,
			BottomLabel:      "",
			BottomLabelAlign: paragraph.LabelAlignLeft}, paragraph.GetBoxPattern(box))
	}

	if accolades != paragraph.AccoladesStyleNone {
		linesIn = linesIn.AutoAccolades(accolades)
	}

	if extraLinesBefore > 0 {
		linesIn = paragraph.NewWithPresetContent("", extraLinesBefore).Append(linesIn)
	}

	if extraLinesAfter > 0 {
		linesIn = linesIn.Append(paragraph.NewWithPresetContent("", extraLinesAfter))
	}

	return linesIn
}

func colorizeLinesHTML(linesIn paragraph.Paragraph, hexColor string) paragraph.Paragraph {
	const trueColor = `<span style="color:%s">`
	const colorOff = "</span>"
	s := fmt.Sprintf(trueColor, hexColor)
	return linesIn.Surround(s, colorOff)
}

func colorizeLinesConsole(linesIn paragraph.Paragraph, hexColor string) paragraph.Paragraph {
	// console true color : https://gist.github.com/XVilka/8346728
	const trueColor = "\033[38;2;%d;%d;%dm"
	const colorOff = "\033[39;49m"
	col, _ := parseHexColorFast(hexColor)
	s := fmt.Sprintf(trueColor, col.R, col.G, col.B)
	return linesIn.Surround(s, colorOff)
}

func getColor(n *tree.Node, prefix string) (hexColor string) {
	if c, ok := n.Values()[prefix+"Color"]; ok {
		hexColor = c.(string)
	} else {
		hexColor = "#FDFDFD"
	}
	return
}

// Found here : https://stackoverflow.com/questions/54197913/parse-hex-string-to-image-color
var errInvalidFormat = errors.New("Invalid #hexa color format")

func parseHexColorFast(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}

// Found here https://stackoverflow.com/questions/15323767/does-go-have-if-x-in-construct-similar-to-python (tanks to minikomi)
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
