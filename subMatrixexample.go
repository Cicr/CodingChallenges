package main

import (
	"fmt"
)

// Define a struct to represent the sub-matrix size (optional)
type SubMatrixSize struct {
	Rows int
	Cols int
}

func findSubMatrices(matrix [][]string, subMatrixSize SubMatrixSize) [][]string {
	rows, cols := len(matrix), len(matrix[0])
	subRowSize, subColSize := subMatrixSize.Rows, subMatrixSize.Cols

	var subMatrices [][]string

	// Iterate through all possible starting positions for the sub-matrix
	for i := 0; i <= rows-subRowSize; i++ {
		for j := 0; j <= cols-subColSize; j++ {
			subMatrix := make([]string, 0) // Create an empty sub-matrix slice
			for r := 0; r < subRowSize; r++ {
				// Extract each element (string) from the sub-row and append
				for c := j; c < j+subColSize; c++ {
					subMatrix = append(subMatrix, matrix[i+r][c])
				}
			}
			subMatrices = append(subMatrices, subMatrix) // Add sub-matrix to results
		}
	}

	return subMatrices
}

func main() {
	matrix := [][]string{
		{"a", "b", "c", "d", "e"},
		{"f", "g", "h", "i", "j"},
		{"k", "l", "m", "n", "o"},
	}

	subMatrixSize := SubMatrixSize{Rows: 2, Cols: 2} // Specify sub-matrix size (2x2)

	subMatrices := findSubMatrices(matrix, subMatrixSize)

	fmt.Println("Sub-matrices:")
	for _, subMatrix := range subMatrices {
		for _, row := range subMatrix {
			fmt.Println(row)
		}
		fmt.Println() // Print a newline after each sub-matrix
	}
}
