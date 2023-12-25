package task

import (
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/param"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type paramTask struct {
	dataPath  string
	name      string
	count     int
	byteOrder binary.ByteOrder
	scanner   scanner.Scanner
	checker   checker.Checker
}

func NewParamTask(pathPrefix string, byteOrder binary.ByteOrder) Task {
	return &paramTask{
		dataPath:  pathPrefix + "/16param/Param%03d/%03d/%03d",
		name:      "Param",
		count:     70,
		byteOrder: byteOrder,
	}
}

func (pt *paramTask) GetCount() int { return pt.count }

func (pt *paramTask) GetName() string { return pt.name }

func (pt *paramTask) SetData(taskNo int, testNo int) error {
	var filePath = fmt.Sprintf(pt.dataPath, taskNo, testNo, testNo)
	scanner, err := scanner.New(filePath + ".dat")
	if err != nil {
		return errors.Wrap(err, "scanner.New")
	}
	checker, err := checker.New(filePath + ".ans")
	if err != nil {
		return errors.Wrap(err, "checker.New")
	}
	pt.scanner = scanner
	pt.checker = checker
	return nil
}

func (pt *paramTask) ScannerEOF() bool { return pt.scanner.EOF() }

func (pt *paramTask) CheckerEOF() bool { return pt.checker.EOF() }

func (pt *paramTask) CleanData() {
	pt.scanner.RemoveFiles()
	pt.checker.RemoveFiles()
}

func (pt *paramTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return pt.param1
	case 2:
		return pt.param2
	case 3:
		return pt.param3
	case 4:
		return pt.param4
	case 5:
		return pt.param5
	case 6:
		return pt.param6
	case 7:
		return pt.param7
	case 8:
		return pt.param8
	case 9:
		return pt.param9
	case 10:
		return pt.param10
	case 11:
		return pt.param11
	case 12:
		return pt.param12
	case 13:
		return pt.param13
	case 14:
		return pt.param14
	case 15:
		return pt.param15
	case 16:
		return pt.param16
	case 17:
		return pt.param17
	case 18:
		return pt.param18
	case 19:
		return pt.param19
	case 20:
		return pt.param20
	case 21:
		return pt.param21
	case 22:
		return pt.param22
	case 23:
		return pt.param23
	case 24:
		return pt.param24
	case 25:
		return pt.param25
	case 26:
		return pt.param26
	case 27:
		return pt.param27
	case 28:
		return pt.param28
	case 29:
		return pt.param29
	case 30:
		return pt.param30
	case 31:
		return pt.param31
	case 32:
		return pt.param32
	case 33:
		return pt.param33
	case 34:
		return pt.param34
	case 35:
		return pt.param35
	case 36:
		return pt.param36
	case 37:
		return pt.param37
	case 38:
		return pt.param38
	case 39:
		return pt.param39
	case 40:
		return pt.param40
	case 41:
		return pt.param41
	case 42:
		return pt.param42
	case 43:
		return pt.param43
	case 44:
		return pt.param44
	case 45:
		return pt.param45
	case 46:
		return pt.param46
	case 47:
		return pt.param47
	case 48:
		return pt.param48
	case 49:
		return pt.param49
	case 50:
		return pt.param50
	case 51:
		return pt.param51
	case 52:
		return pt.param52
	case 53:
		return pt.param53
	case 54:
		return pt.param54
	case 55:
		return pt.param55
	case 56:
		return pt.param56
	case 57:
		return pt.param57
	case 58:
		return pt.param58
	case 59:
		return pt.param59
	case 60:
		return pt.param60
	case 61:
		return pt.param61
	case 62:
		return pt.param62
	case 63:
		return pt.param63
	case 64:
		return pt.param64
	case 65:
		return pt.param65
	case 66:
		return pt.param66
	case 67:
		return pt.param67
	case 68:
		return pt.param68
	case 69:
		return pt.param69
	case 70:
		return pt.param70
	default:
		return nil
	}
}

// { "no": 1, "dat": "", "ans": "" }
func (pt *paramTask) param1() bool {
	var (
		n, minimal int
		array      []int
	)
	for index := 0; index < 3; index++ {
		n = pt.scanner.NextInt()
		array = make([]int, n)
		for i := 0; i < n; i++ {
			array[i] = pt.scanner.NextInt()
		}
		minimal = param.MinElem(array, n)
		if !pt.checker.CompareInt(minimal) {
			return false
		}
	}
	return true
}

