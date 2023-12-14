package repository

import "github.com/blue-farid/WebMidExam/model"

func CreateBasket(b *model.Basket) (*model.Basket, error) {
	res := db.Preload("User").Create(b)
	if res.Error != nil {
		return nil, res.Error
	}
	return b, nil
}

func GetBasket(bID int) (*model.Basket, error) {
	var b model.Basket
	res := db.Preload("User").Find(&b, bID)
	if res.Error != nil {
		return nil, res.Error
	}
	return &b, nil
}

func UpdateBasket(b *model.Basket, id int) (*model.Basket, error) {
	oB, err := GetBasket(id)
	if err != nil {
		return nil, err
	}

	if b.Data != "" {
		oB.Data = b.Data
	}

	if b.State != "" {
		oB.State = b.State
	}

	res := db.Preload("User").Save(oB)
	if res.Error != nil {
		return nil, res.Error
	}
	return oB, nil
}

func GetAllBaskets(uID uint) ([]model.Basket, error) {
	var baskets []model.Basket
	result := db.Preload("User").Where("userId = ?", uID)
	if result.Error != nil {
		return nil, result.Error
	}
	return baskets, nil
}

func DeleteBasket(id int) error {
	var b model.Basket
	res := db.Delete(&b, id)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
