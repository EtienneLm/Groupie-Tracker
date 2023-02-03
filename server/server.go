package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	println("Hello, World!")
	Inisialistion()
}

func Inisialistion() {
	Port := "8080"                                          //We choose port 8080
	fmt.Println("The serveur start on port " + Port + " 🔥") //We print this when the server is online
	styles := http.FileServer(http.Dir("template/css"))
	http.Handle("/styles/", http.StripPrefix("/styles", styles)) //We link the css with http.Handle
	http.HandleFunc("/", MainPage)                               //We create the main page , the only function who use a template
	http.ListenAndServe(":"+Port, nil)                           //We start the server
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/index.html")) //We link the template and the html file
	tmpl.Execute(w)                                                     //We execute the template and put the Data in
}
