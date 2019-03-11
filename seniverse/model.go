package seniverse

import "time"

type Error struct {
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
}

type Results struct {
	Results []struct {
		Location struct {
			ID             string `json:"id"`
			Name           string `json:"name"`
			Country        string `json:"country"`
			Path           string `json:"path"`
			Timezone       string `json:"timezone"`
			TimezoneOffset string `json:"timezone_offset"`
		} `json:"location"`
		Now struct {
			Text        string `json:"text"`
			Code        string `json:"code"`
			Temperature string `json:"temperature"`
		} `json:"now"`
		LastUpdate time.Time `json:"last_update"`
	} `json:"results"`
}

//type Location struct {
//	Id             string `json:"id"`
//	Name           string `json:"name"`
//	Country        string `json:"country"`
//	Path           string `json:"path"`
//	Timezone       string `json:"timezone"`
//	TimezoneOffset string `json:"timezone_offset"`
//}
//
//type Now struct {
//	Text        string `json:"text"`
//	Code        string `json:"code"`
//	Temperature string `json:"temperature"`
//}
//
//type ResultValue struct {
//	Location   Location
//	Now        Now
//	LastUpdate string `json:"last_update"`
//}
