package bestseller

import (
	"fmt"
	"github.com/onurozerdal/hepsiburada-case/etl-process/database"
	"github.com/onurozerdal/hepsiburada-case/etl-process/infrastructure"
	"github.com/onurozerdal/hepsiburada-case/etl-process/repository"
)

func Find() {
	connectionDataDb := database.NewDataDbConnection()
	sqlHandler := infrastructure.NewSqlHandler(connectionDataDb)
	repositoryDataDb := repository.NewRepository(*sqlHandler)
	r, err := repositoryDataDb.FindBestsellers()
	if err != nil {
		panic(err)
	}

	connectionRecommendationDb := database.NewRecommendedDbConnection()
	sqlHandlerRecommendation := infrastructure.NewSqlHandler(connectionRecommendationDb)
	repositoryRecommendationDb := repository.NewRepository(*sqlHandlerRecommendation)

	repositoryRecommendationDb.ClearBestseller()

	for i := range r {
		fmt.Println(i, ": ", r[i])
		repositoryRecommendationDb.SaveBestseller(r[i])
	}
}
