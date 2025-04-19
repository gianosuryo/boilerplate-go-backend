package helper

import (
	"bytes"
	"html/template"
	"path"
)

func RenderHtmlToString(htmlFilePath []string, data any) string {
	// Get body from html template file
	var filepath = path.Join(htmlFilePath...)
	var tmpl, err = template.ParseFiles(filepath)
	PanicIfError(err)

	var outputTemplate bytes.Buffer
	err = tmpl.Execute(&outputTemplate, data)
	PanicIfError(err)

	return outputTemplate.String()
}
