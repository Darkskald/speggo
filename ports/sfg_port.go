package ports

import "speggo/domain"

// todo: implement all the filtering here

type SfgPort interface {
	GetSfgById(id string) (domain.SFG, error)
	ListSfgSpectra() ([]domain.SFG, error)
}
