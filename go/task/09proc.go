package task

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/proc"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type procTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewProcTask(pathPrefix string) Task {
	return &procTask{
		dataPath: pathPrefix + "/09proc/Proc%03d/%03d/%03d",
		name:     "Proc",
		count:    60,
	}
}

func (pt *procTask) GetCount() int { return pt.count }

func (pt *procTask) GetName() string { return pt.name }

func (pt *procTask) SetData(taskNo int, testNo int) error {
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

func (pt *procTask) ScannerEOF() bool { return pt.scanner.EOF() }

func (pt *procTask) CheckerEOF() bool { return pt.checker.EOF() }

func (pt *procTask) CleanData() {
	pt.scanner.RemoveFiles()
	pt.checker.RemoveFiles()
}

func (pt *procTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return pt.proc1
	case 2:
		return pt.proc2
	case 3:
		return pt.proc3
	case 4:
		return pt.proc4
	case 5:
		return pt.proc5
	case 6:
		return pt.proc6
	case 7:
		return pt.proc7
	case 8:
		return pt.proc8
	case 9:
		return pt.proc9
	case 10:
		return pt.proc10
	case 11:
		return pt.proc11
	case 12:
		return pt.proc12
	case 13:
		return pt.proc13
	case 14:
		return pt.proc14
	case 15:
		return pt.proc15
	case 16:
		return pt.proc16
	case 17:
		return pt.proc17
	case 18:
		return pt.proc18
	case 19:
		return pt.proc19
	case 20:
		return pt.proc20
	case 21:
		return pt.proc21
	case 22:
		return pt.proc22
	case 23:
		return pt.proc23
	case 24:
		return pt.proc24
	case 25:
		return pt.proc25
	case 26:
		return pt.proc26
	case 27:
		return pt.proc27
	case 28:
		return pt.proc28
	case 29:
		return pt.proc29
	case 30:
		return pt.proc30
	case 31:
		return pt.proc31
	case 32:
		return pt.proc32
	case 33:
		return pt.proc33
	case 34:
		return pt.proc34
	case 35:
		return pt.proc35
	case 36:
		return pt.proc36
	case 37:
		return pt.proc37
	case 38:
		return pt.proc38
	case 39:
		return pt.proc39
	case 40:
		return pt.proc40
	case 41:
		return pt.proc41
	case 42:
		return pt.proc42
	case 43:
		return pt.proc43
	case 44:
		return pt.proc44
	case 45:
		return pt.proc45
	case 46:
		return pt.proc46
	case 47:
		return pt.proc47
	case 48:
		return pt.proc48
	case 49:
		return pt.proc49
	case 50:
		return pt.proc50
	case 51:
		return pt.proc51
	case 52:
		return pt.proc52
	case 53:
		return pt.proc53
	case 54:
		return pt.proc54
	case 55:
		return pt.proc55
	case 56:
		return pt.proc56
	case 57:
		return pt.proc57
	case 58:
		return pt.proc58
	case 59:
		return pt.proc59
	case 60:
		return pt.proc60
	default:
		return nil
	}
}

// { "no": 1, "dat": "2", "ans": "2" }
func (pt *procTask) proc1() bool {
	var number, degree float64
	for index := 0; index < 5; index++ {
		number = pt.scanner.NextFloat64()
		proc.PowerA3(number, &degree)
		if !pt.checker.CompareFloat64(2, degree) {
			return false
		}
	}
	return true
}

// { "no": 2, "dat": "2", "ans": "2" }
func (pt *procTask) proc2() bool {
	var number, degree2, degree3, degree4 float64
	for index := 0; index < 5; index++ {
		number = pt.scanner.NextFloat64()
		proc.PowerA234(number, &degree2, &degree3, &degree4)
		if !pt.checker.CompareFloat64(2, degree2, degree3, degree4) {
			return false
		}
	}
	return true
}

// { "no": 3, "dat": "2", "ans": "2" }
func (pt *procTask) proc3() bool {
	var a = pt.scanner.NextFloat64()
	var b, aMean, gMean float64
	for index := 0; index < 3; index++ {
		b = pt.scanner.NextFloat64()
		proc.Mean(a, b, &aMean, &gMean)
		if !pt.checker.CompareFloat64(2, aMean, gMean) {
			return false
		}
	}
	return true
}

