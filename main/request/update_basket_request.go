package request

import v "github.com/go-ozzo/ozzo-validation"

type UpdateBasket struct {
	Data  string
	State string
}

func (u *UpdateBasket) Validate() error {
	return v.ValidateStruct(u, v.Field(&u.State, v.In("PENDING", "COMPLETE", "")))
}
