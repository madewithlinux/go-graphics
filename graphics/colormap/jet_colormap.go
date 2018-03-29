package colormap

import "image/color"

func JetColorMap() []color.Color {
	out := make([]color.Color, len(jet_colors))
	for i := 0; i < len(jet_colors); i++ {
		out[i] = color.RGBA{
			uint8(255 * jet_colors[i][0]),
			uint8(255 * jet_colors[i][1]),
			uint8(255 * jet_colors[i][2]),
			255,
		}
	}
	return out
}

var jet_colors = [][3]float64{
	{   0.00000,   0.00000,   0.51562},
	{   0.00000,   0.00000,   0.53125},
	{   0.00000,   0.00000,   0.54688},
	{   0.00000,   0.00000,   0.56250},
	{   0.00000,   0.00000,   0.57812},
	{   0.00000,   0.00000,   0.59375},
	{   0.00000,   0.00000,   0.60938},
	{   0.00000,   0.00000,   0.62500},
	{   0.00000,   0.00000,   0.64062},
	{   0.00000,   0.00000,   0.65625},
	{   0.00000,   0.00000,   0.67188},
	{   0.00000,   0.00000,   0.68750},
	{   0.00000,   0.00000,   0.70312},
	{   0.00000,   0.00000,   0.71875},
	{   0.00000,   0.00000,   0.73438},
	{   0.00000,   0.00000,   0.75000},
	{   0.00000,   0.00000,   0.76562},
	{   0.00000,   0.00000,   0.78125},
	{   0.00000,   0.00000,   0.79688},
	{   0.00000,   0.00000,   0.81250},
	{   0.00000,   0.00000,   0.82812},
	{   0.00000,   0.00000,   0.84375},
	{   0.00000,   0.00000,   0.85938},
	{   0.00000,   0.00000,   0.87500},
	{   0.00000,   0.00000,   0.89062},
	{   0.00000,   0.00000,   0.90625},
	{   0.00000,   0.00000,   0.92188},
	{   0.00000,   0.00000,   0.93750},
	{   0.00000,   0.00000,   0.95312},
	{   0.00000,   0.00000,   0.96875},
	{   0.00000,   0.00000,   0.98438},
	{   0.00000,   0.00000,   1.00000},
	{   0.00000,   0.01562,   1.00000},
	{   0.00000,   0.03125,   1.00000},
	{   0.00000,   0.04688,   1.00000},
	{   0.00000,   0.06250,   1.00000},
	{   0.00000,   0.07812,   1.00000},
	{   0.00000,   0.09375,   1.00000},
	{   0.00000,   0.10938,   1.00000},
	{   0.00000,   0.12500,   1.00000},
	{   0.00000,   0.14062,   1.00000},
	{   0.00000,   0.15625,   1.00000},
	{   0.00000,   0.17188,   1.00000},
	{   0.00000,   0.18750,   1.00000},
	{   0.00000,   0.20312,   1.00000},
	{   0.00000,   0.21875,   1.00000},
	{   0.00000,   0.23438,   1.00000},
	{   0.00000,   0.25000,   1.00000},
	{   0.00000,   0.26562,   1.00000},
	{   0.00000,   0.28125,   1.00000},
	{   0.00000,   0.29688,   1.00000},
	{   0.00000,   0.31250,   1.00000},
	{   0.00000,   0.32812,   1.00000},
	{   0.00000,   0.34375,   1.00000},
	{   0.00000,   0.35938,   1.00000},
	{   0.00000,   0.37500,   1.00000},
	{   0.00000,   0.39062,   1.00000},
	{   0.00000,   0.40625,   1.00000},
	{   0.00000,   0.42188,   1.00000},
	{   0.00000,   0.43750,   1.00000},
	{   0.00000,   0.45312,   1.00000},
	{   0.00000,   0.46875,   1.00000},
	{   0.00000,   0.48438,   1.00000},
	{   0.00000,   0.50000,   1.00000},
	{   0.00000,   0.51562,   1.00000},
	{   0.00000,   0.53125,   1.00000},
	{   0.00000,   0.54688,   1.00000},
	{   0.00000,   0.56250,   1.00000},
	{   0.00000,   0.57812,   1.00000},
	{   0.00000,   0.59375,   1.00000},
	{   0.00000,   0.60938,   1.00000},
	{   0.00000,   0.62500,   1.00000},
	{   0.00000,   0.64062,   1.00000},
	{   0.00000,   0.65625,   1.00000},
	{   0.00000,   0.67188,   1.00000},
	{   0.00000,   0.68750,   1.00000},
	{   0.00000,   0.70312,   1.00000},
	{   0.00000,   0.71875,   1.00000},
	{   0.00000,   0.73438,   1.00000},
	{   0.00000,   0.75000,   1.00000},
	{   0.00000,   0.76562,   1.00000},
	{   0.00000,   0.78125,   1.00000},
	{   0.00000,   0.79688,   1.00000},
	{   0.00000,   0.81250,   1.00000},
	{   0.00000,   0.82812,   1.00000},
	{   0.00000,   0.84375,   1.00000},
	{   0.00000,   0.85938,   1.00000},
	{   0.00000,   0.87500,   1.00000},
	{   0.00000,   0.89062,   1.00000},
	{   0.00000,   0.90625,   1.00000},
	{   0.00000,   0.92188,   1.00000},
	{   0.00000,   0.93750,   1.00000},
	{   0.00000,   0.95312,   1.00000},
	{   0.00000,   0.96875,   1.00000},
	{   0.00000,   0.98438,   1.00000},
	{   0.00000,   1.00000,   1.00000},
	{   0.01562,   1.00000,   0.98438},
	{   0.03125,   1.00000,   0.96875},
	{   0.04688,   1.00000,   0.95312},
	{   0.06250,   1.00000,   0.93750},
	{   0.07812,   1.00000,   0.92188},
	{   0.09375,   1.00000,   0.90625},
	{   0.10938,   1.00000,   0.89062},
	{   0.12500,   1.00000,   0.87500},
	{   0.14062,   1.00000,   0.85938},
	{   0.15625,   1.00000,   0.84375},
	{   0.17188,   1.00000,   0.82812},
	{   0.18750,   1.00000,   0.81250},
	{   0.20312,   1.00000,   0.79688},
	{   0.21875,   1.00000,   0.78125},
	{   0.23438,   1.00000,   0.76562},
	{   0.25000,   1.00000,   0.75000},
	{   0.26562,   1.00000,   0.73438},
	{   0.28125,   1.00000,   0.71875},
	{   0.29688,   1.00000,   0.70312},
	{   0.31250,   1.00000,   0.68750},
	{   0.32812,   1.00000,   0.67188},
	{   0.34375,   1.00000,   0.65625},
	{   0.35938,   1.00000,   0.64062},
	{   0.37500,   1.00000,   0.62500},
	{   0.39062,   1.00000,   0.60938},
	{   0.40625,   1.00000,   0.59375},
	{   0.42188,   1.00000,   0.57812},
	{   0.43750,   1.00000,   0.56250},
	{   0.45312,   1.00000,   0.54688},
	{   0.46875,   1.00000,   0.53125},
	{   0.48438,   1.00000,   0.51562},
	{   0.50000,   1.00000,   0.50000},
	{   0.51562,   1.00000,   0.48438},
	{   0.53125,   1.00000,   0.46875},
	{   0.54688,   1.00000,   0.45312},
	{   0.56250,   1.00000,   0.43750},
	{   0.57812,   1.00000,   0.42188},
	{   0.59375,   1.00000,   0.40625},
	{   0.60938,   1.00000,   0.39062},
	{   0.62500,   1.00000,   0.37500},
	{   0.64062,   1.00000,   0.35938},
	{   0.65625,   1.00000,   0.34375},
	{   0.67188,   1.00000,   0.32812},
	{   0.68750,   1.00000,   0.31250},
	{   0.70312,   1.00000,   0.29688},
	{   0.71875,   1.00000,   0.28125},
	{   0.73438,   1.00000,   0.26562},
	{   0.75000,   1.00000,   0.25000},
	{   0.76562,   1.00000,   0.23438},
	{   0.78125,   1.00000,   0.21875},
	{   0.79688,   1.00000,   0.20312},
	{   0.81250,   1.00000,   0.18750},
	{   0.82812,   1.00000,   0.17188},
	{   0.84375,   1.00000,   0.15625},
	{   0.85938,   1.00000,   0.14062},
	{   0.87500,   1.00000,   0.12500},
	{   0.89062,   1.00000,   0.10938},
	{   0.90625,   1.00000,   0.09375},
	{   0.92188,   1.00000,   0.07812},
	{   0.93750,   1.00000,   0.06250},
	{   0.95312,   1.00000,   0.04688},
	{   0.96875,   1.00000,   0.03125},
	{   0.98438,   1.00000,   0.01562},
	{   1.00000,   1.00000,   0.00000},
	{   1.00000,   0.98438,   0.00000},
	{   1.00000,   0.96875,   0.00000},
	{   1.00000,   0.95312,   0.00000},
	{   1.00000,   0.93750,   0.00000},
	{   1.00000,   0.92188,   0.00000},
	{   1.00000,   0.90625,   0.00000},
	{   1.00000,   0.89062,   0.00000},
	{   1.00000,   0.87500,   0.00000},
	{   1.00000,   0.85938,   0.00000},
	{   1.00000,   0.84375,   0.00000},
	{   1.00000,   0.82812,   0.00000},
	{   1.00000,   0.81250,   0.00000},
	{   1.00000,   0.79688,   0.00000},
	{   1.00000,   0.78125,   0.00000},
	{   1.00000,   0.76562,   0.00000},
	{   1.00000,   0.75000,   0.00000},
	{   1.00000,   0.73438,   0.00000},
	{   1.00000,   0.71875,   0.00000},
	{   1.00000,   0.70312,   0.00000},
	{   1.00000,   0.68750,   0.00000},
	{   1.00000,   0.67188,   0.00000},
	{   1.00000,   0.65625,   0.00000},
	{   1.00000,   0.64062,   0.00000},
	{   1.00000,   0.62500,   0.00000},
	{   1.00000,   0.60938,   0.00000},
	{   1.00000,   0.59375,   0.00000},
	{   1.00000,   0.57812,   0.00000},
	{   1.00000,   0.56250,   0.00000},
	{   1.00000,   0.54688,   0.00000},
	{   1.00000,   0.53125,   0.00000},
	{   1.00000,   0.51562,   0.00000},
	{   1.00000,   0.50000,   0.00000},
	{   1.00000,   0.48438,   0.00000},
	{   1.00000,   0.46875,   0.00000},
	{   1.00000,   0.45312,   0.00000},
	{   1.00000,   0.43750,   0.00000},
	{   1.00000,   0.42188,   0.00000},
	{   1.00000,   0.40625,   0.00000},
	{   1.00000,   0.39062,   0.00000},
	{   1.00000,   0.37500,   0.00000},
	{   1.00000,   0.35938,   0.00000},
	{   1.00000,   0.34375,   0.00000},
	{   1.00000,   0.32812,   0.00000},
	{   1.00000,   0.31250,   0.00000},
	{   1.00000,   0.29688,   0.00000},
	{   1.00000,   0.28125,   0.00000},
	{   1.00000,   0.26562,   0.00000},
	{   1.00000,   0.25000,   0.00000},
	{   1.00000,   0.23438,   0.00000},
	{   1.00000,   0.21875,   0.00000},
	{   1.00000,   0.20312,   0.00000},
	{   1.00000,   0.18750,   0.00000},
	{   1.00000,   0.17188,   0.00000},
	{   1.00000,   0.15625,   0.00000},
	{   1.00000,   0.14062,   0.00000},
	{   1.00000,   0.12500,   0.00000},
	{   1.00000,   0.10938,   0.00000},
	{   1.00000,   0.09375,   0.00000},
	{   1.00000,   0.07812,   0.00000},
	{   1.00000,   0.06250,   0.00000},
	{   1.00000,   0.04688,   0.00000},
	{   1.00000,   0.03125,   0.00000},
	{   1.00000,   0.01562,   0.00000},
	{   1.00000,   0.00000,   0.00000},
	{   0.98438,   0.00000,   0.00000},
	{   0.96875,   0.00000,   0.00000},
	{   0.95312,   0.00000,   0.00000},
	{   0.93750,   0.00000,   0.00000},
	{   0.92188,   0.00000,   0.00000},
	{   0.90625,   0.00000,   0.00000},
	{   0.89062,   0.00000,   0.00000},
	{   0.87500,   0.00000,   0.00000},
	{   0.85938,   0.00000,   0.00000},
	{   0.84375,   0.00000,   0.00000},
	{   0.82812,   0.00000,   0.00000},
	{   0.81250,   0.00000,   0.00000},
	{   0.79688,   0.00000,   0.00000},
	{   0.78125,   0.00000,   0.00000},
	{   0.76562,   0.00000,   0.00000},
	{   0.75000,   0.00000,   0.00000},
	{   0.73438,   0.00000,   0.00000},
	{   0.71875,   0.00000,   0.00000},
	{   0.70312,   0.00000,   0.00000},
	{   0.68750,   0.00000,   0.00000},
	{   0.67188,   0.00000,   0.00000},
	{   0.65625,   0.00000,   0.00000},
	{   0.64062,   0.00000,   0.00000},
	{   0.62500,   0.00000,   0.00000},
	{   0.60938,   0.00000,   0.00000},
	{   0.59375,   0.00000,   0.00000},
	{   0.57812,   0.00000,   0.00000},
	{   0.56250,   0.00000,   0.00000},
	{   0.54688,   0.00000,   0.00000},
	{   0.53125,   0.00000,   0.00000},
	{   0.51562,   0.00000,   0.00000},
	{   0.50000,   0.00000,   0.00000},
}