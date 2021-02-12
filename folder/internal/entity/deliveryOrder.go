package entity

import (
	"errors"
	"html"
	"strings"

	"github.com/gofrs/uuid"
)

type DeliveryOrder struct {
	ID          string `gorm:"primary_key" json:"id"`
	CodBill     string `gorm:"size:255;not null" json:"codBill"`
	Address     string `gorm:"size:255;not null" json:"address"`
	Responsable string `gorm:"size:255;not null" json:"responsable"`
}

func (d *DeliveryOrder) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.CodBill = html.EscapeString(strings.TrimSpace(d.CodBill))
	d.Responsable = html.EscapeString(strings.TrimSpace(d.Responsable))
	d.Address = html.EscapeString(strings.TrimSpace(d.Address))

}

func (d *DeliveryOrder) Validate() error {
	if d.CodBill == "" {
		return errors.New("required bill code")
	}
	if d.Address == "" {
		return errors.New("required address")
	}
	if d.Responsable == "" {
		return errors.New("required responsable")
	}

	return nil
}

type FindAllDeliveryOrderRequest struct {
}

type GetByIDDeliveryOrderRequest struct {
	ID string `json:"id"`
}

type CreateDeliveryOrderRequest struct {
	DeliveryOrder DeliveryOrder `json:"deliveryOrder"`
}

type DeleteDeliveryOrderRequest struct {
	ID string `json:"id"`
}

type FindAllDeliveryOrderResponse struct {
	DeliveryOrders []DeliveryOrder `json:"deliveryOrders"`
	Err            error           `json:"error,omitempty"`
}

func (r FindAllDeliveryOrderResponse) error() error { return r.Err }

type GetByIDDeliveryOrderResponse struct {
	DeliveryOrder DeliveryOrder `json:"deliveryOrder"`
	Err           error         `json:"error,omitempty"`
}

func (r GetByIDDeliveryOrderResponse) error() error { return r.Err }

type CreateDeliveryOrderResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateDeliveryOrderResponse) error() error { return r.Err }

type DeleteDeliveryOrderResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteDeliveryOrderResponse) error() error { return r.Err }
