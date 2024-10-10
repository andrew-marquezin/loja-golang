package controllers

import (
	"alura_loja/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.SelectAll()

	temp.ExecuteTemplate(w, "Index", produtos)
}
