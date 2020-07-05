package main

import (
	"log"

	registryController "./controllers"
	consumer "./messages"
)

func main() {
	log.Print("Startup Service...")
	consumer.Consume()
	registryController.RegistryController()
}
