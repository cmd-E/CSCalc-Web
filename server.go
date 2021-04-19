package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/cmd-e/cscalc-web/calculator"
	"github.com/cmd-e/cscalc-web/tools"
)

func init() {
	tools.Templates = template.Must(template.ParseGlob("./static/html/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/calculate", calculate)
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
	AverageMark float32 `json:"averageMark"`
	ExamMark    float32 `json:"examMark"`
}

type responseData struct {
	FinalMark float32 `json:"finalMark"`
}

func calculate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
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
	d := data{}
	err = json.Unmarshal(requestBody, &d)
	if err != nil {
		fmt.Fprintf(w, "Error occurred while unmarshaling json: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var errStruct tools.ErrStruct
	if errStruct = tools.MarksAreValid(d.AverageMark, d.ExamMark); errStruct.IsError {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errStruct)
		return
	}
	fm := responseData{FinalMark: calculator.CalculateFinal(d.AverageMark, d.ExamMark)}
	json.NewEncoder(w).Encode(fm)
}
