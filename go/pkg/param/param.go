package param

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/shahzodshafizod/ptaskbook/go/pkg/array"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/matrix"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/mylog"
)

func MinElem(a []int, n int) int {
	var minimal = 0
	for index, elem := range a {
		if index == 0 || elem < minimal {
			minimal = elem
		}
	}
	return minimal
}

func MaxNum(a []float64, n int) int {
	var num int
	for index, elem := range a {
		if index == 0 || elem > a[num] {
			num = index
		}
	}
	num++
	return num
}

func MinmaxNum(a []float64, n int, nMin, nMax *int) {
	*nMin, *nMax = 0, 0
	for index, elem := range a {
		if index > 0 && elem < a[*nMin] {
			*nMin = index
		} else if index > 0 && elem > a[*nMax] {
			*nMax = index
		}
	}
	(*nMin)++
	(*nMax)++
}

func Inv(a []float64, n int) {
	tempA := array.Float64(a)
	tempA.Reverse()
	a = tempA
}

func Smooth1(a []float64, n int) {
	var sum float64
	for index, elem := range a {
		sum += elem
		a[index] = sum / float64(index+1)
	}
}

func Smooth2(a []float64, n int) {
	if n <= 0 {
		return
	}
	var prevElem = a[0]
	for index := 1; index < n; index++ {
		a[index], prevElem = (prevElem+a[index])/2, a[index]
	}
}

func Smooth3(a []float64, n int) {
	if n <= 1 {
		return
	}
	var prev, curr = a[0], a[1]
	var sum = prev + curr
	a[0] = sum / 2
	for index := 1; index < n-1; index++ {
		curr = a[index]
		sum += a[index+1]
		a[index] = sum / 3
		sum -= prev
		prev = curr
	}
	a[n-1] = sum / 2
}

func RemoveX(a *[]int, n *int, x int) {
	for index := 0; index < *n; index++ {
		if (*a)[index] == x {
			*a = append((*a)[:index], (*a)[index+1:]...)
			(*n)--
			index--
		}
	}
}

func RemoveForInc(a *[]float64, n *int) {
	for index := 1; index < *n; index++ {
		if (*a)[index-1] > (*a)[index] {
			*a = append((*a)[:index], (*a)[index+1:]...)
			(*n)--
			index--
		}
	}
}

func DoubleX(a *[]int, n *int, x int) {
	for index := 0; index < *n; index++ {
		if (*a)[index] == x {
			*a = append((*a)[:index+1], (*a)[index:]...)
			(*n)++
			index++
		}
	}
}

func SortArray(a []float64, n int) {
	tempA := sort.Float64Slice(a)
	tempA.Sort()
	a = tempA
}

func SortIndex(a []float64, n int, indexes *[]int) {
	*indexes = make([]int, n)
	for index := 0; index < n; index++ {
		(*indexes)[index] = index
	}
	for i := 0; i < n-1; i++ {
		for index := 1; index < n-i; index++ {
			if a[(*indexes)[index-1]] > a[(*indexes)[index]] {
				(*indexes)[index-1], (*indexes)[index] = (*indexes)[index], (*indexes)[index-1]
			}
		}
	}
	for index := 0; index < n; index++ {
		(*indexes)[index]++
	}
}

func Hill(a []float64, n int) {
	for i := 0; i < n/2; i++ {
		// move to the beginning;
		for index := n - 1 - i; index > i; index-- {
			if a[index] < a[index-1] {
				a[index], a[index-1] = a[index-1], a[index]
			}
		}
		// move to the end;
		for index := i + 1; index < n-1-i; index++ {
			if a[index] < a[index+1] {
				a[index], a[index+1] = a[index+1], a[index]
			}
		}
	}
}

func Split1(a []float64, na int, b *[]float64, nb *int, c *[]float64, nc *int) {
	*nb, *nc = 0, 0
	*b = make([]float64, *nb)
	*c = make([]float64, *nc)
	for index := 0; index < na; index += 2 {
		*b = append(*b, a[index])
		(*nb)++
	}
	for index := 1; index < na; index += 2 {
		*c = append(*c, a[index])
		(*nc)++
	}
}

