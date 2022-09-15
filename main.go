package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Capitals struct {
	Id      int    `form:"id" json:"id"`
	Country string `form:"country" json:"country", binding: "required"`
	Capital string `form:"capital" json:"capital", binding: "required"`
}

var id int
var country, capital string

func dbConn() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Mysqlpassword"
	dbName := "capitals"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Println("Couldnt connect to the DB due to error:", err.Error())
		return nil
	}
	log.Println("DB Connectoin successful")
	return db
}

func HomePage(c *gin.Context) {
	db := dbConn()
	if db == nil {
		log.Println("Couldnt establish connection with the database")
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM capitals ORDER BY id DESC")
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	item := Capitals{}
	res := []Capitals{}
	for selDB.Next() {
		err = selDB.Scan(&id, &country, &capital)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
		}
		item.Id = id
		item.Country = country
		item.Capital = capital
		res = append(res, item)
	}
	c.HTML(http.StatusOK, "Index.html", gin.H{"result": res})
}

func ShowCapital(c *gin.Context) {
	db := dbConn()
	defer db.Close()
	id, _ = strconv.Atoi(c.Param("id"))
	selDB, err := db.Query("SELECT * FROM Capitals WHERE id=?", id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	item := Capitals{}
	for selDB.Next() {
		err = selDB.Scan(&id, &country, &capital)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
		}
	}

	item.Id = id
	item.Country = country
	item.Capital = capital

	c.HTML(http.StatusOK, "Show.html", item)
}

func NewCapital(c *gin.Context) {

	c.HTML(http.StatusOK, "New.html", "")
}

func EditCapital(c *gin.Context) {
	id, _ = strconv.Atoi(c.Param("id"))

	db := dbConn()
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM Capitals WHERE id=?", id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	item := Capitals{}
	for selDB.Next() {
		err = selDB.Scan(&id, &country, &capital)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
		}
	}
	item.Id = id
	item.Country = country
	item.Capital = capital
	c.HTML(http.StatusOK, "Edit.html", item)

}

func InsertCapital(c *gin.Context) {

	db := dbConn()
	defer db.Close()
	log.Println("Insert capital")

	var input Capitals
	input.Country = c.PostForm("country")
	input.Capital = c.PostForm("capital")
	log.Println(input)

	if input.Country == "" || input.Capital == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	insForm, err := db.Prepare("INSERT INTO capitals (country, capital) VALUES (?,?)")
	if err != nil {
		return
	}
	insForm.Exec(&input.Country, &input.Capital)
	c.Redirect(http.StatusFound, "/")
}

func UpdateCapital(c *gin.Context) {
	db := dbConn()
	defer db.Close()

	_ = c.Request.ParseMultipartForm(10)
	id, _ = strconv.Atoi(c.Param("id"))
	country = c.Request.FormValue("country")
	capital = c.Request.FormValue("capital")

	if country == "" || capital == "" || id == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	updateForm, err := db.Prepare("UPDATE Capitals SET country=?, capital=? WHERE id=? ;")
	if err != nil {
		return
	}
	updateForm.Exec(country, capital, id)

	c.Redirect(http.StatusFound, "/")
}
func DeleteHTML(c *gin.Context) {

	id, _ = strconv.Atoi(c.Param("id"))

	db := dbConn()
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM Capitals WHERE id=?", id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	item := Capitals{}
	for selDB.Next() {
		err = selDB.Scan(&id, &country, &capital)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
		}
	}
	item.Id = id
	item.Country = country
	item.Capital = capital
	c.HTML(http.StatusOK, "Delete.html", item)
}

func DeleteCapital(c *gin.Context) {

	db := dbConn()
	defer db.Close()
	log.Println("delete")

	id, _ = strconv.Atoi(c.Param("id"))
	delForm, err := db.Prepare("DELETE FROM capitals WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	c.Redirect(http.StatusFound, "/")
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("forms/*")

	router.GET("/", HomePage)
	router.GET("/show/:id", ShowCapital)
	router.GET("/new", NewCapital)
	router.GET("/edit/:id", EditCapital)
	router.GET("/delete/:id", DeleteHTML)
	router.POST("/insert", InsertCapital)
	router.POST("/update/:id", UpdateCapital)
	router.POST("delete/:id", DeleteCapital)

	log.Fatal(router.Run("localhost:3333"))
}
