package api

import "context"

//go:generate go run go.uber.org/mock/mockgen -source=api.go -destination=mockuber/api.go -package=mockuber
type WeatherApi interface {
	CurrentWeather(ctx context.Context, latitude, longitude string) (WeatherResult, error)
	ForecastNextFourDays(ctx context.Context, latitude, longitude string) (WeatherForecastResult, error)
}
