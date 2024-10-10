package models

import "alura_loja/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func SelectAll() []Produto {
	db := db.ConnectDb()

	selectAll, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectAll.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAll.Scan(&id, &nome, &descricao, &preco, &quantidade)
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

func CreateProduct(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDb()

	insertData, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDb()

	delete, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Produto {
	db := db.ConnectDb()

	dbProduct, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Produto{}

	for dbProduct.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := dbProduct.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		product.Id = id
		product.Nome = nome
		product.Descricao = descricao
		product.Preco = preco
		product.Quantidade = quantidade
	}
	defer db.Close()
	return product
}

func UpdateProduct(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDb()

	update, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	update.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
