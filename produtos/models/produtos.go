package models

import "curso-go-alura/produtos/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaBanco()

	selectProdutos, err := db.Query("select * from produtos order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

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

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()

	insereDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDados.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletarProduto(id string) {
	db := db.ConectaBanco()

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarProduto.Exec(id)

	defer db.Close()
}

func BuscaProdutoPorId(id string) Produto {
	db := db.ConectaBanco()

	produto, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoRetorno := Produto{}

	for produto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = produto.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtoRetorno.Id = id
		produtoRetorno.Nome = nome
		produtoRetorno.Descricao = descricao
		produtoRetorno.Quantidade = quantidade
		produtoRetorno.Preco = preco
	}
	defer db.Close()
	return produtoRetorno
}

func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()
	Atualiza, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	Atualiza.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()

}
