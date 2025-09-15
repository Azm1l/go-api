package repository

import (
	"github.com/Azm1l/rest-api-go/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	ShowAll() ([]entity.User, error)
	FindById(id int64) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	DeleteUser(id int64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *entity.User) (*entity.User, error) {

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) ShowAll() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindById(id int64) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (r *userRepository) Update(user *entity.User) (*entity.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(id int64) error {
	var user entity.User
	if err := r.db.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}
