package task

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type beginTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewBeginTask(pathPrefix string) Task {
	return &beginTask{
		dataPath: pathPrefix + "/01begin/Begin%03d/%03d/%03d",
		name:     "Begin",
		count:    40,
	}
}

func (bt *beginTask) GetCount() int { return bt.count }

func (bt *beginTask) GetName() string { return bt.name }

func (bt *beginTask) SetData(taskNo int, testNo int) error {
	var filePath = fmt.Sprintf(bt.dataPath, taskNo, testNo, testNo)
	scanner, err := scanner.New(filePath + ".dat")
	if err != nil {
		return errors.Wrap(err, "scanner.New")
	}
	checker, err := checker.New(filePath + ".ans")
	if err != nil {
		return errors.Wrap(err, "checker.New")
	}
	bt.scanner = scanner
	bt.checker = checker
	return nil
}

func (bt *beginTask) ScannerEOF() bool { return bt.scanner.EOF() }

func (bt *beginTask) CheckerEOF() bool { return bt.checker.EOF() }

func (bt *beginTask) CleanData() {
	bt.scanner.RemoveFiles()
	bt.checker.RemoveFiles()
}

func (bt *beginTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return bt.begin1
	case 2:
		return bt.begin2
	case 3:
		return bt.begin3
	case 4:
		return bt.begin4
	case 5:
		return bt.begin5
	case 6:
		return bt.begin6
	case 7:
		return bt.begin7
	case 8:
		return bt.begin8
	case 9:
		return bt.begin9
	case 10:
		return bt.begin10
	case 11:
		return bt.begin11
	case 12:
		return bt.begin12
	case 13:
		return bt.begin13
	case 14:
		return bt.begin14
	case 15:
		return bt.begin15
	case 16:
		return bt.begin16
	case 17:
		return bt.begin17
	case 18:
		return bt.begin18
	case 19:
		return bt.begin19
	case 20:
		return bt.begin20
	case 21:
		return bt.begin21
	case 22:
		return bt.begin22
	case 23:
		return bt.begin23
	case 24:
		return bt.begin24
	case 25:
		return bt.begin25
	case 26:
		return bt.begin26
	case 27:
		return bt.begin27
	case 28:
		return bt.begin28
	case 29:
		return bt.begin29
	case 30:
		return bt.begin30
	case 31:
		return bt.begin31
	case 32:
		return bt.begin32
	case 33:
		return bt.begin33
	case 34:
		return bt.begin34
	case 35:
		return bt.begin35
	case 36:
		return bt.begin36
	case 37:
		return bt.begin37
	case 38:
		return bt.begin38
	case 39:
		return bt.begin39
	case 40:
		return bt.begin40
	default:
		return nil
	}
}

// { "no": 1, "dat": "2", "ans": "2" }
func (bt *beginTask) begin1() bool {
	var a = bt.scanner.NextFloat64()
	var p = 4 * a
	return bt.checker.CompareFloat64(2, p)
}

// { "no": 2, "dat": "2", "ans": "2" }
func (bt *beginTask) begin2() bool {
	var a = bt.scanner.NextFloat64()
	var s = a * a
	return bt.checker.CompareFloat64(2, s)
}

// { "no": 3, "dat": "2", "ans": "2" }
func (bt *beginTask) begin3() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	var s = a * b
	var p = 2 * (a + b)
	return bt.checker.CompareFloat64(2, s, p)
}

// { "no": 4, "dat": "2", "ans": "2" }
func (bt *beginTask) begin4() bool {
	var d = bt.scanner.NextFloat64()
	const PI = 3.14
	var l = PI * d
	return bt.checker.CompareFloat64(2, l)
}

// { "no": 5, "dat": "2", "ans": "3" }
func (bt *beginTask) begin5() bool {
	var a = bt.scanner.NextFloat64()
	var v = a * a * a
	var s = 6 * a * a
	return bt.checker.CompareFloat64(3, v, s)
}

// { "no": 6, "dat": "2", "ans": "3" }
func (bt *beginTask) begin6() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	var c = bt.scanner.NextFloat64()
	var v = a * b * c
	var s = 2 * (a*b + b*c + a*c)
	return bt.checker.CompareFloat64(3, v, s)
}

