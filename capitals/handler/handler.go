package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/the_clean_architecture_crud/capitals"
	"github.com/the_clean_architecture_crud/model"
)

type Handler struct {
	CapitalsUsecase capitals.CapitalsUsecase
}

func CreateHandler(router *gin.Engine, capitalsUsecase capitals.CapitalsUsecase) {
	Handler := &Handler{capitalsUsecase}
	//router.LoadHTMLGlob("github.com/the_clean_architecture_crud/forms/*")

	router.GET("/", Handler.HomePage)
	router.GET("/show/:id", Handler.ShowById)
	router.GET("/new", NewHTML)
	router.GET("/edit/:id", Handler.EditHTML)
	router.GET("/delete/:id", Handler.DeleteHTML)
	router.POST("/insert", Handler.InsertCapital)
	router.POST("/update/:id", Handler.UpdateCapital)
	router.POST("delete/:id", Handler.DeleteCapital)
}

func (h *Handler) HomePage(c *gin.Context) {
	capitals, err := h.CapitalsUsecase.ShowAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.HTML(200, "Index.html", gin.H{"result": capitals})
}

func (h *Handler) ShowById(c *gin.Context) {
	id_string := c.Param("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		log.Println("ID you entered is not valid", err)
		return
	}
	capital, err := h.CapitalsUsecase.ShowById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.HTML(200, "Show.html", capital)
}

func (h *Handler) InsertCapital(c *gin.Context) {
	var input model.Capitals
	input.Country = c.PostForm("country")
	input.Capital = c.PostForm("capital")

	if input.Country == "" || input.Capital == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	_, err := h.CapitalsUsecase.InsertCapital(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(302, "/")
}

func (h *Handler) UpdateCapital(c *gin.Context) {

	var input model.Capitals
	input.Id, _ = strconv.Atoi(c.Param("id"))
	input.Country = c.PostForm("country")
	input.Capital = c.PostForm("capital")

	if input.Country == "" || input.Capital == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err := h.CapitalsUsecase.UpdateCapital(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(302, "/")
}

func (h *Handler) DeleteCapital(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("ID you entered is not valid", err)
		return
	}
	_, err = h.CapitalsUsecase.DeleteCapital(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

//router.GET("/edit/:id", EditCapitalHTML)
func (h *Handler) EditHTML(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("ID you entered is not valid", err)
		return
	}
	capital, err := h.CapitalsUsecase.EditHTML(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "Edit.html", capital)
}

//router.GET("/delete/:id", DeleteHTML)

func (h *Handler) DeleteHTML(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("ID you entered is not valid", err)
		return
	}
	capital, err := h.CapitalsUsecase.ShowById(id)
	c.HTML(http.StatusOK, "Delete.html", capital)
}

//router.GET("/new", NewCapitalHTML)
func NewHTML(c *gin.Context) {

	c.HTML(http.StatusOK, "New.html", "")
}
