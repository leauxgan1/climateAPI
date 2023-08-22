
package types 

import (
  "fmt"
  "testing"
)


func TestClimateYear(t *testing.T) {
  mockTestingYear := NewClimateYear("Somalia", 1984) 

  fmt.Println(mockTestingYear.CountryName)
  fmt.Println(mockTestingYear.Year)
  for i := range mockTestingYear.Readings {
    fmt.Println(mockTestingYear.Readings[i])
  }

}
