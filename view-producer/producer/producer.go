package producer

import (
	"bufio"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	cc "github.com/onurozerdal/github.com/onurozerdal/hepsiburada-case/view-producer/config"
	"log"
	"os"
	"time"
)

var config cc.Config

func Produce() {
	file, err := os.Open("product-views.json")

	if err != nil {
		log.Fatalf("failed to open")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	config.Read()
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.Kafka.Bootstrap})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := config.Kafka.Topic
	for scanner.Scan() {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(scanner.Text()),
		}, nil)
		time.Sleep(1 * time.Second)
	}

	p.Flush(15 * 1000)

	file.Close()
}
