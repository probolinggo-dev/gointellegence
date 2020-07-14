package normalization

// MinMax : geting Min and Max
func MinMax(in []float64) (min, max float64) {
	if len(in) < 1 {
		return 0, 0
	}
	min = in[0]
	for _, i := range in {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return
}

// MinMaxScaller :(each data - min.all_data) / (max.all_data - min. all_data)
func MinMaxScaller(in []float64) []float64 {
	min, max := MinMax(in)
	scale := max - min
	if scale == 0 {
		return in
	}
	for i := range in {
		in[i] = (in[i] - min) / scale
	}
	return in
}
