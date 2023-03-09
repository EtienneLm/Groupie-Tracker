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

type ArtistsToDisplay struct {
	Id           int
	Image        string
	Name         string
	Members      []Member
	CreationDate int
	FirstAlbum   string
	Concert      []Concert
	SpotifyId    string
}
type Member struct {
	Member string
}
type Concert struct {
	Location string
	Date     []DateConcert
}
type DateConcert struct {
	Date string
}
type AdminCheck struct {
	IsConnected bool
	IsBadInput  bool
}

// type SortName struct {
// 	Id   int
// 	Name string
// }

type Cards struct {
	Array        []Artists
	ToDisplay    []Artists
	NotFirstPage bool
	NotLastPage  bool
	IdPage       int
	PreviousPage int
	NexPage      int
	IsCardIn     bool
}

// todo le reste des artistes 30 au total
var Spotify = map[int]string{
	1:  "1dfeR4HaWDbWqFHLkxsg1d",
	2:  "2vaWvC8suCFkRXejDOK7EE",
	3:  "0k17h0D3J5VfsdmQ1iZtE9",
	4:  "27T030eWyCQRmDyuvr1kxY",
	5:  "15UsOTVnJzReFVN1VCnxy4",
	6:  "4LLpKhyESsyAXpc4laK94U",
	7:  "6C1ohJrd5VydigQtaGy5Wa",
	8:  "2YZyLoL8N0Wb9xBt1NhZWg",
	9:  "711MCceyCBcFnzjGY4Q7Un",
	10: "1w5Kfo2jwwIPruYS2UWh56",
	11: "6jJ0s89eD6GaHleKKya26X",
	12: "5pKCCKE2ajJHZ9KAiaK11H",
	13: "3CkvROUTQ6nRi9yQOcsB50",
	14: "4lxfqrEsLX6N1N4OCSkILp",
	15: "36QJpDe2go2KgaRleHCDTp",
	16: "776Uo845nYHJpNaStv1Ds4",
	17: "1LZEQNv7sE11VDY3SdxQeN",
	18: "568ZhdwyaiCyOGJRtNYhWf",
	19: "7Ey4PD4MYsKc5I2dolUwbH",
	20: "0WwSkZ7LtFUFjGjMZBMt6T",
	21: "5Q9RKJrjHdfpWVxzv45XTJ",
	22: "0RqtSIYZmd4fiBKVFqyIqD",
	23: "53XhwfbYqKCa1cC15pYq2q",
	24: "4MCBfE4596Uoi2O4DtmEMz",
	25: "4xRYI6VqpkE3UwrDrAZL8L",
	26: "5IH6FPUwQTxPSXurCrcIov",
	27: "2FjkZT851ez950cyPjeYid",
	28: "6cEuCEZu7PAE9ZSzLLc2oQ",
	29: "246dkjvS1zLTtiykXe5h60",
	30: "0Y5tJX1MQlPlqiwlOH1tJY",
	31: "6l3HvQ5sa6mXTsMTB19rO5",
	32: "6deZN1bslXzeGvOLaLMOIF",
	33: "6O2zJ0tId7g07yzHtX0yap",
	34: "3qm84nBOXUEQ2vnTfUTTFC",
	35: "4EnEZVjo3w1cwcQYePccay",
	36: "51Blml2LZPmy7TTiAg47vQ",
	37: "7Ln80lUS6He07XvHI8qqHH",
	38: "4UXqAaa6dQYAk18Lv7PEgX",
	39: "3AA28KZvwAUcZuOKwyblJQ",
	40: "0ECwFtbIWEVNwjlrfc6xoL",
	41: "6XyY86QOPPrYVGvF9ch6wz",
	42: "0L8ExT028jH3ddEcZwqJJ5",
	43: "7dGJo4pcD2V6oG8kP0tJRR",
	44: "7oPftvlwr6VrsViSDV7fJY",
	45: "2ye2Wgw4gimLv2eAKyk1NB",
	46: "4gzpq5DPGxSnKTe4SA8HAU",
	47: "04gDigrS5kc9YWfZHwBETP",
	48: "3YQKmKGau1PzlVlkL1iodx",
	49: "22bE4uQ6baNwSHPVcDxLCe",
	50: "12Chz98pHFMPJEknJQMWvI",
	51: "7jy3rLJdDQY21OgRLCZ9sD",
	52: "69GGBxA162lTqCwzJG5jLp"}
