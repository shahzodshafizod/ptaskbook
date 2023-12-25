package task

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type caseTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewCaseTask(pathPrefix string) Task {
	return &caseTask{
		dataPath: pathPrefix + "/05case/Case%03d/%03d/%03d",
		name:     "Case",
		count:    20,
	}
}

func (ct *caseTask) GetCount() int { return ct.count }

func (ct *caseTask) GetName() string { return ct.name }

func (ct *caseTask) SetData(taskNo int, testNo int) error {
	var filePath = fmt.Sprintf(ct.dataPath, taskNo, testNo, testNo)
	scanner, err := scanner.New(filePath + ".dat")
	if err != nil {
		return errors.Wrap(err, "scanner.New")
	}
	checker, err := checker.New(filePath + ".ans")
	if err != nil {
		return errors.Wrap(err, "checker.New")
	}
	ct.scanner = scanner
	ct.checker = checker
	return nil
}

func (ct *caseTask) ScannerEOF() bool { return ct.scanner.EOF() }

func (ct *caseTask) CheckerEOF() bool { return ct.checker.EOF() }

func (ct *caseTask) CleanData() {
	ct.scanner.RemoveFiles()
	ct.checker.RemoveFiles()
}

func (ct *caseTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return ct.case1
	case 2:
		return ct.case2
	case 3:
		return ct.case3
	case 4:
		return ct.case4
	case 5:
		return ct.case5
	case 6:
		return ct.case6
	case 7:
		return ct.case7
	case 8:
		return ct.case8
	case 9:
		return ct.case9
	case 10:
		return ct.case10
	case 11:
		return ct.case11
	case 12:
		return ct.case12
	case 13:
		return ct.case13
	case 14:
		return ct.case14
	case 15:
		return ct.case15
	case 16:
		return ct.case16
	case 17:
		return ct.case17
	case 18:
		return ct.case18
	case 19:
		return ct.case19
	case 20:
		return ct.case20
	default:
		return nil
	}
}

// { "no": 1, "dat": "", "ans": "" }
func (ct *caseTask) case1() bool {
	var number = ct.scanner.NextInt()
	var weekDayName string
	switch number {
	case 1:
		weekDayName = "Monday"
	case 2:
		weekDayName = "Tuesday"
	case 3:
		weekDayName = "Wednesday"
	case 4:
		weekDayName = "Thursday"
	case 5:
		weekDayName = "Friday"
	case 6:
		weekDayName = "Saturday"
	case 7:
		weekDayName = "Sunday"
	}
	return ct.checker.CompareWord(weekDayName)
}

// { "no": 2, "dat": "", "ans": "" }
func (ct *caseTask) case2() bool {
	var k = ct.scanner.NextInt()
	var mark string
	switch k {
	case 1:
		mark = "bad"
	case 2:
		mark = "unsatisfactory"
	case 3:
		mark = "mediocre"
	case 4:
		mark = "good"
	case 5:
		mark = "excellent"
	default:
		mark = "error"
	}
	return ct.checker.CompareWord(mark)
}

// { "no": 3, "dat": "", "ans": "" }
func (ct *caseTask) case3() bool {
	var monthNo = ct.scanner.NextInt()
	var season string
	switch monthNo {
	case 1, 2, 12:
		season = "Winter"
	case 3, 4, 5:
		season = "Spring"
	case 6, 7, 8:
		season = "Summer"
	case 9, 10, 11:
		season = "Autumn"
	}
	return ct.checker.CompareWord(season)
}

// { "no": 4, "dat": "", "ans": "" }
func (ct *caseTask) case4() bool {
	var monthNo = ct.scanner.NextInt()
	var daysCount int
	switch monthNo {
	case 1, 3, 5, 7, 8, 10, 12:
		daysCount = 31
	case 4, 6, 9, 11:
		daysCount = 30
	case 2:
		daysCount = 28
	}
	return ct.checker.CompareInt(daysCount)
}

