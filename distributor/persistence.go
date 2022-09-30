package distributor

type Persistence interface {
	AddDistributor(DTO) error
	GetDistributors() (Distributors, error)
	DeleteDistributor(Name) error
}

type Distributors map[Name]DTO
