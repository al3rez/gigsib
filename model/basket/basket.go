package basket

import basket_item "gigsib/model/basket_item"

type Type struct {
	Items []*basket_item.Type
}

func (b *Type) TotalPrice() float64 {
	total := 0.0
	for _, item := range b.Items {
		total += item.Price
	}
	return total
}

func (b *Type) AddItemBang(item *basket_item.Type) {
	b.Items = append(b.Items, item)
}

func (b *Type) AddItem(item *basket_item.Type) []*basket_item.Type {
	return append(b.Items, item)
}

func New(items []*basket_item.Type) *Type {
	return &Type{items}
}
