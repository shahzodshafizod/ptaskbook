package task

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type forTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewForTask(pathPrefix string) Task {
	return &forTask{
		dataPath: pathPrefix + "/06for/For%03d/%03d/%03d",
		name:     "For",
		count:    40,
	}
}

func (ft *forTask) GetCount() int { return ft.count }

func (ft *forTask) GetName() string { return ft.name }

func (ft *forTask) SetData(taskNo int, testNo int) error {
	var filePath = fmt.Sprintf(ft.dataPath, taskNo, testNo, testNo)
	scanner, err := scanner.New(filePath + ".dat")
	if err != nil {
		return errors.Wrap(err, "scanner.New")
	}
	checker, err := checker.New(filePath + ".ans")
	if err != nil {
		return errors.Wrap(err, "checker.New")
	}
	ft.scanner = scanner
	ft.checker = checker
	return nil
}

func (ft *forTask) ScannerEOF() bool { return ft.scanner.EOF() }

func (ft *forTask) CheckerEOF() bool { return ft.checker.EOF() }

func (ft *forTask) CleanData() {
	ft.scanner.RemoveFiles()
	ft.checker.RemoveFiles()
}

func (ft *forTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return ft.for1
	case 2:
		return ft.for2
	case 3:
		return ft.for3
	case 4:
		return ft.for4
	case 5:
		return ft.for5
	case 6:
		return ft.for6
	case 7:
		return ft.for7
	case 8:
		return ft.for8
	case 9:
		return ft.for9
	case 10:
		return ft.for10
	case 11:
		return ft.for11
	case 12:
		return ft.for12
	case 13:
		return ft.for13
	case 14:
		return ft.for14
	case 15:
		return ft.for15
	case 16:
		return ft.for16
	case 17:
		return ft.for17
	case 18:
		return ft.for18
	case 19:
		return ft.for19
	case 20:
		return ft.for20
	case 21:
		return ft.for21
	case 22:
		return ft.for22
	case 23:
		return ft.for23
	case 24:
		return ft.for24
	case 25:
		return ft.for25
	case 26:
		return ft.for26
	case 27:
		return ft.for27
	case 28:
		return ft.for28
	case 29:
		return ft.for29
	case 30:
		return ft.for30
	case 31:
		return ft.for31
	case 32:
		return ft.for32
	case 33:
		return ft.for33
	case 34:
		return ft.for34
	case 35:
		return ft.for35
	case 36:
		return ft.for36
	case 37:
		return ft.for37
	case 38:
		return ft.for38
	case 39:
		return ft.for39
	case 40:
		return ft.for40
	default:
		return nil
	}
}

// { "no": 1, "dat": "", "ans": "" }
func (ft *forTask) for1() bool {
	var k, n = ft.scanner.NextInt(), ft.scanner.NextInt()
	for index := 0; index < n; index++ {
		if !ft.checker.CompareInt(k) {
			return false
		}
	}
	return true
}

// { "no": 2, "dat": "", "ans": "" }
func (ft *forTask) for2() bool {
	var a = ft.scanner.NextInt()
	var b = ft.scanner.NextInt()
	var n int
	for index := a; index <= b; index++ {
		if !ft.checker.CompareInt(index) {
			return false
		}
		n++
	}
	return ft.checker.CompareInt(n)
}

// { "no": 3, "dat": "", "ans": "" }
func (ft *forTask) for3() bool {
	var a = ft.scanner.NextInt()
	var b = ft.scanner.NextInt()
	var n int
	for index := b - 1; index > a; index-- {
		if !ft.checker.CompareInt(index) {
			return false
		}
		n++
	}
	return ft.checker.CompareInt(n)
}

// { "no": 4, "dat": "2", "ans": "2" }
func (ft *forTask) for4() bool {
	var kg = ft.scanner.NextFloat64()
	var price float64
	for index := 1.0; index <= 10; index++ {
		price = index * kg
		if !ft.checker.CompareFloat64(2, price) {
			return false
		}
	}
	return true
}

// { "no": 5, "dat": "2", "ans": "2" }
func (ft *forTask) for5() bool {
	var kg = ft.scanner.NextFloat64()
	var price float64
	for index := 0.1; index <= 1; index += 0.1 {
		price = index * kg
		if !ft.checker.CompareFloat64(2, price) {
			return false
		}
	}
	return true
}

// { "no": 6, "dat": "2", "ans": "2" }
func (ft *forTask) for6() bool {
	var kg = ft.scanner.NextFloat64()
	var price float64
	for index := 1.2; index <= 2; index += 0.2 {
		price = index * kg
		if !ft.checker.CompareFloat64(2, price) {
			return false
		}
	}
	return true
}

