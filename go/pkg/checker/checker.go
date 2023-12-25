package checker

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/hash"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/mylog"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/node"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
	"github.com/shopspring/decimal"
)

type Checker interface {
	EOF() bool
	Skip()
	CompareWord(args ...string) bool
	CompareLine(args ...string) bool
	CompareByte(args ...byte) bool
	CompareRune(args ...rune) bool
	CompareBool(args ...bool) bool
	CompareInt(args ...int) bool
	CompareInt32(args ...int32) bool
	CompareInt64(args ...int64) bool
	CompareFloat32(precision int32, args ...float32) bool
	CompareFloat64(precision int32, args ...float64) bool
	CompareBinaryFile(endian binary.ByteOrder, fileNames ...string) bool
	CompareTextFile(fileNames ...string) bool
	CompareStack(tops ...*node.Node) bool
	CompareQueue(head, tail *node.Node) bool
	CompareList(checkFirst, checkLast, checkCurrent *node.Node) bool
	CompareBarrierList(barrier, current *node.Node) bool
	CompareTree(root *node.Node, currents ...*node.Node) bool
	RemoveFiles()
}

type checker struct {
	filePath string
	scanner  scanner.Scanner
}

// New is a constructor for making a new Checker;
func New(fileName string) (Checker, error) {
	sc, err := scanner.New(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "scanner.New")
	}
	return &checker{
		filePath: filepath.Dir(fileName),
		scanner:  sc,
	}, nil
}

// EOF checks if it is the end of the source;
func (chck *checker) EOF() bool {
	return chck.scanner.EOF()
}

// Skip skips one symbol of the source;
func (chck *checker) Skip() {
	chck.scanner.Skip()
}

// CompareWord checks the source words with argument words;
func (chck *checker) CompareWord(args ...string) bool {
	for _, arg := range args {
		if chck.EOF() {
			if len(arg) == 0 {
				continue
			}
			mylog.Log("checker", "CompareWord", "chck.EOF", chck.EOF())
			return false
		}
		value := chck.scanner.NextWord()
		if value != arg {
			mylog.Log("checker", "CompareWord", "expected:", value)
			mylog.Log("checker", "CompareWord", "have got:", arg)
			return false
		}
	}
	return true
}

// CompareLine checks the source lines with argument lines;
func (chck *checker) CompareLine(args ...string) bool {
	for _, arg := range args {
		// if chck.EOF() {
		// 	if len(arg) == 0 {
		// 		continue
		// 	}
		// 	mylog.Log("checker", "CompareLine", "chck.EOF", chck.EOF())
		// 	ok = false
		// 	return
		// }
		value := chck.scanner.NextLine()
		if value != arg {
			mylog.Log("checker", "CompareLine", "expected:", value)
			mylog.Log("checker", "CompareLine", "have got:", arg)
			return false
		}
	}
	return true
}

// CompareByte checks the source bytes with argument bytes;
func (chck *checker) CompareByte(args ...byte) bool {
	for _, arg := range args {
		value := chck.scanner.NextByte()
		if value == 0 {
			mylog.Log("checker", "CompareByte", "chck.EOF", chck.EOF())
			return true
		}
		if value != arg {
			mylog.Log("checker", "CompareByte", "expected:", string(value))
			mylog.Log("checker", "CompareByte", "have got:", string(arg))
			return false
		}
	}
	return true
}

// CompareRune checks the source runes with argument runes;
func (chck *checker) CompareRune(args ...rune) bool {
	for _, arg := range args {
		value, size := chck.scanner.NextRune()
		if size == 0 {
			mylog.Log("checker", "CompareRune", "chck.EOF", chck.EOF())
			return true
		}
		if value != arg {
			mylog.Log("checker", "CompareRune", "expected:", string(value))
			mylog.Log("checker", "CompareRune", "have got:", string(arg))
			return false
		}
	}
	return true
}

// CompareBool checks the source bools with argument bools;
func (chck *checker) CompareBool(args ...bool) bool {
	for _, arg := range args {
		if chck.EOF() {
			mylog.Log("checker", "CompareBool", "chck.EOF", chck.EOF())
			return false
		}
		value := chck.scanner.NextBool()
		if value != arg {
			mylog.Log("checker", "CompareBool", "expected:", value)
			mylog.Log("checker", "CompareBool", "have got:", arg)
			return false
		}
	}
	return true
}

// CompareInt checks the source ints with argument ints;
func (chck *checker) CompareInt(args ...int) bool {
	for _, arg := range args {
		if chck.EOF() {
			mylog.Log("checker", "CompareInt", "chck.EOF", chck.EOF())
			return false
		}
		value := chck.scanner.NextInt()
		if value != arg {
			mylog.Log("checker", "CompareInt", "expected:", value)
			mylog.Log("checker", "CompareInt", "have got:", arg)
			return false
		}
	}
	return true
}

