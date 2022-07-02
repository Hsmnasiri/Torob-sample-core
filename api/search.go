package api

type searchInput struct {
	Types string `json:"types"`
	Price uint   `json:"price"`
}
