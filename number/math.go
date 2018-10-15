package number

import "math"

func Round(val float64, places int) float64 {
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}

	return t
}

func RoundInt(val1, val2 int) int {
	return int(Round(float64(val1)/float64(val2), 0))
}

func MaxInt(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func MinInt(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
