package task

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/mylog"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type textTask struct {
	dataPath  string
	name      string
	count     int
	byteOrder binary.ByteOrder
	scanner   scanner.Scanner
	checker   checker.Checker
}

func NewTextTask(pathPrefix string, byteOrder binary.ByteOrder) Task {
	return &textTask{
		dataPath:  pathPrefix + "/15text/Text%03d/%03d/%03d",
		name:      "Text",
		count:     60,
		byteOrder: byteOrder,
	}
}

func (tt *textTask) GetCount() int { return tt.count }

func (tt *textTask) GetName() string { return tt.name }

func (tt *textTask) SetData(taskNo int, testNo int) error {
	var filePath = fmt.Sprintf(tt.dataPath, taskNo, testNo, testNo)
	scanner, err := scanner.New(filePath + ".dat")
	if err != nil {
		return errors.Wrap(err, "scanner.New")
	}
	checker, err := checker.New(filePath + ".ans")
	if err != nil {
		return errors.Wrap(err, "checker.New")
	}
	tt.scanner = scanner
	tt.checker = checker
	return nil
}

func (tt *textTask) ScannerEOF() bool { return tt.scanner.EOF() }

func (tt *textTask) CheckerEOF() bool { return tt.checker.EOF() }

func (tt *textTask) CleanData() {
	tt.scanner.RemoveFiles()
	tt.checker.RemoveFiles()
}

func (tt *textTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return tt.text1
	case 2:
		return tt.text2
	case 3:
		return tt.text3
	case 4:
		return tt.text4
	case 5:
		return tt.text5
	case 6:
		return tt.text6
	case 7:
		return tt.text7
	case 8:
		return tt.text8
	case 9:
		return tt.text9
	case 10:
		return tt.text10
	case 11:
		return tt.text11
	case 12:
		return tt.text12
	case 13:
		return tt.text13
	case 14:
		return tt.text14
	case 15:
		return tt.text15
	case 16:
		return tt.text16
	case 17:
		return tt.text17
	case 18:
		return tt.text18
	case 19:
		return tt.text19
	case 20:
		return tt.text20
	case 21:
		return tt.text21
	case 22:
		return tt.text22
	case 23:
		return tt.text23
	case 24:
		return tt.text24
	case 25:
		return tt.text25
	case 26:
		return tt.text26
	case 27:
		return tt.text27
	case 28:
		return tt.text28
	case 29:
		return tt.text29
	case 30:
		return tt.text30
	case 31:
		return tt.text31
	case 32:
		return tt.text32
	case 33:
		return tt.text33
	case 34:
		return tt.text34
	case 35:
		return tt.text35
	case 36:
		return tt.text36
	case 37:
		return tt.text37
	case 38:
		return tt.text38
	case 39:
		return tt.text39
	case 40:
		return tt.text40
	case 41:
		return tt.text41
	case 42:
		return tt.text42
	case 43:
		return tt.text43
	case 44:
		return tt.text44
	case 45:
		return tt.text45
	case 46:
		return tt.text46
	case 47:
		return tt.text47
	case 48:
		return tt.text48
	case 49:
		return tt.text49
	case 50:
		return tt.text50
	case 51:
		return tt.text51
	case 52:
		return tt.text52
	case 53:
		return tt.text53
	case 54:
		return tt.text54
	case 55:
		return tt.text55
	case 56:
		return tt.text56
	case 57:
		return tt.text57
	case 58:
		return tt.text58
	case 59:
		return tt.text59
	case 60:
		return tt.text60
	default:
		return nil
	}
}

func (tt *textTask) text1() bool {
	var fileName = tt.scanner.NextLine()
	var n = tt.scanner.NextInt()
	var k = tt.scanner.NextInt()
	var file, err = os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text001", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	for index := 0; index < n; index++ {
		_, err = file.WriteString(strings.Repeat("*", k) + "\n")
		if err != nil {
			mylog.Log("task", "text001", "file.WriteString", err)
			return false
		}
	}
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text2() bool {
	var fileName = tt.scanner.NextLine()
	var n = tt.scanner.NextInt()
	var file, err = os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text002", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var letter = 'a'
	var str string
	for index := 0; index < n; index++ {
		str += string(letter)
		_, err = file.WriteString(str + "\n")
		if err != nil {
			mylog.Log("task", "text002", "file.WriteString", err)
			return false
		}
		letter++
	}
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text3() bool {
	var fileName = tt.scanner.NextLine()
	var n = tt.scanner.NextInt()
	var file, err = os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text003", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var letter = 'A'
	var data = []rune(strings.Repeat("*", n))
	for index := 0; index < n; index++ {
		data[index] = letter
		_, err = file.WriteString(string(data) + "\n")
		if err != nil {
			mylog.Log("task", "text003", "file.WriteString", err)
			return false
		}
		letter++
	}
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text4() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text004", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var charactersCount, linesCount = 0, 0
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		charactersCount += len(line)
		linesCount++
	}
	return tt.checker.CompareInt(charactersCount, linesCount)
}

