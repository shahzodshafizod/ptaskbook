package task

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/array"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/matrix"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type matrixTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewMatrixTask(pathPrefix string) Task {
	return &matrixTask{
		dataPath: pathPrefix + "/12matrix/Matrix%03d/%03d/%03d",
		name:     "Matrix",
		count:    100,
	}
}

func (mt *matrixTask) GetCount() int { return mt.count }

func (mt *matrixTask) GetName() string { return mt.name }

func (mt *matrixTask) SetData(taskNo int, testNo int) error {
	var filePath = fmt.Sprintf(mt.dataPath, taskNo, testNo, testNo)
	scanner, err := scanner.New(filePath + ".dat")
	if err != nil {
		return errors.Wrap(err, "scanner.New")
	}
	checker, err := checker.New(filePath + ".ans")
	if err != nil {
		return errors.Wrap(err, "checker.New")
	}
	mt.scanner = scanner
	mt.checker = checker
	return nil
}

func (mt *matrixTask) ScannerEOF() bool { return mt.scanner.EOF() }

func (mt *matrixTask) CheckerEOF() bool { return mt.checker.EOF() }

func (mt *matrixTask) CleanData() {
	mt.scanner.RemoveFiles()
	mt.checker.RemoveFiles()
}

func (mt *matrixTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return mt.matrix1
	case 2:
		return mt.matrix2
	case 3:
		return mt.matrix3
	case 4:
		return mt.matrix4
	case 5:
		return mt.matrix5
	case 6:
		return mt.matrix6
	case 7:
		return mt.matrix7
	case 8:
		return mt.matrix8
	case 9:
		return mt.matrix9
	case 10:
		return mt.matrix10
	case 11:
		return mt.matrix11
	case 12:
		return mt.matrix12
	case 13:
		return mt.matrix13
	case 14:
		return mt.matrix14
	case 15:
		return mt.matrix15
	case 16:
		return mt.matrix16
	case 17:
		return mt.matrix17
	case 18:
		return mt.matrix18
	case 19:
		return mt.matrix19
	case 20:
		return mt.matrix20
	case 21:
		return mt.matrix21
	case 22:
		return mt.matrix22
	case 23:
		return mt.matrix23
	case 24:
		return mt.matrix24
	case 25:
		return mt.matrix25
	case 26:
		return mt.matrix26
	case 27:
		return mt.matrix27
	case 28:
		return mt.matrix28
	case 29:
		return mt.matrix29
	case 30:
		return mt.matrix30
	case 31:
		return mt.matrix31
	case 32:
		return mt.matrix32
	case 33:
		return mt.matrix33
	case 34:
		return mt.matrix34
	case 35:
		return mt.matrix35
	case 36:
		return mt.matrix36
	case 37:
		return mt.matrix37
	case 38:
		return mt.matrix38
	case 39:
		return mt.matrix39
	case 40:
		return mt.matrix40
	case 41:
		return mt.matrix41
	case 42:
		return mt.matrix42
	case 43:
		return mt.matrix43
	case 44:
		return mt.matrix44
	case 45:
		return mt.matrix45
	case 46:
		return mt.matrix46
	case 47:
		return mt.matrix47
	case 48:
		return mt.matrix48
	case 49:
		return mt.matrix49
	case 50:
		return mt.matrix50
	case 51:
		return mt.matrix51
	case 52:
		return mt.matrix52
	case 53:
		return mt.matrix53
	case 54:
		return mt.matrix54
	case 55:
		return mt.matrix55
	case 56:
		return mt.matrix56
	case 57:
		return mt.matrix57
	case 58:
		return mt.matrix58
	case 59:
		return mt.matrix59
	case 60:
		return mt.matrix60
	case 61:
		return mt.matrix61
	case 62:
		return mt.matrix62
	case 63:
		return mt.matrix63
	case 64:
		return mt.matrix64
	case 65:
		return mt.matrix65
	case 66:
		return mt.matrix66
	case 67:
		return mt.matrix67
	case 68:
		return mt.matrix68
	case 69:
		return mt.matrix69
	case 70:
		return mt.matrix70
	case 71:
		return mt.matrix71
	case 72:
		return mt.matrix72
	case 73:
		return mt.matrix73
	case 74:
		return mt.matrix74
	case 75:
		return mt.matrix75
	case 76:
		return mt.matrix76
	case 77:
		return mt.matrix77
	case 78:
		return mt.matrix78
	case 79:
		return mt.matrix79
	case 80:
		return mt.matrix80
	case 81:
		return mt.matrix81
	case 82:
		return mt.matrix82
	case 83:
		return mt.matrix83
	case 84:
		return mt.matrix84
	case 85:
		return mt.matrix85
	case 86:
		return mt.matrix86
	case 87:
		return mt.matrix87
	case 88:
		return mt.matrix88
	case 89:
		return mt.matrix89
	case 90:
		return mt.matrix90
	case 91:
		return mt.matrix91
	case 92:
		return mt.matrix92
	case 93:
		return mt.matrix93
	case 94:
		return mt.matrix94
	case 95:
		return mt.matrix95
	case 96:
		return mt.matrix96
	case 97:
		return mt.matrix97
	case 98:
		return mt.matrix98
	case 99:
		return mt.matrix99
	case 100:
		return mt.matrix100
	default:
		return nil
	}
}

