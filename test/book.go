package test

type Book struct {
	Total int
	Name  string
}

func (b *Book) GetTotal() int {
	return b.Total
}

func (b *Book) GetName() string {
	return b.Name
}

func (b *Book) SetName(Name string) {
	b.Name = Name
}

func (b *Book) SetTotal(Total int) {
	b.Total = Total
}
