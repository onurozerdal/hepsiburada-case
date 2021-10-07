package repository

import (
	"github.com/onurozerdal/hepsiburada-case/stream-reader/infrastructure"
	"github.com/onurozerdal/hepsiburada-case/stream-reader/model"
	"time"
)

type Repository struct {
}

var dbHandler infrastructure.SqlHandler

func NewRepository(iDbHandler infrastructure.SqlHandler) *Repository {
	dbHandler = iDbHandler
	return &Repository{}
}

func (repository *Repository) Save(message model.ProductView) {
	sqlInsert := `INSERT INTO history (message_id, user_id, product_id, source, date) VALUES ($1, $2, $3, $4, $5) RETURNING message_id`
	dbHandler.QueryRow(sqlInsert, message.MessageId, message.UserId, message.Properties.ProductId, message.Context.Source, time.Now()).Scan(&message.MessageId)
}
