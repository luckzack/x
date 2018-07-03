package slice

// Eliminate the members of src2 from src
func EliminateStrings(src, src2 []string) []string {
	if src2 == nil || len(src2) < 1 {
		return src
	}
	var dst = []string{}

	for _, v1 := range src {
		has := false
		for _, v2 := range src2 {
			if v1 == v2 {
				has = true
				break
			}
		}

		if !has {
			dst = append(dst, v1)
		}
	}

	return dst
}