// { "no": 1, "dat": "", "ans": "" }
func (mt *matrixTask) matrix1() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]int, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = 10 * (row + 1)
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareInt(matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 2, "dat": "", "ans": "" }
func (mt *matrixTask) matrix2() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]int, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = 5 * (col + 1)
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareInt(matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 3, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix3() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var array = make([]float64, m)
	for index := 0; index < m; index++ {
		array[index] = mt.scanner.NextFloat64()
	}
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = array[row]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 4, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix4() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = mt.scanner.NextFloat64()
	}
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = array[col]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 5, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix5() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var d = mt.scanner.NextFloat64()
	var array = make([]float64, m)
	for index := 0; index < m; index++ {
		array[index] = mt.scanner.NextFloat64()
	}
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		matrix[row][0] = array[row]
		for col := 1; col < n; col++ {
			matrix[row][col] = matrix[row][col-1] + d
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 6, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix6() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var r = mt.scanner.NextFloat64()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = mt.scanner.NextFloat64()
	}
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			if row == 0 {
				matrix[row][col] = array[col]
			} else {
				matrix[row][col] = matrix[row-1][col] * r
			}
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 7, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix7() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var k = mt.scanner.NextInt()
	if k >= 1 && k <= m {
		if !mt.checker.CompareFloat64(2, matrix[k-1]...) {
			return false
		}
	}
	return true
}

// { "no": 8, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix8() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var k = mt.scanner.NextInt()
	if k >= 1 && k <= n {
		for row := 0; row < m; row++ {
			if !mt.checker.CompareFloat64(2, matrix[row][k-1]) {
				return false
			}
		}
	}
	return true
}

// { "no": 9, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix9() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := 1; row < m; row += 2 {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 10, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix10() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for col := 0; col < n; col += 2 {
		for row := 0; row < m; row++ {
			if !mt.checker.CompareFloat64(2, matrix[row][col]) {
				return false
			}
		}
	}
	return true
}

// { "no": 11, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix11() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := 0; row < m; row++ {
		if row%2 == 0 {
			for col := 0; col < n; col++ {
				if !mt.checker.CompareFloat64(2, matrix[row][col]) {
					return false
				}
			}
		} else {
			for col := n - 1; col >= 0; col-- {
				if !mt.checker.CompareFloat64(2, matrix[row][col]) {
					return false
				}
			}
		}
	}
	return true
}

// { "no": 12, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix12() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for col := 0; col < n; col++ {
		if col%2 == 0 {
			for row := 0; row < m; row++ {
				if !mt.checker.CompareFloat64(2, matrix[row][col]) {
					return false
				}
			}
		} else {
			for row := m - 1; row >= 0; row-- {
				if !mt.checker.CompareFloat64(2, matrix[row][col]) {
					return false
				}
			}
		}
	}
	return true
}

// { "no": 13, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix13() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	for index := 0; index < m; index++ {
		for col := 0; col < m-index; col++ {
			if !mt.checker.CompareFloat64(2, a[index][col]) {
				return false
			}
		}
		for row := index + 1; row < m; row++ {
			if !mt.checker.CompareFloat64(2, a[row][m-1-index]) {
				return false
			}
		}
	}
	return true
}

// { "no": 14, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix14() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	for index := 0; index < m; index++ {
		for row := 0; row < m-index; row++ {
			if !mt.checker.CompareFloat64(2, a[row][index]) {
				return false
			}
		}
		for col := index + 1; col < m; col++ {
			if !mt.checker.CompareFloat64(2, a[m-1-index][col]) {
				return false
			}
		}
	}
	return true
}

// { "no": 15, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix15() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	for index := 0; index < m/2; index++ {
		for col := index; col < m-index; col++ {
			if !mt.checker.CompareFloat64(2, a[index][col]) {
				return false
			}
		}
		for row := index + 1; row < m-index; row++ {
			if !mt.checker.CompareFloat64(2, a[row][m-1-index]) {
				return false
			}
		}
		for col := m - 2 - index; col >= index; col-- {
			if !mt.checker.CompareFloat64(2, a[m-1-index][col]) {
				return false
			}
		}
		for row := m - 2 - index; row > index; row-- {
			if !mt.checker.CompareFloat64(2, a[row][index]) {
				return false
			}
		}
	}
	if m%2 != 0 {
		if !mt.checker.CompareFloat64(2, a[m/2][m/2]) {
			return false
		}
	}
	return true
}

// { "no": 16, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix16() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	for index := 0; index < m/2; index++ {
		for row := index; row < m-index; row++ {
			if !mt.checker.CompareFloat64(2, a[row][index]) {
				return false
			}
		}
		for col := index + 1; col < m-index; col++ {
			if !mt.checker.CompareFloat64(2, a[m-1-index][col]) {
				return false
			}
		}
		for row := m - 2 - index; row >= index; row-- {
			if !mt.checker.CompareFloat64(2, a[row][m-1-index]) {
				return false
			}
		}
		for col := m - 2 - index; col > index; col-- {
			if !mt.checker.CompareFloat64(2, a[index][col]) {
				return false
			}
		}
	}
	if m%2 != 0 {
		if !mt.checker.CompareFloat64(2, a[m/2][m/2]) {
			return false
		}
	}
	return true
}

// { "no": 17, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix17() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var k = mt.scanner.NextInt()
	var sum, pro float64 = 0, 1
	if k >= 1 && k <= m {
		for col := 0; col < n; col++ {
			sum += matrix[k-1][col]
			pro *= matrix[k-1][col]
		}
	}
	return mt.checker.CompareFloat64(2, sum, pro)
}

// { "no": 18, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix18() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var k = mt.scanner.NextInt()
	var sum, pro float64 = 0, 1
	if k >= 1 && k <= n {
		for row := 0; row < m; row++ {
			sum += matrix[row][k-1]
			pro *= matrix[row][k-1]
		}
	}
	return mt.checker.CompareFloat64(2, sum, pro)
}

// { "no": 19, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix19() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var sum float64
	for row := 0; row < m; row++ {
		sum = 0
		for col := 0; col < n; col++ {
			sum += matrix[row][col]
		}
		if !mt.checker.CompareFloat64(2, sum) {
			return false
		}
	}
	return true
}

