package application

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED = "enabled"
)

type Product struct {
	ID string
	Name string
	Price float64
	Status string
}

func (p *Product) IsValid() (bool, error) {
	return false, nil
}

func (p *Product) Enable() error {
	panic("implement me")
}

func (p *Product) Disable() error {
	panic("implement me")
}

func (p *Product) GetID() string {
	panic("implement me")
}

func (p *Product) GetName() string {
	panic("implement me")
}

func (p *Product) GetStatus() string {
	panic("implement me")
}

func (p *Product) GetPrice() float64 {
	panic("implement me")
}
