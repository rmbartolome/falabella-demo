package repository

import (
	"context"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (User, error)
}
type userRepository struct {
	conn Connection
}

func NewUserRepository(conn Connection) UserRepository {
	return &userRepository{
		conn: conn,
	}
}
func GetByID(ID string) (found datos.User) {
	user := datos.User{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&user, ID)
	return user
}

func Delete(ID string) (found datos.User) {
	user := datos.User{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&user.ID)
	db.Delete(user)
	return user
}

func Create(User) (found datos.User) {
	user := datos.User{}
	db := utils.GetConnection()
	defer db.Close()
	err = db.Create(&user).Error
	return user
}

func Update(ID string) (found datos.User) {
	find := datos.User{}
	data := datos.User{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&find, ID)
	db.Model(&find).Updates(data)

}

func FindAll() (found datos.User) {
	user := datos.User{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&user)
	return user

}
