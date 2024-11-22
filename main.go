package main

import (
	"fmt"
	"net/http"

	f "groupie/groupie"
	// tu va dans le package corps, et tu prend tout ce que contient ce repertoir
	//"groupie/groupie"
)

// var client *http.Client
const port = ":8080"

func main() {
	// client = &http.Client{Timeout: 10 * time.Second}
	// corps.GetArtists("queen")
	// f.GetArtistData("10")
	http.HandleFunc("/", f.Home)
	http.HandleFunc("/artist", f.Artist)
	http.HandleFunc("/submit", f.Submit)
	http.HandleFunc("/search", f.SearchHandler)
	// http.HandleFunc("/search", f.Search)
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	// f.Search2("Jahseh Dwayne Ricardo Onfroy -- members")
	// f.GetArtists1(5)
	// f.DisplayArtist(10)
	// tabl := []int{7}
	// tabp := []int{27, 36}
	// tabk := []int{14, 27, 36}
	// tabj := []int{36}
	// f.GetFilterDate2(1995, 2000, []int{2, 21, 22, 32, 35, 39, 41, 43, 46})
	// f.GetFilterDate(1975, 1985, tabk)
	// f.GetNbrMembers(tabl, tabp)
	// f.GetLocation("MUMBAI, INDIA", tabj)
	fmt.Println("http://localhost:8080) - server started on port", port)
	// f.GetArtistByDate("2010", "2015", "2010", "2015")
	http.ListenAndServe(port, nil)
}
