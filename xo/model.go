package xo

type Game struct{
	ID int `json:"id"`
	Field string `json:"field"`
	StartPlayer int `json:"start_player"`
	Round int `json:"round"`
	Winner string `json:"winner"`
}