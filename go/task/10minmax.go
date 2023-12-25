package task

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/proc"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type minmaxTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewMinmaxTask(pathPrefix string) Task {
	return &minmaxTask{
		dataPath: pathPrefix + "/10minmax/Minmax%03d/%03d/%03d",
		name:     "Minmax",
		count:    30,
	}
}

func (mt *minmaxTask) GetCount() int { return mt.count }

func (mt *minmaxTask) GetName() string { return mt.name }

func (mt *minmaxTask) SetData(taskNo int, testNo int) error {
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

func (mt *minmaxTask) ScannerEOF() bool { return mt.scanner.EOF() }

func (mt *minmaxTask) CheckerEOF() bool { return mt.checker.EOF() }

func (mt *minmaxTask) CleanData() {
	mt.scanner.RemoveFiles()
	mt.checker.RemoveFiles()
}

func (mt *minmaxTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return mt.minmax1
	case 2:
		return mt.minmax2
	case 3:
		return mt.minmax3
	case 4:
		return mt.minmax4
	case 5:
		return mt.minmax5
	case 6:
		return mt.minmax6
	case 7:
		return mt.minmax7
	case 8:
		return mt.minmax8
	case 9:
		return mt.minmax9
	case 10:
		return mt.minmax10
	case 11:
		return mt.minmax11
	case 12:
		return mt.minmax12
	case 13:
		return mt.minmax13
	case 14:
		return mt.minmax14
	case 15:
		return mt.minmax15
	case 16:
		return mt.minmax16
	case 17:
		return mt.minmax17
	case 18:
		return mt.minmax18
	case 19:
		return mt.minmax19
	case 20:
		return mt.minmax20
	case 21:
		return mt.minmax21
	case 22:
		return mt.minmax22
	case 23:
		return mt.minmax23
	case 24:
		return mt.minmax24
	case 25:
		return mt.minmax25
	case 26:
		return mt.minmax26
	case 27:
		return mt.minmax27
	case 28:
		return mt.minmax28
	case 29:
		return mt.minmax29
	case 30:
		return mt.minmax30
	default:
		return nil
	}
}

// { "no": 1, "dat": "2", "ans": "2" }
func (mt *minmaxTask) minmax1() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextFloat64()
	var minimal, maximal = number, number
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextFloat64()
		if number < minimal {
			minimal = number
		} else if number > maximal {
			maximal = number
		}
	}
	return mt.checker.CompareFloat64(2, minimal, maximal)
}

// { "no": 2, "dat": "2", "ans": "2" }
func (mt *minmaxTask) minmax2() bool {
	var n = mt.scanner.NextInt()
	var a = mt.scanner.NextFloat64()
	var b = mt.scanner.NextFloat64()
	var minArea = a * b
	var area float64
	for index := 2; index <= n; index++ {
		a = mt.scanner.NextFloat64()
		b = mt.scanner.NextFloat64()
		area = a * b
		if area < minArea {
			minArea = area
		}
	}
	return mt.checker.CompareFloat64(2, minArea)
}

// { "no": 3, "dat": "2", "ans": "2" }
func (mt *minmaxTask) minmax3() bool {
	var n = mt.scanner.NextInt()
	var a = mt.scanner.NextFloat64()
	var b = mt.scanner.NextFloat64()
	var maxPerimeter = 2 * (a + b)
	var perimeter float64
	for index := 2; index <= n; index++ {
		a = mt.scanner.NextFloat64()
		b = mt.scanner.NextFloat64()
		perimeter = 2 * (a + b)
		if perimeter > maxPerimeter {
			maxPerimeter = perimeter
		}
	}
	return mt.checker.CompareFloat64(2, maxPerimeter)
}

// { "no": 4, "dat": "2", "ans": "" }
func (mt *minmaxTask) minmax4() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextFloat64()
	var minimal = number
	var minimalIndex = 1
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextFloat64()
		if number <= minimal {
			minimal, minimalIndex = number, index
		}
	}
	return mt.checker.CompareInt(minimalIndex)
}

