package repository

type Repository struct {
	pr *ProductRepo
}

func Init() *Repository {
	return &Repository{
		pr: &ProductRepo{},
	}
}

func (r *Repository) GetProductRepo() ProductRepo {
	return *r.pr
}
