package accuweather

//placeholder for a struct we don't need so we're not adding right now...
type TBD interface{}

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
	Temperature Temperature

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
	Link        string
	Night       Icon
	Day         Icon
	Date        string
	EpochDate   int
	MobileLink  string
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
