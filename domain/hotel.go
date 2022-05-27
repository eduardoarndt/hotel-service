package domain

type Hotel struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	City    string  `json:"city"`
	Reviews int     `json:"reviews"`
	Rating  float32 `json:"rating"`
}
