package domain

import (
	"fmt"
	"time"
)

type SFG struct {
	Name         string    `json:"name"`
	MeasuredTime time.Time `json:"measured_time"`
	Wavenumbers  []float64 `json:"wavenumbers"`
	SfgIntensity []float64 `json:"sfg_intensity"`
	Ir           []float64 `json:"ir"`
	Vis          []float64 `json:"vis"`
}

type SfgFilter struct {
	Type string    `form:"type"`
	From time.Time `form:"from" time_format:"2006-01-02"`
	To   time.Time `form:"to" time_format:"2006-01-02"`
}

func (sf SfgFilter) GetQuery() string {
	if sf.Type == "" {
		return ""
	}
	return fmt.Sprintf("WHERE type='%s'", sf.Type)
}
