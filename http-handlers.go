package main

import (
	"fmt"
	"net/http"
	"strings"
	"html/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func pageEditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, _err := loadPage(title)
	if _err != nil {
		p = &Page{Title: title}
	}
	templatePath := "edit-template.html"
	renderTemplate(w, templatePath, p)
}

func pageSaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	_err := p.save()
	if _err != nil {
		fmt.Fprintf(w, "<h1>500</h1><div>Internal server error</div>")
	} else {
		http.Redirect(w, r, "/view/"+title, http.StatusFound)
	}
}

func pageViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	templatePath := "view-template.html"
	p, _err := loadPage(title)

	if _err != nil {
		title = fmt.Sprintf("Page %s does not exist (yet)", title)
		body := "Use the \"Edit\" button above to create this page"
		p = &Page{Title: title, Body: []byte(body)}
	}
	renderTemplate(w, templatePath, p)
}

func pageErrorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>404</h1><div>Page not found</div>")
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the Wiki!</h1><div>This wiki is the place where amazing things are born</div>")
}

func renderTemplate(w http.ResponseWriter, tmplPath string, p *Page) {
	_err := templates.ExecuteTemplate(w,tmplPath,p)
	if _err != nil {
		fmt.Fprintf(w, "<h1>500</h1><div>Internal server error</div>")
	}
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		mainPageHandler(w, r)
		return
	}
	if strings.Contains(r.URL.Path, "/view/") {
		pageViewHandler(w, r)
		return
	}
	if strings.Contains(r.URL.Path, "/edit/") {
		pageEditHandler(w, r)
		return
	}
	if strings.Contains(r.URL.Path, "/save/") {
		pageSaveHandler(w, r)
		return
	}
	pageErrorHandler(w, r)
}