func (tt *textTask) text5() bool {
	var s = tt.scanner.NextLine()
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text005", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	file.Seek(0, 2)
	_, err = file.WriteString(s)
	if err != nil {
		mylog.Log("task", "text005", "file.WriteString", err)
		return false
	}
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text6() bool {
	var fileName1 = tt.scanner.NextTextFileName()
	var fileName2 = tt.scanner.NextTextFileName()
	var file1, file2 *os.File
	var err error
	file1, err = os.OpenFile(fileName1, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text006", "os.OpenFile(fileName1)", err)
		return false
	}
	defer file1.Close()
	file2, err = os.OpenFile(fileName2, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text006", "os.OpenFile(fileName2)", err)
		return false
	}
	defer file2.Close()
	file1.Seek(0, 2)
	_, err = io.Copy(file1, file2)
	if err != nil {
		mylog.Log("task", "text006", "io.Copy", err)
		return false
	}
	return tt.checker.CompareTextFile(fileName1)
}

func (tt *textTask) text7() bool {
	var s = tt.scanner.NextLine()
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text007", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text07.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text007", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	_, err = tempFile.WriteString(s + "\n")
	if err != nil {
		mylog.Log("task", "text007", "tempFile.WriteString", err)
		return false
	}
	_, err = io.Copy(tempFile, file)
	if err != nil {
		mylog.Log("task", "text007", "io.Copy", err)
		return false
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text8() bool {
	var fileName1 = tt.scanner.NextTextFileName()
	var fileName2 = tt.scanner.NextTextFileName()
	var file1, file2 *os.File
	var err error
	file1, err = os.OpenFile(fileName1, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text008", "os.OpenFile(fileName1)", err)
		return false
	}
	defer file1.Close()
	file2, err = os.OpenFile(fileName2, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text008", "os.OpenFile(fileName2)", err)
		return false
	}
	defer file2.Close()
	var tempFileName = "$tempFileName_text08.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text008", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	_, err = io.Copy(tempFile, file2)
	if err != nil {
		mylog.Log("task", "text008", "io.Copy(tempFile, file2)", err)
		return false
	}
	_, err = io.Copy(tempFile, file1)
	if err != nil {
		mylog.Log("task", "text008", "io.Copy(tempFile, file1)", err)
		return false
	}
	os.Remove(fileName1)
	os.Rename(tempFileName, fileName1)
	return tt.checker.CompareTextFile(fileName1)
}

func (tt *textTask) text9() bool {
	var k = tt.scanner.NextInt()
	tt.scanner.Skip()
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text009", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text09.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text009", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var line string
	var index = 1
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line = scanner.Text()
		if index == k {
			_, err = tempFile.Write([]byte{'\n'})
			if err != nil {
				mylog.Log("task", "text009", "tempFile.WriteString1", err)
				return false
			}
		}
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text009", "tempFile.WriteString2", err)
			return false
		}
		index++
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text10() bool {
	var k = tt.scanner.NextInt()
	tt.scanner.Skip()
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text010", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text09.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text010", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var line string
	var index = 1
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line = scanner.Text()
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text010", "tempFile.WriteString2", err)
			return false
		}
		if index == k {
			_, err = tempFile.Write([]byte{'\n'})
			if err != nil {
				mylog.Log("task", "text010", "tempFile.WriteString1", err)
				return false
			}
		}
		index++
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text11() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text011", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text11.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text011", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			line = "\n"
		}
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text011", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text12() bool {
	var s = tt.scanner.NextLine()
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text012", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text11.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text012", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			line = s
		}
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text012", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text13() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text013", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text13.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text013", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	var index = 0
	for scanner.Scan() {
		line = scanner.Text()
		index++
		if index > 1 {
			_, err = tempFile.WriteString(line + "\n")
			if err != nil {
				mylog.Log("task", "text013", "tempFile.WriteString", err)
				return false
			}
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text14() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text014", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text13.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text014", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var prevLine, currLine string
	var index = 0
	for scanner.Scan() {
		currLine = scanner.Text()
		index++
		if index > 1 {
			_, err = tempFile.WriteString(prevLine + "\n")
			if err != nil {
				mylog.Log("task", "text014", "tempFile.WriteString", err)
				return false
			}
		}
		prevLine = currLine
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text15() bool {
	var k = tt.scanner.NextInt()
	tt.scanner.Skip()
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text015", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text13.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text015", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	var index = 0
	for scanner.Scan() {
		line = scanner.Text()
		index++
		if index == k {
			continue
		}
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text015", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text16() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text016", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text16.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text016", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			continue
		}
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text016", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text17() bool {
	var fileName1 = tt.scanner.NextTextFileName()
	var fileName2 = tt.scanner.NextTextFileName()
	var file1, file2, tempFile *os.File
	var err error
	file1, err = os.OpenFile(fileName1, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text017", "os.OpenFile(fileName1)", err)
		return false
	}
	defer file1.Close()
	file2, err = os.OpenFile(fileName2, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text017", "os.OpenFile(fileName2)", err)
		return false
	}
	defer file2.Close()
	var tempFileName = "$tempFileName_text17.tmp$"
	tempFile, err = os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text017", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner1 = bufio.NewScanner(file1)
	scanner1.Split(bufio.ScanLines)
	var scanner2 = bufio.NewScanner(file2)
	scanner2.Split(bufio.ScanLines)
	var line1, line2 string
	for scanner1.Scan() {
		line1 = scanner1.Text()
		line2 = ""
		if scanner2.Scan() {
			line2 = scanner2.Text()
		}
		_, err = tempFile.WriteString(line1 + line2 + "\n")
		if err != nil {
			mylog.Log("task", "text017", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName1)
	os.Rename(tempFileName, fileName1)
	return tt.checker.CompareTextFile(fileName1)
}

func (tt *textTask) text18() bool {
	var k = tt.scanner.NextInt()
	tt.scanner.Skip()
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text018", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text18.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text018", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) < k {
			line = ""
		} else {
			line = line[k:]
		}
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text018", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text19() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "text019", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text19.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text019", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		line = strings.Map(func(r rune) rune {
			if unicode.Is(unicode.Latin, r) {
				if unicode.IsUpper(r) {
					r = unicode.ToLower(r)
				} else if unicode.IsLower(r) {
					r = unicode.ToUpper(r)
				}
			}
			return r
		}, line)
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text20() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text020", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text20.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text020", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	var words []string
	for scanner.Scan() {
		line = scanner.Text()
		words = strings.Fields(line)
		line = strings.Join(words, " ")
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text020", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text21() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text021", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text21.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text021", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	const length = 4
	var lines = make([]string, length)
	var index, linesCount = 0, 0
	for scanner.Scan() {
		lines[index] = scanner.Text()
		index = (index + 1) % length
		linesCount++
		if linesCount >= length {
			_, err = tempFile.WriteString(lines[index] + "\n")
			if err != nil {
				mylog.Log("task", "text021", "tempFile.WriteString", err)
				return false
			}
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text22() bool {
	var k = tt.scanner.NextInt()
	tt.scanner.Skip()
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text022", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text22.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text022", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var length = k + 1
	var lines = make([]string, length)
	var index, linesCount = 0, 0
	for scanner.Scan() {
		lines[index] = scanner.Text()
		index = (index + 1) % length
		linesCount++
		if linesCount >= length {
			_, err = tempFile.WriteString(lines[index] + "\n")
			if err != nil {
				mylog.Log("task", "text022", "tempFile.WriteString", err)
				return false
			}
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text23() bool {
	var k = tt.scanner.NextInt()
	tt.scanner.Skip()
	var inFileName = tt.scanner.NextTextFileName()
	var outFileName = tt.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text023", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text023", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var scanner = bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var lines = make([]string, k)
	var index, lastIndex = 0, -1
	for scanner.Scan() {
		lines[index] = scanner.Text()
		lastIndex = index
		index = (index + 1) % k
	}
	if lastIndex >= 0 {
		for {
			_, err = outFile.WriteString(lines[index] + "\n")
			if err != nil {
				mylog.Log("task", "text023", "outFile.WriteString", err)
				return false
			}
			if index == lastIndex {
				break
			}
			index = (index + 1) % k
		}
	}
	return tt.checker.CompareTextFile(outFileName)
}

func (tt *textTask) text24() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text024", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var prevLine, currLine string
	var paragraphsCount = 0
	for scanner.Scan() {
		currLine = scanner.Text()
		if len(currLine) != 0 {
			if len(prevLine) == 0 {
				paragraphsCount++
			}
		}
		prevLine = currLine
	}
	return tt.checker.CompareInt(paragraphsCount)
}

func (tt *textTask) text25() bool {
	var k = tt.scanner.NextInt()
	tt.scanner.Skip()
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text025", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text25.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text025", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var prevLine, currLine string
	var paragraphsCount = 0
	for scanner.Scan() {
		currLine = scanner.Text()
		if len(currLine) != 0 {
			if len(prevLine) == 0 {
				paragraphsCount++
			}
		}
		if len(currLine) == 0 || paragraphsCount != k {
			_, err = tempFile.WriteString(currLine + "\n")
			if err != nil {
				mylog.Log("task", "text025", "tempFile.WriteString", err)
				return false
			}
		}
		prevLine = currLine
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text26() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text026", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	var paragraphsCount = 0
	for scanner.Scan() {
		line = scanner.Text()
		if strings.Index(line, "     ") == 0 {
			paragraphsCount++
		}
	}
	return tt.checker.CompareInt(paragraphsCount)
}

func (tt *textTask) text27() bool {
	var k = tt.scanner.NextInt()
	tt.scanner.Skip()
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text027", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text27.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text027", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	var paragraphsCount = 0
	for scanner.Scan() {
		line = scanner.Text()
		if strings.Index(line, "     ") == 0 {
			paragraphsCount++
		}
		if len(line) != 0 && paragraphsCount == k {
			continue
		}
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text027", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text28() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text028", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text28.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text028", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	var linesCount = 0
	for scanner.Scan() {
		line = scanner.Text()
		linesCount++
		if linesCount > 1 && strings.Index(line, "     ") == 0 {
			line = "\n" + line
		}
		line += "\n"
		_, err = tempFile.WriteString(line)
		if err != nil {
			mylog.Log("task", "text028", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text29() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text029", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var word, longWord string
	for scanner.Scan() {
		word = scanner.Text()
		if len(word) > len(longWord) {
			longWord = word
		}
	}
	return tt.checker.CompareWord(longWord)
}

func (tt *textTask) text30() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text030", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var word, shortWord string
	for scanner.Scan() {
		word = scanner.Text()
		if len(word) > 0 && (len(shortWord) == 0 || len(word) <= len(shortWord)) {
			shortWord = word
		}
	}
	return tt.checker.CompareWord(shortWord)
}

func (tt *textTask) text31() bool {
	var k = tt.scanner.NextInt()
	tt.scanner.Skip()
	var textFileName = tt.scanner.NextTextFileName()
	var textFile, err = os.OpenFile(textFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text031", "os.OpenFile(textFileName)", err)
		return false
	}
	defer textFile.Close()
	var binaryFileName = tt.scanner.NextLine()
	binaryFile, err := os.OpenFile(binaryFileName, os.O_CREATE|os.O_TRUNC|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text031", "os.OpenFile(binaryFileName)", err)
		return false
	}
	defer binaryFile.Close()
	var scanner = bufio.NewScanner(textFile)
	scanner.Split(bufio.ScanWords)
	var word string
	var words []string
	for scanner.Scan() {
		word = scanner.Text()
		word = strings.TrimFunc(word, func(r rune) bool { return unicode.IsPunct(r) })
		words = strings.FieldsFunc(word, func(r rune) bool { return r == '-' || r == '.' })
		for _, w := range words {
			if utf8.RuneCountInString(w) == k {
				var data = []byte(w)
				data = append(data, make([]byte, 80-k)...)
				data[k] = 0
				err = binary.Write(binaryFile, tt.byteOrder, data)
				if err != nil {
					mylog.Log("task", "text031", "binary.Write", err)
					return false
				}
			}
		}
	}
	return tt.checker.CompareBinaryFile(tt.byteOrder, binaryFileName)
}

func (tt *textTask) text32() bool {
	var c, _ = tt.scanner.NextRune()
	tt.scanner.Skip()
	var textFileName = tt.scanner.NextTextFileName()
	var binaryFileName = tt.scanner.NextLine()
	var textFile, binaryFile *os.File
	var err error
	textFile, err = os.OpenFile(textFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text032", "os.OpenFile(textFileName)", err)
		return false
	}
	defer textFile.Close()
	binaryFile, err = os.OpenFile(binaryFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text032", "os.OpenFile(binaryFileName)", err)
		return false
	}
	defer binaryFile.Close()
	var lower rune
	if unicode.IsUpper(c) {
		lower = unicode.ToLower(c)
	} else if unicode.IsLower(c) {
		lower = unicode.ToUpper(c)
	}
	var scanner = bufio.NewScanner(textFile)
	scanner.Split(bufio.ScanWords)
	var word string
	var words []string
	var startsWith bool
	for scanner.Scan() {
		word = scanner.Text()
		word = strings.TrimFunc(word, func(r rune) bool { return unicode.IsPunct(r) })
		words = strings.FieldsFunc(word, func(r rune) bool { return r == '-' || r == '.' })
		for _, w := range words {
			startsWith = false
			for _, r := range w {
				startsWith = r == c || r == lower
				break
			}
			if startsWith {
				var data = []byte(w)
				data = append(data, make([]byte, 80-len(w))...)
				data[len(w)] = 0
				err = binary.Write(binaryFile, tt.byteOrder, data)
				if err != nil {
					mylog.Log("task", "text032", "binary.Write", err)
					return false
				}
			}
		}
	}
	return tt.checker.CompareBinaryFile(tt.byteOrder, binaryFileName)
}

func (tt *textTask) text33() bool {
	var c, _ = tt.scanner.NextRune()
	tt.scanner.Skip()
	var textFileName = tt.scanner.NextTextFileName()
	var binaryFileName = tt.scanner.NextLine()
	var textFile, binaryFile *os.File
	var err error
	textFile, err = os.OpenFile(textFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text033", "os.OpenFile(textFileName)", err)
		return false
	}
	defer textFile.Close()
	binaryFile, err = os.OpenFile(binaryFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text033", "os.OpenFile(binaryFileName)", err)
		return false
	}
	defer binaryFile.Close()
	if unicode.Is(unicode.Latin, c) && unicode.IsLower(c) {
		var scanner = bufio.NewScanner(textFile)
		scanner.Split(bufio.ScanWords)
		var word string
		var words []string
		for scanner.Scan() {
			word = strings.TrimFunc(scanner.Text(), func(r rune) bool { return unicode.IsPunct(r) })
			words = strings.FieldsFunc(word, func(r rune) bool { return r == '-' || r == '.' })
			for _, w := range words {
				if strings.ContainsRune(strings.ToLower(w), c) {
					var data = []byte(w)
					data = append(data, make([]byte, 80-len(w))...)
					data[len(w)] = 0
					err = binary.Write(binaryFile, tt.byteOrder, data)
					if err != nil {
						mylog.Log("task", "text033", "binary.Write", err)
						return false
					}
				}
			}
		}
	}
	return tt.checker.CompareBinaryFile(tt.byteOrder, binaryFileName)
}

func (tt *textTask) text34() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text034", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text34.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text034", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	var length int
	for scanner.Scan() {
		line = scanner.Text()
		length = len(line)
		if length > 0 && length <= 50 {
			line = strings.Repeat(" ", 50-length) + line
		}
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text034", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text35() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text035", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text34.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text035", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	var length int
	for scanner.Scan() {
		line = scanner.Text()
		length = len(line)
		if length > 0 && length <= 50 {
			line = strings.Repeat(" ", (50-length)/2+length%2) + line
		}
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text035", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text36() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text036", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text34.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text036", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	var length, spacesCount int
	for scanner.Scan() {
		line = scanner.Text()
		length = len(line)
		if length > 0 && length <= 50 {
			spacesCount = 0
			for _, r := range line {
				if !unicode.IsSpace(r) {
					break
				}
				spacesCount++
			}
			line = line[spacesCount/2+spacesCount%2:]
		}
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text036", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text37() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text037", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_text37.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text037", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var prev, curr string
	const LENGTH = 50
	var index = 0
	for scanner.Scan() {
		curr = scanner.Text()
		index++
		if index > 1 {
			if len(prev) > 0 && len(curr) > 0 {
				for len(prev) < LENGTH {
					var position = len(prev)
					for position >= 0 && len(prev) < LENGTH {
						position = strings.LastIndex(prev[:position], " ")
						if position < 0 {
							break
						}
						prev = prev[:position+1] + prev[position:]
						for position >= 0 && prev[position] == ' ' {
							position--
						}
					}
				}
			}
			_, err = tempFile.WriteString(prev + "\n")
			if err != nil {
				mylog.Log("task", "text037", "tempFile.WriteString", err)
				return false
			}
		}
		prev = curr
	}
	_, err = tempFile.WriteString(prev + "\n")
	if err != nil {
		mylog.Log("task", "text037", "tempFile.WriteString", err)
		return false
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text38() bool {
	var k = tt.scanner.NextInt()
	tt.scanner.Skip()
	var inFileName = tt.scanner.NextTextFileName()
	var outFileName = tt.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text038", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text038", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var scanner = bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var line, remained string
	var position int
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			if len(remained) != 0 {
				remained += "\n"
			}
			_, err = outFile.WriteString(remained + "\n")
			if err != nil {
				mylog.Log("task", "text038", "outFile.WriteString1", err)
				return false
			}
			remained = ""
		} else {
			if len(remained) != 0 {
				remained += " "
			}
			remained += line
			var length = len(remained)
			for length >= k {
				if length > k {
					position = strings.LastIndexByte(remained[:k+1], ' ')
					if position != k {
						position = -1
					}
				} else if length == k {
					position = k
				}
				if position < 1 {
					position = strings.LastIndexByte(remained[:k], ' ')
				}
				line = remained[:position]
				_, err = outFile.WriteString(line + "\n")
				if err != nil {
					mylog.Log("task", "text038", "outFile.WriteString2", err)
					return false
				}
				if len(remained) > position {
					position++
				}
				remained = remained[position:]
				length = len(remained)
				position = -1
			}
		}
	}
	if len(remained) != 0 {
		_, err = outFile.WriteString(remained + "\n")
		if err != nil {
			mylog.Log("task", "text038", "outFile.WriteString3", err)
			return false
		}
	}
	return tt.checker.CompareTextFile(outFileName)
}

func (tt *textTask) text39() bool {
	var k = tt.scanner.NextInt()
	tt.scanner.Skip()
	var inFileName = tt.scanner.NextTextFileName()
	var outFileName = tt.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text039", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text039", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var scanner = bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var line, remained string
	var position int
	for scanner.Scan() {
		line = scanner.Text()
		if strings.Index(line, "     ") == 0 && len(remained) != 0 {
			_, err = outFile.WriteString(remained + "\n")
			if err != nil {
				mylog.Log("task", "text039", "outFile.WriteString1", err)
				return false
			}
			remained = ""
		}
		if len(remained) != 0 {
			remained += " "
		}
		remained += line
		var length = len(remained)
		for length >= k {
			if length > k {
				position = strings.LastIndexByte(remained[:k+1], ' ')
				if position != k {
					position = -1
				}
			} else if length == k {
				position = k
			}
			if position < 1 {
				position = strings.LastIndexByte(remained[:k], ' ')
			}
			line = remained[:position]
			_, err = outFile.WriteString(line + "\n")
			if err != nil {
				mylog.Log("task", "text039", "outFile.WriteString2", err)
				return false
			}
			if len(remained) > position {
				position++
			}
			remained = remained[position:]
			length = len(remained)
			position = -1
		}
	}
	if len(remained) != 0 {
		_, err = outFile.WriteString(remained + "\n")
		if err != nil {
			mylog.Log("task", "text039", "outFile.WriteString3", err)
			return false
		}
	}
	return tt.checker.CompareTextFile(outFileName)
}

func (tt *textTask) text40() bool {
	var inFileName1 = tt.scanner.NextBinaryFileName(tt.byteOrder)
	var inFileName2 = tt.scanner.NextBinaryFileName(tt.byteOrder)
	var outFileName = tt.scanner.NextLine()
	var inFile1, inFile2, outFile *os.File
	var err error
	inFile1, err = os.OpenFile(inFileName1, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text040", "os.OpenFile(inFileName1)", err)
		return false
	}
	defer inFile1.Close()
	inFile2, err = os.OpenFile(inFileName2, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text040", "os.OpenFile(inFileName2)", err)
		return false
	}
	defer inFile2.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text040", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var data1, data2 int32
	for {
		err = binary.Read(inFile1, tt.byteOrder, &data1)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "text040", "binary.Read(inFile1)", err)
			return false
		}
		err = binary.Read(inFile2, tt.byteOrder, &data2)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "text040", "binary.Read(inFile2)", err)
			return false
		}
		_, err = outFile.WriteString(fmt.Sprintf("|% 30d% 30d|\n", data1, data2))
		if err != nil {
			mylog.Log("task", "text040", "outFile.WriteString", err)
			return false
		}
	}
	return tt.checker.CompareTextFile(outFileName)
}

func (tt *textTask) text41() bool {
	var inFileName1 = tt.scanner.NextBinaryFileName(tt.byteOrder)
	var inFileName2 = tt.scanner.NextBinaryFileName(tt.byteOrder)
	var inFileName3 = tt.scanner.NextBinaryFileName(tt.byteOrder)
	var outFileName = tt.scanner.NextLine()
	var inFile1, inFile2, inFile3, outFile *os.File
	var err error
	inFile1, err = os.OpenFile(inFileName1, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text041", "os.OpenFile(inFileName1)", err)
		return false
	}
	defer inFile1.Close()
	inFile2, err = os.OpenFile(inFileName2, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text041", "os.OpenFile(inFileName2)", err)
		return false
	}
	defer inFile2.Close()
	inFile3, err = os.OpenFile(inFileName3, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text041", "os.OpenFile(inFileName3)", err)
		return false
	}
	defer inFile3.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text041", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var data1, data2, data3 int32
	for {
		err = binary.Read(inFile1, tt.byteOrder, &data1)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "text041", "binary.Read(inFile1)", err)
			return false
		}
		err = binary.Read(inFile2, tt.byteOrder, &data2)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "text041", "binary.Read(inFile2)", err)
			return false
		}
		err = binary.Read(inFile3, tt.byteOrder, &data3)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "text041", "binary.Read(inFile3)", err)
			return false
		}
		_, err = outFile.WriteString(fmt.Sprintf("|%-20d%-20d%-20d|\n", data1, data2, data3))
		if err != nil {
			mylog.Log("task", "text041", "outFile.WriteString", err)
			return false
		}
	}
	return tt.checker.CompareTextFile(outFileName)
}

func (tt *textTask) text42() bool {
	var a = tt.scanner.NextFloat64()
	var b = tt.scanner.NextFloat64()
	var n = tt.scanner.NextInt()
	tt.scanner.Skip()
	var fileName = tt.scanner.NextLine()
	var file, err = os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text042", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var h = (b - a) / float64(n)
	var x = a
	for indeX := 0; indeX <= n; indeX++ {
		_, err = file.WriteString(fmt.Sprintf("% 10.4f% 15.8f\n", x, math.Sqrt(x)))
		if err != nil {
			mylog.Log("task", "text042", "file.WriteString", err)
			return false
		}
		x += h
	}
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text43() bool {
	var a = tt.scanner.NextFloat64()
	var b = tt.scanner.NextFloat64()
	var n = tt.scanner.NextInt()
	tt.scanner.Skip()
	var fileName = tt.scanner.NextLine()
	var file, err = os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text043", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var h = (b - a) / float64(n)
	var x = a
	for indeX := 0; indeX <= n; indeX++ {
		_, err = file.WriteString(fmt.Sprintf("% 8.4f% 12.8f% 12.8f\n", x, math.Sin(x), math.Cos(x)))
		if err != nil {
			mylog.Log("task", "text043", "file.WriteString", err)
			return false
		}
		x += h
	}
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text44() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text044", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var count, sum = 0, 0
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var number int
	for scanner.Scan() {
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			mylog.Log("task", "text044", "strconv.Atoi", err)
			return false
		}
		count++
		sum += number
	}
	return tt.checker.CompareInt(count, sum)
}

func (tt *textTask) text45() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text045", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var count = 0
	var sum float64
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var number float64
	for scanner.Scan() {
		number, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			mylog.Log("task", "text045", "strconv.ParseFloat", err)
			return false
		}
		if number-float64(int(number)) != 0 {
			count++
			sum += number
		}
	}
	return tt.checker.CompareInt(count) && tt.checker.CompareFloat64(2, sum)
}

func (tt *textTask) text46() bool {
	var textFileName = tt.scanner.NextTextFileName()
	var textFile, err = os.OpenFile(textFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text046", "os.OpenFile(textFileName)", err)
		return false
	}
	defer textFile.Close()
	var binaryFileName = tt.scanner.NextLine()
	binaryFile, err := os.OpenFile(binaryFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text046", "os.OpenFile(binaryFileName)", err)
		return false
	}
	defer binaryFile.Close()
	var scanner = bufio.NewScanner(textFile)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		var word = scanner.Text()
		number, err := strconv.ParseFloat(word, 64)
		if err != nil {
			continue
		}
		if number-float64(int64(number)) == 0 {
			continue
		}
		err = binary.Write(binaryFile, tt.byteOrder, number)
		if err != nil {
			mylog.Log("task", "text046", "binary.Write", err)
			return false
		}
	}
	return tt.checker.CompareBinaryFile(tt.byteOrder, binaryFileName)
}

func (tt *textTask) text47() bool {
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text047", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var count, sum = 0, 0
	var number int
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		count++
		sum += number
	}
	return tt.checker.CompareInt(count, sum)
}

func (tt *textTask) text48() bool {
	var inFileName = tt.scanner.NextTextFileName()
	var outFileName = tt.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text048", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text048", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var scanner = bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanWords)
	var number int
	for scanner.Scan() {
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		err = binary.Write(outFile, tt.byteOrder, int32(number))
		if err != nil {
			mylog.Log("task", "text048", "binary.Write", err)
			return false
		}
	}
	return tt.checker.CompareBinaryFile(tt.byteOrder, outFileName)
}

func (tt *textTask) text49() bool {
	var fileName1 = tt.scanner.NextTextFileName()
	var fileName2 = tt.scanner.NextBinaryFileName(tt.byteOrder)
	var tempFileName = "$tempFileName_text49.tmp$"
	var file1, file2, tempFile *os.File
	var err error
	file1, err = os.OpenFile(fileName1, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text049", "os.OpenFile(fileName1)", err)
		return false
	}
	defer file1.Close()
	file2, err = os.OpenFile(fileName2, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text049", "os.OpenFile(fileName2)", err)
		return false
	}
	defer file2.Close()
	tempFile, err = os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text049", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file1)
	scanner.Split(bufio.ScanLines)
	var number int32
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if err == nil {
			err = binary.Read(file2, tt.byteOrder, &number)
			if err == nil {
				line += strconv.Itoa(int(number))
			}
		}
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("task", "text049", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName1)
	os.Rename(tempFileName, fileName1)
	return tt.checker.CompareTextFile(fileName1)
}

func (tt *textTask) text50() bool {
	var textFileName = tt.scanner.NextTextFileName()
	var stringFileName = tt.scanner.NextLine()
	var floatFileName = tt.scanner.NextLine()
	var textFile, stringFile, floatFile *os.File
	var err error
	textFile, err = os.OpenFile(textFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text050", "os.OpenFile(textFileName)", err)
		return false
	}
	defer textFile.Close()
	stringFile, err = os.OpenFile(stringFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text050", "os.OpenFile(stringFileName)", err)
		return false
	}
	defer stringFile.Close()
	floatFile, err = os.OpenFile(floatFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text050", "os.OpenFile(floatFileName)", err)
		return false
	}
	defer floatFile.Close()
	var scanner = bufio.NewScanner(textFile)
	scanner.Split(bufio.ScanLines)
	var line string
	var number float64
	for scanner.Scan() {
		line = scanner.Text()
		var data = []byte(line[:30])
		data = append(data, make([]byte, 50)...)
		err = binary.Write(stringFile, tt.byteOrder, data)
		if err != nil {
			mylog.Log("task", "text050", "binary.Write(stringFile)", err)
			return false
		}
		number, err = strconv.ParseFloat(line[30:], 64)
		if err != nil {
			continue
		}
		err = binary.Write(floatFile, tt.byteOrder, number)
		if err != nil {
			mylog.Log("task", "text050", "binary.Write(floatFile)", err)
			return false
		}
	}
	return tt.checker.CompareBinaryFile(tt.byteOrder, stringFileName, floatFileName)
}

func (tt *textTask) text51() bool {
	var inFileName = tt.scanner.NextTextFileName()
	var inFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text051", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()

	var outFileNames [3]string
	var outFiles [3]*os.File
	for index := 0; index < 3; index++ {
		outFileNames[index] = tt.scanner.NextLine()
		outFiles[index], err = os.OpenFile(outFileNames[index], os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
		if err != nil {
			mylog.Log("task", "text051", fmt.Sprintf("os.OpenFile(outFileNames[%d])", index), err)
			return false
		}
		defer outFiles[index].Close()
	}
	var scanner = bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var line string
	var words []string
	var number float64
	for scanner.Scan() {
		line = scanner.Text()
		words = strings.Fields(line)
		for index, word := range words {
			number, err = strconv.ParseFloat(word, 64)
			if err != nil {
				continue
			}
			err = binary.Write(outFiles[index], tt.byteOrder, number)
			if err != nil {
				mylog.Log("task", "text051", fmt.Sprintf("binary.Write(outFiles[%d])", index), err)
				return false
			}
		}
	}
	return tt.checker.CompareBinaryFile(tt.byteOrder, outFileNames[0], outFileNames[1], outFileNames[2])
}

func (tt *textTask) text52() bool {
	var inFileName = tt.scanner.NextTextFileName()
	var outFileName = tt.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text052", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text052", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var scanner = bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		separator := rune(line[0])
		words := strings.FieldsFunc(line, func(r rune) bool { return r == separator })
		var sum int32
		for _, word := range words {
			word = strings.TrimSpace(word)
			number, err := strconv.Atoi(word)
			if err != nil {
				continue
			}
			sum += int32(number)
		}
		err = binary.Write(outFile, tt.byteOrder, sum)
		if err != nil {
			mylog.Log("task", "text052", "binary.Write", err)
			return false
		}
	}
	return tt.checker.CompareBinaryFile(tt.byteOrder, outFileName)
}

func (tt *textTask) text53() bool {
	var inFileName = tt.scanner.NextTextFileName()
	var outFileName = tt.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text053", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text053", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var scanner = bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var character rune
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		for _, character = range line {
			if unicode.IsPunct(character) && character != '\'' {
				err = binary.Write(outFile, tt.byteOrder, byte(character))
				if err != nil {
					mylog.Log("task", "text053", "binary.Write", err)
					return false
				}
			}
		}
	}
	return tt.checker.CompareBinaryFile(tt.byteOrder, outFileName)
}

func (tt *textTask) text54() bool {
	var inFileName = tt.scanner.NextTextFileName()
	var outFileName = tt.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text054", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "text054", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var scanner = bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var byteArray = make([]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		for _, character := range line {
			exists := false
			for _, bayt := range byteArray {
				if bayt == byte(character) {
					exists = true
					break
				}
			}
			if !exists {
				byteArray = append(byteArray, byte(character))
			}
		}
	}
	for _, bayt := range byteArray {
		err = binary.Write(outFile, tt.byteOrder, bayt)
		if err != nil {
			mylog.Log("task", "text054", "binary.Write", err)
			return false
		}
	}
	return tt.checker.CompareBinaryFile(tt.byteOrder, outFileName)
}

