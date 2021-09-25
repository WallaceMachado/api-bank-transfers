package utils

func GenerateString(n int) string {
	a := "a"
	s := ""
	for i := 0; i < n; i++ {
		s = s + a
	}

	return s
}
