package repository

import "github.com/blue-farid/WebMidExam/model"

func FindUser(username string) (*model.User, error) {
	var u model.User

	res := db.Where("username = ?", username).First(&u)

	if res.Error != nil {
		return nil, res.Error
	}

	return &u, nil
}

func SaveUser(u *model.User) (*model.User, error) {
	res := db.Create(u)
	if res.Error != nil {
		return nil, res.Error
	}
	return u, nil
}

func GetUser(uID uint) (*model.User, error) {
	var u model.User

	res := db.Find(&u, uID)

	if res.Error != nil {
		return nil, res.Error
	}

	return &u, nil
}
