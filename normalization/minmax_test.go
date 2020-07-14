package normalization

import (
	"testing"
)

func TestMinMax(t *testing.T) {
	in := []float64{1, 2, 3, 4, 45, 5, 6}
	o1, o2 := MinMax(in)
	if o1 != 1 {
		t.Errorf("Minimalnya expect %v, got %v", 1, o1)
	}
	if o2 != 45 {
		t.Errorf("Maximalnya expect %v, got %v", 45, o2)

	}

	in = []float64{0.5, .1, .5, 100, 45, .100, .033}
	o1, o2 = MinMax(in)
	if o1 != .033 {
		t.Errorf("Minimalnya expect %v, got %v", .033, o1)
	}
	if o2 != 100 {
		t.Errorf("Maximalnya expect %v, got %v", 100, o2)

	}

	in = []float64{.033}
	o1, o2 = MinMax(in)
	if o1 != .033 {
		t.Errorf("Minimalnya expect %v, got %v", .033, o1)
	}
	if o2 != .033 {
		t.Errorf("Maximalnya expect %v, got %v", .033, o2)

	}

	in = []float64{}
	o1, o2 = MinMax(in)
	if o1 != 0 {
		t.Errorf("Minimalnya expect %v, got %v", 0, o1)
	}
	if o2 != 0 {
		t.Errorf("Maximalnya expect %v, got %v", 0, o1)
	}
}

func equal(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestMinMaxScaller(t *testing.T) {
	in := []float64{1, 2, 3, 4, 45, 5, 6}
	out := MinMaxScaller(in)
	if !equal(out, []float64{0, 0.022727272727272728, 0.045454545454545456, 0.06818181818181818, 1, 0.09090909090909091, 0.11363636363636363}) {
		t.Error("Result is wrong!")
	}
	in = []float64{1, 1, 1}
	out = MinMaxScaller(in)
	if !equal(out, []float64{1, 1, 1}) {
		t.Errorf("Result is wrong %v!", out)
	}

}
