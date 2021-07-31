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

	//for _, snippet := range s {
	//	fmt.Fprintf(w, "%v\n", snippet)
	//}

	//create an instance of a templateData struct holding the slice of snippets
	//data := &templateData{Snippets: s}
	//
	//files := []string{
	//	"./ui/html/home.page.html",
	//	"./ui/html/base.layout.html",
	//	"./ui/html/footer.partial.html",
	//}
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	//log.Println(err.Error())
	//
	//	//method
	//	//app.errorLog.Println(err.Error())
	//	//http.Error(w, "Internal Server Error", 500)
	//
	//	//use the serverError() helper
	//	app.serverError(w, err)
	//	return
	//	}

		//err = ts.Execute(w, data)
		//if err != nil {
		//	//log.Println(err.Error())
		//
		//	//app.errorLog.Println(err.Error())
		//	//http.Error(w, "Internal Server Error", 500)
		//
		//	//use the serverError() helper
		//	app.serverError(w, err)
		//	//}
		//}

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

	////create an instance of a templateData struct holding the snippet data
	//data := &templateData{Snippet: s}
	//
	//
	////initialize a slice containing the paths to the show.page.html file
	////plus the base.layout.html and footer.partial.html that we made earlier
	//files := []string{
	//	"./ui/html/show.page.html",
	//	"./ui/html/base.layout.html",
	//	"./ui/html/footer.partial.html",
	//}
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w,err)
	//	return
	//}
	//
	//err = ts.Execute(w, data)
	//if err != nil {
	//	app.serverError(w, err)
	//}

	//fmt.Fprintf(w, "Display a specific snippet with Id %d...", id)

	//fmt.Fprintf(w, "%v", s)
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