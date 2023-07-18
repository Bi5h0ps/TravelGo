package model

type WeatherData struct {
	Latitude             float64    `json:"latitude"`
	Longitude            float64    `json:"longitude"`
	GenerationTimeMS     float64    `json:"generationtime_ms"`
	UTCOffsetSeconds     int        `json:"utc_offset_seconds"`
	Timezone             string     `json:"timezone"`
	TimezoneAbbreviation string     `json:"timezone_abbreviation"`
	Elevation            float64    `json:"elevation"`
	DailyUnits           DailyUnits `json:"daily_units"`
	Daily                DailyData  `json:"daily"`
}

type DailyUnits struct {
	Time             string `json:"time"`
	Temperature2mMax string `json:"temperature_2m_max"`
	Temperature2mMin string `json:"temperature_2m_min"`
}

type DailyData struct {
	Time             []string  `json:"time"`
	Temperature2mMax []float64 `json:"temperature_2m_max"`
	Temperature2mMin []float64 `json:"temperature_2m_min"`
}