// { "no": 7, "dat": "", "ans": "" }
func (ft *forTask) for7() bool {
	var a = ft.scanner.NextInt()
	var b = ft.scanner.NextInt()
	var sum = 0
	for index := a; index <= b; index++ {
		sum += index
	}
	return ft.checker.CompareInt(sum)
}

// { "no": 8, "dat": "", "ans": "" }
func (ft *forTask) for8() bool {
	var a = ft.scanner.NextInt()
	var b = ft.scanner.NextInt()
	var pro = 1
	for index := a; index <= b; index++ {
		pro *= index
	}
	return ft.checker.CompareInt(pro)
}

// { "no": 9, "dat": "", "ans": "" }
func (ft *forTask) for9() bool {
	var a = ft.scanner.NextInt()
	var b = ft.scanner.NextInt()
	var sum = 0
	for index := a; index <= b; index++ {
		sum += index * index
	}
	return ft.checker.CompareInt(sum)
}

// { "no": 10, "dat": "", "ans": "6" }
func (ft *forTask) for10() bool {
	var n = ft.scanner.NextInt()
	var sum = 0.0
	for index := 1; index <= n; index++ {
		sum += 1.0 / float64(index)
	}
	return ft.checker.CompareFloat64(6, sum)
}

// { "no": 11, "dat": "", "ans": "" }
func (ft *forTask) for11() bool {
	var n = ft.scanner.NextInt()
	var sum = 0
	var number = n
	for index := 0; index <= n; index++ {
		sum += number * number
		number++
	}
	return ft.checker.CompareInt(sum)
}

// { "no": 12, "dat": "", "ans": "6" }
func (ft *forTask) for12() bool {
	var n = ft.scanner.NextInt()
	var pro = 1.0
	var number = 1.0
	for index := 0; index < n; index++ {
		number += 0.1
		pro *= number
	}
	return ft.checker.CompareFloat64(6, pro)
}

// { "no": 13, "dat": "", "ans": "2" }
func (ft *forTask) for13() bool {
	var n = ft.scanner.NextInt()
	var result = 0.0
	var number = 1.0
	var sign = 1.0
	for index := 0; index < n; index++ {
		number += 0.1
		result += sign * number
		sign *= -1
	}
	return ft.checker.CompareFloat64(2, result)
}

// { "no": 14, "dat": "", "ans": "" }
func (ft *forTask) for14() bool {
	var n = ft.scanner.NextInt()
	var squaresSum = 0
	for index := 1; index <= n; index++ {
		squaresSum += 2*index - 1
		if !ft.checker.CompareInt(squaresSum) {
			return false
		}
	}
	return true
}

// { "no": 15, "dat": "2", "ans": "2" }
func (ft *forTask) for15() bool {
	var a = ft.scanner.NextFloat64()
	var n = ft.scanner.NextInt()
	var pro = 1.0
	for index := 0; index < n; index++ {
		pro *= a
	}
	return ft.checker.CompareFloat64(2, pro)
}

// { "no": 16, "dat": "2", "ans": "2" }
func (ft *forTask) for16() bool {
	var a = ft.scanner.NextFloat64()
	var n = ft.scanner.NextInt()
	var pro = 1.0
	for index := 0; index < n; index++ {
		pro *= a
		if !ft.checker.CompareFloat64(2, pro) {
			return false
		}
	}
	return true
}

// { "no": 17, "dat": "2", "ans": "2" }
func (ft *forTask) for17() bool {
	var a = ft.scanner.NextFloat64()
	var n = ft.scanner.NextInt()
	var degree = 1.0
	var sum = 0.0
	for index := 0; index <= n; index++ {
		sum += degree
		degree *= a
	}
	return ft.checker.CompareFloat64(2, sum)
}

// { "no": 18, "dat": "2", "ans": "2" }
func (ft *forTask) for18() bool {
	var (
		a      = ft.scanner.NextFloat64()
		n      = ft.scanner.NextInt()
		degree = 1.0
		result = 0.0
		sign   = 1.0
	)
	for index := 0; index <= n; index++ {
		result += sign * degree
		sign *= -1
		degree *= a
	}
	return ft.checker.CompareFloat64(2, result)
}

// { "no": 19, "dat": "", "ans": "2" }
func (ft *forTask) for19() bool {
	var n = ft.scanner.NextInt()
	var factorial = 1.0
	for index := 1; index <= n; index++ {
		factorial *= float64(index)
	}
	return ft.checker.CompareFloat64(2, factorial)
}

// { "no": 20, "dat": "", "ans": "2" }
func (ft *forTask) for20() bool {
	var n = ft.scanner.NextInt()
	var factorial = 1.0
	var sum = 0.0
	for index := 1; index <= n; index++ {
		factorial *= float64(index)
		sum += factorial
	}
	return ft.checker.CompareFloat64(2, sum)
}