// { "no": 20, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix20() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var pro float64
	for col := 0; col < n; col++ {
		pro = 1
		for row := 0; row < m; row++ {
			pro *= matrix[row][col]
		}
		if !mt.checker.CompareFloat64(2, pro) {
			return false
		}
	}
	return true
}

// { "no": 21, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix21() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var sum, average float64
	for row := 0; row < m; row += 2 {
		sum = 0
		for col := 0; col < n; col++ {
			sum += matrix[row][col]
		}
		average = sum / float64(n)
		if !mt.checker.CompareFloat64(2, average) {
			return false
		}
	}
	return true
}

// { "no": 22, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix22() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var sum float64
	for col := 1; col < n; col += 2 {
		sum = 0
		for row := 0; row < m; row++ {
			sum += matrix[row][col]
		}
		if !mt.checker.CompareFloat64(2, sum) {
			return false
		}
	}
	return true
}

// { "no": 23, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix23() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var minimal float64
	for row := 0; row < m; row++ {
		minimal = 0
		for col := 0; col < n; col++ {
			if col == 0 || matrix[row][col] < minimal {
				minimal = matrix[row][col]
			}
		}
		if !mt.checker.CompareFloat64(2, minimal) {
			return false
		}
	}
	return true
}

// { "no": 24, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix24() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var maximal float64
	for col := 0; col < n; col++ {
		maximal = 0
		for row := 0; row < m; row++ {
			if row == 0 || matrix[row][col] > maximal {
				maximal = matrix[row][col]
			}
		}
		if !mt.checker.CompareFloat64(2, maximal) {
			return false
		}
	}
	return true
}

// { "no": 25, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix25() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var sum, maxSum float64
	var maxSumRowIndex = -1
	for row := 0; row < m; row++ {
		sum = 0
		for col := 0; col < n; col++ {
			sum += matrix[row][col]
		}
		if row == 0 || sum > maxSum {
			maxSum = sum
			maxSumRowIndex = row
		}
	}
	maxSumRowIndex++
	if !mt.checker.CompareInt(maxSumRowIndex) {
		return false
	}
	return mt.checker.CompareFloat64(2, maxSum)
}

// { "no": 26, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix26() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var pro, minPro float64
	var minProColIndex = -1
	for col := 0; col < n; col++ {
		pro = 1
		for row := 0; row < m; row++ {
			pro *= matrix[row][col]
		}
		if pro != 0 && (col == 0 || pro < minPro) {
			minPro = pro
			minProColIndex = col
		}
	}
	minProColIndex++
	if !mt.checker.CompareInt(minProColIndex) {
		return false
	}
	return mt.checker.CompareFloat64(2, minPro)
}

// { "no": 27, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix27() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var minimal, maximal float64
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if col == 0 || matrix[row][col] < minimal {
				minimal = matrix[row][col]
			}
		}
		if row == 0 || minimal > maximal {
			maximal = minimal
		}
	}
	return mt.checker.CompareFloat64(2, maximal)
}

// { "no": 28, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix28() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var maximal, minimal float64
	for col := 0; col < n; col++ {
		for row := 0; row < m; row++ {
			if row == 0 || matrix[row][col] > maximal {
				maximal = matrix[row][col]
			}
		}
		if col == 0 || maximal < minimal {
			minimal = maximal
		}
	}
	return mt.checker.CompareFloat64(2, minimal)
}

// { "no": 29, "dat": "2", "ans": "" }
func (mt *matrixTask) matrix29() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	var averages = make([]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		averages[row] = 0
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
			averages[row] += matrix[row][col]
		}
		averages[row] /= float64(n)
	}
	var count int
	for row := 0; row < m; row++ {
		count = 0
		for col := 0; col < n; col++ {
			if matrix[row][col] < averages[row] {
				count++
			}
		}
		if !mt.checker.CompareInt(count) {
			return false
		}
	}
	return true
}

// { "no": 30, "dat": "2", "ans": "" }
func (mt *matrixTask) matrix30() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	var averages = make([]float64, n)
	for col := 0; col < n; col++ {
		averages[col] = 0
	}
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
			averages[col] += matrix[row][col]
		}
	}
	var count int
	for col := 0; col < n; col++ {
		averages[col] /= float64(m)
		count = 0
		for row := 0; row < m; row++ {
			if matrix[row][col] > averages[col] {
				count++
			}
		}
		if !mt.checker.CompareInt(count) {
			return false
		}
	}
	return true
}

// { "no": 31, "dat": "2", "ans": "" }
func (mt *matrixTask) matrix31() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	var sum = 0.0
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
			sum += matrix[row][col]
		}
	}
	var average = sum / float64(m*n)
	var rowIndex, colIndex = -1, -1
	var distance, minDistance float64
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			distance = math.Abs(matrix[row][col] - average)
			if (row == 0 && col == 0) || distance <= minDistance {
				minDistance = distance
				rowIndex = row
				colIndex = col
			}
		}
	}
	return mt.checker.CompareInt(rowIndex+1, colIndex+1)
}

// { "no": 32, "dat": "", "ans": "" }
func (mt *matrixTask) matrix32() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]int, m)
	var positivesCount, negativesCount int
	var index = -1
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		positivesCount, negativesCount = 0, 0
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextInt()
			if matrix[row][col] > 0 {
				positivesCount++
			} else if matrix[row][col] < 0 {
				negativesCount++
			}
		}
		if index < 0 && positivesCount == negativesCount {
			index = row
		}
	}
	index++
	return mt.checker.CompareInt(index)
}

// { "no": 33, "dat": "", "ans": "" }
func (mt *matrixTask) matrix33() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]int, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextInt()
		}
	}
	var positivesCount, negativesCount int
	var index = -1
	for col := 0; col < n; col++ {
		positivesCount, negativesCount = 0, 0
		for row := 0; row < m; row++ {
			if matrix[row][col] > 0 {
				positivesCount++
			} else if matrix[row][col] < 0 {
				negativesCount++
			}
		}
		if negativesCount == positivesCount {
			index = col
		}
	}
	index++
	return mt.checker.CompareInt(index)
}

