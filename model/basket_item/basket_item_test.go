package basket_item_test

import (
	BasketItem "shopping-center/model/basket_item"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	basketItem := &BasketItem.Type{}

	newBasketItem := BasketItem.New("", 0.0)

	assert.Equal(t, newBasketItem, basketItem)
}
