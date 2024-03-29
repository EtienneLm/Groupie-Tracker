package main

import (
	"fmt"
	"groupietrackers"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var Cards groupietrackers.Cards //* All the data

var CardsPagination []groupietrackers.Cards       //* Use in pagination in main page
var SortedCardsPagination []groupietrackers.Cards //* Use in pagination when search bar used

var SelectedCard int       //* Use to communicate the chosen cards id , we can't use navigate like in JSX
var NumberOfCards int = 10 //* The number of cards in page the server start

var ArtistsToDisplay groupietrackers.ArtistsToDisplay //*

var SpotifyToken string //* Use to stock the Token and note spam the Spotify API

var AdminMail string = "admin"
var AdminPassword string = "admin"

func main() {
	/*
	* We call the APICall to extract artists for the main page
	 */
	var wg sync.WaitGroup

	groupietrackers.APICall("https://groupietrackers.herokuapp.com/api/artists", &Cards.Array)
	wg.Add(1) //*We create a secondary chanel
	go FastServerStart()
	Inisialistion()
	wg.Wait()
}

func Inisialistion() {
	/*
	* We inialise the server and all roads
	* We anisialise the css and display when the server is online
	 */
	styles := http.FileServer(http.Dir("template/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", styles)) //We link the css with http.Handle
	http.HandleFunc("/", MainPage)                               //We create the main page , the only function who use a template
	http.HandleFunc("/artistPage", artistPage)
	http.HandleFunc("/searchName", searchName)
	http.HandleFunc("/concert", concertPage)
	http.HandleFunc("/contactUs", contactUsPage)
	http.HandleFunc("/changePage", ChangePage)
	http.HandleFunc("/adminLog", AdminLog)
	http.HandleFunc("/adminpage", Adminpage)
	http.HandleFunc("/NbrInPageChange", NbrInPageChange)
	http.HandleFunc("/changeMap", MapUpDate)
	http.HandleFunc("/reloadAPI", ReloadAPI)
	http.HandleFunc("/sorting", SortingList)

	Port := "8080"                                          //We choose port 8080
	fmt.Println("The serveur start on port " + Port + " 🔥") //We print this when the server is online
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":"+Port, nil) //We start the server
}

func ReloadAPI(w http.ResponseWriter, r *http.Request) {
	/*
	* Function who reload information from the api
	 */
	groupietrackers.APICall("https://groupietrackers.herokuapp.com/api/artists", &Cards.Array)
	FastServerStart()
	Adminpage(w, r)
}

func NbrInPageChange(w http.ResponseWriter, r *http.Request) {
	/*
	* Function who change the number of cards in the main page
	 */
	NewNumberOfCards, err := strconv.Atoi(r.FormValue("mail"))
	if err == nil {
		NumberOfCards = NewNumberOfCards
		var TmpValueForCards = Cards.Array
		CardsPagination = groupietrackers.IntoMultiplePages(&NumberOfCards, TmpValueForCards, 1, &Cards.ForReacherchBar)
		MainPage(w, r)
	} else {
		MainPage(w, r)
	}
}

func MapUpDate(w http.ResponseWriter, r *http.Request) {
	if ArtistsToDisplay.Concert != nil {
		tmpl := template.Must(template.ParseFiles("./template/artistPage.html")) //change the html
		IdMap, _ := strconv.Atoi(r.FormValue("ChangeMap"))
		Output := groupietrackers.Map(ArtistsToDisplay.Concert[IdMap].Location)
		ArtistsToDisplay.X, ArtistsToDisplay.Y = Output[0], Output[1]
		tmpl.Execute(w, ArtistsToDisplay)
	} else {
		MainPage(w, r)
	}

}
func AdminLog(w http.ResponseWriter, r *http.Request) {
	/*
	* Function of redirection to the admin page
	 */
	tmpl := template.Must(template.ParseFiles("./template/AdminLog.html")) //We link the template and the html file
	groupietrackers.Admin.IsConnected = false
	groupietrackers.Admin.IsBadInput = false
	tmpl.Execute(w, groupietrackers.Admin)
}

func Adminpage(w http.ResponseWriter, r *http.Request) {
	/*
	* Function who check if the admin is connected and who check the password and the mail
	 */
	tmpl := template.Must(template.ParseFiles("./template/AdminLog.html")) //We link the template and the html file
	if AdminMail == r.FormValue("mail") && AdminPassword == r.FormValue("psw") || groupietrackers.Admin.IsConnected == true {
		groupietrackers.Admin.IsConnected = true
		groupietrackers.Admin.IsBadInput = false
	} else {
		groupietrackers.Admin.IsConnected = false
		groupietrackers.Admin.IsBadInput = true
	}
	tmpl.Execute(w, groupietrackers.Admin)
}

func MainPage(w http.ResponseWriter, r *http.Request) {

	/*
	* Function who redirecte to the main page
	 */
	tmpl := template.Must(template.ParseFiles("./template/mainPage.html")) //We link the template and the html file
	tmpl.Execute(w, CardsPagination[0])
}

