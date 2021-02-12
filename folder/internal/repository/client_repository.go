package repository

import (
	"context"
	"errors"
	"folder/internal/entity"

	"encoding/json"
	"fmt"
	"net/http"
)

type ClientRepository interface {
	FindAll(ctx context.Context) ([]Client, error)
	GetByID(ctx context.Context, id string) (Client, error)
	Create(ctx context.Context, client Client) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (Client, error)
}
type clientRepository struct {
	conn Connection
}

func NewBillRepository(conn Connection) ClientRepository {
	return &clientRepository{
		conn: conn,
	}
}

func GetByID(ID string) (found datos.Client) {
	client := datos.Client{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&client, ID)
	return client
}

func Delete(ID string) (found datos.Client) {
	client := datos.Client{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&client.ID)
	db.Delete(client)
	return client
}

func Create(Client) (found datos.Client) {
	client := datos.Client{}
	db := utils.GetConnection()
	defer db.Close()
	err = db.Create(&client).Error
	return client
}

func Update(ID string) (found datos.Client) {
	Find := datos.Client{}
	Data := datos.Client{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&Find, ID)
	db.Model(&Find).Updates(Data)

}

func GetFindAll() (found datos.Client) {
	client := datos.Client{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&client)
	return client

}
