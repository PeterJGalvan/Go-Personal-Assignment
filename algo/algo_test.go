package algo

import (
	"reflect"
	"testing"
)

func Test_countAverages(t *testing.T) {
  testCases := []struct {
    name string
    filePath string
    expected []float32
    isError  bool
  }{
    {
      name: "Good Case",
      filePath: "./inputs/good.json",
      expected: []float32{138.5, 0, 66, 237},
      isError: false,
    },
    {
      name: "Bad Case",
      filePath: "./inputs/bad.json",
      isError:  true,
    },
  }

  for _, currentTestCase := range testCases {
    t.Run(currentTestCase.name, func(t *testing.T) {
      inputs, error := countAverages(currentTestCase.filePath)

      // checking if there is a error and there should be no error
      if(error != nil && !currentTestCase.isError) {
        t.Errorf("Expected No Errors, got: %v", error)
      }

      // checking if there is no errors and there should be a error
      if(error == nil && currentTestCase.isError) {
        t.Errorf("Expected No Errors, got: %v", error)
      }

      // checking for expected output to equal returned inputs
      if !reflect.DeepEqual(inputs, currentTestCase.expected) {
        t.Errorf("Expected inputs: %v, got: %v", currentTestCase.expected, inputs)
      }
    })
  }
}


func Test_readInputs(t *testing.T) {
  testCases := []struct {
    name     string
    filePath string
    expected [][]float32
    isError  bool
  }{
    {
      name: "Good Case",
      filePath: "./inputs/good.json",
      expected: [][]float32{
        {87, 23, 57, 87, 23},
        {},
        {75, 34, 23},
        {23, 45, 83, 239, 28, 56},
      },
      isError: false,
    },
    {
      name: "Bad Case",
      filePath: "./inputs/bad.json",
      isError: true,
    },
  }

  for _, currentTestCase := range testCases {
    t.Run(currentTestCase.name, func(t *testing.T) {
      inputs, error := readInputs(currentTestCase.filePath)

      // checking if there is a error and there should be no error
      if(error != nil && !currentTestCase.isError) {
        t.Errorf("Expected No Errors, got: %v", error)
      }

      // checking if there is no errors and there should be a error
      if(error == nil && currentTestCase.isError) {
        t.Errorf("Expected No Errors, got: %v", error)
      }

      // checking for expected output to equal returned inputs
      if !reflect.DeepEqual(inputs, currentTestCase.expected) {
        t.Errorf("Expected inputs: %v, got: %v", currentTestCase.expected, inputs)
      }
    })
	}
}
