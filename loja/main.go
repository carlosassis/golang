package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func conectaBD() *sql.DB {
	conexao := "user=postgres dbname=storedb password=mysecretpassword host=localhost ssl=disabled"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade, Id  int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaBD()
	p := Produto{}
	produtos := []Produto{}

	buscaProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	for buscaProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = buscaProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)

	defer db.Close()
}
