package http_error

type HttpError struct {
	Status  int
	Message map[string]string
}