// { "no": 34, "dat": "", "ans": "" }
func (mt *matrixTask) matrix34() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]int, m)
	var index = -1
	var isEvenRow bool
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		isEvenRow = true
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextInt()
			if isEvenRow && matrix[row][col]%2 != 0 {
				isEvenRow = false
			}
		}
		if isEvenRow {
			index = row
		}
	}
	index++
	return mt.checker.CompareInt(index)
}

// { "no": 35, "dat": "", "ans": "" }
func (mt *matrixTask) matrix35() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]int, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextInt()
		}
	}
	var isOddColumn bool
	var index = -1
	for col := 0; col < n; col++ {
		isOddColumn = true
		for row := 0; row < m; row++ {
			if matrix[row][col]%2 == 0 {
				isOddColumn = false
			}
		}
		if isOddColumn {
			index = col
			break
		}
	}
	index++
	return mt.checker.CompareInt(index)
}

// { "no": 36, "dat": "", "ans": "" }
func (mt *matrixTask) matrix36() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]int, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextInt()
		}
	}
	var firstRow [101]bool
	if m > 1 {
		for col := 0; col < n; col++ {
			firstRow[matrix[0][col]] = true
		}
	}
	var similarRowsCount = 0
	var isSimilar bool
	for row := 1; row < m; row++ {
		var currentRow [101]bool
		for col := 0; col < n; col++ {
			currentRow[matrix[row][col]] = true
		}
		isSimilar = true
		for index := 0; index < 101; index++ {
			if firstRow[index] != currentRow[index] {
				isSimilar = false
				break
			}
		}
		if isSimilar {
			similarRowsCount++
		}
	}
	return mt.checker.CompareInt(similarRowsCount)
}

// { "no": 37, "dat": "", "ans": "" }
func (mt *matrixTask) matrix37() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]int, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextInt()
		}
	}
	var lastCol [101]bool
	if n > 1 {
		for row := 0; row < m; row++ {
			lastCol[matrix[row][n-1]] = true
		}
	}
	var similarColsCount = 0
	var isSimilar bool
	for col := 0; col < n-1; col++ {
		var currentCol [101]bool
		for row := 0; row < m; row++ {
			currentCol[matrix[row][col]] = true
		}
		isSimilar = true
		for index := 0; index < 101; index++ {
			if currentCol[index] != lastCol[index] {
				isSimilar = false
				break
			}
		}
		if isSimilar {
			similarColsCount++
		}
	}
	return mt.checker.CompareInt(similarColsCount)
}

// { "no": 38, "dat": "", "ans": "" }
func (mt *matrixTask) matrix38() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]int, m)
	var notEqualRowsCount = 0
	var exists bool
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		var equals [101]int
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextInt()
			equals[matrix[row][col]]++
		}
		exists = false
		for index := 0; index < 101; index++ {
			if equals[index] > 1 {
				exists = true
				break
			}
		}
		if !exists {
			notEqualRowsCount++
		}
	}
	return mt.checker.CompareInt(notEqualRowsCount)
}

// { "no": 39, "dat": "", "ans": "" }
func (mt *matrixTask) matrix39() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]int, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextInt()
		}
	}
	var exists bool
	var notEqualColsCount = 0
	for col := 0; col < n; col++ {
		var equals [101]int
		for row := 0; row < m; row++ {
			equals[matrix[row][col]]++
		}
		exists = false
		for index := 0; index < 101; index++ {
			if equals[index] > 1 {
				exists = true
				break
			}
		}
		if !exists {
			notEqualColsCount++
		}
	}
	return mt.checker.CompareInt(notEqualColsCount)
}

// { "no": 40, "dat": "", "ans": "" }
func (mt *matrixTask) matrix40() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]int, m)
	var maximal int
	var maximalRowIndex = -1
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		var equals [101]int
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextInt()
			equals[matrix[row][col]]++
		}
		for index := 0; index < 101; index++ {
			if (row == 0 && index == 0) || equals[index] >= maximal {
				maximal = equals[index]
				maximalRowIndex = row
			}
		}
	}
	maximalRowIndex++
	return mt.checker.CompareInt(maximalRowIndex)
}

// { "no": 41, "dat": "", "ans": "" }
func (mt *matrixTask) matrix41() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]int, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextInt()
		}
	}
	var maximal int
	var maximalColIndex = -1
	for col := 0; col < n; col++ {
		var equals [101]int
		for row := 0; row < m; row++ {
			equals[matrix[row][col]]++
		}
		for index := 0; index < 101; index++ {
			if (col == 0 && index == 0) || equals[index] > maximal {
				maximal = equals[index]
				maximalColIndex = col
			}
		}
	}
	maximalColIndex++
	return mt.checker.CompareInt(maximalColIndex)
}

// { "no": 42, "dat": "2", "ans": "" }
func (mt *matrixTask) matrix42() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	var isSortedAsc bool
	var sortedRowsCount = 0
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		isSortedAsc = true
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
			if col > 0 && matrix[row][col] < matrix[row][col-1] {
				isSortedAsc = false
			}
		}
		if isSortedAsc {
			sortedRowsCount++
		}
	}
	return mt.checker.CompareInt(sortedRowsCount)
}

// { "no": 43, "dat": "2", "ans": "" }
func (mt *matrixTask) matrix43() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var isSortedDesc bool
	var sortedColsCount = 0
	for col := 0; col < n; col++ {
		isSortedDesc = true
		for row := 0; row < m; row++ {
			if row > 0 && matrix[row-1][col] < matrix[row][col] {
				isSortedDesc = false
				break
			}
		}
		if isSortedDesc {
			sortedColsCount++
		}
	}
	return mt.checker.CompareInt(sortedColsCount)
}

