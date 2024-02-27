package repository

import (
	"fmt"

	"gorm.io/gorm"
	"newser.app/infra/dao"
	"newser.app/infra/dto"
	"newser.app/model"
)

type UserGormRepo struct {
	DB *gorm.DB
}

func NewUserGormRepo(db *gorm.DB) UserGormRepo {
	return UserGormRepo{DB: db}
}

func (r UserGormRepo) Get(udto dto.UserDTO) (model.User, error) {
	var u model.User
	ug := dao.UserGorm{
		Email:          udto.Email,
		HashedPassword: udto.HashedPassword,
	}
	res := r.DB.First(&ug, udto.Id)
	fmt.Println(res)
	if res.Error != nil {
		return u, res.Error
	}
	// u.Id = res
	return u, nil
}

func (r UserGormRepo) FindByEmail(email string) (model.User, error) {
	var ug dao.UserGorm
	var u model.User
	res := r.DB.Where("email = ?", email).First(&ug)
	if res.Error != nil {
		return u, res.Error
	}
	u.Email = ug.Email
	u.Id = ug.ID
	return u, nil
}

func (r UserGormRepo) GetHashedPasswordByEmail(email string) (string, error) {
	var ug dao.UserGorm
	res := r.DB.Where("email = ?", email).First(&ug)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("email not found")
		}
	}
	return ug.HashedPassword, nil
}

func (r UserGormRepo) FindById(id uint) (model.User, error) {
	var ug dao.UserGorm
	var u model.User
	res := r.DB.First(&ug, id)
	fmt.Println("result: ", res)
	if res.Error != nil {
		return u, res.Error
	}
	return u, nil
}

func (r UserGormRepo) Create(udto dto.UserDTO) (model.User, error) {
	ug := dao.UserGorm{
		Email:          udto.Email,
		HashedPassword: udto.HashedPassword,
	}
	res := r.DB.Create(&ug)
	u := model.User{
		Email: ug.Email,
	}
	if res.Error != nil {
		return u, res.Error
	}
	u.Id = ug.ID

	return u, nil
}

func (r UserGormRepo) All() []model.User {
	return []model.User{}
}

func (r UserGormRepo) Update(udto dto.UserDTO) (model.User, error) {
	return model.User{}, nil
}

func (r UserGormRepo) Delete(id uint) error {
	return nil
}