// { "no": 5, "dat": "2", "ans": "2" }
func (ct *caseTask) case5() bool {
	var n = ct.scanner.NextInt()
	var a = ct.scanner.NextFloat64()
	var b = ct.scanner.NextFloat64()
	var result float64
	switch n {
	case 1:
		result = a + b
	case 2:
		result = a - b
	case 3:
		result = a * b
	case 4:
		result = a / b
	}
	return ct.checker.CompareFloat64(2, result)
}

// { "no": 6, "dat": "2", "ans": "5" }
func (ct *caseTask) case6() bool {
	var n = ct.scanner.NextInt()
	var l = ct.scanner.NextFloat64()
	var meters float64
	switch n {
	case 1:
		meters = l / 10
	case 2:
		meters = l * 1000
	case 3:
		meters = l
	case 4:
		meters = l / 1000
	case 5:
		meters = l / 100
	}
	return ct.checker.CompareFloat64(5, meters)
}

// { "no": 7, "dat": "2", "ans": "8" }
func (ct *caseTask) case7() bool {
	var n = ct.scanner.NextInt()
	var m = ct.scanner.NextFloat64()
	var kg float64
	switch n {
	case 1:
		kg = m
	case 2:
		kg = m / 1000000
	case 3:
		kg = m / 1000
	case 4:
		kg = m * 1000
	case 5:
		kg = m * 100
	}
	return ct.checker.CompareFloat64(8, kg)
}

// { "no": 8, "dat": "", "ans": "" }
func (ct *caseTask) case8() bool {
	var d, m = ct.scanner.NextInt(), ct.scanner.NextInt()
	switch d {
	case 1:
		switch m {
		case 1:
			d, m = 31, 13
		case 3:
			d = 28
		case 2, 4, 6, 8, 9, 11:
			d = 31
		case 5, 7, 10, 12:
			d = 30
		}
		m--
	default:
		d--
	}
	return ct.checker.CompareInt(d, m)
}

// { "no": 9, "dat": "", "ans": "" }
func (ct *caseTask) case9() bool {
	var d, m = ct.scanner.NextInt(), ct.scanner.NextInt()
	switch d {
	case 31:
		d = 0
		switch m {
		case 12:
			m = 0
		}
	case 30:
		switch m {
		case 4, 6, 9, 11:
			d = 0
		}
	case 28:
		switch m {
		case 2:
			d = 0
		}
	}
	if d == 0 {
		m++
	}
	d++
	return ct.checker.CompareInt(d, m)
}

// { "no": 10, "dat": "", "ans": "" }
func (ct *caseTask) case10() bool {
	var c, _ = ct.scanner.NextRune()
	var n = ct.scanner.NextInt()
	switch n {
	case 1:
		switch c {
		case 'N':
			c = 'W'
		case 'E':
			c = 'N'
		case 'S':
			c = 'E'
		case 'W':
			c = 'S'
		}
	case -1:
		switch c {
		case 'N':
			c = 'E'
		case 'E':
			c = 'S'
		case 'S':
			c = 'W'
		case 'W':
			c = 'N'
		}
	}
	return ct.checker.CompareRune(c)
}

// { "no": 11, "dat": "", "ans": "" }
func (ct *caseTask) case11() bool {
	var c, _ = ct.scanner.NextRune()
	var n1, n2 = ct.scanner.NextInt(), ct.scanner.NextInt()
	switch n1 + n2 {
	case 1: // left;
		switch c {
		case 'N':
			c = 'W'
		case 'E':
			c = 'N'
		case 'S':
			c = 'E'
		case 'W':
			c = 'S'
		}
	case 3: // right;
		switch c {
		case 'N':
			c = 'E'
		case 'E':
			c = 'S'
		case 'S':
			c = 'W'
		case 'W':
			c = 'N'
		}
	case 2, -2: // 180;
		switch c {
		case 'N':
			c = 'S'
		case 'E':
			c = 'W'
		case 'S':
			c = 'N'
		case 'W':
			c = 'E'
		}
	}
	return ct.checker.CompareRune(c)
}

