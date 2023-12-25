package task

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/mylog"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type fileTask struct {
	dataPath  string
	name      string
	count     int
	byteOrder binary.ByteOrder
	scanner   scanner.Scanner
	checker   checker.Checker
}

func NewFileTask(pathPrefix string, byteOrder binary.ByteOrder) Task {
	return &fileTask{
		dataPath:  pathPrefix + "/14file/File%03d/%03d/%03d",
		name:      "File",
		count:     90,
		byteOrder: byteOrder,
	}
}

func (ft *fileTask) GetCount() int { return ft.count }

func (ft *fileTask) GetName() string { return ft.name }

func (ft *fileTask) SetData(taskNo int, testNo int) error {
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

func (ft *fileTask) ScannerEOF() bool { return ft.scanner.EOF() }

func (ft *fileTask) CheckerEOF() bool { return ft.checker.EOF() }

func (ft *fileTask) CleanData() {
	ft.scanner.RemoveFiles()
	ft.checker.RemoveFiles()
}

func (ft *fileTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return ft.file1
	case 2:
		return ft.file2
	case 3:
		return ft.file3
	case 4:
		return ft.file4
	case 5:
		return ft.file5
	case 6:
		return ft.file6
	case 7:
		return ft.file7
	case 8:
		return ft.file8
	case 9:
		return ft.file9
	case 10:
		return ft.file10
	case 11:
		return ft.file11
	case 12:
		return ft.file12
	case 13:
		return ft.file13
	case 14:
		return ft.file14
	case 15:
		return ft.file15
	case 16:
		return ft.file16
	case 17:
		return ft.file17
	case 18:
		return ft.file18
	case 19:
		return ft.file19
	case 20:
		return ft.file20
	case 21:
		return ft.file21
	case 22:
		return ft.file22
	case 23:
		return ft.file23
	case 24:
		return ft.file24
	case 25:
		return ft.file25
	case 26:
		return ft.file26
	case 27:
		return ft.file27
	case 28:
		return ft.file28
	case 29:
		return ft.file29
	case 30:
		return ft.file30
	case 31:
		return ft.file31
	case 32:
		return ft.file32
	case 33:
		return ft.file33
	case 34:
		return ft.file34
	case 35:
		return ft.file35
	case 36:
		return ft.file36
	case 37:
		return ft.file37
	case 38:
		return ft.file38
	case 39:
		return ft.file39
	case 40:
		return ft.file40
	case 41:
		return ft.file41
	case 42:
		return ft.file42
	case 43:
		return ft.file43
	case 44:
		return ft.file44
	case 45:
		return ft.file45
	case 46:
		return ft.file46
	case 47:
		return ft.file47
	case 48:
		return ft.file48
	case 49:
		return ft.file49
	case 50:
		return ft.file50
	case 51:
		return ft.file51
	case 52:
		return ft.file52
	case 53:
		return ft.file53
	case 54:
		return ft.file54
	case 55:
		return ft.file55
	case 56:
		return ft.file56
	case 57:
		return ft.file57
	case 58:
		return ft.file58
	case 59:
		return ft.file59
	case 60:
		return ft.file60
	case 61:
		return ft.file61
	case 62:
		return ft.file62
	case 63:
		return ft.file63
	case 64:
		return ft.file64
	case 65:
		return ft.file65
	case 66:
		return ft.file66
	case 67:
		return ft.file67
	case 68:
		return ft.file68
	case 69:
		return ft.file69
	case 70:
		return ft.file70
	case 71:
		return ft.file71
	case 72:
		return ft.file72
	case 73:
		return ft.file73
	case 74:
		return ft.file74
	case 75:
		return ft.file75
	case 76:
		return ft.file76
	case 77:
		return ft.file77
	case 78:
		return ft.file78
	case 79:
		return ft.file79
	case 80:
		return ft.file80
	case 81:
		return ft.file81
	case 82:
		return ft.file82
	case 83:
		return ft.file83
	case 84:
		return ft.file84
	case 85:
		return ft.file85
	case 86:
		return ft.file86
	case 87:
		return ft.file87
	case 88:
		return ft.file88
	case 89:
		return ft.file89
	case 90:
		return ft.file90
	default:
		return nil
	}
}

// { "no": 1, "dat": "", "ans": "" }
func (ft *fileTask) file1() bool {
	var s = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var _, err = os.Create(s)
	var created = err == nil
	if created {
		os.Remove(s)
	}
	return ft.checker.CompareBool(created)
}

// { "no": 2, "dat": "", "ans": "" }
func (ft *fileTask) file2() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var n = ft.scanner.NextInt()
	var file, err = os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file002", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var number int32 = 2
	for index := 0; index < n; index++ {
		err = binary.Write(file, ft.byteOrder, number)
		if err != nil {
			mylog.Log("task", "file002", "binary.Write", err)
			return false
		}
		number += 2
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 3, "dat": "2", "ans": "2" }
func (ft *fileTask) file3() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var a = ft.scanner.NextFloat64()
	var d = ft.scanner.NextFloat64()
	var file, err = os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file003", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var number = a
	for index := 0; index < 10; index++ {
		err = binary.Write(file, ft.byteOrder, number)
		if err != nil {
			mylog.Log("task", "file003", "binary.Write", err)
			return false
		}
		number += d
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 4, "dat": "", "ans": "" }
func (ft *fileTask) file4() bool {
	var (
		fileName string
		err      error
		count    = 0
	)
	for index := 0; index < 4; index++ {
		fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
		_, err = os.Stat(fileName)
		if err == nil {
			count++
		}
	}
	return ft.checker.CompareInt(count)
}

// { "no": 5, "dat": "", "ans": "" }
func (ft *fileTask) file5() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var count = -1
	var file, err = os.Stat(fileName)
	if err == nil {
		count = int(file.Size()) / 4
	}
	return ft.checker.CompareInt(count)
}

// { "no": 6, "dat": "", "ans": "" }
func (ft *fileTask) file6() bool {
	var k = ft.scanner.NextInt64()
	ft.scanner.Skip()
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file006", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var data int32 = -1
	var info, _ = file.Stat()
	var fileElements = info.Size() / 4
	if fileElements >= k {
		file.Seek((k-1)*int64(4), 0)
		err = binary.Read(file, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file006", "binary.Read", err)
			return false
		}
	}
	return ft.checker.CompareInt32(data)
}

// { "no": 7, "dat": "", "ans": "" }
func (ft *fileTask) file7() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file007", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var data int32
	for index := 1; index <= 4; index++ {
		err = binary.Read(file, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file007", "binary.Read", err)
			return false
		}
		if !ft.checker.CompareInt32(data) {
			return false
		}
		if index == 2 {
			file.Seek(-8, 2)
		}
	}
	return true
}

