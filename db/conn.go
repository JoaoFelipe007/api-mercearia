package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "aws-0-us-west-1.pooler.supabase.com"
	port     = 6543
	password = "grocery@123"
	user     = "postgres.wuolsxxapvxwpljrotrv"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=require", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
	}

	fmt.Println("Conexão feita com sucesso ao banco de dados!")

	fmt.Println("Conexão feita com sucesso ao " + dbname)

	return db, nil
}
