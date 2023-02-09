package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Dataform struct {
	Artists   string
	Locations string
	Dates     string
	Relation  string
}

var Data Dataform

type Datatest struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}
type DataArray struct {
	Array []Datatest
}

var Dataarray DataArray
var datatest Datatest
var test2 Datatest

func main() {
	datatest.ID = 1
	datatest.Image = "https://groupietrackers.herokuapp.com/api/images/queen.jpeg"
	datatest.Name = "The Beatles"
	datatest.Members = []string{"John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr"}
	datatest.CreationDate = 1960
	datatest.FirstAlbum = "Please Please Me"
	datatest.Locations = "Liverpool, England"
	datatest.ConcertDates = "1962-1966"
	test2.ID = 1
	test2.Image = "https://groupietrackers.herokuapp.com/api/images/queen.jpeg"
	test2.Name = "oui"
	test2.Members = []string{"John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr"}
	test2.CreationDate = 1960
	test2.FirstAlbum = "Please Please Me"
	test2.Locations = "Liverpool, England"
	test2.ConcertDates = "1962-1966"
	Dataarray.Array = append(Dataarray.Array, datatest)
	Dataarray.Array = append(Dataarray.Array, test2)

	//OpenAPI("https://groupietrackers.herokuapp.com/api")
	fmt.Println(Data.Artists)
	Inisialistion()
}

func Inisialistion() {
	Port := "8080" //We choose port 8080
	fmt.Println("The serveur start on port " + Port + " 🔥")
	fmt.Println("url: http://localhost:8080") //We print this when the server is online
	styles := http.FileServer(http.Dir("template/css"))
	http.Handle("/styles/", http.StripPrefix("/styles", styles)) //We link the css with http.Handle
	http.HandleFunc("/", MainPage)                               //We create the main page , the only function who use a template
	http.ListenAndServe(":"+Port, nil)                           //We start the server
}

func OpenAPI(url string) {
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
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	err = json.Unmarshal(data, &Data)
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/index.html")) //We link the template and the html file
	tmpl.Execute(w, Dataarray)
}
