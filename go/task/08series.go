package task

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type seriesTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewSeriesTask(pathPrefix string) Task {
	return &seriesTask{
		dataPath: pathPrefix + "/08series/Series%03d/%03d/%03d",
		name:     "Series",
		count:    40,
	}
}

func (st *seriesTask) GetCount() int { return st.count }

func (st *seriesTask) GetName() string { return st.name }

func (st *seriesTask) SetData(taskNo int, testNo int) error {
	var filePath = fmt.Sprintf(st.dataPath, taskNo, testNo, testNo)
	scanner, err := scanner.New(filePath + ".dat")
	if err != nil {
		return errors.Wrap(err, "scanner.New")
	}
	checker, err := checker.New(filePath + ".ans")
	if err != nil {
		return errors.Wrap(err, "checker.New")
	}
	st.scanner = scanner
	st.checker = checker
	return nil
}

func (st *seriesTask) ScannerEOF() bool { return st.scanner.EOF() }

func (st *seriesTask) CheckerEOF() bool { return st.checker.EOF() }

func (st *seriesTask) CleanData() {
	st.scanner.RemoveFiles()
	st.checker.RemoveFiles()
}

func (st *seriesTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return st.series1
	case 2:
		return st.series2
	case 3:
		return st.series3
	case 4:
		return st.series4
	case 5:
		return st.series5
	case 6:
		return st.series6
	case 7:
		return st.series7
	case 8:
		return st.series8
	case 9:
		return st.series9
	case 10:
		return st.series10
	case 11:
		return st.series11
	case 12:
		return st.series12
	case 13:
		return st.series13
	case 14:
		return st.series14
	case 15:
		return st.series15
	case 16:
		return st.series16
	case 17:
		return st.series17
	case 18:
		return st.series18
	case 19:
		return st.series19
	case 20:
		return st.series20
	case 21:
		return st.series21
	case 22:
		return st.series22
	case 23:
		return st.series23
	case 24:
		return st.series24
	case 25:
		return st.series25
	case 26:
		return st.series26
	case 27:
		return st.series27
	case 28:
		return st.series28
	case 29:
		return st.series29
	case 30:
		return st.series30
	case 31:
		return st.series31
	case 32:
		return st.series32
	case 33:
		return st.series33
	case 34:
		return st.series34
	case 35:
		return st.series35
	case 36:
		return st.series36
	case 37:
		return st.series37
	case 38:
		return st.series38
	case 39:
		return st.series39
	case 40:
		return st.series40
	default:
		return nil
	}
}

// { "no": 1, "dat": "2", "ans": "2" }
func (st *seriesTask) series1() bool {
	var number float64
	var sum = 0.0
	for index := 0; index < 10; index++ {
		number = st.scanner.NextFloat64()
		sum += number
	}
	return st.checker.CompareFloat64(2, sum)
}

// { "no": 2, "dat": "2", "ans": "2" }
func (st *seriesTask) series2() bool {
	var number float64
	var pro = 1.0
	for index := 0; index < 10; index++ {
		number = st.scanner.NextFloat64()
		pro *= number
	}
	return st.checker.CompareFloat64(2, pro)
}

// { "no": 3, "dat": "2", "ans": "2" }
func (st *seriesTask) series3() bool {
	var number float64
	var sum = 0.0
	for index := 0; index < 10; index++ {
		number = st.scanner.NextFloat64()
		sum += number
	}
	var average = sum / 10
	return st.checker.CompareFloat64(2, average)
}

// { "no": 4, "dat": "2", "ans": "2" }
func (st *seriesTask) series4() bool {
	var n = st.scanner.NextInt()
	var number float64
	var sum, pro = 0.0, 1.0
	for index := 0; index < n; index++ {
		number = st.scanner.NextFloat64()
		sum += number
		pro *= number
	}
	return st.checker.CompareFloat64(2, sum, pro)
}

// { "no": 5, "dat": "2", "ans": "2" }
func (st *seriesTask) series5() bool {
	var n = st.scanner.NextInt()
	var number, integerPart float64
	var sum = 0.0
	for index := 0; index < n; index++ {
		number = st.scanner.NextFloat64()
		integerPart = float64(int(number))
		if !st.checker.CompareFloat64(2, integerPart) {
			return false
		}
		sum += integerPart
	}
	return st.checker.CompareFloat64(2, sum)
}

