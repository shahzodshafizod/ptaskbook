package task

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type whileTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewWhileTask(pathPrefix string) Task {
	return &whileTask{
		dataPath: pathPrefix + "/07while/While%03d/%03d/%03d",
		name:     "While",
		count:    30,
	}
}

func (wt *whileTask) GetCount() int { return wt.count }

func (wt *whileTask) GetName() string { return wt.name }

func (wt *whileTask) SetData(taskNo int, testNo int) error {
	var filePath = fmt.Sprintf(wt.dataPath, taskNo, testNo, testNo)
	scanner, err := scanner.New(filePath + ".dat")
	if err != nil {
		return errors.Wrap(err, "scanner.New")
	}
	checker, err := checker.New(filePath + ".ans")
	if err != nil {
		return errors.Wrap(err, "checker.New")
	}
	wt.scanner = scanner
	wt.checker = checker
	return nil
}

func (wt *whileTask) ScannerEOF() bool { return wt.scanner.EOF() }

func (wt *whileTask) CheckerEOF() bool { return wt.checker.EOF() }

func (wt *whileTask) CleanData() {
	wt.scanner.RemoveFiles()
	wt.checker.RemoveFiles()
}

func (wt *whileTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return wt.while1
	case 2:
		return wt.while2
	case 3:
		return wt.while3
	case 4:
		return wt.while4
	case 5:
		return wt.while5
	case 6:
		return wt.while6
	case 7:
		return wt.while7
	case 8:
		return wt.while8
	case 9:
		return wt.while9
	case 10:
		return wt.while10
	case 11:
		return wt.while11
	case 12:
		return wt.while12
	case 13:
		return wt.while13
	case 14:
		return wt.while14
	case 15:
		return wt.while15
	case 16:
		return wt.while16
	case 17:
		return wt.while17
	case 18:
		return wt.while18
	case 19:
		return wt.while19
	case 20:
		return wt.while20
	case 21:
		return wt.while21
	case 22:
		return wt.while22
	case 23:
		return wt.while23
	case 24:
		return wt.while24
	case 25:
		return wt.while25
	case 26:
		return wt.while26
	case 27:
		return wt.while27
	case 28:
		return wt.while28
	case 29:
		return wt.while29
	case 30:
		return wt.while30
	default:
		return nil
	}
}

// { "no": 1, "dat": "2", "ans": "2" }
func (wt *whileTask) while1() bool {
	var a = wt.scanner.NextFloat64()
	var b = wt.scanner.NextFloat64()
	for a >= b {
		a -= b
	}
	var unusedPart = a
	return wt.checker.CompareFloat64(2, unusedPart)
}

// { "no": 2, "dat": "2", "ans": "" }
func (wt *whileTask) while2() bool {
	var a = wt.scanner.NextFloat64()
	var b = wt.scanner.NextFloat64()
	var segments = 0
	for a >= b {
		a -= b
		segments++
	}
	return wt.checker.CompareInt(segments)
}

// { "no": 3, "dat": "", "ans": "" }
func (wt *whileTask) while3() bool {
	var n = wt.scanner.NextInt()
	var k = wt.scanner.NextInt()
	var segments = 0
	for n >= k {
		n -= k
		segments++
	}
	var unusedPart = n
	return wt.checker.CompareInt(segments, unusedPart)
}

// { "no": 4, "dat": "", "ans": "" }
func (wt *whileTask) while4() bool {
	var n = wt.scanner.NextInt()
	var degree3 = 1
	for degree3 < n {
		degree3 *= 3
	}
	var isDegree3 = degree3 == n
	return wt.checker.CompareBool(isDegree3)
}

// { "no": 5, "dat": "", "ans": "" }
func (wt *whileTask) while5() bool {
	var n = wt.scanner.NextInt()
	var degree2 = 1
	var k = 0
	for degree2 < n {
		degree2 *= 2
		k++
	}
	return wt.checker.CompareInt(k)
}

// { "no": 6, "dat": "", "ans": "2" }
func (wt *whileTask) while6() bool {
	var n = wt.scanner.NextInt()
	var doubleFactorial = 1.0
	for n > 0 {
		doubleFactorial *= float64(n)
		n -= 2
	}
	return wt.checker.CompareFloat64(2, doubleFactorial)
}

