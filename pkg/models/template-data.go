package models

// a template for holding diffrent kinds of  data sent from  handlers to template,this will give a reponse on the page
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} //interface{} implements data we arent sure of the datatype that will be parsed/declared
	CSRFToken string                 //a security token that handles forms
	Flash     string
	Warning   string
	Error     string
}
