package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cmd-e/cscalc-web/calculator"
	"github.com/cmd-e/cscalc-web/tools"
)

func init() {
	tools.Templates = template.Must(template.ParseGlob("./static/html/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/calculate", calculate)
	log.Println("Listening at :8080...")
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js")))) // handles css folder
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	err := tools.Templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Fprintf(w, "main.go tools.Templates.ExecuteTemplate error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type data struct {
	AverageMark int `json:"averageMark,string"`
	ExamMark    int `json:"examMark,string"`
}

func calculate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error occurred while reading data: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(string(requestBody))

	d := data{}
	err = json.Unmarshal(requestBody, &d)
	if err != nil {
		fmt.Fprintf(w, "Error occurred while unmarshaling json: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	finalMark := calculator.CalculateFinal(d.AverageMark, d.ExamMark)
	json.NewEncoder(w).Encode(finalMark)
}
