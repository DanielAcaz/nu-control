package tests

import (
	financeRegistry "../models"
	registryService "../services"
	"testing"
)


func TestFilterByStream(test *testing.T) {

	registry := financeRegistry.FinanceRegistry{ ID: 1, Title: "Title to Test", FamilyCategory: "FamilyCategoryTest",
		MyCategory: "MyCategoryTest", Amount: 100.0 }


	registries := []financeRegistry.FinanceRegistry{
		{ID: 1, Title: "Title to Testing", FamilyCategory: "FamilyCategoryTest",
			MyCategory: "MyCategoryTest", Amount: 100.0},
		{ID: 2, Title: "Title to Error", FamilyCategory: "FamilyCategoryTest",
			MyCategory: "MyCategoryTest", Amount: 100.0},
		{ID: 3, Title: "Error", FamilyCategory: "FamilyCategoryTest",
			MyCategory: "MyCategoryTest", Amount: 100.0},
		{ID: 4, Title: "Title to Test", FamilyCategory: "FamilyCategoryTest",
			MyCategory: "MyCategoryTest", Amount: 100.0},
		{ID: 5, Title: "Error", FamilyCategory: "FamilyCategoryTest",
			MyCategory: "MyCategoryTest", Amount: 100.0},
		{ID: 6, Title: "Title for Test", FamilyCategory: "FamilyCategoryTest",
			MyCategory: "MyCategoryTest", Amount: 100.0},
	}

	OneHundredPercentCommon := registryService.FilterByCommonTitle(registry, registries, 100.0)
	SeventyPercentCommon := registryService.FilterByCommonTitle(registry, registries, 70.0)
	FortyPercentCommon := registryService.FilterByCommonTitle(registry, registries, 40.0)

	if len(OneHundredPercentCommon) != 1  {
		test.Errorf("Expected 1 but found %d", len(OneHundredPercentCommon) )
	}

	if len(SeventyPercentCommon) != 3  {
		test.Errorf("Expected 3 but found %d", len(SeventyPercentCommon) )
	}

	if len(FortyPercentCommon) != 4  {
		test.Errorf("Expected 4 but found %d", len(FortyPercentCommon) )
	}

	if FortyPercentCommon[0].ID != 4 {
		test.Errorf("Expected id 4 but found %d", FortyPercentCommon[0].ID  )
	}

	test.Log("All tests success")



}