// { "no": 7, "dat": "", "ans": "" }
func (wt *whileTask) while7() bool {
	var n = wt.scanner.NextInt()
	var k = 1
	for k*k <= n {
		k++
	}
	return wt.checker.CompareInt(k)
}

// { "no": 8, "dat": "", "ans": "" }
func (wt *whileTask) while8() bool {
	var n = wt.scanner.NextInt()
	var k = 1
	for k*k <= n {
		k++
	}
	k--
	return wt.checker.CompareInt(k)
}

// { "no": 9, "dat": "", "ans": "" }
func (wt *whileTask) while9() bool {
	var n = wt.scanner.NextInt()
	var k = 0
	var degree3 = 1
	for degree3 <= n {
		degree3 *= 3
		k++
	}
	return wt.checker.CompareInt(k)
}

// { "no": 10, "dat": "", "ans": "" }
func (wt *whileTask) while10() bool {
	var n = wt.scanner.NextInt()
	var k = 0
	var degree3 = 1
	for degree3 < n {
		degree3 *= 3
		k++
	}
	k--
	return wt.checker.CompareInt(k)
}

// { "no": 11, "dat": "", "ans": "" }
func (wt *whileTask) while11() bool {
	var n = wt.scanner.NextInt()
	var sum = 0
	var k = 0
	for sum < n {
		k++
		sum += k
	}
	return wt.checker.CompareInt(k, sum)
}

// { "no": 12, "dat": "", "ans": "" }
func (wt *whileTask) while12() bool {
	var n = wt.scanner.NextInt()
	var sum = 0
	var k = 0
	for sum <= n {
		k++
		sum += k
	}
	sum -= k
	k--
	return wt.checker.CompareInt(k, sum)
}

// { "no": 13, "dat": "5", "ans": "5" }
func (wt *whileTask) while13() bool {
	var a = wt.scanner.NextFloat64()
	var sum = 0.0
	var k = 0
	for sum <= a {
		k++
		sum += 1.0 / float64(k)
	}
	if !wt.checker.CompareInt(k) {
		return false
	}
	return wt.checker.CompareFloat64(5, sum)
}

// { "no": 14, "dat": "5", "ans": "5" }
func (wt *whileTask) while14() bool {
	var a = wt.scanner.NextFloat64()
	var sum = 0.0
	var k = 0
	for sum < a {
		k++
		sum += 1.0 / float64(k)
	}
	sum -= 1.0 / float64(k)
	k--
	if !wt.checker.CompareInt(k) {
		return false
	}
	return wt.checker.CompareFloat64(5, sum)
}

// { "no": 15, "dat": "2", "ans": "2" }
func (wt *whileTask) while15() bool {
	var amount = 1000.0
	var percent = wt.scanner.NextFloat64()
	var years = 0
	for amount <= 1100 {
		years++
		amount += amount * percent / 100
	}
	if !wt.checker.CompareInt(years) {
		return false
	}
	return wt.checker.CompareFloat64(2, amount)
}

// { "no": 16, "dat": "2", "ans": "3" }
func (wt *whileTask) while16() bool {
	var run = 10.0
	var percent = wt.scanner.NextFloat64()
	var days = 0
	var totalRun = 0.0
	for totalRun <= 200 {
		totalRun += run
		days++
		run += run * percent / 100
	}
	if !wt.checker.CompareInt(days) {
		return false
	}
	return wt.checker.CompareFloat64(3, totalRun)
}

// { "no": 17, "dat": "", "ans": "" }
func (wt *whileTask) while17() bool {
	var n = wt.scanner.NextInt()
	var digit int
	for n > 0 {
		digit = n % 10
		if !wt.checker.CompareInt(digit) {
			return false
		}
		n /= 10
	}
	return true
}

// { "no": 18, "dat": "", "ans": "" }
func (wt *whileTask) while18() bool {
	var n = wt.scanner.NextInt()
	var sum, count = 0, 0
	for n > 0 {
		count++
		sum += n % 10
		n /= 10
	}
	return wt.checker.CompareInt(count, sum)
}

// { "no": 19, "dat": "", "ans": "" }
func (wt *whileTask) while19() bool {
	var n = wt.scanner.NextInt()
	var chappa = 0
	for n > 0 {
		chappa = chappa*10 + n%10
		n /= 10
	}
	n = chappa
	return wt.checker.CompareInt(n)
}

