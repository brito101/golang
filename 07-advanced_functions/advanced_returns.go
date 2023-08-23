package main

func main() {
	rSum, _, rMultiplication := math_operations(1, 4)
	println(rSum)
	// println(result_sub)
	println(rMultiplication)
}

func math_operations(a, b int) (sum, sub, multiplication int) {
	sum = a + b
	sub = a - b
	multiplication = a * b
	return
}
