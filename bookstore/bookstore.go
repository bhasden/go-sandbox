package bookstore

type Book struct {
	Title string
	Author string
	PriceCents int
	DiscountPercent int
}

func NetPrice(b Book) int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}

func (b Book) NetPrice() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving }

type Catalog []Book

func (c Catalog) GetAllBooks() []Book {
	return c
}

func (c Catalog) Len() int {
	return len(c)
}