// { "no": 12, "dat": "2", "ans": "2" }
func (ct *caseTask) case12() bool {
	var index = ct.scanner.NextInt()
	var value = ct.scanner.NextFloat64()
	var r, d, l, s float64
	const pi = 3.14
	switch index {
	case 1:
		r = value
		d = 2 * r
		l = 2 * pi * r
		s = pi * r * r
	case 2:
		d = value
		r = d / 2
		l = 2 * pi * r
		s = pi * r * r
	case 3:
		l = value
		r = l / (2 * pi)
		d = 2 * r
		s = pi * r * r
	case 4:
		s = value
		r = math.Sqrt(s / pi)
		d = 2 * r
		l = 2 * pi * r
	}
	if index != 1 {
		if !ct.checker.CompareFloat64(2, r) {
			return false
		}
	}
	if index != 2 {
		if !ct.checker.CompareFloat64(2, d) {
			return false
		}
	}
	if index != 3 {
		if !ct.checker.CompareFloat64(2, l) {
			return false
		}
	}
	if index != 4 {
		if !ct.checker.CompareFloat64(2, s) {
			return false
		}
	}
	return true
}

// { "no": 13, "dat": "2", "ans": "2" }
func (ct *caseTask) case13() bool {
	var index = ct.scanner.NextInt()
	var value = ct.scanner.NextFloat64()
	var a, c, h, s float64
	switch index {
	case 1:
		a = value
		c = a * math.Sqrt(2)
		h = c / 2
		s = c * h / 2
	case 2:
		c = value
		a = c / math.Sqrt(2)
		h = c / 2
		s = c * h / 2
	case 3:
		h = value
		c = 2 * h
		a = c / math.Sqrt(2)
		s = c * h / 2
	case 4:
		s = value
		c = math.Sqrt(4 * s)
		a = c / math.Sqrt(2)
		h = c / 2
	}
	if index != 1 {
		if !ct.checker.CompareFloat64(2, a) {
			return false
		}
	}
	if index != 2 {
		if !ct.checker.CompareFloat64(2, c) {
			return false
		}
	}
	if index != 3 {
		if !ct.checker.CompareFloat64(2, h) {
			return false
		}
	}
	if index != 4 {
		if !ct.checker.CompareFloat64(2, s) {
			return false
		}
	}
	return true
}

// { "no": 14, "dat": "2", "ans": "2" }
func (ct *caseTask) case14() bool {
	var index = ct.scanner.NextInt()
	var value = ct.scanner.NextFloat64()
	var a, r1, r2, s float64
	switch index {
	case 1:
		a = value
		r1 = a * math.Sqrt(3) / 6
		r2 = 2 * r1
		s = a * a * math.Sqrt(3) / 4
	case 2:
		r1 = value
		a = 6 / math.Sqrt(3) * r1
		r2 = 2 * r1
		s = a * a * math.Sqrt(3) / 4
	case 3:
		r2 = value
		r1 = r2 / 2
		a = 6 / math.Sqrt(3) * r1
		s = a * a * math.Sqrt(3) / 4
	case 4:
		s = value
		a = math.Sqrt(4 / math.Sqrt(3) * s)
		r1 = a * math.Sqrt(3) / 6
		r2 = 2 * r1
	}
	if index != 1 {
		if !ct.checker.CompareFloat64(2, a) {
			return false
		}
	}
	if index != 2 {
		if !ct.checker.CompareFloat64(2, r1) {
			return false
		}
	}
	if index != 3 {
		if !ct.checker.CompareFloat64(2, r2) {
			return false
		}
	}
	if index != 4 {
		if !ct.checker.CompareFloat64(2, s) {
			return false
		}
	}
	return true
}

// { "no": 15, "dat": "", "ans": "" }
func (ct *caseTask) case15() bool {
	var cardValue = ct.scanner.NextInt()
	var cardSuit = ct.scanner.NextInt()
	var description string
	switch cardValue {
	case 6:
		description += "six"
	case 7:
		description += "seven"
	case 8:
		description += "eight"
	case 9:
		description += "nine"
	case 10:
		description += "ten"
	case 11:
		description += "jack"
	case 12:
		description += "queen"
	case 13:
		description += "king"
	case 14:
		description += "ace"
	}
	description += " of "
	switch cardSuit {
	case 1:
		description += "spades"
	case 2:
		description += "clubs"
	case 3:
		description += "diamonds"
	case 4:
		description += "hearts"
	}
	return ct.checker.CompareLine(description)
}

