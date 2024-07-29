package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"web-golang/src/main/mrizkisaputra"
)

type M map[string]any

//go:embed templates/*.gohtml
var TEMPLATES embed.FS

//go:embed assets/css/*.css
var STATIC_CSS embed.FS
var tmpl = template.Must(template.ParseFS(TEMPLATES, "templates/*.gohtml"))
var serveMux = http.NewServeMux()

func homeHandler(w http.ResponseWriter, request *http.Request) {
	//data := M{
	//	"title": "Halaman Beranda",
	//	"body":  "ini adalah halaman beranda home.html",
	//}
	user := mrizkisaputra.User{
		Username: "kiki", Password: "rahasia",
	}
	renderTemplate(w, "home", user)
}

func aboutHandler(w http.ResponseWriter, request *http.Request) {
	data := M{
		"title": "Halaman About",
		"body":  "ini adalah halaman About about.html",
	}
	renderTemplate(w, "about", data)
}

func renderTemplate(writer http.ResponseWriter, templateName string, data interface{}) {
	err := tmpl.ExecuteTemplate(writer, templateName, data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func staticCssFileHandler() http.Handler {
	sub, err := fs.Sub(STATIC_CSS, "assets/css")
	if err != nil {
		http.Error(nil, err.Error(), http.StatusInternalServerError)
	}
	return http.FileServer(http.FS(sub))
}

func crossSiteScripting(w http.ResponseWriter, request *http.Request) {
	query := request.URL.Query().Get("test")
	if query != "" {

		tmpl.ExecuteTemplate(w, "home",
			map[string]interface{}{"query": template.HTML(query)},
		)
	}
}

var funcs = template.FuncMap{
	"avg": func(n ...int) int {
		var result = 0
		for _, e := range n {
			result += e
		}
		return result / len(n)
	},
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/home", http.StatusTemporaryRedirect)
	//template.Must(template.ParseFiles("index.gohtml")).Funcs(funcs)
	//tmpl.Execute(writer, nil)
}

func main() {
	serveMux.HandleFunc("/XSS", crossSiteScripting)
	serveMux.HandleFunc("/", indexHandler)
	serveMux.HandleFunc("/home", homeHandler)
	serveMux.HandleFunc("/about", aboutHandler)
	serveMux.Handle("/static/css/", http.StripPrefix("/static/css/", staticCssFileHandler()))
	var server = http.Server{
		Addr:    "localhost:9000",
		Handler: serveMux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

////go:embed templates/*.gohtml
//var TEMPLATES embed.FS
////go:embed assets/css
//var STATIC_CSS embed.FS
//func main() {
//	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
//		tmpl, err := template.ParseFS(TEMPLATES, "templates/index.gohtml")
//		if err != nil {
//			http.Error(writer, err.Error(), http.StatusInternalServerError)
//		}
//		var data map[string]any = map[string]any{
//			"title": "Belajar pemograman web",
//			"name":  "John Doe",
//		}
//		err = tmpl.Execute(writer, data)
//		if err != nil {
//			http.Error(writer, err.Error(), http.StatusInternalServerError)
//		}
//	})
//
//	sub, _ := fs.Sub(STATIC_CSS, "assets/css")
//	fileServer := http.FileServer(http.FS(sub))
//	serveMux.Handle("/static/css/", http.StripPrefix("/static/css", fileServer))
//
//	var server = http.Server{
//		Addr:    "localhost:9000",
//		Handler: serveMux,
//	}
//	server.ListenAndServe()
//}
