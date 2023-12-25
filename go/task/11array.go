package task

import (
	"fmt"
	"math"
	"sort"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/array"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/point"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/proc"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type arrayTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewArrayTask(pathPrefix string) Task {
	return &arrayTask{
		dataPath: pathPrefix + "/11array/Array%03d/%03d/%03d",
		name:     "Array",
		count:    140,
	}
}

func (at *arrayTask) GetCount() int { return at.count }

func (at *arrayTask) GetName() string { return at.name }

func (at *arrayTask) SetData(taskNo int, testNo int) error {
	var filePath = fmt.Sprintf(at.dataPath, taskNo, testNo, testNo)
	scanner, err := scanner.New(filePath + ".dat")
	if err != nil {
		return errors.Wrap(err, "scanner.New")
	}
	checker, err := checker.New(filePath + ".ans")
	if err != nil {
		return errors.Wrap(err, "checker.New")
	}
	at.scanner = scanner
	at.checker = checker
	return nil
}

func (at *arrayTask) ScannerEOF() bool { return at.scanner.EOF() }

func (at *arrayTask) CheckerEOF() bool { return at.checker.EOF() }

func (at *arrayTask) CleanData() {
	at.scanner.RemoveFiles()
	at.checker.RemoveFiles()
}

func (at *arrayTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return at.array1
	case 2:
		return at.array2
	case 3:
		return at.array3
	case 4:
		return at.array4
	case 5:
		return at.array5
	case 6:
		return at.array6
	case 7:
		return at.array7
	case 8:
		return at.array8
	case 9:
		return at.array9
	case 10:
		return at.array10
	case 11:
		return at.array11
	case 12:
		return at.array12
	case 13:
		return at.array13
	case 14:
		return at.array14
	case 15:
		return at.array15
	case 16:
		return at.array16
	case 17:
		return at.array17
	case 18:
		return at.array18
	case 19:
		return at.array19
	case 20:
		return at.array20
	case 21:
		return at.array21
	case 22:
		return at.array22
	case 23:
		return at.array23
	case 24:
		return at.array24
	case 25:
		return at.array25
	case 26:
		return at.array26
	case 27:
		return at.array27
	case 28:
		return at.array28
	case 29:
		return at.array29
	case 30:
		return at.array30
	case 31:
		return at.array31
	case 32:
		return at.array32
	case 33:
		return at.array33
	case 34:
		return at.array34
	case 35:
		return at.array35
	case 36:
		return at.array36
	case 37:
		return at.array37
	case 38:
		return at.array38
	case 39:
		return at.array39
	case 40:
		return at.array40
	case 41:
		return at.array41
	case 42:
		return at.array42
	case 43:
		return at.array43
	case 44:
		return at.array44
	case 45:
		return at.array45
	case 46:
		return at.array46
	case 47:
		return at.array47
	case 48:
		return at.array48
	case 49:
		return at.array49
	case 50:
		return at.array50
	case 51:
		return at.array51
	case 52:
		return at.array52
	case 53:
		return at.array53
	case 54:
		return at.array54
	case 55:
		return at.array55
	case 56:
		return at.array56
	case 57:
		return at.array57
	case 58:
		return at.array58
	case 59:
		return at.array59
	case 60:
		return at.array60
	case 61:
		return at.array61
	case 62:
		return at.array62
	case 63:
		return at.array63
	case 64:
		return at.array64
	case 65:
		return at.array65
	case 66:
		return at.array66
	case 67:
		return at.array67
	case 68:
		return at.array68
	case 69:
		return at.array69
	case 70:
		return at.array70
	case 71:
		return at.array71
	case 72:
		return at.array72
	case 73:
		return at.array73
	case 74:
		return at.array74
	case 75:
		return at.array75
	case 76:
		return at.array76
	case 77:
		return at.array77
	case 78:
		return at.array78
	case 79:
		return at.array79
	case 80:
		return at.array80
	case 81:
		return at.array81
	case 82:
		return at.array82
	case 83:
		return at.array83
	case 84:
		return at.array84
	case 85:
		return at.array85
	case 86:
		return at.array86
	case 87:
		return at.array87
	case 88:
		return at.array88
	case 89:
		return at.array89
	case 90:
		return at.array90
	case 91:
		return at.array91
	case 92:
		return at.array92
	case 93:
		return at.array93
	case 94:
		return at.array94
	case 95:
		return at.array95
	case 96:
		return at.array96
	case 97:
		return at.array97
	case 98:
		return at.array98
	case 99:
		return at.array99
	case 100:
		return at.array100
	case 101:
		return at.array101
	case 102:
		return at.array102
	case 103:
		return at.array103
	case 104:
		return at.array104
	case 105:
		return at.array105
	case 106:
		return at.array106
	case 107:
		return at.array107
	case 108:
		return at.array108
	case 109:
		return at.array109
	case 110:
		return at.array110
	case 111:
		return at.array111
	case 112:
		return at.array112
	case 113:
		return at.array113
	case 114:
		return at.array114
	case 115:
		return at.array115
	case 116:
		return at.array116
	case 117:
		return at.array117
	case 118:
		return at.array118
	case 119:
		return at.array119
	case 120:
		return at.array120
	case 121:
		return at.array121
	case 122:
		return at.array122
	case 123:
		return at.array123
	case 124:
		return at.array124
	case 125:
		return at.array125
	case 126:
		return at.array126
	case 127:
		return at.array127
	case 128:
		return at.array128
	case 129:
		return at.array129
	case 130:
		return at.array130
	case 131:
		return at.array131
	case 132:
		return at.array132
	case 133:
		return at.array133
	case 134:
		return at.array134
	case 135:
		return at.array135
	case 136:
		return at.array136
	case 137:
		return at.array137
	case 138:
		return at.array138
	case 139:
		return at.array139
	case 140:
		return at.array140
	default:
		return nil
	}
}

// { "no": 1, "dat": "", "ans": "" }
func (at *arrayTask) array1() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	array[0] = 1
	for index := 1; index < n; index++ {
		array[index] = array[index-1] + 2
	}
	return at.checker.CompareInt(array...)
}

// { "no": 2, "dat": "", "ans": "" }
func (at *arrayTask) array2() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	array[0] = 2
	for index := 1; index < n; index++ {
		array[index] = array[index-1] * 2
	}
	return at.checker.CompareInt(array...)
}

// { "no": 3, "dat": "2", "ans": "2" }
func (at *arrayTask) array3() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	var a = at.scanner.NextFloat64()
	var d = at.scanner.NextFloat64()
	array[0] = a
	for index := 1; index < n; index++ {
		array[index] = array[index-1] + d
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 4, "dat": "2", "ans": "2" }
func (at *arrayTask) array4() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	var a = at.scanner.NextFloat64()
	var r = at.scanner.NextFloat64()
	array[0] = a
	for index := 1; index < n; index++ {
		array[index] = array[index-1] * r
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 5, "dat": "", "ans": "" }
func (at *arrayTask) array5() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	array[0], array[1] = 1, 1
	for index := 2; index < n; index++ {
		array[index] = array[index-1] + array[index-2]
	}
	return at.checker.CompareInt(array...)
}

// { "no": 6, "dat": "", "ans": "" }
func (at *arrayTask) array6() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	var a = at.scanner.NextInt()
	var b = at.scanner.NextInt()
	array[0], array[1], array[2] = a, b, a+b
	for index := 3; index < n; index++ {
		array[index] = array[index-1] * 2
	}
	return at.checker.CompareInt(array...)
}

// { "no": 7, "dat": "2", "ans": "2" }
func (at *arrayTask) array7() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := n - 1; index >= 0; index-- {
		if !at.checker.CompareFloat64(2, array[index]) {
			return false
		}
	}
	return true
}

// { "no": 8, "dat": "", "ans": "" }
func (at *arrayTask) array8() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	var oddsCount = 0
	for index := 0; index < n; index++ {
		if array[index]%2 != 0 {
			if !at.checker.CompareInt(array[index]) {
				return false
			}
			oddsCount++
		}
	}
	return at.checker.CompareInt(oddsCount)
}

// { "no": 9, "dat": "", "ans": "" }
func (at *arrayTask) array9() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	var evensCount = 0
	for index := n - 1; index >= 0; index-- {
		if array[index]%2 == 0 {
			if !at.checker.CompareInt(array[index]) {
				return false
			}
			evensCount++
		}
	}
	return at.checker.CompareInt(evensCount)
}

// { "no": 10, "dat": "", "ans": "" }
func (at *arrayTask) array10() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	// output all odds;
	for index := 0; index < n; index++ {
		if array[index]%2 == 0 {
			if !at.checker.CompareInt(array[index]) {
				return false
			}
		}
	}
	// output all evens;
	for index := n - 1; index >= 0; index-- {
		if array[index]%2 != 0 {
			if !at.checker.CompareInt(array[index]) {
				return false
			}
		}
	}
	return true
}

// { "no": 11, "dat": "2", "ans": "2" }
func (at *arrayTask) array11() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	for index := k - 1; index < n; index += k {
		if !at.checker.CompareFloat64(2, array[index]) {
			return false
		}
	}
	return true
}

