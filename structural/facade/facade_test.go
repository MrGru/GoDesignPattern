package facade

import (
	"fmt"
	"testing"
)

func TestOpenWeatherMap_responseParse(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{APIkey: ""}
	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}
	if weather.ID != 3117735 {
		t.Errorf("Madrid id is 3117735, not %d\n", weather.ID)
	}

	weatherMap := CurrentWeatherData{APIkey: "913064a5171d70ea4f82d5ebfa784352"}
	weather, err = weatherMap.GetByCityAndCountryCode("Hanoi", "VN")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Temperature in Ha Noi is %f celsius\n", weather.Main.Temp-273.15)
}
