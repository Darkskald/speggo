package sqlite

import (
	"speggo/domain"
)

type RawSFG struct {
	Id           int
	Name         string
	MeasuredTime string
	Wavenumbers  string
	Sfg          string
	Ir           string
	Vis          string
}

func (r RawSFG) ToSFG() (domain.SFG, error) {
	measuredTime := ParseMeasuredTime(r.MeasuredTime)
	wavenumbers := NumbersFromText(r.Wavenumbers)
	sfgIntensity := NumbersFromText(r.Sfg)
	ir := NumbersFromText(r.Ir)
	vis := NumbersFromText(r.Vis)

	temp := domain.SFG{
		Name:         r.Name,
		MeasuredTime: measuredTime,
		Wavenumbers:  wavenumbers,
		SfgIntensity: sfgIntensity,
		Ir:           ir,
		Vis:          vis,
	}
	return temp, nil

}
