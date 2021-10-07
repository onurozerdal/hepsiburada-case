package repository

import (
	"fmt"
	"hepsiburada-case/etl-process/infrastructure"
	"hepsiburada-case/etl-process/model"
)

type Repository struct {
}

var dbHandler infrastructure.SqlHandler

func NewRepository(iDbHandler infrastructure.SqlHandler) *Repository {
	dbHandler = iDbHandler
	return &Repository{}
}

func (repository *Repository) FindBestsellers() ([]model.Bestseller, error) {
	var bestsellers []model.Bestseller
	statement := fmt.Sprintf(
		`select i.product_id as productId, count(distinct o.user_id) as quantity from order_items i, orders o
				where i.order_id = o.order_id
				group by i.product_id
				order by count(distinct o.user_id) desc`)

	// Execute query
	rows, err := dbHandler.QueryContext(statement)
	if err != nil {
		return bestsellers, err
	}

	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		var productId string
		var quantity int

		// Get values from row.
		err := rows.Scan(&productId, &quantity)
		if err != nil {
			return bestsellers, err
		}
		bestseller := model.Bestseller{productId, quantity}
		bestsellers = append(bestsellers, bestseller)
	}

	return bestsellers, err
}

func (repository *Repository) ProductsCategories() ([]model.ProductCategory, error) {
	var productCategories []model.ProductCategory
	statement := fmt.Sprintf(
		`select p.product_id as productId, p.category_id as categoryId from products p`)

	// Execute query
	rows, err := dbHandler.QueryContext(statement)
	if err != nil {
		return productCategories, err
	}

	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		var productId, categoryId string

		// Get values from row.
		err := rows.Scan(&productId, &categoryId)
		if err != nil {
			return productCategories, err
		}
		productCategory := model.ProductCategory{productId, categoryId}
		productCategories = append(productCategories, productCategory)
	}

	return productCategories, err
}

func (repository *Repository) SaveBestseller(bestseller model.Bestseller) {
	sqlInsert:= `INSERT INTO bestseller (product_id, quantity) VALUES ($1, $2) RETURNING product_id`
	dbHandler.QueryRow(sqlInsert, bestseller.ProductId, bestseller.Quantity).Scan(&bestseller.ProductId)
}

func (repository *Repository) SaveProductCategory(productCategory model.ProductCategory) {
	sqlInsert:= `INSERT INTO products (product_id, category_id) VALUES ($1, $2) RETURNING product_id`
	dbHandler.QueryRow(sqlInsert, productCategory.ProductId, productCategory.CategoryId).Scan(&productCategory.ProductId)
}

func (repository *Repository) ClearBestseller() {
	sql:= `truncate bestseller`
	dbHandler.Execute(sql)
}

func (repository *Repository) ClearProductCategory() {
	sql:= `truncate products`
	dbHandler.Execute(sql)
}