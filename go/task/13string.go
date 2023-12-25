package task

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/checker"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/scanner"
)

type stringTask struct {
	dataPath string
	name     string
	count    int
	scanner  scanner.Scanner
	checker  checker.Checker
}

func NewStringTask(pathPrefix string) Task {
	return &stringTask{
		dataPath: pathPrefix + "/13string/String%03d/%03d/%03d",
		name:     "String",
		count:    70,
	}
}

func (st *stringTask) GetCount() int { return st.count }

func (st *stringTask) GetName() string { return st.name }

func (st *stringTask) SetData(taskNo int, testNo int) error {
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

func (st *stringTask) ScannerEOF() bool { return st.scanner.EOF() }

func (st *stringTask) CheckerEOF() bool { return st.checker.EOF() }

func (st *stringTask) CleanData() {
	st.scanner.RemoveFiles()
	st.checker.RemoveFiles()
}

func (st *stringTask) RemoveScannerFiles() { st.scanner.RemoveFiles() }

func (st *stringTask) RemoveCheckFiles() { st.checker.RemoveFiles() }

func (st *stringTask) GetFn(taskNo int) func() bool {
	switch taskNo {
	case 1:
		return st.string1
	case 2:
		return st.string2
	case 3:
		return st.string3
	case 4:
		return st.string4
	case 5:
		return st.string5
	case 6:
		return st.string6
	case 7:
		return st.string7
	case 8:
		return st.string8
	case 9:
		return st.string9
	case 10:
		return st.string10
	case 11:
		return st.string11
	case 12:
		return st.string12
	case 13:
		return st.string13
	case 14:
		return st.string14
	case 15:
		return st.string15
	case 16:
		return st.string16
	case 17:
		return st.string17
	case 18:
		return st.string18
	case 19:
		return st.string19
	case 20:
		return st.string20
	case 21:
		return st.string21
	case 22:
		return st.string22
	case 23:
		return st.string23
	case 24:
		return st.string24
	case 25:
		return st.string25
	case 26:
		return st.string26
	case 27:
		return st.string27
	case 28:
		return st.string28
	case 29:
		return st.string29
	case 30:
		return st.string30
	case 31:
		return st.string31
	case 32:
		return st.string32
	case 33:
		return st.string33
	case 34:
		return st.string34
	case 35:
		return st.string35
	case 36:
		return st.string36
	case 37:
		return st.string37
	case 38:
		return st.string38
	case 39:
		return st.string39
	case 40:
		return st.string40
	case 41:
		return st.string41
	case 42:
		return st.string42
	case 43:
		return st.string43
	case 44:
		return st.string44
	case 45:
		return st.string45
	case 46:
		return st.string46
	case 47:
		return st.string47
	case 48:
		return st.string48
	case 49:
		return st.string49
	case 50:
		return st.string50
	case 51:
		return st.string51
	case 52:
		return st.string52
	case 53:
		return st.string53
	case 54:
		return st.string54
	case 55:
		return st.string55
	case 56:
		return st.string56
	case 57:
		return st.string57
	case 58:
		return st.string58
	case 59:
		return st.string59
	case 60:
		return st.string60
	case 61:
		return st.string61
	case 62:
		return st.string62
	case 63:
		return st.string63
	case 64:
		return st.string64
	case 65:
		return st.string65
	case 66:
		return st.string66
	case 67:
		return st.string67
	case 68:
		return st.string68
	case 69:
		return st.string69
	case 70:
		return st.string70
	default:
		return nil
	}
}

func (st *stringTask) string1() bool {
	var c = st.scanner.NextByte()
	var n = int(c)
	return st.checker.CompareInt(n)
}

func (st *stringTask) string2() bool {
	var n = st.scanner.NextInt()
	var c = byte(n)
	return st.checker.CompareByte(c)
}

func (st *stringTask) string3() bool {
	var c = st.scanner.NextByte()
	var prevC = c - 1
	var nextC = c + 1
	return st.checker.CompareByte(prevC, nextC)
}

func (st *stringTask) string4() bool {
	var n = st.scanner.NextInt()
	var c byte = 'A'
	for index := 0; index < n; index++ {
		if !st.checker.CompareByte(c) {
			return false
		}
		c++
	}
	return true
}

func (st *stringTask) string5() bool {
	var n = st.scanner.NextInt()
	var c byte = 'z'
	for index := 0; index < n; index++ {
		if !st.checker.CompareByte(c) {
			return false
		}
		c--
	}
	return true
}

func (st *stringTask) string6() bool {
	var c, _ = st.scanner.NextRune()
	var result string
	if unicode.IsDigit(c) {
		result = "digit"
	} else if unicode.IsUpper(c) {
		result = "capital"
	} else if unicode.IsLower(c) {
		result = "small"
	}
	return st.checker.CompareWord(result)
}

func (st *stringTask) string7() bool {
	var str = st.scanner.NextLine()
	var firstCode = int(str[0])
	var lastCode = int(str[len(str)-1])
	return st.checker.CompareInt(firstCode, lastCode)
}

func (st *stringTask) string8() bool {
	var n = st.scanner.NextInt()
	var c = st.scanner.NextByte()
	var str = ""
	for index := 0; index < n; index++ {
		str += string(c)
	}
	return st.checker.CompareLine(str)
}

func (st *stringTask) string9() bool {
	var n = st.scanner.NextInt()
	var c1 = st.scanner.NextByte()
	var c2 = st.scanner.NextByte()
	var str = ""
	for index := 0; index < n; index++ {
		str += string(c1)
		index++
		if index >= n {
			break
		}
		str += string(c2)
	}
	return st.checker.CompareLine(str)
}

func (st *stringTask) string10() bool {
	var str = st.scanner.NextLine()
	var reversedStr = ""
	for _, s := range str {
		reversedStr = string(s) + reversedStr
	}
	str = reversedStr
	return st.checker.CompareLine(str)
}

func (st *stringTask) string11() bool {
	var str = st.scanner.NextLine()
	var tempStr = ""
	for index, s := range str {
		if index > 0 {
			tempStr += " "
		}
		tempStr += string(s)
	}
	str = tempStr
	return st.checker.CompareLine(str)
}

func (st *stringTask) string12() bool {
	var str = st.scanner.NextLine()
	var n = st.scanner.NextInt()
	var separatedStr = ""
	for index := 0; index < n; index++ {
		separatedStr += "*"
	}
	var tempStr = ""
	for index, s := range str {
		if index > 0 {
			tempStr += separatedStr
		}
		tempStr += string(s)
	}
	str = tempStr
	return st.checker.CompareLine(str)
}

func (st *stringTask) string13() bool {
	var str = st.scanner.NextLine()
	var digitsCount = 0
	for _, s := range str {
		if unicode.IsDigit(s) {
			digitsCount++
		}
	}
	return st.checker.CompareInt(digitsCount)
}

func (st *stringTask) string14() bool {
	var str = st.scanner.NextLine()
	var latinCapitalsCount = 0
	for _, s := range str {
		if unicode.Is(unicode.Latin, s) && unicode.IsUpper(s) {
			latinCapitalsCount++
		}
	}
	return st.checker.CompareInt(latinCapitalsCount)
}

func (st *stringTask) string15() bool {
	var str = st.scanner.NextLine()
	var latinLettersCount = 0
	for _, s := range str {
		if unicode.Is(unicode.Latin, s) {
			latinLettersCount++
		}
	}
	return st.checker.CompareInt(latinLettersCount)
}

func (st *stringTask) string16() bool {
	var str = st.scanner.NextLine()
	str = strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Latin, r) && unicode.IsUpper(r) {
			return unicode.ToLower(r)
		}
		return r
	}, str)
	return st.checker.CompareLine(str)
}

