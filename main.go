package main

import (
	"html/template"
	"log"
	"net/http"
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
	http.NotFound(w,r)
}
func gridHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET url %.25s", r.URL.Path)
	log.Printf("GET params were: %s", r.URL.Query())

	update := r.URL.Query().Get("update")
	if len(update) > 0 {
		http.ServeFile(w,r,"static/sampledata.json")
	} else {
		http.ServeFile(w,r,"static/index.html")
	}
	//q := TemplateData{Box1: fmt.Sprintf("%04d", rand.Intn(10)),
	//	Box2: fmt.Sprintf("%04d", rand.Intn(10)),
	//	Box3: fmt.Sprintf("%04d", rand.Intn(10)),
	//	Box4: fmt.Sprintf("%04d", rand.Intn(10))}
	//tpl.Execute(w, q)
}

func favicon(w http.ResponseWriter, r *http.Request) {
// does nothing cos we dont have an icon yet#
log.Printf("Favicon")
http.ServeFile(w,r,"static/favicon.ico")
	//w.Header().Set("Content-Type", "image/x-icon")
	//w.Header().Set("Cache-Control", "public, max-age=7776000")
	//fmt.Fprintln(w, "data:image/x-icon;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQEAYAAABPYyMiAAAABmJLR0T///////8JWPfcAAAACXBIWXMAAABIAAAASABGyWs+AAAAF0lEQVRIx2NgGAWjYBSMglEwCkbBSAcACBAAAeaR9cIAAAAASUVORK5CYII=\n")
}

func main() {
	port := "3000"

	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/grid/", gridHandler)
	mux.HandleFunc("/favicon.ico", favicon)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}