// { "no": 8, "dat": "2", "ans": "2" }
func (ft *fileTask) file8() bool {
	var fileName1 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var fileName2 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file1, err = os.OpenFile(fileName1, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file008", "os.OpenFile 1", err)
		return false
	}
	defer file1.Close()
	file2, err := os.OpenFile(fileName2, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file008", "os.OpenFile 2", err)
		return false
	}
	defer file2.Close()
	var data float64
	for index := 1; index <= 2; index++ {
		err = binary.Read(file1, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file008", "binary.Read", err)
			return false
		}
		err = binary.Write(file2, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file008", "binary.Write", err)
			return false
		}
		if index == 1 {
			info, err := file1.Stat()
			if err != nil {
				mylog.Log("task", "file008", "file1.Stat", err)
				return false
			}
			if info.Size() > 8 {
				file1.Seek(-8, 2)
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName2)
}

// { "no": 9, "dat": "2", "ans": "2" }
func (ft *fileTask) file9() bool {
	var fileName1 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var fileName2 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file1, err = os.OpenFile(fileName1, os.O_RDONLY, 0777)
	if err != nil {
		fileName1, fileName2 = fileName2, fileName1
		file1, err = os.OpenFile(fileName1, os.O_RDONLY, 0777)
		if err != nil {
			mylog.Log("task", "file009", "os.OpenFile(fileName1)", err)
			return false
		}
	}
	defer file1.Close()
	file2, err := os.OpenFile(fileName2, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file009", "os.OpenFile(fileName2)", err)
		return false
	}
	defer file2.Close()
	var data float64
	for index := 1; index <= 2; index++ {
		if index == 1 {
			info, err := file1.Stat()
			if err != nil {
				mylog.Log("task", "file009", "file1.Stat", err)
				return false
			}
			if info.Size() > 8 {
				file1.Seek(-8, 2)
			}
		} else {
			file1.Seek(0, 0)
		}
		err = binary.Read(file1, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file009", "binary.Read", err)
			return false
		}
		err = binary.Write(file2, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file009", "binary.Write", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName2)
}

// { "no": 10, "dat": "", "ans": "" }
func (ft *fileTask) file10() bool {
	var inputFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outputFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file1, err = os.OpenFile(inputFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file010", "os.OpenFile(inputFileName)", err)
		return false
	}
	defer file1.Close()
	file2, err := os.OpenFile(outputFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file010", "os.OpenFile(outputFileName)", err)
		return false
	}
	defer file2.Close()
	info, err := file1.Stat()
	if err != nil {
		mylog.Log("task", "file010", "file1.Stat", err)
		return false
	}
	var data int32
	for position := info.Size() - 4; position >= 0; position -= 4 {
		file1.Seek(position, 0)
		err = binary.Read(file1, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file010", "binary.Read", err)
			return false
		}
		err = binary.Write(file2, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file010", "binary.Write", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outputFileName)
}

// { "no": 11, "dat": "2", "ans": "2" }
func (ft *fileTask) file11() bool {
	var inputFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outputFileName1 = ft.scanner.NextLine()
	var outputFileName2 = ft.scanner.NextLine()
	var inFile, err = os.OpenFile(inputFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file011", "os.OpenFile", err)
		return false
	}
	defer inFile.Close()
	var outputFile1, outputFile2 *os.File
	outputFile1, err = os.OpenFile(outputFileName1, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file011", "os.OpenFile(outputFileName1)", err)
		return false
	}
	defer outputFile1.Close()
	outputFile2, err = os.OpenFile(outputFileName2, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file011", "os.OpenFile(outputFileName2)", err)
		return false
	}
	defer outputFile2.Close()
	var data float64
	var index = 0
	for {
		err = binary.Read(inFile, ft.byteOrder, &data)
		if err != nil {
			break
		}
		if index%2 == 0 {
			err = binary.Write(outputFile1, ft.byteOrder, data)
		} else {
			err = binary.Write(outputFile2, ft.byteOrder, data)
		}
		if err != nil {
			mylog.Log("task", "file011", "binary.Write", err)
			return false
		}
		index++
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outputFileName1, outputFileName2)
}

// { "no": 12, "dat": "", "ans": "" }
func (ft *fileTask) file12() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName1 = ft.scanner.NextLine()
	var outFileName2 = ft.scanner.NextLine()
	var inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file012", "os.OpenFile", err)
		return false
	}
	defer inFile.Close()
	var outFile1, outFile2 *os.File
	outFile1, err = os.OpenFile(outFileName1, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file012", "os.OpenFile(outFileName1)", err)
		return false
	}
	defer outFile1.Close()
	outFile2, err = os.OpenFile(outFileName2, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file012", "os.OpenFile(outFileName2)", err)
		return false
	}
	defer outFile2.Close()
	var data int32
	for {
		err = binary.Read(inFile, ft.byteOrder, &data)
		if err != nil {
			break
		}
		if data%2 == 0 {
			err = binary.Write(outFile1, ft.byteOrder, data)
		} else {
			err = binary.Write(outFile2, ft.byteOrder, data)
		}
		if err != nil {
			mylog.Log("task", "file012", "binary.Write", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName1, outFileName2)
}

// { "no": 13, "dat": "", "ans": "" }
func (ft *fileTask) file13() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName1 = ft.scanner.NextLine()
	var outFileName2 = ft.scanner.NextLine()
	var inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file013", "os.OpenFile", err)
		return false
	}
	defer inFile.Close()
	var outFile1, outFile2 *os.File
	outFile1, err = os.OpenFile(outFileName1, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file013", "os.OpenFile(outFileName1)", err)
		return false
	}
	defer outFile1.Close()
	outFile2, err = os.OpenFile(outFileName2, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file013", "os.OpenFile(outFileName2)", err)
		return false
	}
	defer outFile2.Close()
	var info, _ = inFile.Stat()
	var position = info.Size()
	var data int32
	for position > 3 {
		position -= 4
		inFile.Seek(position, 0)
		err = binary.Read(inFile, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file013", "binary.Read", err)
			return false
		}
		if data >= 0 {
			err = binary.Write(outFile1, ft.byteOrder, data)
		} else {
			err = binary.Write(outFile2, ft.byteOrder, data)
		}
		if err != nil {
			mylog.Log("task", "file013", "binary.Write", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName1, outFileName2)
}

// { "no": 14, "dat": "2", "ans": "2" }
func (ft *fileTask) file14() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file014", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var elementsCount = int(info.Size()) / 8
	var sum = 0.0
	var data float64
	for index := 0; index < elementsCount; index++ {
		err = binary.Read(file, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file014", "binary.Read", err)
			return false
		}
		sum += data
	}
	var average = sum / float64(elementsCount)
	return ft.checker.CompareFloat64(2, average)
}

// { "no": 15, "dat": "2", "ans": "2" }
func (ft *fileTask) file15() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file015", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var elementsCount = info.Size()
	var data float64
	var sum = 0.0
	var position int64
	for position = 8; position < elementsCount; position += 16 {
		file.Seek(position, 0)
		err = binary.Read(file, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file015", "binary.Read", err)
			return false
		}
		sum += data
	}
	return ft.checker.CompareFloat64(2, sum)
}

// { "no": 16, "dat": "", "ans": "" }
func (ft *fileTask) file16() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file016", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var prev, curr int32
	var elementsCount = int(info.Size()) / 4
	var seriesCount = 0
	for index := 0; index < elementsCount; index++ {
		err = binary.Read(file, ft.byteOrder, &curr)
		if err != nil {
			mylog.Log("task", "file016", "binary.Read", err)
			return false
		}
		if index == 0 || prev != curr {
			seriesCount++
		}
		prev = curr
	}
	return ft.checker.CompareInt(seriesCount)
}

// { "no": 17, "dat": "", "ans": "" }
func (ft *fileTask) file17() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file017", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file017", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var info, _ = inFile.Stat()
	var elementsCount = int(info.Size()) / 4
	var count int32
	var prev, curr int32
	for index := 0; index < elementsCount; index++ {
		err = binary.Read(inFile, ft.byteOrder, &curr)
		if err != nil {
			mylog.Log("task", "file017", "binary.Read", err)
			return false
		}
		if index > 0 && prev != curr {
			err = binary.Write(outFile, ft.byteOrder, count)
			if err != nil {
				mylog.Log("task", "file017", "binary.Write", err)
				return false
			}
			count = 0
		}
		count++
		prev = curr
	}
	err = binary.Write(outFile, ft.byteOrder, count)
	if err != nil {
		mylog.Log("task", "file017", "binary.Write", err)
		return false
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 18, "dat": "2", "ans": "2" }
func (ft *fileTask) file18() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file018", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var elementsCount = int(info.Size()) / 8
	var prev, curr, next, localMininum float64
	var found = false
	index := 0
	for ; index < elementsCount; index++ {
		err = binary.Read(file, ft.byteOrder, &next)
		if err != nil {
			mylog.Log("task", "file018", "binary.Read", err)
			return false
		}
		if index == 1 && curr < next {
			localMininum = curr
			found = true
			break
		} else if index > 1 && curr < prev && curr < next {
			localMininum = curr
			found = true
			break
		}
		prev = curr
		curr = next
	}
	if !found && index > 1 && curr < prev {
		localMininum = curr
	}
	return ft.checker.CompareFloat64(2, localMininum)
}

// { "no": 19, "dat": "2", "ans": "2" }
func (ft *fileTask) file19() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file019", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var elementsCount = int(info.Size()) / 8
	var prev, curr, next, localMaximum float64
	for index := 0; index < elementsCount; index++ {
		err = binary.Read(file, ft.byteOrder, &next)
		if err != nil {
			mylog.Log("task", "file019", "binary.Read", err)
			return false
		}
		if index == 1 && curr > next {
			localMaximum = curr
		} else if index > 1 && curr > prev && curr > next {
			localMaximum = curr
		}
		prev = curr
		curr = next
	}
	if curr > prev {
		localMaximum = curr
	}
	return ft.checker.CompareFloat64(2, localMaximum)
}

// { "no": 20, "dat": "2", "ans": "" }
func (ft *fileTask) file20() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file020", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var elementsCount = int(info.Size()) / 8
	var prev, curr, next float64
	var count = 0
	for index := 0; index < elementsCount; index++ {
		err = binary.Read(file, ft.byteOrder, &next)
		if err != nil {
			mylog.Log("task", "file020", "binary.Read", err)
			return false
		}
		if index == 1 && curr != next {
			count++
		} else if index > 1 {
			if curr > prev && curr > next || curr < prev && curr < next {
				count++
			}
		}
		prev = curr
		curr = next
	}
	if curr != prev {
		count++
	}
	return ft.checker.CompareInt(count)
}

