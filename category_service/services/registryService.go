package services

import (
	"github.com/vladimirvivien/automi/collectors"
	"github.com/vladimirvivien/automi/emitters"
	"github.com/vladimirvivien/automi/stream"
	"log"
	"time"

	algorithms "github.com/daniel-acaz/nubank-control/category_service/commons"
	model "github.com/daniel-acaz/nubank-control/category_service/models"
	registryRepository "github.com/daniel-acaz/nubank-control/category_service/repository"
)

func CreateRegistry(registry model.FinanceRegistry) model.FinanceRegistry {
	registryRepository.SaveRegistryAndApproveStatistic(registry)
	return registry
}

func CreateCategoryByBaseRegistries(registry model.FinanceRegistry) model.FinanceRegistry {

	startDate := time.Now().AddDate(0, -6, 0)

	baseRegistries := GetRegistriesByStartDate(startDate)

	baseRegistries = FilterByCommonTitle(registry, baseRegistries, 70)

	if len(baseRegistries) == 0 {
		return registry
	}

	registry.MyCategory = baseRegistries[0].MyCategory
	registry.Accuracy = baseRegistries[0].Accuracy

	return registry

}

func FilterByCommonTitle(registry model.FinanceRegistry, registries []model.FinanceRegistry,
	percentage float64) []model.FinanceRegistry {

	data := emitters.Slice(registries)

	stream := stream.New(data)

	var mapRegistries = make(map[model.FinanceRegistry]float64)

	for _, item := range registries {
		sequencePercentage := algorithms.LongestCommonSubSequencePercentage([]byte(registry.Title), []byte(item.Title))
		mapRegistries[item] = sequencePercentage
	}

	stream.Filter(func(item model.FinanceRegistry) bool {
		return mapRegistries[item] >= percentage
	})

	stream.Batch().SortWith(func(registries interface{}, i, j int) bool {
		sortRegistries := registries.([]model.FinanceRegistry)

		return mapRegistries[sortRegistries[i]] > mapRegistries[sortRegistries[j]]

	})

	var filteredRegistries []model.FinanceRegistry
	stream.Into(collectors.Func(func(data interface{}) error {
		filteredRegistries = data.([]model.FinanceRegistry)
		return nil
	}))

	if err := <-stream.Open(); err != nil {
		log.Fatal(err)
		return nil
	}

	var accurateRegistries []model.FinanceRegistry
	for _, registry := range filteredRegistries {
		registry.Accuracy = mapRegistries[registry]
		accurateRegistries = append(accurateRegistries, registry)
	}

	return accurateRegistries
}

func GetRegistriesByStartDate(startDate time.Time) []model.FinanceRegistry {

	registries := registryRepository.GetRegistriesByStartDate(startDate.Year(), int(startDate.Month()), startDate.Day())

	return registries
}
