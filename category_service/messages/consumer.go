package messages

import (
	"encoding/json"
	"log"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	financeRegistry "../models"
	registryService "../services"
)

func Consume() {
	log.Print("Start Consuming...")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "1",
	})

	defer c.Close()

	if err != nil {
		log.Panic("Error To Connect")
		panic(err)
	}

	c.SubscribeTopics([]string{"create-cartegory-topic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var registry financeRegistry.FinanceRegistry
			json.Unmarshal(msg.Value, &registry)
			registry = registryService.CreateCategoryByBaseRegistries(registry)

			log.Printf("Registry Date: %v", registry.Date)
			log.Printf("Registry Title: %v", registry.Title)
			log.Printf("Registry Category: %v", registry.Category)
			log.Printf("Registry My Category: %v", registry.MyCategory)

		} else {
			log.Fatalf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
