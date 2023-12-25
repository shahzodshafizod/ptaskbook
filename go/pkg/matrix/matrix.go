package matrix

import "github.com/shahzodshafizod/ptaskbook/go/pkg/array"

////////////////////////////Int////////////////////////////

// Int struct type;
type Int [][]int

// NewInt is a constructor for Int struct type;
func NewInt(rowsCount, colsCount int) Int {
	var matrix = make([][]int, rowsCount)
	for index := 0; index < rowsCount; index++ {
		matrix[index] = make([]int, colsCount)
	}
	return matrix
}

// ReverseRows reverses all rows of the Int matrix;
func (matrix *Int) ReverseRows(rowsCount, colsCount int) {
	for col := 0; col < colsCount/2; col++ {
		for row := 0; row < rowsCount; row++ {
			(*matrix)[row][col], (*matrix)[row][colsCount-1-col] = (*matrix)[row][colsCount-1-col], (*matrix)[row][col]
		}
	}
}

// ReverseCols reverses all cols of the Int matrix;
func (matrix *Int) ReverseCols(rowsCount, colsCount int) {
	for row := 0; row < rowsCount/2; row++ {
		for col := 0; col < colsCount; col++ {
			(*matrix)[row][col], (*matrix)[rowsCount-1-row][col] = (*matrix)[rowsCount-1-row][col], (*matrix)[row][col]
		}
	}
}

// InsertRowBefore inserts new rows before the index position in the Int matrix;
func (matrix *Int) InsertRowBefore(rowsCount, colsCount *int, index int, values ...[]int) {
	if index < 0 || index >= *rowsCount {
		return
	}
	(*rowsCount) += len(values)
	*matrix = append((*matrix)[:index], append(values, (*matrix)[index:]...)...)
}

// InsertRowAfter inserts new rows after the index position in the Int matrix;
func (matrix *Int) InsertRowAfter(rowsCount, colsCount *int, index int, values ...[]int) {
	if index < 0 || index >= *rowsCount {
		return
	}
	(*rowsCount) += len(values)
	*matrix = append((*matrix)[:index+1], append(values, (*matrix)[index+1:]...)...)
}

// InsertColBefore inserts new cols before the index position in the Int matrix;
func (matrix *Int) InsertColBefore(rowsCount, colsCount *int, index int, values ...[]int) {
	if index < 0 || index >= *colsCount {
		return
	}
	(*colsCount) += len(values)
	for _, value := range values {
		for row := 0; row < *rowsCount; row++ {
			tempRow := array.Int((*matrix)[row])
			tempRow.InsertBefore(index, value[row])
			(*matrix)[row] = tempRow
		}
	}
}

// InsertColAfter inserts new cols after the index position in the Int matrix;
func (matrix *Int) InsertColAfter(rowsCount, colsCount *int, index int, values ...[]int) {
	if index < 0 || index >= *colsCount {
		return
	}
	(*colsCount) += len(values)
	for _, value := range values {
		for row := 0; row < *rowsCount; row++ {
			tempRow := array.Int((*matrix)[row])
			tempRow.InsertAfter(index, value[row])
			(*matrix)[row] = tempRow
		}
	}
}

// DeleteRowAt deletes the row of the Int matrix at the index position;
func (matrix *Int) DeleteRowAt(rowsCount, colsCount *int, index int) []int {
	if index < 0 || index >= *rowsCount {
		return nil
	}
	(*rowsCount)--
	var elements = (*matrix)[index]
	*matrix = append((*matrix)[:index], (*matrix)[index+1:]...)
	return elements
}

// DeleteColAt deletes the col of the Int matrix at the index position;
func (matrix *Int) DeleteColAt(rowsCount, colsCount *int, index int) []int {
	if index < 0 || index >= *colsCount {
		return nil
	}
	(*colsCount)--
	var elements = make([]int, 0)
	for rowIndex, row := range *matrix {
		tempRow := array.Int(row)
		elements = append(elements, tempRow.DeleteAt(index))
		(*matrix)[rowIndex] = tempRow
	}
	return elements
}

////////////////////////////Float64////////////////////////////

// Float64 struct type;
type Float64 [][]float64

// NewFloat64 is a constructor for Float64 struct type;
func NewFloat64(rowsCount, colsCount int) Float64 {
	var matrix = make([][]float64, rowsCount)
	for index := 0; index < rowsCount; index++ {
		matrix[index] = make([]float64, colsCount)
	}
	return matrix
}

// ReverseRows reverses all rows of the Float64 matrix;
func (matrix *Float64) ReverseRows(rowsCount, colsCount int) {
	for col := 0; col < colsCount/2; col++ {
		for row := 0; row < rowsCount; row++ {
			(*matrix)[row][col], (*matrix)[row][colsCount-1-col] = (*matrix)[row][colsCount-1-col], (*matrix)[row][col]
		}
	}
}

// ReverseCols reverses all cols of the Float64 matrix;
func (matrix *Float64) ReverseCols(rowsCount, colsCount int) {
	for row := 0; row < rowsCount/2; row++ {
		for col := 0; col < colsCount; col++ {
			(*matrix)[row][col], (*matrix)[rowsCount-1-row][col] = (*matrix)[rowsCount-1-row][col], (*matrix)[row][col]
		}
	}
}

// InsertRowBefore inserts new rows before the index position in the Float64 matrix;
func (matrix *Float64) InsertRowBefore(rowsCount, colsCount *int, index int, values ...[]float64) {
	if index < 0 || index >= *rowsCount {
		return
	}
	(*rowsCount) += len(values)
	*matrix = append((*matrix)[:index], append(values, (*matrix)[index:]...)...)
}

