package groupietrackers

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
	DatesLocations [][]string
}

type Location struct {
	Id       int
	Location []string
	Dates    Dates
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
	FirstAlbum   []string
	Locations    Location
	ConcertDates Dates
	Relations    Relations
}

type Cards struct {
	Array []Artists
}
