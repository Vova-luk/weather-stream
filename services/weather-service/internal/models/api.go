package models

type WeatherInformation struct {
	Current struct {
		TempC    float64 `json:"temp_c"`
		Humidity int32   `json:"humidity"`
		WindMph  float64 `json:"wind_mph"`
		Pressure float64 `json:"pressure_in"`
		Precip   float64 `json:"precip_mm"`
		Cloud    int32   `json:"cloud"`
	} `json:"current"`
}

func (w *WeatherInformation) ToWeatherData(locationId int32) *WeatherData {
	return &WeatherData{
		LocationId:  locationId,
		Temperature: w.Current.TempC,
		Humidity:    w.Current.Humidity,
		WindSpeed:   w.Current.WindMph,
		Pressure:    w.Current.Pressure,
		Precip:      w.Current.Precip,
		Cloud:       w.Current.Cloud,
	}
}
