package groupie

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// defini le port
// const port = ":8080"
var (
	client      *http.Client
	artists     []Artists
	userlocate  UserLocate
	userconcert UserConcert
	// userrelation UserRelation
)

// structure API google maap
type Geo struct {
	Geolocation int
}

// Structure Qui récupère les donnée de l'API Artist
type UserArtist struct {
	Results []Artists
}
type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// Structure qui récupère les données de l'API location
type UserLocate struct {
	Results []Locate `json:"index"`
}
type Locate struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// // Structure qui récupère les données de l'API dates
type UserConcert struct {
	Results []Concerts `json:"index"`
}
type Concerts struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Structure qui récupère les données de l'API relations
type UserRelation struct {
	Results []Relations `json:"index"`
}
type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// // Structure qui récupère les données de toutes les API
type ArtistData struct {
	Artist       Artists
	Locations    []Locate
	Dates        []Concerts
	Relations    []Relations
	TabRelations []string
}

// func GetGoogle() []Geo {
// }

// Fonction qui permet de ressortir les infos de l'API artist
func GetArtists() []Artists {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &artists)
	if err != nil {
		fmt.Printf("erreur : %s\n", err.Error())
	} 
	return artists
}

// Fonction qui permet de ressortir les infos de l'API location
func GetLocate() UserLocate {
	url := "https://groupietrackers.herokuapp.com/api/locations"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &userlocate)
	if err != nil {
		fmt.Printf("erreur : %s\n", err.Error())
		// return
	}
	
	return userlocate
}

// Fonction qui permet de ressortir les infos de l'API dates
func GetConcerts() UserConcert {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &userconcert)
	if err != nil {
		fmt.Printf("erreur : %s\n", err.Error())
		// return
	}
	
	return userconcert
}

// Fonction qui permet de ressortir les infos de l'API relation
func GetRelations() UserRelation {
	url := "https://groupietrackers.herokuapp.com/api/relation"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	var user UserRelation
	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		fmt.Printf("erreur : %s\n", err.Error())
		// return
	} /*else {
		// fmt.Print(user.Results[0])
		// fmt.Print("\n")
	}*/
	return user
}

// Fonction pour récupérer toutes les informations liées à un artiste depuis les différentes API
func GetArtistData(artistID string) ArtistData {
	// Récupérer les informations de l'API Artist
	artist := GetArtistByID(artistID)
	// Récupérer les informations de l'API Location pour cet artiste
	locations := GetLocationsByArtistID(artistID)
	// Récupérer les informations de l'API Dates pour cet artiste
	dates := GetDatesByArtistID(artistID)
	// Récupérer les informations de l'API Relations pour cet artiste
	relations := GetRelationsByArtistID(artistID)
	test := GetRelationsByArtistID(artistID)
	var tab []string
	for _, r := range test {
		for _, tabMap := range r.DatesLocations {
			tab = append(tab, tabMap...)
		}
	}
	// Rassembler toutes les informations dans une structure de données commune
	artistData := ArtistData{
		Artist:       artist,
		Locations:    locations,
		Dates:        dates,
		Relations:    relations,
		TabRelations: tab,
	}
	return artistData
}

// Fonction pour récupérer les informations de l'API Artist par ID
func GetArtistByID(artistID string) Artists {
	// Récupérer tous les artistes
	artists := GetArtists()
	// Rechercher l'artiste par ID dans la liste
	for _, artist := range artists {
		if strconv.Itoa(artist.Id) == artistID {
			return artist
		}
	}
	// Si l'artiste n'a pas été trouvé, renvoyer une structure Artists vide ou gérer l'erreur autrement selon tes besoins
	return Artists{}
}

// Fonction pour récupérer les informations de l'API Location pour un artiste par ID
func GetLocationsByArtistID(artistID string) []Locate {
	// Récupérer toutes les informations de l'API Location
	userlocate := GetLocate()
	// Filtrer les informations pour obtenir celles liées à l'artiste par ID
	var locations []Locate
	for _, locate := range userlocate.Results {
		if strconv.Itoa(locate.Id) == artistID {
			locations = append(locations, locate)
		}
	}
	return locations
}

