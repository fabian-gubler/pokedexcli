package models

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