// { "no": 16, "dat": "", "ans": "" }
func (ct *caseTask) case16() bool {
	var age = ct.scanner.NextInt()
	var alphabetic string
	switch age / 10 {
	case 2:
		alphabetic += "twenty"
	case 3:
		alphabetic += "thirty"
	case 4:
		alphabetic += "fourty"
	case 5:
		alphabetic += "fifty"
	case 6:
		alphabetic += "sixty"
	}
	switch age % 10 {
	case 1:
		alphabetic += "-one"
	case 2:
		alphabetic += "-two"
	case 3:
		alphabetic += "-three"
	case 4:
		alphabetic += "-four"
	case 5:
		alphabetic += "-five"
	case 6:
		alphabetic += "-six"
	case 7:
		alphabetic += "-seven"
	case 8:
		alphabetic += "-eight"
	case 9:
		alphabetic += "-nine"
	}
	alphabetic += " years"
	return ct.checker.CompareLine(alphabetic)
}

// { "no": 17, "dat": "", "ans": "" }
func (ct *caseTask) case17() bool {
	var taskIndex = ct.scanner.NextInt()
	var tens = taskIndex / 10
	var ones = taskIndex % 10
	var alphabetic = "the "
	switch tens {
	case 1:
		switch ones {
		case 0:
			alphabetic += "tenth"
		case 1:
			alphabetic += "eleventh"
		case 2:
			alphabetic += "twelfth"
		case 3:
			alphabetic += "thirteenth"
		case 4:
			alphabetic += "fourteenth"
		case 5:
			alphabetic += "fifteenth"
		case 6:
			alphabetic += "sixteenth"
		case 7:
			alphabetic += "seventeenth"
		case 8:
			alphabetic += "eighteenth"
		case 9:
			alphabetic += "nineteenth"
		}
	case 2:
		alphabetic += "twent"
	case 3:
		alphabetic += "thirt"
	case 4:
		alphabetic += "fort"
	}
	if tens != 1 {
		switch ones {
		case 1:
			alphabetic += "y-first"
		case 2:
			alphabetic += "y-second"
		case 3:
			alphabetic += "y-third"
		case 4:
			alphabetic += "y-fourth"
		case 5:
			alphabetic += "y-fifth"
		case 6:
			alphabetic += "y-sixth"
		case 7:
			alphabetic += "y-seventh"
		case 8:
			alphabetic += "y-eighth"
		case 9:
			alphabetic += "y-ninth"
		default:
			alphabetic += "ieth"
		}
	}
	alphabetic += " task"
	return ct.checker.CompareLine(alphabetic)
}

// { "no": 18, "dat": "", "ans": "" }
func (ct *caseTask) case18() bool {
	var number = ct.scanner.NextInt()
	var hundreds = number / 100
	var tens = number / 10 % 10
	var ones = number % 10
	var alphabetic string
	switch hundreds {
	case 1:
		alphabetic += "one"
	case 2:
		alphabetic += "two"
	case 3:
		alphabetic += "three"
	case 4:
		alphabetic += "four"
	case 5:
		alphabetic += "five"
	case 6:
		alphabetic += "six"
	case 7:
		alphabetic += "seven"
	case 8:
		alphabetic += "eight"
	case 9:
		alphabetic += "nine"
	}
	if hundreds != 0 {
		alphabetic += " hundred"
	}
	if tens != 0 || ones != 0 {
		alphabetic += " and "
	}
	switch tens {
	case 1:
		switch ones {
		case 0:
			alphabetic += "ten"
		case 1:
			alphabetic += "eleven"
		case 2:
			alphabetic += "twelve"
		case 3:
			alphabetic += "thirteen"
		case 4:
			alphabetic += "fourteen"
		case 5:
			alphabetic += "fifteen"
		case 6:
			alphabetic += "sixteen"
		case 7:
			alphabetic += "seventeen"
		case 8:
			alphabetic += "eighteen"
		case 9:
			alphabetic += "nineteen"
		}
	case 2:
		alphabetic += "twenty"
	case 3:
		alphabetic += "thirty"
	case 4:
		alphabetic += "forty"
	case 5:
		alphabetic += "fifty"
	case 6:
		alphabetic += "sixty"
	case 7:
		alphabetic += "seventy"
	case 8:
		alphabetic += "eighty"
	case 9:
		alphabetic += "ninety"
	}
	if tens > 1 && ones != 0 {
		alphabetic += "-"
	}
	if tens != 1 {
		switch ones {
		case 1:
			alphabetic += "one"
		case 2:
			alphabetic += "two"
		case 3:
			alphabetic += "three"
		case 4:
			alphabetic += "four"
		case 5:
			alphabetic += "five"
		case 6:
			alphabetic += "six"
		case 7:
			alphabetic += "seven"
		case 8:
			alphabetic += "eight"
		case 9:
			alphabetic += "nine"
		}
	}
	return ct.checker.CompareLine(alphabetic)
}

