package repository

import (
	"errors"
	"gorm.io/gorm"
	"server-side/database"
	"server-side/model"
	"server-side/utility"
	"time"
)

type IUserRepository interface {
	Migrate() error
	InsertUser(userCreateInput model.UserCreate) (*model.User, error)
	GetUser(id int64) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	UpdateUser(id int64, updateData model.UserUpdate) (*model.User, error)
	DeleteUser(id int64) error
}

var UserRepositoryInstance IUserRepository

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func InitUserRepository() error {
	UserRepositoryInstance = NewUserRepository(database.GormDB)
	err := UserRepositoryInstance.Migrate()
	return err
}

func (r *UserRepository) Migrate() error {
	return r.db.AutoMigrate(&model.User{})
}

func (r *UserRepository) InsertUser(userCreateInput model.UserCreate) (*model.User, error) {
	hashedPassword, err := utility.HashPassword(userCreateInput.Password)
	if err != nil {
		return nil, err
	}
	currentTime := time.Now()
	user := &model.User{
		ID:        0,
		Username:  userCreateInput.Username,
		Password:  hashedPassword,
		Name:      userCreateInput.Name,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUser(id int64) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(id int64, updateData model.UserUpdate) (*model.User, error) {
	existingRecord := model.User{}
	existingRecord.ID = id
	if err := r.db.First(&existingRecord).Error; err != nil {
		return nil, err // Record not found
	}
	if updateData.Name != nil {
		existingRecord.Name = *updateData.Name
	}
	if updateData.Password != nil {
		hashedPassword, err := utility.HashPassword(*updateData.Password)
		if err != nil {
			return nil, err
		}
		existingRecord.Password = hashedPassword
	}
	err := r.db.Save(&existingRecord).Error
	return &existingRecord, err
}

func (r *UserRepository) DeleteUser(id int64) error {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
