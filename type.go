package shared

type Item struct {
	UPC int
}

type Order struct {
	OrderNumber string
	Items       []Item
}

type Cubby Order

type Session struct {
	ID      int
	Cubbies []Cubby
}
