package task

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type ifTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewIfTask(pathPrefix string) Task {
	return &ifTask{
		dataPath: pathPrefix + "/04if/If%03d/%03d/%03d",
		name:     "If",
		count:    30,
	}
}

func (it *ifTask) GetCount() int { return it.count }

func (it *ifTask) GetName() string { return it.name }

func (it *ifTask) SetData(taskNo int, testNo int) error {
	var filePath = fmt.Sprintf(it.dataPath, taskNo, testNo, testNo)
	scanner, err := scanner.New(filePath + ".dat")
	if err != nil {
		return errors.Wrap(err, "scanner.New")
	}
	checker, err := checker.New(filePath + ".ans")
	if err != nil {
		return errors.Wrap(err, "checker.New")
	}
	it.scanner = scanner
	it.checker = checker
	return nil
}

func (it *ifTask) ScannerEOF() bool { return it.scanner.EOF() }

func (it *ifTask) CheckerEOF() bool { return it.checker.EOF() }

func (it *ifTask) CleanData() {
	it.scanner.RemoveFiles()
	it.checker.RemoveFiles()
}

func (it *ifTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return it.if1
	case 2:
		return it.if2
	case 3:
		return it.if3
	case 4:
		return it.if4
	case 5:
		return it.if5
	case 6:
		return it.if6
	case 7:
		return it.if7
	case 8:
		return it.if8
	case 9:
		return it.if9
	case 10:
		return it.if10
	case 11:
		return it.if11
	case 12:
		return it.if12
	case 13:
		return it.if13
	case 14:
		return it.if14
	case 15:
		return it.if15
	case 16:
		return it.if16
	case 17:
		return it.if17
	case 18:
		return it.if18
	case 19:
		return it.if19
	case 20:
		return it.if20
	case 21:
		return it.if21
	case 22:
		return it.if22
	case 23:
		return it.if23
	case 24:
		return it.if24
	case 25:
		return it.if25
	case 26:
		return it.if26
	case 27:
		return it.if27
	case 28:
		return it.if28
	case 29:
		return it.if29
	case 30:
		return it.if30
	default:
		return nil
	}
}

// { "no": 1, "dat": "", "ans": "" }
func (it *ifTask) if1() bool {
	var number = it.scanner.NextInt()
	if number > 0 {
		number -= 8
	}
	return it.checker.CompareInt(number)
}

// { "no": 2, "dat": "", "ans": "" }
func (it *ifTask) if2() bool {
	var number = it.scanner.NextInt()
	if number > 0 {
		number -= 8
	} else {
		number += 6
	}
	return it.checker.CompareInt(number)
}

// { "no": 3, "dat": "", "ans": "" }
func (it *ifTask) if3() bool {
	var number = it.scanner.NextInt()
	if number > 0 {
		number -= 8
	} else if number < 0 {
		number += 6
	} else {
		number = 10
	}
	return it.checker.CompareInt(number)
}

// { "no": 4, "dat": "", "ans": "" }
func (it *ifTask) if4() bool {
	var a = it.scanner.NextInt()
	var b = it.scanner.NextInt()
	var c = it.scanner.NextInt()
	var positivesCount int
	if a > 0 {
		positivesCount++
	}
	if b > 0 {
		positivesCount++
	}
	if c > 0 {
		positivesCount++
	}
	return it.checker.CompareInt(positivesCount)
}

// { "no": 5, "dat": "", "ans": "" }
func (it *ifTask) if5() bool {
	var a = it.scanner.NextInt()
	var b = it.scanner.NextInt()
	var c = it.scanner.NextInt()
	var positivesCount, negativesCount int
	if a > 0 {
		positivesCount++
	} else if a < 0 {
		negativesCount++
	}
	if b > 0 {
		positivesCount++
	} else if b < 0 {
		negativesCount++
	}
	if c > 0 {
		positivesCount++
	} else if c < 0 {
		negativesCount++
	}
	return it.checker.CompareInt(positivesCount, negativesCount)
}

// { "no": 6, "dat": "2", "ans": "2" }
func (it *ifTask) if6() bool {
	var a = it.scanner.NextFloat64()
	var b = it.scanner.NextFloat64()
	var larger = a
	if b > a {
		larger = b
	}
	return it.checker.CompareFloat64(2, larger)
}

// { "no": 7, "dat": "2", "ans": "" }
func (it *ifTask) if7() bool {
	var a = it.scanner.NextFloat64()
	var b = it.scanner.NextFloat64()
	var smallerIndex = 1
	if b < a {
		smallerIndex = 2
	}
	return it.checker.CompareInt(smallerIndex)
}

// { "no": 8, "dat": "2", "ans": "2" }
func (it *ifTask) if8() bool {
	var a = it.scanner.NextFloat64()
	var b = it.scanner.NextFloat64()
	var smaller, larger = a, b
	if a > b {
		smaller, larger = b, a
	}
	return it.checker.CompareFloat64(2, larger, smaller)
}

