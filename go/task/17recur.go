package task

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/proc"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type recurTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewRecurTask(pathPrefix string) Task {
	return &recurTask{
		dataPath: pathPrefix + "/17recur/Recur%03d/%03d/%03d",
		name:     "Recur",
		count:    30,
	}
}

func (rt *recurTask) GetCount() int { return rt.count }

func (rt *recurTask) GetName() string { return rt.name }

func (rt *recurTask) SetData(taskNo int, testNo int) error {
	var filePath = fmt.Sprintf(rt.dataPath, taskNo, testNo, testNo)
	scanner, err := scanner.New(filePath + ".dat")
	if err != nil {
		return errors.Wrap(err, "scanner.New")
	}
	checker, err := checker.New(filePath + ".ans")
	if err != nil {
		return errors.Wrap(err, "checker.New")
	}
	rt.scanner = scanner
	rt.checker = checker
	return nil
}

func (rt *recurTask) ScannerEOF() bool { return rt.scanner.EOF() }

func (rt *recurTask) CheckerEOF() bool { return rt.checker.EOF() }

func (rt *recurTask) CleanData() {
	rt.scanner.RemoveFiles()
	rt.checker.RemoveFiles()
}

func (rt *recurTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return rt.recur1
	case 2:
		return rt.recur2
	case 3:
		return rt.recur3
	case 4:
		return rt.recur4
	case 5:
		return rt.recur5
	case 6:
		return rt.recur6
	case 7:
		return rt.recur7
	case 8:
		return rt.recur8
	case 9:
		return rt.recur9
	case 10:
		return rt.recur10
	case 11:
		return rt.recur11
	case 12:
		return rt.recur12
	case 13:
		return rt.recur13
	case 14:
		return rt.recur14
	case 15:
		return rt.recur15
	case 16:
		return rt.recur16
	case 17:
		return rt.recur17
	case 18:
		return rt.recur18
	case 19:
		return rt.recur19
	case 20:
		return rt.recur20
	case 21:
		return rt.recur21
	case 22:
		return rt.recur22
	case 23:
		return rt.recur23
	case 24:
		return rt.recur24
	case 25:
		return rt.recur25
	case 26:
		return rt.recur26
	case 27:
		return rt.recur27
	case 28:
		return rt.recur28
	case 29:
		return rt.recur29
	case 30:
		return rt.recur30
	default:
		return nil
	}
}

// { "no": 1, "dat": "", "ans": "1" }
func (rt *recurTask) recur1() bool {
	for index := 0; index < 5; index++ {
		var number = rt.scanner.NextInt()
		var factorial = proc.Fact(number)
		if !rt.checker.CompareFloat64(1, factorial) {
			return false
		}
	}
	return true
}

// { "no": 2, "dat": "", "ans": "1" }
func (rt *recurTask) recur2() bool {
	return false
}

// { "no": 3, "dat": "", "ans": "E (OR 15)" }
func (rt *recurTask) recur3() bool {
	return false
}

// { "no": 4, "dat": "", "ans": "" }
func (rt *recurTask) recur4() bool {
	return false
}

// { "no": 5, "dat": "", "ans": "" }
func (rt *recurTask) recur5() bool {
	return false
}

// { "no": 6, "dat": "", "ans": "" }
func (rt *recurTask) recur6() bool {
	return false
}

// { "no": 7, "dat": "", "ans": "" }
func (rt *recurTask) recur7() bool {
	return false
}

// { "no": 8, "dat": "2", "ans": "8" }
func (rt *recurTask) recur8() bool {
	return false
}

// { "no": 9, "dat": "", "ans": "" }
func (rt *recurTask) recur9() bool {
	return false
}

// { "no": 10, "dat": "", "ans": "" }
func (rt *recurTask) recur10() bool {
	return false
}

// { "no": 11, "dat": "", "ans": "" }
func (rt *recurTask) recur11() bool {
	return false
}

// { "no": 12, "dat": "", "ans": "" }
func (rt *recurTask) recur12() bool {
	return false
}

// { "no": 13, "dat": "", "ans": "" }
func (rt *recurTask) recur13() bool {
	return false
}

// { "no": 14, "dat": "", "ans": "" }
func (rt *recurTask) recur14() bool {
	return false
}

// { "no": 15, "dat": "", "ans": "" }
func (rt *recurTask) recur15() bool {
	return false
}

// { "no": 16, "dat": "", "ans": "" }
func (rt *recurTask) recur16() bool {
	return false
}

// { "no": 17, "dat": "", "ans": "" }
func (rt *recurTask) recur17() bool {
	return false
}

// { "no": 18, "dat": "", "ans": "" }
func (rt *recurTask) recur18() bool {
	return false
}

// { "no": 19, "dat": "", "ans": "" }
func (rt *recurTask) recur19() bool {
	return false
}

// { "no": 20, "dat": "", "ans": "" }
func (rt *recurTask) recur20() bool {
	return false
}

// { "no": 21, "dat": "", "ans": "" }
func (rt *recurTask) recur21() bool {
	return false
}

// { "no": 22, "dat": "", "ans": "" }
func (rt *recurTask) recur22() bool {
	return false
}

// { "no": 23, "dat": "", "ans": "" }
func (rt *recurTask) recur23() bool {
	return false
}

// { "no": 24, "dat": "", "ans": "" }
func (rt *recurTask) recur24() bool {
	return false
}

// { "no": 25, "dat": "", "ans": "" }
func (rt *recurTask) recur25() bool {
	return false
}

// { "no": 26, "dat": "", "ans": "" }
func (rt *recurTask) recur26() bool {
	return false
}

// { "no": 27, "dat": "", "ans": "" }
func (rt *recurTask) recur27() bool {
	return false
}

// { "no": 28, "dat": "", "ans": "" }
func (rt *recurTask) recur28() bool {
	return false
}

// { "no": 29, "dat": "", "ans": "" }
func (rt *recurTask) recur29() bool {
	return false
}

// { "no": 30, "dat": "", "ans": "" }
func (rt *recurTask) recur30() bool {
	return false
}