// { "no": 44, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix44() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	var minimal float64
	var sortedAsc, sortedDesc bool
	var inited = false
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		sortedAsc, sortedDesc = true, true
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
			if col > 0 {
				if matrix[row][col-1] < matrix[row][col] {
					sortedDesc = false
				} else if matrix[row][col-1] > matrix[row][col] {
					sortedAsc = false
				}
			}
		}
		if sortedAsc {
			if !inited || matrix[row][0] < minimal {
				minimal = matrix[row][0]
				inited = true
			}
		} else if sortedDesc {
			if !inited || matrix[row][n-1] < minimal {
				minimal = matrix[row][n-1]
				inited = true
			}
		}
	}
	return mt.checker.CompareFloat64(2, minimal)
}

// { "no": 45, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix45() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var isSortedAsc, isSortedDesc bool
	var maximal float64
	var inited = false
	var compareRowIndex int
	for col := 0; col < n; col++ {
		isSortedAsc, isSortedDesc = true, true
		for row := 0; row < m; row++ {
			if row > 0 {
				if matrix[row-1][col] < matrix[row][col] {
					isSortedDesc = false
				} else if matrix[row-1][col] > matrix[row][col] {
					isSortedAsc = false
				}
			}
		}
		compareRowIndex = -1
		if isSortedAsc {
			compareRowIndex = m - 1
		} else if isSortedDesc {
			compareRowIndex = 0
		}
		if compareRowIndex >= 0 {
			if !inited || matrix[compareRowIndex][col] > maximal {
				maximal = matrix[compareRowIndex][col]
				inited = true
			}
		}
	}
	return mt.checker.CompareFloat64(2, maximal)
}

// { "no": 46, "dat": "", "ans": "" }
func (mt *matrixTask) matrix46() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]int, m)
	var maxsInRows = make([]int, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextInt()
			if col == 0 || matrix[row][col] > maxsInRows[row] {
				maxsInRows[row] = matrix[row][col]
			}
		}
	}
	var wantedElement = 0
	for row := 0; row < m; row++ {
	LOOP:
		for col := 0; col < n; col++ {
			if matrix[row][col] == maxsInRows[row] {
				var ismin = true
				for r := 0; r < m; r++ {
					if r == row {
						continue
					}
					if matrix[r][col] < maxsInRows[row] {
						ismin = false
						break
					}
				}
				if ismin {
					wantedElement = matrix[row][col]
					break LOOP
				}
			}
		}
	}
	return mt.checker.CompareInt(wantedElement)
}