// { "no": 2, "dat": "2", "ans": "" }
func (pt *paramTask) param2() bool {
	var (
		n, num int
		array  []float64
	)
	for index := 0; index < 3; index++ {
		n = pt.scanner.NextInt()
		array = make([]float64, n)
		for i := 0; i < n; i++ {
			array[i] = pt.scanner.NextFloat64()
		}
		num = param.MaxNum(array, n)
		if !pt.checker.CompareInt(num) {
			return false
		}
	}
	return true
}

// { "no": 3, "dat": "2", "ans": "" }
func (pt *paramTask) param3() bool {
	var (
		n, nMin, nMax int
		array         []float64
	)
	for index := 0; index < 3; index++ {
		n = pt.scanner.NextInt()
		array = make([]float64, n)
		for i := 0; i < n; i++ {
			array[i] = pt.scanner.NextFloat64()
		}
		param.MinmaxNum(array, n, &nMin, &nMax)
		if !pt.checker.CompareInt(nMin, nMax) {
			return false
		}
	}
	return true
}

// { "no": 4, "dat": "2", "ans": "2" }
func (pt *paramTask) param4() bool {
	var (
		n     int
		array []float64
	)
	for index := 0; index < 3; index++ {
		n = pt.scanner.NextInt()
		array = make([]float64, n)
		for i := 0; i < n; i++ {
			array[i] = pt.scanner.NextFloat64()
		}
		param.Inv(array, n)
		if !pt.checker.CompareFloat64(2, array...) {
			return false
		}
	}
	return true
}

// { "no": 5, "dat": "2", "ans": "2" }
func (pt *paramTask) param5() bool {
	var n = pt.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = pt.scanner.NextFloat64()
	}
	for index := 0; index < 5; index++ {
		param.Smooth1(array, n)
		if !pt.checker.CompareFloat64(2, array...) {
			return false
		}
	}
	return true
}

// { "no": 6, "dat": "2", "ans": "2" }
func (pt *paramTask) param6() bool {
	var n = pt.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = pt.scanner.NextFloat64()
	}
	for index := 0; index < 5; index++ {
		param.Smooth2(array, n)
		if !pt.checker.CompareFloat64(2, array...) {
			return false
		}
	}
	return true
}

// { "no": 7, "dat": "2", "ans": "2" }
func (pt *paramTask) param7() bool {
	var n = pt.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = pt.scanner.NextFloat64()
	}
	for index := 0; index < 5; index++ {
		param.Smooth3(array, n)
		if !pt.checker.CompareFloat64(2, array...) {
			return false
		}
	}
	return true
}

// { "no": 8, "dat": "", "ans": "" }
func (pt *paramTask) param8() bool {
	var (
		x, n  int
		array []int
	)
	for index := 0; index < 3; index++ {
		x = pt.scanner.NextInt()
		n = pt.scanner.NextInt()
		array = make([]int, n)
		for i := 0; i < n; i++ {
			array[i] = pt.scanner.NextInt()
		}
		param.RemoveX(&array, &n, x)
		if !pt.checker.CompareInt(n) || !pt.checker.CompareInt(array...) {
			return false
		}
	}
	return true
}

// { "no": 9, "dat": "2", "ans": "2" }
func (pt *paramTask) param9() bool {
	var (
		n     int
		array []float64
	)
	for index := 0; index < 3; index++ {
		n = pt.scanner.NextInt()
		array = make([]float64, n)
		for i := 0; i < n; i++ {
			array[i] = pt.scanner.NextFloat64()
		}
		param.RemoveForInc(&array, &n)
		if !pt.checker.CompareInt(n) || !pt.checker.CompareFloat64(2, array...) {
			return false
		}
	}
	return true
}

// { "no": 10, "dat": "", "ans": "" }
func (pt *paramTask) param10() bool {
	var (
		x, n  int
		array []int
	)
	for index := 0; index < 3; index++ {
		x = pt.scanner.NextInt()
		n = pt.scanner.NextInt()
		array = make([]int, n)
		for i := 0; i < n; i++ {
			array[i] = pt.scanner.NextInt()
		}
		param.DoubleX(&array, &n, x)
		if !pt.checker.CompareInt(n) || !pt.checker.CompareInt(array...) {
			return false
		}
	}
	return true
}

