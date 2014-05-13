package accuweather

import (
	"testing"
)

func init() {
	InitClient("http://apidev.accuweather.com", "dev")
}

func Test_GetLocationByCity(t *testing.T) {
	location, e := GetLocationByCountryCity("IL", "Tel Aviv")

	if e != nil {
		t.Fatalf("Got error while resolving location key: %s", e)
	}

	if len((*location)) == 0 {
		t.Fatalf("Got empty response")
	}

	var expected = "215854"
	if (*location)[0].Key != expected {
		t.Fatalf("Got wrong key (got: %s, expected: %s)", (*location)[0].Key, expected)
	}
}

func Test_GetLocationByLatLon(t *testing.T) {
	location, e := GetLocationByLatLon(32.05, 34.7)

	if e != nil {
		t.Fatalf("Got error while resolving location key: %s", e)
	}

	var expected = "215781"
	if (*location).Key != expected {
		t.Fatalf("Got wrong key (got: %s, expected: %s)", (*location).Key, expected)
	}
}

func Test_GetLocalWeather(t *testing.T) {
	weather, e := GetLocalWeather("215854", "en", false)

	if e != nil {
		t.Fatalf("Got error while requesting local weather: %s", e)
	}

	if weather.CurrentConditions.Temperature.Unit != "F" {
		t.Fatalf("Expected imperial units, but got metric (%s)", weather.CurrentConditions.Temperature.Unit)
	}
}

func testLocale(testValue, expected string, t *testing.T) {
	var actual = verifyLocale(testValue)

	if actual != expected {
		t.Fatalf("verifyLocale returned wrong locale (expected: %s, actual: %s)", expected, actual)
	}
}

func Test_verifyLocaleWithGoodLocale(t *testing.T) {
	testLocale("en-us", "en-us", t)
}

func Test_verifyLocaleWithGoodLocaleUpperCase(t *testing.T) {
	testLocale("en-US", "en-us", t)
}

func Test_verifyLocaleWithGoodLanguageBadCountryCode(t *testing.T) {
	testLocale("en-IL", "en", t)
}

func Test_verifyLocaleWithUnknownLocale(t *testing.T) {
	testLocale("zz-ZZ", "en", t)
}

func Test_GetLocalWeatherMetric(t *testing.T) {
	weather, e := GetLocalWeather("215854", "en", true)

	if e != nil {
		t.Fatalf("Got error while requesting local weather: %s", e)
	}

	if weather.CurrentConditions.Temperature.Unit != "C" {
		t.Fatalf("Expected metric units, but got imperial (%s)", weather.CurrentConditions.Temperature.Unit)
	}
}
