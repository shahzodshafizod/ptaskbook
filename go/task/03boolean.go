package task

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type booleanTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewBooleanTask(pathPrefix string) Task {
	return &booleanTask{
		dataPath: pathPrefix + "/03boolean/Boolean%03d/%03d/%03d",
		name:     "Boolean",
		count:    40,
	}
}

func (bt *booleanTask) GetCount() int { return bt.count }

func (bt *booleanTask) GetName() string { return bt.name }

func (bt *booleanTask) SetData(taskNo int, testNo int) error {
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

func (bt *booleanTask) ScannerEOF() bool { return bt.scanner.EOF() }

func (bt *booleanTask) CheckerEOF() bool { return bt.checker.EOF() }

func (bt *booleanTask) CleanData() {
	bt.scanner.RemoveFiles()
	bt.checker.RemoveFiles()
}

func (bt *booleanTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return bt.boolean1
	case 2:
		return bt.boolean2
	case 3:
		return bt.boolean3
	case 4:
		return bt.boolean4
	case 5:
		return bt.boolean5
	case 6:
		return bt.boolean6
	case 7:
		return bt.boolean7
	case 8:
		return bt.boolean8
	case 9:
		return bt.boolean9
	case 10:
		return bt.boolean10
	case 11:
		return bt.boolean11
	case 12:
		return bt.boolean12
	case 13:
		return bt.boolean13
	case 14:
		return bt.boolean14
	case 15:
		return bt.boolean15
	case 16:
		return bt.boolean16
	case 17:
		return bt.boolean17
	case 18:
		return bt.boolean18
	case 19:
		return bt.boolean19
	case 20:
		return bt.boolean20
	case 21:
		return bt.boolean21
	case 22:
		return bt.boolean22
	case 23:
		return bt.boolean23
	case 24:
		return bt.boolean24
	case 25:
		return bt.boolean25
	case 26:
		return bt.boolean26
	case 27:
		return bt.boolean27
	case 28:
		return bt.boolean28
	case 29:
		return bt.boolean29
	case 30:
		return bt.boolean30
	case 31:
		return bt.boolean31
	case 32:
		return bt.boolean32
	case 33:
		return bt.boolean33
	case 34:
		return bt.boolean34
	case 35:
		return bt.boolean35
	case 36:
		return bt.boolean36
	case 37:
		return bt.boolean37
	case 38:
		return bt.boolean38
	case 39:
		return bt.boolean39
	case 40:
		return bt.boolean40
	default:
		return nil
	}
}

// { "no": 1, "dat": "", "ans": "" }
func (bt *booleanTask) boolean1() bool {
	var a = bt.scanner.NextInt()
	var isPositive = a > 0
	return bt.checker.CompareBool(isPositive)
}

// { "no": 2, "dat": "", "ans": "" }
func (bt *booleanTask) boolean2() bool {
	var a = bt.scanner.NextInt()
	var isOdd = a%2 != 0
	return bt.checker.CompareBool(isOdd)
}

// { "no": 3, "dat": "", "ans": "" }
func (bt *booleanTask) boolean3() bool {
	var a = bt.scanner.NextInt()
	var isEven = a%2 == 0
	return bt.checker.CompareBool(isEven)
}

// { "no": 4, "dat": "", "ans": "" }
func (bt *booleanTask) boolean4() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var result = a > 2 && b <= 3
	return bt.checker.CompareBool(result)
}

// { "no": 5, "dat": "", "ans": "" }
func (bt *booleanTask) boolean5() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var result = a >= 0 || b < -2
	return bt.checker.CompareBool(result)
}

// { "no": 6, "dat": "", "ans": "" }
func (bt *booleanTask) boolean6() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var c = bt.scanner.NextInt()
	var isOrderedAsc = a < b && b < c
	return bt.checker.CompareBool(isOrderedAsc)
}

// { "no": 7, "dat": "", "ans": "" }
func (bt *booleanTask) boolean7() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var c = bt.scanner.NextInt()
	var isOrdered = a < b && b < c || a > b && b > c
	return bt.checker.CompareBool(isOrdered)
}

// { "no": 8, "dat": "", "ans": "" }
func (bt *booleanTask) boolean8() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var eachOdd = (a*b)%2 != 0
	return bt.checker.CompareBool(eachOdd)
}

// { "no": 9, "dat": "", "ans": "" }
func (bt *booleanTask) boolean9() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var anyOdd = a%2 != 0 || b%2 != 0
	return bt.checker.CompareBool(anyOdd)
}

// { "no": 10, "dat": "", "ans": "" }
func (bt *booleanTask) boolean10() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var oneOdd = (a+b)%2 != 0
	return bt.checker.CompareBool(oneOdd)
}

// { "no": 11, "dat": "", "ans": "" }
func (bt *booleanTask) boolean11() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var equalParity = (a+b)%2 == 0
	return bt.checker.CompareBool(equalParity)
}

// { "no": 12, "dat": "", "ans": "" }
func (bt *booleanTask) boolean12() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var c = bt.scanner.NextInt()
	var eachPositive = a > 0 && b > 0 && c > 0
	return bt.checker.CompareBool(eachPositive)
}

