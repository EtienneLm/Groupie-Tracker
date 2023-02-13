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
	DatesLocations map[string][]string
}

type Location struct {
	Id        int
	Locations []string
	Dates     Dates
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
	SpotifyId    string
}

type Cards struct {
	Array []Artists
}

// todo le reste des artistes 12
var Spotify = map[int]string{1: "1dfeR4HaWDbWqFHLkxsg1d", 2: "2vaWvC8suCFkRXejDOK7EE", 3: "0k17h0D3J5VfsdmQ1iZtE9", 4: "27T030eWyCQRmDyuvr1kxY", 5: "15UsOTVnJzReFVN1VCnxy4", 6: "4LLpKhyESsyAXpc4laK94U", 7: "6C1ohJrd5VydigQtaGy5Wa", 8: "2YZyLoL8N0Wb9xBt1NhZWg", 9: "711MCceyCBcFnzjGY4Q7Un", 10: "1w5Kfo2jwwIPruYS2UWh56", 11: "6jJ0s89eD6GaHleKKya26X", 12: "5pKCCKE2ajJHZ9KAiaK11H"}