// { "no": 11, "dat": "2", "ans": "2" }
func (pt *paramTask) param11() bool {
	var (
		n     int
		array []float64
	)
	for index := 0; index < 3; index++ {
		n = pt.scanner.NextInt()
		array = make([]float64, n)
		for i := 0; i < n; i++ {
			array[i] = pt.scanner.NextFloat64()
		}
		param.SortArray(array, n)
		if !pt.checker.CompareFloat64(2, array...) {
			return false
		}
	}
	return true
}

// { "no": 12, "dat": "2", "ans": "" }
func (pt *paramTask) param12() bool {
	var (
		n       int
		array   []float64
		indexes []int
	)
	for index := 0; index < 3; index++ {
		n = pt.scanner.NextInt()
		array = make([]float64, n)
		for i := 0; i < n; i++ {
			array[i] = pt.scanner.NextFloat64()
		}
		param.SortIndex(array, n, &indexes)
		if !pt.checker.CompareInt(indexes...) {
			return false
		}
	}
	return true
}

// { "no": 13, "dat": "2", "ans": "2" }
func (pt *paramTask) param13() bool {
	var (
		n     int
		array []float64
	)
	for index := 0; index < 3; index++ {
		n = pt.scanner.NextInt()
		array = make([]float64, n)
		for i := 0; i < n; i++ {
			array[i] = pt.scanner.NextFloat64()
		}
		param.Hill(array, n)
		if !pt.checker.CompareFloat64(2, array...) {
			return false
		}
	}
	return true
}

// { "no": 14, "dat": "2", "ans": "2" }
func (pt *paramTask) param14() bool {
	var (
		na, nb, nc int
		a, b, c    []float64
	)
	na = pt.scanner.NextInt()
	a = make([]float64, na)
	for index := 0; index < na; index++ {
		a[index] = pt.scanner.NextFloat64()
	}
	param.Split1(a, na, &b, &nb, &c, &nc)
	if !pt.checker.CompareInt(nb) || !pt.checker.CompareFloat64(2, b...) {
		return false
	}
	if !pt.checker.CompareInt(nc) || !pt.checker.CompareFloat64(2, c...) {
		return false
	}
	return true
}

// { "no": 15, "dat": "", "ans": "" }
func (pt *paramTask) param15() bool {
	var (
		na, nb, nc int
		a, b, c    []int
	)
	na = pt.scanner.NextInt()
	a = make([]int, na)
	for index := 0; index < na; index++ {
		a[index] = pt.scanner.NextInt()
	}
	param.Split2(a, na, &b, &nb, &c, &nc)
	if !pt.checker.CompareInt(nb) || !pt.checker.CompareInt(b...) {
		return false
	}
	if !pt.checker.CompareInt(nc) || !pt.checker.CompareInt(c...) {
		return false
	}
	return true
}

// { "no": 16, "dat": "2", "ans": "2" }
func (pt *paramTask) param16() bool {
	var (
		k, m, n int
		array   []float64
		matrix  [][]float64
	)
	k = pt.scanner.NextInt()
	array = make([]float64, k)
	for index := 0; index < k; index++ {
		array[index] = pt.scanner.NextFloat64()
	}
	m = pt.scanner.NextInt()
	n = pt.scanner.NextInt()
	param.ArrayToMatrRow(array, k, m, n, &matrix)
	for row := 0; row < m; row++ {
		if !pt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 17, "dat": "2", "ans": "2" }
func (pt *paramTask) param17() bool {
	var (
		k, m, n int
		array   []float64
		matrix  [][]float64
	)
	k = pt.scanner.NextInt()
	array = make([]float64, k)
	for index := 0; index < k; index++ {
		array[index] = pt.scanner.NextFloat64()
	}
	m = pt.scanner.NextInt()
	n = pt.scanner.NextInt()
	param.ArrayToMatrCol(array, k, m, n, &matrix)
	for row := 0; row < m; row++ {
		if !pt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 18, "dat": "", "ans": "" }
func (pt *paramTask) param18() bool {
	var m = pt.scanner.NextInt()
	var n = pt.scanner.NextInt()
	var matrix [][]int
	param.Chessboard(m, n, &matrix)
	for row := 0; row < m; row++ {
		if !pt.checker.CompareInt(matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 19, "dat": "2", "ans": "2" }
func (pt *paramTask) param19() bool {
	var m = pt.scanner.NextInt()
	var n = pt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = pt.scanner.NextFloat64()
		}
		if !pt.checker.CompareFloat64(2, param.Norm1(matrix, row+1, n)) {
			return false
		}
	}
	return true
}

// { "no": 20, "dat": "2", "ans": "2" }
func (pt *paramTask) param20() bool {
	var m = pt.scanner.NextInt()
	var n = pt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = pt.scanner.NextFloat64()
		}
	}
	for row := 0; row < m; row++ {
		if !pt.checker.CompareFloat64(2, param.Norm2(matrix, row+1, n)) {
			return false
		}
	}
	return true
}

