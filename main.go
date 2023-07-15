package main

import "TravelGo/backend/http"

func main() {
	service := http.NewRouter()
	service.StartServer()
}
