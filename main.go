package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
)

// http://127.0.0.1:3000/?param1=2&param2=apple

type TemplateData struct {
	Box1 string
	Box2 string
	Box3 string
	Box4 string
}

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {

	//fmt.Println("GET params were:", r.URL.Query())
	//v := r.URL.Query().Get("param1")
	//fmt.Println(v)

	q := TemplateData{Box1: fmt.Sprintf("%04d", rand.Intn(10)),
		Box2: fmt.Sprintf("%04d", rand.Intn(10)),
		Box3: fmt.Sprintf("%04d", rand.Intn(10)),
		Box4: fmt.Sprintf("%04d", rand.Intn(10))}

	tpl.Execute(w, q)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)

}



