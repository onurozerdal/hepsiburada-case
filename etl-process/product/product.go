package product

import (
	"fmt"
	"github.com/onurozerdal/hepsiburada-case/etl-process/database"
	"github.com/onurozerdal/hepsiburada-case/etl-process/infrastructure"
	"github.com/onurozerdal/hepsiburada-case/etl-process/repository"
)

func Process() {
	connectionDataDb := database.NewDataDbConnection()
	sqlHandler := infrastructure.NewSqlHandler(connectionDataDb)
	repositoryDataDb := repository.NewRepository(*sqlHandler)
	r, err := repositoryDataDb.ProductsCategories()
	if err != nil {
		panic(err)
	}

	connectionRecommendationDb := database.NewRecommendedDbConnection()
	sqlHandlerRecommendation := infrastructure.NewSqlHandler(connectionRecommendationDb)
	repositoryRecommendationDb := repository.NewRepository(*sqlHandlerRecommendation)

	repositoryRecommendationDb.ClearProductCategory()

	for i := range r {
		fmt.Println(i, ": ", r[i])
		repositoryRecommendationDb.SaveProductCategory(r[i])
	}
}
