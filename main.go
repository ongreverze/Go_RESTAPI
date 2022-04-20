package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"fmt"
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

		fmt.Println(rs)

		if (err != nil){
			fmt.Println(err)
		}

		// close connection
		defer db.Close()

		c.JSON(http.StatusCreated, "accepted")
		
		// unique_id := c.Request.FormValue("unique_id")
		// name := c.Request.FormValue("name")
		// str_age := c.Request.FormValue("age")
		// age ,err:= strconv.ParseInt(str_age, 0, 64)
		// fmt.Println(err)

		// d := Data{Unique_id: unique_id, Name: name, Age: age}
		// err := d.addData()
		// if err := c.ShouldBindJSON(&d); err != nil {
		// 	c.JSON(http.StatusNotFound, gin.H{
		// 		"status" : http.StatusNotFound , 
		// 		"error": "Invalid input!"})
		// 	return
		// }
		// msg := fmt.Sprintf("insert successful")
		// c.JSON(http.StatusAccepted , gin.H{
		// "status" : http.StatusAccepted,
		// "message": msg})
	})
	router.Run(":8001")
}