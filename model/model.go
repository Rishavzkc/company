package model

type Company struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	Location string `json:"location"`
}

type Response struct {
	StatusCode int
	Message    string
}

type ServiceError struct {
	StatusCode int
	ErrorCode  int
	Error      string
}