// { "no": 13, "dat": "", "ans": "" }
func (bt *booleanTask) boolean13() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var c = bt.scanner.NextInt()
	var anyPositive = a > 0 || b > 0 || c > 0
	return bt.checker.CompareBool(anyPositive)
}

// { "no": 14, "dat": "", "ans": "" }
func (bt *booleanTask) boolean14() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var c = bt.scanner.NextInt()
	var onePositive = (a > 0 && b <= 0 && c <= 0) ||
		(a <= 0 && b > 0 && c <= 0) ||
		(a <= 0 && b <= 0 && c > 0)
	return bt.checker.CompareBool(onePositive)
}

// { "no": 15, "dat": "", "ans": "" }
func (bt *booleanTask) boolean15() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var c = bt.scanner.NextInt()
	var twoPositives = (a > 0 && b > 0 && c <= 0) ||
		(a > 0 && b <= 0 && c > 0) ||
		(a <= 0 && b > 0 && c > 0)
	return bt.checker.CompareBool(twoPositives)
}

// { "no": 16, "dat": "", "ans": "" }
func (bt *booleanTask) boolean16() bool {
	var number = bt.scanner.NextInt()
	var isTwoDigitalEven = number > 9 && number < 100 && number%2 == 0
	return bt.checker.CompareBool(isTwoDigitalEven)
}

// { "no": 17, "dat": "", "ans": "" }
func (bt *booleanTask) boolean17() bool {
	var number = bt.scanner.NextInt()
	var threeDigitalOdd = number > 99 && number < 1000 && number%2 != 0
	return bt.checker.CompareBool(threeDigitalOdd)
}

// { "no": 18, "dat": "", "ans": "" }
func (bt *booleanTask) boolean18() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var c = bt.scanner.NextInt()
	var anyPair = a == b || b == c || c == a
	return bt.checker.CompareBool(anyPair)
}

// { "no": 19, "dat": "", "ans": "" }
func (bt *booleanTask) boolean19() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var c = bt.scanner.NextInt()
	var anyOppositePair = a == -b || b == -c || c == -a
	return bt.checker.CompareBool(anyOppositePair)
}

// { "no": 20, "dat": "", "ans": "" }
func (bt *booleanTask) boolean20() bool {
	var number = bt.scanner.NextInt()
	var hundreds = number / 100
	var tens = number / 10 % 10
	var ones = number % 10
	var areDifferent = hundreds != tens && hundreds != ones && tens != ones
	return bt.checker.CompareBool(areDifferent)
}

// { "no": 21, "dat": "", "ans": "" }
func (bt *booleanTask) boolean21() bool {
	var number = bt.scanner.NextInt()
	var hundreds = number / 100
	var tens = number / 10 % 10
	var ones = number % 10
	var areOrderedAsc = hundreds < tens && tens < ones
	return bt.checker.CompareBool(areOrderedAsc)
}

// { "no": 22, "dat": "", "ans": "" }
func (bt *booleanTask) boolean22() bool {
	var number = bt.scanner.NextInt()
	var hundreds = number / 100
	var tens = number / 10 % 10
	var ones = number % 10
	var areOrdered = hundreds < tens && tens < ones ||
		hundreds > tens && tens > ones
	return bt.checker.CompareBool(areOrdered)
}

// { "no": 23, "dat": "", "ans": "" }
func (bt *booleanTask) boolean23() bool {
	var number = bt.scanner.NextInt()
	var firstHalf = number / 100
	var secondHalf = number%10*10 + number/10%10
	var isPalindrome = firstHalf == secondHalf
	return bt.checker.CompareBool(isPalindrome)
}

// { "no": 24, "dat": "2", "ans": "" }
func (bt *booleanTask) boolean24() bool {
	var a = bt.scanner.NextFloat32()
	var b = bt.scanner.NextFloat32()
	var c = bt.scanner.NextFloat32()
	var d = b*b - 4*a*c
	var hasRoots = d >= 0
	return bt.checker.CompareBool(hasRoots)
}

// { "no": 25, "dat": "2", "ans": "" }
func (bt *booleanTask) boolean25() bool {
	var x = bt.scanner.NextFloat32()
	var y = bt.scanner.NextFloat32()
	var inTheSecondQuarter = x < 0 && y > 0
	return bt.checker.CompareBool(inTheSecondQuarter)
}

// { "no": 26, "dat": "2", "ans": "" }
func (bt *booleanTask) boolean26() bool {
	var x = bt.scanner.NextFloat32()
	var y = bt.scanner.NextFloat32()
	var inTheFourthQuarter = x > 0 && y < 0
	return bt.checker.CompareBool(inTheFourthQuarter)
}

// { "no": 27, "dat": "2", "ans": "" }
func (bt *booleanTask) boolean27() bool {
	var x = bt.scanner.NextFloat32()
	var y = bt.scanner.NextFloat32()
	var result = x < 0 && y != 0
	return bt.checker.CompareBool(result)
}