// { "no": 21, "dat": "2", "ans": "" }
func (ft *fileTask) file21() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file021", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file021", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var info, _ = inFile.Stat()
	var elementsCount = int32(info.Size() / 8)
	var prev, curr, next float64
	var indeX int32
	for indeX = 1; indeX <= elementsCount; indeX++ {
		err = binary.Read(inFile, ft.byteOrder, &next)
		if err != nil {
			mylog.Log("task", "file021", "binary.Read", err)
			return false
		}
		if indeX == 2 && curr > next {
			err = binary.Write(outFile, ft.byteOrder, indeX-1)
			if err != nil {
				mylog.Log("task", "file021", "binary.Write", err)
				return false
			}
		} else if indeX > 2 && curr > prev && curr > next {
			err = binary.Write(outFile, ft.byteOrder, indeX-1)
			if err != nil {
				mylog.Log("task", "file021", "binary.Write", err)
				return false
			}
		}
		prev = curr
		curr = next
	}
	if curr > prev {
		err = binary.Write(outFile, ft.byteOrder, elementsCount)
		if err != nil {
			mylog.Log("task", "file021", "binary.Write", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 22, "dat": "2", "ans": "" }
func (ft *fileTask) file22() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file022", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file022", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var prev, curr, next float64
	var info, _ = inFile.Stat()
	var elementsCount = int32(info.Size()) / 8
	var index, localIndex int32
	for index = elementsCount; index > 0; index-- {
		inFile.Seek(int64(index)*8-8, 0)
		err = binary.Read(inFile, ft.byteOrder, &prev)
		if err != nil {
			mylog.Log("task", "file022", "binary.Read", err)
			return false
		}
		localIndex = -1
		if index == elementsCount-1 && curr != prev {
			localIndex = index + 1
		} else if curr > prev && curr > next || curr < prev && curr < next {
			localIndex = index + 1
		}
		if localIndex > 0 {
			err = binary.Write(outFile, ft.byteOrder, localIndex)
			if err != nil {
				mylog.Log("task", "file022", "binary.Write", err)
				return false
			}
		}
		next = curr
		curr = prev
	}
	if elementsCount > 1 && curr != next {
		localIndex = 1
		err = binary.Write(outFile, ft.byteOrder, localIndex)
		if err != nil {
			mylog.Log("task", "file022", "binary.Write", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 23, "dat": "2", "ans": "" }
func (ft *fileTask) file23() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file023", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file023", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var prev, curr float64
	var info, _ = inFile.Stat()
	var elementsCount = int(info.Size()) / 8
	var count int32
	for index := 1; index <= elementsCount; index++ {
		err = binary.Read(inFile, ft.byteOrder, &curr)
		if err != nil {
			mylog.Log("task", "file023", "binary.Read", err)
			return false
		}
		if index > 1 && prev <= curr {
			if count > 1 {
				err = binary.Write(outFile, ft.byteOrder, count)
				if err != nil {
					mylog.Log("task", "file023", "binary.Write", err)
					return false
				}
			}
			count = 0
		}
		count++
		prev = curr
	}
	if count > 1 {
		err = binary.Write(outFile, ft.byteOrder, count)
		if err != nil {
			mylog.Log("task", "file023", "binary.Write", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 24, "dat": "2", "ans": "" }
func (ft *fileTask) file24() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file024", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file024", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var info, _ = inFile.Stat()
	var elementsCount = int(info.Size()) / 8
	var prev, curr float64
	var ascCount, descCount int32
	for index := 1; index <= elementsCount; index++ {
		err = binary.Read(inFile, ft.byteOrder, &curr)
		if err != nil {
			mylog.Log("task", "file024", "binary.Read", err)
			return false
		}
		if index > 1 {
			if prev < curr {
				if descCount > 1 {
					err = binary.Write(outFile, ft.byteOrder, descCount)
					if err != nil {
						mylog.Log("task", "file024", "binary.Write", err)
						return false
					}
				}
				descCount = 0
			} else if prev > curr {
				if ascCount > 1 {
					err = binary.Write(outFile, ft.byteOrder, ascCount)
					if err != nil {
						mylog.Log("task", "file024", "binary.Write", err)
						return false
					}
				}
				ascCount = 0
			}
		}
		ascCount++
		descCount++
		prev = curr
	}
	if descCount > 1 {
		err = binary.Write(outFile, ft.byteOrder, descCount)
		if err != nil {
			mylog.Log("task", "file024", "binary.Write", err)
			return false
		}
	} else if ascCount > 1 {
		err = binary.Write(outFile, ft.byteOrder, ascCount)
		if err != nil {
			mylog.Log("task", "file024", "binary.Write", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 25, "dat": "2", "ans": "2" }
func (ft *fileTask) file25() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file025", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var elementsCount = int(info.Size()) / 8
	var data float64
	for index := 0; index < elementsCount; index++ {
		err = binary.Read(file, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file025", "binary.Read", err)
			return false
		}
		data *= data
		file.Seek(-8, 1)
		err = binary.Write(file, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file025", "binary.Write", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 26, "dat": "2", "ans": "2" }
func (ft *fileTask) file26() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file026", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var elementsCount = info.Size() / 8
	var index int64
	var maximalIndex, minimalIndex int64 = -1, -1
	var maximal, minimal, data float64
	for index = 0; index < elementsCount; index++ {
		err = binary.Read(file, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file026", "binary.Read", err)
			return false
		}
		if index == 0 {
			minimal, maximal = data, data
			minimalIndex, maximalIndex = index, index
		} else if data < minimal {
			minimal = data
			minimalIndex = index
		} else if data > maximal {
			maximal = data
			maximalIndex = index
		}
	}
	if minimalIndex >= 0 && maximalIndex >= 0 {
		file.Seek(minimalIndex*8, 0)
		err = binary.Write(file, ft.byteOrder, maximal)
		if err != nil {
			mylog.Log("task", "file026", "binary.Write(maximal)", err)
			return false
		}
		file.Seek(maximalIndex*8, 0)
		err = binary.Write(file, ft.byteOrder, minimal)
		if err != nil {
			mylog.Log("task", "file026", "binary.Write(minimal)", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 27, "dat": "", "ans": "" }
func (ft *fileTask) file27() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file027", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var elementsCount = int(info.Size()) / 4
	var array = make([]int32, elementsCount)
	for indeX := 0; indeX < elementsCount; indeX++ {
		err = binary.Read(file, ft.byteOrder, &array[indeX])
		if err != nil {
			mylog.Log("task", "file027", "binary.Read", err)
			return false
		}
	}
	file.Seek(0, 0)
	for indeX := 0; indeX < elementsCount/2; indeX++ {
		err = binary.Write(file, ft.byteOrder, array[indeX])
		if err != nil {
			mylog.Log("task", "file027", "binary.Write(array[indeX])", err)
			return false
		}
		err = binary.Write(file, ft.byteOrder, array[elementsCount-1-indeX])
		if err != nil {
			mylog.Log("task", "file027", "binary.Write(array[elementsCount-1-indeX])", err)
			return false
		}
	}
	if elementsCount%2 != 0 {
		err = binary.Write(file, ft.byteOrder, array[elementsCount/2])
		if err != nil {
			mylog.Log("task", "file027", "binary.Write(array[elementsCount/2-1])", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 28, "dat": "2", "ans": "2" }
func (ft *fileTask) file28() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file028", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var elementsCount = info.Size() / 8
	var prev, curr, next, average float64
	var indeX int64
	for indeX = 0; indeX < elementsCount; indeX++ {
		err = binary.Read(file, ft.byteOrder, &next)
		if err != nil {
			mylog.Log("task", "file028", "binary.Read", err)
			return false
		}
		if indeX > 1 {
			average = (prev + curr + next) / 3
			file.Seek(-16, 1)
			err = binary.Write(file, ft.byteOrder, average)
			if err != nil {
				mylog.Log("task", "file028", "file.Write", err)
				return false
			}
			file.Seek(8, 1)
		}
		prev = curr
		curr = next
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 29, "dat": "", "ans": "" }
func (ft *fileTask) file29() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file029", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	file.Truncate(50 * 4)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 30, "dat": "", "ans": "" }
func (ft *fileTask) file30() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file030", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var halfFileSize = info.Size() / 4 / 2 * 4
	file.Truncate(halfFileSize)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 31, "dat": "", "ans": "" }
func (ft *fileTask) file31() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file031", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var data = make([]byte, 50*4)
	file.ReadAt(data, info.Size()-(50*4))
	file.Truncate(0)
	file.Write(data)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 32, "dat": "", "ans": "" }
func (ft *fileTask) file32() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file032", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var fileHalfSize = info.Size() / 2
	var data = make([]byte, fileHalfSize)
	file.ReadAt(data, info.Size()-fileHalfSize)
	file.Truncate(0)
	file.Write(data)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 33, "dat": "", "ans": "" }
func (ft *fileTask) file33() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file033", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_file33.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file033", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var data int32
	for {
		err = binary.Read(file, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file033", "binary.Read", err)
			return false
		}
		err = binary.Write(tempFile, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file033", "binary.Write", err)
			return false
		}
		file.Seek(4, 1)
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 34, "dat": "", "ans": "" }
func (ft *fileTask) file34() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file034", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_file34.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file034", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var data int32
	for {
		err = binary.Read(file, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file034", "binary.Read", err)
			return false
		}
		if data >= 0 {
			err = binary.Write(tempFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file034", "binary.Write", err)
				return false
			}
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 35, "dat": "", "ans": "" }
func (ft *fileTask) file35() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file035", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var elementsCount = int(info.Size()) / 4
	if 50-elementsCount > 0 {
		var tempFileName = "$tempFileName_file35.tmp$"
		tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			mylog.Log("task", "file035", "os.OpenFile(tempFileName)", err)
			return false
		}
		defer tempFile.Close()
		tempFile.Write(make([]byte, (50-elementsCount)*4))
		io.Copy(tempFile, file)
		os.Remove(fileName)
		os.Rename(tempFileName, fileName)
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 36, "dat": "", "ans": "" }
func (ft *fileTask) file36() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file036", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var fileSize = info.Size()
	var position int64
	var data int32
	for position = 0; position < fileSize; position += 4 {
		file.Seek(position, 0)
		err = binary.Read(file, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file036", "binary.Read", err)
			return false
		}
		file.Seek(0, 2)
		err = binary.Write(file, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file036", "binary.Write", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 37, "dat": "", "ans": "" }
func (ft *fileTask) file37() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file037", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var data int32
	for position := info.Size() - 4; position >= 0; position -= 4 {
		file.Seek(position, 0)
		err = binary.Read(file, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file037", "binary.Read", err)
			return false
		}
		file.Seek(0, 2)
		err = binary.Write(file, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file037", "binary.Write", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 38, "dat": "", "ans": "" }
func (ft *fileTask) file38() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file038", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_file38.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file038", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var data int32
	var indeX = 1
	for {
		err = binary.Read(file, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file038", "binary.Read", err)
			return false
		}
		err = binary.Write(tempFile, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file038", "binary.Write", err)
			return false
		}
		if indeX%2 != 0 {
			err = binary.Write(tempFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file038", "binary.Write", err)
				return false
			}
		}
		indeX++
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 39, "dat": "", "ans": "" }
func (ft *fileTask) file39() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file039", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_file39.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file039", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var data int32
	for {
		err = binary.Read(file, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file039", "binary.Read", err)
			return false
		}
		err = binary.Write(tempFile, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file039", "binary.Write", err)
			return false
		}
		if data >= 5 && data <= 10 {
			err = binary.Write(tempFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file039", "binary.Write", err)
				return false
			}
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 40, "dat": "", "ans": "" }
func (ft *fileTask) file40() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file040", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_file40.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file040", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var data int32
	var indeX = 1
	for {
		err = binary.Read(file, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file040", "binary.Read", err)
			return false
		}
		if indeX%2 == 0 {
			data = 0
			err = binary.Write(tempFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file040", "binary.Write", err)
				return false
			}
		}
		err = binary.Write(tempFile, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file040", "binary.Write", err)
			return false
		}
		indeX++
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 41, "dat": "", "ans": "" }
func (ft *fileTask) file41() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file041", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	var tempFileName = "$tempFileName_file40.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file041", "os.OpenFile(tempFileName)", err)
		return false
	}
	defer tempFile.Close()
	var data int32
	for {
		err = binary.Read(file, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file041", "binary.Read", err)
			return false
		}
		if data > 0 {
			data = 0
			err = binary.Write(tempFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file041", "binary.Write", err)
				return false
			}
			err = binary.Write(tempFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file041", "binary.Write", err)
				return false
			}
		}
		err = binary.Write(tempFile, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file041", "binary.Write", err)
			return false
		}
	}
	os.Remove(fileName)
	os.Rename(tempFileName, fileName)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 42, "dat": "2", "ans": "2" }
func (ft *fileTask) file42() bool {
	var fileName1 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var fileName2 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file1, err = os.OpenFile(fileName1, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file042", "os.OpenFile(fileName1)", err)
		return false
	}
	defer file1.Close()
	var info, _ = file1.Stat()
	var file1Size = info.Size()
	file2, err := os.OpenFile(fileName2, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file042", "os.OpenFile(fileName2)", err)
		return false
	}
	defer file2.Close()
	info, _ = file2.Stat()
	var file2Size = info.Size()
	var (
		data1, data2 byte
		err1, err2   error
	)
	for {
		if err1 != io.EOF {
			err1 = binary.Read(file1, ft.byteOrder, &data1)
			if err1 != nil && err1 != io.EOF {
				mylog.Log("task", "file042", "binary.Read(file1) error:", err1)
				return false
			}
		}
		if err2 != io.EOF {
			err2 = binary.Read(file2, ft.byteOrder, &data2)
			if err2 != nil && err2 != io.EOF {
				mylog.Log("task", "file042", "binary.Read(file2) error:", err2)
				return false
			}
		}
		if err1 == io.EOF && err2 == io.EOF {
			break
		}
		if err1 != io.EOF {
			if err2 != io.EOF {
				file2.Seek(-1, 1)
			}
			err = binary.Write(file2, ft.byteOrder, data1)
			if err != nil {
				mylog.Log("task", "file042", "binary.Write(file2)", err)
				return false
			}
		}
		if err2 != io.EOF {
			if err1 != io.EOF {
				file1.Seek(-1, 1)
			}
			err = binary.Write(file1, ft.byteOrder, data2)
			if err != nil {
				mylog.Log("task", "file042", "binary.Write(file1)", err)
				return false
			}
		}
	}
	file1.Truncate(file2Size)
	file2.Truncate(file1Size)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName1, fileName2)
}

// { "no": 43, "dat": "2", "ans": "2" }
func (ft *fileTask) file43() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var s = ft.scanner.NextLine()
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file043", "os.OpenFile(fileName)", err)
		return false
	}
	defer file.Close()
	fileOut, err := os.OpenFile(s, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file043", "os.OpenFile(s)", err)
		return false
	}
	defer fileOut.Close()
	var data byte
	for {
		err = binary.Read(file, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file043", "binary.Read", err)
			return false
		}
		err = binary.Write(fileOut, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file043", "binary.Write", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, s)
}

// { "no": 44, "dat": "2", "ans": "2" }
func (ft *fileTask) file44() bool {
	var fileName1 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var fileName2 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var fileName3 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file1, file2, file3 *os.File
	var err error
	file1, err = os.OpenFile(fileName1, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file044", "os.OpenFile(fileName1)", err)
		return false
	}
	defer file1.Close()
	file2, err = os.OpenFile(fileName2, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file044", "os.OpenFile(fileName2)", err)
		return false
	}
	defer file2.Close()
	file3, err = os.OpenFile(fileName3, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file044", "os.OpenFile(fileName3)", err)
		return false
	}
	defer file3.Close()
	var info, _ = file1.Stat()
	var file1Size = info.Size()
	info, _ = file2.Stat()
	var file2Size = info.Size()
	info, _ = file3.Stat()
	var file3Size = info.Size()
	var fileFrom, fileTo *os.File
	if file1Size < file2Size && file1Size < file3Size {
		fileFrom = file1
	} else if file2Size < file3Size {
		fileFrom = file2
	} else {
		fileFrom = file3
	}
	if file1Size > file2Size && file1Size > file3Size {
		fileTo = file1
	} else if file2Size > file3Size {
		fileTo = file2
	} else {
		fileTo = file3
	}
	fileTo.Truncate(0)
	io.Copy(fileTo, fileFrom)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName1, fileName2, fileName3)
}

// { "no": 45, "dat": "2", "ans": "2" }
func (ft *fileTask) file45() bool {
	var fileName1 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var fileName2 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var fileName3 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file1, file2, file3 *os.File
	var err error
	file1, err = os.OpenFile(fileName1, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file045", "os.OpenFile(fileName1)", err)
		return false
	}
	defer file1.Close()
	file2, err = os.OpenFile(fileName2, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file045", "os.OpenFile(fileName2)", err)
		return false
	}
	defer file2.Close()
	file3, err = os.OpenFile(fileName3, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file045", "os.OpenFile(fileName3)", err)
		return false
	}
	defer file3.Close()
	var info, _ = file1.Stat()
	var file1Size = info.Size()
	info, _ = file2.Stat()
	var file2Size = info.Size()
	info, _ = file3.Stat()
	var file3Size = info.Size()
	var fileFrom, fileTo *os.File
	if file1Size > file2Size && file1Size > file3Size {
		fileFrom = file1
	} else if file2Size > file3Size {
		fileFrom = file2
	} else {
		fileFrom = file3
	}
	if file1Size < file2Size && file1Size < file3Size {
		fileTo = file1
	} else if file2Size < file3Size {
		fileTo = file2
	} else {
		fileTo = file3
	}
	fileTo.Truncate(0)
	io.Copy(fileTo, fileFrom)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName1, fileName2, fileName3)
}

// { "no": 46, "dat": "2", "ans": "2" }
func (ft *fileTask) file46() bool {
	var s0 = ft.scanner.NextLine()
	var n = ft.scanner.NextInt()
	ft.scanner.Skip()
	fileTo, err := os.OpenFile(s0, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file046", "os.OpenFile(s0)", err)
		return false
	}
	defer fileTo.Close()
	var fileFrom *os.File
	var s string
	for indeX := 0; indeX < n; indeX++ {
		s = ft.scanner.NextBinaryFileName(ft.byteOrder)
		fileFrom, err = os.OpenFile(s, os.O_RDONLY, 0777)
		if err != nil {
			mylog.Log("task", "file046", "os.OpenFile", err)
			return false
		}
		defer fileFrom.Close()
		io.Copy(fileTo, fileFrom)
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, s0)
}

// { "no": 47, "dat": "2", "ans": "2" }
func (ft *fileTask) file47() bool {
	var fileName1 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var fileName2 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file1, err = os.OpenFile(fileName1, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file047", "os.OpenFile(fileName1)", err)
		return false
	}
	defer file1.Close()
	file2, err := os.OpenFile(fileName2, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file047", "os.OpenFile(fileName2)", err)
		return false
	}
	defer file2.Close()
	var info, _ = file1.Stat()
	var file1Size = info.Size()
	file1.Seek(0, 2)
	io.Copy(file1, file2)
	file1.Seek(0, 0)
	io.CopyN(file2, file1, file1Size)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName1, fileName2)
}

// { "no": 48, "dat": "", "ans": "" }
func (ft *fileTask) file48() bool {
	var s string
	var files [3]*os.File
	var err error
	for indeX := 0; indeX < 3; indeX++ {
		s = ft.scanner.NextBinaryFileName(ft.byteOrder)
		files[indeX], err = os.OpenFile(s, os.O_RDONLY, 0777)
		if err != nil {
			mylog.Log("task", "file048", "os.OpenFile", err)
		}
		defer files[indeX].Close()
	}
	var sD = ft.scanner.NextLine()
	fileD, err := os.OpenFile(sD, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file048", "os.OpenFile(sD)", err)
		return false
	}
	defer fileD.Close()
	var data int32
LOOP:
	for {
		for indeX := 0; indeX < 3; indeX++ {
			err = binary.Read(files[indeX], ft.byteOrder, &data)
			if err == io.EOF {
				break LOOP
			} else if err != nil {
				mylog.Log("task", "file048", "binary.Read", err)
				return false
			}
			err = binary.Write(fileD, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file048", "binary.Write(fileD)", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, sD)
}

// { "no": 49, "dat": "", "ans": "" }
func (ft *fileTask) file49() bool {
	var s string
	var files [4]*os.File
	var err error
	for indeX := 0; indeX < 4; indeX++ {
		s = ft.scanner.NextBinaryFileName(ft.byteOrder)
		files[indeX], err = os.OpenFile(s, os.O_RDONLY, 0777)
		if err != nil {
			mylog.Log("task", "file049", "os.OpenFile", err)
			return false
		}
		defer files[indeX].Close()
	}
	var sE = ft.scanner.NextLine()
	fileE, err := os.OpenFile(sE, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file049", "os.OpenFile(sE)", err)
		return false
	}
	defer fileE.Close()
	var data [4]int32
LOOP:
	for {
		for indeX := 0; indeX < 4; indeX++ {
			err = binary.Read(files[indeX], ft.byteOrder, &data[indeX])
			if err == io.EOF {
				break LOOP
			} else if err != nil {
				mylog.Log("task", "file049", "binary.Read", err)
				return false
			}
		}
		for indeX := 0; indeX < 4; indeX++ {
			err = binary.Write(fileE, ft.byteOrder, data[indeX])
			if err != nil {
				mylog.Log("task", "file049", "binary.Write", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, sE)
}

// { "no": 50, "dat": "2", "ans": "2" }
func (ft *fileTask) file50() bool {
	var s1 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var s2 = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var s3 = ft.scanner.NextLine()
	var file1, file2, file3 *os.File
	var err error
	file1, err = os.OpenFile(s1, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file050", "os.OpenFile(s1)", err)
		return false
	}
	defer file1.Close()
	file2, err = os.OpenFile(s2, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file050", "os.OpenFile(s2)", err)
		return false
	}
	defer file2.Close()
	file3, err = os.OpenFile(s3, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file050", "os.OpenFile(s3)", err)
		return false
	}
	defer file3.Close()
	var data1, data2, data float64
	var err1, err2 error
	for {
		if err1 != io.EOF {
			err1 = binary.Read(file1, ft.byteOrder, &data1)
			if err1 != nil && err1 != io.EOF {
				mylog.Log("task", "file050", "binary.Read(file1) error:", err1)
				return false
			}
		}
		if err2 != io.EOF {
			err2 = binary.Read(file2, ft.byteOrder, &data2)
			if err2 != nil && err2 != io.EOF {
				mylog.Log("task", "file050", "binary.Read(file2) error:", err2)
				return false
			}
		}
		if err1 == io.EOF && err2 == io.EOF {
			break
		} else if err1 == io.EOF {
			data = data2
		} else if err2 == io.EOF {
			data = data1
		} else if data1 < data2 {
			data = data1
			file2.Seek(-8, 1)
		} else {
			data = data2
			file1.Seek(-8, 1)
		}
		err = binary.Write(file3, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file050", "binary.Write", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, s3)
}

// { "no": 51, "dat": "2", "ans": "2" }
func (ft *fileTask) file51() bool {
	var s string
	const filesCount = 3
	var files [filesCount]*os.File
	var err error
	for indeX := 0; indeX < 3; indeX++ {
		s = ft.scanner.NextBinaryFileName(ft.byteOrder)
		files[indeX], err = os.OpenFile(s, os.O_RDONLY, 0777)
		if err != nil {
			mylog.Log("task", "file051", "os.OpenFile", err)
			return false
		}
		defer files[indeX].Close()
	}
	var s4 = ft.scanner.NextLine()
	file4, err := os.OpenFile(s4, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file051", "os.OpenFile(s4)", err)
		return false
	}
	defer file4.Close()
	var errs [filesCount]error
	var data [filesCount]float64
	var exitable bool
	var maximalIndex int
LOOP:
	for {
		exitable = true
		maximalIndex = -1
		for indeX := 0; indeX < filesCount; indeX++ {
			if errs[indeX] != io.EOF {
				errs[indeX] = binary.Read(files[indeX], ft.byteOrder, &data[indeX])
				if errs[indeX] == nil {
					exitable = false
					if maximalIndex < 0 || data[indeX] > data[maximalIndex] {
						maximalIndex = indeX
					}
				} else if errs[indeX] != nil && errs[indeX] != io.EOF {
					mylog.Log("task", "file051", "binary.Read", errs[indeX])
					return false
				}
			}
		}
		if exitable {
			break LOOP
		}
		err = binary.Write(file4, ft.byteOrder, data[maximalIndex])
		if err != nil {
			mylog.Log("task", "file051", "binary.Write(file4)", err)
			return false
		}
		for indeX := 0; indeX < filesCount; indeX++ {
			if errs[indeX] == io.EOF || indeX == maximalIndex {
				continue
			}
			files[indeX].Seek(-8, 1)
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, s4)
}

// { "no": 52, "dat": "", "ans": "" }
func (ft *fileTask) file52() bool {
	var s0 = ft.scanner.NextLine()
	var n = ft.scanner.NextInt32()
	ft.scanner.Skip()
	var file0, err = os.OpenFile(s0, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file052", "os.OpenFile(s0)", err)
		return false
	}
	defer file0.Close()
	err = binary.Write(file0, ft.byteOrder, n)
	if err != nil {
		mylog.Log("task", "file052", "binary.Write(file0, n)", err)
		return false
	}
	var files = make([]*os.File, n)
	var indeX int32
	for indeX = 0; indeX < n; indeX++ {
		s := ft.scanner.NextBinaryFileName(ft.byteOrder)
		files[indeX], err = os.OpenFile(s, os.O_RDONLY, 0777)
		if err != nil {
			mylog.Log("task", "file052", "os.OpenFile", err)
			return false
		}
		defer files[indeX].Close()
		var info, _ = files[indeX].Stat()
		var fileElements = int32(info.Size()) / 4
		err = binary.Write(file0, ft.byteOrder, fileElements)
		if err != nil {
			mylog.Log("task", "file052", "binary.Write(file0, fileElements)", err)
			return false
		}
	}
	var data int32
	for indeX = 0; indeX < n; indeX++ {
		for {
			err = binary.Read(files[indeX], ft.byteOrder, &data)
			if err == io.EOF {
				break
			} else if err != nil {
				mylog.Log("task", "file052", "binary.Read", err)
				return false
			}
			err = binary.Write(file0, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file052", "binary.Write(file0, data)", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, s0)
}

// { "no": 53, "dat": "", "ans": "" }
func (ft *fileTask) file53() bool {
	var s = ft.scanner.NextLine()
	var n = ft.scanner.NextInt32()
	ft.scanner.Skip()
	var archivalFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(s, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file053", "os.OpenFile(s)", err)
		return false
	}
	defer file.Close()
	archivalFile, err := os.OpenFile(archivalFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file053", "os.OpenFile(archivalFileName)", err)
		return false
	}
	defer archivalFile.Close()
	var filesCount int32
	err = binary.Read(archivalFile, ft.byteOrder, &filesCount)
	if err != nil {
		mylog.Log("task", "file053", "binary.Read(archivalFile, &filesCount)", err)
		return false
	}
	if n > 0 && n <= filesCount {
		var elementsCounts = make([]int32, filesCount)
		var index int32
		for index = 0; index < filesCount; index++ {
			err = binary.Read(archivalFile, ft.byteOrder, &elementsCounts[index])
			if err != nil {
				mylog.Log("task", "file053", "binary.Read(archivalFile)", err)
				return false
			}
		}
		for index = 0; index < n-1; index++ {
			archivalFile.Seek(int64(elementsCounts[index])*4, 1)
		}
		var data int32
		for index = 0; index < elementsCounts[n-1]; index++ {
			err = binary.Read(archivalFile, ft.byteOrder, &data)
			if err != nil {
				mylog.Log("task", "file053", "binary.Read(archivalFile, &data)", err)
				return false
			}
			err = binary.Write(file, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file053", "binary.Write(file)", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, s)
}

// { "no": 54, "dat": "", "ans": "2" }
func (ft *fileTask) file54() bool {
	var s = ft.scanner.NextLine()
	var archivalFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(s, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file054", "os.OpenFile(s)", err)
		return false
	}
	defer file.Close()
	archivalFile, err := os.OpenFile(archivalFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file054", "os.OpenFile(archivalFileName)", err)
		return false
	}
	defer archivalFile.Close()
	var filesCount int32
	err = binary.Read(archivalFile, ft.byteOrder, &filesCount)
	if err != nil {
		mylog.Log("task", "file054", "binary.Read(archivalFile, &filesCount)", err)
		return false
	}
	var elementsCounts = make([]int32, filesCount)
	var indeX int32
	for indeX = 0; indeX < filesCount; indeX++ {
		err = binary.Read(archivalFile, ft.byteOrder, &elementsCounts[indeX])
		if err != nil {
			mylog.Log("task", "file054", "binary.Read(archivalFile)", err)
			return false
		}
	}
	var data, i int32
	var sum, average float64
	for indeX = 0; indeX < filesCount; indeX++ {
		sum = 0
		for i = 0; i < elementsCounts[indeX]; i++ {
			err = binary.Read(archivalFile, ft.byteOrder, &data)
			if err != nil {
				mylog.Log("task", "file054", "binary.Read(archivalFile, &data)", err)
				return false
			}
			sum += float64(data)
		}
		average = sum / float64(elementsCounts[indeX])
		err = binary.Write(file, ft.byteOrder, average)
		if err != nil {
			mylog.Log("task", "file054", "binary.Write(file, average)", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, s)
}

// { "no": 55, "dat": "", "ans": "" }
func (ft *fileTask) file55() bool {
	var s0 = ft.scanner.NextLine()
	var n = ft.scanner.NextInt()
	ft.scanner.Skip()
	var s string
	var files = make([]*os.File, n)
	var err error
	for indeX := 0; indeX < n; indeX++ {
		s = ft.scanner.NextBinaryFileName(ft.byteOrder)
		files[indeX], err = os.OpenFile(s, os.O_RDONLY, 0777)
		if err != nil {
			mylog.Log("task", "file055", "os.OpenFile(s)", err)
			return false
		}
		defer files[indeX].Close()
	}
	file, err := os.OpenFile(s0, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file055", "os.OpenFile(s0)", err)
		return false
	}
	defer file.Close()
	var info os.FileInfo
	var elementsCount, data, i int32
	for indeX := 0; indeX < n; indeX++ {
		info, _ = files[indeX].Stat()
		elementsCount = int32(info.Size()) / 4
		err = binary.Write(file, ft.byteOrder, elementsCount)
		if err != nil {
			mylog.Log("task", "file055", "binary.Write(file, elementsCount)", err)
			return false
		}
		for i = 0; i < elementsCount; i++ {
			err = binary.Read(files[indeX], ft.byteOrder, &data)
			if err != nil {
				mylog.Log("task", "file055", "binary.Read(files[indeX], &data)", err)
				return false
			}
			err = binary.Write(file, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file055", "binary.Write(file, data)", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, s0)
}

// { "no": 56, "dat": "", "ans": "" }
func (ft *fileTask) file56() bool {
	var s = ft.scanner.NextLine()
	var n = ft.scanner.NextInt()
	ft.scanner.Skip()
	var archivalFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(s, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file056", "os.OpenFile(s)", err)
		return false
	}
	defer file.Close()
	archivalFile, err := os.OpenFile(archivalFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file056", "os.OpenFile(archivalFileName)", err)
		return false
	}
	defer archivalFile.Close()
	var elementsCount int32
	for indeX := 1; true; indeX++ {
		err = binary.Read(archivalFile, ft.byteOrder, &elementsCount)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file056", "binary.Read(archivalFile, &elementsCount)", err)
			return false
		}
		if indeX == n {
			io.CopyN(file, archivalFile, int64(elementsCount)*4)
			break
		}
		archivalFile.Seek(int64(elementsCount)*4, 1)
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, s)
}

// { "no": 57, "dat": "", "ans": "" }
func (ft *fileTask) file57() bool {
	var s1 = ft.scanner.NextLine()
	var s2 = ft.scanner.NextLine()
	var archivalFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file1, file2, archivalFile *os.File
	var err error
	file1, err = os.OpenFile(s1, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file057", "os.OpenFile(s1)", err)
		return false
	}
	defer file1.Close()
	file2, err = os.OpenFile(s2, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file057", "os.OpenFile(s2)", err)
		return false
	}
	defer file2.Close()
	archivalFile, err = os.OpenFile(archivalFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file057", "os.OpenFile(archivalFileName)", err)
		return false
	}
	defer archivalFile.Close()
	var elementsCount, data int32
	for {
		err = binary.Read(archivalFile, ft.byteOrder, &elementsCount)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file057", "binary.Read(archivalFile, &elementsCount)", err)
			return false
		}
		if elementsCount > 2 {
			err = binary.Read(archivalFile, ft.byteOrder, &data)
			if err != nil {
				mylog.Log("task", "file057", "binary.Read(archivalFile, &firstData)", err)
				return false
			}
			err = binary.Write(file1, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file057", "binary.Write(file1, data)", err)
				return false
			}
			archivalFile.Seek((int64(elementsCount)-2)*4, 1)
			err = binary.Read(archivalFile, ft.byteOrder, &data)
			if err != nil {
				mylog.Log("task", "file057", "binary.Read(archivalFile, &lastData)", err)
				return false
			}
			err = binary.Write(file2, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file057", "binary.Write(file2, data)", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, s1, s2)
}

// { "no": 58, "dat": "", "ans": "" }
func (ft *fileTask) file58() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file058", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var data byte
	var position int64
	for {
		err = binary.Read(file, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file058", "binary.Read", err)
			return false
		}
		if data == ' ' {
			break
		}
		position++
	}
	file.Truncate(position)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 59, "dat": "", "ans": "" }
func (ft *fileTask) file59() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file059", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var data byte
	var position, lastBlankPosition int64 = 0, -1
	for {
		err = binary.Read(file, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file059", "binary.Read", err)
			return false
		}
		position++
		if data == ' ' {
			lastBlankPosition = position
		}
	}
	if lastBlankPosition > 0 {
		file.Truncate(lastBlankPosition - 1)
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 60, "dat": "", "ans": "" }
func (ft *fileTask) file60() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file060", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var position int64
	var data byte
	for {
		err = binary.Read(file, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file060", "binary.Read", err)
			return false
		}
		position++
		if data == ' ' {
			break
		}
	}
	var info, _ = file.Stat()
	var fileSize = info.Size()
	var farq = fileSize - position
	var indeX int64
	for indeX = 0; indeX < farq; indeX++ {
		file.Seek(position+indeX, 0)
		err = binary.Read(file, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file060", "binary.Read", err)
			return false
		}
		file.Seek(indeX, 0)
		err = binary.Write(file, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file060", "binary.Write", err)
			return false
		}
	}
	file.Truncate(farq)
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 61, "dat": "", "ans": "" }
func (ft *fileTask) file61() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file061", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var position, lastPosition int64 = 0, -1
	var data byte
	for {
		err = binary.Read(file, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file061", "binary.Read", err)
			return false
		}
		position++
		if data == ' ' {
			lastPosition = position
		}
	}
	if lastPosition > 0 {
		var info, _ = file.Stat()
		var fileSize = info.Size()
		var farq = fileSize - lastPosition
		var indeX int64
		for indeX = 0; indeX < farq; indeX++ {
			file.Seek(lastPosition+indeX, 0)
			err = binary.Read(file, ft.byteOrder, &data)
			if err != nil {
				mylog.Log("task", "file061", "binary.Read", err)
				return false
			}
			file.Seek(indeX, 0)
			err = binary.Write(file, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file061", "binary.Write", err)
				return false
			}
		}
		file.Truncate(farq)
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 62, "dat": "", "ans": "" }
func (ft *fileTask) file62() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file062", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var isSorted bool
	var prev, curr byte
	var index int
	for {
		isSorted = true
		file.Seek(0, 0)
		index = 0
		for {
			err = binary.Read(file, ft.byteOrder, &curr)
			if err == io.EOF {
				break
			} else if err != nil {
				mylog.Log("task", "file062", "binary.Read", err)
				return false
			}
			index++
			if index > 1 && prev > curr {
				file.Seek(-2, 1)
				err = binary.Write(file, ft.byteOrder, curr)
				if err != nil {
					mylog.Log("task", "file062", "binary.Write(file, curr)", err)
					return false
				}
				err = binary.Write(file, ft.byteOrder, prev)
				if err != nil {
					mylog.Log("task", "file062", "binary.Write(file, prev)", err)
					return false
				}
				curr = prev
				isSorted = false
			}
			prev = curr
		}
		if isSorted {
			break
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, fileName)
}

// { "no": 63, "dat": "", "ans": "" }
func (ft *fileTask) file63() bool {
	var k = ft.scanner.NextInt()
	ft.scanner.Skip()
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName1 = ft.scanner.NextLine()
	var outFileName2 = ft.scanner.NextLine()
	var inFile, outFile1, outFile2 *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file063", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile1, err = os.OpenFile(outFileName1, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file063", "os.OpenFile(outFileName1)", err)
		return false
	}
	defer outFile1.Close()
	outFile2, err = os.OpenFile(outFileName2, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file063", "os.OpenFile(outFileName2)", err)
		return false
	}
	defer outFile2.Close()
	for {
		var data = make([]byte, 80)
		err = binary.Read(inFile, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file063", "binary.Read", err)
			return false
		}
		length := bytes.IndexByte(data, 0)
		var character byte = ' '
		if length >= k && k <= 80 {
			data[k] = 0
			character = data[k-1]
		}
		err = binary.Write(outFile1, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file063", "binary.Write(outFile1)", err)
			return false
		}
		err = binary.Write(outFile2, ft.byteOrder, character)
		if err != nil {
			mylog.Log("task", "file063", "binary.Write(outFile2)", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName1, outFileName2)
}

// { "no": 64, "dat": "", "ans": "" }
func (ft *fileTask) file64() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file064", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file064", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var length, minLength = 0, 0
	for {
		var data = make([]byte, 80)
		err = binary.Read(inFile, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file064", "binary.Read(inFile)", err)
			return false
		}
		data = bytes.TrimRight(data, " ")
		length = len(data)
		if length != 0 && (minLength == 0 || length < minLength) {
			minLength = length
		}
	}
	inFile.Seek(0, 0)
	for {
		var data = make([]byte, 80)
		err = binary.Read(inFile, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file064", "binary.Read", err)
			return false
		}
		if len(bytes.TrimRight(data, " ")) == minLength {
			err = binary.Write(outFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file064", "binary.Write", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 65, "dat": "", "ans": "" }
func (ft *fileTask) file65() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file065", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file065", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var length, maxLength = 0, 0
	for {
		var data = make([]byte, 80)
		err = binary.Read(inFile, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file065", "binary.Read(inFile)", err)
			return false
		}
		data = bytes.TrimRight(data, " ")
		length = len(data)
		if length > maxLength {
			maxLength = length
		}
	}
	var info, _ = inFile.Stat()
	var fileSize = info.Size()
	var position = fileSize - 80
	for position >= 0 {
		inFile.Seek(position, 0)
		var data = make([]byte, 80)
		err = binary.Read(inFile, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file065", "binary.Read", err)
			return false
		}
		if len(bytes.TrimRight(data, " ")) == maxLength {
			err = binary.Write(outFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file065", "binary.Write", err)
				return false
			}
		}
		position -= 80
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 66, "dat": "", "ans": "" }
func (ft *fileTask) file66() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, err = os.OpenFile(inFileName, os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file066", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file066", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	io.Copy(outFile, inFile)
	var prev, curr []byte
	var isSorted bool
	var index int
	for {
		index = 0
		isSorted = true
		outFile.Seek(0, 0)
		for {
			curr = make([]byte, 80)
			err = binary.Read(outFile, ft.byteOrder, &curr)
			if err == io.EOF {
				break
			} else if err != nil {
				mylog.Log("task", "file066", "binary.Read", err)
				return false
			}
			index++
			if index > 1 && bytes.Compare(prev, curr) > 0 {
				outFile.Seek(-160, 1)
				err = binary.Write(outFile, ft.byteOrder, curr)
				if err != nil {
					mylog.Log("task", "file066", "binary.Write(file, curr)", err)
					return false
				}
				err = binary.Write(outFile, ft.byteOrder, prev)
				if err != nil {
					mylog.Log("task", "file066", "binary.Write(file, prev)", err)
					return false
				}
				curr = prev
				isSorted = false
			}
			prev = curr
		}
		if isSorted {
			break
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 67, "dat": "", "ans": "" }
func (ft *fileTask) file67() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName1 = ft.scanner.NextLine()
	var outFileName2 = ft.scanner.NextLine()
	var inFile, outFile1, outFile2 *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file067", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile1, err = os.OpenFile(outFileName1, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file067", "os.OpenFile(outFileName1)", err)
		return false
	}
	defer outFile1.Close()
	outFile2, err = os.OpenFile(outFileName2, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file067", "os.OpenFile(outFileName2)", err)
		return false
	}
	defer outFile2.Close()
	var day, month int
	for {
		var data [80]byte
		err = binary.Read(inFile, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file067", "binary.Read", err)
			return false
		}
		day, _ = strconv.Atoi(string(data[:2]))
		month, _ = strconv.Atoi(string(data[3:5]))
		err = binary.Write(outFile1, ft.byteOrder, int32(day))
		if err != nil {
			mylog.Log("task", "file067", "binary.Write(outFile1)", err)
			return false
		}
		err = binary.Write(outFile2, ft.byteOrder, int32(month))
		if err != nil {
			mylog.Log("task", "file067", "binary.Write(outFile2)", err)
			return false
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName1, outFileName2)
}

// { "no": 68, "dat": "", "ans": "" }
func (ft *fileTask) file68() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName1 = ft.scanner.NextLine()
	var outFileName2 = ft.scanner.NextLine()
	var inFile, outFile1, outFile2 *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file068", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile1, err = os.OpenFile(outFileName1, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file068", "os.OpenFile(outFileName1)", err)
		return false
	}
	defer outFile1.Close()
	outFile2, err = os.OpenFile(outFileName2, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file068", "os.OpenFile(outFileName2)", err)
		return false
	}
	defer outFile2.Close()
	var info, _ = inFile.Stat()
	var position = info.Size() - 80
	var month, year int
	for position >= 0 {
		inFile.Seek(position, 0)
		var data [80]byte
		err = binary.Read(inFile, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file068", "binary.Read", err)
			return false
		}
		month, _ = strconv.Atoi(string(data[3:5]))
		year, _ = strconv.Atoi(string(data[6:10]))
		err = binary.Write(outFile1, ft.byteOrder, int32(month))
		if err != nil {
			mylog.Log("task", "file068", "binary.Write(outFile1)", err)
			return false
		}
		err = binary.Write(outFile2, ft.byteOrder, int32(year))
		if err != nil {
			mylog.Log("task", "file068", "binary.Write(outFile2)", err)
			return false
		}
		position -= 80
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName1, outFileName2)
}

// { "no": 69, "dat": "", "ans": "" }
func (ft *fileTask) file69() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file069", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file069", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var month int
	for {
		var data [80]byte
		err = binary.Read(inFile, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file069", "binary.Read", err)
			return false
		}
		month, _ = strconv.Atoi(string(data[3:5]))
		switch month {
		case 6, 7, 8:
			err = binary.Write(outFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file069", "binary.Write", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 70, "dat": "", "ans": "" }
func (ft *fileTask) file70() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file070", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file070", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var info, _ = inFile.Stat()
	var position = info.Size() - 80
	var month int
	for position >= 0 {
		inFile.Seek(position, 0)
		var data [80]byte
		err = binary.Read(inFile, ft.byteOrder, &data)
		if err != nil {
			mylog.Log("task", "file070", "binary.Read", err)
			return false
		}
		month, _ = strconv.Atoi(string(data[3:5]))
		switch month {
		case 1, 2, 12:
			err = binary.Write(outFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file070", "binary.Write", err)
				return false
			}
		}
		position -= 80
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 71, "dat": "", "ans": "" }
func (ft *fileTask) file71() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file071", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var sana time.Time
	var theDate string
	var minUnixNano int64
	for {
		var data = make([]byte, 80)
		err = binary.Read(file, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file071", "binary.Read", err)
			return false
		}
		data = data[:10]
		sana, err = time.Parse("02/01/2006", string(data))
		if err != nil {
			mylog.Log("task", "file071", "time.Parse", err)
			return false
		}
		switch sana.Month() {
		case time.March, time.April, time.May:
			if minUnixNano == 0 || sana.UnixNano() < minUnixNano {
				minUnixNano = sana.UnixNano()
				theDate = string(data)
			}
		}
	}
	return ft.checker.CompareWord(theDate)
}

// { "no": 72, "dat": "", "ans": "" }
func (ft *fileTask) file72() bool {
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file072", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var sana time.Time
	var theDate string
	var maxUnixNano int64
	for {
		var data = make([]byte, 80)
		err = binary.Read(file, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file072", "binary.Read", err)
			return false
		}
		data = data[:10]
		sana, err = time.Parse("02/01/2006", string(data))
		if err != nil {
			mylog.Log("task", "file072", "time.Parse", err)
			return false
		}
		switch sana.Month() {
		case time.September, time.October, time.November:
			if maxUnixNano == 0 || sana.UnixNano() > maxUnixNano {
				maxUnixNano = sana.UnixNano()
				theDate = string(data)
			}
		}
	}
	return ft.checker.CompareWord(theDate)
}

// { "no": 73, "dat": "", "ans": "" }
func (ft *fileTask) file73() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file073", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		mylog.Log("task", "file073", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	for {
		var data = make([]byte, 80)
		err = binary.Read(inFile, ft.byteOrder, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("task", "file073", "binary.Read(inFile)", err)
			return false
		}
		err = binary.Write(outFile, ft.byteOrder, data)
		if err != nil {
			mylog.Log("task", "file073", "binary.Write(outFile)", err)
			return false
		}
	}
	var isSorted bool
	var prev, curr []byte
	var prevDate, currDate time.Time
	var index int
	for {
		isSorted = true
		index = 0
		outFile.Seek(0, 0)
		for {
			curr = make([]byte, 80)
			err = binary.Read(outFile, ft.byteOrder, &curr)
			if err == io.EOF {
				break
			} else if err != nil {
				mylog.Log("task", "file073", "binary.Read(outFile)", err)
				return false
			}
			currDate, err = time.Parse("02/01/2006", string(curr[:10]))
			if err != nil {
				mylog.Log("task", "file073", "time.Parse", err)
				return false
			}
			index++
			if index > 1 && prevDate.Before(currDate) {
				isSorted = false
				outFile.Seek(-160, 1)
				err = binary.Write(outFile, ft.byteOrder, curr)
				if err != nil {
					mylog.Log("task", "file073", "binary.Write", err)
					return false
				}
				err = binary.Write(outFile, ft.byteOrder, prev)
				if err != nil {
					mylog.Log("task", "file073", "binary.Write", err)
					return false
				}
				curr = prev
				currDate = prevDate
			}
			prev = curr
			prevDate = currDate
		}
		if isSorted {
			break
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 74, "dat": "2", "ans": "2" }
func (ft *fileTask) file74() bool {
	var i = ft.scanner.NextInt()
	var j = ft.scanner.NextInt()
	ft.scanner.Skip()
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file074", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var elementsCount = info.Size() / 8
	var matrixOrder = int(math.Sqrt(float64(elementsCount)))
	var element float64
	if i > 0 && i <= matrixOrder && j > 0 && j <= matrixOrder {
		skipElementsCount := (i-1)*matrixOrder + j - 1
		file.Seek(int64(skipElementsCount*8), 0)
		err = binary.Read(file, ft.byteOrder, &element)
		if err != nil {
			mylog.Log("task", "file074", "binary.Read", err)
			return false
		}
	}
	return ft.checker.CompareFloat64(2, element)
}

// { "no": 75, "dat": "2", "ans": "2" }
func (ft *fileTask) file75() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file075", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file075", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var info, _ = inFile.Stat()
	var matrixOrder = int64(math.Sqrt(float64(info.Size() / 8)))
	var indeX, position int64
	var data float64
	for indeX = 0; indeX < matrixOrder; indeX++ {
		for position = 0; position < matrixOrder; position++ {
			inFile.Seek((position*matrixOrder+indeX)*8, 0)
			err = binary.Read(inFile, ft.byteOrder, &data)
			if err != nil {
				mylog.Log("task", "file075", "binary.Read", err)
				return false
			}
			err = binary.Write(outFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file075", "binary.Write", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 76, "dat": "2", "ans": "2" }
func (ft *fileTask) file76() bool {
	var sA = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var sB = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var sC = ft.scanner.NextLine()
	var fileA, fileB, fileC *os.File
	var err error
	fileA, err = os.OpenFile(sA, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file076", "os.OpenFile(sA)", err)
		return false
	}
	defer fileA.Close()
	var info, _ = fileA.Stat()
	var matrixAOrder = int64(math.Sqrt(float64(info.Size() / 8)))
	fileB, err = os.OpenFile(sB, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file076", "os.OpenFile(sB)", err)
		return false
	}
	defer fileB.Close()
	info, _ = fileB.Stat()
	var matrixBOrder = int64(math.Sqrt(float64(info.Size() / 8)))
	fileC, err = os.OpenFile(sC, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file076", "os.OpenFile(sC)", err)
		return false
	}
	defer fileC.Close()
	if matrixAOrder == matrixBOrder {
		var col, position, row int64
		var dataA, dataB, dataC float64
		for row = 0; row < matrixAOrder; row++ {
			for col = 0; col < matrixAOrder; col++ {
				dataC = 0
				for position = 0; position < matrixAOrder; position++ {
					fileA.Seek((position+row*matrixAOrder)*8, 0)
					err = binary.Read(fileA, ft.byteOrder, &dataA)
					if err != nil {
						mylog.Log("task", "file076", "binary.Read(fileA)", err)
						return false
					}
					fileB.Seek((position*matrixAOrder+col)*8, 0)
					err = binary.Read(fileB, ft.byteOrder, &dataB)
					if err != nil {
						mylog.Log("task", "file076", "binary.Read(fileB)", err)
						return false
					}
					dataC += dataA * dataB
				}
				err = binary.Write(fileC, ft.byteOrder, dataC)
				if err != nil {
					mylog.Log("task", "file076", "binary.Write(fileC)", err)
					return false
				}
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, sC)
}

// { "no": 77, "dat": "2", "ans": "2" }
func (ft *fileTask) file77() bool {
	var i = ft.scanner.NextInt64()
	var j = ft.scanner.NextInt64()
	ft.scanner.Skip()
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file077", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var element float64
	var rowsCount, colsCount int64
	err = binary.Read(file, ft.byteOrder, &element)
	if err != nil {
		mylog.Log("task", "file077", "binary.Read(file, &colsCount)", err)
		return false
	}
	colsCount, element = int64(element), 0
	var info, _ = file.Stat()
	var elementsCount = info.Size()/8 - 1
	rowsCount = elementsCount / colsCount
	if i > 0 && i <= rowsCount && j > 0 && j <= colsCount {
		skipElementsCount := (i-1)*colsCount + j - 1
		file.Seek(skipElementsCount*8, 1)
		err = binary.Read(file, ft.byteOrder, &element)
		if err != nil {
			mylog.Log("task", "file077", "binary.Read(file, &element)", err)
			return false
		}
	}
	return ft.checker.CompareFloat64(2, element)
}

// { "no": 78, "dat": "2", "ans": "2" }
func (ft *fileTask) file78() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file078", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file078", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var element float64
	err = binary.Read(inFile, ft.byteOrder, &element)
	if err != nil {
		mylog.Log("task", "file078", "binary.Read", err)
		return false
	}
	var colsCount = int64(element)
	var info, _ = inFile.Stat()
	var rowsCount = (info.Size()/8 - 1) / colsCount
	if rowsCount > 0 && colsCount > 0 {
		err = binary.Write(outFile, ft.byteOrder, float64(rowsCount))
		if err != nil {
			mylog.Log("task", "file078", "binary.Write(outFile)", err)
			return false
		}
		// continue;
		var indeX, position int64
		for indeX = 0; indeX < colsCount; indeX++ {
			for position = 0; position < rowsCount; position++ {
				inFile.Seek((position*colsCount+indeX+1)*8, 0)
				err = binary.Read(inFile, ft.byteOrder, &element)
				if err != nil {
					mylog.Log("task", "file078", "binary.Read(inFile)", err)
					return false
				}
				err = binary.Write(outFile, ft.byteOrder, element)
				if err != nil {
					mylog.Log("task", "file078", "binary.Write(outFile)", err)
					return false
				}
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 79, "dat": "2", "ans": "2" }
func (ft *fileTask) file79() bool {
	var sA = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var sB = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var sC = ft.scanner.NextLine()
	var fileA, fileB, fileC *os.File
	var err error
	fileA, err = os.OpenFile(sA, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file079", "os.OpenFile(sA)", err)
		return false
	}
	defer fileA.Close()
	var element float64
	err = binary.Read(fileA, ft.byteOrder, &element)
	if err != nil {
		mylog.Log("task", "file079", "binary.Read(fileA)", err)
		return false
	}
	var colsCountA = int64(element)
	var info, _ = fileA.Stat()
	var rowsCountA = (info.Size()/8 - 1) / colsCountA
	fileB, err = os.OpenFile(sB, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file079", "os.OpenFile(sB)", err)
		return false
	}
	defer fileB.Close()
	err = binary.Read(fileB, ft.byteOrder, &element)
	if err != nil {
		mylog.Log("task", "file079", "binary.Read(fileB)", err)
		return false
	}
	var colsCountB = int64(element)
	info, _ = fileB.Stat()
	var rowsCountB = (info.Size()/8 - 1) / colsCountB
	fileC, err = os.OpenFile(sC, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file079", "os.OpenFile(sC)", err)
		return false
	}
	defer fileC.Close()
	if colsCountA == rowsCountB {
		err = binary.Write(fileC, ft.byteOrder, float64(colsCountB))
		if err != nil {
			mylog.Log("task", "file079", "binary.Write(fileC)", err)
			return false
		}
		var col, position, row int64
		var dataA, dataB, dataC float64
		for row = 0; row < rowsCountA; row++ {
			for col = 0; col < colsCountB; col++ {
				dataC = 0
				for position = 0; position < colsCountA; position++ {
					fileA.Seek((row*colsCountA+position+1)*8, 0)
					err = binary.Read(fileA, ft.byteOrder, &dataA)
					if err != nil {
						mylog.Log("task", "file079", "binary.Read(fileA)", err)
						return false
					}
					fileB.Seek((position*colsCountB+col+1)*8, 0)
					err = binary.Read(fileB, ft.byteOrder, &dataB)
					if err != nil {
						mylog.Log("task", "file079", "binary.Read(fileB)", err)
						return false
					}
					dataC += dataA * dataB
				}
				err = binary.Write(fileC, ft.byteOrder, dataC)
				if err != nil {
					mylog.Log("task", "file079", "binary.Write(fileC)", err)
					return false
				}
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, sC)
}

// { "no": 80, "dat": "2", "ans": "2" }
func (ft *fileTask) file80() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file080", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file080", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var info, _ = inFile.Stat()
	var matrixOrder = int(math.Sqrt(float64(info.Size() / 8)))
	var data float64
	for row := 0; row < matrixOrder; row++ {
		for col := 0; col < matrixOrder; col++ {
			err = binary.Read(inFile, ft.byteOrder, &data)
			if err != nil {
				mylog.Log("task", "file080", "binary.Read", err)
				return false
			}
			if row > col {
				continue
			}
			err = binary.Write(outFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file080", "binary.Write", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 81, "dat": "2", "ans": "2" }
func (ft *fileTask) file81() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file081", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file081", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var info, _ = inFile.Stat()
	var matrixOrder = int(math.Sqrt(float64(info.Size() / 8)))
	var data float64
	for row := 0; row < matrixOrder; row++ {
		for col := 0; col < matrixOrder; col++ {
			err = binary.Read(inFile, ft.byteOrder, &data)
			if err != nil {
				mylog.Log("task", "file081", "binary.Read", err)
				return false
			}
			if row < col {
				continue
			}
			err = binary.Write(outFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file081", "binary.Write", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 82, "dat": "2", "ans": "2" }
func (ft *fileTask) file82() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file082", "os.OpenFile(inFileName)", err)
		return false
	}
	defer inFile.Close()
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file082", "os.OpenFile(outFileName)", err)
		return false
	}
	defer outFile.Close()
	var info, _ = inFile.Stat()
	var matrixOrder = int(math.Sqrt(float64(info.Size() / 8)))
	var data float64
	for row := 0; row < matrixOrder; row++ {
		for col := 0; col < matrixOrder; col++ {
			err = binary.Read(inFile, ft.byteOrder, &data)
			if err != nil {
				mylog.Log("task", "file082", "binary.Read", err)
				return false
			}
			if math.Abs(float64(row-col)) > 1 {
				continue
			}
			err = binary.Write(outFile, ft.byteOrder, data)
			if err != nil {
				mylog.Log("task", "file082", "binary.Write", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 83, "dat": "2", "ans": "2" }
func (ft *fileTask) file83() bool {
	var i = ft.scanner.NextInt()
	var j = ft.scanner.NextInt()
	ft.scanner.Skip()
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file083", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var matrixOrder = int(info.Size() / 8)
	for indeX := 1; matrixOrder-indeX > 0; indeX++ {
		matrixOrder -= indeX
	}
	var element float64 = -1
	if i > 0 && i <= matrixOrder && j > 0 && j <= matrixOrder {
		if i > j {
			element = 0
		} else {
		LOOP:
			for row := 0; row < i; row++ {
				for col := 0; col < matrixOrder; col++ {
					if row > col {
						continue
					}
					err = binary.Read(file, ft.byteOrder, &element)
					if err != nil {
						mylog.Log("task", "file083", "binary.Read", err)
						return false
					}
					if row == i-1 && col == j-1 {
						break LOOP
					}
				}
			}
		}
	}
	if !ft.checker.CompareInt(matrixOrder) {
		return false
	}
	return ft.checker.CompareFloat64(2, element)
}

// { "no": 84, "dat": "2", "ans": "2" }
func (ft *fileTask) file84() bool {
	var i = ft.scanner.NextInt()
	var j = ft.scanner.NextInt()
	ft.scanner.Skip()
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file084", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var matrixOrder = int(info.Size() / 8)
	for indeX := 1; matrixOrder-indeX > 0; indeX++ {
		matrixOrder -= indeX
	}
	var element float64 = -1
	if i > 0 && i <= matrixOrder && j > 0 && j <= matrixOrder {
		if i < j {
			element = 0
		} else {
		LOOP:
			for row := 0; row < i; row++ {
				for col := 0; col < matrixOrder; col++ {
					if row < col {
						continue
					}
					err = binary.Read(file, ft.byteOrder, &element)
					if err != nil {
						mylog.Log("task", "file084", "binary.Read", err)
						return false
					}
					if row == i-1 && col == j-1 {
						break LOOP
					}
				}
			}
		}
	}
	if !ft.checker.CompareInt(matrixOrder) {
		return false
	}
	return ft.checker.CompareFloat64(2, element)
}

// { "no": 85, "dat": "2", "ans": "2" }
func (ft *fileTask) file85() bool {
	var i = ft.scanner.NextInt()
	var j = ft.scanner.NextInt()
	ft.scanner.Skip()
	var fileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file085", "os.OpenFile", err)
		return false
	}
	defer file.Close()
	var info, _ = file.Stat()
	var elementsCount = int(info.Size() / 8)
	var matrixOrder = 0
	for indeX := 1; elementsCount > 0; indeX++ {
		if indeX < 3 {
			elementsCount -= 2
		} else {
			elementsCount -= 3
		}
		matrixOrder++
	}
	var element float64 = -1
	if i > 0 && i <= matrixOrder && j > 0 && j <= matrixOrder {
		if math.Abs(float64(i-j)) > 1 {
			element = 0
		} else {
		LOOP:
			for row := 0; row < i; row++ {
				for col := 0; col < matrixOrder; col++ {
					if math.Abs(float64(row-col)) > 1 {
						continue
					}
					err = binary.Read(file, ft.byteOrder, &element)
					if err != nil {
						mylog.Log("task", "file085", "binary.Read", err)
						return false
					}
					if row == i-1 && col == j-1 {
						break LOOP
					}
				}
			}
		}
	}
	if !ft.checker.CompareInt(matrixOrder) {
		return false
	}
	return ft.checker.CompareFloat64(2, element)
}

// { "no": 86, "dat": "2", "ans": "2" }
func (ft *fileTask) file86() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file086", "os.OpenFile", err)
		return false
	}
	defer inFile.Close()
	var info, _ = inFile.Stat()
	var matrixOrder = int(info.Size() / 8)
	for indeX := 1; matrixOrder-indeX > 0; indeX++ {
		matrixOrder -= indeX
	}
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file086", "os.OpenFile", err)
		return false
	}
	defer outFile.Close()
	var element float64
	for row := 0; row < matrixOrder; row++ {
		for col := 0; col < matrixOrder; col++ {
			element = 0
			if row <= col {
				err = binary.Read(inFile, ft.byteOrder, &element)
				if err != nil {
					mylog.Log("task", "file086", "binary.Read(inFile)", err)
					return false
				}
			}
			err = binary.Write(outFile, ft.byteOrder, element)
			if err != nil {
				mylog.Log("task", "file086", "binary.Write(outFile)", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 87, "dat": "2", "ans": "2" }
func (ft *fileTask) file87() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file087", "os.OpenFile", err)
		return false
	}
	defer inFile.Close()
	var info, _ = inFile.Stat()
	var matrixOrder = int(info.Size() / 8)
	for indeX := 1; matrixOrder-indeX > 0; indeX++ {
		matrixOrder -= indeX
	}
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file087", "os.OpenFile", err)
		return false
	}
	defer outFile.Close()
	var element float64
	for row := 0; row < matrixOrder; row++ {
		for col := 0; col < matrixOrder; col++ {
			element = 0
			if row >= col {
				err = binary.Read(inFile, ft.byteOrder, &element)
				if err != nil {
					mylog.Log("task", "file087", "binary.Read(inFile)", err)
					return false
				}
			}
			err = binary.Write(outFile, ft.byteOrder, element)
			if err != nil {
				mylog.Log("task", "file087", "binary.Write(outFile)", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 88, "dat": "2", "ans": "2" }
func (ft *fileTask) file88() bool {
	var inFileName = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var outFileName = ft.scanner.NextLine()
	var inFile, outFile *os.File
	var err error
	inFile, err = os.OpenFile(inFileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file088", "os.OpenFile", err)
		return false
	}
	defer inFile.Close()
	var info, _ = inFile.Stat()
	var elementsCount = int(info.Size() / 8)
	var matrixOrder = 0
	for indeX := 1; elementsCount > 0; indeX++ {
		if indeX < 3 {
			elementsCount -= 2
		} else {
			elementsCount -= 3
		}
		matrixOrder++
	}
	outFile, err = os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file088", "os.OpenFile", err)
		return false
	}
	defer outFile.Close()
	var element float64
	for row := 0; row < matrixOrder; row++ {
		for col := 0; col < matrixOrder; col++ {
			element = 0
			if math.Abs(float64(row-col)) <= 1 {
				err = binary.Read(inFile, ft.byteOrder, &element)
				if err != nil {
					mylog.Log("task", "file088", "binary.Read(inFile)", err)
					return false
				}
			}
			err = binary.Write(outFile, ft.byteOrder, element)
			if err != nil {
				mylog.Log("task", "file088", "binary.Write(outFile)", err)
				return false
			}
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, outFileName)
}

// { "no": 89, "dat": "2", "ans": "2" }
func (ft *fileTask) file89() bool {
	var sA = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var sB = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var sC = ft.scanner.NextLine()
	var fileA, fileB, fileC *os.File
	var err error
	fileA, err = os.OpenFile(sA, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file089", "os.OpenFile(sA)", err)
		return false
	}
	defer fileA.Close()
	var info, _ = fileA.Stat()
	var matrixAOrder = info.Size() / 8
	var indeX int64
	for indeX = 1; matrixAOrder-indeX > 0; indeX++ {
		matrixAOrder -= indeX
	}
	fileB, err = os.OpenFile(sB, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file089", "os.OpenFile(sB)", err)
		return false
	}
	defer fileB.Close()
	info, _ = fileB.Stat()
	var matrixBOrder = info.Size() / 8
	for indeX = 1; matrixBOrder-indeX > 0; indeX++ {
		matrixBOrder -= indeX
	}
	fileC, err = os.OpenFile(sC, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file089", "os.OpenFile(sC)", err)
		return false
	}
	defer fileC.Close()
	if matrixAOrder == matrixBOrder {
		var col, position, row, skipBElementsSum, skipBElements, skipAElementsSum, skipAElements int64
		var dataA, dataB, dataC float64
		skipAElementsSum, skipAElements = 0, matrixAOrder
		for row = 0; row < matrixAOrder; row++ {
			for col = 0; col < matrixAOrder; col++ {
				dataC = 0
				skipBElementsSum, skipBElements = 0, matrixAOrder
				for position = 0; position < matrixAOrder; position++ {
					dataA = 0
					if row <= position {
						fileA.Seek((skipAElementsSum+(position-row))*8, 0)
						err = binary.Read(fileA, ft.byteOrder, &dataA)
						if err != nil {
							mylog.Log("task", "file089", "binary.Read(fileA)", err)
							return false
						}
					}
					dataB = 0
					if position <= col {
						fileB.Seek((skipBElementsSum+(col-position))*8, 0)
						err = binary.Read(fileB, ft.byteOrder, &dataB)
						if err != nil {
							mylog.Log("task", "file089", "binary.Read(fileB)", err)
							return false
						}
					}
					dataC += dataA * dataB
					skipBElementsSum += skipBElements
					skipBElements--
				}
				if row <= col {
					err = binary.Write(fileC, ft.byteOrder, dataC)
					if err != nil {
						mylog.Log("task", "file089", "binary.Write(fileC)", err)
						return false
					}
				}
			}
			skipAElementsSum += skipAElements
			skipAElements--
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, sC)
}

// { "no": 90, "dat": "2", "ans": "2" }
func (ft *fileTask) file90() bool {
	var sA = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var sB = ft.scanner.NextBinaryFileName(ft.byteOrder)
	var sC = ft.scanner.NextLine()
	var fileA, fileB, fileC *os.File
	var err error
	fileA, err = os.OpenFile(sA, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file090", "os.OpenFile(sA)", err)
		return false
	}
	defer fileA.Close()
	var info, _ = fileA.Stat()
	var matrixAOrder = info.Size() / 8
	var indeX int64
	for indeX = 1; matrixAOrder-indeX > 0; indeX++ {
		matrixAOrder -= indeX
	}
	fileB, err = os.OpenFile(sB, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("task", "file090", "os.OpenFile(sB)", err)
		return false
	}
	defer fileB.Close()
	info, _ = fileB.Stat()
	var matrixBOrder = info.Size() / 8
	for indeX = 1; matrixBOrder-indeX > 0; indeX++ {
		matrixBOrder -= indeX
	}
	fileC, err = os.OpenFile(sC, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("task", "file090", "os.OpenFile(sC)", err)
		return false
	}
	defer fileC.Close()
	if matrixAOrder == matrixBOrder {
		var col, position, row, skipBElementsSum, skipBElements, skipAElementsSum, skipAElements int64
		var dataA, dataB, dataC float64
		skipAElementsSum, skipAElements = 0, 1
		for row = 0; row < matrixAOrder; row++ {
			for col = 0; col < matrixAOrder; col++ {
				dataC = 0
				skipBElementsSum, skipBElements = 0, 1
				for position = 0; position < matrixAOrder; position++ {
					dataA = 0
					if row >= position {
						fileA.Seek((skipAElementsSum+position)*8, 0)
						err = binary.Read(fileA, ft.byteOrder, &dataA)
						if err != nil {
							mylog.Log("task", "file090", "binary.Read(fileA)", err)
							return false
						}
					}
					dataB = 0
					if position >= col {
						fileB.Seek((skipBElementsSum+col)*8, 0)
						err = binary.Read(fileB, ft.byteOrder, &dataB)
						if err != nil {
							mylog.Log("task", "file090", "binary.Read(fileB)", err)
							return false
						}
					}
					dataC += dataA * dataB
					skipBElementsSum += skipBElements
					skipBElements++
				}
				if row >= col {
					err = binary.Write(fileC, ft.byteOrder, dataC)
					if err != nil {
						mylog.Log("task", "file090", "binary.Write(fileC)", err)
						return false
					}
				}
			}
			skipAElementsSum += skipAElements
			skipAElements++
		}
	}
	return ft.checker.CompareBinaryFile(ft.byteOrder, sC)
}
