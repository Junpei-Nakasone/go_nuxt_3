package domain

// RequestParam contains RequestParam
type RequestParam struct {
	ID int `json:"id"`
}

// Todo contains todo
type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
