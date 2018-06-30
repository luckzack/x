package utils

import (
	"fmt"
)

func ToBandwidth(raw float64) string {
	raw = raw * 8
	var t float64 = 1024
	var d float64 = 1

	if raw < t {
		return fmt.Sprintf("%.1fbps", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fKbps", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fMbps", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fGbps", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fTbps", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fPbps", raw/d)
	}

	return "TooLarge"
}

// 1024=>1,1k
func HumanizeNumber(raw float64) string{
	var t float64 = 1024
	var d float64 = 1

	if raw < t {
		return fmt.Sprintf("%.1fB", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fK", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fM", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fG", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fT", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.1fP", raw/d)
	}

	return "TooLarge"


}