// { "no": 6, "dat": "2", "ans": "2{N}, 6" }
func (st *seriesTask) series6() bool {
	var n = st.scanner.NextInt()
	var number, fractionalPart float64
	var pro = 1.0
	for index := 0; index < n; index++ {
		number = st.scanner.NextFloat64()
		fractionalPart = number - float64(int(number))
		if !st.checker.CompareFloat64(2, fractionalPart) {
			return false
		}
		pro *= fractionalPart
	}
	return st.checker.CompareFloat64(6, pro)
}

// { "no": 7, "dat": "2", "ans": "" }
func (st *seriesTask) series7() bool {
	var n = st.scanner.NextInt()
	var number float64
	var rounded int
	var sum = 0
	for index := 0; index < n; index++ {
		number = st.scanner.NextFloat64()
		rounded = int(math.Round(number))
		if !st.checker.CompareInt(rounded) {
			return false
		}
		sum += rounded
	}
	return st.checker.CompareInt(sum)
}

// { "no": 8, "dat": "", "ans": "" }
func (st *seriesTask) series8() bool {
	var n = st.scanner.NextInt()
	var number int
	var k = 0
	for index := 0; index < n; index++ {
		number = st.scanner.NextInt()
		if number%2 == 0 {
			if !st.checker.CompareInt(number) {
				return false
			}
			k++
		}
	}
	return st.checker.CompareInt(k)
}

// { "no": 9, "dat": "", "ans": "" }
func (st *seriesTask) series9() bool {
	var n = st.scanner.NextInt()
	var number int
	var k = 0
	for index := 1; index <= n; index++ {
		number = st.scanner.NextInt()
		if number%2 != 0 {
			if !st.checker.CompareInt(index) {
				return false
			}
			k++
		}
	}
	return st.checker.CompareInt(k)
}

// { "no": 10, "dat": "", "ans": "" }
func (st *seriesTask) series10() bool {
	var n = st.scanner.NextInt()
	var number int
	var hasPositive = false
	for index := 0; index < n; index++ {
		number = st.scanner.NextInt()
		if number > 0 && !hasPositive {
			hasPositive = true
		}
	}
	return st.checker.CompareBool(hasPositive)
}

// { "no": 11, "dat": "", "ans": "" }
func (st *seriesTask) series11() bool {
	var k = st.scanner.NextInt()
	var n = st.scanner.NextInt()
	var number int
	var hasLessThanK = false
	for index := 0; index < n; index++ {
		number = st.scanner.NextInt()
		if number < k && !hasLessThanK {
			hasLessThanK = true
		}
	}
	return st.checker.CompareBool(hasLessThanK)
}

// { "no": 12, "dat": "", "ans": "" }
func (st *seriesTask) series12() bool {
	var number int
	var count = 0
	for {
		number = st.scanner.NextInt()
		if number == 0 {
			break
		}
		count++
	}
	return st.checker.CompareInt(count)
}

// { "no": 13, "dat": "", "ans": "" }
func (st *seriesTask) series13() bool {
	var number int
	var sum = 0
	for {
		number = st.scanner.NextInt()
		if number == 0 {
			break
		}
		if number > 0 && number%2 == 0 {
			sum += number
		}
	}
	return st.checker.CompareInt(sum)
}

// { "no": 14, "dat": "", "ans": "" }
func (st *seriesTask) series14() bool {
	var k = st.scanner.NextInt()
	var number int
	var lessThanKCount = 0
	for {
		number = st.scanner.NextInt()
		if number == 0 {
			break
		}
		if number < k {
			lessThanKCount++
		}
	}
	return st.checker.CompareInt(lessThanKCount)
}

// { "no": 15, "dat": "", "ans": "" }
func (st *seriesTask) series15() bool {
	var k = st.scanner.NextInt()
	var number int
	var index, firstGreaterThanKIndex = 0, 0
	for {
		number = st.scanner.NextInt()
		if number == 0 {
			break
		}
		index++
		if number > k && firstGreaterThanKIndex == 0 {
			firstGreaterThanKIndex = index
		}
	}
	return st.checker.CompareInt(firstGreaterThanKIndex)
}

// { "no": 16, "dat": "", "ans": "" }
func (st *seriesTask) series16() bool {
	var k = st.scanner.NextInt()
	var number int
	var index, lastGreaterThanKIndex = 0, 0
	for {
		number = st.scanner.NextInt()
		if number == 0 {
			break
		}
		index++
		if number > k {
			lastGreaterThanKIndex = index
		}
	}
	return st.checker.CompareInt(lastGreaterThanKIndex)
}