// { "no": 4, "dat": "2", "ans": "2" }
func (pt *procTask) proc4() bool {
	var a, p, s float64
	for index := 0; index < 3; index++ {
		a = pt.scanner.NextFloat64()
		proc.TrianglePS(a, &p, &s)
		if !pt.checker.CompareFloat64(2, p, s) {
			return false
		}
	}
	return true
}

// { "no": 5, "dat": "2", "ans": "2" }
func (pt *procTask) proc5() bool {
	var x1, y1, x2, y2, p, s float64
	for index := 0; index < 3; index++ {
		x1 = pt.scanner.NextFloat64()
		y1 = pt.scanner.NextFloat64()
		x2 = pt.scanner.NextFloat64()
		y2 = pt.scanner.NextFloat64()
		proc.RectPS(x1, y1, x2, y2, &p, &s)
		if !pt.checker.CompareFloat64(2, p, s) {
			return false
		}
	}
	return true
}

// { "no": 6, "dat": "", "ans": "" }
func (pt *procTask) proc6() bool {
	var number, count, sum int
	for index := 0; index < 5; index++ {
		number = pt.scanner.NextInt()
		proc.DigitCountSum(number, &count, &sum)
		if !pt.checker.CompareInt(count, sum) {
			return false
		}
	}
	return true
}

// { "no": 7, "dat": "", "ans": "" }
func (pt *procTask) proc7() bool {
	var number int
	for index := 0; index < 5; index++ {
		number = pt.scanner.NextInt()
		proc.InvDigits(&number)
		if !pt.checker.CompareInt(number) {
			return false
		}
	}
	return true
}

// { "no": 8, "dat": "", "ans": "" }
func (pt *procTask) proc8() bool {
	var k = pt.scanner.NextInt()
	var d int
	for index := 0; index < 2; index++ {
		d = pt.scanner.NextInt()
		proc.AddRightDigit(d, &k)
		if !pt.checker.CompareInt(k) {
			return false
		}
	}
	return true
}

// { "no": 9, "dat": "", "ans": "" }
func (pt *procTask) proc9() bool {
	var k = pt.scanner.NextInt()
	var d int
	for index := 0; index < 2; index++ {
		d = pt.scanner.NextInt()
		proc.AddLeftDigit(d, &k)
		if !pt.checker.CompareInt(k) {
			return false
		}
	}
	return true
}

// { "no": 10, "dat": "2", "ans": "2" }
func (pt *procTask) proc10() bool {
	var a = pt.scanner.NextFloat64()
	var b = pt.scanner.NextFloat64()
	var c = pt.scanner.NextFloat64()
	var d = pt.scanner.NextFloat64()
	proc.Swap(&a, &b)
	proc.Swap(&c, &d)
	proc.Swap(&b, &c)
	return pt.checker.CompareFloat64(2, a, b, c, d)
}

// { "no": 11, "dat": "2", "ans": "2" }
func (pt *procTask) proc11() bool {
	var a = pt.scanner.NextFloat64()
	var b = pt.scanner.NextFloat64()
	var c = pt.scanner.NextFloat64()
	var d = pt.scanner.NextFloat64()
	proc.Minmax(&a, &b)
	proc.Minmax(&c, &d)
	proc.Minmax(&a, &c)
	proc.Minmax(&b, &d)
	return pt.checker.CompareFloat64(2, a, d)
}

// { "no": 12, "dat": "2", "ans": "2" }
func (pt *procTask) proc12() bool {
	var a, b, c float64
	for index := 0; index < 2; index++ {
		a = pt.scanner.NextFloat64()
		b = pt.scanner.NextFloat64()
		c = pt.scanner.NextFloat64()
		proc.SortInc3(&a, &b, &c)
		if !pt.checker.CompareFloat64(2, a, b, c) {
			return false
		}
	}
	return true
}

// { "no": 13, "dat": "2", "ans": "2" }
func (pt *procTask) proc13() bool {
	var a, b, c float64
	for index := 0; index < 2; index++ {
		a = pt.scanner.NextFloat64()
		b = pt.scanner.NextFloat64()
		c = pt.scanner.NextFloat64()
		proc.SortDec3(&a, &b, &c)
		if !pt.checker.CompareFloat64(2, a, b, c) {
			return false
		}
	}
	return true
}

