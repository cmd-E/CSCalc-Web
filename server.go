package main

import (
	"encoding/json"
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
		log.Fatal("main.go tools.Templates.ExecuteTemplate error: " + err.Error())
		// tools.ExecuteError(w, http.StatusInternalServerError, "Internal server error. Template error")
	}
}

type data struct {
	AverageMark int `json:"averageMark,string"`
	ExamMark    int `json:"examMark,string"`
}

func calculate(w http.ResponseWriter, r *http.Request) {
	log.Println("/calculate accessed")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("server.go, sendScore. Error occured: %v", err.Error())
	}
	log.Println(string(requestBody))

	d := data{}
	err = json.Unmarshal(requestBody, &d)
	log.Println(d)
	if err != nil {
		//TODO: handle error
		log.Println(err.Error())
	}
	finalMark := calculator.CalculateFinal(d.AverageMark, d.ExamMark)
	json.NewEncoder(w).Encode(finalMark)
	// averageMark := r.FormValue("averageMark")
	// var amInt int
	// examMark := r.FormValue("examMark")
	// var emInt int
	// if tempAmInt, err := strconv.Atoi(averageMark); err != nil {
	// 	// TODO: Execute error
	// } else {
	// 	amInt = tempAmInt
	// }
	// if tempEmInt, err := strconv.Atoi(examMark); err != nil {
	// 	// TODO: Execute error
	// } else {
	// 	emInt = tempEmInt
	// }
	// finalMark := calculator.CalculateFinal(amInt, emInt)
	// json.NewEncoder(w).Encode(finalMark)
}
