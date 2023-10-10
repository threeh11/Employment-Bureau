package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	_, err := w.Write([]byte("Привет из Бюро вакансий"))
	if err != nil {
		return
	}
}

func admin(w http.ResponseWriter, r *http.Request) {

}