// { "no": 19, "dat": "", "ans": "" }
func (ct *caseTask) case19() bool {
	var year = ct.scanner.NextInt()
	year -= 4
	if year < 0 {
		year = -year
	}
	year %= 60
	var yearName = "The "
	// colors;
	switch {
	case year < 12:
		yearName += "Green"
	case year < 24:
		yearName += "Red"
	case year < 36:
		yearName += "Yellow"
	case year < 48:
		yearName += "White"
	default:
		yearName += "Black"
	}
	yearName += " "
	// animals;
	switch year % 60 % 12 {
	case 0:
		yearName += "Rat"
	case 1:
		yearName += "Cow"
	case 2:
		yearName += "Tiger"
	case 3:
		yearName += "Hare"
	case 4:
		yearName += "Dragon"
	case 5:
		yearName += "Snake"
	case 6:
		yearName += "Horse"
	case 7:
		yearName += "Sheep"
	case 8:
		yearName += "Monkey"
	case 9:
		yearName += "Hen"
	case 10:
		yearName += "Dog"
	case 11:
		yearName += "Pig"
	}
	yearName += "'s year"
	return ct.checker.CompareLine(yearName)
}

// { "no": 20, "dat": "", "ans": "" }
func (ct *caseTask) case20() bool {
	var d, m = ct.scanner.NextInt(), ct.scanner.NextInt()
	var zodiacalName string
	switch m {
	case 1:
		switch {
		case d < 20:
			zodiacalName = "Capricorn"
		default:
			zodiacalName = "Aquarius"
		}
	case 2:
		switch {
		case d < 19:
			zodiacalName = "Aquarius"
		default:
			zodiacalName = "Pisces"
		}
	case 3:
		switch {
		case d < 21:
			zodiacalName = "Pisces"
		default:
			zodiacalName = "Aries"
		}
	case 4:
		switch {
		case d < 20:
			zodiacalName = "Aries"
		default:
			zodiacalName = "Taurus"
		}
	case 5:
		switch {
		case d < 21:
			zodiacalName = "Taurus"
		default:
			zodiacalName = "Gemini"
		}
	case 6:
		switch {
		case d < 22:
			zodiacalName = "Gemini"
		default:
			zodiacalName = "Cancer"
		}
	case 7:
		switch {
		case d < 23:
			zodiacalName = "Cancer"
		default:
			zodiacalName = "Leo"
		}
	case 8:
		switch {
		case d < 23:
			zodiacalName = "Leo"
		default:
			zodiacalName = "Virgo"
		}
	case 9:
		switch {
		case d < 23:
			zodiacalName = "Virgo"
		default:
			zodiacalName = "Libra"
		}
	case 10:
		switch {
		case d < 23:
			zodiacalName = "Libra"
		default:
			zodiacalName = "Scorpio"
		}
	case 11:
		switch {
		case d < 23:
			zodiacalName = "Scorpio"
		default:
			zodiacalName = "Sagittarius"
		}
	case 12:
		switch {
		case d < 22:
			zodiacalName = "Sagittarius"
		default:
			zodiacalName = "Capricorn"
		}
	}
	return ct.checker.CompareWord(zodiacalName)
}
