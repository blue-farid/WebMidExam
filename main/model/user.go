package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	v "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string          `gorm:"unique,not null"`
	Password string          `gorm:"not null"`
	Roles    JSONStringArray `gorm:"type:json"`
}

type JSONStringArray []string

func (u *User) CreationValidate() error {
	return v.ValidateStruct(u, v.Field(&u.Username, v.Required), v.Field(&u.Password, v.Required))
}

func (u *User) HasRole(r string) bool {
	for _, role := range u.Roles {
		if role == r {
			return true
		}
	}
	return false
}

func (j *JSONStringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSON value")
	}

	var result []string
	if err := json.Unmarshal(bytes, &result); err != nil {
		return err
	}

	*j = result
	return nil
}

func (j JSONStringArray) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}
