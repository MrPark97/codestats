package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Config struct {
	Port    string   `yaml:"port"`
	Handles []string `yaml:"handles"`
}

var (
	port   = os.Getenv("PORT")
	config Config
)

func init() {
	yamlFile, err := ioutil.ReadFile("./config.yaml")

	if err != nil {
		log.Println("Error while reading config file ", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Println("Error while unmarshalling config ", err)
	}

	if port == "" {
		port = config.Port
	}
}

func main() {
	http.HandleFunc("/rating", ratingPage)
	http.HandleFunc("/submissions", submissionsPage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func header(w http.ResponseWriter) {
	t, _ := template.ParseFiles("./templates/header.tpl")
	t.Execute(w, struct{}{})
}

func footer(w http.ResponseWriter) {
	t, _ := template.ParseFiles("./templates/footer.tpl")
	t.Execute(w, struct{}{})
}
