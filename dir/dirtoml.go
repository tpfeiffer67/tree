package dir

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/pelletier/go-toml"
)

const errDelimitedSection = "Bad delimitation of the section"

// Here we assume (know) that the given file exists
func getDescriptionFromTomlFile(fileName string) (m map[string]interface{}, err error) {
	tml, err := toml.LoadFile(fileName)
	if err == nil {
		m := tml.ToMap()
		return m, err
	}
	return nil, err
}

// Here we assume (know) that the given file exists
func getDescriptionFromTomlSectionInFile(fileName string, sectionDelimiterBegin string, sectionDelimiterEnd string) (m map[string]interface{}, err error) {

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var section string
	section, err = extractDelimitedSection(string(content), sectionDelimiterBegin, sectionDelimiterEnd)
	if err != nil {
		return nil, err
	}

	tml, err := toml.Load(section)
	if err != nil {
		return nil, err
	}

	m = tml.ToMap()
	return m, err
}

// extractDelimitedSection estracts a delimited section from a text.
// For example, this is used to extract a TOML section from a .md file used for HUGO (https://gohugo.io/templates/files/).
// It's also possible to extract a commented TOML section from a GO source code.
// Only one delimited section is allowed in a file.
func extractDelimitedSection(s string, begin string, end string) (section string, err error) {
	// Serching for the limits of the section
	posBegin := strings.Index(s, begin) + len(begin)
	posEnd := strings.LastIndex(s, end)

	if posBegin > posEnd || posBegin == -1 || posEnd == -1 {
		err = errors.New(errDelimitedSection)
		return "", err
	} else {
		section = s[posBegin:posEnd]
		// If the start or stop marker is still in the extracted section, it's an error
		if strings.Index(section, begin) != -1 || strings.Index(section, end) != -1 {
			err = errors.New(errDelimitedSection)
			return "", err
		}
		return section, nil
	}
}
