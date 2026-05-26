package mint

func trimLeadingZeros(s string) string {
	if s == "" || s == "0" {
		return "0"
	}
	i := 0
	for i < len(s)-1 && s[i] == '0' {
		i++
	}
	return s[i:]
}
