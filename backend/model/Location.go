package model

type Location struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Elevation   float64 `json:"elevation"`
	FeatureCode string  `json:"feature_code"`
	CountryCode string  `json:"country_code"`
	Admin1ID    int     `json:"admin1_id"`
	Admin2ID    int     `json:"admin2_id"`
	Admin3ID    int     `json:"admin3_id,omitempty"`
	Timezone    string  `json:"timezone"`
	Population  int     `json:"population,omitempty"`
	CountryID   int     `json:"country_id"`
	Country     string  `json:"country"`
	Admin1      string  `json:"admin1"`
	Admin2      string  `json:"admin2"`
	Admin3      string  `json:"admin3,omitempty"`
}

type LocationResults struct {
	Results          []Location `json:"results"`
	GenerationTimeMS float64    `json:"generationtime_ms"`
}
