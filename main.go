package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

var (
	indexTemplate = template.Must(template.ParseFiles("static/index.html"))
)

func main() {
	git, err := exec.LookPath("git")
	if err != nil {
		log.Fatal("installing git is in your future")
	}
	cmd := exec.Command(git, "log", "--all", "-M", "-C", "--numstat", "--date=short", "--pretty=format:'--%h--%cd--%cn'")
	logs, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	parser := NewParser(logs)
	entries, err := parser.Parse()
	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	churns := AbsoluteTrends(Flatten(entries))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := indexTemplate.Execute(w, nil); err != nil {
			renderError(w, err)
		}
	})
	http.HandleFunc("/bundle.js", func(w http.ResponseWriter, r *http.Request) {
		// Hot load for development
		js, err := ioutil.ReadFile("static/bundle.js")
		if err != nil {
			renderError(w, err)
		}
		fmt.Fprint(w, string(js))
	})
	http.HandleFunc("/churns", func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		err := encoder.Encode(churns)
		if err != nil {
			fmt.Fprint(w, err)
		}
	})
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func renderError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, err)
}
