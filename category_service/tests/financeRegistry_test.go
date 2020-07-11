package tests

import (
	"encoding/json"
	"testing"
	"time"

	model "github.com/daniel-acaz/nubank-control/category_service/models"
)

func TestConvertStringToTime(test *testing.T) {

	mockRegistryJSON := `{"date":"2020-07-01","category":"casa","title":"Deposito Sao Francisco","amount":10,"my_category":"Familia","family_category":"Outros"}`

	var registry model.FinanceRegistry
	json.Unmarshal([]byte(mockRegistryJSON), &registry)

	if registry == (model.FinanceRegistry{}) {
		test.Errorf("registry couldn't be parse from json, cataegory value is: %v", registry.Category)
	} else if registry.Date.Weekday() != time.Wednesday {
		test.Errorf("data parse with some error, date is: %v", registry.Date.Weekday())
	} else {
		test.Logf("registry was parsed with success")
	}

}
