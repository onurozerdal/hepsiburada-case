package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	cc "github.com/onurozerdal/hepsiburada-case/stream-reader/config"
	"github.com/onurozerdal/hepsiburada-case/stream-reader/database"
	"github.com/onurozerdal/hepsiburada-case/stream-reader/infrastructure"
	"github.com/onurozerdal/hepsiburada-case/stream-reader/model"
	"github.com/onurozerdal/hepsiburada-case/stream-reader/repository"
)

var config cc.Config

func Consume() {
	config.Read()
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Kafka.Bootstrap,
		"group.id":          config.Kafka.Group,
		"auto.offset.reset": config.Kafka.Offset,
	})

	if err != nil {
		panic(err)
	}

	c.Subscribe(config.Kafka.Topic, nil)

	connection := database.NewSqlConnection()
	sqlHandler := infrastructure.NewSqlHandler(connection)
	repository := repository.NewRepository(*sqlHandler)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			d := model.ProductView{}

			// Unmarshal or Decode the JSON to the interface.
			json.Unmarshal(msg.Value, &d)
			repository.Save(d)
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	defer c.Close()
}