// { "no": 14, "dat": "2", "ans": "2" }
func (pt *procTask) proc14() bool {
	var a, b, c float64
	for index := 0; index < 2; index++ {
		a = pt.scanner.NextFloat64()
		b = pt.scanner.NextFloat64()
		c = pt.scanner.NextFloat64()
		proc.ShiftRight3(&a, &b, &c)
		if !pt.checker.CompareFloat64(2, a, b, c) {
			return false
		}
	}
	return true
}

// { "no": 15, "dat": "2", "ans": "2" }
func (pt *procTask) proc15() bool {
	var a, b, c float64
	for index := 0; index < 2; index++ {
		a = pt.scanner.NextFloat64()
		b = pt.scanner.NextFloat64()
		c = pt.scanner.NextFloat64()
		proc.ShiftLeft3(&a, &b, &c)
		if !pt.checker.CompareFloat64(2, a, b, c) {
			return false
		}
	}
	return true
}

// { "no": 16, "dat": "2", "ans": "" }
func (pt *procTask) proc16() bool {
	var a = pt.scanner.NextFloat64()
	var b = pt.scanner.NextFloat64()
	var result = proc.Sign(a) + proc.Sign(b)
	return pt.checker.CompareInt(result)
}

// { "no": 17, "dat": "2", "ans": "" }
func (pt *procTask) proc17() bool {
	var a, b, c float64
	var count int
	for index := 0; index < 3; index++ {
		a = pt.scanner.NextFloat64()
		b = pt.scanner.NextFloat64()
		c = pt.scanner.NextFloat64()
		count = proc.RootCount(a, b, c)
		if !pt.checker.CompareInt(count) {
			return false
		}
	}
	return true
}

// { "no": 18, "dat": "2", "ans": "2" }
func (pt *procTask) proc18() bool {
	var radius, area float64
	for index := 0; index < 3; index++ {
		radius = pt.scanner.NextFloat64()
		area = proc.CircleS(radius)
		if !pt.checker.CompareFloat64(2, area) {
			return false
		}
	}
	return true
}

// { "no": 19, "dat": "2", "ans": "2" }
func (pt *procTask) proc19() bool {
	var outerRadius, innerRadius, area float64
	for index := 0; index < 3; index++ {
		outerRadius = pt.scanner.NextFloat64()
		innerRadius = pt.scanner.NextFloat64()
		area = proc.RingS(outerRadius, innerRadius)
		if !pt.checker.CompareFloat64(2, area) {
			return false
		}
	}
	return true
}

// { "no": 20, "dat": "2", "ans": "2" }
func (pt *procTask) proc20() bool {
	var baseSide, altitude, perimeter float64
	for index := 0; index < 3; index++ {
		baseSide = pt.scanner.NextFloat64()
		altitude = pt.scanner.NextFloat64()
		perimeter = proc.TriangleP(baseSide, altitude)
		if !pt.checker.CompareFloat64(2, perimeter) {
			return false
		}
	}
	return true
}

// { "no": 21, "dat": "", "ans": "" }
func (pt *procTask) proc21() bool {
	var a = pt.scanner.NextInt()
	var b = pt.scanner.NextInt()
	var c = pt.scanner.NextInt()
	var sum = proc.SumRange(a, b)
	if !pt.checker.CompareInt(sum) {
		return false
	}
	sum = proc.SumRange(b, c)
	return pt.checker.CompareInt(sum)
}

// { "no": 22, "dat": "2", "ans": "2" }
func (pt *procTask) proc22() bool {
	var a = pt.scanner.NextFloat64()
	var b = pt.scanner.NextFloat64()
	var n int
	var result float64
	for index := 0; index < 3; index++ {
		n = pt.scanner.NextInt()
		result = proc.Calc(a, b, n)
		if !pt.checker.CompareFloat64(2, result) {
			return false
		}
	}
	return true
}

// { "no": 23, "dat": "2", "ans": "" }
func (pt *procTask) proc23() bool {
	var x, y float64
	var number int
	for index := 0; index < 3; index++ {
		x = pt.scanner.NextFloat64()
		y = pt.scanner.NextFloat64()
		number = proc.Quarter(x, y)
		if !pt.checker.CompareInt(number) {
			return false
		}
	}
	return true
}

