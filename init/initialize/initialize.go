package initialize

import (
	cc "github.com/onurozerdal/github.com/onurozerdal/hepsiburada-case/init/config"
	"github.com/onurozerdal/github.com/onurozerdal/hepsiburada-case/init/database"
	"github.com/onurozerdal/github.com/onurozerdal/hepsiburada-case/init/infrastructure"
	"github.com/onurozerdal/github.com/onurozerdal/hepsiburada-case/init/repository"
)

var config cc.Config

func Initialize() {
	config.Read()

	connection := database.NewConnection()
	sqlHandler := infrastructure.NewSqlHandler(connection)
	repositoryCreate := repository.NewRepository(*sqlHandler)
	repositoryCreate.CreateTable()

	connectionDb := database.NewDbConnection()
	sqlHandlerDb := infrastructure.NewSqlHandler(connectionDb)
	repositoryDb := repository.NewRepository(*sqlHandlerDb)
	repositoryDb.CreateTable()
	repositoryDb.CreateHistory()
	repositoryDb.CreateBestseller()
	repositoryDb.CreateProducts()
}