func Split2(a []int, na int, b *[]int, nb *int, c *[]int, nc *int) {
	*nb, *nc = 0, 0
	*b = make([]int, *nb)
	*c = make([]int, *nc)
	for index := 0; index < na; index++ {
		if a[index]%2 == 0 {
			*b = append(*b, a[index])
			(*nb)++
		} else {
			*c = append(*c, a[index])
			(*nc)++
		}
	}
}

func ArrayToMatrRow(a []float64, k, m, n int, b *[][]float64) {
	*b = make([][]float64, m)
	var i = 0
	for row := 0; row < m; row++ {
		(*b)[row] = make([]float64, n)
		for col := 0; col < n; col++ {
			if i < k {
				(*b)[row][col] = a[i]
				i++
			}
		}
	}
}

func ArrayToMatrCol(a []float64, k, m, n int, b *[][]float64) {
	*b = make([][]float64, m)
	for row := 0; row < m; row++ {
		(*b)[row] = make([]float64, n)
	}
	var i = 0
	for col := 0; col < n; col++ {
		for row := 0; row < m; row++ {
			if i < k {
				(*b)[row][col] = a[i]
				i++
			}
		}
	}
}

func Chessboard(m, n int, a *[][]int) {
	*a = make([][]int, m)
	for row := 0; row < m; row++ {
		(*a)[row] = make([]int, n)
		for col := 0; col < n; col++ {
			(*a)[row][col] = (row + col) % 2
		}
	}
}

func Norm1(a [][]float64, m, n int) float64 {
	var max float64 = 0
	var sum float64
	for col := 0; col < n; col++ {
		sum = 0
		for row := 0; row < m; row++ {
			sum += a[row][col]
		}
		if col == 0 || sum > max {
			max = sum
		}
	}
	return max
}

func Norm2(a [][]float64, m, n int) float64 {
	var max float64 = 0
	var sum float64
	for row := 0; row < m; row++ {
		sum = 0
		for col := 0; col < n; col++ {
			sum += a[row][col]
		}
		if row == 0 || sum > max {
			max = sum
		}
	}
	return max
}

func SumRow(a [][]float64, m, n, k int) float64 {
	var sum float64 = 0
	if k > 0 && k <= m {
		for col := 0; col < n; col++ {
			sum += a[k-1][col]
		}
	}
	return sum
}

func SumCol(a [][]float64, m, n, k int) float64 {
	var sum float64 = 0
	if k > 0 && k <= n {
		for row := 0; row < m; row++ {
			sum += a[row][k-1]
		}
	}
	return sum
}

func SwapRow(a [][]float64, m, n, k1, k2 int) {
	if k1 < 1 || k1 > m || k2 < 1 || k2 > m {
		return
	}
	for col := 0; col < n; col++ {
		a[k1-1][col], a[k2-1][col] = a[k2-1][col], a[k1-1][col]
	}
}

func SwapCol(a [][]float64, m, n, k1, k2 int) {
	if k1 < 1 || k1 > n || k2 < 1 || k2 > n {
		return
	}
	for row := 0; row < m; row++ {
		a[row][k1-1], a[row][k2-1] = a[row][k2-1], a[row][k1-1]
	}
}

func Transp(a [][]float64, m int) {
	for row := 0; row < m; row++ {
		for col := row; col < m; col++ {
			a[row][col], a[col][row] = a[col][row], a[row][col]
		}
	}
}

func RemoveRows(a *[][]float64, m *int, n, k1, k2 int) {
	if k1 <= 0 {
		k1 = 1
	}
	if k2 > *m {
		k2 = *m
	}
	var matrix = matrix.Float64(*a)
	for index := k1; index <= k2; index++ {
		matrix.DeleteRowAt(m, n, k1-1)
	}
	*a = matrix
}

func RemoveCols(a *[][]float64, m int, n *int, k1, k2 int) {
	if k1 <= 0 {
		k1 = 1
	}
	if k2 > *n {
		k2 = *n
	}
	var matrix = matrix.Float64(*a)
	for index := k1; index <= k2; index++ {
		matrix.DeleteColAt(m, n, k1-1)
	}
	*a = matrix
}

func RemoveRowCol(a *[][]float64, m, n *int, k, l int) {
	if k < 1 || k > *m || l < 1 || l > *n {
		return
	}
	var matrix = matrix.Float64(*a)
	matrix.DeleteRowAt(m, *n, k-1)
	matrix.DeleteColAt(*m, n, l-1)
	*a = matrix
}