func (st *stringTask) string17() bool {
	var str = st.scanner.NextLine()
	str = strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Latin, r) && unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		return r
	}, str)
	return st.checker.CompareLine(str)
}

func (st *stringTask) string18() bool {
	var str = st.scanner.NextLine()
	str = strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Latin, r) {
			switch {
			case unicode.IsLower(r):
				return unicode.ToUpper(r)
			case unicode.IsUpper(r):
				return unicode.ToLower(r)
			}
		}
		return r
	}, str)
	return st.checker.CompareLine(str)
}

func (st *stringTask) string19() bool {
	var str = st.scanner.NextLine()
	var result = 0
	if _, err := strconv.Atoi(str); err == nil {
		result = 1
	} else if _, err := strconv.ParseFloat(str, 64); err == nil {
		result = 2
	}
	return st.checker.CompareInt(result)
}

func (st *stringTask) string20() bool {
	var n = st.scanner.NextInt()
	var digits = make([]byte, 0)
	for n > 0 {
		digits = append(digits, strconv.Itoa(n % 10)[0])
		n /= 10
	}
	for index := len(digits) - 1; index >= 0; index-- {
		if !st.checker.CompareByte(digits[index]) {
			return false
		}
	}
	return true
}

func (st *stringTask) string21() bool {
	var n = st.scanner.NextInt()
	for n > 0 {
		if !st.checker.CompareByte(strconv.Itoa(n % 10)[0]) {
			return false
		}
		n /= 10
	}
	return true
}

