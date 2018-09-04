package slice

func IntersectStrings(slice1, slice2 []string) (s []string) {
	if len(slice2) < 1 || len(slice1) < 1 {
		return
	}

	for _, v := range slice1 {

		for _, v2 := range slice2 {
			if v2 == v {

				s = append(s, v)
				break
			}
		}

	}

	return
}