// Fonction pour récupérer les informations de l'API Location pour un artiste par ID
func GetReallyLocationsByArtistID(artistID string) []string {
	// Récupérer toutes les informations de l'API Location
	bite := GetLocate()
	id, _ := strconv.Atoi(artistID)
	return bite.Results[id-1].Locations
}

// Fonction pour récupérer les informations de l'API Dates pour un artiste par ID
func GetDatesByArtistID(artistID string) []Concerts {
	// Récupérer toutes les informations de l'API Dates
	userconcert := GetConcerts()
	// Filtrer les informations pour obtenir celles liées à l'artiste par ID
	var dates []Concerts
	for _, concert := range userconcert.Results {
		if strconv.Itoa(concert.Id) == artistID {
			dates = append(dates, concert)
		}
	}
	return dates
}

// Fonction pour récupérer les informations de l'API Relations pour un artiste par ID
func GetRelationsByArtistID(artistID string) []Relations {
	// Récupérer toutes les informations de l'API Relations
	userrelation := GetRelations()
	// Filtrer les informations pour obtenir celles liées à l'artiste par ID
	var relations []Relations
	for _, relation := range userrelation.Results {
		if strconv.Itoa(relation.Id) == artistID {
			relations = append(relations, relation)
		}
	}
	return relations
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

// ------------FONCTION FILTER-------------------//
func GetFirstAlbum(enter, end int) []int {
	var result []int
	r := GetArtists()
	for id, artist := range r {
		var annee string
		for i := 6; i < len(artist.FirstAlbum); i++ {
			annee += string(artist.FirstAlbum[i])
		}
		valeur, _ := strconv.Atoi(annee)
		if valeur >= enter && valeur <= end {
			result = append(result, id+1)
		}
	}

	return result
}

func GetFilterDate(enter, end int, tabId []int) []int {
	var result []int
	for _, id := range tabId {
		artist := GetArtistByID(strconv.Itoa(id))
		if artist.CreationDate >= enter && artist.CreationDate <= end {
			result = append(result, id)
		}
	}
	return result
}



func GetNbrMembers(s []int, tabId []int) []int {
	var result []int
	for _, id := range tabId {
		artist := GetArtistByID(strconv.Itoa(id))
		for i := 0; i < len(s); i++ {
			if len(artist.Members) == s[i] {
				result = append(result, id)
				break
			}
		}
	}

	return result
}

func Sst(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "," && string(s[i+1]) == " " {
			result += "-"
			i++
		} else if string(s[i]) == " " {
			result += "_"
		} else {
			result += string(s[i])
		}
	}
	return result
}

func GetLocation(s string, tabId []int) []int {
	var result []int
	for _, id := range tabId {
		artist := GetReallyLocationsByArtistID(strconv.Itoa(id))
		for i := 0; i < len(artist); i++ {
			s = strings.ToLower(s)
			s = Sst(s)
			if s == artist[i] {
				result = append(result, id)
			}
		}
	}
	return result
}

func Global(enter, end, enter1, end1 int, tab []int, locate string) []string {
	var y []string
	s := GetFirstAlbum(enter, end)
	v := GetFilterDate(enter1, end1, s)
	h := GetNbrMembers(tab, v)
	u := GetLocation(locate, h)
	if len(u) == 0 {
		u = h
	}
	for _, t := range u {
		y = append(y, strconv.Itoa(t))
	}
	return y
}

// <-----------------SEARCH BAR------------------>

func ArtistSuggestions(artists []Artists, query string) []Artists {
	suggestions := GetArtists()
	for _, artist := range artists {
		if strings.HasPrefix(strings.ToLower(artist.Name), query) {
			suggestions = append(suggestions, artist)
		}
		for _, mbr := range artist.Members {
			if strings.HasPrefix(strings.ToLower(mbr), query) {
				suggestions = append(suggestions, artist)
			}
		}
	}
	return suggestions
}

