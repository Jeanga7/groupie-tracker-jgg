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
type Cache struct {
	ArtistsURI   string `json:"artists"`
	LocationsURI string `json:"locations"`
	DatesURI     string `json:"dates"`
	RelationURI  string `json:"relation"`

	Artists   []Artist
	Locations Locations
	Dates     Dates
	Relation  Relation
}