// { "no": 5, "dat": "2", "ans": "2" }
func (mt *minmaxTask) minmax5() bool {
	var n = mt.scanner.NextInt()
	var m = mt.scanner.NextFloat64()
	var v = mt.scanner.NextFloat64()
	var maxDensity = m / v
	var density float64
	var maxDensityIndex = 1
	for index := 2; index <= n; index++ {
		m = mt.scanner.NextFloat64()
		v = mt.scanner.NextFloat64()
		density = m / v
		if density > maxDensity {
			maxDensity, maxDensityIndex = density, index
		}
	}
	if !mt.checker.CompareInt(maxDensityIndex) {
		return false
	}
	return mt.checker.CompareFloat64(2, maxDensity)
}

// { "no": 6, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax6() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextInt()
	var firstMinimal, lastMaximal = number, number
	var firstMinimalIndex, lastMaximalIndex = 1, 1
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number < firstMinimal {
			firstMinimal, firstMinimalIndex = number, index
		} else if number >= lastMaximal {
			lastMaximal, lastMaximalIndex = number, index
		}
	}
	return mt.checker.CompareInt(firstMinimalIndex, lastMaximalIndex)
}

// { "no": 7, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax7() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextInt()
	var firstMaximal, lastMinimal = number, number
	var firstMaximalIndex, lastMinimalIndex = 1, 1
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number > firstMaximal {
			firstMaximal, firstMaximalIndex = number, index
		} else if number <= lastMinimal {
			lastMinimal, lastMinimalIndex = number, index
		}
	}
	return mt.checker.CompareInt(firstMaximalIndex, lastMinimalIndex)
}

// { "no": 8, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax8() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextInt()
	var minimal = number
	var firstMinimalIndex, lastMinimalIndex = 1, 1
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number == minimal {
			lastMinimalIndex = index
		} else if number < minimal {
			minimal = number
			firstMinimalIndex, lastMinimalIndex = index, index
		}
	}
	return mt.checker.CompareInt(firstMinimalIndex, lastMinimalIndex)
}

// { "no": 9, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax9() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextInt()
	var maximal = number
	var firstMaximalIndex, lastMaximalIndex = 1, 1
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number == maximal {
			lastMaximalIndex = index
		} else if number > maximal {
			maximal = number
			firstMaximalIndex, lastMaximalIndex = index, index
		}
	}
	return mt.checker.CompareInt(firstMaximalIndex, lastMaximalIndex)
}

// { "no": 10, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax10() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextInt()
	var minimal, maximal = number, number
	var firstMinimalIndex, firstMaximalIndex = 1, 1
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number < minimal {
			minimal = number
			firstMinimalIndex = index
		} else if number > maximal {
			maximal = number
			firstMaximalIndex = index
		}
	}
	var firstExtremalIndex = firstMinimalIndex
	if firstMaximalIndex < firstMinimalIndex {
		firstExtremalIndex = firstMaximalIndex
	}
	return mt.checker.CompareInt(firstExtremalIndex)
}

// { "no": 11, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax11() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextInt()
	var minimal, maximal = number, number
	var lastMinimalIndex, lastMaximalIndex = 1, 1
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number <= minimal {
			minimal = number
			lastMinimalIndex = index
		} else if number >= maximal {
			maximal = number
			lastMaximalIndex = index
		}
	}
	var lastExtremalIndex = lastMinimalIndex
	if lastMaximalIndex > lastMinimalIndex {
		lastExtremalIndex = lastMaximalIndex
	}
	return mt.checker.CompareInt(lastExtremalIndex)
}

// { "no": 12, "dat": "2", "ans": "2" }
func (mt *minmaxTask) minmax12() bool {
	var n = mt.scanner.NextInt()
	var number float64
	var minPositive = 0.0
	for index := 0; index < n; index++ {
		number = mt.scanner.NextFloat64()
		if number > 0 && (minPositive == 0.0 || number < minPositive) {
			minPositive = number
		}
	}
	return mt.checker.CompareFloat64(2, minPositive)
}

// { "no": 13, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax13() bool {
	var n = mt.scanner.NextInt()
	var number, maxOdd int
	var firstMaxOddIndex = 0
	for index := 1; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number%2 != 0 && (firstMaxOddIndex == 0 || number > maxOdd) {
			maxOdd = number
			firstMaxOddIndex = index
		}
	}
	return mt.checker.CompareInt(firstMaxOddIndex)
}

