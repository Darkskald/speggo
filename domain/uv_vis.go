package domain

type UvVis struct {
	Name       string    `json:"name"`
	WaveLength []float64 `json:"wavelength"`
	Absorbance []float64 `json:"absorbance"`
}
