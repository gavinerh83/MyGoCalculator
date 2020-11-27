package calc

//Add adds 2 integers and returns an integer
func Add(x, y int) int {
	return x + y
}

//Subtract subtract second integer from the first
func Subtract(x, y int) int {
	return x - y
}

//Divide divide second integer from the first
func Divide(x, y int) float64 {
	return float64(x) / float64(y)
}

//Multiply multiplies x by y
func Multiply(x, y int) int {
	return x * y
}