// { "no": 47, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix47() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var k1 = mt.scanner.NextInt()
	var k2 = mt.scanner.NextInt()
	if k1 >= 1 && k1 < k2 && k2 <= m {
		for col := 0; col < n; col++ {
			matrix[k1-1][col], matrix[k2-1][col] = matrix[k2-1][col], matrix[k1-1][col]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 48, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix48() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var k1 = mt.scanner.NextInt()
	var k2 = mt.scanner.NextInt()
	if k1 >= 1 && k1 < k2 && k2 <= n {
		for row := 0; row < m; row++ {
			matrix[row][k1-1], matrix[row][k2-1] = matrix[row][k2-1], matrix[row][k1-1]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 49, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix49() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	var maximalIndex, minimalIndex int
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
			if col == 0 || matrix[row][col] >= matrix[row][maximalIndex] {
				maximalIndex = col
			}
			if col == 0 || matrix[row][col] <= matrix[row][minimalIndex] {
				minimalIndex = col
			}
		}
		matrix[row][minimalIndex], matrix[row][maximalIndex] = matrix[row][maximalIndex], matrix[row][minimalIndex]
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 50, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix50() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var minimalIndex, maximalIndex int
	for col := 0; col < n; col++ {
		for row := 0; row < m; row++ {
			if row == 0 || matrix[row][col] > matrix[maximalIndex][col] {
				maximalIndex = row
			}
			if row == 0 || matrix[row][col] < matrix[minimalIndex][col] {
				minimalIndex = row
			}
		}
		matrix[maximalIndex][col], matrix[minimalIndex][col] = matrix[minimalIndex][col], matrix[maximalIndex][col]
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 51, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix51() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	var minimalRowIndex, maximalRowIndex = -1, -1
	var maximal, minimal float64
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
			if row == 0 && col == 0 {
				maximalRowIndex, minimalRowIndex = row, row
				maximal, minimal = matrix[row][col], matrix[row][col]
			} else {
				if matrix[row][col] >= maximal {
					maximalRowIndex = row
					maximal = matrix[row][col]
				} else if matrix[row][col] <= minimal {
					minimalRowIndex = row
					minimal = matrix[row][col]
				}
			}
		}
	}
	if minimalRowIndex >= 0 && maximalRowIndex >= 0 && minimalRowIndex != maximalRowIndex {
		for col := 0; col < n; col++ {
			matrix[minimalRowIndex][col], matrix[maximalRowIndex][col] = matrix[maximalRowIndex][col], matrix[minimalRowIndex][col]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 52, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix52() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var maximal, minimal float64
	var maximalColIndex, minimalColIndex = -1, -1
	for col := 0; col < n; col++ {
		for row := 0; row < m; row++ {
			if row == 0 && col == 0 {
				maximal, minimal = matrix[row][col], matrix[row][col]
				maximalColIndex, minimalColIndex = col, col
			} else {
				if matrix[row][col] > maximal {
					maximal = matrix[row][col]
					maximalColIndex = col
				} else if matrix[row][col] < minimal {
					minimal = matrix[row][col]
					minimalColIndex = col
				}
			}
		}
	}
	if minimalColIndex >= 0 && maximalColIndex >= 0 && minimalColIndex != maximalColIndex {
		for row := 0; row < m; row++ {
			matrix[row][minimalColIndex], matrix[row][maximalColIndex] = matrix[row][maximalColIndex], matrix[row][minimalColIndex]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 53, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix53() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var positiveColIndex = -1
	var isPositive bool
	for col := 1; col < n; col++ {
		isPositive = true
		for row := 0; row < m; row++ {
			if matrix[row][col] <= 0 {
				isPositive = false
				break
			}
		}
		if isPositive {
			positiveColIndex = col
		}
	}
	if positiveColIndex > 0 {
		for row := 0; row < m; row++ {
			matrix[row][0], matrix[row][positiveColIndex] = matrix[row][positiveColIndex], matrix[row][0]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 54, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix54() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var negativeColIndex = -1
	for col := 0; col < n-1; col++ {
		var isNegative = true
		for row := 0; row < m; row++ {
			if matrix[row][col] > 0 {
				isNegative = false
				break
			}
		}
		if isNegative {
			negativeColIndex = col
			break
		}
	}
	if negativeColIndex >= 0 {
		for row := 0; row < m; row++ {
			matrix[row][n-1], matrix[row][negativeColIndex] = matrix[row][negativeColIndex], matrix[row][n-1]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 55, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix55() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := 0; row < m/2; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col], matrix[m/2+row][col] = matrix[m/2+row][col], matrix[row][col]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 56, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix56() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for col := 0; col < n/2; col++ {
		for row := 0; row < m; row++ {
			matrix[row][col], matrix[row][n/2+col] = matrix[row][n/2+col], matrix[row][col]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 57, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix57() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := 0; row < m/2; row++ {
		for col := 0; col < n/2; col++ {
			matrix[row][col], matrix[m/2+row][n/2+col] = matrix[m/2+row][n/2+col], matrix[row][col]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 58, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix58() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := 0; row < m/2; row++ {
		for col := 0; col < n/2; col++ {
			matrix[row][n/2+col], matrix[m/2+row][col] = matrix[m/2+row][col], matrix[row][n/2+col]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 59, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix59() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	matrix.ReverseCols(m, n)
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 60, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix60() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	matrix.ReverseRows(m, n)
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 61, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix61() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var k = mt.scanner.NextInt()
	matrix.DeleteRowAt(&m, n, k-1)
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 62, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix62() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var k = mt.scanner.NextInt()
	matrix.DeleteColAt(m, &n, k-1)
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 63, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix63() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	var minimal float64
	var minimalRowIndex = -1
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
			if minimalRowIndex < 0 || matrix[row][col] <= minimal {
				minimal = matrix[row][col]
				minimalRowIndex = row
			}
		}
	}
	matrix.DeleteRowAt(&m, n, minimalRowIndex)
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 64, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix64() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var maximal float64
	var maximalColIndex = -1
	for col := 0; col < n; col++ {
		for row := 0; row < m; row++ {
			if maximalColIndex < 0 || matrix[row][col] >= maximal {
				maximal = matrix[row][col]
				maximalColIndex = col
			}
		}
	}
	matrix.DeleteColAt(m, &n, maximalColIndex)
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 65, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix65() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var isPositiveCol bool
	for col := 0; col < n; col++ {
		isPositiveCol = true
		for row := 0; row < m; row++ {
			if matrix[row][col] <= 0 {
				isPositiveCol = false
				break
			}
		}
		if isPositiveCol {
			matrix.DeleteColAt(m, &n, col)
			break
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 66, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix66() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var isNegativeCol bool
	var negativeColIndex = -1
	for col := 0; col < n; col++ {
		isNegativeCol = true
		for row := 0; row < m; row++ {
			if matrix[row][col] >= 0 {
				isNegativeCol = false
				break
			}
		}
		if isNegativeCol {
			negativeColIndex = col
		}
	}
	if negativeColIndex >= 0 {
		matrix.DeleteColAt(m, &n, negativeColIndex)
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 67, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix67() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var isPositiveCol bool
	for col := 0; col < n; col++ {
		isPositiveCol = true
		for row := 0; row < m; row++ {
			if matrix[row][col] <= 0 {
				isPositiveCol = false
				break
			}
		}
		if isPositiveCol {
			matrix.DeleteColAt(m, &n, col)
			col--
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 68, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix68() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var k = mt.scanner.NextInt()
	matrix.InsertRowBefore(&m, &n, k-1, array.NewFloat64(n, 0))
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 69, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix69() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var k = mt.scanner.NextInt()
	matrix.InsertColAfter(&m, &n, k-1, array.NewFloat64(m, 1))
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 70, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix70() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	var maximal float64
	var maximalRowIndex = -1
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
			if maximalRowIndex < 0 || matrix[row][col] > maximal {
				maximal = matrix[row][col]
				maximalRowIndex = row
			}
		}
	}
	if maximalRowIndex >= 0 {
		matrix.InsertRowAfter(&m, &n, maximalRowIndex, matrix[maximalRowIndex])
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 71, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix71() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	var minimal float64
	var minimalColIndex = -1
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
			if minimalColIndex < 0 || matrix[row][col] < minimal {
				minimal = matrix[row][col]
				minimalColIndex = col
			}
		}
	}
	if minimalColIndex >= 0 {
		var col = make([]float64, m)
		for index := 0; index < m; index++ {
			col[index] = matrix[index][minimalColIndex]
		}
		matrix.InsertColAfter(&m, &n, minimalColIndex, col)
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 72, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix72() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var isPositiveCol bool
	for col := 0; col < n; col++ {
		isPositiveCol = true
		for row := 0; row < m; row++ {
			if matrix[row][col] <= 0 {
				isPositiveCol = false
				break
			}
		}
		if isPositiveCol {
			matrix.InsertColBefore(&m, &n, col, array.NewFloat64(m, 1))
			break
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 73, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix73() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = matrix.NewFloat64(m, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var isNegativeCol bool
	var negativeColIndex = -1
	for col := 0; col < n; col++ {
		isNegativeCol = true
		for row := 0; row < m; row++ {
			if matrix[row][col] >= 0 {
				isNegativeCol = false
				break
			}
		}
		if isNegativeCol {
			negativeColIndex = col
		}
	}
	if negativeColIndex >= 0 {
		matrix.InsertColAfter(&m, &n, negativeColIndex, array.NewFloat64(m, 0))
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 74, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix74() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matr = make([][]float64, m)
	var tempMatrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matr[row] = make([]float64, n)
		tempMatrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matr[row][col] = mt.scanner.NextFloat64()
			tempMatrix[row][col] = matr[row][col]
		}
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix.IsLocalMinimum(tempMatrix, m, n, row, col) {
				matr[row][col] = 0
			}
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matr[row]...) {
			return false
		}
	}
	return true
}

// { "no": 75, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix75() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matr = make([][]float64, m)
	var tempMatrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matr[row] = make([]float64, n)
		tempMatrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matr[row][col] = mt.scanner.NextFloat64()
			tempMatrix[row][col] = matr[row][col]
		}
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix.IsLocalMaximum(tempMatrix, m, n, row, col) {
				matr[row][col] *= -1
			}
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matr[row]...) {
			return false
		}
	}
	return true
}

// { "no": 76, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix76() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	// the exchange sort method (the bubble sorting);
	for index := 0; index < m-1; index++ {
		for row := 1; row < m-index; row++ {
			if matrix[row-1][0] > matrix[row][0] {
				matrix[row-1], matrix[row] = matrix[row], matrix[row-1]
			}
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 77, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix77() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var indexes = make([]int, n)
	for index := 0; index < n; index++ {
		indexes[index] = index
	}
	for index := 0; index < n-1; index++ {
		for col := 1; col < n-index; col++ {
			if matrix[m-1][indexes[col-1]] < matrix[m-1][indexes[col]] {
				indexes[col-1], indexes[col] = indexes[col], indexes[col-1]
			}
		}
	}
	for row := 0; row < m; row++ {
		var tempRow = make([]float64, n)
		for col := 0; col < n; col++ {
			tempRow[col] = matrix[row][indexes[col]]
		}
		for col := 0; col < n; col++ {
			matrix[row][col] = tempRow[col]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 78, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix78() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	var minimalsIndexes = make([]int, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
			if col == 0 || matrix[row][col] < matrix[row][minimalsIndexes[row]] {
				minimalsIndexes[row] = col
			}
		}
	}
	for index := 0; index < m-1; index++ {
		for row := 1; row < m-index; row++ {
			if matrix[row-1][minimalsIndexes[row-1]] < matrix[row][minimalsIndexes[row]] {
				matrix[row-1], matrix[row] = matrix[row], matrix[row-1]
				minimalsIndexes[row-1], minimalsIndexes[row] = minimalsIndexes[row], minimalsIndexes[row-1]
			}
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 79, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix79() bool {
	var m = mt.scanner.NextInt()
	var n = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	var maximalsIndexes = make([]int, n)
	for col := 0; col < n; col++ {
		for row := 0; row < m; row++ {
			if row == 0 || matrix[row][col] > matrix[maximalsIndexes[col]][col] {
				maximalsIndexes[col] = row
			}
		}
	}
	for index := 0; index < n-1; index++ {
		for col := 1; col < n-index; col++ {
			if matrix[maximalsIndexes[col-1]][col-1] > matrix[maximalsIndexes[col]][col] {
				for row := 0; row < m; row++ {
					matrix[row][col-1], matrix[row][col] = matrix[row][col], matrix[row][col-1]
				}
				maximalsIndexes[col-1], maximalsIndexes[col] = maximalsIndexes[col], maximalsIndexes[col-1]
			}
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 80, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix80() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	var sum = 0.0
	for index := 0; index < m; index++ {
		sum += a[index][index]
	}
	return mt.checker.CompareFloat64(2, sum)
}

// { "no": 81, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix81() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	var sum = 0.0
	for index := 0; index < m; index++ {
		sum += a[index][m-1-index]
	}
	var average = sum / float64(m)
	return mt.checker.CompareFloat64(2, average)
}

// { "no": 82, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix82() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	var sum float64
	for index := m - 1; index >= 0; index-- {
		sum = 0
		for row, col := 0, index; col < m; row, col = row+1, col+1 {
			sum += a[row][col]
		}
		if !mt.checker.CompareFloat64(2, sum) {
			return false
		}
	}
	for index := 1; index < m; index++ {
		sum = 0
		for row, col := index, 0; row < m; row, col = row+1, col+1 {
			sum += a[row][col]
		}
		if !mt.checker.CompareFloat64(2, sum) {
			return false
		}
	}
	return true
}

// { "no": 83, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix83() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	var sum float64
	for index := 0; index < m; index++ {
		sum = 0
		for row, col := 0, index; col >= 0; row, col = row+1, col-1 {
			sum += a[row][col]
		}
		if !mt.checker.CompareFloat64(2, sum) {
			return false
		}
	}
	for index := 1; index < m; index++ {
		sum = 0
		for row, col := index, m-1; row < m; row, col = row+1, col-1 {
			sum += a[row][col]
		}
		if !mt.checker.CompareFloat64(2, sum) {
			return false
		}
	}
	return true
}

// { "no": 84, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix84() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	var sum, average float64
	var count int
	for index := m - 1; index >= 0; index-- {
		sum = 0
		count = 0
		for row, col := 0, index; col < m; row, col = row+1, col+1 {
			sum += a[row][col]
			count++
		}
		average = sum / float64(count)
		if !mt.checker.CompareFloat64(2, average) {
			return false
		}
	}
	for index := 1; index < m; index++ {
		sum = 0
		count = 0
		for row, col := index, 0; row < m; row, col = row+1, col+1 {
			sum += a[row][col]
			count++
		}
		average = sum / float64(count)
		if !mt.checker.CompareFloat64(2, average) {
			return false
		}
	}
	return true
}

// { "no": 85, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix85() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	var sum, average float64
	var count int
	for index := 0; index < m; index++ {
		sum = 0
		count = 0
		for row, col := 0, index; col >= 0; row, col = row+1, col-1 {
			sum += a[row][col]
			count++
		}
		average = sum / float64(count)
		if !mt.checker.CompareFloat64(2, average) {
			return false
		}
	}
	for index := 1; index < m; index++ {
		sum = 0
		count = 0
		for row, col := index, m-1; row < m; row, col = row+1, col-1 {
			sum += a[row][col]
			count++
		}
		average = sum / float64(count)
		if !mt.checker.CompareFloat64(2, average) {
			return false
		}
	}
	return true
}

// { "no": 86, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix86() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	var minimal float64
	for index := m - 1; index >= 0; index-- {
		for row, col := 0, index; col < m; row, col = row+1, col+1 {
			if row == 0 || a[row][col] < minimal {
				minimal = a[row][col]
			}
		}
		if !mt.checker.CompareFloat64(2, minimal) {
			return false
		}
	}
	for index := 1; index < m; index++ {
		for row, col := index, 0; row < m; row, col = row+1, col+1 {
			if col == 0 || a[row][col] < minimal {
				minimal = a[row][col]
			}
		}
		if !mt.checker.CompareFloat64(2, minimal) {
			return false
		}
	}
	return true
}

// { "no": 87, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix87() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	var maximal float64
	for index := 0; index < m; index++ {
		for row, col := 0, index; col >= 0; row, col = row+1, col-1 {
			if row == 0 || a[row][col] > maximal {
				maximal = a[row][col]
			}
		}
		if !mt.checker.CompareFloat64(2, maximal) {
			return false
		}
	}
	for index := 1; index < m; index++ {
		for row, col := index, m-1; row < m; row, col = row+1, col-1 {
			if col == m-1 || a[row][col] > maximal {
				maximal = a[row][col]
			}
		}
		if !mt.checker.CompareFloat64(2, maximal) {
			return false
		}
	}
	return true
}

// { "no": 88, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix88() bool {
	var m = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := 0; row < m; row++ {
		for col := row - 1; col >= 0; col-- {
			matrix[row][col] = 0
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 89, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix89() bool {
	var m = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := 0; row < m; row++ {
		for col := m - 2 - row; col >= 0; col-- {
			matrix[row][col] = 0
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 90, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix90() bool {
	var m = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := 0; row < m; row++ {
		for col := m - 1 - row; col < m; col++ {
			matrix[row][col] = 0
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 91, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix91() bool {
	var m = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := 0; row < m; row++ {
		for col := row; col < m; col++ {
			matrix[row][col] = 0
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 92, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix92() bool {
	var m = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := 0; row < m/2; row++ {
		for col := row + 1; col < m-1-row; col++ {
			matrix[row][col] = 0
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 93, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix93() bool {
	var m = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for col := m - 1; col >= m/2; col-- {
		for row := m - col; row < col; row++ {
			matrix[row][col] = 0
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 94, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix94() bool {
	var m = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for col := 0; col <= m/2; col++ {
		for row := col; row < m-col; row++ {
			matrix[row][col] = 0
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 95, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix95() bool {
	var m = mt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			matrix[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := m - 1; row >= m/2; row-- {
		for col := m - row - 1; col <= row; col++ {
			matrix[row][col] = 0
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 96, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix96() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := 0; row < m; row++ {
		for col := row; col < m; col++ {
			a[row][col], a[col][row] = a[col][row], a[row][col]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, a[row]...) {
			return false
		}
	}
	return true
}

// { "no": 97, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix97() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	for index := 0; index < m; index++ {
		for row, col := index, m-1-index; row < m; row, col = row+1, col-1 {
			a[row][m-1-index], a[index][col] = a[index][col], a[row][m-1-index]
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, a[row]...) {
			return false
		}
	}
	return true
}

// { "no": 98, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix98() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	for index := 0; index < m/2; index++ {
		row, col := index, index
		for col < m-index {
			a[row][col], a[m-1-row][m-1-col] = a[m-1-row][m-1-col], a[row][col]
			col++
		}
		row++
		col--
		for row < m-1-index {
			a[row][col], a[m-1-row][m-1-col] = a[m-1-row][m-1-col], a[row][col]
			row++
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, a[row]...) {
			return false
		}
	}
	return true
}

// { "no": 99, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix99() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := 0; row < m/2; row++ {
		col := row
		for col < m-1-row {
			a[row][col], a[col][m-1-row], a[m-1-row][m-1-col], a[m-1-col][row] =
				a[col][m-1-row], a[m-1-row][m-1-col], a[m-1-col][row], a[row][col]
			col++
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, a[row]...) {
			return false
		}
	}
	return true
}

// { "no": 100, "dat": "2", "ans": "2" }
func (mt *matrixTask) matrix100() bool {
	var m = mt.scanner.NextInt()
	var a = make([][]float64, m)
	for row := 0; row < m; row++ {
		a[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			a[row][col] = mt.scanner.NextFloat64()
		}
	}
	for row := 0; row < m/2; row++ {
		col := row
		for col < m-1-row {

			a[row][col], a[m-1-col][row], a[m-1-row][m-1-col], a[col][m-1-row] =
				a[m-1-col][row], a[m-1-row][m-1-col], a[col][m-1-row], a[row][col]

			col++
		}
	}
	for row := 0; row < m; row++ {
		if !mt.checker.CompareFloat64(2, a[row]...) {
			return false
		}
	}
	return true
}
