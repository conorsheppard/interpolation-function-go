package interpolation_function

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_InterpolationFunction(t *testing.T) {
	tests := []struct {
		inputJSON      string
		inputMap       map[string]string
		expectedResult string
	}{
		{
			inputJSON:      "test/assets/input/test-input-1.json",
			inputMap:       map[string]string{"customerFirstName": "John", "catName": "Mittens"},
			expectedResult: "test/assets/result/test-result-1.json",
		},
		{
			inputJSON:      "test/assets/input/test-input-2.json",
			inputMap:       map[string]string{"customerFirstName": "Dave", "catName": "Felix", "meal1": "KatKin Chicken Nibbles", "meal2": "KatKin Splash Whitefish"},
			expectedResult: "test/assets/result/test-result-2.json",
		},
	}

	for i, test := range tests {
		fmt.Printf("test: %d\n", i)

		inputContent, expectedResult := readFile(test.inputJSON), readFile(test.expectedResult)

		result := interpolate(inputContent, test.inputMap)

		if result != expectedResult {
			t.Errorf("FAIL: expected %s, got %s\n", expectedResult, result)
		} else {
			fmt.Printf("\tPASS\n")
		}
	}
}

func readFile(filename string) string {
	inputBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("fatal: failed to read test file " + filename)
	}
	return string(inputBytes)
}
