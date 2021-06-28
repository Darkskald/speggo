package sqlite

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

func NumbersFromText(s string) []float64 {
	temp := strings.Split(s, ",")
	out := make([]float64, 0)

	for _, i := range temp {
		s, err := strconv.ParseFloat(i, 64)
		if err != nil {
			log.Println(err)
		} else {
			out = append(out, convertInvalid(s)) // 3.14159265
		}
	}
	return out
}

func ParseMeasuredTime(s string) time.Time {
	const layout = "2006-01-02T15:04:05Z"
	t, err := time.Parse(layout, s)

	if err != nil {
		log.Println(err)
	}
	return t
}

func convertInvalid(f float64) float64 {
	var out float64
	switch {
	case math.IsInf(f, 1):
		out = 0
	case math.IsInf(f, -1):
		out = 0
	case math.IsNaN(f):
		out = 0
	default:
		out = f
	}
	return out
}
