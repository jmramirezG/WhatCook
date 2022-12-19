/*
WhatCook

Copyright (C) 2022  José María Ramírez González \<jm.ramirez.gonza@gmail.com\>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package WhatCook

import (
	"testing"

	"gotest.tools/v3/assert"

	. "pdftron"
)

func TestPathExists(t *testing.T) {
	doc := Documento{"./nonexistingfile", "replace"}

	doc_exists := Documento{"./documento.go", "replace"}

	assert.Assert(t, !doc.pathExists(), "Path exists when it should not")

	assert.Assert(t, doc_exists.pathExists(), "Path does not exists when it should")
}

func TestGenerateDocument(t *testing.T) {
	doc := Documento{"../testData/testTemplate.docx", `{"name": "test"}`}

	assert.Assert(t, doc.pathExists(), "Path does not exists when it should")

	doc.generateDocument("../testData/testPdf.ignore.pdf")

	generated_pdf := NewPDFDoc("../testData/testPdf.ignore.pdf")
	page := generated_pdf.GetPage(1)

	txt := NewTextExtractor()
	txt.Begin(page) // Read the page

	wordString := ""

	// Extract words one by one.
	word := NewWord()
	line := txt.GetFirstLine()
	for line.IsValid() {
		word = line.GetFirstWord()
		for word.IsValid() {
			wordString = word.GetString()
		}
	}

	assert.Equal(t, wordString, "test", "The text written from the template does not match the text we were supposed to write")
}
