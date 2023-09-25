package product

type ProductService struct {
}

var allProducts = []Product{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "qwe"},
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) List() []Product {
	return allProducts
}
