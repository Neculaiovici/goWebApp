package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// first method to load static files from directory
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "public"+r.URL.Path)
	// })

	// second method
	// fmt.Println("Server start")
	// http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))

	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestFile := r.URL.Path[1:]
		t := templates.Lookup(requestFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
}

//Template loading function
func populateTemplates() *template.Template {
	result := template.New("template")
	const basePath = "template"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
