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
	"sync"
)

var Cards groupietrackers.Cards
var LocationEx groupietrackers.ExtractLocation
var DatesEx groupietrackers.ExtractDates
var RelationEx groupietrackers.ExtractRelation
var SelectedCard int
var CardsPagination []groupietrackers.Cards
var SortedCardsPagination []groupietrackers.Cards

func main() {

	var wg sync.WaitGroup //We use this value for the invisilble api calls
	//We call artists from API
	wg.Add(1)               //We create a secondary chanel
	go FastServerStart(&wg) //We run AddNewWord in this chanel because the function is slow
	Inisialistion()
	wg.Wait() //We stop the chanel
}

func IntoMultiplePages(NumberOfCards int, Entry []groupietrackers.Artists, toTurnNegative int) []groupietrackers.Cards {
	var CardPagiantion []groupietrackers.Cards
	var TmpCardsArray groupietrackers.Cards
	TmpCardsArray.NotLastPage = true
	var TmpIndex int
	NbrPage := 0
	if toTurnNegative == -1 {
		NbrPage = 1
	}
	for index := range Entry {
		TmpIndex++
		TmpCardsArray.Array = append(TmpCardsArray.Array, Entry[index])
		if TmpIndex == NumberOfCards {
			TmpIndex = 0
			TmpCardsArray.PreviousPage = (NbrPage - 1) * toTurnNegative
			TmpCardsArray.NexPage = (NbrPage + 1) * toTurnNegative
			if toTurnNegative == -1 {
				TmpCardsArray.IdPage = NbrPage
			} else {
				TmpCardsArray.IdPage = NbrPage + 1
			}
			CardPagiantion = append(CardPagiantion, TmpCardsArray)
			TmpCardsArray.Array = nil
			TmpCardsArray.NotFirstPage = true
			NbrPage++
		}
	}
	TmpCardsArray.NotLastPage = false
	TmpCardsArray.PreviousPage = (NbrPage - 1) * toTurnNegative
	TmpCardsArray.IdPage = NbrPage + 1
	TmpCardsArray.NexPage = (NbrPage + 1) * toTurnNegative
	CardPagiantion = append(CardPagiantion, TmpCardsArray)
	return CardPagiantion
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
	http.HandleFunc("/changePage", ChangePage)
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

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/mainPage.html")) //We link the template and the html file
	tmpl.Execute(w, CardsPagination[0])
}
func ChangePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/mainPage.html")) //We link the template and the html file
	ToPageNbr, _ := strconv.Atoi(r.FormValue("topage"))
	if ToPageNbr < 0 {
		tmpl.Execute(w, SortedCardsPagination[(ToPageNbr*-1)-1])
	} else {
		tmpl.Execute(w, CardsPagination[ToPageNbr])
	}

}

func artistPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/artistPage.html")) //change the html
	index, err := strconv.Atoi(r.FormValue("cardButton"))
	if err != nil {
		MainPage(w, r)
	} else {
		SelectedCard = index - 1
		ArtistsToDisplay := DataToFunctionnalData(SelectedCard)
		tmpl.Execute(w, ArtistsToDisplay)
	}

}

func searchName(w http.ResponseWriter, r *http.Request) {
	NewDataForInput := groupietrackers.Cards{}
	InputSeachBar := r.FormValue("searchName")
	if InputSeachBar == "" {
		MainPage(w, r)
	} else {
		for _, value := range Cards.Array {
			if strings.Contains(strings.ToLower(value.Name), strings.ToLower(InputSeachBar)) {
				NewDataForInput.Array = append(NewDataForInput.Array, value)
			}
		}
		SortedCardsPagination = IntoMultiplePages(5, NewDataForInput.Array, -1)
		tmpl := template.Must(template.ParseFiles("./template/mainPage.html")) //We link the template and the html file
		tmpl.Execute(w, SortedCardsPagination[0])
	}
}

func DataToFunctionnalData(IdArstist int) groupietrackers.ArtistsToDisplay {
	/**********We create a new struct to display the data in the html with exploitable data in template**********/
	var ArtistsToDisplay groupietrackers.ArtistsToDisplay
	ArtistsToDisplay.Id = Cards.Array[SelectedCard].Id
	ArtistsToDisplay.Image = Cards.Array[SelectedCard].Image
	ArtistsToDisplay.Name = Cards.Array[SelectedCard].Name
	ArtistsToDisplay.SpotifyId = Cards.Array[SelectedCard].SpotifyId
	ArtistsToDisplay.CreationDate = Cards.Array[SelectedCard].CreationDate
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

func FastServerStart(wg *sync.WaitGroup) { // We enter the DB and the word to add for add the word into the target DB
	defer wg.Done() //We use defer for close wg in the end of the function
	data := APICall("https://groupietrackers.herokuapp.com/api/artists")
	json.Unmarshal(data, &Cards.Array)
	data = APICall("https://groupietrackers.herokuapp.com/api/locations")
	json.Unmarshal(data, &LocationEx)
	data = APICall("https://groupietrackers.herokuapp.com/api/dates")
	json.Unmarshal(data, &DatesEx)
	data = APICall("https://groupietrackers.herokuapp.com/api/relation")
	json.Unmarshal(data, &RelationEx)
	for index := range Cards.Array {
		Cards.Array[index].SpotifyId = groupietrackers.Spotify[Cards.Array[index].Id]
		Cards.Array[index].Locations = LocationEx.Index[index].Locations
		Cards.Array[index].ConcertDates = DatesEx.Index[index].Dates
		Cards.Array[index].Relations = RelationEx.Index[index].DatesLocations
	}

	var TmpValueForCards = Cards.Array
	CardsPagination = IntoMultiplePages(10, TmpValueForCards, 1)
}