func (st *stringTask) string22() bool {
	var str = st.scanner.NextWord()
	var sum = 0
	for _, s := range str {
		digit, err := strconv.Atoi(string(s))
		if err != nil {
			return false
		}
		sum += digit
	}
	return st.checker.CompareInt(sum)
}

func (st *stringTask) string23() bool {
	var str = st.scanner.NextWord()
	var result = 0
	var length = len(str)
	for index := 0; index < length; index += 2 {
		digit, err := strconv.Atoi(string(str[index]))
		if err != nil {
			return false
		}
		if index == 0 {
			result = digit
		} else {
			switch str[index-1] {
			case '+':
				result += digit
			case '-':
				result -= digit
			}
		}
	}
	return st.checker.CompareInt(result)
}

func (st *stringTask) string24() bool {
	var str = st.scanner.NextWord()
	var value, _ = strconv.ParseInt(str, 2, 64)
	str = strconv.FormatInt(value, 10)
	return st.checker.CompareWord(str)
}

func (st *stringTask) string25() bool {
	var str = st.scanner.NextWord()
	var value, _ = strconv.ParseInt(str, 10, 64)
	str = strconv.FormatInt(value, 2)
	return st.checker.CompareWord(str)
}

func (st *stringTask) string26() bool {
	var n = st.scanner.NextInt()
	var s = st.scanner.NextWord()
	var length = len(s)
	if length > n {
		s = s[length-n:]
	} else if length < n {
		for index := n - length; index > 0; index-- {
			s = "." + s
		}
	}
	return st.checker.CompareLine(s)
}

func (st *stringTask) string27() bool {
	var n1 = st.scanner.NextInt()
	var n2 = st.scanner.NextInt()
	var s1 = st.scanner.NextWord()
	var s2 = st.scanner.NextWord()
	var str = ""
	if len(s1) > n1 {
		str += s1[:n1]
	} else {
		str += s1
	}
	if len(s2) > n2 {
		str += s2[len(s2)-n2:]
	} else {
		str += s2
	}
	return st.checker.CompareLine(str)
}

func (st *stringTask) string28() bool {
	var c = st.scanner.NextByte()
	var str = st.scanner.NextWord()
	str = strings.ReplaceAll(str, string(c), string(c)+string(c))
	return st.checker.CompareLine(str)
}

