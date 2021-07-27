package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func pull() {
	smhi_endpoint := "https://opendata-download-metfcst.smhi.se/api/category/pmp3g/version/2/geotype/point/lon/11.9746/lat/57.7089/data.json"
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	//sb := string(body)
	log.Printf(string(body))
}
