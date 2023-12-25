package task

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type integerTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewIntegerTask(pathPrefix string) Task {
	return &integerTask{
		dataPath: pathPrefix + "/02integer/Integer%03d/%03d/%03d",
		name:     "Integer",
		count:    30,
	}
}

func (it *integerTask) GetCount() int { return it.count }

func (it *integerTask) GetName() string { return it.name }

func (it *integerTask) SetData(taskNo int, testNo int) error {
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

func (it *integerTask) ScannerEOF() bool { return it.scanner.EOF() }

func (it *integerTask) CheckerEOF() bool { return it.checker.EOF() }

func (it *integerTask) CleanData() {
	it.scanner.RemoveFiles()
	it.checker.RemoveFiles()
}

func (it *integerTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return it.integer1
	case 2:
		return it.integer2
	case 3:
		return it.integer3
	case 4:
		return it.integer4
	case 5:
		return it.integer5
	case 6:
		return it.integer6
	case 7:
		return it.integer7
	case 8:
		return it.integer8
	case 9:
		return it.integer9
	case 10:
		return it.integer10
	case 11:
		return it.integer11
	case 12:
		return it.integer12
	case 13:
		return it.integer13
	case 14:
		return it.integer14
	case 15:
		return it.integer15
	case 16:
		return it.integer16
	case 17:
		return it.integer17
	case 18:
		return it.integer18
	case 19:
		return it.integer19
	case 20:
		return it.integer20
	case 21:
		return it.integer21
	case 22:
		return it.integer22
	case 23:
		return it.integer23
	case 24:
		return it.integer24
	case 25:
		return it.integer25
	case 26:
		return it.integer26
	case 27:
		return it.integer27
	case 28:
		return it.integer28
	case 29:
		return it.integer29
	case 30:
		return it.integer30
	default:
		return nil
	}
}

func (it *integerTask) integer1() bool {
	var l = it.scanner.NextInt()
	var meters = l / 100
	return it.checker.CompareInt(meters)
}

func (it *integerTask) integer2() bool {
	var m = it.scanner.NextInt()
	var tons = m / 1000
	return it.checker.CompareInt(tons)
}

func (it *integerTask) integer3() bool {
	var b = it.scanner.NextInt()
	var kBytes = b / 1024
	return it.checker.CompareInt(kBytes)
}

func (it *integerTask) integer4() bool {
	var a = it.scanner.NextInt()
	var b = it.scanner.NextInt()
	var segments = a / b
	return it.checker.CompareInt(segments)
}

func (it *integerTask) integer5() bool {
	var a = it.scanner.NextInt()
	var b = it.scanner.NextInt()
	var unusedPart = a % b
	return it.checker.CompareInt(unusedPart)
}

func (it *integerTask) integer6() bool {
	var number = it.scanner.NextInt()
	var tens = number / 10
	var ones = number % 10
	return it.checker.CompareInt(tens, ones)
}

func (it *integerTask) integer7() bool {
	var number = it.scanner.NextInt()
	var tens = number / 10
	var ones = number % 10
	var sum = tens + ones
	var pro = tens * ones
	return it.checker.CompareInt(sum, pro)
}

func (it *integerTask) integer8() bool {
	var number = it.scanner.NextInt()
	var tens = number / 10
	var ones = number % 10
	number = ones*10 + tens
	return it.checker.CompareInt(number)
}

func (it *integerTask) integer9() bool {
	var number = it.scanner.NextInt()
	var hundreds = number / 100
	return it.checker.CompareInt(hundreds)
}

func (it *integerTask) integer10() bool {
	var number = it.scanner.NextInt()
	var ones = number % 10
	var tens = number / 10 % 10
	return it.checker.CompareInt(ones, tens)
}

func (it *integerTask) integer11() bool {
	var number = it.scanner.NextInt()
	var hundreds = number / 100
	var tens = number / 10 % 10
	var ones = number % 10
	var sum = hundreds + tens + ones
	var pro = hundreds * tens * ones
	return it.checker.CompareInt(sum, pro)
}

