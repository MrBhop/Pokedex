package pokeapi

type MapResponse struct {
	Count    int       `json:"count"`
	Next     *string    `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
