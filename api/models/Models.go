package models


type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Relation struct {
	Index []struct {
		ID             int `json:"id"`
		DatesLocations map[string][]string
	} `json:"index"`
}

//__________________________________________________________________________________________
//                                                                   all data api
type Cache struct {
	ArtistsURL   string `json:"artists"`
	LocationsURL string `json:"locations"`
	DatesURL     string `json:"dates"`
	RelationURL  string `json:"relation"`

	Artists   []Artist
	Locations Locations
	Dates     Dates
	Relation  Relation
}

type Data struct {
	ArtistsID     int
	Image         string
	Name          string
	Members       []string
	CreationDate  int
	FirstAlbum    string
	LocationsLink string
	ConcertDates  string
	Relations     string
	Locations      []string
	LocationsDates string
	RelationStruct map[string][]string
	Dates []string

	ErrorCode int
	Error     string

	//FoundBy []string
	//JSONLen int
}