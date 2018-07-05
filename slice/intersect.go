package slice

func IntersectStrings(slice1, slice2 []string) (s []string) {
	if len(slice2) < 1 {
		return
	}

	m := map[string]bool{}

	for _, v := range slice1 {
		m[v] = true
	}

	for _, v := range slice2 {
		if m[v] {
			m[v] = true
		} else {
			m[v] = false
		}
	}

	for k, b := range m {
		if b {
			s = append(s, k)
		}
	}

	return
}
