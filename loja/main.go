package main

import (
	"net/http"

	"github.com/carlos/git/golang/loja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
