package product

type Servise struct {
}

func NewService() *Servise {
	return &Servise{}
}

func (s *Servise) List() []Product {
	return allProducts
}
