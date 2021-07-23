package ports

import (
	"speggo/domain"
)

// todo: implement all the filtering here, probably pass a query map to the List functions

type SfgPort interface {
	GetSfgById(id string) (domain.SFG, error)
	GetSfgByName(name string) (domain.SFG, error)
	ListSfgSpectra(filter string) ([]domain.SFG, error)
	FuzzySearch(filter string) ([]domain.MinimalSfg, error)
}