// { "no": 7, "dat": "2", "ans": "3" }
func (bt *beginTask) begin7() bool {
	var r = bt.scanner.NextFloat64()
	const PI = 3.14
	var l = 2 * PI * r
	var s = PI * r * r
	return bt.checker.CompareFloat64(3, l, s)
}

// { "no": 8, "dat": "2", "ans": "2" }
func (bt *beginTask) begin8() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	var average = (a + b) / 2
	return bt.checker.CompareFloat64(2, average)
}

// { "no": 9, "dat": "2", "ans": "2" }
func (bt *beginTask) begin9() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	var gmean = math.Sqrt(a * b)
	return bt.checker.CompareFloat64(2, gmean)
}

// { "no": 10, "dat": "2", "ans": "2" }
func (bt *beginTask) begin10() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	a *= a
	b *= b
	var sum = a + b
	var dif = a - b
	var mul = a * b
	var div = a / b
	return bt.checker.CompareFloat64(2, sum, dif, mul, div)
}

// { "no": 11, "dat": "2", "ans": "2" }
func (bt *beginTask) begin11() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	a = math.Abs(a)
	b = math.Abs(b)
	var sum = a + b
	var dif = a - b
	var mul = a * b
	var div = a / b
	return bt.checker.CompareFloat64(2, sum, dif, mul, div)
}

// { "no": 12, "dat": "2", "ans": "2" }
func (bt *beginTask) begin12() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	var c = math.Sqrt(a*a + b*b)
	var p = a + b + c
	return bt.checker.CompareFloat64(2, c, p)
}

// { "no": 13, "dat": "2", "ans": "2" }
func (bt *beginTask) begin13() bool {
	var r1 = bt.scanner.NextFloat64()
	var r2 = bt.scanner.NextFloat64()
	const PI = 3.14
	var s1 = PI * r1 * r1
	var s2 = PI * r2 * r2
	var s3 = s1 - s2
	return bt.checker.CompareFloat64(2, s1, s2, s3)
}

// { "no": 14, "dat": "2", "ans": "2" }
func (bt *beginTask) begin14() bool {
	var l = bt.scanner.NextFloat64()
	const PI = 3.14
	var r = l / (2 * PI)
	var s = PI * r * r
	return bt.checker.CompareFloat64(2, r, s)
}

// { "no": 15, "dat": "2", "ans": "2" }
func (bt *beginTask) begin15() bool {
	var s = bt.scanner.NextFloat64()
	const PI = 3.14
	var d = math.Sqrt(4 * s / PI)
	var l = PI * d
	return bt.checker.CompareFloat64(2, d, l)
}

// { "no": 16, "dat": "2", "ans": "2" }
func (bt *beginTask) begin16() bool {
	var x1 = bt.scanner.NextFloat64()
	var x2 = bt.scanner.NextFloat64()
	var distance = math.Abs(x2 - x1)
	return bt.checker.CompareFloat64(2, distance)
}

// { "no": 17, "dat": "2", "ans": "2" }
func (bt *beginTask) begin17() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	var c = bt.scanner.NextFloat64()
	var ac = math.Abs(c - a)
	var bc = math.Abs(c - b)
	var sum = ac + bc
	return bt.checker.CompareFloat64(2, ac, bc, sum)
}

// { "no": 18, "dat": "2", "ans": "2" }
func (bt *beginTask) begin18() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	var c = bt.scanner.NextFloat64()
	var product = math.Abs(c-a) * math.Abs(c-b)
	return bt.checker.CompareFloat64(2, product)
}

// { "no": 19, "dat": "2", "ans": "2" }
func (bt *beginTask) begin19() bool {
	var x1 = bt.scanner.NextFloat64()
	var y1 = bt.scanner.NextFloat64()
	var x2 = bt.scanner.NextFloat64()
	var y2 = bt.scanner.NextFloat64()
	var a = math.Abs(x2 - x1)
	var b = math.Abs(y2 - y1)
	var p = 2 * (a + b)
	var s = a * b
	return bt.checker.CompareFloat64(2, p, s)
}

