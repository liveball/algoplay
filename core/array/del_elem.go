package array

func sliceDel(s []string, a string) (n []string) {
	// fmt.Println(s)
	for _, v := range s {
		if a != v {
			n = append(n, v)
		}
	}
	// fmt.Println(s, n)
	return
}

func sliceDel2(s []string, a string) (n []string) {
	// fmt.Println(s)
	// n = make([]string, 0, len(s)-1)
	for i := len(s) - 1; i > 0; i-- {
		if a == s[i] {
			n = append(s[:i], s[i+1:]...)
			i--
		}
	}
	// fmt.Println(s, n)
	return
}

func sliceDel3(s []string, a string) (n []string) {
	// fmt.Println(s)
	n = s[:0]
	for _, v := range s {
		if a != v {
			n = append(n, v)
		}
	}
	// fmt.Println(s, n)
	return
}
