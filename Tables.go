package groupietrackers

//* Those structs are used to extracte the data from the database
type ExtractLocation struct {
	Index []Location
}
type ExtractDates struct {
	Index []Dates
}
type ExtractRelation struct {
	Index []Relations
}

var LocationEx ExtractLocation
var DatesEx ExtractDates
var RelationEx ExtractRelation

type Relations struct {
	Id             int
	DatesLocations map[string][]string
}

type Location struct {
	Id        int
	Locations []string
}

type Dates struct {
	Id    int
	Dates []string
}

type Artists struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	ConcertDates []string
	Concert      []Concert
	Relations    map[string][]string
}

//* This struct is used to display the data on artist page
//* X and Y are the coordinates of the next concert of the artists
type ArtistsToDisplay struct {
	Id              int
	Image           string
	Name            string
	Members         []Member
	CreationDate    int
	FirstAlbum      string
	Concert         []Concert
	SelectedConcert Concert
	X               float64
	Y               float64
	SpotifyId       string
	Genre           string
	Followers       int
}
type Member struct {
	Member string
}
type Concert struct {
	Id       int
	Location string
	Date     string
}


var Admin AdminCheck
type AdminCheck struct {
	IsConnected bool
	IsBadInput  bool
}

type Cards struct {
	Array           []Artists
	ToDisplay       []Artists
	ForReacherchBar []Artists
	NotFirstPage    bool
	NotLastPage     bool
	IdPage          int
	PreviousPage    int
	NexPage         int
	IsCardIn        bool
}
