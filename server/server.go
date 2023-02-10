package main

import (
	"encoding/json"
	"fmt"
	"groupietrackers"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

var Cards groupietrackers.Cards
var LocationEx groupietrackers.ExtractLocation
var DatesEx groupietrackers.ExtractDates
var RelationEx groupietrackers.ExtractRelation

func main() {
	InitAPI()
	Inisialistion()
}

func Inisialistion() {
	Port := "8080"                                          //We choose port 8080
	fmt.Println("The serveur start on port " + Port + " 🔥") //We print this when the server is online
	fmt.Println("http://localhost:8080/")
	styles := http.FileServer(http.Dir("template/css"))
	http.Handle("/styles/", http.StripPrefix("/styles", styles)) //We link the css with http.Handle
	http.HandleFunc("/", MainPage)                               //We create the main page , the only function who use a template
	http.HandleFunc("/artistPage", artistPage)
	http.ListenAndServe(":"+Port, nil) //We start the server

}

func APICall(url string) (data []byte) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer res.Body.Close()
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return
}

func InitAPI() {
	//We call artists from API
	data := APICall("https://groupietrackers.herokuapp.com/api/artists")
	json.Unmarshal(data, &Cards.Array)
	data = APICall("https://groupietrackers.herokuapp.com/api/locations")
	json.Unmarshal(data, &LocationEx)
	data = APICall("https://groupietrackers.herokuapp.com/api/dates")
	json.Unmarshal(data, &DatesEx)
	data = APICall("https://groupietrackers.herokuapp.com/api/relation")
	json.Unmarshal(data, &RelationEx)
	for index, _ := range Cards.Array {
		Cards.Array[index].Locations = LocationEx.Index[index]
		Cards.Array[index].ConcertDates = DatesEx.Index[index]
		Cards.Array[index].Relations = RelationEx.Index[index]
	}

}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/mainPage.html")) //We link the template and the html file
	tmpl.Execute(w, Cards)
}
func artistPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/artistPage.html")) //change the html
	fmt.Println(r.FormValue("cardButton"))
	tmpl.Execute(w, nil)
}