// { "no": 12, "dat": "2", "ans": "2" }
func (at *arrayTask) array12() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := 1; index < n; index += 2 {
		if !at.checker.CompareFloat64(2, array[index]) {
			return false
		}
	}
	return true
}

// { "no": 13, "dat": "2", "ans": "2" }
func (at *arrayTask) array13() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := n - 1; index >= 0; index -= 2 {
		if !at.checker.CompareFloat64(2, array[index]) {
			return false
		}
	}
	return true
}

// { "no": 14, "dat": "2", "ans": "2" }
func (at *arrayTask) array14() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := 1; index < n; index += 2 {
		if !at.checker.CompareFloat64(2, array[index]) {
			return false
		}
	}
	for index := 0; index < n; index += 2 {
		if !at.checker.CompareFloat64(2, array[index]) {
			return false
		}
	}
	return true
}

// { "no": 15, "dat": "2", "ans": "2" }
func (at *arrayTask) array15() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := 0; index < n; index += 2 {
		if !at.checker.CompareFloat64(2, array[index]) {
			return false
		}
	}
	for index := n - 1 - n%2; index >= 0; index -= 2 {
		if !at.checker.CompareFloat64(2, array[index]) {
			return false
		}
	}
	return true
}

// { "no": 16, "dat": "2", "ans": "2" }
func (at *arrayTask) array16() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var limit = n/2 + n%2
	for index := 0; index < limit; index++ {
		if !at.checker.CompareFloat64(2, array[index]) {
			return false
		}
		if n-index-1 > index {
			if !at.checker.CompareFloat64(2, array[n-index-1]) {
				return false
			}
		}
	}
	return true
}

// { "no": 17, "dat": "2", "ans": "2" }
func (at *arrayTask) array17() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var index = 0
	var limit = n - 1
LOOP:
	for index <= limit {
		for i := 0; i < 2; i++ {
			if !at.checker.CompareFloat64(2, array[index]) {
				return false
			}
			index++
			if index > limit {
				break LOOP
			}
		}
		for i := 0; i < 2; i++ {
			if !at.checker.CompareFloat64(2, array[limit]) {
				return false
			}
			limit--
			if index > limit {
				break LOOP
			}
		}
	}
	return true
}

// { "no": 18, "dat": "", "ans": "" }
func (at *arrayTask) array18() bool {
	var array [10]int
	for index := 0; index < 10; index++ {
		array[index] = at.scanner.NextInt()
	}
	var result = 0
	for index := 0; index < 9; index++ {
		if array[index] < array[9] {
			result = array[index]
			break
		}
	}
	return at.checker.CompareInt(result)
}

// { "no": 19, "dat": "", "ans": "" }
func (at *arrayTask) array19() bool {
	var array [10]int
	for index := 0; index < 10; index++ {
		array[index] = at.scanner.NextInt()
	}
	var resultIndex = 0
	for index := 1; index < 9; index++ {
		if array[0] < array[index] && array[index] < array[9] {
			resultIndex = index + 1
		}
	}
	return at.checker.CompareInt(resultIndex)
}

// { "no": 20, "dat": "2", "ans": "2" }
func (at *arrayTask) array20() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	var l = at.scanner.NextInt()
	var sum = 0.0
	for index := k - 1; index < l; index++ {
		sum += array[index]
	}
	return at.checker.CompareFloat64(2, sum)
}

// { "no": 21, "dat": "2", "ans": "2" }
func (at *arrayTask) array21() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	var l = at.scanner.NextInt()
	var sum = 0.0
	for index := k - 1; index < l; index++ {
		sum += array[index]
	}
	var average = sum / float64(l-k+1)
	return at.checker.CompareFloat64(2, average)
}

// { "no": 22, "dat": "2", "ans": "2" }
func (at *arrayTask) array22() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	var sum = 0.0
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
		sum += array[index]
	}
	var k = at.scanner.NextInt()
	var l = at.scanner.NextInt()
	for index := k - 1; index < l; index++ {
		sum -= array[index]
	}
	return at.checker.CompareFloat64(2, sum)
}

// { "no": 23, "dat": "2", "ans": "2" }
func (at *arrayTask) array23() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	var sum = 0.0
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
		sum += array[index]
	}
	var k = at.scanner.NextInt()
	var l = at.scanner.NextInt()
	for index := k - 1; index < l; index++ {
		sum -= array[index]
	}
	var average = sum / float64(n-(l-k+1))
	return at.checker.CompareFloat64(2, average)
}

// { "no": 24, "dat": "", "ans": "" }
func (at *arrayTask) array24() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	var difference = 0
	for index := 1; index < n; index++ {
		if index == 1 {
			difference = array[index] - array[index-1]
		} else if array[index]-array[index-1] != difference {
			difference = 0
			break
		}
	}
	return at.checker.CompareInt(difference)
}

// { "no": 25, "dat": "", "ans": "" }
func (at *arrayTask) array25() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	var ratio = 0
	for index := 1; index < n; index++ {
		if index == 1 {
			ratio = array[index] / array[index-1]
		} else if array[index]/array[index-1] != ratio || array[index]%array[index-1] != 0 {
			ratio = 0
			break
		}
	}
	return at.checker.CompareInt(ratio)
}

// { "no": 26, "dat": "", "ans": "" }
func (at *arrayTask) array26() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	var errIndex = 0
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if errIndex == 0 && index > 0 && (array[index-1]+array[index])%2 == 0 {
			errIndex = index + 1
		}
	}
	return at.checker.CompareInt(errIndex)
}

// { "no": 27, "dat": "", "ans": "" }
func (at *arrayTask) array27() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	var errIndex = 0
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if errIndex == 0 && index > 0 && array[index-1]*array[index] > 0 {
			errIndex = index + 1
		}
	}
	return at.checker.CompareInt(errIndex)
}

// { "no": 28, "dat": "2", "ans": "2" }
func (at *arrayTask) array28() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var minimum = 0.0
	for index := 1; index < n; index += 2 {
		if index == 1 {
			minimum = array[index]
		} else if array[index] < minimum {
			minimum = array[index]
		}
	}
	return at.checker.CompareFloat64(2, minimum)
}

// { "no": 29, "dat": "2", "ans": "2" }
func (at *arrayTask) array29() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var maximum = 0.0
	for index := 0; index < n; index += 2 {
		if index == 0 {
			maximum = array[index]
		} else if array[index] > maximum {
			maximum = array[index]
		}
	}
	return at.checker.CompareFloat64(2, maximum)
}

// { "no": 30, "dat": "2", "ans": "" }
func (at *arrayTask) array30() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var count = 0
	for index := 0; index < n-1; index++ {
		if array[index] > array[index+1] {
			count++
			if !at.checker.CompareInt(index + 1) {
				return false
			}
		}
	}
	return at.checker.CompareInt(count)
}

// { "no": 31, "dat": "2", "ans": "" }
func (at *arrayTask) array31() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var count = 0
	for index := n - 1; index > 0; index-- {
		if array[index] > array[index-1] {
			count++
			if !at.checker.CompareInt(index + 1) {
				return false
			}
		}
	}
	return at.checker.CompareInt(count)
}

// { "no": 32, "dat": "2", "ans": "" }
func (at *arrayTask) array32() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	var localMinimumIndex = -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := 0; index < n; index++ {
		if index == 0 {
			if n > 1 && array[index] < array[index+1] {
				localMinimumIndex = index
				break
			}
		} else if index == n-1 {
			if n > 1 && array[index] < array[index-1] {
				localMinimumIndex = index
				break
			}
		} else if array[index] < array[index-1] && array[index] < array[index+1] {
			localMinimumIndex = index
			break
		}
	}
	localMinimumIndex++
	return at.checker.CompareInt(localMinimumIndex)
}

// { "no": 33, "dat": "2", "ans": "" }
func (at *arrayTask) array33() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var localMaximumIndex = -1
	for index := 0; index < n; index++ {
		if index == 0 {
			if n > 1 && array[index] > array[index+1] {
				localMaximumIndex = index
			}
		} else if index == n-1 {
			if n > 1 && array[index] > array[index-1] {
				localMaximumIndex = index
			}
		} else if array[index] > array[index-1] && array[index] > array[index+1] {
			localMaximumIndex = index
		}
	}
	localMaximumIndex++
	return at.checker.CompareInt(localMaximumIndex)
}

// { "no": 34, "dat": "2", "ans": "2" }
func (at *arrayTask) array34() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var maximumIndex = -1
	for index := 0; index < n; index++ {
		if index == 0 {
			if n > 1 && array[index] < array[index+1] {
				if maximumIndex < 0 || array[index] > array[maximumIndex] {
					maximumIndex = index
				}
			}
		} else if index == n-1 {
			if n > 1 && array[index] < array[index-1] {
				if maximumIndex < 0 || array[index] > array[maximumIndex] {
					maximumIndex = index
				}
			}
		} else if array[index] < array[index-1] && array[index] < array[index+1] {
			if maximumIndex < 0 || array[index] > array[maximumIndex] {
				maximumIndex = index
			}
		}
	}
	var maximum float64
	if maximumIndex >= 0 {
		maximum = array[maximumIndex]
	}
	return at.checker.CompareFloat64(2, maximum)
}

