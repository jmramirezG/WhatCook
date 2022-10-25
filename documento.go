package WhatCook

import (
	"fmt"
	. "pdftron"
)

import  "pdftron/Samples/LicenseKey/GO"

type Documento struct {
	path      []string
	reemplazo []string
}

func (d Documento) generaDocumento(pathGuardado []string) ([]string) {
	// Create a TemplateDocument object from an input office file.
	templateDoc := ConvertCreateOfficeTemplate(d.path, NewOfficeToPDFOptions());

	// Fill the template with data from a JSON string, producing a PDF document.
	pdfDoc := templateDoc.FillTemplateJson(d.reemplazo);

	// Save the PDF to a file.
	pdfDoc.Save(pathGuardado, uint(SDFDocE_linearized))
}