func (st *stringTask) string29() bool {
	var c = st.scanner.NextByte()
	var s = st.scanner.NextWord()
	var s0 = st.scanner.NextWord()
	s = strings.ReplaceAll(s, string(c), s0+string(c))
	return st.checker.CompareLine(s)
}

func (st *stringTask) string30() bool {
	var c = st.scanner.NextByte()
	var s = st.scanner.NextWord()
	var s0 = st.scanner.NextWord()
	s = strings.ReplaceAll(s, string(c), string(c)+s0)
	return st.checker.CompareLine(s)
}

func (st *stringTask) string31() bool {
	var s = st.scanner.NextLine()
	var s0 = st.scanner.NextLine()
	var isSubstring = strings.Contains(s, s0)
	return st.checker.CompareBool(isSubstring)
}

func (st *stringTask) string32() bool {
	var s = st.scanner.NextLine()
	var s0 = st.scanner.NextLine()
	var count = strings.Count(s, s0)
	return st.checker.CompareInt(count)
}

func (st *stringTask) string33() bool {
	var s = st.scanner.NextLine()
	var s0 = st.scanner.NextLine()
	s = strings.Replace(s, s0, "", 1)
	return st.checker.CompareLine(s)
}

func (st *stringTask) string34() bool {
	var s = st.scanner.NextLine()
	var s0 = st.scanner.NextLine()
	var lastIndex = strings.LastIndex(s, s0)
	if lastIndex >= 0 {
		s = s[:lastIndex] + strings.Replace(s[lastIndex:], s0, "", 1)
	}
	return st.checker.CompareLine(s)
}

func (st *stringTask) string35() bool {
	var s = st.scanner.NextLine()
	var s0 = st.scanner.NextLine()
	s = strings.ReplaceAll(s, s0, "")
	return st.checker.CompareLine(s)
}

func (st *stringTask) string36() bool {
	var s = st.scanner.NextLine()
	var s1 = st.scanner.NextLine()
	var s2 = st.scanner.NextLine()
	s = strings.Replace(s, s1, s2, 1)
	return st.checker.CompareLine(s)
}

func (st *stringTask) string37() bool {
	var s = st.scanner.NextLine()
	var s1 = st.scanner.NextLine()
	var s2 = st.scanner.NextLine()
	var lastIndex = strings.LastIndex(s, s1)
	if lastIndex >= 0 {
		s = s[:lastIndex] + strings.Replace(s[lastIndex:], s1, s2, 1)
	}
	return st.checker.CompareLine(s)
}

func (st *stringTask) string38() bool {
	var s = st.scanner.NextLine()
	var s1 = st.scanner.NextLine()
	var s2 = st.scanner.NextLine()
	s = strings.ReplaceAll(s, s1, s2)
	return st.checker.CompareLine(s)
}

func (st *stringTask) string39() bool {
	var str = st.scanner.NextLine()
	var firstBlankIndex = strings.IndexByte(str, ' ')
	var secondBlankIndex = strings.IndexByte(str[firstBlankIndex+1:], ' ') + firstBlankIndex
	if firstBlankIndex >= 0 && firstBlankIndex < secondBlankIndex {
		str = str[firstBlankIndex+1 : secondBlankIndex+1]
	} else {
		str = ""
	}
	return st.checker.CompareLine(str)
}

func (st *stringTask) string40() bool {
	var str = st.scanner.NextLine()
	var firstBlankIndex = strings.Index(str, " ")
	var lastBlankIndex = strings.LastIndex(str, " ")
	if firstBlankIndex >= 0 && firstBlankIndex < lastBlankIndex {
		str = str[firstBlankIndex+1 : lastBlankIndex]
	} else {
		str = ""
	}
	return st.checker.CompareLine(str)
}

func (st *stringTask) string41() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	var wordsCount = len(words)
	return st.checker.CompareInt(wordsCount)
}

func (st *stringTask) string42() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	var count = 0
	for _, word := range words {
		if word[0] == word[len(word)-1] {
			count++
		}
	}
	return st.checker.CompareInt(count)
}

