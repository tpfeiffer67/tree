package asciitree

/*
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

var TreePatternsCollection = [TreeStyleCount]Branchs{
	{"▣", "=", "-", "~", "┬", "─", "├", "│", "└", ":", "═", "."},                        // Debug
	{"", "", " ", " ", "", "", "", "", "", "", "\t", "\t"},                              // indented text (tab = indentation and space = description)
	{"", "", "⎵", "⎵", "", "", "", "", "", "", "🠚", "🠚"},                                // indented text without invisible characters
	{"*", "-", " ", " ", "+", "-", "|", "|", "`", " ", "-", " "},                        // Ascii1
	{"*", "-", " ", " ", "+", "-", "|", "|", `\`, " ", "-", " "},                        // Ascii2
	{"¤", ".", " ", " ", ".", ".", ":", ":", ":", " ", ".", " "},                        // AsciiDots
	{"¤", "…", " ", " ", "…", "…", "⦙", "⦙", "⦙", " ", "…", " "},                        // Dots
	{"¤", "╌", " ", " ", "╌", "╌", "╎", "╎", "╎", " ", "╌", " "},                        // Dashes22
	{"¤", "╍", " ", " ", "╍", "╍", "╏", "╏", "╏", " ", "╍", " "},                        // Dashes22Bold
	{"¤", "┄", " ", " ", "┄", "┄", "┆", "┆", "┆", " ", "┄", " "},                        // Dashes33
	{"¤", "┅", " ", " ", "┅", "┅", "┇", "┇", "┇", " ", "┅", " "},                        // Dashes33Bold
	{"¤", "┈", " ", " ", "┈", "┈", "┊", "┊", "┊", " ", "┈", " "},                        // Dashes44
	{"¤", "┉", " ", " ", "┉", "┉", "┋", "┋", "┋", " ", "┉", " "},                        // Dashes44Bold
	{"■", "─", " ", " ", "┬", "─", "├", "│", "└", " ", "─", " "},                        // SingleLine
	{"■", "─", " ", " ", "┬", "─", "├", "│", "╰", " ", "─", " "},                        // SingleLineRounded
	{"■", "═", " ", " ", "╤", "═", "╞", "│", "╘", " ", "═", " "},                        // SingleVDoubleH
	{"■", "─", " ", " ", "╥", "─", "╟", "║", "╙", " ", "─", " "},                        // SingleHDoubleV
	{"■", "═", " ", " ", "╦", "═", "╠", "║", "╚", " ", "═", " "},                        // DoubleLine
	{"■", "━", " ", " ", "┳", "━", "┣", "┃", "┗", " ", "━", " "},                        // Heavy
	{"■", "━", " ", " ", "┯", "━", "┝", "│", "┕", " ", "━", " "},                        // HorizontalHeavy
	{"■", "─", " ", " ", "┰", "─", "┠", "┃", "┖", " ", "─", " "},                        // VerticalHeavy
	{"█", " ", " ", " ", "▂", "▂", "▐", "▐", "▐", " ", "▂", " "},                        // Bold
	{"■ ", "─", " ", " ", "┬┬", "──", "├┼", "││", "└┴", "  ", "─", " "},                 // TwoSingleLines
	{"■ ", "─", " ", " ", "┬┬", "──", "├┼", "││", "╰┴", "  ", "─", " "},                 // TwoSingleLinesRounded
	{"■ ", "─", " ", " ", "╥╥", "──", "╟╫", "║║", "╙╨", "  ", "─", " "},                 // TwoDoubleLines
	{"", "  ", " ", " - ", "▂", "▂", "▐▂", "▐", "▐▂", "|", "▂▂", "..."},                 // ExtraBold1
	{"", "  ", "  ", "  ", "▂", "▂", "▐▂", "▐", "▐▂", " ", "▂▂", "   "},                 // ExtraBold2
	{"", " ", " ", " - ", "▂", "▂", "▙▂", "▌", "▚▂", " ", "▂▂", "   "},                  // ExtraBold3
	{"", " ", " ", " - ", "▄", "▄", "▙▄", "▌", "◥▄▄", " ", "▄▄▄", "   "},                // ExtraBold4
	{"", "  ", " ", " - ", "▂", "▂", "▉▄", "▉", "▉▄", "|", "▄▄", "..."},                 // ExtraBold5
	{"▉ ", " ", " ", " - ", "▗▗", "▗▗", "▚", "▘", "▚", " ", "▗▗▗", "   "},               // BoldDots
	{"▉ ", " ", " ", " - ", " ▄", " ▄", "▀▄", "▄ ", "▀▄", " ", " ▄", "   "},             // ExtraBoldDots
	{"▉ ", "▃▃▃ ", "    ", "     ", "▃▃", "▃▃", "██", "██", " ▀", "  ", "▇▆▄▄", "    "}, // SuperBold
}

func GetTreePattern(style TreeStyle) Branchs {
	if style < 0 || style > TreeStyleLastValue+6 {
		return TreePatternsCollection[TreeStyleAscii1]
	}
	return TreePatternsCollection[style]
}
