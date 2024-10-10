package routes

import (
	"alura_loja/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new-product", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.Insert)
}
