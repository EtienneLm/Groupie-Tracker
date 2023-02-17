package main

import (
	"encoding/json"
	"fmt"
	"groupietrackers"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var Cards groupietrackers.Cards
var LocationEx groupietrackers.ExtractLocation
var DatesEx groupietrackers.ExtractDates
var RelationEx groupietrackers.ExtractRelation
var SelectedCard int

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
	http.HandleFunc("/searchName", searchName)
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
		Cards.Array[index].SpotifyId = groupietrackers.Spotify[Cards.Array[index].Id]
		Cards.Array[index].Locations = LocationEx.Index[index].Locations
		Cards.Array[index].ConcertDates = DatesEx.Index[index].Dates
		Cards.Array[index].Relations = RelationEx.Index[index].DatesLocations
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/mainPage.html")) //We link the template and the html file
	tmpl.Execute(w, Cards)
}
func artistPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/artistPage.html")) //change the html
	index, err := strconv.Atoi(r.FormValue("cardButton"))
	if err != nil {
		fmt.Println("Index error in  html value , is not a number")
		MainPage(w, r)
	}
	SelectedCard = index - 1
	ArtistsToDisplay := DataToFunctionnalData(SelectedCard)
	tmpl.Execute(w, ArtistsToDisplay)
}

func searchName(w http.ResponseWriter, r *http.Request) {
	NewDataForInput := groupietrackers.Cards{}
	InputSeachBar := r.FormValue("searchName")
	if InputSeachBar == "" {
		MainPage(w, r)
	}

	for _, value := range Cards.Array {
		if strings.Contains(strings.ToLower(value.Name), strings.ToLower(InputSeachBar)) {
			NewDataForInput.Array = append(NewDataForInput.Array, value)
		}
	}
	tmpl := template.Must(template.ParseFiles("./template/mainPage.html")) //We link the template and the html file
	tmpl.Execute(w, NewDataForInput)
}

func DataToFunctionnalData(IdArstist int) groupietrackers.ArtistsToDisplay {
	/**********We create a new struct to display the data in the html with exploitable data in template**********/
	var ArtistsToDisplay groupietrackers.ArtistsToDisplay
	ArtistsToDisplay.Id = Cards.Array[SelectedCard].Id
	ArtistsToDisplay.Image = Cards.Array[SelectedCard].Image
	ArtistsToDisplay.Name = Cards.Array[SelectedCard].Name
	ArtistsToDisplay.SpotifyId = Cards.Array[SelectedCard].SpotifyId
	for _, value := range Cards.Array[SelectedCard].Members {
		toAppend := new(groupietrackers.Member)
		toAppend.Member = value
		ArtistsToDisplay.Members = append(ArtistsToDisplay.Members, *toAppend)
	}
	var ConcertToAppend []groupietrackers.Concert
	for _, value := range Cards.Array[SelectedCard].Locations {
		toAppend := new(groupietrackers.Concert)
		toAppend.Location = value
		for _, date := range Cards.Array[SelectedCard].Relations[value] {
			toAppenddate := new(groupietrackers.DateConcert)
			toAppenddate.Date = date
			toAppend.Date = append(toAppend.Date, *toAppenddate)
		}
		ConcertToAppend = append(ConcertToAppend, *toAppend)
	}
	ArtistsToDisplay.Concert = ConcertToAppend
	return ArtistsToDisplay
}

// <iframe style="border-radius:12px ;" src="https://open.spotify.com/embed/artist/{{.SpotifyId}}?utm_source=generator&theme=0" width="100%" height="352" frameBorder="0" allowfullscreen="" allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture" loading="lazy"></iframe>
