package algo

import (
	"encoding/json"
	"math"
	"os"
)

func readInputs(fileName string) ([][]float32, error) {  
  fileData, error := os.ReadFile(fileName)
  
  if error != nil {
    return nil, error
  }

  var data struct {
    Input [][]float32
  }

  // Unmarshal returns a error you need to pass a pointer in 2nd parameter 
  // for where the values are returned to
  error = json.Unmarshal(fileData, &data)

  if error != nil {
    return nil, error
  }

  return data.Input, nil
}

func countAverages(fileName string) ([]float32, error) {
  jsonInput, error := readInputs(fileName)

  if error != nil {
    return nil, error
  }

  returnValue := []float32{}

  for _, currentArray := range jsonInput {
    var sum float32 = 0

    for _, currentItem := range currentArray {
      sum += currentItem
    }
    
    averageOfCurrentArray := sum / 2
    
    // checking if averageOfCurrentArray is a number math.IsNaN only takes float64
    if(math.IsNaN(float64(averageOfCurrentArray))) {
      returnValue = append(returnValue, 0)
    } else {
      returnValue = append(returnValue, averageOfCurrentArray)
    }
  }

  return returnValue, nil
}