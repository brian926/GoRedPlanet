package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Apod struct {
	Url         string `json:"url"`
	Explanation string `json:"explanation"`
	Date        string `json:"data"`
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	test := ""
	resp, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=" + test)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll((resp.Body))
	if err != nil {
		log.Fatalln(err)
	}

	var apod Apod
	json.Unmarshal(body, &apod)

	fmt.Println(apod.Url)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