// { "no": 17, "dat": "2", "ans": "2" }
func (st *seriesTask) series17() bool {
	var b = st.scanner.NextFloat64()
	var n = st.scanner.NextInt()
	var number float64
	var outed = false
	for index := 0; index < n; index++ {
		number = st.scanner.NextFloat64()
		if !outed && b < number {
			if !st.checker.CompareFloat64(2, b) {
				return false
			}
			outed = true
		}
		if !st.checker.CompareFloat64(2, number) {
			return false
		}
	}
	if !outed {
		return st.checker.CompareFloat64(2, b)
	}
	return true
}

// { "no": 18, "dat": "", "ans": "" }
func (st *seriesTask) series18() bool {
	var n = st.scanner.NextInt()
	var prev, curr int
	for index := 0; index < n; index++ {
		curr = st.scanner.NextInt()
		if index == 0 {
			if !st.checker.CompareInt(curr) {
				return false
			}
		} else {
			if curr != prev {
				if !st.checker.CompareInt(curr) {
					return false
				}
			}
		}
		prev = curr
	}
	return true
}

// { "no": 19, "dat": "", "ans": "" }
func (st *seriesTask) series19() bool {
	var n = st.scanner.NextInt()
	var prev, curr int
	var k = 0
	for index := 0; index < n; index++ {
		curr = st.scanner.NextInt()
		if index > 0 && curr < prev {
			if !st.checker.CompareInt(curr) {
				return false
			}
			k++
		}
		prev = curr
	}
	return st.checker.CompareInt(k)
}

// { "no": 20, "dat": "", "ans": "" }
func (st *seriesTask) series20() bool {
	var n = st.scanner.NextInt()
	var curr, next int
	var k = 0
	for index := 0; index < n; index++ {
		next = st.scanner.NextInt()
		if index > 0 && curr < next {
			if !st.checker.CompareInt(curr) {
				return false
			}
			k++
		}
		curr = next
	}
	return st.checker.CompareInt(k)
}

// { "no": 21, "dat": "2", "ans": "" }
func (st *seriesTask) series21() bool {
	var n = st.scanner.NextInt()
	var prev, curr float64
	var isAscending = true
	for index := 0; index < n; index++ {
		curr = st.scanner.NextFloat64()
		if isAscending && index > 0 && prev >= curr {
			isAscending = false
		}
		prev = curr
	}
	return st.checker.CompareBool(isAscending)
}

// { "no": 22, "dat": "2", "ans": "" }
func (st *seriesTask) series22() bool {
	var n = st.scanner.NextInt()
	var prev, curr float64
	var errIndex = 0
	for index := 1; index <= n; index++ {
		curr = st.scanner.NextFloat64()
		if errIndex == 0 && index > 1 && prev <= curr {
			errIndex = index
		}
		prev = curr
	}
	return st.checker.CompareInt(errIndex)
}

// { "no": 23, "dat": "2", "ans": "" }
func (st *seriesTask) series23() bool {
	var n = st.scanner.NextInt()
	var prev, curr, next float64
	var errIndex = 0
	for index := 1; index <= n; index++ {
		next = st.scanner.NextFloat64()
		if errIndex == 0 && index > 2 &&
			(prev <= curr && curr <= next || prev >= curr && curr >= next) {
			errIndex = index - 1
		}
		prev, curr = curr, next
	}
	return st.checker.CompareInt(errIndex)
}

// { "no": 24, "dat": "", "ans": "" }
func (st *seriesTask) series24() bool {
	var n = st.scanner.NextInt()
	var number int
	var prevSum, sum = 0, 0
	for index := 0; index < n; index++ {
		number = st.scanner.NextInt()
		if number == 0 {
			prevSum, sum = sum, 0
		}
		sum += number
	}
	sum = prevSum
	return st.checker.CompareInt(sum)
}

// { "no": 25, "dat": "", "ans": "" }
func (st *seriesTask) series25() bool {
	var n = st.scanner.NextInt()
	var number int
	var totalSum, sum = 0, 0
	var firstTime = true
	for index := 0; index < n; index++ {
		number = st.scanner.NextInt()
		if number == 0 {
			if firstTime {
				firstTime = false
			} else {
				totalSum += sum
			}
			sum = 0
		}
		sum += number
	}
	sum = totalSum
	return st.checker.CompareInt(sum)
}