// { "no": 14, "dat": "2", "ans": "2" }
func (mt *minmaxTask) minmax14() bool {
	var b = mt.scanner.NextFloat64()
	var number float64
	var minimal, minimalIndex = 0.0, 0
	for index := 1; index <= 10; index++ {
		number = mt.scanner.NextFloat64()
		if number > b && (minimalIndex == 0 || number < minimal) {
			minimal, minimalIndex = number, index
		}
	}
	if !mt.checker.CompareFloat64(2, minimal) {
		return false
	}
	return mt.checker.CompareInt(minimalIndex)
}

// { "no": 15, "dat": "2", "ans": "2" }
func (mt *minmaxTask) minmax15() bool {
	var b = mt.scanner.NextFloat64()
	var c = mt.scanner.NextFloat64()
	var number float64
	var maximal, maximalIndex = 0.0, 0
	for index := 1; index <= 10; index++ {
		number = mt.scanner.NextFloat64()
		if b < number && number < c && (maximalIndex == 0 || number > maximal) {
			maximal, maximalIndex = number, index
		}
	}
	if !mt.checker.CompareFloat64(2, maximal) {
		return false
	}
	return mt.checker.CompareInt(maximalIndex)
}

// { "no": 16, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax16() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextInt()
	var firstMinimal = number
	var firstMinimalIndex = 1
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number < firstMinimal {
			firstMinimal, firstMinimalIndex = number, index
		}
	}
	return mt.checker.CompareInt(firstMinimalIndex - 1)
}

// { "no": 17, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax17() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextInt()
	var lastMaximal = number
	var lastMaximalIndex = 1
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number >= lastMaximal {
			lastMaximal, lastMaximalIndex = number, index
		}
	}
	return mt.checker.CompareInt(n - lastMaximalIndex)
}

// { "no": 18, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax18() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextInt()
	var maximal = number
	var firstMaximalIndex, lastMaximalIndex = 1, 1
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number == maximal {
			lastMaximalIndex = index
		} else if number > maximal {
			maximal = number
			firstMaximalIndex, lastMaximalIndex = index, index
		}
	}
	var count = 0
	if firstMaximalIndex != lastMaximalIndex {
		count = lastMaximalIndex - firstMaximalIndex - 1
	}
	return mt.checker.CompareInt(count)
}

// { "no": 19, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax19() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextInt()
	var minimal = number
	var minimalsCount = 1
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number == minimal {
			minimalsCount++
		} else if number < minimal {
			minimal, minimalsCount = number, 1
		}
	}
	return mt.checker.CompareInt(minimalsCount)
}

// { "no": 20, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax20() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextInt()
	var minimal, maximal = number, number
	var minimalsCount, maximalsCount = 1, 1
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number == minimal {
			minimalsCount++
		} else if number < minimal {
			minimal, minimalsCount = number, 1
		}
		if number == maximal {
			maximalsCount++
		} else if number > maximal {
			maximal, maximalsCount = number, 1
		}
	}
	var extremalsCount = minimalsCount + maximalsCount
	if minimal == maximal {
		extremalsCount = minimalsCount
	}
	return mt.checker.CompareInt(extremalsCount)
}

// { "no": 21, "dat": "2", "ans": "2" }
func (mt *minmaxTask) minmax21() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextFloat64()
	var minimal, maximal = number, number
	var sum = number
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextFloat64()
		sum += number
		if number < minimal {
			minimal = number
		} else if number > maximal {
			maximal = number
		}
	}
	sum -= minimal + maximal
	sum /= float64(n - 2)
	return mt.checker.CompareFloat64(2, sum)
}

// { "no": 22, "dat": "2", "ans": "2" }
func (mt *minmaxTask) minmax22() bool {
	var n = mt.scanner.NextInt()
	var minimal1 = mt.scanner.NextFloat64()
	var minimal2 = mt.scanner.NextFloat64()
	proc.Minmax(&minimal1, &minimal2)
	var number float64
	for index := 3; index <= n; index++ {
		number = mt.scanner.NextFloat64()
		if number < minimal2 {
			minimal2 = number
			proc.Minmax(&minimal1, &minimal2)
		}
	}
	return mt.checker.CompareFloat64(2, minimal1, minimal2)
}

