package controller

import (
	"TravelGo/backend/model"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type MashUpController struct{}

const ApiKey = "38340278-86c45a30281af96820e2b3f29"

func (c *MashUpController) GetCityTemp(ctx *gin.Context) {
	//first request: CITY -> [longitude, latitude]
	city := ctx.Request.URL.Query().Get("city")
	baseURL := "https://geocoding-api.open-meteo.com/v1/search"
	params := url.Values{}
	params.Set("name", city)

	// Construct the URL with the parameters
	requestURL := baseURL + "?" + params.Encode()

	// Send a GET request to the constructed URL
	resp, err := http.Get(requestURL)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	var results model.LocationResults
	err = json.Unmarshal(body, &results)
	if err != nil || len(results.Results) == 0 {
		ErrorResponse(ctx, errors.New("no location found"))
		return
	}
	longitude := results.Results[0].Longitude
	latitude := results.Results[0].Latitude

	//second request: [longitude, latitude] -> Temperature
	forecastUrl := "https://api.open-meteo.com/v1/forecast"
	forecastParams := url.Values{}
	forecastParams.Set("longitude", strconv.FormatFloat(longitude, 'f', 4, 64))
	forecastParams.Set("latitude", strconv.FormatFloat(latitude, 'f', 4, 64))
	forecastParams.Set("timezone", "GMT")
	forecastParams.Set("daily", "temperature_2m_max")
	forecastParams.Add("daily", "temperature_2m_min")
	forecastParams.Set("forecast_days", "7")
	// Construct the URL with the parameters
	requestURL = forecastUrl + "?" + forecastParams.Encode()

	// Send a GET request to the constructed URL
	resp, err = http.Get(requestURL)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	// Read the response body
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	var forecastResult model.WeatherData
	err = json.Unmarshal(body, &forecastResult)
	if err != nil {
		ErrorResponse(ctx, errors.New("something went wrong"))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"date":     forecastResult.Daily.Time,
		"temp_max": forecastResult.Daily.Temperature2mMax,
		"temp_min": forecastResult.Daily.Temperature2mMin,
	})
}

func (c *MashUpController) GetCityPics(ctx *gin.Context) {
	city := ctx.Request.URL.Query().Get("city")
	baseURL := "https://pixabay.com/api/"
	params := url.Values{}
	params.Set("key", ApiKey)
	params.Set("q", city)
	params.Set("image_type", "photo")
	params.Set("per_page", "5")

	// Construct the URL with the parameters
	requestURL := baseURL + "?" + params.Encode()

	// Send a GET request to the constructed URL
	resp, err := http.Get(requestURL)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	var results model.PixabayResponse
	err = json.Unmarshal(body, &results)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	imgList := make([]string, len(results.Hits))
	for i, v := range results.Hits {
		imgList[i] = v.WebformatURL
	}
	SuccessResponse(ctx, imgList)
}