// { "no": 9, "dat": "2", "ans": "2" }
func (it *ifTask) if9() bool {
	var a = it.scanner.NextFloat64()
	var b = it.scanner.NextFloat64()
	if a > b {
		a, b = b, a
	}
	return it.checker.CompareFloat64(2, a, b)
}

// { "no": 10, "dat": "", "ans": "" }
func (it *ifTask) if10() bool {
	var a = it.scanner.NextInt()
	var b = it.scanner.NextInt()
	if a != b {
		a += b
	} else {
		a = 0
	}
	b = a
	return it.checker.CompareInt(a, b)
}

// { "no": 11, "dat": "", "ans": "" }
func (it *ifTask) if11() bool {
	var a = it.scanner.NextInt()
	var b = it.scanner.NextInt()
	if a != b {
		if b > a {
			a = b
		}
	} else {
		a = 0
	}
	b = a
	return it.checker.CompareInt(a, b)
}

// { "no": 12, "dat": "2", "ans": "2" }
func (it *ifTask) if12() bool {
	var a = it.scanner.NextFloat64()
	var b = it.scanner.NextFloat64()
	var c = it.scanner.NextFloat64()
	var minimal float64
	if a < b && a < c {
		minimal = a
	} else if b < c {
		minimal = b
	} else {
		minimal = c
	}
	return it.checker.CompareFloat64(2, minimal)
}

// { "no": 13, "dat": "2", "ans": "2" }
func (it *ifTask) if13() bool {
	var a = it.scanner.NextFloat64()
	var b = it.scanner.NextFloat64()
	var c = it.scanner.NextFloat64()
	var minimal float64
	if a < b && a < c {
		minimal = a
	} else if b < c {
		minimal = b
	} else {
		minimal = c
	}
	var maximal float64
	if a > b && a > c {
		maximal = a
	} else if b > c {
		maximal = b
	} else {
		maximal = c
	}
	var between = a + b + c - minimal - maximal
	return it.checker.CompareFloat64(2, between)
}

// { "no": 14, "dat": "2", "ans": "2" }
func (it *ifTask) if14() bool {
	var a = it.scanner.NextFloat64()
	var b = it.scanner.NextFloat64()
	var c = it.scanner.NextFloat64()
	var minimal float64
	if a < b && a < c {
		minimal = a
	} else if b < c {
		minimal = b
	} else {
		minimal = c
	}
	var maximal float64
	if a > b && a > c {
		maximal = a
	} else if b > c {
		maximal = b
	} else {
		maximal = c
	}
	return it.checker.CompareFloat64(2, minimal, maximal)
}

// { "no": 15, "dat": "2", "ans": "2" }
func (it *ifTask) if15() bool {
	var a = it.scanner.NextFloat64()
	var b = it.scanner.NextFloat64()
	var c = it.scanner.NextFloat64()
	var minimal float64
	if a < b && a < c {
		minimal = a
	} else if b < c {
		minimal = b
	} else {
		minimal = c
	}
	var sum = a + b + c - minimal
	return it.checker.CompareFloat64(2, sum)
}

// { "no": 16, "dat": "2", "ans": "2" }
func (it *ifTask) if16() bool {
	var a = it.scanner.NextFloat64()
	var b = it.scanner.NextFloat64()
	var c = it.scanner.NextFloat64()
	var multiplier float64
	if a < b && b < c {
		multiplier = 2
	} else {
		multiplier = -1
	}
	a *= multiplier
	b *= multiplier
	c *= multiplier
	return it.checker.CompareFloat64(2, a, b, c)
}

// { "no": 17, "dat": "2", "ans": "2" }
func (it *ifTask) if17() bool {
	var a = it.scanner.NextFloat64()
	var b = it.scanner.NextFloat64()
	var c = it.scanner.NextFloat64()
	var multiplier float64
	if a < b && b < c || a > b && b > c {
		multiplier = 2
	} else {
		multiplier = -1
	}
	a *= multiplier
	b *= multiplier
	c *= multiplier
	return it.checker.CompareFloat64(2, a, b, c)
}

// { "no": 18, "dat": "", "ans": "" }
func (it *ifTask) if18() bool {
	var a = it.scanner.NextInt()
	var b = it.scanner.NextInt()
	var c = it.scanner.NextInt()
	var differIndex int
	if a == b {
		differIndex = 3
	} else if a == c {
		differIndex = 2
	} else {
		differIndex = 1
	}
	return it.checker.CompareInt(differIndex)
}

// { "no": 19, "dat": "", "ans": "" }
func (it *ifTask) if19() bool {
	var a = it.scanner.NextInt()
	var b = it.scanner.NextInt()
	var c = it.scanner.NextInt()
	var d = it.scanner.NextInt()
	_ = d
	var differIndex = 1
	if a == b {
		// differs are c OR d;
		if a == c {
			differIndex = 4
		} else {
			differIndex = 3
		}
	} else {
		// differs are a OR b;
		if a == c {
			differIndex = 2
		} else {
			differIndex = 1
		}
	}
	return it.checker.CompareInt(differIndex)
}

