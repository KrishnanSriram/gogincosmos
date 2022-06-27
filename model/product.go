package model

import (
	"errors"

	"github.com/google/uuid"
)

type Product struct {
  Id string             `json:"id,omitempty" validate:"required"`
  Name string           `json:"name,omitempty" validate:"required"`
  Description string    `json:"description,omitempty" validate:"required"`
}

type ProductRequest struct {
  Name string         `json:"name" validate:"required"`
  Description string  `json:"description" validate:"required"`
}

type ProductResponse struct {
  Status int
  Data interface{}
  Message string
}

func (pr *ProductRequest)ToProduct() *Product {
  return &Product{ Id: uuid.NewString(), Name: pr.Name, Description: pr.Description}
}

func (p *Product)ToProductResponse(status int, message string) *ProductResponse {
 return &ProductResponse{ Data: p, Status: status, Message: message}
}

func (pr ProductRequest)BindError() error {
  prerror := errors.New("ERROR: Name, Description are required fields. Please check your request")
  return prerror
}