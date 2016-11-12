package main

func SquareSumDifference(n uint64) uint64 {
	var squareSum, sumSquares, i uint64 = 0, 0, 1
	for i <= n {
		sumSquares += i * i
		squareSum += i
		i++
	}
	squareSum *= squareSum
	return squareSum - sumSquares
}