// CompareInt32 checks the source int32s with argument int32s;
func (chck *checker) CompareInt32(args ...int32) bool {
	for _, arg := range args {
		if chck.EOF() {
			mylog.Log("checker", "CompareInt32", "chck.EOF", chck.EOF())
			return false
		}
		value := chck.scanner.NextInt32()
		if value != arg {
			mylog.Log("checker", "CompareInt32", "expected:", value)
			mylog.Log("checker", "CompareInt32", "have got:", arg)
			return false
		}
	}
	return true
}

// CompareInt64 checks the source int64s with argument int6s;
func (chck *checker) CompareInt64(args ...int64) bool {
	for _, arg := range args {
		if chck.EOF() {
			mylog.Log("checker", "CompareInt64", "chck.EOF", chck.EOF())
			return false
		}
		value := chck.scanner.NextInt64()
		if value != arg {
			mylog.Log("checker", "CompareInt64", "expected:", value)
			mylog.Log("checker", "CompareInt64", "have got:", arg)
			return false
		}
	}
	return true
}

// CompareFloat32 checks the source float32s with argument float32s;
func (chck *checker) CompareFloat32(precision int32, args ...float32) bool {
	for _, arg := range args {
		if chck.EOF() {
			mylog.Log("checker", "CompareFloat32", "chck.EOF", chck.EOF())
			return false
		}
		value := decimal.NewFromFloat32(chck.scanner.NextFloat32()).Round(precision)
		argValue := decimal.NewFromFloat32(arg).Round(precision)
		if !value.Equal(argValue) {
			mylog.Log("checker", "CompareFloat32", "expected:", value)    // fmt.Sprintf("%.*f", precision, value)
			mylog.Log("checker", "CompareFloat32", "have got:", argValue) // fmt.Sprintf("%.*f", precision, arg)
			mylog.Log("checker", "CompareFloat64", "have got:", arg)
			return false
		}
	}
	return true
}

// CompareFloat64 checks the source float64s with argument float64s;
func (chck *checker) CompareFloat64(precision int32, args ...float64) bool {
	for _, arg := range args {
		if chck.EOF() {
			mylog.Log("checker", "CompareFloat64", "chck.EOF", chck.EOF())
			return false
		}
		value := decimal.NewFromFloat(chck.scanner.NextFloat64()).Round(precision)
		argValue := decimal.NewFromFloat(arg).Round(precision)
		if !value.Equal(argValue) {
			mylog.Log("checker", "CompareFloat64", "expected:", value)
			mylog.Log("checker", "CompareFloat64", "have got:", argValue)
			mylog.Log("checker", "CompareFloat64", "have got:", arg)
			return false
		}
	}
	return true
}

// CompareBinaryFile checks the source binary file with argument binary file;
func (chck *checker) CompareBinaryFile(endian binary.ByteOrder, fileNames ...string) bool {
	for _, fileName := range fileNames {
		_, err := os.Stat(fileName)
		if err != nil {
			continue
		}
		chck.scanner.AddToRemoveQueue(fileName, fileName)
	}
	for _, fileName := range fileNames {
		_, err := os.Stat(fileName)
		if err != nil {
			continue
		}
		if chck.EOF() {
			if len(fileName) == 0 {
				continue
			}
			mylog.Log("checker", "CompareBinaryFile", "chck.EOF", chck.EOF())
			return false
		}
		value := chck.filePath + "/" + chck.scanner.NextLine()
		if endian == binary.BigEndian {
			value += ".big"
		}
		chs1, err := hash.GetCheckSum(value)
		if err != nil {
			mylog.Log("checker", "CompareBinaryFile", "hash.GetCheckSum(value)", err)
			return true
		}
		chs2, err := hash.GetCheckSum(fileName)
		if err != nil {
			mylog.Log("checker", "CompareBinaryFile", "hash.GetCheckSum(fileName)", err)
			return true
		}
		// Compare with hashed checkSums;
		if chs1 != chs2 {
			mylog.Log("checker", "CompareBinaryFile", "", fmt.Sprintf("files %s and %s have different contents", fileName, value))
			// correctContent, _ := os.ReadFile(value)
			// CompareContent, _ := os.ReadFile(fileName)
			// mylog.Log("checker", "CompareWord", "expected:", string(correctContent))
			// mylog.Log("checker", "CompareWord", "have got:", string(checkContent))
			return false
		}
	}
	return true
}

