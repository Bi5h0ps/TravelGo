package model

type PixabayResponse struct {
	Total     int            `json:"total"`
	TotalHits int            `json:"totalHits"`
	Hits      []PixabayImage `json:"hits"`
}

type PixabayImage struct {
	ID              int    `json:"id"`
	PageURL         string `json:"pageURL"`
	Type            string `json:"type"`
	Tags            string `json:"tags"`
	PreviewURL      string `json:"previewURL"`
	PreviewWidth    int    `json:"previewWidth"`
	PreviewHeight   int    `json:"previewHeight"`
	WebformatURL    string `json:"webformatURL"`
	WebformatWidth  int    `json:"webformatWidth"`
	WebformatHeight int    `json:"webformatHeight"`
	LargeImageURL   string `json:"largeImageURL"`
	ImageWidth      int    `json:"imageWidth"`
	ImageHeight     int    `json:"imageHeight"`
	ImageSize       int    `json:"imageSize"`
	Views           int    `json:"views"`
	Downloads       int    `json:"downloads"`
	Collections     int    `json:"collections"`
	Likes           int    `json:"likes"`
	Comments        int    `json:"comments"`
	UserID          int    `json:"user_id"`
	User            string `json:"user"`
	UserImageURL    string `json:"userImageURL"`
}
