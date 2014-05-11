package accuweather

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

type Client struct {
	baseUrl string
	apiKey  string
}

var client *Client = nil

func InitClient(baseUrl, apiKey string) {
	client = &Client{
		apiKey:  apiKey,
		baseUrl: baseUrl,
	}
}

func GetLocationByCountryCity(countryCode, city string) (*MultipleLocationResponse, error) {
	var endpoint = fmt.Sprintf("locations/v1/cities/%s/search.json", countryCode)
	values := url.Values{}

	values.Set("q", city)

	response, e := client.makeRequest(endpoint, &values, reflect.TypeOf(MultipleLocationResponse{}))

	if e != nil {
		return nil, e
	}

	return response.(*MultipleLocationResponse), nil
}

func GetLocationByLatLon(lat, lon float64) (*LocationResponse, error) {
	var endpoint = "locations/v1/cities/geoposition/search.json"
	values := url.Values{}

	values.Set("q", fmt.Sprintf("%.2f,%.2f", lat, lon))

	response, e := client.makeRequest(endpoint, &values, reflect.TypeOf(LocationResponse{}))

	if e != nil {
		return nil, e
	}

	return response.(*LocationResponse), nil
}

func GetLocalWeather(locationKey, locale string, metric bool) (*LocalWeatherResponse, error) {
	var endpoint = fmt.Sprintf("localweather/v1/%s", locationKey)
	values := url.Values{}

	// TODO: check if locale is a valid one?
	values.Set("language", locale)
	values.Set("metric", fmt.Sprintf("%t", metric))

	response, e := client.makeRequest(endpoint, &values, reflect.TypeOf(LocalWeatherResponse{}))

	if e != nil {
		return nil, e
	}

	return response.(*LocalWeatherResponse), nil
}

func parseResponse(body []byte, resType reflect.Type) (interface{}, error) {
	res := reflect.New(resType).Interface()
	e := json.Unmarshal(body, res)

	if e != nil {
		return nil, fmt.Errorf("Error deserializing: %s", e)
	}

	return res, nil
}

func (c *Client) makeRequest(endpoint string, values *url.Values, responseType reflect.Type) (interface{}, error) {
	var reqUrl bytes.Buffer
	reqUrl.WriteString(fmt.Sprintf("%s/%s?", c.baseUrl, endpoint))

	values.Set("apikey", c.apiKey)

	reqUrl.WriteString(values.Encode())

	r, e := http.Get(string(reqUrl.String()))
	if e != nil {
		return nil, e
	}

	defer r.Body.Close()

	if r.StatusCode == 200 {
		body, _ := ioutil.ReadAll(r.Body)
		return parseResponse(body, responseType)
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		return nil, fmt.Errorf("Error making Accuweather reuqest; code:%d error:%v", r.StatusCode, e)
	}
}
