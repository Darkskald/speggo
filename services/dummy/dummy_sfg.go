package dummy

import (
	"speggo/domain"
	"time"
)

type DummySfgService struct {
	Spectrum domain.SFG
}

func NewDummySfgService() DummySfgService {
	return DummySfgService{
		Spectrum: domain.SFG{
			Name:         "BX12_x1_#1",
			MeasuredTime: time.Now(),
			Wavenumbers:  []float64{0.3, 0.4, 0.5, 0.6},
			SfgIntensity: []float64{0.3, 0.4, 0.5, 0.6},
			Ir:           []float64{0.3, 0.4, 0.5, 0.6},
			Vis:          []float64{0.3, 0.4, 0.5, 0.6},
		},
	}
}

func (ds DummySfgService) GetSfgById(id string) (domain.SFG, error) {
	return ds.Spectrum, nil
}

func (ds DummySfgService) ListSfgSpectra() ([]domain.SFG, error){
	return []domain.SFG{ds.Spectrum}, nil
}