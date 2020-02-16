//Web Scraper - Word Counter
package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Starting the application...")
	router := mux.NewRouter()
	router.HandleFunc("/scrape", ScrapeHandler).Methods("GET")
	http.ListenAndServe(":12345", router)
}

func ScrapeHandler(response http.ResponseWriter, request *http.Request) {
	//Scraping Website
	resp, _ := http.Get("http://toscrape.com")
	//Read The Website
	var bytes []byte
	bytes, _ = ioutil.ReadAll(resp.Body)
	domDocTest := html.NewTokenizer(strings.NewReader(string(bytes)))

	var dict = make(map[string]int)
	fmt.Println("\n\n\nStart\n\n\n", dict)
	//Find the HTML Tags & Get the Content inside
	for tokenType := domDocTest.Next(); tokenType != html.ErrorToken; {
		if tokenType != html.TextToken {
			tokenType = domDocTest.Next()
			continue
		}
		//Parse The Content Specifying the Delimiters
		words := strings.FieldsFunc(string(domDocTest.Text()), func(r rune) bool { return strings.ContainsRune(" .,:\t\n(){}", r) })
		//Store the Words into a Dictionary
		for _, v := range words {
			dict[string(v)]++
		}
		tokenType = domDocTest.Next()
	}
	//Print the Dictionary & Encode the Result into JSON Response Body
	fmt.Println("\n\n\nEnds\n\n\n", dict, "Length:", len(dict))
	json.NewEncoder(response).Encode(dict)
}
