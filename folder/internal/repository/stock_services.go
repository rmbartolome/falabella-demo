package repository

import (
	"context"
)

type StockRepository interface {
	FindAll(ctx context.Context) ([]Stock, error)
	GetByID(ctx context.Context, id string) (Stock, error)
	Create(ctx context.Context, stock Stock) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (Stock, error)
}
type stockRepository struct {
	conn Connection
}

func NewStockRepository(conn Connection) StockRepository {
	return &stockRepository{
		conn: conn,
	}
}
func GetByID(ID string) (found datos.Stock) {
	stock := datos.Stock{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&stock, ID)
	return stock
}

func Delete(ID string) (found datos.Stock) {
	stock := datos.Stock{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&stock.ID)
	db.Delete(stock)
	return stock
}

func Create(Stock) (found datos.Stock) {
	stock := datos.Stock{}
	db := utils.GetConnection()
	defer db.Close()
	err = db.Create(&stock).Error
	return stock
}

func Update(ID string) (found datos.Stock) {
	find := datos.Stock{}
	data := datos.Stock{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&find, ID)
	db.Model(&find).Updates(data)

}

func FindAll() (found datos.Stock) {
	stock := datos.Stock{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&stock)
	return stock

}
