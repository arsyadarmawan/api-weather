package api

import "context"

//go:generate go run go.uber.org/mock/mockgen -source=api.go -destination=mockuber/api.go -package=mockuber
type WeatherApi interface {
	CurrentWeather(ctx context.Context, apiKey, latitude, longitude string) (WeatherResult, error)
}
