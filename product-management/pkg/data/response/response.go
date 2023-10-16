package response

import "product-management/pkg/models"

type Response struct {
	Message   string           `json:"message,omitempty"`
	Data      []models.Product `json:"data,omitempty"`
	Error     string           `json:"error,omitempty"`
	PageCount string           `json:"page_count,omitempty"`
}
