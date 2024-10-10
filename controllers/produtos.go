package controllers

import (
	"alura_loja/models"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.SelectAll()

	temp.ExecuteTemplate(w, "Index", produtos)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			fmt.Println("Erro ao converter o preço:", err)
		}

		qtdConv, err := strconv.Atoi(quantidade)
		if err != nil {
			fmt.Println("Erro ao converter a quantidade:", err)
		}

		models.CreateProduct(nome, descricao, precoConv, qtdConv)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Erro ao converter o ID:", err)
		}

		qtdConv, err := strconv.Atoi(quantidade)
		if err != nil {
			fmt.Println("Erro ao converter o ID:", err)
		}

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			fmt.Println("Erro ao converter o ID:", err)
		}

		models.UpdateProduct(idConv, nome, descricao, precoConv, qtdConv)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