func SortCols(a [][]int, m, n int) {
	var maxColIndex int
	for index := n - 1; index > 0; index-- {
		maxColIndex = index
		for col := index - 1; col >= 0; col-- {
			for row := 0; row < m; row++ {
				if a[row][col] < a[row][maxColIndex] {
					break
				} else if a[row][col] > a[row][maxColIndex] {
					maxColIndex = col
				}
			}
		}
		if maxColIndex != index {
			for row := 0; row < m; row++ {
				a[row][maxColIndex], a[row][index] = a[row][index], a[row][maxColIndex]
			}
		}
	}
}

func IsIdent(s string) int {
	if len(s) == 0 {
		return -1
	}
	for index, r := range s {
		switch {
		case unicode.IsDigit(r):
			if index == 0 {
				return -2
			}
		case unicode.Is(unicode.Latin, r), r == '_':
		default:
			return index + 1
		}
	}
	return 0
}

func FillStr(s string, n int) string {
	if len(s) == 0 {
		return ""
	}
	var result = ""
LOOP:
	for {
		for _, r := range s {
			result += string(r)
			if len(result) >= n {
				break LOOP
			}
		}
	}
	return result
}

func UpCaseLat(s *string) {
	*s = strings.Map(func(r rune) rune {
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		return r
	}, *s)
}

func LowCaseLat(s *string) {
	*s = strings.Map(func(r rune) rune {
		if unicode.IsUpper(r) {
			return unicode.ToLower(r)
		}
		return r
	}, *s)
}

func TrimLeftC(s *string, c rune) {
	*s = strings.TrimLeftFunc(*s, func(r rune) bool { return r == c })
}

func TrimRightC(s *string, c rune) {
	*s = strings.TrimRightFunc(*s, func(r rune) bool { return r == c })
}

func InvStr(s string, k, n int) string {
	if len(s) < k {
		return ""
	}
	s = s[k-1:]
	var count = 0
	var inversed = ""
	for _, r := range s {
		inversed = string(r) + inversed
		count++
		if count >= n {
			break
		}
	}
	return inversed
}

func PosSub(s0, s string, k, n int) int {
	if k > len(s) {
		return 0
	}
	if len(s) < k+n {
		n = len(s)
	}
	var subs string
	var index = 0
	for _, r := range s {
		index++
		if index < k {
			continue
		} else if index >= k+n {
			break
		}
		subs += string(r)
	}
	var position = strings.Index(subs, s0) + 1
	if position > 0 {
		position += k - 1
	}
	return position
}

func PosLast(s0, s string) int {
	return strings.LastIndex(s, s0) + 1
}

func PosK(s0, s string, k int) int {
	var position int
	var lengthSum, s0Length = 0, len(s0)
	for index := 0; index < k; index++ {
		position = strings.Index(s, s0)
		if position < 0 {
			break
		}
		position += s0Length
		s = s[position:]
		lengthSum += position
	}
	position++
	if position > 0 {
		position = lengthSum - s0Length + 1
	}
	return position
}

func WordK(s string, k int) string {
	var words = strings.FieldsFunc(s, func(r rune) bool { return unicode.IsSpace(r) })
	var word = ""
	if len(words) >= k {
		word = words[k-1]
	}
	return word
}

func SplitStr(s string, w *[]string, n *int) {
	*w = strings.FieldsFunc(s, func(r rune) bool { return unicode.IsSpace(r) })
	*n = len(*w)
}

func CompressStr(s string) string {
	var count int
	var compressed = ""
	for position, theRune := range s {
		if count > 0 {
			count--
			continue
		}
		if strings.Index(s[position:], strings.Repeat(string(theRune), 5)) == 0 {
			count = 0
			for _, r := range s[position:] {
				if r != theRune {
					break
				}
				count++
			}
			compressed += fmt.Sprintf("%s{%d}", string(theRune), count)
			count--
			continue
		}
		compressed += string(theRune)
	}
	return compressed
}

func DecompressStr(s string) string {
	var prevRune rune
	var skipCount int
	var decompressed = ""
	for position, theRune := range s {
		if skipCount > 0 {
			skipCount--
			continue
		}
		if theRune == '{' {
			countStr := ""
			skipCount = 0
			for _, r := range s[position+1:] {
				skipCount++
				if r == '}' {
					break
				}
				countStr += string(r)
			}
			count, _ := strconv.Atoi(countStr)
			decompressed += strings.Repeat(string(prevRune), count-1)
			continue
		}
		decompressed += string(theRune)
		prevRune = theRune
	}
	return decompressed
}

