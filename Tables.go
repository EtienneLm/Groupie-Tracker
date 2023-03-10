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
	Relations    map[string][]string
	SpotifyId    string
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
	SpotifyId       string
	SelectedConcert Concert
	X               float64
	Y               float64
}
type Member struct {
	Member string
}
type Concert struct {
	Id       int
	Location string
	Date     string
}

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
