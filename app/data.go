package main

// Message json object mapper
type Message struct {
	Username string `json:"username"`
	Body     string `json:"body"`
	Color    struct {
		R int `json:"r"`
		G int `json:"g"`
		B int `json:"b"`
	}
}