// { "no": 21, "dat": "", "ans": "8" }
func (ft *forTask) for21() bool {
	var n = ft.scanner.NextInt()
	var sum = 0.0
	var factorial = 1.0
	for index := 0; index <= n; {
		sum += 1.0 / factorial
		index++
		factorial *= float64(index)
	}
	return ft.checker.CompareFloat64(8, sum)
}

// { "no": 22, "dat": "8", "ans": "8" }
func (ft *forTask) for22() bool {
	var x = ft.scanner.NextFloat64()
	var n = ft.scanner.NextInt()
	var sum = 0.0
	var degree = 1.0
	var factorial = 1.0
	for index := 0; index <= n; {
		sum += degree / factorial
		degree *= x
		index++
		factorial *= float64(index)
	}
	return ft.checker.CompareFloat64(8, sum)
}

// { "no": 23, "dat": "8", "ans": "8" }
func (ft *forTask) for23() bool {
	var x = ft.scanner.NextFloat64()
	var n = ft.scanner.NextInt()
	var result = 0.0
	var degree = 1.0
	var factorial = 1.0
	var sign = 1.0
	for index := 0; index <= n; {
		degree *= x
		result += sign * degree / factorial
		index++
		sign *= -1
		degree *= x
		factorial *= float64(2*index) * float64(2*index+1)
	}
	return ft.checker.CompareFloat64(8, result)
}

// { "no": 24, "dat": "8", "ans": "8" }
func (ft *forTask) for24() bool {
	var x = ft.scanner.NextFloat64()
	var n = ft.scanner.NextInt()
	var result = 0.0
	var degree = 1.0
	var factorial = 1.0
	var sign = 1.0
	for index := 0; index <= n; {
		result += sign * degree / factorial
		index++
		sign *= -1
		degree *= x * x
		factorial *= float64(2*index-1) * float64(2*index)
	}
	return ft.checker.CompareFloat64(8, result)
}

// { "no": 25, "dat": "8", "ans": "8" }
func (ft *forTask) for25() bool {
	var x = ft.scanner.NextFloat64()
	var n = ft.scanner.NextInt()
	var result = 0.0
	var sign = 1.0
	var degree = 1.0
	for index := 1; index <= n; index++ {
		degree *= x
		result += sign * degree / float64(index)
		sign *= -1
	}
	return ft.checker.CompareFloat64(8, result)
}

// { "no": 26, "dat": "8", "ans": "8" }
func (ft *forTask) for26() bool {
	var x = ft.scanner.NextFloat64()
	var n = ft.scanner.NextInt()
	var result = 0.0
	var degree = 1.0
	var sign = 1.0
	for index := 0; index <= n; index++ {
		degree *= x
		result += sign * degree / float64(2*index+1)
		sign *= -1
		degree *= x
	}
	return ft.checker.CompareFloat64(8, result)
}

// { "no": 27, "dat": "8", "ans": "8" }
func (ft *forTask) for27() bool {
	var x = ft.scanner.NextFloat64()
	var n = ft.scanner.NextInt()
	var sum = 0.0
	var oddsPro, evensPro, degree = 1.0, 1.0, 1.0
	for index := 0; index <= n; {
		degree *= x
		sum += oddsPro * degree / (evensPro * float64(2*index+1))
		index++
		evensPro *= float64(2 * index)
		oddsPro *= float64(2*index - 1)
		degree *= x
	}
	return ft.checker.CompareFloat64(8, sum)
}

// { "no": 28, "dat": "8", "ans": "8" }
func (ft *forTask) for28() bool {
	var x = ft.scanner.NextFloat64()
	var n = ft.scanner.NextInt()
	var result = 1.0
	var oddsPro, evensPro, degree = 1.0, 1.0, 1.0
	var sign = 1.0
	for index := 1; index <= n; {
		degree *= x
		evensPro *= float64(2 * index)
		result += sign * oddsPro * degree / evensPro
		index++
		sign *= -1
		oddsPro *= float64(2*index - 3)
	}
	return ft.checker.CompareFloat64(8, result)
}

// { "no": 29, "dat": "2", "ans": "5" }
func (ft *forTask) for29() bool {
	var n = ft.scanner.NextInt()
	var a = ft.scanner.NextFloat64()
	var b = ft.scanner.NextFloat64()
	var h = (b - a) / float64(n)
	if !ft.checker.CompareFloat64(5, h) {
		return false
	}
	var sequence = a
	for index := 0; index <= n; index++ {
		if !ft.checker.CompareFloat64(5, sequence) {
			return false
		}
		sequence += h
	}
	return true
}

