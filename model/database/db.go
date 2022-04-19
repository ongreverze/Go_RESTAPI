package database

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "password"
  dbname   = "go_db"
)

func Open(unique_id, name, age) {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  sqlStatement := `
INSERT INTO users (unique_id, name, age)
VALUES ($1, $2, $3)`
  id := 0
  err = db.QueryRow(sqlStatement, unique_id, name, age)
  if err != nil {
    panic(err)
  }
  fmt.Println("New record ID is:", id)
}