// InsertRowAfter inserts new rows after the index position in the Float64 matrix;
func (matrix *Float64) InsertRowAfter(rowsCount, colsCount *int, index int, values ...[]float64) {
	if index < 0 || index >= *rowsCount {
		return
	}
	(*rowsCount) += len(values)
	*matrix = append((*matrix)[:index+1], append(values, (*matrix)[index+1:]...)...)
}

// InsertColBefore inserts new cols before the index position in the Float64 matrix;
func (matrix *Float64) InsertColBefore(rowsCount, colsCount *int, index int, values ...[]float64) {
	if index < 0 || index >= *colsCount {
		return
	}
	(*colsCount) += len(values)
	for _, value := range values {
		for row := 0; row < *rowsCount; row++ {
			tempRow := array.Float64((*matrix)[row])
			tempRow.InsertBefore(index, value[row])
			(*matrix)[row] = tempRow
		}
	}
}

// InsertColAfter inserts new cols after the index position in the Float64 matrix;
func (matrix *Float64) InsertColAfter(rowsCount, colsCount *int, index int, values ...[]float64) {
	if index < 0 || index >= *colsCount {
		return
	}
	(*colsCount) += len(values)
	for _, value := range values {
		for row := 0; row < *rowsCount; row++ {
			tempRow := array.Float64((*matrix)[row])
			tempRow.InsertAfter(index, value[row])
			(*matrix)[row] = tempRow
		}
	}
}

// DeleteRowAt deletes the row of the Float64 matrix at the index position;
func (matrix *Float64) DeleteRowAt(rowsCount *int, colsCount int, index int) []float64 {
	if index < 0 || index >= *rowsCount {
		return nil
	}
	(*rowsCount)--
	var elements = (*matrix)[index]
	*matrix = append((*matrix)[:index], (*matrix)[index+1:]...)
	return elements
}

// DeleteColAt deletes the col of the Float64 matrix at the index position;
func (matrix *Float64) DeleteColAt(rowsCount int, colsCount *int, index int) []float64 {
	if index < 0 || index >= *colsCount {
		return nil
	}
	(*colsCount)--
	var elements = make([]float64, 0)
	for rowIndex, row := range *matrix {
		tempRow := array.Float64(row)
		elements = append(elements, tempRow.DeleteAt(index))
		(*matrix)[rowIndex] = tempRow
	}
	return elements
}

func IsLocalMinimum(matrix [][]float64, rowsCount, colsCount, rowIndex, colIndex int) bool {
	// check with left;
	if rowIndex > 0 && !(matrix[rowIndex][colIndex] < matrix[rowIndex-1][colIndex]) {
		return false
	}
	// check with right;
	if rowIndex < rowsCount-1 && !(matrix[rowIndex][colIndex] < matrix[rowIndex+1][colIndex]) {
		return false
	}
	// check with top;
	if colIndex > 0 && !(matrix[rowIndex][colIndex] < matrix[rowIndex][colIndex-1]) {
		return false
	}
	// check with down;
	if colIndex < colsCount-1 && !(matrix[rowIndex][colIndex] < matrix[rowIndex][colIndex+1]) {
		return false
	}
	// check with top + left;
	if rowIndex > 0 && colIndex > 0 && !(matrix[rowIndex][colIndex] < matrix[rowIndex-1][colIndex-1]) {
		return false
	}
	// check with top + right;
	if rowIndex > 0 && colIndex < colsCount-1 && !(matrix[rowIndex][colIndex] < matrix[rowIndex-1][colIndex+1]) {
		return false
	}
	// check with down + left;
	if rowIndex < rowsCount-1 && colIndex > 0 && !(matrix[rowIndex][colIndex] < matrix[rowIndex+1][colIndex-1]) {
		return false
	}
	// check with down + right;
	if rowIndex < rowsCount-1 && colIndex < colsCount-1 && !(matrix[rowIndex][colIndex] < matrix[rowIndex+1][colIndex+1]) {
		return false
	}
	return true
}

func IsLocalMaximum(matrix [][]float64, rowsCount, colsCount, rowIndex, colIndex int) bool {
	// check with top
	if rowIndex > 0 && !(matrix[rowIndex][colIndex] > matrix[rowIndex-1][colIndex]) {
		return false
	}
	// check with bottom
	if rowIndex < rowsCount-1 && !(matrix[rowIndex][colIndex] > matrix[rowIndex+1][colIndex]) {
		return false
	}
	// check with left
	if colIndex > 0 && !(matrix[rowIndex][colIndex] > matrix[rowIndex][colIndex-1]) {
		return false
	}
	// check with right
	if colIndex < colsCount-1 && !(matrix[rowIndex][colIndex] > matrix[rowIndex][colIndex+1]) {
		return false
	}
	// check with top + left;
	if rowIndex > 0 && colIndex > 0 && !(matrix[rowIndex][colIndex] > matrix[rowIndex-1][colIndex-1]) {
		return false
	}
	// check with top + right;
	if rowIndex > 0 && colIndex < colsCount-1 && !(matrix[rowIndex][colIndex] > matrix[rowIndex-1][colIndex+1]) {
		return false
	}
	// check with down + left;
	if rowIndex < rowsCount-1 && colIndex > 0 && !(matrix[rowIndex][colIndex] > matrix[rowIndex+1][colIndex-1]) {
		return false
	}
	// check with down + right;
	if rowIndex < rowsCount-1 && colIndex < colsCount-1 && !(matrix[rowIndex][colIndex] > matrix[rowIndex+1][colIndex+1]) {
		return false
	}
	return true
}
