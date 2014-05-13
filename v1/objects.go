package accuweather

import "fmt"

//placeholder for a struct we don't need so we're not adding right now...
type TBD interface{}

func IconUrl(icon int) string {
	return fmt.Sprintf("http://api.accuweather.com/developers/Media/Default/WeatherIcons/%02d-s.png", icon)
}

// Location Data structs:
type Region struct {
	ID            string
	LocalizedName string
	EnglishName   string
}

type Country struct {
	ID            string
	LocalizedName string
	EnglishName   string
}

type AdministrativeArea struct {
	ID            string
	LocalizedName string
	EnglishName   string
	Level         int
	LocalizedType string
	EnglishType   string
	CountryID     string
}

type TimeZone struct {
	Code             string
	Name             string
	GmtOffset        float64
	IsDaylightSaving bool
	NextOffsetChange string
}

type ElevationDetails struct {
	Value    float64
	Unit     string
	UnitType int
}

type Elevation struct {
	Metric   ElevationDetails
	Imperial ElevationDetails
}

type GeoPosition struct {
	Latitude  float64
	Longitude float64
	Elevation Elevation
}

type Location struct {
	Version                int
	Key                    string
	Type                   string
	Rank                   int
	LocalizedName          string
	EnglishName            string
	PrimaryPostalCode      string
	Region                 Region
	Country                Country
	AdministrativeArea     AdministrativeArea
	TimeZone               TimeZone
	GeoPosition            GeoPosition
	IsAlias                bool
	SupplementalAdminAreas TBD
}

// Weather data structs
type MapImage struct {
	Date string
	Url  string
}

type SatelliteImages struct {
	Images []MapImage
	Size   string
}

type Maps struct {
	link       string
	Satellite  SatelliteImages
	MobileLink string
	Radar      TBD
}

type Temperature struct {
	UnitType int
	Unit     string
	Value    float64
}

type TemperatureMinMax struct {
	Minimum Temperature
	Maximum Temperature
}

type Icon struct {
	IconPhrase string
	Icon       int
}

type Conditions struct {
	Temperature       Temperature
	TemperatureMinMax TemperatureMinMax

	IsDaylight    bool
	DateTime      string
	EpochDateTime int
	IconPhrase    string
	WeatherIcon   int

	Link       string
	MobileLink string
}

type CurrentConditions struct {
	Temperature Temperature

	IsDayTime                bool
	LocalObservationDateTime string
	EpochTime                int
	WeatherText              string
	WeatherIcon              int

	Link       string
	MobileLink string
}

func (c *CurrentConditions) GetConditions() *Conditions {
	var conditions = &Conditions{
		Temperature:   c.Temperature,
		IsDaylight:    c.IsDayTime,
		DateTime:      c.LocalObservationDateTime,
		EpochDateTime: c.EpochTime,
		IconPhrase:    c.WeatherText,
		WeatherIcon:   c.WeatherIcon,
		Link:          c.Link,
		MobileLink:    c.MobileLink,
	}

	return conditions
}

type HourlyForecast struct {
	Conditions

	PrecipitationProbability float64
}

type DailyForecast struct {
	Temperature TemperatureMinMax
	Sources     []string
	Night       Icon
	Day         Icon
	EpochDate   int
	Date        string
	Link        string
	MobileLink  string
}

func (df *DailyForecast) GetConditions() *Conditions {
	var conditions = &Conditions{
		Temperature:       df.Temperature.Maximum,
		TemperatureMinMax: df.Temperature,
		DateTime:          df.Date,
		EpochDateTime:     df.EpochDate,
		IconPhrase:        df.Day.IconPhrase,
		WeatherIcon:       df.Day.Icon,
		Link:              df.Link,
		MobileLink:        df.MobileLink,
	}

	return conditions
}

type Headline struct {
	Category           string
	EndDate            string
	Severity           int
	EffectiveDate      string
	Text               string
	EndEpochDate       int
	EffectiveEpochDate int
	Link               string
	MobileLink         string
}

type ForecastSummary struct {
	Headline        Headline
	DailyForecasts  []DailyForecast
	HourlyForecasts []HourlyForecast
}

// Responses definitions
type LocalWeatherResponse struct {
	Maps              Maps
	Location          Location
	CurrentConditions CurrentConditions
	ForecastSummary   ForecastSummary
}

type MultipleLocationResponse []Location
type LocationResponse Location

// Supported language codes list. Exported on 13/5/2014 using:
// $ curl "http://apidev.accuweather.com/translations/v1/languages/?apikey={your key}" | jq ".[].ISO"
var SupportedLanguages = map[string]bool{
	"ar":    true,
	"ar-ae": true,
	"ar-bh": true,
	"ar-dz": true,
	"ar-eg": true,
	"ar-iq": true,
	"ar-jo": true,
	"ar-kw": true,
	"ar-lb": true,
	"ar-ly": true,
	"ar-ma": true,
	"ar-om": true,
	"ar-qa": true,
	"ar-sa": true,
	"ar-sy": true,
	"ar-tn": true,
	"ar-ye": true,
	"bg":    true,
	"ca":    true,
	"cs":    true,
	"da":    true,
	"de":    true,
	"de-at": true,
	"de-ch": true,
	"de-li": true,
	"de-lu": true,
	"el":    true,
	"en":    true,
	"en-au": true,
	"en-bz": true,
	"en-ca": true,
	"en-gb": true,
	"en-ie": true,
	"en-jm": true,
	"en-nz": true,
	"en-tt": true,
	"en-us": true,
	"en-za": true,
	"es":    true,
	"es-ar": true,
	"es-bo": true,
	"es-cl": true,
	"es-co": true,
	"es-cr": true,
	"es-do": true,
	"es-ec": true,
	"es-gt": true,
	"es-hn": true,
	"es-mx": true,
	"es-ni": true,
	"es-pa": true,
	"es-pe": true,
	"es-pr": true,
	"es-py": true,
	"es-sv": true,
	"es-uy": true,
	"es-ve": true,
	"et":    true,
	"fa":    true,
	"fi":    true,
	"fil":   true,
	"fr":    true,
	"fr-be": true,
	"fr-ca": true,
	"fr-ch": true,
	"fr-lu": true,
	"he":    true,
	"hi":    true,
	"hr":    true,
	"hu":    true,
	"id":    true,
	"it":    true,
	"it-ch": true,
	"ja":    true,
	"kk":    true,
	"ko":    true,
	"lt":    true,
	"lv":    true,
	"mk":    true,
	"ms":    true,
	"nl":    true,
	"nl-be": true,
	"no":    true,
	"ph":    true,
	"pl":    true,
	"pt":    true,
	"pt-br": true,
	"ro":    true,
	"ro-mo": true,
	"ru":    true,
	"ru-mo": true,
	"sk":    true,
	"sl":    true,
	"sr":    true,
	"sv":    true,
	"sv-fi": true,
	"th":    true,
	"tl":    true,
	"tl-ph": true,
	"tr":    true,
	"uk":    true,
	"vi":    true,
	"zh":    true,
	"zh-cn": true,
	"zh-hk": true,
	"zh-sg": true,
	"zh-tw": true,
}
