package handler

import (
	"html/template"
	"net/http"
)

func HandlerWorkOutPage(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("templates/workout.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "could not execute template",err)
		return
	}
}