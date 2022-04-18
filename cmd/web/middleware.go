package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//a middleware that writes to the console when somebody hits a page
//var app config.AppConfig

//Nosurf adds CSRF protection to POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)        //gen
	csrfHandler.SetBaseCookie(http.Cookie{ //cookie generated to identify user when they visit a page or website
		HttpOnly: true,
		Path:     "/",              //cookie path which applies to the entire site
		Secure:   app.InProduction, //the app.InProduction refers to the variable defined in the main package,     production use true but in development set it to false
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

//adding a middleware that tells the webserver to remember a state using sessions
//this func loads and save the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
