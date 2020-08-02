package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

var results Results

var tpl = template.Must(template.ParseFiles("index.html")) // creates the index.html page

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil) // display template
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String()) //gets string from URL and splits it up
	if err != nil {                     //if errors exist
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	params := u.Query() // end of URL
	searchKey := params.Get("q") // user inputted search item
	location := params.Get("location") //user inputted location
	searchKey = strings.ReplaceAll(searchKey, " ", "") // get rid of spaces
	location = strings.ReplaceAll(location, " ", "")

	// Getting the Data
	client := &http.Client{}
	//request to api
	req, err := http.NewRequest("GET", "https://api.yelp.com/v3/businesses/search?location=" + location + "&term=" + searchKey, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Bearer rJJQr__Caaoo4_hJLTqceXT60zz3dmbOnntMUlpHkR-333e6408R2bloGZzw1JzOYULdq6sOHc2rBiK8cn87dFAmbWQ2EOctsy4QFFwxrLn4RsLO84nTEQbvXCEnX3Yx") //auth token
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewDecoder(response.Body).Decode(&results) //parse response

	err = tpl.Execute(w, results)


	fmt.Println("Search Query is: ", searchKey)
	fmt.Println("Location is ", location)
}

func main() {
	fmt.Println("App Started")

	mux := http.NewServeMux() // helps to call the correct handler based on the URL

	fs := http.FileServer(http.Dir("assets"))  // put static files into server (i.e. styles.css)
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler) // what function to call on main page
	mux.HandleFunc("/search", searchHandler) // mux calls SearchHandler when /search is seen in URL
	http.ListenAndServe(":3000", mux) // start a local server to run files
}

type Business struct {
	Rating     int    `json:"rating"`
	Price      string `json:"price"`
	Phone      string `json:"phone"`
	ID         string `json:"id"`
	Alias      string `json:"alias"`
	IsClosed   bool   `json:"is_closed"`
	Categories []struct {
		Alias string `json:"alias"`
		Title string `json:"title"`
	} `json:"categories"`
	ReviewCount int    `json:"review_count"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	ImageURL string `json:"image_url"`
	Location struct {
		City     string `json:"city"`
		Country  string `json:"country"`
		Address2 string `json:"address2"`
		Address3 string `json:"address3"`
		State    string `json:"state"`
		Address1 string `json:"address1"`
		ZipCode  string `json:"zip_code"`
	} `json:"location"`
	Distance     float64  `json:"distance"`
	Transactions []string `json:"transactions"`
}

type Results struct {
	Total      int        `json:"total"`
	Businesses []Business `json:"businesses"`
}