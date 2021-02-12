package entity

import (
	"errors"
	"html"
	"strings"

	"github.com/gofrs/uuid"
)

type Stock struct {
	ID            string  `gorm:"primary_key" json:"id"`
	ProductName   string  `gorm:"size:255;not null" json:"productname"`
	TypeProdName  string  `gorm:"size:255;not null" json:"typeprodname"`
	QuantityExist float64 `json:"quantityexist"`
	PurchasePrice float64 `json:"purchaseprice"`
	SalePrice     float64 `json:"saleprice"`
}

func (d *Stock) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.ProductName = html.EscapeString(strings.TrimSpace(d.ProductName))
	d.TypeProdName = html.EscapeString(strings.TrimSpace(d.TypeProdName))

}

func (d *Stock) Validate() error {
	if d.ProductName == "" {
		return errors.New("required type user")
	}
	if d.TypeProdName == "" {
		return errors.New("required type product name")
	}
	if d.QuantityExist == 0 {
		return errors.New("required quantity exist")
	}
	if d.PurchasePrice == 0 {
		return errors.New("required purchase price")
	}
	if d.SalePrice == 0 {
		return errors.New("required sales price")
	}

	return nil
}

type FindAllStockRequest struct {
}

type GetByIDStockRequest struct {
	ID string `json:"id"`
}

type CreateStockRequest struct {
	Stock Stock `json:"stock"`
}

type DeleteStockRequest struct {
	ID string `json:"id"`
}

type FindAllStockResponse struct {
	Stocks []Stock `json:"stocks"`
	Err    error   `json:"error,omitempty"`
}

func (r FindAllStockResponse) error() error { return r.Err }

type GetByIDStockResponse struct {
	Stock Stock `json:"stock"`
	Err   error `json:"error,omitempty"`
}

func (r GetByIDStockResponse) error() error { return r.Err }

type CreateStockResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateStockResponse) error() error { return r.Err }

type DeleteStockResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteStockResponse) error() error { return r.Err }