// { "no": 24, "dat": "", "ans": "" }
func (pt *procTask) proc24() bool {
	var number int
	var evensCount = 0
	for index := 0; index < 10; index++ {
		number = pt.scanner.NextInt()
		if proc.Even(number) {
			evensCount++
		}
	}
	return pt.checker.CompareInt(evensCount)
}

// { "no": 25, "dat": "", "ans": "" }
func (pt *procTask) proc25() bool {
	var number int
	var squaresCount = 0
	for index := 0; index < 10; index++ {
		number = pt.scanner.NextInt()
		if proc.IsSquare(number) {
			squaresCount++
		}
	}
	return pt.checker.CompareInt(squaresCount)
}

// { "no": 26, "dat": "", "ans": "" }
func (pt *procTask) proc26() bool {
	var number int
	var fifthPowersCount = 0
	for index := 0; index < 10; index++ {
		number = pt.scanner.NextInt()
		if proc.IsPower5(number) {
			fifthPowersCount++
		}
	}
	return pt.checker.CompareInt(fifthPowersCount)
}

// { "no": 27, "dat": "", "ans": "" }
func (pt *procTask) proc27() bool {
	var n = pt.scanner.NextInt()
	var number int
	var powersCount = 0
	for index := 0; index < 10; index++ {
		number = pt.scanner.NextInt()
		if proc.IsPowerN(number, n) {
			powersCount++
		}
	}
	return pt.checker.CompareInt(powersCount)
}

// { "no": 28, "dat": "", "ans": "" }
func (pt *procTask) proc28() bool {
	var number int
	var primesCount = 0
	for index := 0; index < 10; index++ {
		number = pt.scanner.NextInt()
		if proc.IsPrime(number) {
			primesCount++
		}
	}
	return pt.checker.CompareInt(primesCount)
}

// { "no": 29, "dat": "", "ans": "" }
func (pt *procTask) proc29() bool {
	var number, count int
	for index := 0; index < 5; index++ {
		number = pt.scanner.NextInt()
		count = proc.DigitCount(number)
		if !pt.checker.CompareInt(count) {
			return false
		}
	}
	return true
}

// { "no": 30, "dat": "", "ans": "" }
func (pt *procTask) proc30() bool {
	var number, digit int
	for index := 0; index < 5; index++ {
		number = pt.scanner.NextInt()
		for i := 1; i <= 5; i++ {
			digit = proc.DigitN(number, i)
			if !pt.checker.CompareInt(digit) {
				return false
			}
		}
	}
	return true
}

// { "no": 31, "dat": "", "ans": "" }
func (pt *procTask) proc31() bool {
	var number int
	var palindromesCount = 0
	for index := 0; index < 10; index++ {
		number = pt.scanner.NextInt()
		if proc.IsPalindrome(number) {
			palindromesCount++
		}
	}
	return pt.checker.CompareInt(palindromesCount)
}

// { "no": 32, "dat": "2", "ans": "2" }
func (pt *procTask) proc32() bool {
	var degree, radian float64
	for index := 0; index < 5; index++ {
		degree = pt.scanner.NextFloat64()
		radian = proc.DegToRad(degree)
		if !pt.checker.CompareFloat64(2, radian) {
			return false
		}
	}
	return true
}

// { "no": 33, "dat": "2", "ans": "2" }
func (pt *procTask) proc33() bool {
	var degree, radian float64
	for index := 0; index < 5; index++ {
		radian = pt.scanner.NextFloat64()
		degree = proc.RadToDeg(radian)
		if !pt.checker.CompareFloat64(2, degree) {
			return false
		}
	}
	return true
}

// { "no": 34, "dat": "", "ans": "2" }
func (pt *procTask) proc34() bool {
	var number int
	var factorial float64
	for index := 0; index < 5; index++ {
		number = pt.scanner.NextInt()
		factorial = proc.Fact(number)
		if !pt.checker.CompareFloat64(2, factorial) {
			return false
		}
	}
	return true
}

