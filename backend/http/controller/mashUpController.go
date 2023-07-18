package controller

import (
	"TravelGo/backend/model"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type MashUpController struct{}

const HotWireApiKey = "9y2jx6mw8hgtp5ex7mfx8p9m"

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

func (c *MashUpController) GetHotelDeals(ctx *gin.Context) {
	//first request: CITY -> [longitude, latitude]
	city := ctx.Request.URL.Query().Get("city")
	startDate := ctx.Request.URL.Query().Get("start_date")
	endDate := ctx.Request.URL.Query().Get("end_date")
	baseURL := "http://api.hotwire.com/v1/deal/hotel"
	params := url.Values{}
	params.Set("apiKey", HotWireApiKey)
	params.Set("limit", "10")
	params.Set("dest", city)

	// Construct the URL with the parameters
	requestURL := baseURL + "?" + params.Encode() + "&startdate=" + startDate + "&enddate=" + endDate

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
	var results model.Hotwire
	err = xml.Unmarshal(body, &results)
	dealList := make([]map[string]string, len(results.Result.HotelDeals))
	for i, v := range results.Result.HotelDeals {
		dealList[i] = map[string]string{
			"title": v.Headline,
			"url":   v.Url,
		}
	}
	SuccessResponse(ctx, dealList)
}
