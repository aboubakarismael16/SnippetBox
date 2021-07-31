package main

import (
	"SnippetBox/pkg/models/mysql"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	//session *sessions2.Session
	snippets *mysql.SnippetModel // use the SnippetModel available in pkg/models
	templateCache map[string]*template.Template
}

func openDB(dsn string) (*sql.DB, error)  {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn","root:13628@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	//secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret")
	//flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	//Initialize a new template cache...
	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	//session := sessions2.New([]byte(*secret))
	//session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		//session: session,
		snippets: &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}


	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),// Call the new app.routes() method
	}

	//log.Printf("Starting Sever on %s", *addr)
	infoLog.Printf("Starting Sever on %s", *addr)
	err = srv.ListenAndServe()
	//log.Fatal(err)
	errorLog.Fatal(err)
}
