package slice

func UnionStrings(slice1, slice2 []string) (s []string) {
	if len(slice2) < 1 || len(slice1) < 1 {
		return
	}

	m := map[string]bool{}

	for _, v := range slice1 {
		m[v] = true
		s = append(s, v)
	}

	for _, v := range slice2 {
		if !m[v] {
			m[v] = true
			s = append(s, v)
		}
	}

	m = nil
	return
}
