package whoisinspace

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Username string
	Password string
}

type Response struct {
	NumOfPeopleInSpace int      `json:"number"`
	People             []Person `json:"people"`
}

type Person struct {
	Name           string
	Biophoto       string
	Biophotowidth  int
	Biophotoheight int
	Country        string
	Countryflag    string
	Launchdate     string
	Careerdays     int
	Title          string
	Location       string
	Bio            string
	Biolink        string
	Twitter        string
}

func Astronauts() *Response {
	resp, err := http.Get("https://www.howmanypeopleareinspacerightnow.com/peopleinspace.json")
	if err != nil {
		panic("There was an error when trying to retrieve the amount of people in space json file")
	}
	var response Response
	json.NewDecoder(resp.Body).Decode(&response)
	return &response
}

func (response *Response) GetAllNamesByCountry(country string) []string {
	var astronauts []string
	for _, val := range response.People {
		if val.Country == country {
			astronauts = append(astronauts, val.Name)
		}
	}
	return astronauts
}

func (response *Response) GetAllNamesOfAstronauts() []string {
	var names []string
	for _, astronaut := range response.People {
		names = append(names, astronaut.Name)
	}
	return names
}

func (response *Response) GetAllCountries() map[string]int {
	countryMap := make(map[string]int)
	for _, astronaut := range response.People {
		countryMap[astronaut.Country] += 1
	}
	return countryMap
}
