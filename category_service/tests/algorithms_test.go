package tests

import (
	algorithms "github.com/daniel-acaz/nubank-control/category_service/commons"
	"testing"
)

func TestLongestCommonSubSequenceAlgorithm(test *testing.T) {

	firstWorld := "ATGTTAT"
	secondWorld := "ATCGTAC"

	sequence := algorithms.LongestCommonSubSequence([]byte(firstWorld), []byte(secondWorld))

	if sequence != 5{
		test.Errorf("Value should be 5 but is %d", sequence)
	}
	test.Logf("Expected 5 and found %d with success", sequence)

}


func TestLongestCommonPercentageAlgorithm(test *testing.T) {

	firstWorld := "ATGTTAT"
	secondWorld := "ATCGTAC"

	sequence := algorithms.LongestCommonSubSequencePercentage([]byte(firstWorld), []byte(secondWorld))

	if sequence < 55 || sequence > 56  {
		test.Errorf("Value should be 55.* but is %f", sequence)
	}
	test.Logf("Expected 55.* and found %f with success", sequence)

}