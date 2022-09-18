package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/the_clean_architecture_crud/capitals/handler"
	"github.com/the_clean_architecture_crud/capitals/repo"
	"github.com/the_clean_architecture_crud/capitals/usecase"
	"github.com/the_clean_architecture_crud/config"
	"github.com/the_clean_architecture_crud/utils"
)

func main() {

	db := config.DbConn()
	defer db.Close()

	server := `localhost`
	serverPort := 3333

	router := utils.GetRouter(true)

	capitalRepo := repo.CreateCapitalsRepo(db)
	capitalService := usecase.CreateCapitalsUsecase(capitalRepo)
	handler.CreateHandler(router, capitalService)

	utils.RunRouter(router, server, serverPort)
}