// { "no": 35, "dat": "2", "ans": "2" }
func (at *arrayTask) array35() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var minimumIndex = -1
	for index := 0; index < n; index++ {
		if index == 0 {
			if n > 1 && array[index] > array[index+1] {
				if minimumIndex < 0 || array[index] < array[minimumIndex] {
					minimumIndex = index
				}
			}
		} else if index == n-1 {
			if n > 1 && array[index] > array[index-1] {
				if minimumIndex < 0 || array[index] < array[minimumIndex] {
					minimumIndex = index
				}
			}
		} else if array[index] > array[index-1] && array[index] > array[index+1] {
			if minimumIndex < 0 || array[index] < array[minimumIndex] {
				minimumIndex = index
			}
		}
	}
	var minimum float64
	if minimumIndex >= 0 {
		minimum = array[minimumIndex]
	}
	return at.checker.CompareFloat64(2, minimum)
}

// { "no": 36, "dat": "2", "ans": "2" }
func (at *arrayTask) array36() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var maximumIndex = -1
	for index := 1; index < n-1; index++ {
		if array[index-1] < array[index] && array[index] < array[index+1] ||
			array[index-1] > array[index] && array[index] > array[index+1] {
			if maximumIndex < 0 || array[index] > array[maximumIndex] {
				maximumIndex = index
			}
		}
	}
	var maximum float64
	if maximumIndex >= 0 {
		maximum = array[maximumIndex]
	}
	return at.checker.CompareFloat64(2, maximum)
}

// { "no": 37, "dat": "2", "ans": "" }
func (at *arrayTask) array37() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var count = 0
	var isProgress = false
	for index := 1; index < n; index++ {
		if array[index-1] < array[index] {
			isProgress = true
		} else if isProgress {
			count++
			isProgress = false
		}
	}
	if isProgress {
		count++
		isProgress = false
	}
	return at.checker.CompareInt(count)
}

// { "no": 38, "dat": "2", "ans": "" }
func (at *arrayTask) array38() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var count = 0
	var isRegress = false
	for index := 1; index < n; index++ {
		if array[index-1] > array[index] {
			isRegress = true
		} else if isRegress {
			count++
			isRegress = false
		}
	}
	if isRegress {
		count++
		isRegress = false
	}
	return at.checker.CompareInt(count)
}

// { "no": 39, "dat": "2", "ans": "" }
func (at *arrayTask) array39() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var count = 0
	var isProgress, isRegress = false, false
	for index := 1; index < n; index++ {
		if array[index-1] < array[index] {
			isProgress = true
		} else if isProgress {
			count++
			isProgress = false
		}
		if array[index-1] > array[index] {
			isRegress = true
		} else if isRegress {
			count++
			isRegress = false
		}
	}
	if isProgress || isRegress {
		count++
	}
	return at.checker.CompareInt(count)
}

// { "no": 40, "dat": "2", "ans": "2" }
func (at *arrayTask) array40() bool {
	var r = at.scanner.NextFloat64()
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	var distance, minimalDistance float64
	var minimal float64
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
		distance = math.Abs(array[index] - r)
		if index == 0 || distance < minimalDistance {
			minimalDistance = distance
			minimal = array[index]
		}
	}
	return at.checker.CompareFloat64(2, minimal)
}

// { "no": 41, "dat": "2", "ans": "2" }
func (at *arrayTask) array41() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var sum, maxSum float64
	var sumIndex = -1
	for index := 1; index < n; index++ {
		sum = array[index-1] + array[index]
		if sumIndex < 0 || sum > maxSum {
			maxSum = sum
			sumIndex = index
		}
	}
	var element1, element2 float64
	if sumIndex > 0 {
		element1, element2 = array[sumIndex-1], array[sumIndex]
	}
	return at.checker.CompareFloat64(2, element1, element2)
}

// { "no": 42, "dat": "2", "ans": "2" }
func (at *arrayTask) array42() bool {
	var r = at.scanner.NextFloat64()
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var distance, minDistance float64
	var distanceIndex = -1
	for index := 1; index < n; index++ {
		distance = math.Abs(array[index-1] + array[index] - r)
		if distanceIndex < 0 || distance < minDistance {
			minDistance = distance
			distanceIndex = index
		}
	}
	var element1, element2 float64
	if distanceIndex > 0 {
		element1, element2 = array[distanceIndex-1], array[distanceIndex]
	}
	return at.checker.CompareFloat64(2, element1, element2)
}

// { "no": 43, "dat": "", "ans": "" }
func (at *arrayTask) array43() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	var distinctsCount = 0
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if index == 0 || array[index] != array[index-1] {
			distinctsCount++
		}
	}
	return at.checker.CompareInt(distinctsCount)
}

// { "no": 44, "dat": "", "ans": "" }
func (at *arrayTask) array44() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	var index1 int
	var index2 = -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if index2 < 0 {
			for index1 = index - 1; index1 >= 0; index1-- {
				if array[index1] == array[index] {
					index2 = index
					break
				}
			}
		}
	}
	index1++
	index2++
	return at.checker.CompareInt(index1, index2)
}

// { "no": 45, "dat": "2", "ans": "" }
func (at *arrayTask) array45() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	var index1, index2 = -1, -1
	var distance, minDistance float64
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
		for i := index - 1; i >= 0; i-- {
			distance = math.Abs(array[index] - array[i])
			if index1 < 0 || distance < minDistance {
				minDistance = distance
				index1, index2 = i, index
			}
		}
	}
	index1++
	index2++
	return at.checker.CompareInt(index1, index2)
}

// { "no": 46, "dat": "2", "ans": "2" }
func (at *arrayTask) array46() bool {
	var r = at.scanner.NextFloat64()
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	var distance, minDistance float64
	var index1, index2 = -1, -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
		for i := index - 1; i >= 0; i-- {
			distance = math.Abs(array[i] + array[index] - r)
			if index1 < 0 || distance < minDistance {
				minDistance = distance
				index1, index2 = i, index
			}
		}
	}
	var element1, element2 float64
	if index1 >= 0 && index2 >= 0 {
		element1, element2 = array[index1], array[index2]
	}
	return at.checker.CompareFloat64(2, element1, element2)
}

// { "no": 47, "dat": "", "ans": "" }
func (at *arrayTask) array47() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	var isDistinct bool
	var count = 0
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		isDistinct = true
		for i := index - 1; i >= 0; i-- {
			if array[i] == array[index] {
				isDistinct = false
				break
			}
		}
		if isDistinct {
			count++
		}
	}
	return at.checker.CompareInt(count)
}

// { "no": 48, "dat": "", "ans": "" }
func (at *arrayTask) array48() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	var count int
	var maxCount = 0
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		count = 1
		for i := index - 1; i >= 0; i-- {
			if array[i] == array[index] {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return at.checker.CompareInt(maxCount)
}

// { "no": 49, "dat": "", "ans": "" }
func (at *arrayTask) array49() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	var errIndex = -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if errIndex < 0 {
			if array[index] < 1 || array[index] > n {
				errIndex = index
			} else {
				for i := index - 1; i >= 0; i-- {
					if array[i] == array[index] {
						errIndex = index
						break
					}
				}
			}
		}
	}
	errIndex++
	return at.checker.CompareInt(errIndex)
}

// { "no": 50, "dat": "", "ans": "" }
func (at *arrayTask) array50() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	var count = 0
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		for i := index - 1; i >= 0; i-- {
			if array[i] > array[index] {
				count++
			}
		}
	}
	return at.checker.CompareInt(count)
}

// { "no": 51, "dat": "2", "ans": "2" }
func (at *arrayTask) array51() bool {
	var n = at.scanner.NextInt()
	var a = make([]float64, n)
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextFloat64()
	}
	var b = make([]float64, n)
	for index := 0; index < n; index++ {
		b[index] = at.scanner.NextFloat64()
	}
	// a, b = b, a
	for index := 0; index < n; index++ {
		a[index], b[index] = b[index], a[index]
	}
	if !at.checker.CompareFloat64(2, a...) {
		return false
	}
	return at.checker.CompareFloat64(2, b...)
}

// { "no": 52, "dat": "2", "ans": "2" }
func (at *arrayTask) array52() bool {
	var n = at.scanner.NextInt()
	var a = make([]float64, n)
	var b = make([]float64, n)
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextFloat64()
		if a[index] < 5 {
			b[index] = 2 * a[index]
		} else {
			b[index] = a[index] / 2
		}
	}
	return at.checker.CompareFloat64(2, b...)
}

// { "no": 53, "dat": "2", "ans": "2" }
func (at *arrayTask) array53() bool {
	var n = at.scanner.NextInt()
	var a = make([]float64, n)
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextFloat64()
	}
	var b = make([]float64, n)
	for index := 0; index < n; index++ {
		b[index] = at.scanner.NextFloat64()
	}
	var c = make([]float64, n)
	for index := 0; index < n; index++ {
		if a[index] > b[index] {
			c[index] = a[index]
		} else {
			c[index] = b[index]
		}
	}
	return at.checker.CompareFloat64(2, c...)
}

