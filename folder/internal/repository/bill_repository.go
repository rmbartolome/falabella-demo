package repository

import (
	"context"
	"errors"

	"encoding/json"
	"fmt"
	"net/http"
)

type BillRepository interface {
	FindAll(ctx context.Context) ([]Bill, error)
	GetByID(ctx context.Context, id string) (Bill, error)
	Create(ctx context.Context, bill Bill) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (Bill, error)
}
type billRepository struct {
	conn Connection
}

func NewBillRepository(conn Connection) BillRepository {
	return &billRepository{
		conn: conn,
	}
}

func GetByID(ID string) (found datos.Bill) {
	bill := datos.Bill{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&bill, ID)
	return bill
}

func Delete(ID string) (found datos.Bill) {
	bill := datos.Bill{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&bill.ID)
	db.Delete(bill)
	return bill
}

func Create(Bill) (found datos.Bill) {
	bill := datos.Bill{}
	db := utils.GetConnection()
	defer db.Close()
	err = db.Create(&bill).Error
	return bill
}

func Update(ID string) (found datos.Bill) {
	billFind := datos.Bill{}
	billData := datos.Bill{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&billFind, ID)
	db.Model(&billFind).Updates(billData)

}

func GetFindAll() (found datos.Bill) {
	bill := datos.Bill{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&bill)
	return bill

}