// { "no": 26, "dat": "2", "ans": "2" }
func (st *seriesTask) series26() bool {
	var k = st.scanner.NextInt()
	var n = st.scanner.NextInt()
	var number, degreeK float64
	for index := 0; index < n; index++ {
		number = st.scanner.NextFloat64()
		degreeK = 1.0
		for i := 0; i < k; i++ {
			degreeK *= number
		}
		if !st.checker.CompareFloat64(2, degreeK) {
			return false
		}
	}
	return true
}

// { "no": 27, "dat": "2", "ans": "2" }
func (st *seriesTask) series27() bool {
	var n = st.scanner.NextInt()
	var number, degreeK float64
	for index := 1; index <= n; index++ {
		number = st.scanner.NextFloat64()
		degreeK = 1.0
		for i := 0; i < index; i++ {
			degreeK *= number
		}
		if !st.checker.CompareFloat64(2, degreeK) {
			return false
		}
	}
	return true
}

// { "no": 28, "dat": "2", "ans": "2" }
func (st *seriesTask) series28() bool {
	var n = st.scanner.NextInt()
	var number, degreeK float64
	for index := 1; index <= n; index++ {
		number = st.scanner.NextFloat64()
		degreeK = 1.0
		for i := 0; i <= n-index; i++ {
			degreeK *= number
		}
		if !st.checker.CompareFloat64(2, degreeK) {
			return false
		}
	}
	return true
}

// { "no": 29, "dat": "", "ans": "" }
func (st *seriesTask) series29() bool {
	var k, n = st.scanner.NextInt(), st.scanner.NextInt()
	var number int
	var sum = 0
	for index := 0; index < k; index++ {
		for i := 0; i < n; i++ {
			number = st.scanner.NextInt()
			sum += number
		}
	}
	return st.checker.CompareInt(sum)
}

// { "no": 30, "dat": "", "ans": "" }
func (st *seriesTask) series30() bool {
	var k, n = st.scanner.NextInt(), st.scanner.NextInt()
	var number, sum int
	for index := 0; index < k; index++ {
		sum = 0
		for i := 0; i < n; i++ {
			number = st.scanner.NextInt()
			sum += number
		}
		if !st.checker.CompareInt(sum) {
			return false
		}
	}
	return true
}

// { "no": 31, "dat": "", "ans": "" }
func (st *seriesTask) series31() bool {
	var k, n = st.scanner.NextInt(), st.scanner.NextInt()
	var number int
	var hasTwo bool
	var count = 0
	for index := 0; index < k; index++ {
		hasTwo = false
		for i := 0; i < n; i++ {
			number = st.scanner.NextInt()
			if !hasTwo && number == 2 {
				hasTwo = true
			}
		}
		if hasTwo {
			count++
		}
	}
	return st.checker.CompareInt(count)
}

// { "no": 32, "dat": "", "ans": "" }
func (st *seriesTask) series32() bool {
	var k, n = st.scanner.NextInt(), st.scanner.NextInt()
	var number, firstTwoIndex int
	for index := 0; index < k; index++ {
		firstTwoIndex = 0
		for i := 1; i <= n; i++ {
			number = st.scanner.NextInt()
			if firstTwoIndex == 0 && number == 2 {
				firstTwoIndex = i
			}
		}
		if !st.checker.CompareInt(firstTwoIndex) {
			return false
		}
	}
	return true
}

// { "no": 33, "dat": "", "ans": "" }
func (st *seriesTask) series33() bool {
	var k, n = st.scanner.NextInt(), st.scanner.NextInt()
	var number, firstTwoIndex int
	for index := 0; index < k; index++ {
		firstTwoIndex = 0
		for i := 1; i <= n; i++ {
			number = st.scanner.NextInt()
			if number == 2 {
				firstTwoIndex = i
			}
		}
		if !st.checker.CompareInt(firstTwoIndex) {
			return false
		}
	}
	return true
}

// { "no": 34, "dat": "", "ans": "" }
func (st *seriesTask) series34() bool {
	var k, n = st.scanner.NextInt(), st.scanner.NextInt()
	var number, sum int
	var hasTwo bool
	for index := 0; index < k; index++ {
		sum = 0
		hasTwo = false
		for i := 1; i <= n; i++ {
			number = st.scanner.NextInt()
			if !hasTwo && number == 2 {
				hasTwo = true
			}
			sum += number
		}
		if !hasTwo {
			sum = 0
		}
		if !st.checker.CompareInt(sum) {
			return false
		}
	}
	return true
}

