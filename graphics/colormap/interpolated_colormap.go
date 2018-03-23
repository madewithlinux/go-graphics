package colormap

import (
	"github.com/joshua-wright/go-graphics/graphics/interpolation"
	"image/color"
)

type InterpColormap struct {
	r, g, b *interpolation.CubicEquidistant
}

func NewInterpColormap(colors []color.Color) *InterpColormap {
	nPts := len(colors)
	rs := make([]float64, nPts)
	gs := make([]float64, nPts)
	bs := make([]float64, nPts)
	for i := 0; i < nPts; i++ {
		r, g, b, a := colors[i].RGBA()
		rs[i] = float64(r) / float64(a)
		gs[i] = float64(g) / float64(a)
		bs[i] = float64(b) / float64(a)
	}

	icm := InterpColormap{}
	icm.r = interpolation.NewCubicEquidistantInterpolatorFromSlices(0, 1, nPts, rs)
	icm.g = interpolation.NewCubicEquidistantInterpolatorFromSlices(0, 1, nPts, gs)
	icm.b = interpolation.NewCubicEquidistantInterpolatorFromSlices(0, 1, nPts, bs)
	return &icm
}

// 0 <= x < 1
func (icm *InterpColormap) GetColor(x float64) color.RGBA {
	return color.RGBA{
		R: uint8(icm.r.Get(x) * 255),
		G: uint8(icm.g.Get(x) * 255),
		B: uint8(icm.b.Get(x) * 255),
		A: 255,
	}
}