// { "no": 54, "dat": "", "ans": "" }
func (at *arrayTask) array54() bool {
	var n = at.scanner.NextInt()
	var a = make([]int, n)
	var b = make([]int, 0)
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextInt()
		if a[index]%2 == 0 {
			b = append(b, a[index])
		}
	}
	if !at.checker.CompareInt(len(b)) {
		return false
	}
	return at.checker.CompareInt(b...)
}

// { "no": 55, "dat": "", "ans": "" }
func (at *arrayTask) array55() bool {
	var n = at.scanner.NextInt()
	var a = make([]int, n)
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextInt()
	}
	var bSize = n/2 + n%2
	var b = make([]int, bSize)
	for index := 0; index < bSize; index++ {
		b[index] = a[2*index]
	}
	if !at.checker.CompareInt(bSize) {
		return false
	}
	return at.checker.CompareInt(b...)
}

// { "no": 56, "dat": "", "ans": "" }
func (at *arrayTask) array56() bool {
	var n = at.scanner.NextInt()
	var a = make([]int, n)
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextInt()
	}
	var bSize = n / 3
	var b = make([]int, bSize)
	for i, j := 0, 2; j < n; i, j = i+1, j+3 {
		b[i] = a[j]
	}
	if !at.checker.CompareInt(bSize) {
		return false
	}
	return at.checker.CompareInt(b...)
}

// { "no": 57, "dat": "", "ans": "" }
func (at *arrayTask) array57() bool {
	var n = at.scanner.NextInt()
	var a = make([]int, n)
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextInt()
	}
	var b = make([]int, n)
	index := 0
	for i := 1; i < n; i += 2 {
		b[index] = a[i]
		index++
	}
	for i := 0; i < n; i += 2 {
		b[index] = a[i]
		index++
	}
	return at.checker.CompareInt(b...)
}

// { "no": 58, "dat": "2", "ans": "2" }
func (at *arrayTask) array58() bool {
	var n = at.scanner.NextInt()
	var a = make([]float64, n)
	var b = make([]float64, n)
	var sum = 0.0
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextFloat64()
		sum += a[index]
		b[index] = sum
	}
	return at.checker.CompareFloat64(2, b...)
}

// { "no": 59, "dat": "2", "ans": "2" }
func (at *arrayTask) array59() bool {
	var n = at.scanner.NextInt()
	var a = make([]float64, n)
	var b = make([]float64, n)
	var sum = 0.0
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextFloat64()
		sum += a[index]
		b[index] = sum / float64(index+1)
	}
	return at.checker.CompareFloat64(2, b...)
}

// { "no": 60, "dat": "2", "ans": "2" }
func (at *arrayTask) array60() bool {
	var n = at.scanner.NextInt()
	var a = make([]float64, n)
	var sum = 0.0
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextFloat64()
		sum += a[index]
	}
	var b = make([]float64, n)
	for index := 0; index < n; index++ {
		b[index] = sum
		sum -= a[index]
	}
	return at.checker.CompareFloat64(2, b...)
}

// { "no": 61, "dat": "2", "ans": "2" }
func (at *arrayTask) array61() bool {
	var n = at.scanner.NextInt()
	var a = make([]float64, n)
	var sum = 0.0
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextFloat64()
		sum += a[index]
	}
	var b = make([]float64, n)
	for index := 0; index < n; index++ {
		b[index] = sum / float64(n-index)
		sum -= a[index]
	}
	return at.checker.CompareFloat64(2, b...)
}

// { "no": 62, "dat": "2", "ans": "2" }
func (at *arrayTask) array62() bool {
	var n = at.scanner.NextInt()
	var a = make([]float64, n)
	var b = make([]float64, 0)
	var c = make([]float64, 0)
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextFloat64()
		if a[index] >= 0 {
			b = append(b, a[index])
		} else {
			c = append(c, a[index])
		}
	}
	if !at.checker.CompareInt(len(b)) {
		return false
	}
	if !at.checker.CompareFloat64(2, b...) {
		return false
	}
	if !at.checker.CompareInt(len(c)) {
		return false
	}
	return at.checker.CompareFloat64(2, c...)
}

// { "no": 63, "dat": "2", "ans": "2" }
func (at *arrayTask) array63() bool {
	var a, b [5]float64
	var c = make([]float64, 0)
	for index := 0; index < 5; index++ {
		a[index] = at.scanner.NextFloat64()
		c = append(c, a[index])
	}
	for index := 0; index < 5; index++ {
		b[index] = at.scanner.NextFloat64()
		c = append(c, b[index])
	}
	sort.Float64s(c)
	return at.checker.CompareFloat64(2, c...)
}

// { "no": 64, "dat": "", "ans": "" }
func (at *arrayTask) array64() bool {
	var nA = at.scanner.NextInt()
	var d = make([]int, 0)
	var a = make([]int, nA)
	for index := 0; index < nA; index++ {
		a[index] = at.scanner.NextInt()
	}
	d = append(d, a...)
	var nB = at.scanner.NextInt()
	var b = make([]int, nB)
	for index := 0; index < nB; index++ {
		b[index] = at.scanner.NextInt()
	}
	d = append(d, b...)
	var nC = at.scanner.NextInt()
	var c = make([]int, nC)
	for index := 0; index < nC; index++ {
		c[index] = at.scanner.NextInt()
	}
	d = append(d, c...)
	var nD = nA + nB + nC
	sort.Ints(d)
	for index := 0; index < nD/2; index++ {
		d[index], d[nD-index-1] = d[nD-index-1], d[index]
	}
	return at.checker.CompareInt(d...)
}

// { "no": 65, "dat": "2", "ans": "2" }
func (at *arrayTask) array65() bool {
	var n = at.scanner.NextInt()
	var a = make([]float64, n)
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	for index := 0; index < n; index++ {
		if index == k-1 {
			continue
		}
		a[index] += a[k-1]
	}
	if k >= 1 && k <= n {
		a[k-1] += a[k-1]
	}
	return at.checker.CompareFloat64(2, a...)
}

// { "no": 66, "dat": "", "ans": "" }
func (at *arrayTask) array66() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	var evenIndex = -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if evenIndex < 0 && array[index]%2 == 0 {
			evenIndex = index
		}
	}
	if evenIndex >= 0 {
		for index := 0; index < n; index++ {
			if index == evenIndex {
				continue
			}
			if array[index]%2 == 0 {
				array[index] += array[evenIndex]
			}
		}
		array[evenIndex] += array[evenIndex]
	}
	return at.checker.CompareInt(array...)
}

// { "no": 67, "dat": "", "ans": "" }
func (at *arrayTask) array67() bool {
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	var oddIndex = -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if array[index]%2 != 0 {
			oddIndex = index
		}
	}
	if oddIndex >= 0 {
		for index := 0; index < n; index++ {
			if index == oddIndex {
				continue
			}
			if array[index]%2 != 0 {
				array[index] += array[oddIndex]
			}
		}
		array[oddIndex] += array[oddIndex]
	}
	return at.checker.CompareInt(array...)
}

