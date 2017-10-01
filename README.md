Golang Is Good, Structure Is Bad
===============
I begin with `imports`. Nomarlly you just import a package and the last element of the package path would be the package you want to use.

```go
import "foo/bar/baz"
```

Problem
-------

So what if there's (probably is) type called `Baz` and you use it:

```go
import "foo/bar/baz"

func main() {
        baz := baz.New()
}
```

![wat.](http://i0.kym-cdn.com/photos/images/newsfeed/000/173/576/Wat8.jpg)

You just shadowed the package name with calling a variable `baz` and you writing baz twice `baz.Baz` in linguistics terms which is not good (remember DRY!) and probably you'd write a factory function called `NewBaz` that's even worse.

```go
package baz

type Baz struct{}

func NewBaz() {
        // do something
}
```

Solution
--------

I found a solution or trick to naming things in decent way IMO. So you just name your package whatever you want and there must be only one exported type inside the package and you should call it `Type` and you `must` separate each type into different packages and you `must` name packages explicitly with in `CamelCase`.

```go
import Baz "foo/bar/baz"

var baz Baz.Type{}
```

Why?
----

With that trick I can use packages as metaphor to classes and you can write:

```go
package main

import Basket "gigsib/model/basket"
import BasketItem "gigsib/model/basket_item"

func main() {
        basket := Basket.New(nil)
        basket.AddItemBang(BasketItem.New("NikeLab ACG 07 KMTR", "1100"))
        basket.AddItemBang(BasketItem.New("Nike Kobe A.D. Triple Black", "750"))
        basket.TotalPrice()
}
```

```ruby
    require "gigsib/model/basket"
    require "gigsib/model/basket_item"

    basket = Basket.new(nil)
    basket.add_item!(BasketItem.new("NikeLab ACG 07 KMTR", "1100"))
    basket.add_item!(BasketItem.new("Nike Kobe A.D. Triple Black", "750"))
    basket.total_price()
```

Do you see a big difference between ruby and golang? (expect ! and naming convention and entrypoint)