// { "no": 35, "dat": "", "ans": "2" }
func (pt *procTask) proc35() bool {
	var number int
	var factorial float64
	for index := 0; index < 5; index++ {
		number = pt.scanner.NextInt()
		factorial = proc.Fact2(number)
		if !pt.checker.CompareFloat64(2, factorial) {
			return false
		}
	}
	return true
}

// { "no": 36, "dat": "", "ans": "" }
func (pt *procTask) proc36() bool {
	var n, fn int
	for index := 0; index < 5; index++ {
		n = pt.scanner.NextInt()
		fn = proc.Fib(n)
		if !pt.checker.CompareInt(fn) {
			return false
		}
	}
	return true
}

// { "no": 37, "dat": "2", "ans": "2" }
func (pt *procTask) proc37() bool {
	var p = pt.scanner.NextFloat64()
	var a, result float64
	for index := 0; index < 3; index++ {
		a = pt.scanner.NextFloat64()
		result = proc.Power1(a, p)
		if !pt.checker.CompareFloat64(2, result) {
			return false
		}
	}
	return true
}

// { "no": 38, "dat": "2", "ans": "2" }
func (pt *procTask) proc38() bool {
	var a = pt.scanner.NextFloat64()
	var n int
	var result float64
	for index := 0; index < 3; index++ {
		n = pt.scanner.NextInt()
		result = proc.Power2(a, n)
		if !pt.checker.CompareFloat64(2, result) {
			return false
		}
	}
	return true
}

// { "no": 39, "dat": "2", "ans": "2" }
func (pt *procTask) proc39() bool {
	var p = pt.scanner.NextFloat64()
	var a, result float64
	for index := 0; index < 3; index++ {
		a = pt.scanner.NextFloat64()
		result = proc.Power3(a, p)
		if !pt.checker.CompareFloat64(2, result) {
			return false
		}
	}
	return true
}

// { "no": 40, "dat": "7", "ans": "7" }
func (pt *procTask) proc40() bool {
	var x = pt.scanner.NextFloat64()
	var e, result float64
	for index := 0; index < 6; index++ {
		e = pt.scanner.NextFloat64()
		result = proc.Exp1(x, e)
		if !pt.checker.CompareFloat64(7, result) {
			return false
		}
	}
	return true
}

// { "no": 41, "dat": "7", "ans": "7" }
func (pt *procTask) proc41() bool {
	var x = pt.scanner.NextFloat64()
	var e, result float64
	for index := 0; index < 6; index++ {
		e = pt.scanner.NextFloat64()
		result = proc.Sin1(x, e)
		if !pt.checker.CompareFloat64(7, result) {
			return false
		}
	}
	return true
}

// { "no": 42, "dat": "7", "ans": "7" }
func (pt *procTask) proc42() bool {
	var x = pt.scanner.NextFloat64()
	var e, result float64
	for index := 0; index < 6; index++ {
		e = pt.scanner.NextFloat64()
		result = proc.Cos1(x, e)
		if !pt.checker.CompareFloat64(7, result) {
			return false
		}
	}
	return true
}

// { "no": 43, "dat": "7", "ans": "7" }
func (pt *procTask) proc43() bool {
	var x = pt.scanner.NextFloat64()
	var e, result float64
	for index := 0; index < 6; index++ {
		e = pt.scanner.NextFloat64()
		result = proc.Ln1(x, e)
		if !pt.checker.CompareFloat64(7, result) {
			return false
		}
	}
	return true
}

// { "no": 44, "dat": "7", "ans": "7" }
func (pt *procTask) proc44() bool {
	var x = pt.scanner.NextFloat64()
	var e, result float64
	for index := 0; index < 6; index++ {
		e = pt.scanner.NextFloat64()
		result = proc.Atan1(x, e)
		if !pt.checker.CompareFloat64(7, result) {
			return false
		}
	}
	return true
}

// { "no": 45, "dat": "7", "ans": "7" }
func (pt *procTask) proc45() bool {
	var x = pt.scanner.NextFloat64()
	var a = pt.scanner.NextFloat64()
	var e, result float64
	for index := 0; index < 6; index++ {
		e = pt.scanner.NextFloat64()
		result = proc.Power4(x, a, e)
		if !pt.checker.CompareFloat64(7, result) {
			return false
		}
	}
	return true
}

