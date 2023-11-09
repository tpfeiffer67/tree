package asciitree

import (
	"errors"
	"strings"
)

type TreeStyle int

const (
	TreeStyleCount     = 33
	TreeStyleMaxIndex  = int(TreeStyleSuperBold)
	TreeStyleLastValue = TreeStyleSuperBold
)

const (
	TreeStyleDebug TreeStyle = iota
	TreeStyleIndented
	TreeStyleIndentedVisible
	TreeStyleAscii1
	TreeStyleAscii2
	TreeStyleAsciiDots
	TreeStyleDots
	TreeStyleDashes22
	TreeStyleDashes22Bold
	TreeStyleDashes33
	TreeStyleDashes33Bold
	TreeStyleDashes44
	TreeStyleDashes44Bold
	TreeStyleSingleLine
	TreeStyleSingleLineRounded
	TreeStyleSingleVDoubleH
	TreeStyleSingleHDoubleV
	TreeStyleDoubleLine
	TreeStyleHeavy
	TreeStyleHorizontalHeavy
	TreeStyleVerticalHeavy
	TreeStyleBold
	TreeStyleTwoVerticalLines
	TreeStyleTwoVerticalLinesRounded
	TreeStyleTwoVerticalDoubleLines
	TreeStyleExtraBold1
	TreeStyleExtraBold2
	TreeStyleExtraBold3
	TreeStyleExtraBold4
	TreeStyleExtraBold5
	TreeStyleBoldDots
	TreeStyleExtraBoldDots
	TreeStyleSuperBold
)

func (v TreeStyle) String() string {
	return [...]string{
		"TreeStyleDebug",
		"TreeStyleIndented",
		"TreeStyleIndentedVisible",
		"TreeStyleAscii1",
		"TreeStyleAscii2",
		"TreeStyleAsciiDots",
		"TreeStyleDots",
		"TreeStyleDashes22",
		"TreeStyleDashes22Bold",
		"TreeStyleDashes33",
		"TreeStyleDashes33Bold",
		"TreeStyleDashes44",
		"TreeStyleDashes44Bold",
		"TreeStyleSingleLine",
		"TreeStyleSingleLineRounded",
		"TreeStyleSingleVDoubleH",
		"TreeStyleSingleHDoubleV",
		"TreeStyleDoubleLine",
		"TreeStyleHeavy",
		"TreeStyleHorizontalHeavy",
		"TreeStyleVerticalHeavy",
		"TreeStyleBold",
		"TreeStyleTwoVerticalLines",
		"TreeStyleTwoVerticalLinesRounded",
		"TreeStyleTwoVerticalDoubleLines",
		"TreeStyleExtraBold1",
		"TreeStyleExtraBold2",
		"TreeStyleExtraBold3",
		"TreeStyleExtraBold4",
		"TreeStyleExtraBold5",
		"TreeStyleBoldDots",
		"TreeStyleExtraBoldDots",
		"TreeStyleSuperBold",
	}[v]
}

func TreeStyleFromString(s string) (TreeStyle, error) {
	var suffix string
	if strings.HasPrefix(s, "TreeStyle") {
		l := len("TreeStyle")
		if l < len(s) {
			suffix = s[l:]
		}
	} else {
		suffix = s
	}
	switch suffix {
	case "Debug":
		return TreeStyleDebug, nil
	case "Indented":
		return TreeStyleIndented, nil
	case "IndentedVisible":
		return TreeStyleIndentedVisible, nil
	case "Ascii1":
		return TreeStyleAscii1, nil
	case "Ascii2":
		return TreeStyleAscii2, nil
	case "AsciiDots":
		return TreeStyleAsciiDots, nil
	case "Dots":
		return TreeStyleDots, nil
	case "Dashes22":
		return TreeStyleDashes22, nil
	case "Dashes22Bold":
		return TreeStyleDashes22Bold, nil
	case "Dashes33":
		return TreeStyleDashes33, nil
	case "Dashes33Bold":
		return TreeStyleDashes33Bold, nil
	case "Dashes44":
		return TreeStyleDashes44, nil
	case "Dashes44Bold":
		return TreeStyleDashes44Bold, nil
	case "SingleLine":
		return TreeStyleSingleLine, nil
	case "SingleLineRounded":
		return TreeStyleSingleLineRounded, nil
	case "SingleVDoubleH":
		return TreeStyleSingleVDoubleH, nil
	case "SingleHDoubleV":
		return TreeStyleSingleHDoubleV, nil
	case "DoubleLine":
		return TreeStyleDoubleLine, nil
	case "Heavy":
		return TreeStyleHeavy, nil
	case "HorizontalHeavy":
		return TreeStyleHorizontalHeavy, nil
	case "VerticalHeavy":
		return TreeStyleVerticalHeavy, nil
	case "Bold":
		return TreeStyleBold, nil
	case "TwoVerticalLines":
		return TreeStyleTwoVerticalLines, nil
	case "TwoVerticalLinesRounded":
		return TreeStyleTwoVerticalLinesRounded, nil
	case "TwoVerticalDoubleLines":
		return TreeStyleTwoVerticalDoubleLines, nil
	case "ExtraBold1":
		return TreeStyleExtraBold1, nil
	case "ExtraBold2":
		return TreeStyleExtraBold2, nil
	case "ExtraBold3":
		return TreeStyleExtraBold3, nil
	case "ExtraBold4":
		return TreeStyleExtraBold4, nil
	case "ExtraBold5":
		return TreeStyleExtraBold5, nil
	case "BoldDots":
		return TreeStyleBoldDots, nil
	case "ExtraBoldDots":
		return TreeStyleExtraBoldDots, nil
	case "SuperBold":
		return TreeStyleSuperBold, nil
	}
	return TreeStyle(0), errors.New("String does not correspond to any existing TreeStyle values")
}
