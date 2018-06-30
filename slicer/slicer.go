package slicer

const default_size = 10

func SliceStrings(src []string, args ...int) (dst [][]string) {
	var size int
	if len(args) < 1 {
		size = default_size
	} else {
		if args[0] == 0 {
			size = default_size
		}
		size = args[0]
	}

	src_len := len(src)
	dst_len := (src_len-1)/size + 1
	dst = make([][]string, dst_len)

	for i := 0; i < src_len; i += size {
		if len(src) > i+size {
			//	dst = append(dst, src[i:i+size])
			dst[i/size] = src[i : i+size]
		} else {
			//dst = append(dst, src[i:])
			dst[i/size] = src[i:]
		}
	}

	return
}
