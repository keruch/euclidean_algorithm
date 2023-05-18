package euclidean_algorithm

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	for a%b != 0 {
		a = a % b
		a, b = b, a
	}
	if b < 0 {
		return -b // GCD is always positive
	}
	return b
}
