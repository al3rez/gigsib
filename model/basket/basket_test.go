package basket_test

import (
	Basket "shopping-center/model/basket"
	BasketItem "shopping-center/model/basket_item"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	basket := &Basket.Type{}

	newBasket := Basket.New(nil)

	assert.Equal(t, newBasket, basket)
}

func TestTotalPrice(t *testing.T) {
	basket := Basket.New(nil)
	basket.AddItemBang(BasketItem.New("foo", 1.25))
	basket.AddItemBang(BasketItem.New("bar", 1.69))

	totalPrice := basket.TotalPrice()

	assert.Equal(t, totalPrice, 1.25+1.69)
}
func TestAddItemBang(t *testing.T) {
	basket := Basket.New(nil)

	basket.AddItemBang(BasketItem.New("foo", 1.25))
	basket.AddItemBang(BasketItem.New("bar", 1.69))

	assert.Len(t, basket.Items, 2)
}

func TestAddItem(t *testing.T) {
	basket := Basket.New(nil)

	basket.AddItem(BasketItem.New("foo", 1.25))
	basket.AddItem(BasketItem.New("bar", 1.69))

	assert.Len(t, basket.Items, 0)
}