// { "no": 46, "dat": "", "ans": "" }
func (pt *procTask) proc46() bool {
	var a = pt.scanner.NextInt()
	var b, result int
	for index := 0; index < 3; index++ {
		b = pt.scanner.NextInt()
		result = proc.Gcd2(a, b)
		if !pt.checker.CompareInt(result) {
			return false
		}
	}
	return true
}

// { "no": 47, "dat": "", "ans": "" }
func (pt *procTask) proc47() bool {
	var a = pt.scanner.NextInt()
	var b = pt.scanner.NextInt()
	var c, d, p, q int
	for index := 0; index < 3; index++ {
		c, d = pt.scanner.NextInt(), pt.scanner.NextInt()
		proc.Frac1(a*d+c*b, b*d, &p, &q)
		if !pt.checker.CompareInt(p, q) {
			return false
		}
	}
	return true
}

// { "no": 48, "dat": "", "ans": "" }
func (pt *procTask) proc48() bool {
	var a = pt.scanner.NextInt()
	var b, result int
	for index := 0; index < 3; index++ {
		b = pt.scanner.NextInt()
		result = proc.Lcm2(a, b)
		if !pt.checker.CompareInt(result) {
			return false
		}
	}
	return true
}

// { "no": 49, "dat": "", "ans": "" }
func (pt *procTask) proc49() bool {
	var a = pt.scanner.NextInt()
	var b = pt.scanner.NextInt()
	var c = pt.scanner.NextInt()
	var d = pt.scanner.NextInt()
	var greatestCommonDivisorABC = proc.Gcd3(a, b, c)
	var greatestCommonDivisorACD = proc.Gcd3(a, c, d)
	var greatestCommonDivisorBCD = proc.Gcd3(b, c, d)
	return pt.checker.CompareInt(
		greatestCommonDivisorABC,
		greatestCommonDivisorACD,
		greatestCommonDivisorBCD,
	)
}

// { "no": 50, "dat": "", "ans": "" }
func (pt *procTask) proc50() bool {
	var t, hours, minutes, seconds int
	for index := 0; index < 5; index++ {
		t = pt.scanner.NextInt()
		proc.TimeToHMS(t, &hours, &minutes, &seconds)
		if !pt.checker.CompareInt(hours, minutes, seconds) {
			return false
		}
	}
	return true
}

// { "no": 51, "dat": "", "ans": "" }
func (pt *procTask) proc51() bool {
	var hours = pt.scanner.NextInt()
	var minutes = pt.scanner.NextInt()
	var seconds = pt.scanner.NextInt()
	var t = pt.scanner.NextInt()
	proc.IncTime(&hours, &minutes, &seconds, t)
	return pt.checker.CompareInt(hours, minutes, seconds)
}

// { "no": 52, "dat": "", "ans": "" }
func (pt *procTask) proc52() bool {
	var year int
	var isLeap bool
	for index := 0; index < 5; index++ {
		year = pt.scanner.NextInt()
		isLeap = proc.IsLeapYear(year)
		if !pt.checker.CompareBool(isLeap) {
			return false
		}
	}
	return true
}

// { "no": 53, "dat": "", "ans": "" }
func (pt *procTask) proc53() bool {
	var year = pt.scanner.NextInt()
	var monthNo, daysCount int
	for index := 0; index < 3; index++ {
		monthNo = pt.scanner.NextInt()
		daysCount = proc.MonthDays(monthNo, year)
		if !pt.checker.CompareInt(daysCount) {
			return false
		}
	}
	return true
}

// { "no": 54, "dat": "", "ans": "" }
func (pt *procTask) proc54() bool {
	var day, month, year int
	for index := 0; index < 3; index++ {
		day = pt.scanner.NextInt()
		month = pt.scanner.NextInt()
		year = pt.scanner.NextInt()
		proc.PrevDate(&day, &month, &year)
		if !pt.checker.CompareInt(day, month, year) {
			return false
		}
	}
	return true
}

// { "no": 55, "dat": "", "ans": "" }
func (pt *procTask) proc55() bool {
	var day, month, year int
	for index := 0; index < 3; index++ {
		day = pt.scanner.NextInt()
		month = pt.scanner.NextInt()
		year = pt.scanner.NextInt()
		proc.NextDate(&day, &month, &year)
		if !pt.checker.CompareInt(day, month, year) {
			return false
		}
	}
	return true
}

