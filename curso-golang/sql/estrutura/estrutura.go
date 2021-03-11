package main

//para instalar o sql driver: go get -u github.com/go-sql-driver/mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //importe que não será explicitamente usado, mas necessário na app
)

//função exec recebe como parâmetro um db do tipo ponteiro de sql.Db e uma string, retornando uma função sql.Result (pacote database/sql)
func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql) //a vari;ável resultado vai receber e executar esse db, com a string sql
	if err != nil {
		panic(err)
	}
	return result //caso não haja erros, retornará o resultado do banco de dados
}

//o baco de dados necessita estar sendo executado
func main() {
	db, err := sql.Open("mysql", "root: @/") //abertura de conexão com banco de dados, chamada de mysql (porém ainda não existe), passando o user:senha e @/ significa que não iremos nos conectar a um banco de dados específico
	if err != nil {
		panic(err)
	}
	defer db.Close()
	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if existes usuarios")
	exec(db, `create table usuarios(
		id nteger auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
