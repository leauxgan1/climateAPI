package types


type MonthlyClimateReading struct {
  Month string `json:"month"`
  Temperature float64 `json:"temperature"`
}

type ClimateYear struct {
  CountryName string `json:"country-name"`
  Year int `json:"year"`
  Readings  []MonthlyClimateReading `json:"readings"`
}

func NewClimateYear(country string, year int) *ClimateYear {
  monthlyReadings := make([]MonthlyClimateReading,12)
  return &ClimateYear{
    country,
    year,
    monthlyReadings,
  }
}

var Store = []ClimateYear{
  {
    "India",
    2023,
    []MonthlyClimateReading{
      {
        
        "January",
        32.0,
      },
      {
        
        "February",
        34.0,
      },
      {
        
        "March",
        38.0,
      },
      {
        
        "April",
        40.0,
      },
      {
        
        "May",
        45.0,
      },
      {
        
        "June",
        48.0,
      },
      {
        
        "July",
        50.0,
      },
      {
        
        "August",
        52.0,
      },
      {
        
        "September",
        47.0,
      },
      {
        
        "October",
        42.0,
      },
      {
        
        "November",
        38.0,
      },
      {
        
        "December",
        32.0,
      },
    },
  },
}


