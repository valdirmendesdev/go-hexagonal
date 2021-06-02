package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{
		Persistence: persistence,
	}
}

func (p *ProductService) Get(id string) (ProductInterface, error) {
	product, err := p.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	if _, err := product.IsValid(); err != nil {
		return &Product{}, err
	}
	result, err := p.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (p *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return &Product{}, err
	}
	result, err := p.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (p *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return &Product{}, err
	}
	result, err := p.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