// { "no": 68, "dat": "2", "ans": "2" }
func (at *arrayTask) array68() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	var minimalIndex, maximalIndex = -1, -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
		if minimalIndex < 0 || array[index] <= array[minimalIndex] {
			minimalIndex = index
		}
		if maximalIndex < 0 || array[index] >= array[maximalIndex] {
			maximalIndex = index
		}
	}
	if minimalIndex >= 0 && maximalIndex >= 0 {
		array[minimalIndex], array[maximalIndex] = array[maximalIndex], array[minimalIndex]
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 69, "dat": "2", "ans": "2" }
func (at *arrayTask) array69() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := 1; index < n; index += 2 {
		array[index-1], array[index] = array[index], array[index-1]
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 70, "dat": "2", "ans": "2" }
func (at *arrayTask) array70() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var halfN = n / 2
	for index := 0; index < halfN; index++ {
		array[index], array[halfN+index] = array[halfN+index], array[index]
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 71, "dat": "2", "ans": "2" }
func (at *arrayTask) array71() bool {
	var n = at.scanner.NextInt()
	var array = array.NewFloat64(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	array.Reverse()
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 72, "dat": "2", "ans": "2" }
func (at *arrayTask) array72() bool {
	var n = at.scanner.NextInt()
	var a = make([]float64, n)
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	var l = at.scanner.NextInt()
	k--
	l--
	var limit = (l - k + 1) / 2
	for index := 0; index < limit; index++ {
		a[k], a[l] = a[l], a[k]
		k++
		l--
	}
	return at.checker.CompareFloat64(2, a...)
}

// { "no": 73, "dat": "2", "ans": "2" }
func (at *arrayTask) array73() bool {
	var n = at.scanner.NextInt()
	var a = make([]float64, n)
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	var l = at.scanner.NextInt()
	l -= 2
	var limit = (l - k + 1) / 2
	for index := 0; index < limit; index++ {
		a[k], a[l] = a[l], a[k]
		k++
		l--
	}
	return at.checker.CompareFloat64(2, a...)
}

// { "no": 74, "dat": "2", "ans": "2" }
func (at *arrayTask) array74() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	var minimalIndex, maximalIndex = -1, -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
		if minimalIndex < 0 || array[index] < array[minimalIndex] {
			minimalIndex = index
		}
		if maximalIndex < 0 || array[index] > array[maximalIndex] {
			maximalIndex = index
		}
	}
	if minimalIndex > maximalIndex {
		minimalIndex, maximalIndex = maximalIndex, minimalIndex
	}
	for index := minimalIndex + 1; index < maximalIndex; index++ {
		array[index] = 0
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 75, "dat": "2", "ans": "2" }
func (at *arrayTask) array75() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	var minimalIndex, maximalIndex = -1, -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
		if minimalIndex < 0 || array[index] < array[minimalIndex] {
			minimalIndex = index
		}
		if maximalIndex < 0 || array[index] > array[maximalIndex] {
			maximalIndex = index
		}
	}
	if minimalIndex > maximalIndex {
		minimalIndex, maximalIndex = maximalIndex, minimalIndex
	}
	var limit = (maximalIndex - minimalIndex + 1) / 2
	for index := 0; index < limit; index++ {
		array[minimalIndex], array[maximalIndex] = array[maximalIndex], array[minimalIndex]
		minimalIndex++
		maximalIndex--
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 76, "dat": "2", "ans": "2" }
func (at *arrayTask) array76() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var prevValue float64
	var prevIndex = -1
	for index := 0; index < n; index++ {
		if index == 0 {
			if n > 1 && array[index] > array[index+1] {
				prevValue = array[index]
				prevIndex = index
				array[index] = 0
			}
		} else if index == n-1 {
			if n > 1 && array[index] > prevValue {
				array[index] = 0
			}
		} else if array[index] > prevValue && array[index] > array[index+1] {
			prevValue = array[index]
			prevIndex = index
			array[index] = 0
		}
		if prevIndex != index {
			prevValue = array[index]
		}
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 77, "dat": "2", "ans": "2" }
func (at *arrayTask) array77() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var prevValue float64
	var prevIndex = -1
	for index := 0; index < n; index++ {
		if index == 0 {
			if n > 1 && array[index] < array[index+1] {
				prevValue = array[index]
				prevIndex = index
				array[index] *= array[index]
			}
		} else if index == n-1 {
			if n > 1 && array[index] < prevValue {
				array[index] *= array[index]
			}
		} else if array[index] < prevValue && array[index] < array[index+1] {
			prevValue = array[index]
			prevIndex = index
			array[index] *= array[index]
		}
		if prevIndex != index {
			prevValue = array[index]
		}
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 78, "dat": "2", "ans": "2" }
func (at *arrayTask) array78() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var prevValue, sum float64
	for index := 0; index < n; index++ {
		if index == 0 {
			if n > 1 {
				sum = array[index] + array[index+1]
				prevValue = array[index]
				array[index] = sum / 2
			}
		} else if index == n-1 {
			if n > 1 {
				sum = array[index] + prevValue
				array[index] = sum / 2
			}
		} else {
			sum = prevValue + array[index] + array[index+1]
			prevValue = array[index]
			array[index] = sum / 3
		}
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 79, "dat": "2", "ans": "2" }
func (at *arrayTask) array79() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := n - 1; index > 0; index-- {
		array[index] = array[index-1]
	}
	if n > 0 {
		array[0] = 0
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 80, "dat": "2", "ans": "2" }
func (at *arrayTask) array80() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := 0; index < n-1; index++ {
		array[index] = array[index+1]
	}
	if n > 0 {
		array[n-1] = 0
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 81, "dat": "2", "ans": "2" }
func (at *arrayTask) array81() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	for index := n - 1; index >= k; index-- {
		array[index] = array[index-k]
	}
	for index := 0; index < k; index++ {
		array[index] = 0
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 82, "dat": "2", "ans": "2" }
func (at *arrayTask) array82() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	for index := 0; index < n-k; index++ {
		array[index] = array[index+k]
	}
	for index := n - k; index < n; index++ {
		array[index] = 0
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 83, "dat": "2", "ans": "2" }
func (at *arrayTask) array83() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var lastValue float64
	if n > 0 {
		lastValue = array[n-1]
	}
	for index := n - 1; index > 0; index-- {
		array[index] = array[index-1]
	}
	if n > 0 {
		array[0] = lastValue
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 84, "dat": "2", "ans": "2" }
func (at *arrayTask) array84() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var firstValue float64
	if n > 0 {
		firstValue = array[0]
	}
	for index := 0; index < n-1; index++ {
		array[index] = array[index+1]
	}
	if n > 0 {
		array[n-1] = firstValue
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 85, "dat": "2", "ans": "2" }
func (at *arrayTask) array85() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	var tempArray = make([]float64, k)
	for index := n - k; index < n; index++ {
		tempArray[index-n+k] = array[index]
	}
	for index := n - 1; index >= k; index-- {
		array[index] = array[index-k]
	}
	for index := 0; index < k; index++ {
		array[index] = tempArray[index]
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 86, "dat": "2", "ans": "2" }
func (at *arrayTask) array86() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	var tempArray = make([]float64, k)
	for index := 0; index < k; index++ {
		tempArray[index] = array[index]
	}
	for index := 0; index < n-k; index++ {
		array[index] = array[index+k]
	}
	for index := n - k; index < n; index++ {
		array[index] = tempArray[index-n+k]
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 87, "dat": "2", "ans": "2" }
func (at *arrayTask) array87() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := 0; index < n-1; index++ {
		if array[index] > array[index+1] {
			array[index], array[index+1] = array[index+1], array[index]
		} else {
			break
		}
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 88, "dat": "2", "ans": "2" }
func (at *arrayTask) array88() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := n - 2; index >= 0; index-- {
		if array[index] > array[index+1] {
			array[index], array[index+1] = array[index+1], array[index]
		} else {
			break
		}
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 89, "dat": "2", "ans": "2" }
func (at *arrayTask) array89() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	var errIndex = -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
		if errIndex < 0 && index > 0 && array[index-1] < array[index] {
			errIndex = index
		}
	}
	if errIndex >= 0 {
		// move to left from errIndex;
		for index := errIndex; index > 0; index-- {
			if array[index-1] < array[index] {
				array[index-1], array[index] = array[index], array[index-1]
			}
		}
		// move to right from errIndex-1;
		for index := errIndex - 1; index < n-1; index++ {
			if array[index] < array[index+1] {
				array[index], array[index+1] = array[index+1], array[index]
			}
		}
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 90, "dat": "2", "ans": "2" }
func (at *arrayTask) array90() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	if k >= 1 && k <= n {
		array = append(array[0:k-1], array[k:n]...)
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 91, "dat": "2", "ans": "2" }
func (at *arrayTask) array91() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	var l = at.scanner.NextInt()
	if 1 <= k && k < l && l <= n {
		array = append(array[:k-1], array[l:]...)
	}
	if !at.checker.CompareInt(len(array)) {
		return false
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 92, "dat": "", "ans": "" }
func (at *arrayTask) array92() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	for index := 0; index < n; {
		if array[index]%2 != 0 {
			array.DeleteAt(index)
			n--
		} else {
			index++
		}
	}
	if !at.checker.CompareInt(n) {
		return false
	}
	return at.checker.CompareInt(array...)
}

// { "no": 93, "dat": "", "ans": "" }
func (at *arrayTask) array93() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	for index := 1; index < n; index++ {
		array.DeleteAt(index)
		n--
	}
	return at.checker.CompareInt(array...)
}

// { "no": 94, "dat": "", "ans": "" }
func (at *arrayTask) array94() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	for index := 0; index < n; index++ {
		array.DeleteAt(index)
		n--
	}
	return at.checker.CompareInt(array...)
}

// { "no": 95, "dat": "", "ans": "" }
func (at *arrayTask) array95() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	for index := 1; index < n; index++ {
		if array[index] == array[index-1] {
			array.DeleteAt(index)
			n--
			index--
		}
	}
	return at.checker.CompareInt(array...)
}

// { "no": 96, "dat": "", "ans": "" }
func (at *arrayTask) array96() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	var hasEqual bool
	for index := 1; index < n; index++ {
		hasEqual = false
		for i := index - 1; i >= 0; i-- {
			if array[i] == array[index] {
				hasEqual = true
				break
			}
		}
		if hasEqual {
			array.DeleteAt(index)
			n--
			index--
		}
	}
	return at.checker.CompareInt(array...)
}

// { "no": 97, "dat": "", "ans": "" }
func (at *arrayTask) array97() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	var hasEqual bool
	for index := n - 2; index >= 0; index-- {
		hasEqual = false
		for i := index + 1; i < n; i++ {
			if array[i] == array[index] {
				hasEqual = true
				break
			}
		}
		if hasEqual {
			array.DeleteAt(index)
			n--
		}
	}
	return at.checker.CompareInt(array...)
}

// { "no": 98, "dat": "", "ans": "" }
func (at *arrayTask) array98() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	var count = make(map[int]int)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if value, exists := count[array[index]]; exists {
			count[array[index]] = value + 1
		} else {
			count[array[index]] = 1
		}
	}
	for index := 0; index < n; index++ {
		if count[array[index]] < 3 {
			array.DeleteAt(index)
			n--
			index--
		}
	}
	if !at.checker.CompareInt(n) {
		return false
	}
	return at.checker.CompareInt(array...)
}

