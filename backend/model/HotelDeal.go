package model

import "encoding/xml"

type Hotwire struct {
	XMLName xml.Name `xml:"Hotwire"`
	Result  Result   `xml:"Result"`
}

type Result struct {
	HotelDeals []HotelDeal `xml:"HotelDeal"`
}

type HotelDeal struct {
	FoundDate             string  `xml:"FoundDate"`
	CurrencyCode          string  `xml:"CurrencyCode"`
	NightDuration         float64 `xml:"NightDuration"`
	EndDate               string  `xml:"EndDate"`
	Headline              string  `xml:"Headline"`
	IsWeekendStay         bool    `xml:"IsWeekendStay"`
	Price                 float64 `xml:"Price"`
	StartDate             string  `xml:"StartDate"`
	Url                   string  `xml:"Url"`
	City                  string  `xml:"City"`
	CountryCode           string  `xml:"CountryCode"`
	NeighborhoodLatitude  string  `xml:"NeighborhoodLatitude"`
	NeighborhoodLongitude string  `xml:"NeighborhoodLongitude"`
	Neighborhood          string  `xml:"Neighborhood"`
	NeighborhoodId        int     `xml:"NeighborhoodId"`
	SavingsPercentage     int     `xml:"SavingsPercentage"`
	StarRating            float64 `xml:"StarRating"`
	StateCode             string  `xml:"StateCode"`
}
