package item

type Item struct {
	name     string
	price    float64
	quantity int
}

func New(name string, price float64, quantity int) *Item {
	if price <= 0 || quantity <= 0 || len(name) == 0 {
		return nil
	}
	return &Item{name, price, quantity}
}

// getter
func (t *Item) Name() string {
	return t.name
}

// ! setter
func (t *Item) SetName(n string) {
	if len(n) != 0 {
		t.name = n
	}
}

// getter
func (t *Item) Price() float64 {
	return t.price
}

// ! setter
func (t *Item) SetPrice(p float64) {
	if p > 0 {
		t.price = p
	}
}

// getter
func (t *Item) Quantity() int {
	return t.quantity
}

// ! setter
func (t *Item) SetQuantity(q int) {
	if q > 0 {
		t.quantity = q
	}
}
