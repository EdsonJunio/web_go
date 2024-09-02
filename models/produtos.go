package models

import (
	"web.go/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}
	defer selectDeTodosOsProdutos.Close()

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
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
	db := db.ConectaComBancoDeDados()

	query := "INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES(?, ?, ?, ?)"
	insereDadosNoBanco, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	defer insereDadosNoBanco.Close()

	_, err = insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