// { "no": 20, "dat": "2", "ans": "2" }
func (bt *beginTask) begin20() bool {
	var x1 = bt.scanner.NextFloat64()
	var y1 = bt.scanner.NextFloat64()
	var x2 = bt.scanner.NextFloat64()
	var y2 = bt.scanner.NextFloat64()
	var distance = math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return bt.checker.CompareFloat64(2, distance)
}

// { "no": 21, "dat": "2", "ans": "2" }
func (bt *beginTask) begin21() bool {
	var x1 = bt.scanner.NextFloat64()
	var y1 = bt.scanner.NextFloat64()
	var x2 = bt.scanner.NextFloat64()
	var y2 = bt.scanner.NextFloat64()
	var x3 = bt.scanner.NextFloat64()
	var y3 = bt.scanner.NextFloat64()
	var a = math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	var b = math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	var c = math.Sqrt(math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2))
	var p = a + b + c
	var halfP = p / 2
	var s = math.Sqrt(halfP * (halfP - a) * (halfP - b) * (halfP - c))
	return bt.checker.CompareFloat64(2, p, s)
}

// { "no": 22, "dat": "2", "ans": "2" }
func (bt *beginTask) begin22() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	a += b
	b = a - b
	a -= b
	return bt.checker.CompareFloat64(2, a, b)
}

// { "no": 23, "dat": "2", "ans": "2" }
func (bt *beginTask) begin23() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	var c = bt.scanner.NextFloat64()
	var d = a
	a = c
	c = b
	b = d
	return bt.checker.CompareFloat64(2, a, b, c)
}

// { "no": 24, "dat": "2", "ans": "2" }
func (bt *beginTask) begin24() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	var c = bt.scanner.NextFloat64()
	var d = a
	a = b
	b = c
	c = d
	return bt.checker.CompareFloat64(2, a, b, c)
}

// { "no": 25, "dat": "2", "ans": "2" }
func (bt *beginTask) begin25() bool {
	var x = bt.scanner.NextFloat64()
	var y = 3*math.Pow(x, 6) - 6*x*x - 7
	return bt.checker.CompareFloat64(2, y)
}

// { "no": 26, "dat": "2", "ans": "2" }
func (bt *beginTask) begin26() bool {
	var x = bt.scanner.NextFloat64()
	var thregree = math.Pow(x-3, 3)
	var y = 4*thregree*thregree - 7*thregree + 2
	return bt.checker.CompareFloat64(2, y)
}

// { "no": 27, "dat": "2", "ans": "2" }
func (bt *beginTask) begin27() bool {
	var a = bt.scanner.NextFloat64()
	var b = a * a
	if !bt.checker.CompareFloat64(2, b) {
		return false
	}
	b *= b
	if !bt.checker.CompareFloat64(2, b) {
		return false
	}
	b *= b
	return bt.checker.CompareFloat64(2, b)
}

// { "no": 28, "dat": "2", "ans": "2" }
func (bt *beginTask) begin28() bool {
	var a = bt.scanner.NextFloat64()
	var b = a * a // 2
	if !bt.checker.CompareFloat64(2, b) {
		return false
	}
	var c = b * a // 3
	if !bt.checker.CompareFloat64(2, c) {
		return false
	}
	b *= c // 5
	if !bt.checker.CompareFloat64(2, b) {
		return false
	}
	c = b * b // 10
	if !bt.checker.CompareFloat64(2, c) {
		return false
	}
	b *= c // 15
	return bt.checker.CompareFloat64(2, b)
}

// { "no": 29, "dat": "2", "ans": "2" }
func (bt *beginTask) begin29() bool {
	var deg = bt.scanner.NextFloat64()
	const PI = 3.14
	var rad = deg * PI / 180
	return bt.checker.CompareFloat64(2, rad)
}

// { "no": 30, "dat": "2", "ans": "2" }
func (bt *beginTask) begin30() bool {
	var rad = bt.scanner.NextFloat64()
	const PI = 3.14
	var deg = rad * 180 / PI
	return bt.checker.CompareFloat64(2, deg)
}

