package handlers

import (
	"net/http"

	"github.com/Philip-21/temp/pkg/config"
	"github.com/Philip-21/temp/pkg/models"
	"github.com/Philip-21/temp/pkg/renders"
)

//always start a go function with a block letter so that it can be easily imported into anther directory e.g 	renders.RenderTemplate(w, "home.page.html")
//handlers create response and receives request for the clients

//Repository helps to swap contents of our application with a minimal changes requiredin the code base
//Repo is the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHadlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) { // (m*Repository)
	//grab the remote ip address of the visitor and store it in a home page
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	renders.RenderTemplate(w, "home.page.html", &models.TemplateData{}) //parsing the response for the homepage

}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//performs some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello,again"

	//pulling the value out of the session
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip") //gets the string from the context and the key to look it up which is "remote_ip"
	stringMap["remote_ip"] = remoteIP

	//sends the data to the template
	renders.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
