package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	//Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin.
	if name := r.FormValue("name"); name != "" {
		welcome.Name = name
	}
	//If errors show an internal server error message
	//I also pass the welcome struct to the welcome-template.html file.
	if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexHandler)

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8090", nil))
}
