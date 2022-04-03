package repositories

import (
	"github.com/RyaWcksn/go-api/entities"
	"gorm.io/gorm"
)

type UserInterface interface {
    CreatUser(user entities.User) (entities.User, error)
    GetUser(id int) (entities.User, error)
    GetUsers() ([]entities.User, error)
}

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db}
}

func (repo *UserRepository) CreatUser(user entities.User) (entities.User, error) {
    if err := repo.db.Create(&user).Error; err != nil {
        return user, err
    }
    return user, nil
}

func (repo *UserRepository) GetUser(id int) (entities.User, error) {
    var user entities.User
    if err := repo.db.First(&user, id).Error; err != nil {
        return user, err
    }
    return user, nil
}

func (repo *UserRepository) GetUsers() ([]entities.User, error) {
    var users []entities.User
    if err := repo.db.Find(&users).Error; err != nil {
        return users, err
    }
    return users, nil
}
