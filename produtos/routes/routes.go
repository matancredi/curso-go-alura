package routes

import (
	"curso-go-alura/produtos/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
