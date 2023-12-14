package request

import v "github.com/go-ozzo/ozzo-validation"

type CreateBasket struct {
	Data  string
	State string
}

func (c *CreateBasket) Validate() error {
	return v.ValidateStruct(c, v.Field(&c.State, v.In("PENDING", "COMPLETE")), v.Field(&c.Data, v.Required))
}
