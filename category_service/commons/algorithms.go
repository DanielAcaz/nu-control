package commons

import (
	lcs "github.com/yudai/golcs"
)

func LongestCommonSubSequencePercentage(x, y []byte) float64 {

	subSequence := LongestCommonSubSequence(x, y)
	xLength := len(x)
	yLength := len(y)

	return (float64(subSequence) / float64(subSequence+(xLength+yLength-2*subSequence))) * 100.0

}

func LongestCommonSubSequence(leftBytes, rightBytes []byte) int {

	var left = make([]interface{}, len(leftBytes))
	for i, v := range leftBytes {
		left[i] = v
	}

	var right = make([]interface{}, len(rightBytes))
	for i, v := range rightBytes {
		right[i] = v
	}

	return lcs.New(left, right).Length()

}