func DecToBin(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func DecToHex(n int) string {
	return strings.ToUpper(strconv.FormatInt(int64(n), 16))
}

func BinToDec(s string) int {
	tempN, _ := strconv.ParseInt(s, 2, 32)
	return int(tempN)
}

func HexToDec(s string) int {
	tempN, _ := strconv.ParseInt(s, 16, 32)
	return int(tempN)
}

func IntFileSize(s string) int64 {
	var info, err = os.Stat(s)
	if err != nil {
		return -1
	}
	return info.Size() / 4
}

func LineCount(s string) int {
	var file, err = os.OpenFile(s, os.O_RDONLY, 0777)
	if err != nil {
		return -1
	}
	defer file.Close()
	var scanner = bufio.NewScanner(file)
	var count = 0
	for scanner.Scan() {
		count++
	}
	return count
}

func InvIntFile(s string, order binary.ByteOrder) {
	var file, err = os.OpenFile(s, os.O_RDWR, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	var info, _ = file.Stat()
	var size = info.Size()
	var position int64
	var number1, number2 int32
	for position = 0; position < size/2; position += 4 {
		// read an element from the beginning of the file;
		file.Seek(position, 0)
		err = binary.Read(file, order, &number1)
		if err != nil {
			mylog.Log("param", "invIntFile", "binary.Read(file, &number1) error:", err)
			return
		}
		// read an element from the end of the file;
		file.Seek(-position-4, 2)
		err = binary.Read(file, order, &number2)
		if err != nil {
			mylog.Log("param", "invIntFile", "binary.Read(file, &number2) error:", err)
			return
		}
		// write to the end of the file;
		file.Seek(-4, 1)
		err = binary.Write(file, order, number1)
		if err != nil {
			mylog.Log("param", "invIntFile", "binary.Write(file, number1) error:", err)
			return
		}
		// write to the beginning of the file;
		file.Seek(position, 0)
		err = binary.Write(file, order, number2)
		if err != nil {
			mylog.Log("param", "invIntFile", "binary.Write(file, number2) error:", err)
			return
		}
	}
}

func AddLineNumbers(s string, n, k, l int) {
	var file, err = os.OpenFile(s, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("param", "addLineNumbers", "os.OpenFile(s) error:", err)
		return
	}
	defer file.Close()
	var tempFileName = "$tempFileName.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("param", "addLineNumbers", "os.OpenFile(tempFileName) error:", err)
		return
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		_, err = tempFile.WriteString(fmt.Sprintf("%*d", k, n))
		if err != nil {
			mylog.Log("param", "addLineNumbers", "tempFile.WriteString 1 error:", err.Error())
			return
		}
		if len(line) != 0 {
			_, err = tempFile.WriteString(strings.Repeat(" ", l) + line)
			if err != nil {
				mylog.Log("param", "addLineNumbers", "tempFile.WriteString 2 error:", err.Error())
				return
			}
		}
		_, err = tempFile.WriteString("\n")
		if err != nil {
			mylog.Log("param", "addLineNumbers", "tempFile.WriteString 3 error:", err.Error())
			return
		}
		n++
	}
	os.Remove(s)
	os.Rename(tempFileName, s)
}

func RemoveLineNumbers(s string) {
	var file, err = os.OpenFile(s, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("param", "removeLineNumbers", "os.OpenFile(s) error:", err)
		return
	}
	defer file.Close()
	var tempFileName = "$tempFileName.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("param", "removeLineNumbers", "os.OpenFile(tempFileName) error:", err)
		return
	}
	defer tempFile.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line, word string
	var prevLineNumber, lineNumber int
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			return
		}
		word = strings.Fields(line)[0]
		lineNumber, err = strconv.Atoi(word)
		if err != nil {
			return
		}
		if prevLineNumber != 0 && prevLineNumber+1 != lineNumber {
			return
		}
		prevLineNumber = lineNumber
		line = line[strings.Index(line, word)+len(word):]
		line = strings.TrimLeftFunc(line, func(r rune) bool { return unicode.IsSpace(r) })
		_, err = tempFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("param", "removeLineNumbers", "tempFile.WriteString error:", err)
			return
		}
	}
	os.Remove(s)
	os.Rename(tempFileName, s)
}

