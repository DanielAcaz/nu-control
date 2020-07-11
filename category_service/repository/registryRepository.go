package repository

import (
	"bytes"
	"context"
	"encoding/json"
	databaseConfig "github.com/daniel-acaz/nubank-control/category_service/config"
	model "github.com/daniel-acaz/nubank-control/category_service/models"
	"log"
	"strconv"
)


func GetRegistriesByStartDate(year, month, day int) []model.FinanceRegistry {

	var db = databaseConfig.GetConnection()

	var stringMonth string
	if month <= 6 {
		stringMonth = "0" + strconv.Itoa(month)
	} else {
		stringMonth = strconv.Itoa(month)
	}

	startDate := strconv.Itoa(year) + "-" + stringMonth + "-" + strconv.Itoa(day) + "T00:00:00.000"

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"range": map[string]interface{}{
				"date": map[string]interface{}{
					"gte": startDate,
				},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := db.Search(
		db.Search.WithContext(context.Background()),
		db.Search.WithIndex("registries_index"),
		db.Search.WithSize(600),
		db.Search.WithBody(&buf),
		db.Search.WithTrackTotalHits(true),
		db.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var result databaseConfig.SearchResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	var registries []model.FinanceRegistry
	for _, hit := range result.Hits.Hits {
		registry := hit.Source
		registries = append(registries, registry)
	}

	log.Printf("going repository with result size: %d", len(registries))
	return registries
}
