package repository

import (
	"context"
)

type PuchaseOrderRepository interface {
	FindAll(ctx context.Context) ([]PuchaseOrder, error)
	GetByID(ctx context.Context, id string) (PuchaseOrder, error)
	Create(ctx context.Context, puchaseOrder PuchaseOrder) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (PuchaseOrder, error)
}
type puchaseOrderRepository struct {
	conn Connection
}

func NewPuchaseOrderRepository(conn Connection) PuchaseOrderRepository {
	return &puchaseOrderRepository{
		conn: conn,
	}
}
func GetByID(ID string) (found datos.PuchaseOrder) {
	puchaseOrder := datos.PuchaseOrder{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&puchaseOrder, ID)
	return puchaseOrder
}

func Delete(ID string) (found datos.PuchaseOrder) {
	puchaseOrder := datos.PuchaseOrder{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&puchaseOrder.ID)
	db.Delete(puchaseOrder)
	return puchaseOrder
}

func Create(PuchaseOrder) (found datos.PuchaseOrder) {
	puchaseOrder := datos.PuchaseOrder{}
	db := utils.GetConnection()
	defer db.Close()
	err = db.Create(&puchaseOrder).Error
	return puchaseOrder
}

func Update(ID string) (found datos.PuchaseOrder) {
	find := datos.PuchaseOrder{}
	data := datos.PuchaseOrder{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&find, ID)
	db.Model(&find).Updates(data)

}

func FindAll() (found datos.PuchaseOrder) {
	puchaseOrder := datos.PuchaseOrder{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&puchaseOrder)
	return puchaseOrder

}