// { "no": 21, "dat": "2", "ans": "2" }
func (pt *paramTask) param21() bool {
	var m = pt.scanner.NextInt()
	var n = pt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = pt.scanner.NextFloat64()
		}
	}
	var (
		k   int
		sum float64
	)
	for index := 0; index < 3; index++ {
		k = pt.scanner.NextInt()
		sum = param.SumRow(matrix, m, n, k)
		if !pt.checker.CompareFloat64(2, sum) {
			return false
		}
	}
	return true
}

// { "no": 22, "dat": "2", "ans": "2" }
func (pt *paramTask) param22() bool {
	var m = pt.scanner.NextInt()
	var n = pt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = pt.scanner.NextFloat64()
		}
	}
	var (
		k   int
		sum float64
	)
	for index := 0; index < 3; index++ {
		k = pt.scanner.NextInt()
		sum = param.SumCol(matrix, m, n, k)
		if !pt.checker.CompareFloat64(2, sum) {
			return false
		}
	}
	return true
}

// { "no": 23, "dat": "2", "ans": "2" }
func (pt *paramTask) param23() bool {
	var m = pt.scanner.NextInt()
	var n = pt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = pt.scanner.NextFloat64()
		}
	}
	var k1 = pt.scanner.NextInt()
	var k2 = pt.scanner.NextInt()
	param.SwapRow(matrix, m, n, k1, k2)
	for row := 0; row < m; row++ {
		if !pt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 24, "dat": "2", "ans": "2" }
func (pt *paramTask) param24() bool {
	var m = pt.scanner.NextInt()
	var n = pt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = pt.scanner.NextFloat64()
		}
	}
	var k1 = pt.scanner.NextInt()
	var k2 = pt.scanner.NextInt()
	param.SwapCol(matrix, m, n, k1, k2)
	for row := 0; row < m; row++ {
		if !pt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 25, "dat": "2", "ans": "2" }
func (pt *paramTask) param25() bool {
	var m = pt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, m)
		for col := 0; col < m; col++ {
			matrix[row][col] = pt.scanner.NextFloat64()
		}
	}
	param.Transp(matrix, m)
	for row := 0; row < m; row++ {
		if !pt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 26, "dat": "2", "ans": "2" }
func (pt *paramTask) param26() bool {
	var m = pt.scanner.NextInt()
	var n = pt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = pt.scanner.NextFloat64()
		}
	}
	var k1 = pt.scanner.NextInt()
	var k2 = pt.scanner.NextInt()
	param.RemoveRows(&matrix, &m, n, k1, k2)
	if !pt.checker.CompareInt(m, n) {
		return false
	}
	for row := 0; row < m; row++ {
		if !pt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 27, "dat": "2", "ans": "2" }
func (pt *paramTask) param27() bool {
	var m = pt.scanner.NextInt()
	var n = pt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = pt.scanner.NextFloat64()
		}
	}
	var k1 = pt.scanner.NextInt()
	var k2 = pt.scanner.NextInt()
	param.RemoveCols(&matrix, m, &n, k1, k2)
	if !pt.checker.CompareInt(m, n) {
		return false
	}
	for row := 0; row < m; row++ {
		if !pt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 28, "dat": "2", "ans": "2" }
