package whoisinspace

import (
	"fmt"
	"reflect"
	"testing"
)

func getResponse() Response {
	var astronauts []Person
	astronauts = append(astronauts, Person{"Matthew", "test", 10, 10, "country", "countryFlag", "launchdate", 10, "title", "location", "bio", "biolink", "twitter"})
	astronauts = append(astronauts, Person{"James", "test", 10, 10, "country1", "countryFlag", "launchdate", 10, "title", "location", "bio", "biolink", "twitter"})
	astronauts = append(astronauts, Person{"Boyd", "test", 10, 10, "country", "countryFlag", "launchdate", 10, "title", "location", "bio", "biolink", "twitter"})

	response := Response{3, astronauts}

	return response
}

func TestAstronautFunction(t *testing.T) {

	want := Astronauts()
	fmt.Println(reflect.TypeOf(want.People).Kind())
	if reflect.TypeOf(want.NumOfPeopleInSpace).Kind() != reflect.Int {
		t.Fatalf("Number of people is not of type int")

	}
	if reflect.TypeOf(want.People).Kind() != reflect.Slice {
		t.Fatalf("People is not of type []Person")
	}
}

func TestGetAllAstronautNames(t *testing.T) {
	response := getResponse()
	got := response.GetAllNamesOfAstronauts()
	want := []string{"Matthew", "James", "Boyd"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Wrong result back for GetAllAstronautNames")
	}

}

func TestGetAllAstronautsNamesByCountry(t *testing.T) {
	response := getResponse()

	got := response.GetAllNamesByCountry("country")
	want := []string{"Matthew", "Boyd"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Failed on GetAllNamesBycountry, got : %s was not the same as want: %s", got, want)
	}
}

func TestGetAllCountriesOfAstronauts(t *testing.T) {
	response := getResponse()
	result_map := make(map[string]int)
	result_map["country"] = 2
	result_map["country1"] = 1
	got := response.GetAllCountries()
	want := result_map

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Failed on getting all countries of astronauts, got : %v was not the same as want: %v", got, want)
	}
}
