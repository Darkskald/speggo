package ports

import "speggo/domain"

type UvVisPort interface {
	GetUvVisById(id string) (domain.UvVis, error)
	ListUvVis() ([]domain.UvVis, error)
}
