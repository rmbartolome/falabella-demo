package repository

import (
	"context"
)

type DeliveryOrderRepository interface {
	FindAll(ctx context.Context) ([]DeliveryOrder, error)
	GetByID(ctx context.Context, id string) (DeliveryOrder, error)
	Create(ctx context.Context, deliveryOrder DeliveryOrder) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (DeliveryOrder, error)
}
type deliveryOrderRepository struct {
	conn Connection
}

func NewBDeliveryOrderRepository(conn Connection) DeliveryOrderRepository {
	return &deliveryOrderRepository{
		conn: conn,
	}
}

func GetByID(ID string) (found datos.DeliveryOrder) {
	deliveryOrder := datos.DeliveryOrder{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&deliveryOrder, ID)
	return deliveryOrder
}

func Delete(ID string) (found datos.DeliveryOrder) {
	deliveryOrder := datos.DeliveryOrder{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&deliveryOrder.ID)
	db.Delete(deliveryOrder)
	return deliveryOrder
}

func Create(DeliveryOrder) (found datos.DeliveryOrder) {
	deliveryOrder := datos.DeliveryOrder{}
	db := utils.GetConnection()
	defer db.Close()
	err = db.Create(&deliveryOrder).Error
	return deliveryOrder
}

func Update(ID string) (found datos.DeliveryOrder) {
	find := datos.DeliveryOrder{}
	data := datos.DeliveryOrder{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&find, ID)
	db.Model(&find).Updates(data)

}

func FindAll() (found datos.DeliveryOrder) {
	deliveryOrder := datos.DeliveryOrder{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&deliveryOrder)
	return deliveryOrder

}
