package groupie

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

// Fonction qui permet d'afficher les résultats de la requètes sur la page index.html
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // Erreur 404 si le lien est différent de /submit
		t, _ := template.ParseFiles("./html/404.html")
		t.Execute(w, nil)
	} else {
		// renderTemplate(w, "home")
		t, err := template.ParseFiles("./html/index.html")
		if err != nil {
			// crée une erreur avec le module http ! ne pas oublier le return
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := SearchData()
		
		t.Execute(w, data)
	}
}

// Fonction qui permet d'afficher les résultats des requètes sur la page artist.html
func Artist(w http.ResponseWriter, r *http.Request) {
	// récupération de l'ID de l'artist dans le HTML
	btn := r.FormValue("btn")
	// r.URL.Path = "/artist" + button

	// Condition qui permet d'afficher la page 404 si l'URL est différent de /artist
	if r.URL.Path != "/artist" {
		t, _ := template.ParseFiles("./html/404.html")
		t.Execute(w, nil)
	} else {

		artistData := GetArtistData(btn)

		t, err := template.ParseFiles("./html/artist.html")
		if err != nil {
			// crée une erreur avec le module http ! ne pas oublier le return
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, artistData)
	}
}

func Submit(w http.ResponseWriter, r *http.Request) {
	var v int
	var t int
	var u int
	var p int
	var b []Artists
	var tab []string
	var test int
	var tab1 []int
	first1 := r.FormValue("first_album1")
	first2 := r.FormValue("first_album2")
	creadate1 := r.FormValue("crea_date1")
	creadate2 := r.FormValue("crea_date2")
	locatef := r.FormValue("locate")
	badrequest := ErrorBadRequest(locatef)
	for i := 0; i < 8; i++ {
		tab = append(tab, r.FormValue("members-"+strconv.Itoa(i)))
		test, _ = strconv.Atoi(tab[i])
		tab1 = append(tab1, test)
	}
	v, _ = strconv.Atoi(first1)
	t, _ = strconv.Atoi(first2)
	u, _ = strconv.Atoi(creadate1)
	p, _ = strconv.Atoi(creadate2)
	// mbrs := r.FormValue("members")
	// locate := r.FormValue("locate")
	// Condition qui permet d'afficher la page 404 si l'URL est différent de /submit
	if r.URL.Path != "/submit" {
		t, _ := template.ParseFiles("./html/404.html")
		t.Execute(w, nil)
	} else if badrequest != nil {
		t, _ := template.ParseFiles("./html/400.html")
		t.Execute(w, badrequest)
	} else {
		// artistFilter := Global(v, t, u, p)
		o := Global(v, t, u, p, tab1, locatef)
		for i := 0; i < len(o); i++ {
			b = append(b, GetArtistByID(o[i]))
		}
		t, err := template.ParseFiles("./html/submit.html")
		if err != nil {
			// crée une erreur avec le module http ! ne pas oublier le return
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, b)
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var final []ArtistData
	query := strings.ToLower(r.FormValue("query"))

	badrequest := ErrorBadRequest(query)
	if badrequest != nil {
		t, _ := template.ParseFiles("./html/400.html")
		t.Execute(w, badrequest)
	} else if r.URL.Path != "/search" {
		t, _ := template.ParseFiles("./html/404.html")
		t.Execute(w, nil)
	} else {
		t, _ := template.ParseFiles("./html/search.html")

		tabid := Search2(query)
		for _, id := range tabid {
			final = append(final, GetArtistData((id)))
		}
		err := t.Execute(w, final)
		if err != nil {
			fmt.Println(err)
		}
	}
}