// { "no": 56, "dat": "2", "ans": "2" }
func (pt *procTask) proc56() bool {
	var xA = pt.scanner.NextFloat64()
	var yA = pt.scanner.NextFloat64()
	var xB, yB, result float64
	for index := 0; index < 3; index++ {
		xB = pt.scanner.NextFloat64()
		yB = pt.scanner.NextFloat64()
		result = proc.Leng(xA, yA, xB, yB)
		if !pt.checker.CompareFloat64(2, result) {
			return false
		}
	}
	return true
}

// { "no": 57, "dat": "2", "ans": "2" }
func (pt *procTask) proc57() bool {
	var xA = pt.scanner.NextFloat64()
	var yA = pt.scanner.NextFloat64()
	var xB = pt.scanner.NextFloat64()
	var yB = pt.scanner.NextFloat64()
	var xC = pt.scanner.NextFloat64()
	var yC = pt.scanner.NextFloat64()
	var xD = pt.scanner.NextFloat64()
	var yD = pt.scanner.NextFloat64()
	var perimeterABC = proc.Perim(xA, yA, xB, yB, xC, yC)
	var perimeterABD = proc.Perim(xA, yA, xB, yB, xD, yD)
	var perimeterACD = proc.Perim(xA, yA, xC, yC, xD, yD)
	return pt.checker.CompareFloat64(2, perimeterABC, perimeterABD, perimeterACD)
}

// { "no": 58, "dat": "2", "ans": "2" }
func (pt *procTask) proc58() bool {
	var xA = pt.scanner.NextFloat64()
	var yA = pt.scanner.NextFloat64()
	var xB = pt.scanner.NextFloat64()
	var yB = pt.scanner.NextFloat64()
	var xC = pt.scanner.NextFloat64()
	var yC = pt.scanner.NextFloat64()
	var xD = pt.scanner.NextFloat64()
	var yD = pt.scanner.NextFloat64()
	var areaABC = proc.Area(xA, yA, xB, yB, xC, yC)
	var areaABD = proc.Area(xA, yA, xB, yB, xD, yD)
	var areaACD = proc.Area(xA, yA, xC, yC, xD, yD)
	return pt.checker.CompareFloat64(2, areaABC, areaABD, areaACD)
}

// { "no": 59, "dat": "2", "ans": "2" }
func (pt *procTask) proc59() bool {
	var xP = pt.scanner.NextFloat64()
	var yP = pt.scanner.NextFloat64()
	var xA = pt.scanner.NextFloat64()
	var yA = pt.scanner.NextFloat64()
	var xB = pt.scanner.NextFloat64()
	var yB = pt.scanner.NextFloat64()
	var xC = pt.scanner.NextFloat64()
	var yC = pt.scanner.NextFloat64()
	var pab = proc.Dist(xP, yP, xA, yA, xB, yB)
	var pac = proc.Dist(xP, yP, xA, yA, xC, yC)
	var pbc = proc.Dist(xP, yP, xB, yB, xC, yC)
	return pt.checker.CompareFloat64(2, pab, pac, pbc)
}

// { "no": 60, "dat": "2", "ans": "2" }
func (pt *procTask) proc60() bool {
	var xA = pt.scanner.NextFloat64()
	var yA = pt.scanner.NextFloat64()
	var xB = pt.scanner.NextFloat64()
	var yB = pt.scanner.NextFloat64()
	var xC = pt.scanner.NextFloat64()
	var yC = pt.scanner.NextFloat64()
	var xD = pt.scanner.NextFloat64()
	var yD = pt.scanner.NextFloat64()
	var hA, hB, hC float64
	proc.Alts(xA, yA, xB, yB, xC, yC, &hA, &hB, &hC)
	if !pt.checker.CompareFloat64(2, hA, hB, hC) {
		return false
	}
	proc.Alts(xA, yA, xB, yB, xD, yD, &hA, &hB, &hC)
	if !pt.checker.CompareFloat64(2, hA, hB, hC) {
		return false
	}
	proc.Alts(xA, yA, xC, yC, xD, yD, &hA, &hB, &hC)
	return pt.checker.CompareFloat64(2, hA, hB, hC)
}
