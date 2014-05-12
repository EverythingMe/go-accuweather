package accuweather

import (
	"testing"
)

func compare(t *testing.T, a, b interface{}, failure string) {
	if a != b {
		t.Fatalf(failure, a, b)
	}
}

func Test_GetConditions(t *testing.T) {
	var current = CurrentConditions{
		Temperature: Temperature{
			Value:    10,
			UnitType: 0,
			Unit:     "F",
		},
		IsDayTime:                true,
		LocalObservationDateTime: "2014",
		EpochTime:                10,
		WeatherText:              "Text",
		WeatherIcon:              1,
		Link:                     "http://...",
		MobileLink:               "http://...",
	}

	var conditions = current.GetConditions()

	compare(t, conditions.Temperature, current.Temperature, "Temperature not equal: %+v != %+v")
}
