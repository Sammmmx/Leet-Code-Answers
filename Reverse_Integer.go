func reverse(x int) int {
    var result int
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}

	if result < -2147483648 || result > 2147483647 {
		return 0
	}

	return result
}
