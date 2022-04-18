package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Philip-21/temp/pkg/config"
	"github.com/Philip-21/temp/pkg/handlers"
	"github.com/Philip-21/temp/pkg/renders"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

//the main application function
func main() {

	app.InProduction = false //change to true when in production,this is just a develpment phase
	session = scs.New()
	session.Lifetime = 24 * time.Hour              //session lasts for 24hours
	session.Cookie.Persist = true                  //storing sessions in cookies,cookie persists after the browser or window is closed
	session.Cookie.SameSite = http.SameSiteLaxMode //specifying the site where the cookie applies to
	session.Cookie.Secure = app.InProduction       // insists the cookie is being crypted and the connection is from https //production use true but in development set it to false
	app.Session = session                          //defined for handlers to have access to

	tc, err := renders.CreateTemplateCache() //new templates which are stored in the createtemplatecache are defined as tc
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc //defining the app as tc which stores template cache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	renders.NewTemplates(&app) //pointing to a *config.AppConfig in the render dir

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil) // //ListenAndServe starts an HTTP server with a given address or port num and handler and connects to the web page

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app), //handler routes
	}
	err = srv.ListenAndServe()
	log.Fatal(err) //if theres an error it logs fatal

}