func (tt *textTask) text55() bool {
	var inFileName = tt.scanner.NextTextFileName()
	var outFileName = tt.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text055", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text055", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var array = make([]byte, 0)
	var data = make([]byte, 1)
	var character byte
	var exists bool
	for {
		_, err = inFile.Read(data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "text055", "inFile.Read", err)
			return false
		}
		character = data[0]
		if character == '\n' || character == '\r' {
			continue
		}
		exists = false
		for index, fileChar := range array {
			if fileChar == character {
				exists = true
				break
			} else if fileChar > character {
				array[index], character = character, fileChar
			}
		}
		if !exists {
			array = append(array, character)
		}
	}
	for _, fileChar := range array {
		err = binary.Write(outFile, tt.byteOrder, fileChar)
		if err != nil {
			mylog.Log("task", "text055", "binary.Write", err)
			return false
		}
	}
	return tt.checker.CompareBinaryFile(tt.byteOrder, outFileName)
}

func (tt *textTask) text56() bool {
	var inFileName = tt.scanner.NextTextFileName()
	var outFileName = tt.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text056", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text056", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var array = make([]byte, 0)
	var data = make([]byte, 1)
	var character byte
	var exists bool
	for {
		_, err = inFile.Read(data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "text056", "inFile.Read", err)
			return false
		}
		character = data[0]
		if character == '\n' || character == '\r' {
			continue
		}
		exists = false
		for index, fileChar := range array {
			if fileChar == character {
				exists = true
				break
			} else if fileChar < character {
				array[index], character = character, fileChar
			}
		}
		if !exists {
			array = append(array, character)
		}
	}
	for _, fileChar := range array {
		err = binary.Write(outFile, tt.byteOrder, fileChar)
		if err != nil {
			mylog.Log("task", "text056", "binary.Write", err)
			return false
		}
	}
	return tt.checker.CompareBinaryFile(tt.byteOrder, outFileName)
}

