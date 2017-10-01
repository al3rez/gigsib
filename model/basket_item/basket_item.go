package basket_item

type Type struct {
	Id    string
	Price float64
}

func New(id string, price float64) *Type {
	return &Type{id, price}
}