func ChangePage(w http.ResponseWriter, r *http.Request) {
	/*
	* Function who change the page on the main page ( Little arrow on the bottom of the page )
	 */
	tmpl := template.Must(template.ParseFiles("./template/mainPage.html")) //We link the template and the html file
	ToPageNbr, _ := strconv.Atoi(r.FormValue("topage"))
	if ToPageNbr < 0 {

		tmpl.Execute(w, SortedCardsPagination[(ToPageNbr*-1)-1])
	} else {
		tmpl.Execute(w, CardsPagination[ToPageNbr])
	}

}

func artistPage(w http.ResponseWriter, r *http.Request) {
	/*
	* We redirecte to the artist page and we send the data of the artist to the html
	 */
	tmpl := template.Must(template.ParseFiles("./template/artistPage.html")) //change the html
	index, err := strconv.Atoi(r.FormValue("cardButton"))
	if err != nil {
		MainPage(w, r)
	} else {
		SelectedCard = index - 1
		ArtistsToDisplay = DataToFunctionnalData(SelectedCard)
		Output := groupietrackers.Map(ArtistsToDisplay.Concert[0].Location)
		ArtistsToDisplay.X, ArtistsToDisplay.Y = Output[0], Output[1]
		fmt.Println(ArtistsToDisplay.Name, "loaded")
		tmpl.Execute(w, ArtistsToDisplay)
	}
}

func searchName(w http.ResponseWriter, r *http.Request) {
	/*
	* Function who search the name and display the result on the main page
	 */
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
		SortedCardsPagination = groupietrackers.IntoMultiplePages(&NumberOfCards, NewDataForInput.Array, -1, &Cards.ForReacherchBar)
		tmpl := template.Must(template.ParseFiles("./template/mainPage.html")) //We link the template and the html file
		tmpl.Execute(w, SortedCardsPagination[0])
	}
}

func concertPage(w http.ResponseWriter, r *http.Request) {
	/*
	* Redirecte to the concert page
	 */
	AllConcertsInit()
	tmpl := template.Must(template.ParseFiles("./template/concertPage.html"))
	tmpl.Execute(w, Cards)
}

func contactUsPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/contactUsPage.html"))
	tmpl.Execute(w, r)
}

func FastServerStart() {
	/*
	* We call the function to get the data from the api in a secondary channel for run the server faster
	* We open SB/Spotify.txt to get the spotify id of the artist
	 */

	groupietrackers.APICall("https://groupietrackers.herokuapp.com/api/locations", &groupietrackers.LocationEx)
	groupietrackers.APICall("https://groupietrackers.herokuapp.com/api/dates", &groupietrackers.DatesEx)
	groupietrackers.APICall("https://groupietrackers.herokuapp.com/api/relation", &groupietrackers.RelationEx)

	for index := range Cards.Array {
		Cards.Array[index].Locations = groupietrackers.LocationEx.Index[index].Locations
		Cards.Array[index].ConcertDates = groupietrackers.DatesEx.Index[index].Dates
		Cards.Array[index].Relations = groupietrackers.RelationEx.Index[index].DatesLocations
	}
	SortByAlphabetical(Cards.Array)
	var TmpValueForCards = Cards.Array
	Cards.ForReacherchBar = Cards.Array
	CardsPagination = groupietrackers.IntoMultiplePages(&NumberOfCards, TmpValueForCards, 1, &Cards.ForReacherchBar)
	fmt.Println("loading API 100%")
}