func (tt *textTask) text57() bool {
	var inFileName = tt.scanner.NextTextFileName()
	var outFileName = tt.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text057", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text057", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var from, to byte = 'a', 'z'
	var lowerLatins = make([]byte, 0)
	var lowerLatinCounts = make([]int, 0)
	for letter := from; letter <= to; letter++ {
		lowerLatins = append(lowerLatins, letter)
		lowerLatinCounts = append(lowerLatinCounts, 0)
	}
	var data = make([]byte, 1)
	var character byte
	for {
		_, err = inFile.Read(data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "text057", "inFile.Read", err)
			return false
		}
		character = data[0]
		for index, letter := range lowerLatins {
			if character == letter {
				lowerLatinCounts[index]++
			}
		}
	}
	var countStr string
	var pos int
	for index, letter := range lowerLatins {
		if lowerLatinCounts[index] > 0 {
			data = make([]byte, 80)
			data[0] = letter
			data[1] = '-'
			countStr = strconv.Itoa(lowerLatinCounts[index])
			pos = 2
			for _, digit := range countStr {
				data[pos] = byte(digit)
				pos++
			}
			err = binary.Write(outFile, tt.byteOrder, data)
			if err != nil {
				mylog.Log("task", "text057", "binary.Write", err)
				return false
			}
		}
	}
	return tt.checker.CompareBinaryFile(tt.byteOrder, outFileName)
}

