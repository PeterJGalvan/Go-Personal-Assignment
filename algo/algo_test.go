package algo

import (
	"reflect"
	"testing"
)

func Test_countAverages(t *testing.T) {
  testCases := []struct {
    name string
    filePath string
    want []float32
    isError  bool
  }{
    {
      name: "Good Case Test_countAverages",
      filePath: "./inputs/good.json",
      want: []float32{138.5, 0, 66, 237},
      isError: false,
    },
    {
      name: "Bad Case Test_countAverages",
      filePath: "./inputs/bad.json",
      isError:  true,
    },
  }

  for _, currentTestCase := range testCases {
    t.Run(currentTestCase.name, func(t *testing.T) {
      got, error := countAverages(currentTestCase.filePath)

      // checking if there is a error and there should be no error
      if(error != nil && !currentTestCase.isError) {
        t.Errorf("Expected No Errors, got: %v", error)
      }

      // checking if there is no errors and there should be a error
      if(error == nil && currentTestCase.isError) {
        t.Errorf("Expected No Errors, got: %v", error)
      }

      // checking for expected output to equal returned inputs used deep equal 
      // because it is a comparing slices
      if !reflect.DeepEqual(got, currentTestCase.want) {
        t.Errorf("Expected inputs: %v, got: %v", currentTestCase.want, got)
      }
    })
  }
}


func Test_readInputs(t *testing.T) {
  testCases := []struct {
    name     string
    filePath string
    want [][]float32
    isError  bool
  }{
    {
      name: "Good Case Test_readInputs",
      filePath: "./inputs/good.json",
      want: [][]float32{
        {87, 23, 57, 87, 23},
        {},
        {75, 34, 23},
        {23, 45, 83, 239, 28, 56},
      },
      isError: false,
    },
    {
      name: "Bad Case Test_readInputs",
      filePath: "./inputs/bad.json",
      isError: true,
    },
  }

  for _, currentTestCase := range testCases {
    t.Run(currentTestCase.name, func(t *testing.T) {
      got, error := readInputs(currentTestCase.filePath)

      // checking if there is a error and there should be no error
      if(error != nil && !currentTestCase.isError) {
        t.Errorf("Expected No Errors, got: %v", error)
      }

      // checking if there is no errors and there should be a error
      if(error == nil && currentTestCase.isError) {
        t.Errorf("Expected No Errors, got: %v", error)
      }

      // checking for expected output to equal returned inputs used deep equal 
      // because it is a comparing slices
      if !reflect.DeepEqual(got, currentTestCase.want) {
        t.Errorf("Expected inputs: %v, got: %v", currentTestCase.want, got)
      }
    })
	}
}