// { "no": 20, "dat": "2", "ans": "2" }
func (it *ifTask) if20() bool {
	var a = it.scanner.NextFloat64()
	var b = it.scanner.NextFloat64()
	var c = it.scanner.NextFloat64()
	var ab = math.Abs(b - a)
	var ac = math.Abs(c - a)
	var point, distance = b, ab
	if ac < ab {
		point = c
		distance = ac
	}
	return it.checker.CompareFloat64(2, point, distance)
}

// { "no": 21, "dat": "", "ans": "" }
func (it *ifTask) if21() bool {
	var x = it.scanner.NextInt()
	var y = it.scanner.NextInt()
	var liePlace int
	if x == 0 && y == 0 {
		liePlace = 0
	} else if y == 0 {
		liePlace = 1
	} else if x == 0 {
		liePlace = 2
	} else {
		liePlace = 3
	}
	return it.checker.CompareInt(liePlace)
}

// { "no": 22, "dat": "2", "ans": "" }
func (it *ifTask) if22() bool {
	var x = it.scanner.NextFloat64()
	var y = it.scanner.NextFloat64()
	var quarter int
	if x > 0 {
		if y > 0 {
			quarter = 1
		} else {
			quarter = 4
		}
	} else if y > 0 {
		quarter = 2
	} else {
		quarter = 3
	}
	return it.checker.CompareInt(quarter)
}

// { "no": 23, "dat": "", "ans": "" }
func (it *ifTask) if23() bool {
	var x1 = it.scanner.NextInt()
	var y1 = it.scanner.NextInt()
	var x2 = it.scanner.NextInt()
	var y2 = it.scanner.NextInt()
	var x3 = it.scanner.NextInt()
	var y3 = it.scanner.NextInt()
	var x4, y4 int
	if x1 == x2 {
		x4 = x3
	} else if x1 == x3 {
		x4 = x2
	} else {
		x4 = x1
	}
	if y1 == y2 {
		y4 = y3
	} else if y1 == y3 {
		y4 = y2
	} else {
		y4 = y1
	}
	return it.checker.CompareInt(x4, y4)
}

// { "no": 24, "dat": "2", "ans": "2" }
func (it *ifTask) if24() bool {
	var x = it.scanner.NextFloat64()
	var fx float64
	if x > 0 {
		fx = 2 * math.Sin(x)
	} else {
		fx = 6 - x
	}
	return it.checker.CompareFloat64(2, fx)
}

// { "no": 25, "dat": "", "ans": "" }
func (it *ifTask) if25() bool {
	var x = it.scanner.NextInt()
	var fx int
	if x < -2 || x > 2 {
		fx = 2 * x
	} else {
		fx = -3 * x
	}
	return it.checker.CompareInt(fx)
}

// { "no": 26, "dat": "2", "ans": "2" }
func (it *ifTask) if26() bool {
	var x = it.scanner.NextFloat64()
	var fx float64
	if x <= 0 {
		fx = -x
	} else if x >= 2 {
		fx = 4
	} else {
		fx = x * x
	}
	return it.checker.CompareFloat64(2, fx)
}

// { "no": 27, "dat": "2", "ans": "" }
func (it *ifTask) if27() bool {
	var x = it.scanner.NextFloat64()
	var fx int
	if x < 0 {
		fx = 0
	} else if int(x)%2 == 0 {
		fx = 1
	} else {
		fx = -1
	}
	return it.checker.CompareInt(fx)
}

// { "no": 28, "dat": "", "ans": "" }
func (it *ifTask) if28() bool {
	var year = it.scanner.NextInt()
	var daysCount int
	if year%400 == 0 {
		daysCount = 366
	} else if year%100 == 0 {
		daysCount = 365
	} else if year%4 == 0 {
		daysCount = 366
	} else {
		daysCount = 365
	}
	return it.checker.CompareInt(daysCount)
}

// { "no": 29, "dat": "", "ans": "" }
func (it *ifTask) if29() bool {
	var number = it.scanner.NextInt()
	var description string
	if number > 0 {
		description = "positive"
	} else if number < 0 {
		description = "negative"
	} else {
		description = "zero"
	}
	if number != 0 {
		if number%2 == 0 {
			description += " even"
		} else {
			description += " odd"
		}
	}
	description += " number"
	return it.checker.CompareLine(description)
}

// { "no": 30, "dat": "", "ans": "" }
func (it *ifTask) if30() bool {
	var number = it.scanner.NextInt()
	var description string
	if number < 10 {
		description = "one"
	} else if number < 100 {
		description = "two"
	} else {
		description = "three"
	}
	description += "-digit"
	if number%2 == 0 {
		description += " even"
	} else {
		description += " odd"
	}
	description += " number"
	return it.checker.CompareLine(description)
}
