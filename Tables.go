package groupietrackers

type Artists struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

/*{"id":1,
"image":"https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
"name":"Queen",
"members":["Freddie Mercury","Brian May","John Daecon","Roger Meddows-Taylor","Mike Grose","Barry Mitchell","Doug Fogie"],
"creationDate":1970,
"firstAlbum":"14-12-1973",
"locations":"https://groupietrackers.herokuapp.com/api/locations/1",
"concertDates":"https://groupietrackers.herokuapp.com/api/dates/1",
"relations":"https://groupietrackers.herokuapp.com/api/relation/1"}*/
