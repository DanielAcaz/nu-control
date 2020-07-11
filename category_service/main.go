package main

import (
	"log"

	registryController "github.com/daniel-acaz/nubank-control/category_service/controllers"
	consumer "github.com/daniel-acaz/nubank-control/category_service/messages"
)

func main() {
	log.Print("Startup Service...")
	//consumer.ConsumeCreateCategory()
	consumer.ConsumeApprovedCategory()
	registryController.RegistryController()
}
