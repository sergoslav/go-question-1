package repository

type ProductRepo struct {
}

func (p *ProductRepo) GetProduct() string {
	return "product name"
}
