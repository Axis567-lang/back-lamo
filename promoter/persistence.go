package promoter

type Persistence interface {
	GetPromoters() (Promoters, error)
	AddPromoter(Promoter) error
	DeletePromoter(Name) error
}

type Promoters map[Name]DTO