func (tt *textTask) text58() bool {
	var inFileName = tt.scanner.NextTextFileName()
	var outFileName = tt.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text058", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text058", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var from, to byte = 'a', 'z'
	var lowerLatins = make([]byte, 0)
	var lowerLatinCounts = make([]int, 0)
	for letter := from; letter <= to; letter++ {
		lowerLatins = append(lowerLatins, letter)
		lowerLatinCounts = append(lowerLatinCounts, 0)
	}
	var data = make([]byte, 1)
	var character byte
	for {
		_, err = inFile.Read(data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "text058", "inFile.Read", err)
			return false
		}
		character = data[0]
		for index, letter := range lowerLatins {
			if character == letter {
				lowerLatinCounts[index]++
			}
		}
	}
	var arrayLength = len(lowerLatins)
	var indexes = make([]int, arrayLength)
	for index := 0; index < arrayLength; index++ {
		indexes[index] = index
	}
	for i := 0; i < arrayLength-1; i++ {
		for index := 1; index < arrayLength-i; index++ {
			if lowerLatinCounts[indexes[index-1]] < lowerLatinCounts[indexes[index]] {
				indexes[index-1], indexes[index] = indexes[index], indexes[index-1]
			}
		}
	}
	var countStr string
	var pos int
	for index := range lowerLatins {
		if lowerLatinCounts[indexes[index]] > 0 {
			data = make([]byte, 80)
			data[0] = lowerLatins[indexes[index]]
			data[1] = '-'
			countStr = strconv.Itoa(lowerLatinCounts[indexes[index]])
			pos = 2
			for _, digit := range countStr {
				data[pos] = byte(digit)
				pos++
			}
			err = binary.Write(outFile, tt.byteOrder, data)
			if err != nil {
				mylog.Log("task", "text058", "binary.Write", err)
				return false
			}
		}
	}
	if !tt.checker.CompareBinaryFile(tt.byteOrder, outFileName) {
		time.Sleep(time.Hour)
		return false
	}
	return true
}

