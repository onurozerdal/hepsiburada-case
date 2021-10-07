package repository

import (
	"hepsiburada-case/init/infrastructure"
)

type Repository struct {
}

var dbHandler infrastructure.SqlHandler

func NewRepository(iDbHandler infrastructure.SqlHandler) *Repository {
	dbHandler = iDbHandler
	return &Repository{}
}

func (repository *Repository) CreateTable() {
	statement:= `CREATE DATABASE recommendation`
	dbHandler.Exec(statement)
}

func (repository *Repository) CreateHistory() {
	statement:= `create table history
		(
			message_id varchar not null,
			user_id varchar not null,
			product_id varchar not null,
			source varchar not null,
			date timestamp not null
		);
		create unique index history_message_id_uindex
		on history (message_id);
		alter table history
		add constraint history_pk
		primary key (message_id);`
	dbHandler.Exec(statement)
}

func (repository *Repository) CreateBestseller() {
	statement:= `create table bestseller
		(
			product_id varchar not null,
			quantity int not null
		);
		create unique index bestseller_product_id_uindex
		on bestseller (product_id);
		alter table bestseller
		add constraint bestseller_pk
		primary key (product_id);`
	dbHandler.Exec(statement)
}

func (repository *Repository) CreateProducts() {
	statement:= `create table products
		(
			product_id varchar not null,
			category_id varchar not null
		);`
	dbHandler.Exec(statement)
}