// { "no": 20, "dat": "", "ans": "" }
func (wt *whileTask) while20() bool {
	var n = wt.scanner.NextInt()
	var hasTwo = false
	for n > 0 && !hasTwo {
		if n%10 == 2 {
			hasTwo = true
		}
		n /= 10
	}
	return wt.checker.CompareBool(hasTwo)
}

// { "no": 21, "dat": "", "ans": "" }
func (wt *whileTask) while21() bool {
	var n = wt.scanner.NextInt()
	var hasOdd = false
	for n > 0 && !hasOdd {
		if n%10%2 != 0 {
			hasOdd = true
		}
		n /= 10
	}
	return wt.checker.CompareBool(hasOdd)
}

// { "no": 22, "dat": "", "ans": "" }
func (wt *whileTask) while22() bool {
	var n = wt.scanner.NextInt()
	var divisor = int(math.Sqrt(float64(n)))
	var isPrime = true
	for divisor > 1 && isPrime {
		if n%divisor == 0 {
			isPrime = false
		}
		divisor--
	}
	return wt.checker.CompareBool(isPrime)
}

// { "no": 23, "dat": "", "ans": "" }
func (wt *whileTask) while23() bool {
	var a = wt.scanner.NextInt()
	var b = wt.scanner.NextInt()
	for b != 0 {
		a, b = b, a%b
	}
	return wt.checker.CompareInt(a)
}

// { "no": 24, "dat": "", "ans": "" }
func (wt *whileTask) while24() bool {
	var n = wt.scanner.NextInt()
	var prevF, nextF = 1, 1
	for nextF < n {
		prevF, nextF = nextF, prevF+nextF
	}
	var isFibonacci = nextF == n
	return wt.checker.CompareBool(isFibonacci)
}

// { "no": 25, "dat": "", "ans": "" }
func (wt *whileTask) while25() bool {
	var n = wt.scanner.NextInt()
	var prevF, nextF = 1, 1
	for nextF <= n {
		prevF, nextF = nextF, prevF+nextF
	}
	return wt.checker.CompareInt(nextF)
}

// { "no": 26, "dat": "", "ans": "" }
func (wt *whileTask) while26() bool {
	var n = wt.scanner.NextInt()
	var prevF, nextF = 1, 1
	for nextF < n {
		prevF, nextF = nextF, prevF+nextF
	}
	nextF += prevF
	return wt.checker.CompareInt(prevF, nextF)
}

// { "no": 27, "dat": "", "ans": "" }
func (wt *whileTask) while27() bool {
	var n = wt.scanner.NextInt()
	var prevF, nextF = 1, 1
	var k = 2
	for nextF < n {
		prevF, nextF = nextF, prevF+nextF
		k++
	}
	return wt.checker.CompareInt(k)
}

// { "no": 28, "dat": "6", "ans": "6" }
func (wt *whileTask) while28() bool {
	var eps = wt.scanner.NextFloat64()
	var prevA = 2.0
	var currA float64
	var k = 1
	for {
		currA = 2 + 1/prevA
		k++
		if math.Abs(currA-prevA) < eps {
			break
		}
		prevA = currA
	}
	if !wt.checker.CompareInt(k) {
		return false
	}
	return wt.checker.CompareFloat64(6, prevA, currA)
}

// { "no": 29, "dat": "6", "ans": "6" }
func (wt *whileTask) while29() bool {
	var eps = wt.scanner.NextFloat64()
	var prevA, currA = 1.0, 2.0
	var nextA float64
	var k = 1
	for {
		nextA = (prevA + 2*currA) / 3
		k++
		if math.Abs(prevA-currA) < eps {
			break
		}
		prevA, currA = currA, nextA
	}
	if !wt.checker.CompareInt(k) {
		return false
	}
	return wt.checker.CompareFloat64(6, prevA, currA)
}

// { "no": 30, "dat": "2", "ans": "" }
func (wt *whileTask) while30() bool {
	var a = wt.scanner.NextFloat64()
	var b = wt.scanner.NextFloat64()
	var c = wt.scanner.NextFloat64()
	var cInA = 0
	for a >= c {
		cInA++
		a -= c
	}
	var squaresCount = 0
	for b >= c {
		squaresCount += cInA
		b -= c
	}
	return wt.checker.CompareInt(squaresCount)
}
