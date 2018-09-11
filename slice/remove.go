package slice

func RemoveString(sl []string, v string) []string {
	for i, vv := range sl {
		if vv == v {
			return append(sl[0:i], sl[i+1:]...)
		}
	}
	return sl
}
