package weather

// http://weather.livedoor.com/weather_hacks/webservice

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const baseAPIURL = "http://weather.livedoor.com/forecast/webservice/json/v1"

// PublicTime type
type PublicTime struct {
	Time time.Time
}

// UnmarshalJSON method
func (pt *PublicTime) UnmarshalJSON(data []byte) error {
	result, err := time.Parse(`"2006-01-02T15:04:05-0700"`, strings.Replace(string(data), "\\u002b", "+", -1))
	if err != nil {
		return err
	}
	pt.Time = result
	return nil
}

// Result type
type Result struct {
	Title       string      `json:"title"`
	Description Description `json:"description"`
	Link        string      `json:"link"`
	PublicTime  *PublicTime `json:"publicTime"`
	Location    Location    `json:"location"`
	Forecasts   []Forecast  `json:"forecasts"`
}

// Description type
type Description struct {
	Text       string      `json:"text"`
	PublicTime *PublicTime `json:"publicTime"`
}

// Location type
type Location struct {
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
	City       string `json:"city"`
}

// Forecast type
type Forecast struct {
	Date        string      `json:"date"`
	DateLabel   string      `json:"dateLabel"`
	Telop       string      `json:"telop"`
	Temperature Temperature `json:"temperature"`
	Image       Image       `json:"image"`
}

// Temperature type
type Temperature struct {
	Min *TemperatureValue `json:"min"`
	Max *TemperatureValue `json:"max"`
}

// TemperatureValue type
type TemperatureValue struct {
	Celsius    string `json:"celsius"`
	Fahrenheit string `json:"fahrenheit"`
}

// Image type
type Image struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// FetchForecast function
func FetchForecast(cityID int) (*Result, error) {
	u, err := url.ParseRequestURI(baseAPIURL)
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Set("city", strconv.Itoa(cityID))
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := &Result{}
	if json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}
