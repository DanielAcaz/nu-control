package services

import (
	"github.com/vladimirvivien/automi/collectors"
	"github.com/vladimirvivien/automi/emitters"
	"github.com/vladimirvivien/automi/stream"
	"log"
	"time"

	algorithms "../commons"
	financeRegistry "../models"
	registryRepository "../repository"
)

func CreateRegistry(registry financeRegistry.FinanceRegistry) financeRegistry.FinanceRegistry {
	registry.ID = 1
	return registry
}

func CreateCategoryByBaseRegistries(registry financeRegistry.FinanceRegistry) financeRegistry.FinanceRegistry {

	startDate := time.Now().AddDate(0, -6, 0)

	baseRegistries := GetRegistriesByStartDate(startDate)

	baseRegistries = FilterByCommonTitle(registry, baseRegistries, 70)

	if len(baseRegistries) == 0 {
		return registry
	}

	registry.MyCategory = baseRegistries[0].MyCategory

	return registry

}

func FilterByCommonTitle(registry financeRegistry.FinanceRegistry, registries []financeRegistry.FinanceRegistry,
	percentage float64) []financeRegistry.FinanceRegistry {

	data := emitters.Slice(registries)

	stream := stream.New(data)

	var mapRegistries = make(map[financeRegistry.FinanceRegistry]float64)

	for _, item := range registries {
		sequencePercentage := algorithms.LongestCommonSubSequencePercentage([]byte(registry.Title), []byte(item.Title))
		mapRegistries[item] = sequencePercentage
	}

	stream.Filter(func(item financeRegistry.FinanceRegistry) bool {
		return mapRegistries[item] >= percentage
	})

	stream.Batch().SortWith(func(registries interface{}, i, j int) bool {
		sortRegistries := registries.([]financeRegistry.FinanceRegistry)

		return mapRegistries[sortRegistries[i]] > mapRegistries[sortRegistries[j]]

	})

	var filteredRegistries []financeRegistry.FinanceRegistry
	stream.Into(collectors.Func(func(data interface{}) error {
		filteredRegistries = data.([]financeRegistry.FinanceRegistry)
		return nil
	}))

	if err := <-stream.Open(); err != nil {
		log.Fatal(err)
		return nil
	}

	return filteredRegistries
}

func GetRegistriesByStartDate(startDate time.Time) []financeRegistry.FinanceRegistry {

	entities := registryRepository.GetRegistriesByStartDate(startDate.Year(), int(startDate.Month()), startDate.Day())

	var registries []financeRegistry.FinanceRegistry
	for _, entity := range entities {
		registries = append(registries, financeRegistry.FinanceRegistry{
			ID:             entity.ID,
			Date:           financeRegistry.RegistryDate{entity.Date},
			Category:       entity.Category,
			Title:          entity.Title,
			Amount:         entity.Amount,
			MyCategory:     entity.MyCategory,
			FamilyCategory: entity.FamilyCategory})
	}

	return registries
}