func (st *stringTask) string43() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	var count = 0
	for _, word := range words {
		if strings.ContainsRune(word, 'E') {
			count++
		}
	}
	return st.checker.CompareInt(count)
}

func (st *stringTask) string44() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	var count = 0
	for _, word := range words {
		if strings.Count(word, "E") == 3 {
			count++
		}
	}
	return st.checker.CompareInt(count)
}

func (st *stringTask) string45() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	var length, minLength = -1, -1
	for _, word := range words {
		length = len(word)
		if minLength < 0 || length < minLength {
			minLength = length
		}
	}
	return st.checker.CompareInt(minLength)
}

func (st *stringTask) string46() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	var length, maxLength = -1, -1
	for _, word := range words {
		length = len(word)
		if maxLength < 0 || length > maxLength {
			maxLength = length
		}
	}
	return st.checker.CompareInt(maxLength)
}

func (st *stringTask) string47() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	str = strings.Join(words, ".")
	return st.checker.CompareWord(str)
}

func (st *stringTask) string48() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	var index = 0
	for _, word := range words {
		index += strings.Index(str[index:], word)
		word = string(word[0]) + strings.ReplaceAll(word[1:], string(word[0]), ".")
		str = str[:index] + word + str[index+len(word):]
		index += len(word)
	}
	return st.checker.CompareLine(str)
}

func (st *stringTask) string49() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	var index = 0
	var length int
	for _, word := range words {
		index += strings.Index(str[index:], word)
		length = len(word)
		word = strings.ReplaceAll(word[:length-1], string(word[length-1]), ".") + string(word[length-1])
		str = str[:index] + word + str[index+len(word):]
		index += len(word)
	}
	return st.checker.CompareLine(str)
}

func (st *stringTask) string50() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	var length = len(words)
	for index := length/2 - 1; index >= 0; index-- {
		words[index], words[length-index-1] = words[length-index-1], words[index]
	}
	str = strings.Join(words, " ")
	return st.checker.CompareLine(str)
}

func (st *stringTask) string51() bool {
	var str = st.scanner.NextLine()
	var words = sort.StringSlice(strings.Fields(str))
	words.Sort()
	str = strings.Join(words, " ")
	return st.checker.CompareLine(str)
}

func (st *stringTask) string52() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	var index = 0
	var length int
	for _, word := range words {
		index += strings.Index(str[index:], word)
		length = len(word)
		if length > 0 && unicode.IsLetter(rune(word[0])) {
			word = strings.ToUpper(string(word[0])) + word[1:]
		}
		str = str[:index] + word + str[index+len(word):]
		index += len(word)
	}
	return st.checker.CompareLine(str)
}

func (st *stringTask) string53() bool {
	var str = st.scanner.NextLine()
	var count = 0
	for _, r := range str {
		if unicode.IsPunct(r) {
			count++
		}
	}
	return st.checker.CompareInt(count)
}

func (st *stringTask) string54() bool {
	var str = st.scanner.NextLine()
	var count = 0
	for _, r := range str {
		if strings.ContainsRune("aieouAIEOU", r) {
			count++
		}
	}
	return st.checker.CompareInt(count)
}

func (st *stringTask) string55() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	var length, maxLength = -1, -1
	var theWord = ""
	for _, word := range words {
		word = strings.Map(func(r rune) rune {
			if unicode.IsPunct(r) {
				return -1
			}
			return r
		}, word)
		length = len(word)
		if maxLength < 0 || length > maxLength {
			maxLength = length
			theWord = word
		}
	}
	return st.checker.CompareWord(theWord)
}

func (st *stringTask) string56() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	var length, minLength = -1, -1
	var theWord = ""
	for _, word := range words {
		word = strings.Map(func(r rune) rune {
			if unicode.IsPunct(r) {
				return -1
			}
			return r
		}, word)
		length = len(word)
		if minLength < 0 || length <= minLength {
			minLength = length
			theWord = word
		}
	}
	return st.checker.CompareWord(theWord)
}

