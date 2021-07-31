package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	//dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()
	//mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	//mux.Get("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	//mux.Post("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet))
	//mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static",fileServer))

	//return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	return standardMiddleware.Then(mux)
}