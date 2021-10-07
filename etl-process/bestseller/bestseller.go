package bestseller

import (
	"fmt"
	"hepsiburada-case/etl-process/database"
	"hepsiburada-case/etl-process/infrastructure"
	"hepsiburada-case/etl-process/repository"
)

func Find() {
	connectionDataDb := database.NewDataDbConnection()
	sqlHandler := infrastructure.NewSqlHandler(connectionDataDb)
	repositoryDataDb := repository.NewRepository(*sqlHandler)
	r , err := repositoryDataDb.FindBestsellers()
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