// { "no": 99, "dat": "", "ans": "" }
func (at *arrayTask) array99() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	var count = make(map[int]int)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if value, exists := count[array[index]]; exists {
			count[array[index]] = value + 1
		} else {
			count[array[index]] = 1
		}
	}
	for index := 0; index < n; index++ {
		if count[array[index]] > 2 {
			array.DeleteAt(index)
			n--
			index--
		}
	}
	if !at.checker.CompareInt(n) {
		return false
	}
	return at.checker.CompareInt(array...)
}

// { "no": 100, "dat": "", "ans": "" }
func (at *arrayTask) array100() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	var count = make(map[int]int)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if value, exists := count[array[index]]; exists {
			count[array[index]] = value + 1
		} else {
			count[array[index]] = 1
		}
	}
	for index := 0; index < n; index++ {
		if count[array[index]] == 2 {
			array.DeleteAt(index)
			n--
			index--
		}
	}
	if !at.checker.CompareInt(n) {
		return false
	}
	return at.checker.CompareInt(array...)
}

// { "no": 101, "dat": "2", "ans": "2" }
func (at *arrayTask) array101() bool {
	var n = at.scanner.NextInt()
	var array = array.NewFloat64(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	if k >= 1 && k <= n {
		array.InsertBefore(k-1, 0)
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 102, "dat": "2", "ans": "2" }
func (at *arrayTask) array102() bool {
	var n = at.scanner.NextInt()
	var array = array.NewFloat64(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	if k >= 1 && k <= n {
		array.InsertAfter(k-1, 0)
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 103, "dat": "2", "ans": "2" }
func (at *arrayTask) array103() bool {
	var n = at.scanner.NextInt()
	var array = array.NewFloat64(n, 0)
	var minimalIndex, maximalIndex = -1, -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
		if minimalIndex < 0 || array[index] <= array[minimalIndex] {
			minimalIndex = index
		}
		if maximalIndex < 0 || array[index] >= array[maximalIndex] {
			maximalIndex = index
		}
	}
	if minimalIndex >= 0 && maximalIndex >= 0 {
		array.InsertBefore(minimalIndex, 0)

		if maximalIndex > minimalIndex {
			maximalIndex++
		}

		array.InsertAfter(maximalIndex, 0)
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 104, "dat": "2", "ans": "2" }
func (at *arrayTask) array104() bool {
	var n = at.scanner.NextInt()
	var arr = array.NewFloat64(n, 0)
	for index := 0; index < n; index++ {
		arr[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	var m = at.scanner.NextInt()
	if k >= 1 && k <= n {
		arr.InsertBefore(k-1, array.NewFloat64(m, 0)...)
	}
	return at.checker.CompareFloat64(2, arr...)
}

// { "no": 105, "dat": "2", "ans": "2" }
func (at *arrayTask) array105() bool {
	var n = at.scanner.NextInt()
	var arr = array.NewFloat64(n, 0)
	for index := 0; index < n; index++ {
		arr[index] = at.scanner.NextFloat64()
	}
	var k = at.scanner.NextInt()
	var m = at.scanner.NextInt()
	if k >= 1 && k <= n {
		arr.InsertAfter(k-1, array.NewFloat64(m, 0)...)
	}
	return at.checker.CompareFloat64(2, arr...)
}

// { "no": 106, "dat": "1", "ans": "1" }
func (at *arrayTask) array106() bool {
	var n = at.scanner.NextInt()
	var array = array.NewFloat64(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := 1; index < n; index += 3 {
		array.InsertAfter(index, array[index])
		n++
	}
	return at.checker.CompareFloat64(1, array...)
}

// { "no": 107, "dat": "1", "ans": "1" }
func (at *arrayTask) array107() bool {
	var n = at.scanner.NextInt()
	var array = array.NewFloat64(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := 0; index < n; index += 4 {
		array.InsertAfter(index, array[index], array[index])
		n += 2
	}
	return at.checker.CompareFloat64(1, array...)
}

// { "no": 108, "dat": "2", "ans": "2" }
func (at *arrayTask) array108() bool {
	var n = at.scanner.NextInt()
	var array = array.NewFloat64(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := 0; index < n; index++ {
		if array[index] >= 0 {
			array.InsertBefore(index, 0)
			index++
			n++
		}
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 109, "dat": "2", "ans": "2" }
func (at *arrayTask) array109() bool {
	var n = at.scanner.NextInt()
	var array = array.NewFloat64(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	for index := 0; index < n; index++ {
		if array[index] < 0 {
			array.InsertAfter(index, 0)
			index++
			n++
		}
	}
	return at.checker.CompareFloat64(2, array...)
}

// { "no": 110, "dat": "", "ans": "" }
func (at *arrayTask) array110() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	for index := 0; index < n; index++ {
		if array[index]%2 == 0 {
			array.InsertAfter(index, array[index])
			index++
			n++
		}
	}
	return at.checker.CompareInt(array...)
}

// { "no": 111, "dat": "", "ans": "" }
func (at *arrayTask) array111() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	for index := 0; index < n; index++ {
		if array[index]%2 != 0 {
			array.InsertAfter(index, array[index], array[index])
			index += 2
			n += 2
		}
	}
	return at.checker.CompareInt(array...)
}

// { "no": 112, "dat": "2", "ans": "2" }
func (at *arrayTask) array112() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	// the exchange sort method (the bubble sorting)
	for i := 0; i < n-1; i++ {
		for index := 1; index < n-i; index++ {
			if array[index-1] > array[index] {
				array[index-1], array[index] = array[index], array[index-1]
			}
		}
		if !at.checker.CompareFloat64(2, array...) {
			return false
		}
	}
	return true
}

// { "no": 113, "dat": "2", "ans": "2" }
func (at *arrayTask) array113() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	// the selection sort method
	var maximalIndex int
	for i := 0; i < n-1; i++ {
		maximalIndex = -1
		for index := n - i - 1; index >= 0; index-- {
			if maximalIndex < 0 || array[index] > array[maximalIndex] {
				maximalIndex = index
			}
		}
		array[n-1-i], array[maximalIndex] = array[maximalIndex], array[n-1-i]
		if !at.checker.CompareFloat64(2, array...) {
			return false
		}
	}
	return true
}

// { "no": 114, "dat": "2", "ans": "2" }
func (at *arrayTask) array114() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextFloat64()
	}
	// the insertion sort method
	for i := 1; i < n; i++ {
		for index := i; index > 0; index-- {
			if array[index-1] > array[index] {
				array[index-1], array[index] = array[index], array[index-1]
			}
		}
		if !at.checker.CompareFloat64(2, array...) {
			return false
		}
	}
	return true
}

// { "no": 115, "dat": "2", "ans": "" }
func (at *arrayTask) array115() bool {
	var n = at.scanner.NextInt()
	var array = make([]float64, n)
	var indexes = make([]int, n)
	for index := 0; index < n; index++ {
		indexes[index] = index
		array[indexes[index]] = at.scanner.NextFloat64()
	}
	for i := 0; i < n-1; i++ {
		for index := 1; index < n-i; index++ {
			if array[indexes[index-1]] > array[indexes[index]] {
				indexes[index-1], indexes[index] = indexes[index], indexes[index-1]
			}
		}
	}
	for index := 0; index < n; index++ {
		indexes[index]++
	}
	return at.checker.CompareInt(indexes...)
}

// { "no": 116, "dat": "", "ans": "" }
func (at *arrayTask) array116() bool {
	var n = at.scanner.NextInt()
	var a = make([]int, n)
	for index := 0; index < n; index++ {
		a[index] = at.scanner.NextInt()
	}
	var b = make([]int, 0)
	var c = make([]int, 0)
	var bIndex = -1
	for index := 0; index < n; index++ {
		if index == 0 || a[index] != a[index-1] {
			bIndex++
			b = append(b, 1)
			c = append(c, a[index])
		} else {
			b[bIndex]++
		}
	}
	if !at.checker.CompareInt(b...) {
		return false
	}
	return at.checker.CompareInt(c...)
}

// { "no": 117, "dat": "", "ans": "" }
func (at *arrayTask) array117() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	for index := 0; index < n; index++ {
		if index == 0 || array[index] != array[index-1] {
			array.InsertBefore(index, 0)
			index++
			n++
		}
	}
	return at.checker.CompareInt(array...)
}

// { "no": 118, "dat": "", "ans": "" }
func (at *arrayTask) array118() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	for index := 0; index < n; index++ {
		if index > 0 && array[index] != array[index-1] {
			array.InsertAfter(index-1, 0)
			index++
			n++
		}
	}
	if n > 1 {
		array.InsertAfter(n-1, 0)
		n++
	}
	return at.checker.CompareInt(array...)
}

// { "no": 119, "dat": "", "ans": "" }
func (at *arrayTask) array119() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	for index := 0; index < n; index++ {
		if index == 0 || array[index] != array[index-1] {
			array.InsertBefore(index, array[index])
			index++
			n++
		}
	}
	return at.checker.CompareInt(array...)
}

// { "no": 120, "dat": "", "ans": "" }
func (at *arrayTask) array120() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	for index := 0; index < n; index++ {
		if index > 0 && array[index] != array[index-1] {
			array.DeleteAt(index - 1)
			index--
			n--
		}
	}
	if n > 1 {
		array.DeleteAt(n - 1)
	}
	return at.checker.CompareInt(array...)
}

// { "no": 121, "dat": "", "ans": "" }
func (at *arrayTask) array121() bool {
	var k = at.scanner.NextInt()
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	var seriesIndex = 0
	for index := 0; index < n; index++ {
		if index == 0 || array[index] != array[index-1] {
			seriesIndex++
		}
		if seriesIndex == k {
			array.InsertBefore(index, array[index])
			index++
			n++
		}
	}
	return at.checker.CompareInt(array...)
}

// { "no": 122, "dat": "", "ans": "" }
func (at *arrayTask) array122() bool {
	var k = at.scanner.NextInt()
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	var seriesIndex = 0
	for index := 0; index < n; index++ {
		if index == 0 || array[index] != array[index-1] {
			seriesIndex++
		}
		if seriesIndex == k {
			for {
				element := array.DeleteAt(index)
				n--
				if index == n || element != array[index] {
					break
				}
			}
			break
		}
	}
	return at.checker.CompareInt(array...)
}

// { "no": 123, "dat": "", "ans": "" }
func (at *arrayTask) array123() bool {
	var k = at.scanner.NextInt()
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	var seriesCount = 0
	var kSerieIndex = -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if index == 0 || array[index] != array[index-1] {
			seriesCount++
		}
		if kSerieIndex < 0 && seriesCount == k {
			kSerieIndex = index
		}
	}
	if kSerieIndex > 0 {
		// k -> 1
		var firstSerieIndex = 0
		var element = array[kSerieIndex]
		for kSerieIndex < n && array[kSerieIndex] == element {
			array.InsertBefore(firstSerieIndex, array.DeleteAt(kSerieIndex))
			firstSerieIndex++
			kSerieIndex++
		}
		kSerieIndex--
		// 1 -> k
		if firstSerieIndex < kSerieIndex && kSerieIndex < n {
			element = array[firstSerieIndex]
			for firstSerieIndex < kSerieIndex && array[firstSerieIndex] == element {
				kSerieIndex--
				array.InsertAfter(kSerieIndex, array.DeleteAt(firstSerieIndex))
			}
		}
	}
	return at.checker.CompareInt(array...)
}