// { "no": 23, "dat": "2", "ans": "2" }
func (mt *minmaxTask) minmax23() bool {
	var n = mt.scanner.NextInt()
	var maximal1 = mt.scanner.NextFloat64()
	var maximal2 = mt.scanner.NextFloat64()
	var maximal3 = mt.scanner.NextFloat64()
	proc.SortDec3(&maximal1, &maximal2, &maximal3)
	var number float64
	for index := 4; index <= n; index++ {
		number = mt.scanner.NextFloat64()
		if number > maximal3 {
			maximal3 = number
			proc.SortDec3(&maximal1, &maximal2, &maximal3)
		}
	}
	return mt.checker.CompareFloat64(2, maximal1, maximal2, maximal3)
}

// { "no": 24, "dat": "2", "ans": "2" }
func (mt *minmaxTask) minmax24() bool {
	var n = mt.scanner.NextInt()
	var prev = mt.scanner.NextFloat64()
	var curr, sum float64
	var maxSum = prev
	for index := 2; index <= n; index++ {
		curr = mt.scanner.NextFloat64()
		sum = prev + curr
		if index == 2 || sum > maxSum {
			maxSum = sum
		}
		prev = curr
	}
	return mt.checker.CompareFloat64(2, maxSum)
}

// { "no": 25, "dat": "2", "ans": "" }
func (mt *minmaxTask) minmax25() bool {
	var n = mt.scanner.NextInt()
	var prev = mt.scanner.NextFloat64()
	var pro, curr float64
	var minPro = prev
	var proIndex = 1
	for index := 2; index <= n; index++ {
		curr = mt.scanner.NextFloat64()
		pro = prev * curr
		if index == 2 || pro < minPro {
			minPro = pro
			proIndex = index
		}
		prev = curr
	}
	return mt.checker.CompareInt(proIndex-1, proIndex)
}

// { "no": 26, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax26() bool {
	var n = mt.scanner.NextInt()
	var number int
	var count, maxCount = 0, 0
	for index := 1; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number%2 == 0 {
			count++
		} else if count != 0 {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		}
	}
	if count > maxCount {
		maxCount = count
	}
	return mt.checker.CompareInt(maxCount)
}

// { "no": 27, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax27() bool {
	var n = mt.scanner.NextInt()
	var prev = mt.scanner.NextInt()
	var curr int
	var count, maxCount, maxIndex = 1, 1, 1
	for index := 2; index <= n; index++ {
		curr = mt.scanner.NextInt()
		if prev == curr {
			count++
		} else {
			if count > maxCount {
				maxCount, maxIndex = count, index-count
			}
			count = 1
		}
		prev = curr
	}
	if count > maxCount {
		maxCount, maxIndex = count, n-count+1
	}
	return mt.checker.CompareInt(maxIndex, maxCount)
}

// { "no": 28, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax28() bool {
	var n = mt.scanner.NextInt()
	var count, maxCount, maxIndex = 0, 0, 0
	var number int
	for index := 1; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number == 1 {
			count++
		} else if count != 0 {
			if count >= maxCount {
				maxCount, maxIndex = count, index-count
			}
			count = 0
		}
	}
	if count != 0 && count >= maxCount {
		maxCount, maxIndex = count, n-count+1
	}
	return mt.checker.CompareInt(maxIndex, maxCount)
}

// { "no": 29, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax29() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextInt()
	var minimal = number
	var count, maxCount = 1, 1
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number == minimal {
			count++
		} else if number < minimal {
			minimal = number
			maxCount, count = 1, 1
		} else {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		}
	}
	if count > maxCount {
		maxCount = count
	}
	return mt.checker.CompareInt(maxCount)
}

// { "no": 30, "dat": "", "ans": "" }
func (mt *minmaxTask) minmax30() bool {
	var n = mt.scanner.NextInt()
	var number = mt.scanner.NextInt()
	var maximal = number
	var count, minCount = 1, 0
	for index := 2; index <= n; index++ {
		number = mt.scanner.NextInt()
		if number == maximal {
			count++
		} else if number > maximal {
			maximal = number
			count, minCount = 1, 0
		} else if count != 0 {
			if minCount == 0 || count < minCount {
				minCount = count
			}
			count = 0
		}
	}
	if count != 0 && (minCount == 0 || count < minCount) {
		minCount = count
	}
	return mt.checker.CompareInt(minCount)
}
