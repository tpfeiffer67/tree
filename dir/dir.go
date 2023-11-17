package dir

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tpfeiffer67/tree"
)

// TODO move to external configuration
const (
	DescriptionFileSuffix = "_desc.toml"
)

// key names used in this section
const (
	keyLabel    = "label"
	keyIsDir    = "isDir"
	keyPath     = "path"
	keyFullPath = "fullpath"
	keySize     = "size"
	keySizeSI   = "sizeSI"
	keySizeIEC  = "sizeIEC"
)

// TODO Ajouter config : liste des fichier pouvant traiter du toml, liste de fichiers à exclure, extension description
// TODO Ajouter un ingoreFile regex
func NewDirTree(path string) (node *tree.Node, err error) {
	var fullPath string

	fullPath, err = filepath.Abs(path)
	if err != nil {
		// TODO Find out how to test this condition
		return nil, fmt.Errorf("Unable to create absolute path: %w", err)
	}

	// Verification if the given path exists
	info, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("The specified file does not exist: %w", err)
	}

	// Verification if it is a directory
	if !info.IsDir() {
		return nil, errors.New("Given path is not an existing directory")
	}

	node, err = buildDirTree(fullPath)
	if err != nil {
		return nil, fmt.Errorf("Unable to build dir tree: %w", err) // Should never happen! All the errors in builDirTree end in panic.
	}
	return node, err
}

func foundPath(node *tree.Node, searchedValue interface{}) bool {
	if node.ValueString(keyFullPath, "") == searchedValue.(string) {
		return true
	}
	return false
}

func buildDirTree(fullPath string) (node *tree.Node, err error) {
	shortPathStartIndex := (len(fullPath) - len(filepath.Base(fullPath))) // The short path starts with the last directory name from the full path

	err = filepath.Walk(fullPath, func(path string, info os.FileInfo, err error) error {
		var ignore bool
		// TODO commenter correctement les cas ignorés
		// suppression du dernier élément pour retrouver le chemin du parent
		parentNodePath := path[:len(path)-len(info.Name())-1]
		fileName := info.Name()

		// We ignore the description files (extended by _desc.toml)
		// This file will be processed when
		ignore = strings.HasSuffix(fileName, DescriptionFileSuffix)

		// Ignorer les fichiers .md générés par .odt
		// Hugo-Wrictor
		ignore = ignore || filepath.Ext(fileName) == ".md" && fileExists(changeFileExt(path, ".odt"))

		// TODO mettre en parametre et traiter avec un IIF
		if strings.Index(path, ".git") >= 0 {
			ignore = true
		}

		if !ignore {
			newNode := tree.NewNode()
			if info.IsDir() {
				descFile := path + string(os.PathSeparator) + "_desc.toml"
				if fileExists(descFile) {
					readDescriptionFile(newNode, descFile)
				}
				newNode.SetValue(keyIsDir, "true")
			} else {
				descFile := path + "_desc.toml"
				if fileExists(descFile) {
					readDescriptionFile(newNode, descFile)
				} else {
					if filepath.Ext(fileName) == ".odt" { // Hugo-Wrictor
						mdDescFile := changeFileExt(path, ".md")
						if fileExists(mdDescFile) {
							readDescriptionSection(newNode, mdDescFile, "+++", "+++")
						}
					} else if filepath.Ext(fileName) == ".go" {
						readDescriptionSection(newNode, path, "/*+++", "+++*/")
					}
				}
				newNode.SetValue(keySize, info.Size())
			}
			newNode.SetValue(keyLabel, fileName)
			newNode.SetValue(keyPath, path[shortPathStartIndex:])
			newNode.SetValue(keyFullPath, path)

			// If node == nil, then it is the first time we go through here.
			// So, this means that the newNode is the root directory.
			if node == nil {
				node = newNode
			} else {
				err = node.Attach(newNode, foundPath, parentNodePath)
				// This error should never occur in this context.
				// The searched path must have been added to the tree in a previous iteration.
				if err != nil {
					panic(err)
				}
			}
		}
		return nil
	})

	if err != nil {
		panic(err) // error from filepath.Walk, should never happen.
	}

	return node, err
}

// TODO mettre dans la librairie toolbox
func fileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readDescriptionFile(node *tree.Node, fileName string) {
	description, err := getDescriptionFromTomlFile(fileName)
	// if an error occurs we ignore it. We just don't add any description.
	if err == nil {
		node.SetValues(description)
	}
}

func readDescriptionSection(node *tree.Node, fileName string, sectionDelimiterBegin string, sectionDelimiterEnd string) {
	description, err := getDescriptionFromTomlSectionInFile(fileName, sectionDelimiterBegin, sectionDelimiterEnd)
	// if an error occurs we ignore it. We just don't add any description.
	if err == nil {
		node.SetValues(description)
	}
}

func changeFileExt(fileName string, newExt string) string {
	fileExt := filepath.Ext(fileName)
	return fileName[:len(fileName)-len(fileExt)] + newExt
}

func AddFormatedSize(n *tree.Node) {
	n.RecursiveProcessing(
		func(node *tree.Node) {
			if node.ValueExists(keySize) {
				size := node.ValueInt64(keySize, 0)
				node.SetValue(keySizeIEC, byteCountIEC(size))
				node.SetValue(keySizeSI, byteCountSI(size))
			}
		}, nil)
}

// byteCountSI and byteCountIEC functions has been found here:
// https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/

func byteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func byteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