// CompareTextFile checks the source binary file with argument binary file;
func (chck *checker) CompareTextFile(fileNames ...string) bool {
	for _, fileName := range fileNames {
		chck.scanner.AddToRemoveQueue(fileName, fileName)
	}
	for _, fileName := range fileNames {
		if chck.EOF() {
			if len(fileName) == 0 {
				continue
			}
			mylog.Log("checker", "CompareTextFile", "chck.EOF", chck.EOF())
			return false
		}
		value := chck.filePath + "/" + chck.scanner.NextLine()
		correctFile, err := os.OpenFile(value, os.O_RDONLY, 0777)
		if err != nil {
			mylog.Log("checker", "CompareTextFile", "os.OpenFile(value)", err)
			return true
		}
		defer correctFile.Close()
		checkFile, err := os.OpenFile(fileName, os.O_RDONLY, 0777)
		if err != nil {
			mylog.Log("checker", "CompareTextFile", "os.OpenFile(fileName)", err)
			return true
		}
		defer checkFile.Close()
		correctScanner := bufio.NewScanner(correctFile)
		correctScanner.Split(bufio.ScanLines)
		checkScanner := bufio.NewScanner(checkFile)
		checkScanner.Split(bufio.ScanLines)
		var checkEOF, correctEOF bool
		for {
			checkEOF = !checkScanner.Scan()
			correctEOF = !correctScanner.Scan()
			if checkEOF || correctEOF {
				break
			}
			correctLine := correctScanner.Text()
			checkLine := checkScanner.Text()
			if correctLine != checkLine {
				mylog.Log("checker", "CompareTextFile", "expected:", []byte(correctLine))
				mylog.Log("checker", "CompareTextFile", "have got:", []byte(checkLine))
				return false
			}
		}
		if !correctEOF {
			mylog.Log("checker", "CompareTextFile", "correctEOF", "Correct file is not EOF")
			return false
		}
		if !checkEOF {
			mylog.Log("checker", "CompareTextFile", "checkEOF", "Check file is not EOF")
			return false
		}
	}
	return true
}

// CompareStack checks the source stacks with argument stacks;
func (chck *checker) CompareStack(tops ...*node.Node) bool {
	var correctStackTop *node.Node
	for _, top := range tops {
		if chck.EOF() {
			mylog.Log("checker", "CompareStack", "chck.EOF", chck.EOF())
			return false
		}
		correctStackTop = chck.scanner.NextStack()
		for top != nil || correctStackTop != nil {
			if top == nil {
				mylog.Log("checker", "CompareStack", "", "Check stack is empty")
				return false
			}
			if correctStackTop == nil {
				mylog.Log("checker", "CompareStack", "", "Correct stack is empty")
				return false
			}
			if top.Data != correctStackTop.Data {
				mylog.Log("checker", "CompareStack", "expected:", correctStackTop.Data)
				mylog.Log("checker", "CompareStack", "have got:", top.Data)
				return false
			}
			top = top.Next
			correctStackTop = correctStackTop.Next
		}
	}
	return true
}

// CompareQueue checks the source queues with argument queues;
func (chck *checker) CompareQueue(head, tail *node.Node) bool {
	var correctQueueHead *node.Node
	if chck.EOF() {
		mylog.Log("checker", "CompareQueue", "chck.EOF", chck.EOF())
		return false
	}
	correctQueueHead, _ = chck.scanner.NextQueue()
	for head != nil || correctQueueHead != nil {
		if head == nil {
			mylog.Log("checker", "CompareQueue", "", "Check queue is empty")
			return false
		}
		if correctQueueHead == nil {
			mylog.Log("checker", "CompareQueue", "", "Correct queue is empty")
			return false
		}
		if head.Data != correctQueueHead.Data {
			mylog.Log("checker", "CompareQueue", "expected:", correctQueueHead.Data)
			mylog.Log("checker", "CompareQueue", "have got:", head.Data)
			return false
		}
		head = head.Next
		correctQueueHead = correctQueueHead.Next
	}
	return true
}

// CompareList checks the source list with argument list;
func (chck *checker) CompareList(checkFirst, checkLast, checkCurrent *node.Node) bool {
	if chck.EOF() {
		mylog.Log("checker", "CompareList", "chck.EOF", chck.EOF())
		return false
	}
	var correctFirst, correctLast, correctCurrent = chck.scanner.NextList()
	if checkFirst == nil {
		if checkCurrent != nil {
			checkFirst = checkCurrent
			for checkFirst != nil && checkFirst.Prev != nil {
				checkFirst = checkFirst.Prev
			}
		} else {
			checkFirst = checkLast
			for checkFirst != nil && checkFirst.Prev != nil && checkFirst.Prev != checkLast {
				checkFirst = checkFirst.Prev
			}
		}
	}
	if checkLast == nil {
		if checkCurrent != nil {
			checkLast = checkCurrent
			for checkLast != nil && checkLast.Next != nil {
				checkLast = checkLast.Next
			}
		} else {
			checkLast = checkFirst
			for checkLast != nil && checkLast.Next != nil && checkLast.Next != checkFirst {
				checkLast = checkLast.Next
			}
		}
	}
	if !checkListAsc(checkFirst, checkLast, correctFirst, correctLast) &&
		checkListDesc(checkFirst, checkLast, correctFirst, correctLast) {
		return false
	}
	if correctCurrent != nil && checkCurrent.Data != correctCurrent.Data {
		mylog.Log("checker", "CompareList", "Current expected:", correctCurrent.Data)
		mylog.Log("checker", "CompareList", "Current have got:", checkCurrent.Data)
		return false
	}
	return true
}

