package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"fmt"
	"log"
	"strconv"
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
	User_name string `json:"user_name" binding:"required"`
	Age int `json:"age" binding:"required"`
}

// func (d *Data) addData() (err error) {
// 	rs, err := db.Exec("INSERT INTO Users(unique_id, name, age) VALUES (? ,?, ?)", d.Unique_id, d.Name, d.Age)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

var db *sql.DB

func main(){
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.POST("/v1/svc-rpt/report/rptRequestGenReport", func(c *gin.Context){
		user_name := c.PostForm("user_name")
		age,_ := strconv.Atoi(c.PostForm("age"))

		rs, err := db.Exec(`INSERT INTO request_log(user_name, age, unique_id) VALUES (?, ?, ?);`, user_name, age)
		fmt.Println(rs)
        if err != nil {
            log.Fatalln(err)
        }

        if err != nil {
            log.Fatalln(err)
        }

        msg := fmt.Sprintf("insert successful")
        c.JSON(http.StatusOK, gin.H{
            "msg": msg,
        })
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