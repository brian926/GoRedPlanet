package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Apod struct {
	Url         string `json:"url"`
	Explanation string `json:"explanation"`
	Date        string `json:"data"`
}

func picOfDay() string {
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

	return apod.Url
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))

	data := picOfDay()
	tmpl.Execute(w, data)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
