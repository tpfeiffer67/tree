// The indentedtext package provides two functions for importing tree structures from either a string or a file.
// These tree structures are represented in an indented text format, with each line containing either the label of a node or
// the description associated with that node. The description can be either plain text or formatted in TOML,
// and will be stored in the node values with either the key "description" or the keys specified by the TOML fields, respectively.
// The hierarchy of the tree is determined by the indentation of the text lines.
// The buildTree function is responsible for constructing the tree structure from an indented text string,
// while the Import and ImportFromFile functions handle the input and call buildTree to create the tree.

package indentedtext

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/tpfeiffer67/tree"

	"gopkg.in/karalabe/cookiejar.v2/collections/stack"
)

// charIndentation is the character used to denote indentation levels in the indented text.
const charIndentation = '\t'

// charData is the character used to denote a line containing data in the indented text.
const charData = ' '

// charIndentation and charData are ascii chars (one byte in utf8).
// This way it is not necessary to process the string as runes. Processing the bytes is sufficient.

// processLine parses a line of indented text and returns the indentation level,
// whether the line contains data, and the content of the line.
func processLine(s string) (deepness int, isData bool, text string) {
	deepness = 0
	isData = false
	text = ""

	// The deepness of a node is calculated by
	// the number of charIndentation characters
	// present at the left of the line.
	for i := 0; i < len(s); i++ {
		if s[i] == charIndentation {
			deepness++
		} else {
			break
		}
	}

	// If, after the charIndentation characters,
	// the line starts with a charData character,
	// it is a line containing data for the node.
	if len(s) > deepness {
		if s[deepness] == charData {
			isData = true
			text = s[deepness+1:]
		} else {
			text = s[deepness:]
		}
	}

	return
}

// buildTree constructs a tree structure from an indented text string.
func buildTree(indentedText string) (root *tree.Node, err error) {
	var line string
	var lineIndex int
	var s string
	var deepness int
	var previousDeepness int
	var isData bool
	var aNode *tree.Node

	// create a new stack to hold the nodes
	lifo := stack.New()
	scanner := bufio.NewScanner(strings.NewReader(indentedText))
	for scanner.Scan() {
		line = scanner.Text()
		deepness, isData, s = processLine(line)

		// The first line must be the root node (no indentation)
		if lineIndex == 0 && deepness != 0 {
			lifo.Reset()
			return nil, errors.New("First line must be the root node")
		}

		if deepness > previousDeepness+1 {
			return nil, errors.New("Deepness can only increment by one")
		}
		previousDeepness = deepness

		// if the line contains data, it is added to data field of the current node
		if isData {
			addData(lifo.Top().(*tree.Node), s)
		} else {
			// only one root node is allowed
			aNode = tree.NewNode()
			aNode.SetValue(KEY_LABEL, s)
			if deepness == 0 {
				if lineIndex > 0 {
					lifo.Reset()
					return nil, errors.New("Tree can have only one root")
				}
				root = aNode
			} else {
				if deepness >= lifo.Size() {
					lifo.Top().(*tree.Node).Adopt(aNode)
				} else {
					lifo.Pop()
					for lifo.Size() > deepness {
						lifo.Pop()
					}
					lifo.Top().(*tree.Node).Adopt(aNode)
				}
			}
			lifo.Push(aNode)
		}
		// We count lines just to know if it is the first line or not
		lineIndex++
	}
	return
}

func addData(node *tree.Node, s string) {
	const lineBreak = "\n"
	if node.Object != nil {
		str := node.Object.(string)
		node.Object = str + lineBreak + s
	} else {
		node.Object = s
	}
}

func importValuesFromToml(n *tree.Node) {
	n.RecursiveProcessing(
		func(node *tree.Node) {
			// The data imported from the indented file are
			// stored in the field "Object".
			if node.Object != nil {
				data := node.Object.(string)
				// The data can be plain text or TOML
				tml, err := toml.Load(data)
				// If no error has been reported by unmarshaling the TOML,
				// the unmarshaled fields can be stored as values
				if err == nil {
					m := tml.ToMap()
					for key, value := range m {
						node.SetValue(key, value)
					}
				} else { // if unmarshaling failed, then the data is considered as plain text and copied in the description field
					node.SetValue(KEY_DESCRIPTION, data)
				}
				// The temporary data can be deleted
				node.Object = nil
			}
		}, nil)
}

func Import(indentedText string) (*tree.Node, error) {
	root, err := buildTree(indentedText)
	if err != nil {
		return nil, fmt.Errorf("Import: %w", err)
	}
	importValuesFromToml(root)
	return root, err
}

func ReadFileIntoString(fileName string) (string, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("Read string from file: %w", err)
	}
	return string(b), err
}

func ImportFromFile(fileName string) (*tree.Node, error) {
	indentedText, err := ReadFileIntoString(fileName)
	if err != nil {
		return nil, fmt.Errorf("Import from file: %w", err)
	}
	return Import(indentedText)
}
