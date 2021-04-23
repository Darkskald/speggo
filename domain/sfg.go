package domain

import "time"

type SFG struct {
	Name         string    `json:"name"`
	MeasuredTime time.Time `json:"measured_time"`
	Wavenumbers  []float64 `json:"wavenumbers"`
	SfgIntensity []float64 `json:"sfg_intensity"`
	Ir           []float64 `json:"ir"`
	Vis          []float64 `json:"vis"`
}