// { "no": 28, "dat": "2", "ans": "" }
func (bt *booleanTask) boolean28() bool {
	var x = bt.scanner.NextFloat32()
	var y = bt.scanner.NextFloat32()
	var result = x*y > 0
	return bt.checker.CompareBool(result)
}

// { "no": 29, "dat": "2", "ans": "" }
func (bt *booleanTask) boolean29() bool {
	var x = bt.scanner.NextFloat32()
	var y = bt.scanner.NextFloat32()
	var x1 = bt.scanner.NextFloat32()
	var y1 = bt.scanner.NextFloat32()
	var x2 = bt.scanner.NextFloat32()
	var y2 = bt.scanner.NextFloat32()
	var isInside = x > x1 && x < x2 && y > y2 && y < y1
	return bt.checker.CompareBool(isInside)
}

// { "no": 30, "dat": "", "ans": "" }
func (bt *booleanTask) boolean30() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var c = bt.scanner.NextInt()
	var isEquilAteral = a == b && b == c
	return bt.checker.CompareBool(isEquilAteral)
}

// { "no": 31, "dat": "", "ans": "" }
func (bt *booleanTask) boolean31() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var c = bt.scanner.NextInt()
	var isIsosceles = a == b || b == c || c == a
	return bt.checker.CompareBool(isIsosceles)
}

// { "no": 32, "dat": "", "ans": "" }
func (bt *booleanTask) boolean32() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var c = bt.scanner.NextInt()
	var isRightTriangle = a*a == b*b+c*c ||
		b*b == a*a+c*c || c*c == a*a+b*b
	return bt.checker.CompareBool(isRightTriangle)
}

// { "no": 33, "dat": "", "ans": "" }
func (bt *booleanTask) boolean33() bool {
	var a = bt.scanner.NextInt()
	var b = bt.scanner.NextInt()
	var c = bt.scanner.NextInt()
	var existsTriangle = a+b > c && a+c > b && b+c > a
	return bt.checker.CompareBool(existsTriangle)
}

// { "no": 34, "dat": "", "ans": "" }
func (bt *booleanTask) boolean34() bool {
	var x = bt.scanner.NextInt()
	var y = bt.scanner.NextInt()
	var isWhiteSquare = (x+y)%2 != 0
	return bt.checker.CompareBool(isWhiteSquare)
}

// { "no": 35, "dat": "", "ans": "" }
func (bt *booleanTask) boolean35() bool {
	var x1 = bt.scanner.NextInt()
	var y1 = bt.scanner.NextInt()
	var x2 = bt.scanner.NextInt()
	var y2 = bt.scanner.NextInt()
	var sameColor = (x1+y1)%2 == (x2+y2)%2
	return bt.checker.CompareBool(sameColor)
}

// { "no": 36, "dat": "", "ans": "" }
func (bt *booleanTask) boolean36() bool {
	var x1 = bt.scanner.NextInt()
	var y1 = bt.scanner.NextInt()
	var x2 = bt.scanner.NextInt()
	var y2 = bt.scanner.NextInt()
	var canMoveRook = x1 == x2 || y1 == y2
	return bt.checker.CompareBool(canMoveRook)
}

// { "no": 37, "dat": "", "ans": "" }
func (bt *booleanTask) boolean37() bool {
	var x1 = bt.scanner.NextInt()
	var y1 = bt.scanner.NextInt()
	var x2 = bt.scanner.NextInt()
	var y2 = bt.scanner.NextInt()
	var canMoveKing = math.Abs(float64(x2-x1)) < 2 && math.Abs(float64(y2-y1)) < 2
	return bt.checker.CompareBool(canMoveKing)
}

// { "no": 38, "dat": "", "ans": "" }
func (bt *booleanTask) boolean38() bool {
	var x1 = bt.scanner.NextInt()
	var y1 = bt.scanner.NextInt()
	var x2 = bt.scanner.NextInt()
	var y2 = bt.scanner.NextInt()
	var canMoveBishop = math.Abs(float64(x2-x1)) == math.Abs(float64(y2-y1))
	return bt.checker.CompareBool(canMoveBishop)
}

// { "no": 39, "dat": "", "ans": "" }
func (bt *booleanTask) boolean39() bool {
	var x1 = bt.scanner.NextInt()
	var y1 = bt.scanner.NextInt()
	var x2 = bt.scanner.NextInt()
	var y2 = bt.scanner.NextInt()
	var canMoveQueen = x1 == x2 || y1 == y2 ||
		math.Abs(float64(x2-x1)) == math.Abs(float64(y2-y1))
	return bt.checker.CompareBool(canMoveQueen)
}

// { "no": 40, "dat": "", "ans": "" }
func (bt *booleanTask) boolean40() bool {
	var x1 = bt.scanner.NextInt()
	var y1 = bt.scanner.NextInt()
	var x2 = bt.scanner.NextInt()
	var y2 = bt.scanner.NextInt()
	var xs = math.Abs(float64(x2 - x1))
	var ys = math.Abs(float64(y2 - y1))
	var canMoveKnight = xs*ys == 2 && xs != ys
	return bt.checker.CompareBool(canMoveKnight)
}