// { "no": 35, "dat": "", "ans": "" }
func (st *seriesTask) series35() bool {
	var k = st.scanner.NextInt()
	var number, count int
	var totalCount = 0
	for index := 0; index < k; index++ {
		count = 0
		for {
			number = st.scanner.NextInt()
			if number == 0 {
				break
			}
			count++
		}
		if !st.checker.CompareInt(count) {
			return false
		}
		totalCount += count
	}
	return st.checker.CompareInt(totalCount)
}

// { "no": 36, "dat": "", "ans": "" }
func (st *seriesTask) series36() bool {
	var k = st.scanner.NextInt()
	var prev, curr int
	var isAscending, isFirstElem bool
	var ascendingSequencesCount = 0
	for index := 0; index < k; index++ {
		isAscending = true
		isFirstElem = true
		for {
			curr = st.scanner.NextInt()
			if curr == 0 {
				break
			}
			if isFirstElem {
				isFirstElem = false
			} else if isAscending && prev > curr {
				isAscending = false
			}
			prev = curr
		}
		if isAscending {
			ascendingSequencesCount++
		}
	}
	return st.checker.CompareInt(ascendingSequencesCount)
}

// { "no": 37, "dat": "", "ans": "" }
func (st *seriesTask) series37() bool {
	var k = st.scanner.NextInt()
	var prev, curr int
	var isAscending, isDescending, isFirstElem bool
	var ascendingSequencesCount = 0
	for index := 0; index < k; index++ {
		isAscending, isDescending, isFirstElem = true, true, true
		for {
			curr = st.scanner.NextInt()
			if curr == 0 {
				break
			}
			if isFirstElem {
				isFirstElem = false
			} else if isAscending && prev > curr {
				isAscending = false
			} else if isDescending && prev < curr {
				isDescending = false
			}
			prev = curr
		}
		if isAscending || isDescending {
			ascendingSequencesCount++
		}
	}
	return st.checker.CompareInt(ascendingSequencesCount)
}

// { "no": 38, "dat": "", "ans": "" }
func (st *seriesTask) series38() bool {
	var k = st.scanner.NextInt()
	var prev, curr, result int
	var isAscending, isDescending, isFirstElem bool
	for index := 0; index < k; index++ {
		isAscending, isDescending, isFirstElem = true, true, true
		for {
			curr = st.scanner.NextInt()
			if curr == 0 {
				break
			}
			if isFirstElem {
				isFirstElem = false
			} else if isAscending && prev > curr {
				isAscending = false
			} else if isDescending && prev < curr {
				isDescending = false
			}
			prev = curr
		}
		if isAscending {
			result = 1
		} else if isDescending {
			result = -1
		} else {
			result = 0
		}
		if !st.checker.CompareInt(result) {
			return false
		}
	}
	return true
}

// { "no": 39, "dat": "", "ans": "" }
func (st *seriesTask) series39() bool {
	var k = st.scanner.NextInt()
	var prev, curr, next, index int
	var sawtoothSequencesCount = 0
	var isSawtooth bool
	for seqIndex := 0; seqIndex < k; seqIndex++ {
		isSawtooth = true
		index = 0
		for {
			next = st.scanner.NextInt()
			if next == 0 {
				break
			}
			index++
			if isSawtooth && index > 2 &&
				(prev <= curr && curr <= next || prev >= curr && curr >= next) {
				isSawtooth = false
			}
			prev, curr = curr, next
		}
		if isSawtooth {
			sawtoothSequencesCount++
		}
	}
	return st.checker.CompareInt(sawtoothSequencesCount)
}

// { "no": 40, "dat": "", "ans": "" }
func (st *seriesTask) series40() bool {
	var k = st.scanner.NextInt()
	var prev, curr, next, index int
	var errIndex int
	for seqIndex := 0; seqIndex < k; seqIndex++ {
		errIndex = 0
		index = 0
		for {
			next = st.scanner.NextInt()
			if next == 0 {
				break
			}
			index++
			if errIndex == 0 && index > 2 &&
				(prev <= curr && curr <= next || prev >= curr && curr >= next) {
				errIndex = index - 1
			}
			prev, curr = curr, next
		}
		var result int
		if errIndex == 0 {
			result = index
		} else {
			result = errIndex
		}
		if !st.checker.CompareInt(result) {
			return false
		}
	}
	return true
}