// { "no": 31, "dat": "2", "ans": "2" }
func (bt *beginTask) begin31() bool {
	var fahrenheitTemp = bt.scanner.NextFloat64()
	var centigradeTemp = (fahrenheitTemp - 32) * 5 / 9
	return bt.checker.CompareFloat64(2, centigradeTemp)
}

// { "no": 32, "dat": "2", "ans": "2" }
func (bt *beginTask) begin32() bool {
	var centigradeTemp = bt.scanner.NextFloat64()
	var fahrenheitTemp = centigradeTemp/5*9 + 32
	return bt.checker.CompareFloat64(2, fahrenheitTemp)
}

// { "no": 33, "dat": "3,2,3", "ans": "2" }
func (bt *beginTask) begin33() bool {
	var x = bt.scanner.NextFloat64()
	var a = bt.scanner.NextFloat64()
	var y = bt.scanner.NextFloat64()
	var oneKg = a / x
	var yKg = oneKg * y
	return bt.checker.CompareFloat64(2, oneKg, yKg)
}

// { "no": 34, "dat": "3,2,3,2", "ans": "2" }
func (bt *beginTask) begin34() bool {
	var x = bt.scanner.NextFloat64()
	var a = bt.scanner.NextFloat64()
	var y = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	var chocoPrice = a / x
	var candyPrice = b / y
	var expensiver = chocoPrice / candyPrice
	return bt.checker.CompareFloat64(2, chocoPrice, candyPrice, expensiver)
}

// { "no": 35, "dat": "2", "ans": "2" }
func (bt *beginTask) begin35() bool {
	var v = bt.scanner.NextFloat64()
	var u = bt.scanner.NextFloat64()
	var t1 = bt.scanner.NextFloat64()
	var t2 = bt.scanner.NextFloat64()
	var s = v*t1 + (v-u)*t2
	return bt.checker.CompareFloat64(2, s)
}

// { "no": 36, "dat": "2", "ans": "2" }
func (bt *beginTask) begin36() bool {
	var v1 = bt.scanner.NextFloat64()
	var v2 = bt.scanner.NextFloat64()
	var s = bt.scanner.NextFloat64()
	var t = bt.scanner.NextFloat64()
	s += t * (v1 + v2)
	return bt.checker.CompareFloat64(2, s)
}

// { "no": 37, "dat": "2", "ans": "2" }
func (bt *beginTask) begin37() bool {
	var v1 = bt.scanner.NextFloat64()
	var v2 = bt.scanner.NextFloat64()
	var s = bt.scanner.NextFloat64()
	var t = bt.scanner.NextFloat64()
	s = math.Abs(s - t*(v1+v2))
	return bt.checker.CompareFloat64(2, s)
}

// { "no": 38, "dat": "2", "ans": "2" }
func (bt *beginTask) begin38() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	var x = -b / a
	return bt.checker.CompareFloat64(2, x)
}

// { "no": 39, "dat": "2", "ans": "2" }
func (bt *beginTask) begin39() bool {
	var a = bt.scanner.NextFloat64()
	var b = bt.scanner.NextFloat64()
	var c = bt.scanner.NextFloat64()
	var d = b*b - 4*a*c
	var x1 = (-b - math.Sqrt(d)) / (2 * a)
	var x2 = (-b + math.Sqrt(d)) / (2 * a)
	return bt.checker.CompareFloat64(2, x1, x2)
}

// { "no": 40, "dat": "2", "ans": "2" }
func (bt *beginTask) begin40() bool {
	var a1 = bt.scanner.NextFloat64()
	var b1 = bt.scanner.NextFloat64()
	var c1 = bt.scanner.NextFloat64()
	var a2 = bt.scanner.NextFloat64()
	var b2 = bt.scanner.NextFloat64()
	var c2 = bt.scanner.NextFloat64()
	var d = a1*b2 - a2*b1
	var x = (c1*b2 - c2*b1) / d
	var y = (a1*c2 - a2*c1) / d
	return bt.checker.CompareFloat64(2, x, y)
}
