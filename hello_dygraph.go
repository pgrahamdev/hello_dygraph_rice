package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/GeertJohan/go.rice"
)

var csvString string = "Date, Temp\n"

const templateFile string = "dygraph.html"
const portNum int = 8000

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// find a rice.Box
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		log.Fatal(err)
	}
	// get file contents as string
	templateString, err := templateBox.String(templateFile)
	if err != nil {
		log.Fatal(err)
	}
	// parse and execute the template
	htmlTemplate, err := template.New("dygraph").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}

	myTime := time.Now()
	tmpString := csvString
	for i := 0; i < 100; i++ {
		tmpString = tmpString + fmt.Sprintf("%s, %f\n", myTime.Add(time.Hour*time.Duration(24*i)).Format("01/02/2006"), rand.Float32()*100)
	}
	htmlTemplate.Execute(w, tmpString)
}

func main() {
	http.HandleFunc("/", rootHandler)
	box := rice.MustFindBox("assets")
	assetsFileServer := http.StripPrefix("/assets/", http.FileServer(box.HTTPBox()))
	http.Handle("/assets/", assetsFileServer)
	fmt.Printf("Server listening at http://localhost:%d\n", portNum)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(portNum), nil))
}
