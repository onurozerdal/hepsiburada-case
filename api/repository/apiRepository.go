package repository

import (
	"fmt"
	"github.com/onurozerdal/hepsiburada-case/api/infrastructure"
)

type ApiRepository struct {
	dbHandler infrastructure.SqlHandler
}

func NewApiRepository(iDbHandler infrastructure.SqlHandler) *ApiRepository {
	return &ApiRepository{dbHandler: iDbHandler}
}

func (repository *ApiRepository) BrowsingHistories(userId string) ([]string, error) {
	var products []string

	statement := fmt.Sprintf(`select h.product_id as productId from history h where h.user_id ='%s' order by h.date desc limit 10`, userId)

	rows, err := repository.dbHandler.QueryContext(statement)
	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var productId string

		err := rows.Scan(&productId)
		if err != nil {
			return products, err
		}

		products = append(products, productId)
	}

	return products, err
}

func (repository *ApiRepository) CheckBrowsingHistories(userId string) bool {
	statement := fmt.Sprintf(`select count(*) as count from history h where h.user_id = '%s'`, userId)

	var count int
	repository.dbHandler.QueryRow(statement).Scan(&count)
	return count == 0
}

func (repository *ApiRepository) Bestseller10Products() ([]string, error) {
	var products []string

	statement := fmt.Sprintf(`select b.product_id as productId from bestseller b order by b.quantity desc limit 10`)

	rows, err := repository.dbHandler.QueryContext(statement)
	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var productId string

		err := rows.Scan(&productId)
		if err != nil {
			return products, err
		}

		products = append(products, productId)
	}

	return products, err
}

func (repository *ApiRepository) Bestseller10ProductsByUserInterest(userId string) ([]string, error) {
	var products []string

	statement := fmt.Sprintf(
		`select b.product_id as productId from bestseller b, products p 
				where b.product_id =p.product_id 
				and p.category_id in 
					(select mostPopular.category_id 
					from 
						(select p.category_id, count(*) 
						from history h, products p 
						where h.product_id = p.product_id 
						and h.user_id = '%s'
						group by p.category_id 
						order by count(*) desc limit 3)  as mostPopular) 
					order by b.quantity desc limit 10`, userId)

	rows, err := repository.dbHandler.QueryContext(statement)
	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var productId string

		err := rows.Scan(&productId)
		if err != nil {
			return products, err
		}

		products = append(products, productId)
	}

	return products, err
}

func (repository *ApiRepository) DeleteHistory(userId, productId string) {
	statement := `delete from history h where h.user_id = $1 and h.product_id = $2`
	repository.dbHandler.QueryRow(statement, userId, productId).Scan(&userId)
}