// { "no": 124, "dat": "", "ans": "" }
func (at *arrayTask) array124() bool {
	var k = at.scanner.NextInt()
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	var seriesCount = 0
	var kSerieIndex = -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if index == 0 || array[index] != array[index-1] {
			seriesCount++
		}
		if kSerieIndex < 0 && seriesCount == k {
			kSerieIndex = index
		}
	}
	if kSerieIndex >= 0 {
		// k -> after last;
		var lastSerieLastIndex = n - 1
		var element = array[kSerieIndex]
		for lastSerieLastIndex > kSerieIndex && array[kSerieIndex] == element {
			array.InsertAfter(lastSerieLastIndex-1, array.DeleteAt(kSerieIndex))
			lastSerieLastIndex--
		}
		// last -> k;
		element = array[lastSerieLastIndex]
		for kSerieIndex < lastSerieLastIndex && array[lastSerieLastIndex] == element {
			array.InsertBefore(kSerieIndex, array.DeleteAt(lastSerieLastIndex))
			kSerieIndex++
		}
	}
	return at.checker.CompareInt(array...)
}

// { "no": 125, "dat": "", "ans": "" }
func (at *arrayTask) array125() bool {
	var l = at.scanner.NextInt()
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	var length = 0
	for index := 0; index < n; index++ {
		if index == 0 || array[index] != array[index-1] {
			if length != 0 && length < l {
				array = append(array[:index-length+1], array[index:]...)
				n -= length - 1
				index -= length - 1
				array[index-1] = 0
			}
			length = 0
		}
		length++
	}
	if length != 0 && length < l {
		array = append(array[:n-length], array[n-1:]...)
		n -= length - 1
		array[n-1] = 0
	}
	return at.checker.CompareInt(array...)
}

// { "no": 126, "dat": "", "ans": "" }
func (at *arrayTask) array126() bool {
	var l = at.scanner.NextInt()
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	var length = 0
	for index := 0; index < n; index++ {
		if index == 0 || array[index] != array[index-1] {
			if length != 0 && length == l {
				array = append(array[:index-length+1], array[index:]...)
				n -= length - 1
				index -= length - 1
				array[index-1] = 0
			}
			length = 0
		}
		length++
	}
	if length != 0 && length == l {
		array = append(array[:n-length], array[n-1:]...)
		n -= length - 1
		array[n-1] = 0
	}
	return at.checker.CompareInt(array...)
}

// { "no": 127, "dat": "", "ans": "" }
func (at *arrayTask) array127() bool {
	var l = at.scanner.NextInt()
	var n = at.scanner.NextInt()
	var array = make([]int, n)
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
	}
	var length = 0
	for index := 0; index < n; index++ {
		if index == 0 || array[index] != array[index-1] {
			if length != 0 && length > l {
				array = append(array[:index-length+1], array[index:]...)
				n -= length - 1
				index -= length - 1
				array[index-1] = 0
			}
			length = 0
		}
		length++
	}
	if length != 0 && length > l {
		array = append(array[:n-length], array[n-1:]...)
		n -= length - 1
		array[n-1] = 0
	}
	return at.checker.CompareInt(array...)
}

// { "no": 128, "dat": "", "ans": "" }
func (at *arrayTask) array128() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	var length, maxLength = 0, 0
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if index == 0 || array[index] != array[index-1] {
			if length > 0 && length > maxLength {
				maxLength = length
			}
			length = 0
		}
		length++
	}
	if length > 0 && length > maxLength {
		maxLength = length
	}
	length = 0
	var firstLongSerieIndex = -1
	for index := 0; index < n; index++ {
		if index == 0 || array[index] != array[index-1] {
			if length == maxLength {
				firstLongSerieIndex = index - length
				break
			}
			length = 0
		}
		length++
	}
	if firstLongSerieIndex < 0 {
		firstLongSerieIndex = n - length
	}
	array.InsertBefore(firstLongSerieIndex, array[firstLongSerieIndex])
	return at.checker.CompareInt(array...)
}

// { "no": 129, "dat": "", "ans": "" }
func (at *arrayTask) array129() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	var length, maxLength = 0, 0
	var lastLongSerieIndex = -1
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if index == 0 || array[index] != array[index-1] {
			if length > 0 && length >= maxLength {
				lastLongSerieIndex = index - length
				maxLength = length
			}
			length = 0
		}
		length++
	}
	if length > 0 && length >= maxLength {
		lastLongSerieIndex = n - length
	}
	if lastLongSerieIndex >= 0 {
		array.InsertBefore(lastLongSerieIndex, array[lastLongSerieIndex])
	}
	return at.checker.CompareInt(array...)
}

// { "no": 130, "dat": "", "ans": "" }
func (at *arrayTask) array130() bool {
	var n = at.scanner.NextInt()
	var array = array.NewInt(n, 0)
	var length, maxLength = 0, 0
	for index := 0; index < n; index++ {
		array[index] = at.scanner.NextInt()
		if index == 0 || array[index] != array[index-1] {
			if length > 0 && length > maxLength {
				maxLength = length
			}
			length = 0
		}
		length++
	}
	if length > 0 && length > maxLength {
		maxLength = length
	}
	length = 0
	for index := 0; index < n; index++ {
		if index == 0 || array[index] != array[index-1] {
			if length == maxLength {
				array.InsertBefore(index-1, array[index-1])
				index++
				n++
			}
			length = 0
		}
		length++
	}
	if length == maxLength {
		array.InsertBefore(n-1, array[n-1])
		n++
	}
	return at.checker.CompareInt(array...)
}

// { "no": 131, "dat": "2", "ans": "2" }
func (at *arrayTask) array131() bool {
	var n = at.scanner.NextInt()
	var a = make([]point.Float64, n)
	for index := 0; index < n; index++ {
		a[index].X = at.scanner.NextFloat64()
		a[index].Y = at.scanner.NextFloat64()
	}
	var b point.Float64
	b.X = at.scanner.NextFloat64()
	b.Y = at.scanner.NextFloat64()
	var distance, minDistance float64
	var closerX, closerY float64 = 0, 0
	for index := 0; index < n; index++ {
		distance = math.Sqrt(math.Pow(b.X-a[index].X, 2) + math.Pow(b.Y-a[index].Y, 2))
		if index == 0 || distance < minDistance {
			minDistance = distance
			closerX, closerY = a[index].X, a[index].Y
		}
	}
	return at.checker.CompareFloat64(2, closerX, closerY)
}

