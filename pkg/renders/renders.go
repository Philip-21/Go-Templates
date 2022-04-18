package renders

//RenderTemplate renders(gives out) template using html template
//Render template takes a response writer and html takes in the template we want to render
//when i call a page on a website it will call RenderTemplate to render the page,
//the RenderTemplate gets the CreateTemplateCache in Rendertemplate()
//the CreateTemplateCache() gets called .(the first time a template is used it gets to the template cache for retrieval)
//creates the map find all the necessary pages in templates folder,
//loops through the range of the pages and print out the current page then returns the template cache if there are no errors

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Philip-21/temp/pkg/config"
	"github.com/Philip-21/temp/pkg/models"
)

//we can Use our own methods of creating formats  that isnot built into the  golang template language e.g formating dates,returning a current year   ,So we are creating our own functions and parsing it to the template
var functions = template.FuncMap{ //a func map is a map of functions you can use in a template
}

var app *config.AppConfig

//sets the configuration for a new template ,this helps to optimize our template cache which is stored in a map(TemplateCache)
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td

}

//RenderTemplate renders(gives) template using html template
func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {
	var tc map[string]*template.Template
	//tc which stores template cache
	//template cache stores new templates for retrival
	if app.UseCache {
		tc = app.TemplateCache //gets template cache from the appconfig(config dir)
	} else {
		tc, _ = CreateTemplateCache()
	}

	//gets template
	t, ok := tc[html] //if the template exist it will have a value and ok =true
	if !ok {          //if template doesnt exist it will have no value ok =false
		log.Fatal("could not get template from template cache")
	}
	//creating a buffer for a template that is not in the template dir or disk
	buf := new(bytes.Buffer) //puts the parsed template that is currently in memory into some bytes
	_ = t.Execute(buf, td)   //takes the template  executes dont parse any data and store in the buf variable
	_, err := buf.WriteTo(w) //writing the response to the response writer
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

//CreateTemplateCache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{} //using map to hold a particular data structure,creates 2 entries which are in about.page.html & home.page.html
	//getting all the pages in the template folder
	pages, err := filepath.Glob("./templates/*.page.html") //everything that exists in template folder but in the {.page.html} format
	if err != nil {
		return myCache, err //returns eror
	}
	//blank identifiers are used for identfication for a particular purpose without having to refer or return it golang has a feature to refer and use it
	for _, page := range pages { //blank identifiers  to identify the pages
		name := filepath.Base(page) //extracts the name of the page about.page.html & home.page.html using filepath
		//creating a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err //returns error
		}
		//getting layouts formats from the templates
		matches, err := filepath.Glob("./templates/*.layout.html") //base.layout.html
		if err != nil {
			return myCache, err //returns eror
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err //returns eror
			}
		}
		myCache[name] = ts //taking the template set and adding it to the myCache map,adding template to template cache
	}
	return myCache, nil
}
