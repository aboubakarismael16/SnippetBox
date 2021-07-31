package main

import (
	"SnippetBox/pkg/forms"
	"SnippetBox/pkg/models"
	"fmt"
	"net/http"
	"strconv"
)


func (app *application) home(w http.ResponseWriter, r * http.Request) {
	//if r.URL.Path != "/" {
	//	//http.NotFound(w, r)
	//
	//	//use the notFound() helper
	//	app.notFound(w)
	//	return
	//}

	//w.Write([]byte("Hello from SnippetBox"))

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}




	//use the new render helper
	app.render(w, r, "home.page.html", &templateData{
		Snippets: s,
	})

}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		//http.NotFound(w, r)

		//use the notFound() helper
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	//use the new render helper
	app.render(w, r, "show.page.html", &templateData{
		Snippet: s,
	})


}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request)  {
	app.render(w, r, "create.page.html", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(w, r, "create.page.html", &templateData{
			Form: form,
		})
		return
	}

	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}