func DataToFunctionnalData(IdArstist int) groupietrackers.ArtistsToDisplay {
	/*
	* This function is called when we clic on a artist card , we make a new struct with
	* the data of the artist for golang template because we can't use the struct of the json file
	 */
	//We find the id of the artist in the Cards.Array array
	for index, value := range Cards.Array {
		if value.Id == IdArstist+1 {
			IdArstist = index
			break
		}
	}
	//* We fix the name of the artist for the spotify api
	switch Cards.Array[IdArstist].Name {
	case "ACDC":
		groupietrackers.GetArtist("AC/DC", &SpotifyToken)
	case "The Jimi Hendrix Experience":
		groupietrackers.GetArtist("Jimi Hendrix", &SpotifyToken)
	case "NWA":
		groupietrackers.GetArtist("N.W.A.", &SpotifyToken)
	case "Bobby McFerrins":
		groupietrackers.GetArtist("Bobby McFerrin", &SpotifyToken)
	default:
		groupietrackers.GetArtist(Cards.Array[IdArstist].Name, &SpotifyToken)
	}
	//* We call spotify api

	var ArtistsToDisplay groupietrackers.ArtistsToDisplay
	ArtistsToDisplay.Id = Cards.Array[IdArstist].Id
	ArtistsToDisplay.Image = Cards.Array[IdArstist].Image
	ArtistsToDisplay.Name = Cards.Array[IdArstist].Name
	ArtistsToDisplay.CreationDate = Cards.Array[IdArstist].CreationDate
	ArtistsToDisplay.SpotifyId = groupietrackers.SpotifyInfo.Artists.Items[0].Id
	ArtistsToDisplay.Genre = groupietrackers.SpotifyInfo.Artists.Items[0].Genres[0]
	ArtistsToDisplay.Followers = groupietrackers.SpotifyInfo.Artists.Items[0].Followers.Total

	//* We make a new struct for the members
	for _, value := range Cards.Array[IdArstist].Members {
		toAppend := new(groupietrackers.Member)
		toAppend.Member = value
		ArtistsToDisplay.Members = append(ArtistsToDisplay.Members, *toAppend)
	}

	//* We make a new struct for the concerts
	ArtistsToDisplay.Concert = nil
	for _, value := range Cards.Array[IdArstist].Locations {
		for _, date := range Cards.Array[IdArstist].Relations[value] {
			toAppend := new(groupietrackers.Concert)
			toAppend.Location = value
			toAppend.Date = date
			ArtistsToDisplay.Concert = append(ArtistsToDisplay.Concert, *toAppend)
		}
	}
	//* We sort concert to have the next concert first
	index := 0
	lenght := len(ArtistsToDisplay.Concert) - 1
	for index < lenght {
		if groupietrackers.DateCompare(ArtistsToDisplay.Concert[index].Date, ArtistsToDisplay.Concert[index+1].Date) {
			ArtistsToDisplay.Concert[index], ArtistsToDisplay.Concert[index+1] = ArtistsToDisplay.Concert[index+1], ArtistsToDisplay.Concert[index]
			index = 0
		} else {
			index++
		}
	}
	//* we add id in map for the map change
	for index := range ArtistsToDisplay.Concert {
		ArtistsToDisplay.Concert[index].Id = index
	}
	return ArtistsToDisplay
}

func SortingList(w http.ResponseWriter, r *http.Request) {
	Entry := r.FormValue("sorting") // vas chercher sorting dans mainPage.html
	// reverseAlphabeticalOrder := r.FormValue("sorting")
	// numberOfArtists := r.FormValue("sorting")
	// alphabeticalOrder := SortByAlphabetical()

	if Entry == "alphabeticalOrder" {
		SortByAlphabetical(Cards.Array)
	} else if Entry == "reverseAlphabeticalOrder" {
		SortByReverseAlphabetical(Cards.Array)
	} else if Entry == "numberOfArtists" {
		SortByNumberOfArtist(Cards.Array)
	}

	// if alphabetical order = true --> return main page
	// fmt.Println(alphabeticalOrder)
	// SortByReverseAlphabetical(Cards.Array)
	MainPage(w, r)
}

func SortByAlphabetical(Entry []groupietrackers.Artists) {
	index := 0
	lenght := len(Entry)

	for index < lenght-1 {
		if Entry[index].Name > Entry[index+1].Name {
			Entry[index], Entry[index+1] = Entry[index+1], Entry[index]
			index = 0
		} else {
			index++
		}

	}
	Cards.Array = Entry
	CardsPagination = groupietrackers.IntoMultiplePages(&NumberOfCards, Entry, 1, &Cards.ForReacherchBar)
}

func SortByReverseAlphabetical(Entry []groupietrackers.Artists) {
	index := 0
	lenght := len(Entry)

	for index < lenght-1 {
		if Entry[index].Name < Entry[index+1].Name {
			Entry[index], Entry[index+1] = Entry[index+1], Entry[index]
			index = 0
		} else {
			index++
		}

	}
	Cards.Array = Entry
	CardsPagination = groupietrackers.IntoMultiplePages(&NumberOfCards, Entry, 1, &Cards.ForReacherchBar)
}

func SortByNumberOfArtist(Entry []groupietrackers.Artists) {
	index := 0
	lenght := len(Entry) - 1

	for index < lenght {
		if len(Entry[index].Members) > len(Entry[index+1].Members) {
			Entry[index], Entry[index+1] = Entry[index+1], Entry[index]
			index = 0
		} else {
			index++
		}
	}

	Cards.Array = Entry
	CardsPagination = groupietrackers.IntoMultiplePages(&NumberOfCards, Entry, 1, &Cards.ForReacherchBar)
}

func AllConcertsInit() {
	for i := range Cards.Array {
		Cards.Array[i].Concert = nil
		for _, value := range Cards.Array[i].Locations {
			for _, date := range Cards.Array[i].Relations[value] {
				toAppend := new(groupietrackers.Concert)
				toAppend.Location = value
				toAppend.Date = date
				Cards.Array[i].Concert = append(Cards.Array[i].Concert, *toAppend)
			}
		}
		//* We sort concert to have the next concert first
		index := 0
		lenght := len(Cards.Array[i].Concert) - 1
		for index < lenght {
			if groupietrackers.DateCompare(Cards.Array[i].Concert[index].Date, Cards.Array[i].Concert[index+1].Date) {
				Cards.Array[i].Concert[index], Cards.Array[i].Concert[index+1] = Cards.Array[i].Concert[index+1], Cards.Array[i].Concert[index]
				index = 0
			} else {
				index++
			}
		}
	}
}