// { "no": 30, "dat": "2", "ans": "5" }
func (ft *forTask) for30() bool {
	var n = ft.scanner.NextInt()
	var a = ft.scanner.NextFloat64()
	var b = ft.scanner.NextFloat64()
	var h = (b - a) / float64(n)
	if !ft.checker.CompareFloat64(5, h) {
		return false
	}
	var fx float64
	var sequence = a
	for index := 0; index <= n; index++ {
		fx = 1 - math.Sin(sequence)
		if !ft.checker.CompareFloat64(5, fx) {
			return false
		}
		sequence += h
	}
	return true
}

// { "no": 31, "dat": "", "ans": "6" }
func (ft *forTask) for31() bool {
	var n = ft.scanner.NextInt()
	var a = 2.0
	for index := 0; index < n; index++ {
		a = 2 + 1/a
		if !ft.checker.CompareFloat64(6, a) {
			return false
		}
	}
	return true
}

// { "no": 32, "dat": "", "ans": "6" }
func (ft *forTask) for32() bool {
	var n = ft.scanner.NextInt()
	var a = 1.0
	for index := 1; index <= n; index++ {
		a = (a + 1) / float64(index)
		if !ft.checker.CompareFloat64(6, a) {
			return false
		}
	}
	return true
}

// { "no": 33, "dat": "", "ans": "" }
func (ft *forTask) for33() bool {
	var n = ft.scanner.NextInt()
	var prevF, nextF = 1, 1
	if !ft.checker.CompareInt(prevF) {
		return false
	}
	for index := 2; index <= n; index++ {
		if !ft.checker.CompareInt(nextF) {
			return false
		}
		prevF, nextF = nextF, nextF+prevF
	}
	return true
}

// { "no": 34, "dat": "", "ans": "6" }
func (ft *forTask) for34() bool {
	var n = ft.scanner.NextInt()
	var prevA, nextA = 1.0, 2.0
	if !ft.checker.CompareFloat64(6, prevA) {
		return false
	}
	for index := 2; index <= n; index++ {
		if !ft.checker.CompareFloat64(6, nextA) {
			return false
		}
		prevA, nextA = nextA, (prevA+2*nextA)/3
	}
	return true
}

// { "no": 35, "dat": "", "ans": "" }
func (ft *forTask) for35() bool {
	var n = ft.scanner.NextInt()
	var prevA, currA, nextA = 1, 2, 3
	if !ft.checker.CompareInt(prevA) {
		return false
	}
	for index := 2; index <= n; index++ {
		if !ft.checker.CompareInt(currA) {
			return false
		}
		prevA, currA, nextA = currA, nextA, nextA+currA-2*prevA
	}
	return true
}

// { "no": 36, "dat": "", "ans": "2" }
func (ft *forTask) for36() bool {
	var n = ft.scanner.NextInt()
	var k = ft.scanner.NextInt()
	var sum = 0.0
	var degreeK float64
	for index := 1; index <= n; index++ {
		degreeK = 1.0
		for i := 0; i < k; i++ {
			degreeK *= float64(index)
		}
		sum += degreeK
	}
	return ft.checker.CompareFloat64(2, sum)
}

// { "no": 37, "dat": "", "ans": "2" }
func (ft *forTask) for37() bool {
	var n = ft.scanner.NextInt()
	var sum = 0.0
	var degreeK float64
	for index := 1; index <= n; index++ {
		degreeK = 1.0
		for i := 0; i < index; i++ {
			degreeK *= float64(index)
		}
		sum += degreeK
	}
	return ft.checker.CompareFloat64(2, sum)
}

// { "no": 38, "dat": "", "ans": "2" }
func (ft *forTask) for38() bool {
	var n = ft.scanner.NextInt()
	var k int
	var sum = 0.0
	var degreeK float64
	for index := 1; index <= n; index++ {
		degreeK = 1.0
		k = n - index + 1
		for i := 0; i < k; i++ {
			degreeK *= float64(index)
		}
		sum += degreeK
	}
	return ft.checker.CompareFloat64(2, sum)
}

// { "no": 39, "dat": "", "ans": "" }
func (ft *forTask) for39() bool {
	var a = ft.scanner.NextInt()
	var b = ft.scanner.NextInt()
	for index := a; index <= b; index++ {
		for i := 0; i < index; i++ {
			if !ft.checker.CompareInt(index) {
				return false
			}
		}
	}
	return true
}

// { "no": 40, "dat": "", "ans": "" }
func (ft *forTask) for40() bool {
	var a = ft.scanner.NextInt()
	var b = ft.scanner.NextInt()
	var limit = 1
	for index := a; index <= b; index++ {
		for i := 0; i < limit; i++ {
			if !ft.checker.CompareInt(index) {
				return false
			}
		}
		limit++
	}
	return true
}