func (tt *textTask) text59() bool {
	var s = tt.scanner.NextLine()
	var array [10]int
	var index = 0
	for _, character := range s {
		if unicode.IsDigit(character) {
			array[index] = int(character) - int('0')
		}
		index++
	}
	var fileName = tt.scanner.NextTextFileName()
	var tempFileName = "$tempFileName_text59.tmp$"
	var file, tempFile *os.File
	var err error
	file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text059", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	tempFile, err = os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "text059", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var upperLatins, lowerLatins, seekLatins []byte
	lowerLatins = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	lowerLatins = append(lowerLatins, lowerLatins...)
	upperLatins = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	upperLatins = append(upperLatins, upperLatins...)
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line []byte
	for scanner.Scan() {
		index = 0
		line = scanner.Bytes()
		for i, character := range line {
			if unicode.Is(unicode.Latin, rune(character)) {
				switch {
				case unicode.IsLower(rune(character)):
					seekLatins = lowerLatins
				case unicode.IsUpper(rune(character)):
					seekLatins = upperLatins
				default:
					seekLatins = nil
				}
				for indexOf, bayt := range seekLatins {
					if bayt == character {
						line[i] = seekLatins[indexOf+array[index]]
						break
					}
				}
			}
			index = (index + 1) % 10
		}
		_, err = tempFile.Write(append(line, '\n'))
		if err != nil {
			mylog.Log("task", "text059", "tempFile.WriteString", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return tt.checker.CompareTextFile(fileName)
}

func (tt *textTask) text60() bool {
	var s = tt.scanner.NextLine()
	var fileName = tt.scanner.NextTextFileName()
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "text060", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line []byte
	var array [10]int
	var isValidDecoder bool
	if scanner.Scan() {
		line = scanner.Bytes()
		isValidDecoder = true
		if len(line) == len(s) {
			var farq, index int
			for i, b := range line {
				farq = int(b) - int(s[i])
				if farq != 0 {
					if farq < 0 {
						farq += 26
					}
					if array[index] != 0 && array[index] != farq {
						isValidDecoder = false
						break
					}
					array[index] = farq
				}
				index = (index + 1) % 10
			}
		}
	}
	for index := 0; index < 10; index++ {
		if array[index] == 0 {
			isValidDecoder = false
			break
		}
	}
	if isValidDecoder {
		var upperLatins, lowerLatins, seekLatins []byte
		lowerLatins = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
		lowerLatins = append(lowerLatins, lowerLatins...)
		upperLatins = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
		upperLatins = append(upperLatins, upperLatins...)
		var tempFileName = "$tempFileName_text60.tmp$"
		tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
		if err != nil {
			mylog.Log("task", "text060", "os.OpenFile(tempFileName)", err)
			return false
		}
		defer tempFile.Close()
		file.Seek(0, 0)
		scanner = bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var index int
		for scanner.Scan() {
			index = 0
			line = scanner.Bytes()
			for i, character := range line {
				if unicode.Is(unicode.Latin, rune(character)) {
					switch {
					case unicode.IsLower(rune(character)):
						seekLatins = lowerLatins
					case unicode.IsUpper(rune(character)):
						seekLatins = upperLatins
					default:
						seekLatins = nil
					}
					for indexOf := len(seekLatins) - 1; indexOf >= 0; indexOf-- {
						if seekLatins[indexOf] == character {
							line[i] = seekLatins[indexOf-array[index]]
							break
						}
					}
				}
				index = (index + 1) % 10
			}
			_, err = tempFile.Write(append(line, '\n'))
			if err != nil {
				mylog.Log("task", "text060", "tempFile.WriteString", err)
				return false
			}
		}
		os.Remove(fileName)
		os.Rename(tempFileName, fileName)
	}
	return tt.checker.CompareTextFile(fileName)
}
