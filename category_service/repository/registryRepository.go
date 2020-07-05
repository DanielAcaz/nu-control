package repository

import (
	"log"
	"strconv"

	databaseConfig "../config"
	registryEntity "../entity"
)


func GetRegistriesByStartDate(year, month, day int) []registryEntity.RegistryEntity {
	var db = databaseConfig.GetConnection()
	defer db.Close()

	startDate := strconv.Itoa(year) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(day)

	var registries []registryEntity.RegistryEntity

	db.Where("date > ?", startDate).Find(&registries)

	log.Printf("going repository with result size: %d", len(registries))
	return registries
}
