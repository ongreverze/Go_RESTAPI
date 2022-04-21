package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "rptcomm.postgres.database.azure.com"
	port     = 5432
	user     = "rpt_read"
	password = "rpt1234!!"
	dbname   = "rpt_poc"
  )

type Data struct {
	Unique_id string `json:"unique_id" binding:"required"`
	User_name string `json:"user_name" binding:"required"`
	Age int `json:"age" binding:"required"`
}

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s",
    host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func main(){
	var d Data
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.POST("/v1/svc-rpt/report/rptRequestGenReport", func(c *gin.Context){

		db := OpenConnection()

		if err := c.ShouldBindJSON(&d); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		insertStatement := `INSERT INTO test_ong(unique_id,user_name, age) VALUES ($1,$2,$3)`

		rs,err := db.Exec(insertStatement, d.Unique_id, d.User_name, d.Age)

		if (err != nil){
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
			return
		}

		fmt.Println(rs)
		// close connection
		defer db.Close()

		c.JSON(http.StatusCreated, "accepted")
		
		
	})
	router.Run()
}