func (it *integerTask) integer12() bool {
	var number = it.scanner.NextInt()
	var hundreds = number / 100
	var tens = number / 10 % 10
	var ones = number % 10
	number = ones*100 + tens*10 + hundreds
	return it.checker.CompareInt(number)
}

func (it *integerTask) integer13() bool {
	var number = it.scanner.NextInt()
	var hundreds = number / 100
	number = number%100*10 + hundreds
	return it.checker.CompareInt(number)
}

func (it *integerTask) integer14() bool {
	var number = it.scanner.NextInt()
	var ones = number % 10
	number = ones*100 + number/10
	return it.checker.CompareInt(number)
}

func (it *integerTask) integer15() bool {
	var number = it.scanner.NextInt()
	var hundreds = number / 100
	var tens = number / 10 % 10
	var ones = number % 10
	number = tens*100 + hundreds*10 + ones
	return it.checker.CompareInt(number)
}

func (it *integerTask) integer16() bool {
	var number = it.scanner.NextInt()
	var hundreds = number / 100
	var tens = number / 10 % 10
	var ones = number % 10
	number = hundreds*100 + ones*10 + tens
	return it.checker.CompareInt(number)
}

func (it *integerTask) integer17() bool {
	var number = it.scanner.NextInt()
	var hundreds = number / 100 % 10
	return it.checker.CompareInt(hundreds)
}

func (it *integerTask) integer18() bool {
	var number = it.scanner.NextInt()
	var thousands = number / 1000 % 10
	return it.checker.CompareInt(thousands)
}

func (it *integerTask) integer19() bool {
	var n = it.scanner.NextInt()
	var minutes = n / 60
	return it.checker.CompareInt(minutes)
}

func (it *integerTask) integer20() bool {
	var n = it.scanner.NextInt()
	var hours = n / 3600
	return it.checker.CompareInt(hours)
}

func (it *integerTask) integer21() bool {
	var n = it.scanner.NextInt()
	var seconds = n % 60
	return it.checker.CompareInt(seconds)
}

func (it *integerTask) integer22() bool {
	var n = it.scanner.NextInt()
	var seconds = n % 3600
	return it.checker.CompareInt(seconds)
}

func (it *integerTask) integer23() bool {
	var n = it.scanner.NextInt()
	var minutes = n % 3600 / 60
	return it.checker.CompareInt(minutes)
}

func (it *integerTask) integer24() bool {
	var k = it.scanner.NextInt()
	var weekDay = k % 7
	return it.checker.CompareInt(weekDay)
}

func (it *integerTask) integer25() bool {
	var k = it.scanner.NextInt()
	var weekDay = (k + 3) % 7
	return it.checker.CompareInt(weekDay)
}

func (it *integerTask) integer26() bool {
	var k = it.scanner.NextInt()
	var weekDay = k%7 + 1
	return it.checker.CompareInt(weekDay)
}

func (it *integerTask) integer27() bool {
	var k = it.scanner.NextInt()
	var weekDay = (k+4)%7 + 1
	return it.checker.CompareInt(weekDay)
}

func (it *integerTask) integer28() bool {
	var k = it.scanner.NextInt()
	var n = it.scanner.NextInt()
	var weekDay = (k+n-2)%7 + 1
	return it.checker.CompareInt(weekDay)
}

func (it *integerTask) integer29() bool {
	var a = it.scanner.NextInt()
	var b = it.scanner.NextInt()
	var c = it.scanner.NextInt()
	var cInA = a / c
	var cInB = b / c
	var squares = cInA * cInB
	var unusedArea = a*b - squares*c*c
	return it.checker.CompareInt(squares, unusedArea)
}

func (it *integerTask) integer30() bool {
	var year = it.scanner.NextInt()
	var century = (year-1)/100 + 1
	return it.checker.CompareInt(century)
}