func GetLocations() []Locate {
	// Récupérer toutes les informations de l'API Location
	userlocate := GetLocate()
	// Filtrer les informations pour obtenir celles liées à l'artiste par ID
	var locations []Locate
	for _, locate := range userlocate.Results {
		locations = append(locations, locate)
	}

	return locations
}

func LocateSuggestions(locate []Locate, query string) []Locate {
	suggestions := GetLocations()
	for _, locations := range locate {
		for _, loc := range locations.Locations {
			if strings.HasPrefix(strings.ToLower(loc), query) {
				suggestions = append(suggestions, locations)
			}
		}
	}
	return suggestions
}

// type Suggestion interface{}

type Suggestion struct {
	Artists []Artists
	Locate  []Locate
}

func GlobalSuggestions(locate []Locate, artists []Artists, query string) Suggestion {
	// var suggestions []
	artistSuggestions := ArtistSuggestions(artists, query)
	locateSuggestions := LocateSuggestions(locate, query)

	temp := Suggestion{
		Artists: artistSuggestions,
		Locate:  locateSuggestions,
	}
	return temp
}

type Data struct {
	Artists   []Artists
	Locate    []Locate
	Dates     UserConcert
	Relations UserRelation
}

func SearchData() Data {
	// var tab []string
	// Récupérer les informations de l'API Artist
	artist := GetArtists()

	// Récupérer les informations de l'API Location pour cet artiste
	locations := GetLocations()

	// Récupérer les informations de l'API Dates pour cet artiste
	dates := GetConcerts()
	// Rassembler toutes les informations dans une structure de données commune
	artistData := Data{
		Artists: artist,
		Locate:  locations,
		Dates:   dates,
	}
	return artistData
}

// Fonction erreur bad request
func ErrorBadRequest(s string) error {
	FailsArg := errors.New("bad request")
	for _, r := range s {
		if r < 32 || r > 127 {
			if r == 13 {
				return nil
			}
			return FailsArg
		}
	}
	return nil
}

func Search2(s string) []string {
	var final1 []int
	var final []string
	var stransfo string
	var result []int
	// recupere les donnée API de la structure Artist et Locate
	a := GetArtists()
	d := GetLocate()
	for i := 0; i <= 51; i++ {
		// change la premiere lettre de s en majuscule pour les name artist
		stransfo = strings.Title(s)
		// cherche dans les name

		if strings.Contains(stransfo, a[i].Name) {
			result = append(result, i+1)
		}
		// cherche dans les membres
		for _, t := range a[i].Members {
			if strings.Contains(stransfo, t) {
				result = append(result, i+1)
			}
		}
		// cherche dans les firstalbum
		if stransfo == a[i].FirstAlbum {
			result = append(result, i+1)
		}
		// cherche dans les creation date
		date := strconv.Itoa(a[i].CreationDate)
		if stransfo == date {
			result = append(result, i+1)
		}
		// change la premiere lettre de s en minuscule pour les locations
		if len(s) > 0 {
			lowerFirst := strings.ToLower(string(s[0]))
			s = lowerFirst + s[1:]
		}
		for j := 0; j < len(d.Results[i].Locations); j++ {
			if strings.Contains(d.Results[i].Locations[j], s) {
				result = append(result, i+1)
			} else if strings.Contains(d.Results[i].Locations[j], s) {
				result = append(result, i+1)

			}
		}
	}
	if len(result) != 1 {
		for o := 0; o < len(result); o++ {
			if o != len(result)-1 {
				if result[o] == result[o+1] {
					
					final1 = append(final1, result[o])
					o++
				} else {
					final1 = append(final1, result[o])
					
				}
			} else {
				final1 = append(final1, result[o])
			}
		}
	} else {
		final1 = result
	}
	for _, str := range final1 {
		final = append(final, strconv.Itoa(str))
	}
	fmt.Println("result de search", final)
	return final
}