func SplitIntFile(s0 string, k int, s1, s2 string, order binary.ByteOrder) {
	var inFile, outFile1, outFile2 *os.File
	var err error
	inFile, err = os.OpenFile(s0, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("param", "splitIntFile", "os.OpenFile(s0) error:", err)
		return
	}
	defer inFile.Close()
	outFile1, err = os.OpenFile(s1, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("param", "splitIntFile", "os.OpenFile(s1) error:", err)
		return
	}
	defer outFile1.Close()
	outFile2, err = os.OpenFile(s2, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("param", "splitIntFile", "os.OpenFile(s2) error:", err)
		return
	}
	defer outFile2.Close()
	var number int32
	index := 0
	var outFile *os.File
	for {
		index++
		err = binary.Read(inFile, order, &number)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("param", "splitIntFile", "binary.Read(inFile) error:", err)
			return
		}
		if index <= k {
			outFile = outFile1
		} else {
			outFile = outFile2
		}
		err = binary.Write(outFile, order, number)
		if err != nil {
			mylog.Log("param", "splitIntFile", "binary.Write(outFile) error:", err)
			return
		}
	}
}

func SplitText(s0 string, k int, s1, s2 string) {
	var inFile, outFile, outFile1, outFile2 *os.File
	var err error
	inFile, err = os.OpenFile(s0, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("param", "splitText", "os.OpenFile(s0) error:", err)
		return
	}
	defer inFile.Close()
	outFile1, err = os.OpenFile(s1, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("param", "splitText", "os.OpenFile(s1) error:", err)
		return
	}
	defer outFile1.Close()
	outFile2, err = os.OpenFile(s2, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("param", "splitText", "os.OpenFile(s2) error:", err)
		return
	}
	defer outFile2.Close()
	var scanner = bufio.NewScanner(inFile)
	var line string
	index := 0
	for scanner.Scan() {
		index++
		line = scanner.Text()
		if index <= k {
			outFile = outFile1
		} else {
			outFile = outFile2
		}
		_, err = outFile.WriteString(line + "\n")
		if err != nil {
			mylog.Log("param", "splitText", "outFile.WriteString error:", err)
			return
		}
	}
}

func StringFileToText(s string, order binary.ByteOrder) {
	var binaryFile, err = os.OpenFile(s, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("param", "stringFileToText(s)", "os.OpenFile", err)
		return
	}
	defer binaryFile.Close()
	var tempFileName = "$tempFileName.tmp$"
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("param", "stringFileToText", "os.OpenFile(tempFileName)", err)
		return
	}
	defer tempFile.Close()
	for {
		var data = make([]byte, 80)
		err = binary.Read(binaryFile, order, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("param", "stringFileToText", "binary.Read", err)
			return
		}
		data = bytes.TrimRightFunc(data, func(r rune) bool { return unicode.IsSpace(r) || r == 0 })
		data = append(data, '\n')
		_, err = tempFile.Write(data)
		if err != nil {
			mylog.Log("param", "stringFileToText", "tempFile.WriteString", err)
			return
		}
	}
	os.Remove(s)
	err = os.Rename(tempFileName, s)
	if err != nil {
		mylog.Log("param", "stringFileToText", "os.Rename", err)
		return
	}
}

func TextToStringFile(s string, order binary.ByteOrder) {
	var textFile, err = os.OpenFile(s, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("param", "textToStringFile", "os.OpenFile(s)", err)
		return
	}
	defer textFile.Close()
	var tempFileName = "$tempFileName.tmp$"
	binaryFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("param", "textToStringFile", "os.OpenFile(tempFileName)", err)
		return
	}
	defer binaryFile.Close()
	var scanner = bufio.NewScanner(textFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var line = scanner.Bytes()
		line = append(line, make([]byte, 80-len(line))...)
		err = binary.Write(binaryFile, order, line)
		if err != nil {
			mylog.Log("param", "textToStringFile", "binary.Write", err)
			return
		}
	}
	os.Remove(s)
	err = os.Rename(tempFileName, s)
	if err != nil {
		mylog.Log("param", "textToStringFile", "os.Rename", err)
		return
	}
}
