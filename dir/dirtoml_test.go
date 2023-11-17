package dir

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractDelimitedSection(t *testing.T) {
	section, err := extractDelimitedSection("bla bla +++hello+++ bla blabla bla", "+++", "+++")
	assert.Equal(t, "hello", section)
	assert.Nil(t, err)

	section, err = extractDelimitedSection("bla bla +++hello+++ bla blabla+++ bla", "+++", "+++")
	assert.Equal(t, "Bad delimitation of the section", err.Error())

	section, err = extractDelimitedSection("bla bla +++hello bla blabla bla", "+++", "+++")
	assert.Equal(t, "Bad delimitation of the section", err.Error())

	section, err = extractDelimitedSection("bla bla hello bla blabla bla+++", "+++", "+++")
	assert.Equal(t, "Bad delimitation of the section", err.Error())

	section, err = extractDelimitedSection("bla bla hello bla blabla bla", "+++", "+++")
	assert.Equal(t, "Bad delimitation of the section", err.Error())

	section, err = extractDelimitedSection("bla bla /*+++hello+++*/ bla blabla+++ bla", "/*+++", "+++*/")
	assert.Equal(t, "hello", section)
	assert.Nil(t, err)

	section, err = extractDelimitedSection("++++++", "+++", "+++")
	assert.Equal(t, "", section)
	assert.Nil(t, err)

	section, err = extractDelimitedSection("+++++", "+++", "+++")
	assert.Equal(t, "Bad delimitation of the section", err.Error())
}
