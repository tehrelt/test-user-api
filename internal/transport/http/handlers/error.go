package handlers

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
