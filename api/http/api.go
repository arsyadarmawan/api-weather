package http

import (
	"context"
	"docker/api"
	"docker/http"
	"encoding/json"
	"errors"
	http2 "net/http"
)

type HttpWeatherAPI struct {
	client *http.Client
}

func NewHttpWeatherAPI(client *http.Client) *HttpWeatherAPI {
	return &HttpWeatherAPI{client: client}
}

func (h HttpWeatherAPI) CurrentWeather(ctx context.Context, apiKey, latitude, longitude string) (result api.WeatherResult, err error) {
	req, err := h.client.SetAPIKey(ctx, apiKey).SetQueryParams(map[string]string{
		Latitude:  latitude,
		Longitude: longitude,
	}).Get(CurrentWeatherPath)
	if err != nil {
		return
	}

	switch req.StatusCode() {
	case http2.StatusOK:
		if err = json.Unmarshal(req.Body(), &result); err != nil {
			return
		}
	case http2.StatusBadRequest:
		err = errors.New("invalid request")
	case http2.StatusInternalServerError:
		err = errors.New("invalid request")
	}

	return
}