func (st *stringTask) string57() bool {
	var str = st.scanner.NextLine()
	var words = strings.Fields(str)
	str = strings.Join(words, " ")
	return st.checker.CompareLine(str)
}

func (st *stringTask) string58() bool {
	var str = st.scanner.NextLine()
	var filename = filepath.Base(str)
	filename = strings.Replace(filename, filepath.Ext(filename), "", 1)
	return st.checker.CompareLine(filename)
}

func (st *stringTask) string59() bool {
	var str = st.scanner.NextLine()
	var extension = filepath.Ext(str)
	extension = strings.Replace(extension, ".", "", 1)
	return st.checker.CompareWord(extension)
}

func (st *stringTask) string60() bool {
	var str = st.scanner.NextLine()
	var firstSlashIndex = strings.Index(str, "/")
	var secondSlashIndex = strings.Index(str[firstSlashIndex+1:], "/")
	var home = "/"
	if firstSlashIndex >= 0 && firstSlashIndex < secondSlashIndex {
		home = str[firstSlashIndex+1 : secondSlashIndex+1]
	}
	return st.checker.CompareLine(home)
}

func (st *stringTask) string61() bool {
	var str = st.scanner.NextLine()
	var dir = filepath.Dir(str)
	if len(dir) > 0 && dir != "/" {
		var lastSlashIndex = strings.LastIndex(dir, "/")
		if lastSlashIndex >= 0 {
			dir = dir[lastSlashIndex+1:]
		}
	}
	return st.checker.CompareLine(dir)
}

func (st *stringTask) string62() bool {
	var str = st.scanner.NextLine()
	str = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			if r == 'Z' {
				return 'A'
			} else if r == 'z' {
				return 'a'
			}
			return r + 1
		}
		return r
	}, str)
	return st.checker.CompareLine(str)
}

func (st *stringTask) string63() bool {
	var str = st.scanner.NextLine()
	var k = st.scanner.NextInt()
	str = strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Latin, r) {
			var theRune = rune(int(r) + k)
			if unicode.IsLower(r) && theRune > 'z' {
				theRune = 'a' + theRune - 'z' - 1
			} else if unicode.IsUpper(r) && theRune > 'Z' {
				theRune = 'A' + theRune - 'Z' - 1
			}
			return theRune
		}
		return r
	}, str)
	return st.checker.CompareLine(str)
}

func (st *stringTask) string64() bool {
	var str = st.scanner.NextLine()
	var k = st.scanner.NextInt()
	str = strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Latin, r) {
			var theRune = rune(int(r) - k)
			if unicode.IsLower(r) && theRune < 'a' {
				theRune = 'z' - ('a' - theRune - 1)
			} else if unicode.IsUpper(r) && theRune < 'A' {
				theRune = 'Z' - ('A' - theRune - 1)
			}
			return theRune
		}
		return r
	}, str)
	return st.checker.CompareLine(str)
}

func (st *stringTask) string65() bool {
	var str = st.scanner.NextLine()
	var c, _ = st.scanner.NextRune()
	var k int
	if runes := []rune(str); len(runes) > 0 {
		k = int(runes[0] - c)
		if k < 0 {
			if unicode.IsLower(runes[0]) {
				k = int(runes[0] - 'a' + 'z' - c + 1)
			} else if unicode.IsUpper(runes[0]) {
				k = int(runes[0] - 'A' + 'Z' - c + 1)
			}
		}
	}
	str = strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Latin, r) {
			var theRune = rune(int(r) - k)
			if unicode.IsLower(r) && theRune < 'a' {
				theRune = 'z' - ('a' - theRune - 1)
			} else if unicode.IsUpper(r) && theRune < 'A' {
				theRune = 'Z' - ('A' - theRune - 1)
			}
			return theRune
		}
		return r
	}, str)
	if !st.checker.CompareInt(k) {
		return false
	}
	st.checker.Skip()
	return st.checker.CompareLine(str)
}