func (pt *paramTask) param28() bool {
	var m = pt.scanner.NextInt()
	var n = pt.scanner.NextInt()
	var matrix = make([][]float64, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = pt.scanner.NextFloat64()
		}
	}
	var k = pt.scanner.NextInt()
	var l = pt.scanner.NextInt()
	param.RemoveRowCol(&matrix, &m, &n, k, l)
	if !pt.checker.CompareInt(m, n) {
		return false
	}
	for row := 0; row < m; row++ {
		if !pt.checker.CompareFloat64(2, matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 29, "dat": "", "ans": "" }
func (pt *paramTask) param29() bool {
	var m = pt.scanner.NextInt()
	var n = pt.scanner.NextInt()
	var matrix = make([][]int, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		for col := 0; col < n; col++ {
			matrix[row][col] = pt.scanner.NextInt()
		}
	}
	param.SortCols(matrix, m, n)
	for row := 0; row < m; row++ {
		if !pt.checker.CompareInt(matrix[row]...) {
			return false
		}
	}
	return true
}

// { "no": 30, "dat": "", "ans": "" }
func (pt *paramTask) param30() bool {
	var str string
	for index := 0; index < 5; index++ {
		str = pt.scanner.NextLine()
		if !pt.checker.CompareInt(param.IsIdent(str)) {
			return false
		}
	}
	return true
}

// { "no": 31, "dat": "", "ans": "" }
func (pt *paramTask) param31() bool {
	var n = pt.scanner.NextInt()
	pt.scanner.Skip()
	var str, result string
	for index := 0; index < 5; index++ {
		str = pt.scanner.NextLine()
		result = param.FillStr(str, n)
		if !pt.checker.CompareLine(result) {
			return false
		}
	}
	return true
}

// { "no": 32, "dat": "", "ans": "" }
func (pt *paramTask) param32() bool {
	var str string
	for index := 0; index < 5; index++ {
		str = pt.scanner.NextLine()
		param.UpCaseLat(&str)
		if !pt.checker.CompareLine(str) {
			return false
		}
	}
	return true
}

// { "no": 33, "dat": "", "ans": "" }
func (pt *paramTask) param33() bool {
	var str string
	for index := 0; index < 5; index++ {
		str = pt.scanner.NextLine()
		param.LowCaseLat(&str)
		if !pt.checker.CompareLine(str) {
			return false
		}
	}
	return true
}

// { "no": 34, "dat": "", "ans": "" }
func (pt *paramTask) param34() bool {
	var c, _ = pt.scanner.NextRune()
	pt.scanner.Skip()
	var str string
	for index := 0; index < 5; index++ {
		str = pt.scanner.NextLine()
		param.TrimLeftC(&str, c)
		if !pt.checker.CompareLine(str) {
			return false
		}
	}
	return true
}

// { "no": 35, "dat": "", "ans": "" }
func (pt *paramTask) param35() bool {
	var c, _ = pt.scanner.NextRune()
	pt.scanner.Skip()
	var str string
	for index := 0; index < 5; index++ {
		str = pt.scanner.NextLine()
		param.TrimRightC(&str, c)
		if !pt.checker.CompareLine(str) {
			return false
		}
	}
	return true
}

// { "no": 36, "dat": "", "ans": "" }
func (pt *paramTask) param36() bool {
	var str = pt.scanner.NextLine()
	var n, k int
	var s string
	for index := 0; index < 3; index++ {
		k = pt.scanner.NextInt()
		n = pt.scanner.NextInt()
		s = param.InvStr(str, k, n)
		if !pt.checker.CompareLine(s) {
			return false
		}
	}
	return true
}

// { "no": 37, "dat": "", "ans": "" }
func (pt *paramTask) param37() bool {
	var s0 = pt.scanner.NextLine()
	var s = pt.scanner.NextLine()
	var k, n, position int
	for index := 0; index < 3; index++ {
		k = pt.scanner.NextInt()
		n = pt.scanner.NextInt()
		position = param.PosSub(s0, s, k, n)
		if !pt.checker.CompareInt(position) {
			return false
		}
	}
	return true
}

// { "no": 38, "dat": "", "ans": "" }
func (pt *paramTask) param38() bool {
	var s0, s string
	var position int
	for index := 0; index < 5; index++ {
		s0 = pt.scanner.NextLine()
		s = pt.scanner.NextLine()
		position = param.PosLast(s0, s)
		if !pt.checker.CompareInt(position) {
			return false
		}
	}
	return true
}

// { "no": 39, "dat": "", "ans": "" }
func (pt *paramTask) param39() bool {
	var s0, s string
	var k, position int
	for index := 0; index < 5; index++ {
		s0 = pt.scanner.NextLine()
		s = pt.scanner.NextLine()
		k = pt.scanner.NextInt()
		pt.scanner.Skip()
		position = param.PosK(s0, s, k)
		if !pt.checker.CompareInt(position) {
			return false
		}
	}
	return true
}

// { "no": 40, "dat": "", "ans": "" }
func (pt *paramTask) param40() bool {
	var s = pt.scanner.NextLine()
	var k int
	var word string
	for index := 0; index < 3; index++ {
		k = pt.scanner.NextInt()
		word = param.WordK(s, k)
		if !pt.checker.CompareLine(word) {
			return false
		}
	}
	return true
}

// { "no": 41, "dat": "", "ans": "" }
func (pt *paramTask) param41() bool {
	var s = pt.scanner.NextLine()
	var (
		words []string
		n     int
	)
	param.SplitStr(s, &words, &n)
	if !pt.checker.CompareInt(n) {
		return false
	}
	return pt.checker.CompareWord(words...)
}

// { "no": 42, "dat": "", "ans": "" }
func (pt *paramTask) param42() bool {
	var str string
	for index := 0; index < 5; index++ {
		str = pt.scanner.NextLine()
		str = param.CompressStr(str)
		if !pt.checker.CompareLine(str) {
			return false
		}
	}
	return true
}

// { "no": 43, "dat": "", "ans": "" }
func (pt *paramTask) param43() bool {
	var str string
	for index := 0; index < 5; index++ {
		str = pt.scanner.NextLine()
		str = param.DecompressStr(str)
		if !pt.checker.CompareLine(str) {
			return false
		}
	}
	return true
}

// { "no": 44, "dat": "", "ans": "" }
func (pt *paramTask) param44() bool {
	var (
		n   int
		bin string
	)
	for index := 0; index < 5; index++ {
		n = pt.scanner.NextInt()
		bin = param.DecToBin(n)
		if !pt.checker.CompareLine(bin) {
			return false
		}
	}
	return true
}

// { "no": 45, "dat": "", "ans": "" }
func (pt *paramTask) param45() bool {
	var (
		n   int
		hex string
	)
	for index := 0; index < 5; index++ {
		n = pt.scanner.NextInt()
		hex = param.DecToHex(n)
		if !pt.checker.CompareLine(hex) {
			return false
		}
	}
	return true
}

// { "no": 46, "dat": "", "ans": "" }
func (pt *paramTask) param46() bool {
	var (
		n   int
		bin string
	)
	for index := 0; index < 5; index++ {
		bin = pt.scanner.NextLine()
		n = param.BinToDec(bin)
		if !pt.checker.CompareInt(n) {
			return false
		}
	}
	return true
}

// { "no": 47, "dat": "", "ans": "" }
func (pt *paramTask) param47() bool {
	var (
		n   int
		hex string
	)
	for index := 0; index < 5; index++ {
		hex = pt.scanner.NextLine()
		n = param.HexToDec(hex)
		if !pt.checker.CompareInt(n) {
			return false
		}
	}
	return true
}

// { "no": 48, "dat": "", "ans": "" }
func (pt *paramTask) param48() bool {
	var (
		fileName string
		size     int64
	)
	for index := 0; index < 3; index++ {
		fileName = pt.scanner.NextBinaryFileName(pt.byteOrder)
		size = param.IntFileSize(fileName)
		if !pt.checker.CompareInt64(size) {
			return false
		}
	}
	return true
}

// { "no": 49, "dat": "", "ans": "" }
func (pt *paramTask) param49() bool {
	var (
		fileName string
		count    int
	)
	for index := 0; index < 3; index++ {
		fileName = pt.scanner.NextTextFileName()
		count = param.LineCount(fileName)
		if !pt.checker.CompareInt(count) {
			return false
		}
	}
	return true
}

// { "no": 50, "dat": "", "ans": "" }
func (pt *paramTask) param50() bool {
	var fileName string
	for index := 0; index < 3; index++ {
		fileName = pt.scanner.NextBinaryFileName(pt.byteOrder)
		param.InvIntFile(fileName, pt.byteOrder)
		if !pt.checker.CompareBinaryFile(pt.byteOrder, fileName) {
			return false
		}
	}
	return true
}

// { "no": 51, "dat": "", "ans": "" }
func (pt *paramTask) param51() bool {
	var fileName = pt.scanner.NextTextFileName()
	var n = pt.scanner.NextInt()
	var k = pt.scanner.NextInt()
	var l = pt.scanner.NextInt()
	param.AddLineNumbers(fileName, n, k, l)
	return pt.checker.CompareTextFile(fileName)
}

// { "no": 52, "dat": "", "ans": "" }
func (pt *paramTask) param52() bool {
	var fileName = pt.scanner.NextTextFileName()
	param.RemoveLineNumbers(fileName)
	return pt.checker.CompareTextFile(fileName)
}

// { "no": 53, "dat": "", "ans": "" }
func (pt *paramTask) param53() bool {
	var s0 = pt.scanner.NextBinaryFileName(pt.byteOrder)
	var k = pt.scanner.NextInt()
	pt.scanner.Skip()
	var s1 = pt.scanner.NextLine()
	var s2 = pt.scanner.NextLine()
	param.SplitIntFile(s0, k, s1, s2, pt.byteOrder)
	return pt.checker.CompareBinaryFile(pt.byteOrder, s1, s2)
}

// { "no": 54, "dat": "", "ans": "" }
func (pt *paramTask) param54() bool {
	var s0 = pt.scanner.NextTextFileName()
	var k = pt.scanner.NextInt()
	pt.scanner.Skip()
	var s1 = pt.scanner.NextLine()
	var s2 = pt.scanner.NextLine()
	param.SplitText(s0, k, s1, s2)
	return pt.checker.CompareTextFile(s1, s2)
}

// { "no": 55, "dat": "", "ans": "" }
func (pt *paramTask) param55() bool {
	for index := 0; index < 2; index++ {
		var fileName = pt.scanner.NextBinaryFileName(pt.byteOrder)
		param.StringFileToText(fileName, pt.byteOrder)
		if !pt.checker.CompareTextFile(fileName) {
			return false
		}
	}
	return true
}

// { "no": 56, "dat": "", "ans": "" }
func (pt *paramTask) param56() bool {
	for index := 0; index < 2; index++ {
		var fileName = pt.scanner.NextTextFileName()
		param.TextToStringFile(fileName, pt.byteOrder)
		if !pt.checker.CompareBinaryFile(pt.byteOrder, fileName) {
			return false
		}
	}
	return true
}

// { "no": 57, "dat": "", "ans": "" }
func (pt *paramTask) param57() bool {
	return false
}

// { "no": 58, "dat": "", "ans": "" }
func (pt *paramTask) param58() bool {
	return false
}

// { "no": 59, "dat": "", "ans": "" }
func (pt *paramTask) param59() bool {
	return false
}

// { "no": 60, "dat": "", "ans": "" }
func (pt *paramTask) param60() bool {
	return false
}

// { "no": 61, "dat": "", "ans": "" }
func (pt *paramTask) param61() bool {
	return false
}

// { "no": 62, "dat": "", "ans": "" }
func (pt *paramTask) param62() bool {
	return false
}

// { "no": 63, "dat": "", "ans": "" }
func (pt *paramTask) param63() bool {
	return false
}

// { "no": 64, "dat": "2", "ans": "2" }
func (pt *paramTask) param64() bool {
	return false
}

// { "no": 65, "dat": "2", "ans": "2" }
func (pt *paramTask) param65() bool {
	return false
}

// { "no": 66, "dat": "2", "ans": "2" }
func (pt *paramTask) param66() bool {
	return false
}

// { "no": 67, "dat": "2", "ans": "2" }
func (pt *paramTask) param67() bool {
	return false
}

// { "no": 68, "dat": "2", "ans": "2" }
func (pt *paramTask) param68() bool {
	return false
}

// { "no": 69, "dat": "2", "ans": "2" }
func (pt *paramTask) param69() bool {
	return false
}

// { "no": 70, "dat": "2", "ans": "2" }
func (pt *paramTask) param70() bool {
	return false
}
