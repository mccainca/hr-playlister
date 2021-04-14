package scraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/mccainca/hr-playlister/models"
)

//BaseURL : Site containing tracklisting , need to append date/time/on_demand (i.e. "/2021/02/13?time=22:00&on_demand=1")
const BaseURL = "https://tracklist-api.kcrw.com/Simulcast/date/"

//Scrape : Scrapes data from site
func Scrape(date string) models.TrackList {
	var url = fmt.Sprintf("%s%s?time=22:00&on_demand=1", BaseURL, date)
	fmt.Printf("%s\n", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject models.TrackList
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}
