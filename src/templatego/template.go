package templatego

import (
	"html/template"
	"log"
)

var TemplateMap map[string]*template.Template
var base = "./frontend/"

func init() {

	TemplateMap=make(map[string]*template.Template)
	myTemplates := make(map[string]string)
	myTemplates["questions"] = base + "question.html"
	myTemplates["score"] = base + "score.html"


	for k, v := range myTemplates {
		//t := template.Must(template.New("question").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`))
		t := template.Must(template.ParseFiles(v))
		TemplateMap[k] = t

	}

	//
	//for k, v := range myTemplates {
	//	t := template.Must(template.New(k).ParseFiles(v))
	//	TemplateMap[k] = t
	//}

	log.Print("Templated init complete")
}





// templateFile defines the contents of a template to be stored in a file, for testing.
type templateFile struct {
	name     string
	contents string
}