func checkListAsc(checkFirst, checkLast, correctFirst, correctLast *node.Node) bool {
	for checkFirst != nil || correctFirst != nil {
		if checkFirst == nil {
			mylog.Log("checker", "checkListAsc", "", "check list is empty")
			return false
		}
		if correctFirst == nil {
			mylog.Log("checker", "checkListAsc", "", "correct list is empty")
			return false
		}
		if correctFirst.Data != checkFirst.Data {
			mylog.Log("checker", "checkListAsc", "expected", correctFirst.Data)
			mylog.Log("checker", "checkListAsc", "have got", checkFirst.Data)
			return false
		}
		if checkFirst != checkLast {
			checkFirst = checkFirst.Next
		} else {
			checkFirst = nil
		}
		if correctFirst != correctLast {
			correctFirst = correctFirst.Next
		} else {
			correctFirst = nil
		}
	}
	return true
}

func checkListDesc(checkFirst, checkLast, correctFirst, correctLast *node.Node) bool {
	for checkLast != nil || correctLast != nil {
		if checkLast == nil {
			mylog.Log("checker", "checkListDesc", "", "check list is empty")
			return false
		}
		if correctLast == nil {
			mylog.Log("checker", "checkListDesc", "", "correct list is empty")
			return false
		}
		if checkLast.Data != correctLast.Data {
			mylog.Log("checker", "checkListDesc", "expected:", correctLast.Data)
			mylog.Log("checker", "checkListDesc", "have got:", checkLast.Data)
			return false
		}
		if checkLast != checkFirst {
			checkLast = checkLast.Prev
		} else {
			checkLast = nil
		}
		if correctLast != correctFirst {
			correctLast = correctLast.Prev
		} else {
			correctLast = nil
		}
	}
	return true
}

// CompareBarrierList checks the source barrier list with argument barrier list;
func (chck *checker) CompareBarrierList(barrier, current *node.Node) bool {
	first, last := barrier.Next, barrier.Prev
	if first == last && first == barrier {
		first, last, current = nil, nil, nil
	}
	return chck.CompareList(first, last, current)
}

// CompareTree checks the source barrier list with argument barrier list;
func (chck *checker) CompareTree(root *node.Node, currents ...*node.Node) bool {
	if root == nil {
		mylog.Log("checker", "CompareTree", "", "Root is nil")
		return false
	}
	var correctRoot = chck.scanner.NextTree()
	if !checkTree(root, correctRoot) {
		return false
	}
	for _, current := range currents {
		var correctCurrent = correctRoot
	MOVING:
		for correctCurrent != nil {
			switch chck.scanner.NextInt() {
			case 0:
				break MOVING
			case 1:
				correctCurrent = correctCurrent.Left
			case 2:
				correctCurrent = correctCurrent.Right
			}
		}
		if current == nil {
			mylog.Log("checker", "CompareTree", "", "Check node doesn't exist")
			return false
		} else if correctCurrent == nil {
			mylog.Log("checker", "CompareTree", "", "Correct node doesn't exist")
			return false
		} else if current.Data != correctCurrent.Data {
			mylog.Log("checker", "CompareTree", "expected:", correctCurrent.Data)
			mylog.Log("checker", "CompareTree", "have got:", current.Data)
			return false
		}
	}
	return true
}

func checkTree(checkNode, correctNode *node.Node) bool {
	if checkNode == nil && correctNode == nil {
		return true
	} else if checkNode == nil {
		mylog.Log("checker", "checkTree", "", "Check Tree is not completed")
		return false
	} else if correctNode == nil {
		mylog.Log("checker", "checkTree", "", "Correct Tree is not completed")
		return false
	}
	return checkNode.Data == correctNode.Data &&
		checkTree(checkNode.Left, correctNode.Left) &&
		checkTree(checkNode.Right, correctNode.Right)
}

func (chck *checker) RemoveFiles() {
	chck.scanner.RemoveFiles()
}
