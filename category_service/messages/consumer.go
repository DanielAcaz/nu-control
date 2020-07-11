package messages

import (
	"encoding/json"
	"log"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	model "github.com/daniel-acaz/nubank-control/category_service/models"
	registryService "github.com/daniel-acaz/nubank-control/category_service/services"
)

func ConsumeCreateCategory() {
	log.Print("Start Create Category Consuming...")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "1",
	})

	defer c.Close()

	if err != nil {
		log.Panic("Error To Connect")
		panic(err)
	}

	_ = c.SubscribeTopics([]string{"create-category-topic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var registry model.FinanceRegistry
			_ = json.Unmarshal(msg.Value, &registry)
			registry = registryService.CreateCategoryByBaseRegistries(registry)

			log.Printf("Registry Date: %v", registry.Date)
			log.Printf("Registry Title: %v", registry.Title)
			log.Printf("Registry Category: %v", registry.Category)
			log.Printf("Registry My Category: %v", registry.MyCategory)
			log.Printf("Registry Accuracy: %v", registry.Accuracy)
			log.Printf("Registry Approved: %v", registry.Approved)

		} else {
			log.Fatalf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}


func ConsumeApprovedCategory() {
	log.Print("Start Approved Category Consuming...")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "1",
	})

	defer c.Close()

	if err != nil {
		log.Panic("Error To Connect")
		panic(err)
	}

	_ = c.SubscribeTopics([]string{"approved-category-topic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var registry model.FinanceRegistry
			_ = json.Unmarshal(msg.Value, &registry)
			registry = registryService.CreateRegistry(registry)

			log.Printf("Registry Date: %v", registry.Date)
			log.Printf("Registry Title: %v", registry.Title)
			log.Printf("Registry Category: %v", registry.Category)
			log.Printf("Registry My Category: %v", registry.MyCategory)
			log.Printf("Registry Accuracy: %v", registry.Accuracy)
			log.Printf("Registry Approved: %v", registry.Approved)

		} else {
			log.Fatalf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}



