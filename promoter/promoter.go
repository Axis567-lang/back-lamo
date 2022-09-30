package promoter

type Promoter interface {
	Name() Name
}

func New(name Name) DTO {
	return DTO{PromoterName: name}
}

type Name string

type DTO struct {
	PromoterName Name `json:"name,omitempty"`
}

func (d DTO) Name() Name {
	return d.PromoterName
}
