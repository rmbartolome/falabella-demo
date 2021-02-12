package entity

import (
	"errors"
	"html"
	"strings"

	"github.com/gofrs/uuid"
)

type Product struct {
	ID          string  `gorm:"primary_key" json:"id"`
	CodTypeProd string  `gorm:"size:255;not null" json:"codtypeprod"`
	Name        string  `gorm:"size:255;not null" json:"name"`
	Maker       string  `gorm:"size:255;not null" json:"maker"`
	Model       string  `gorm:"size:255;not null" json:"model"`
	Year        string  `gorm:"size:255;not null" json:"year"`
	Price       float64 `json:"price"`
}

func (d *Product) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.CodTypeProd = html.EscapeString(strings.TrimSpace(d.CodTypeProd))
	d.Name = html.EscapeString(strings.TrimSpace(d.Name))
	d.Maker = html.EscapeString(strings.TrimSpace(d.Maker))
	d.Model = html.EscapeString(strings.TrimSpace(d.Model))
	d.Year = html.EscapeString(strings.TrimSpace(d.Year))

}

func (d *Product) Validate() error {
	if d.CodTypeProd == "" {
		return errors.New("required code type product")
	}
	if d.Maker == "" {
		return errors.New("required code maker")
	}
	if d.Model == "" {
		return errors.New("required code model")
	}
	if d.Name == "" {
		return errors.New("required code name")
	}
	if d.Year == "" {
		return errors.New("required code year")
	}
	if d.Price == 0 {
		return errors.New("required code price")
	}

	return nil
}

type FindAllProductRequest struct {
}

type GetByIDProductRequest struct {
	ID string `json:"id"`
}

type CreateProductRequest struct {
	Product Product `json:"product"`
}

type DeleteProductRequest struct {
	ID string `json:"id"`
}

type FindAllProductResponse struct {
	Products []Product `json:"products"`
	Err      error     `json:"error,omitempty"`
}

func (r FindAllProductResponse) error() error { return r.Err }

type GetByIDProductResponse struct {
	Product Product `json:"product"`
	Err     error   `json:"error,omitempty"`
}

func (r GetByIDProductResponse) error() error { return r.Err }

type CreateProductResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateProductResponse) error() error { return r.Err }

type DeleteProductResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteProductResponse) error() error { return r.Err }