// { "no": 132, "dat": "2", "ans": "2" }
func (at *arrayTask) array132() bool {
	var n = at.scanner.NextInt()
	var a = make([]point.Float64, n)
	var wantedX, wantedY float64 = 0, 0
	var distance, maxDistance float64 = 0, 0
	for index := 0; index < n; index++ {
		a[index].X = at.scanner.NextFloat64()
		a[index].Y = at.scanner.NextFloat64()
		if a[index].X < 0 && a[index].Y > 0 {
			distance = math.Sqrt(a[index].X*a[index].X + a[index].Y*a[index].Y)
			if distance > maxDistance {
				maxDistance = distance
				wantedX = a[index].X
				wantedY = a[index].Y
			}
		}
	}
	return at.checker.CompareFloat64(2, wantedX, wantedY)
}

// { "no": 133, "dat": "2", "ans": "2" }
func (at *arrayTask) array133() bool {
	var n = at.scanner.NextInt()
	var a = make([]point.Float64, n)
	var wantedX, wantedY float64 = 0, 0
	var distance, minDistance float64 = 0, 0
	for index := 0; index < n; index++ {
		a[index].X = at.scanner.NextFloat64()
		a[index].Y = at.scanner.NextFloat64()
		if a[index].X*a[index].Y > 0 {
			distance = math.Sqrt(a[index].X*a[index].X + a[index].Y*a[index].Y)
			if distance > 0 && (minDistance == 0 || distance < minDistance) {
				minDistance = distance
				wantedX = a[index].X
				wantedY = a[index].Y
			}
		}
	}
	return at.checker.CompareFloat64(2, wantedX, wantedY)
}

// { "no": 134, "dat": "2", "ans": "2" }
func (at *arrayTask) array134() bool {
	var n = at.scanner.NextInt()
	var a = make([]point.Float64, n)
	var distance, maxDistance float64 = 0, 0
	var index1, index2 = 0, 0
	for index := 0; index < n; index++ {
		a[index].X = at.scanner.NextFloat64()
		a[index].Y = at.scanner.NextFloat64()
		for i := index - 1; i >= 0; i-- {
			distance = math.Sqrt(math.Pow(a[index].X-a[i].X, 2) + math.Pow(a[index].Y-a[i].Y, 2))
			if distance >= maxDistance {
				maxDistance = distance
				index1, index2 = i, index
			}
		}
	}
	return at.checker.CompareFloat64(2, a[index1].X, a[index1].Y, a[index2].X, a[index2].Y, maxDistance)
}

// { "no": 135, "dat": "2", "ans": "2" }
func (at *arrayTask) array135() bool {
	var n1 = at.scanner.NextInt()
	var a = make([]point.Float64, n1)
	for index := 0; index < n1; index++ {
		a[index].X = at.scanner.NextFloat64()
		a[index].Y = at.scanner.NextFloat64()
	}
	var n2 = at.scanner.NextInt()
	var b = make([]point.Float64, n2)
	var aIndex, bIndex = 0, 0
	var distance, minDistance float64 = 0, 0
	for index := 0; index < n2; index++ {
		b[index].X = at.scanner.NextFloat64()
		b[index].Y = at.scanner.NextFloat64()
		for i := 0; i < n1; i++ {
			distance = math.Sqrt(math.Pow(a[i].X-b[index].X, 2) + math.Pow(a[i].Y-b[index].Y, 2))
			if distance > 0 && (minDistance == 0 || distance < minDistance) {
				minDistance = distance
				aIndex, bIndex = i, index
			}
		}
	}
	return at.checker.CompareFloat64(2, minDistance, a[aIndex].X, a[aIndex].Y, b[bIndex].X, b[bIndex].Y)
}

// { "no": 136, "dat": "2", "ans": "2" }
func (at *arrayTask) array136() bool {
	var n = at.scanner.NextInt()
	var a = make([]point.Float64, n)
	for index := 0; index < n; index++ {
		a[index].X = at.scanner.NextFloat64()
		a[index].Y = at.scanner.NextFloat64()
	}
	var sum, distance, minSum float64
	var minSumIndex = -1
	for index := 0; index < n; index++ {
		sum = 0
		for i := 0; i < n; i++ {
			if i == index {
				continue
			}
			distance = math.Sqrt(math.Pow(a[i].X-a[index].X, 2) + math.Pow(a[i].Y-a[index].Y, 2))
			sum += distance
		}
		if minSumIndex < 0 || sum < minSum {
			minSum = sum
			minSumIndex = index
		}
	}
	return at.checker.CompareFloat64(2, a[minSumIndex].X, a[minSumIndex].Y, minSum)
}

// { "no": 137, "dat": "2", "ans": "2" }
func (at *arrayTask) array137() bool {
	var n = at.scanner.NextInt()
	var a = make([]point.Float64, n)
	for index := 0; index < n; index++ {
		a[index].X = at.scanner.NextFloat64()
		a[index].Y = at.scanner.NextFloat64()
	}
	var side1, side2, side3 float64
	var perimeter, maxPerimeter float64
	var index1, index2, index3 = 0, 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i == j && j == k || k == i {
					continue
				}
				side1 = math.Sqrt(math.Pow(a[i].X-a[j].X, 2) + math.Pow(a[i].Y-a[j].Y, 2))
				side2 = math.Sqrt(math.Pow(a[j].X-a[k].X, 2) + math.Pow(a[j].Y-a[k].Y, 2))
				side3 = math.Sqrt(math.Pow(a[k].X-a[i].X, 2) + math.Pow(a[k].Y-a[i].Y, 2))
				perimeter = side1 + side2 + side3
				if perimeter > maxPerimeter {
					maxPerimeter = perimeter
					index1, index2, index3 = i, j, k
				}
			}
		}
	}
	proc.SortInc(&index1, &index2)
	proc.SortInc(&index1, &index3)
	proc.SortInc(&index2, &index3)
	return at.checker.CompareFloat64(2, maxPerimeter,
		a[index1].X, a[index1].Y,
		a[index2].X, a[index2].Y,
		a[index3].X, a[index3].Y,
	)
}

// { "no": 138, "dat": "2", "ans": "2" }
func (at *arrayTask) array138() bool {
	var n = at.scanner.NextInt()
	var a = make([]point.Float64, n)
	for index := 0; index < n; index++ {
		a[index].X = at.scanner.NextFloat64()
		a[index].Y = at.scanner.NextFloat64()
	}
	var side1, side2, side3 float64
	var perimeter, minPerimeter float64 = 0, 0
	var index1, index2, index3 = 0, 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i == j || j == k || k == i {
					continue
				}
				side1 = math.Sqrt(math.Pow(a[i].X-a[j].X, 2) + math.Pow(a[i].Y-a[j].Y, 2))
				side2 = math.Sqrt(math.Pow(a[j].X-a[k].X, 2) + math.Pow(a[j].Y-a[k].Y, 2))
				side3 = math.Sqrt(math.Pow(a[k].X-a[i].X, 2) + math.Pow(a[k].Y-a[i].Y, 2))
				perimeter = side1 + side2 + side3
				if minPerimeter == 0 || perimeter < minPerimeter {
					minPerimeter = perimeter
					index1, index2, index3 = i, j, k
				}
			}
		}
	}
	proc.SortInc(&index1, &index2)
	proc.SortInc(&index1, &index3)
	proc.SortInc(&index2, &index3)
	return at.checker.CompareFloat64(2, minPerimeter,
		a[index1].X, a[index1].Y,
		a[index2].X, a[index2].Y,
		a[index3].X, a[index3].Y,
	)
}

// { "no": 139, "dat": "", "ans": "" }
func (at *arrayTask) array139() bool {
	var n = at.scanner.NextInt()
	var a = make([]point.Int, n)
	for index := 0; index < n; index++ {
		a[index].X = at.scanner.NextInt()
		a[index].Y = at.scanner.NextInt()
	}
	var isSorted bool
	for index := 0; index < n-1; index++ {
		for i := 1; i < n-index; i++ {
			isSorted = a[i-1].X < a[i].X || (a[i-1].X == a[i].X && a[i-1].Y < a[i].Y)
			if !isSorted {
				a[i-1], a[i] = a[i], a[i-1]
			}
		}
	}
	for index := 0; index < n; index++ {
		if !at.checker.CompareInt(a[index].X, a[index].Y) {
			return false
		}
	}
	return true
}

// { "no": 140, "dat": "", "ans": "" }
func (at *arrayTask) array140() bool {
	var n = at.scanner.NextInt()
	var a = make([]point.Int, n)
	for index := 0; index < n; index++ {
		a[index].X = at.scanner.NextInt()
		a[index].Y = at.scanner.NextInt()
	}
	var isSorted bool
	for index := 0; index < n-1; index++ {
		for i := 1; i < n-index; i++ {
			isSorted = a[i-1].X+a[i-1].Y < a[i].X+a[i].Y ||
				(a[i-1].X+a[i-1].Y == a[i].X+a[i].Y && a[i-1].X < a[i].X)
			if isSorted {
				a[i-1], a[i] = a[i], a[i-1]
			}
		}
	}
	for index := 0; index < n; index++ {
		if !at.checker.CompareInt(a[index].X, a[index].Y) {
			return false
		}
	}
	return true
}