func (st *stringTask) string66() bool {
	var str = st.scanner.NextLine()
	var array = []rune(str)
	var length = len(array)
	str = ""
	var index int
	for index = 1; index < length; index += 2 {
		str += string(array[index])
	}
	if length%2 == 0 {
		index -= 3
	} else {
		index--
	}
	for ; index >= 0; index -= 2 {
		str += string(array[index])
	}
	return st.checker.CompareLine(str)
}

func (st *stringTask) string67() bool {
	var str = st.scanner.NextLine()
	var array = []rune(str)
	var length = len(array)
	str = ""
	for index := 0; index < length/2; index++ {
		str += string(array[length-1-index]) + string(array[index])
	}
	if length%2 != 0 {
		str += string(array[length/2])
	}
	return st.checker.CompareLine(str)
}

func (st *stringTask) string68() bool {
	var str = st.scanner.NextLine()
	var prev, curr rune
	var errIndex = -1
	for index, r := range str {
		if unicode.IsLetter(r) {
			curr = r
			if prev != 0 && prev > curr {
				errIndex = index
				break
			}
			prev = curr
		}
	}
	errIndex++
	return st.checker.CompareInt(errIndex)
}

func (st *stringTask) string69() bool {
	var str = st.scanner.NextLine()
	var parentheses = make([]rune, 0)
	var errIndex = -1
	for index, r := range str {
		switch r {
		case '(':
			parentheses = append(parentheses, r)
		case ')':
			if len(parentheses) > 0 && parentheses[len(parentheses)-1] == '(' {
				parentheses = parentheses[:len(parentheses)-1]
			} else {
				if errIndex < 0 {
					errIndex = index
				}
				parentheses = append(parentheses, r)
			}
		}
	}
	errIndex++
	if len(parentheses) > 0 && parentheses[0] != ')' {
		closeParenthesesCount := 0
		for _, parenthese := range parentheses {
			if parenthese == ')' {
				closeParenthesesCount++
			}
		}
		if closeParenthesesCount < len(parentheses)-closeParenthesesCount {
			errIndex = -1
		}
	}
	return st.checker.CompareInt(errIndex)
}

func (st *stringTask) string70() bool {
	var str = st.scanner.NextLine()
	var parentheses = make([]rune, 0)
	var errIndex = -1
	var closed bool
	var openParenthese, theOpened, theClosed rune
	for index, r := range str {
		switch r {
		case '(', '[', '{':
			parentheses = append(parentheses, r)
		case ')':
			closed = true
			openParenthese = '('
		case ']':
			closed = true
			openParenthese = '['
		case '}':
			closed = true
			openParenthese = '{'
		}
		if closed {
			if len(parentheses) > 0 && parentheses[len(parentheses)-1] == openParenthese {
				parentheses = parentheses[:len(parentheses)-1]
			} else {
				if errIndex < 0 {
					errIndex = index
					theOpened = openParenthese
					theClosed = r
				}
				parentheses = append(parentheses, r)
			}
			closed = false
		}
	}
	errIndex++
	if len(parentheses) > 0 &&
		parentheses[0] != ')' && parentheses[0] != ']' && parentheses[0] != '}' {
		if theOpened == 0 && theClosed == 0 {
			theOpened = parentheses[0]
			switch theOpened {
			case '(':
				theClosed = ')'
			case '[':
				theClosed = ']'
			case '{':
				theClosed = '}'
			}
		}
		var openedCount, closedCount int
		for _, parenthese := range parentheses {
			switch parenthese {
			case theOpened:
				openedCount++
			case theClosed:
				closedCount++
			}
		}
		if theOpened != 0 && theClosed != 0 && closedCount < openedCount {
			errIndex = -1
		}
	}
	return st.checker.CompareInt(errIndex)
}
