package models

import (
	"github.com/carlos/git/golang/loja/bd"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade, Id  int
}

func BuscaProdutos() []Produto {

	db := bd.ConectaBD()
	p := Produto{}
	produtos := []Produto{}

	produtosEncontrados, err := db.Query("select * from produtos order by id asc")

	if err != nil {
		panic(err.Error())
	}

	for produtosEncontrados.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtosEncontrados.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := bd.ConectaBD()

	insereNoBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := bd.ConectaBD()

	removeProduto, err := db.Prepare("delete from produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}

	removeProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := bd.ConectaBD()
	produtoBD, err := db.Query("select * from produtos where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoEncontrado := Produto{}

	for produtoBD.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBD.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtoEncontrado.Id = id
		produtoEncontrado.Nome = nome
		produtoEncontrado.Descricao = descricao
		produtoEncontrado.Preco = preco
		produtoEncontrado.Quantidade = quantidade
	}

	defer db.Close()
	return produtoEncontrado
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := bd.ConectaBD()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()

}
