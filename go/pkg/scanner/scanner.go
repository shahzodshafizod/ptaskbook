package scanner

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/pkg/errors"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/mylog"
	"github.com/shahzodshafizod/ptaskbook/go/pkg/node"
)

type Scanner interface {
	EOF() bool
	Skip()
	NextWord() string
	NextLine() string
	NextByte() byte
	NextRune() (rune, int)
	NextBool() bool
	NextInt() int
	NextInt32() int32
	NextInt64() int64
	NextFloat32() float32
	NextFloat64() float64
	NextBinaryFileName(endian binary.ByteOrder) string
	NextTextFileName() string
	NextStack() *node.Node
	NextQueue() (head *node.Node, tail *node.Node)
	NextList() (first, last, current *node.Node)
	NextBarrierList() (barrier, current *node.Node)
	NextTree() *node.Node
	AddToRemoveQueue(key, value interface{})
	RemoveFiles()
}

type scanner struct {
	data           []byte
	filesForRemove *sync.Map
	scanPath       string
}

// New is a constructor for making a new Scanner;
func New(fileName string) (Scanner, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "os.ReadFile")
	}
	return &scanner{
		data:           data,
		filesForRemove: &sync.Map{},
		scanPath:       filepath.Dir(fileName),
	}, nil
}

// EOF checks if it is the end of the source;
func (sc *scanner) EOF() bool {
	return len(bytes.TrimSpace(sc.data)) == 0
}

// Skip skips one symbol of the source;
func (sc *scanner) Skip() {
	if len(sc.data) > 0 {
		sc.data = (sc.data)[1:]
	}
}

// NextWord reads the next word from the source;
func (sc *scanner) NextWord() string {
	advance, token, err := bufio.ScanWords(sc.data, true)
	if err != nil || advance == 0 {
		mylog.Log("scanner", "NextWord", "", "input EOF")
		return *new(string)
	}
	var word = string(token)
	sc.data = (sc.data)[advance:]
	return word
}

// NextLine reads the next line from the source;
func (sc *scanner) NextLine() string {
	advance, token, err := bufio.ScanLines(sc.data, true)
	if err != nil {
		mylog.Log("scanner", "NextLine", "bufio.ScanLines", err)
		return *new(string)
	}
	var line = string(token)
	if advance == 0 {
		return line
	}
	sc.data = (sc.data)[advance:]
	return line
}

// NextByte reads the next byte from the source;
func (sc *scanner) NextByte() byte {
	var value byte
	for len(sc.data) > 0 {
		advance, token, err := bufio.ScanBytes(sc.data, true)
		if err != nil || advance == 0 {
			mylog.Log("scanner", "NextByte", "", "input EOF")
			return *new(byte)
		}
		value = token[0]
		sc.data = (sc.data)[advance:]
		if value != '\n' {
			break
		}
	}
	// read the splitter symbol from scanner;
	advance, _, err := bufio.ScanBytes(sc.data, true)
	if err != nil || advance == 0 {
		return value
	}
	sc.data = (sc.data)[advance:]
	return value
}

// NextRune reads the next rune from the source;
func (sc *scanner) NextRune() (rune, int) {
	var (
		value rune
		size  int
	)
	for len(sc.data) > 0 {
		advance, token, err := bufio.ScanRunes(sc.data, true)
		if err != nil || advance == 0 {
			mylog.Log("scanner", "NextRune", "", "input EOF")
			return *new(rune), 0
		}
		value, size = utf8.DecodeRune(token)
		sc.data = (sc.data)[advance:]
		if size != 0 && value != '\n' {
			break
		}
	}
	// read the splitter symbol from scanner;
	advance, _, err := bufio.ScanRunes(sc.data, true)
	if err != nil || advance == 0 {
		return value, size
	}
	sc.data = (sc.data)[advance:]
	return value, size
}

// NextBool reads the next bool from the source;
func (sc *scanner) NextBool() bool {
	if sc.EOF() {
		mylog.Log("scanner", "NextBool", "", "input EOF")
		return *new(bool)
	}
	var value, _ = strconv.ParseBool(sc.NextWord())
	return value
}

// NextInt reads the next int from the source;
func (sc *scanner) NextInt() int {
	if sc.EOF() {
		mylog.Log("scanner", "NextInt", "", "input EOF")
		return *new(int)
	}
	var value, _ = strconv.Atoi(sc.NextWord())
	return value
}

// NextInt32 reads the next int32 from the source;
func (sc *scanner) NextInt32() int32 {
	if sc.EOF() {
		mylog.Log("scanner", "NextInt32", "", "input EOF")
		return *new(int32)
	}
	var value, _ = strconv.Atoi(sc.NextWord())
	return int32(value)
}

// NextInt64 reads the next int64 from the source;
func (sc *scanner) NextInt64() int64 {
	if sc.EOF() {
		mylog.Log("scanner", "NextInt64", "", "input EOF")
		return *new(int64)
	}
	var value, _ = strconv.ParseInt(sc.NextWord(), 10, 64)
	return value
}

// NextFloat32 reads the next float32 from the source;
func (sc *scanner) NextFloat32() float32 {
	if sc.EOF() {
		mylog.Log("scanner", "NextFloat32", "", "input EOF")
		return *new(float32)
	}
	return float32(sc.NextFloat64())
}

// NextFloat64 reads the next float64 from the source;
func (sc *scanner) NextFloat64() float64 {
	if sc.EOF() {
		mylog.Log("scanner", "NextFloat64", "", "input EOF")
		return *new(float64)
	}
	var value, _ = strconv.ParseFloat(sc.NextWord(), 64)
	return value
}

// NextBinaryFileName reads the next file name from the source;
func (sc *scanner) NextBinaryFileName(endian binary.ByteOrder) string {
	var fileName = sc.NextLine()
	if endian == binary.BigEndian {
		fileName += ".big"
	}
	var fromName = sc.scanPath + "/" + fileName
	var _, err = os.Stat(fromName)
	if err != nil {
		return fileName
	}
	data, err := os.ReadFile(fromName)
	if err != nil {
		mylog.Log("scanner", "NextBinaryFileName", "os.ReadFile", err)
		return *new(string)
	}
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("scanner", "NextBinaryFileName", "os.OpenFile", err)
		return *new(string)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		mylog.Log("scanner", "NextBinaryFileName", "file.Write", err)
		return *new(string)
	}
	sc.AddToRemoveQueue(fileName, fileName)
	return fileName
}

// NextTextFileName reads the next file name from the source;
func (sc *scanner) NextTextFileName() string {
	var fileName = sc.NextLine()
	sc.AddToRemoveQueue(fileName, fileName)
	var fromName = sc.scanPath + "/" + fileName
	var _, err = os.Stat(fromName)
	if err != nil {
		return *new(string)
	}
	data, err := os.ReadFile(fromName)
	if err != nil {
		mylog.Log("scanner", "NextTextFileName", "os.ReadFile", err)
		return *new(string)
	}
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		mylog.Log("scanner", "NextTextFileName", "os.OpenFile", err)
		return *new(string)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		mylog.Log("scanner", "NextTextFileName", "file.Write", err)
		return *new(string)
	}
	return fileName
}

// NextStack reads the next stack with it's elements from the source;
func (sc *scanner) NextStack() *node.Node {
	var count = sc.NextInt()
	var top *node.Node = nil
	for index := 0; index < count; index++ {
		top = &node.Node{
			Data: sc.NextInt(),
			Next: top,
		}
	}
	return top
}

// NextQueue reads the next queue with it's elements from the source;
func (sc *scanner) NextQueue() (head *node.Node, tail *node.Node) {
	var count = sc.NextInt()
	var newNode *node.Node
	for index := 0; index < count; index++ {
		newNode = &node.Node{Data: sc.NextInt()}
		if tail != nil {
			tail.Next = newNode
		}
		tail = newNode
		if head == nil {
			head = newNode
		}
	}
	return
}

// NextList reads the next list with it's elements from the source;
func (sc *scanner) NextList() (first *node.Node, last *node.Node, current *node.Node) {
	var (
		count              = sc.NextInt()
		currentOrderNumber = sc.NextInt()
	)
	var newNode *node.Node
	for index := 0; index < count; index++ {
		newNode = &node.Node{
			Data: sc.NextInt(),
			Prev: last,
		}
		if last != nil {
			last.Next = newNode
		}
		last = newNode
		if currentOrderNumber-1 == index {
			current = newNode
		}
		if first == nil {
			first = newNode
		}
	}
	return
}

// NextBarrierList reads the next barrier list with it's elements from the source;
func (sc *scanner) NextBarrierList() (barrier *node.Node, current *node.Node) {
	first, last, current := sc.NextList()
	barrier = &node.Node{
		Next: first,
		Prev: last,
	}
	if first != nil {
		first.Prev = barrier
	} else {
		barrier.Next = barrier
	}
	if last != nil {
		last.Next = barrier
	} else {
		barrier.Prev = barrier
	}
	if current == nil {
		current = barrier
	}
	return
}

// NextTree reads the next barrier list with it's elements from the source;
func (sc *scanner) NextTree() *node.Node {
	// read level of tree;
	var level = sc.NextInt()
	sc.Skip()
	// read source tree into words matrix;
	var wordsArray = make([][]string, 0)
	for index := 0; index < level; index++ {
		var line = sc.NextLine()
		var words = strings.Fields(line)
		wordsArray = append(wordsArray, words)
	}
	// make an array for indexes of every row array;
	var colIndexes = make([]int, level)
	// parse tree;
	return nextTree(nil, wordsArray, 0, level, colIndexes)
}

func nextTree(parent *node.Node, wordsArray [][]string, rowIndex int, rowsCount int, colIndexes []int) *node.Node {
	if rowIndex >= rowsCount || colIndexes[rowIndex] >= len(wordsArray[rowIndex]) {
		return nil
	}
	var word = wordsArray[rowIndex][colIndexes[rowIndex]]
	colIndexes[rowIndex]++
	var dataStr = strings.Map(func(r rune) rune {
		if !unicode.IsDigit(r) {
			r = -1
		}
		return r
	}, word)
	var data, _ = strconv.Atoi(dataStr)
	node := &node.Node{
		Data:   data,
		Parent: parent,
	}

	if parent == nil {
		parent = node
	}

	var firstChar, lastChar rune
	for index, char := range word {
		if index == 0 {
			firstChar = char
		}
		lastChar = char
	}

	if firstChar == '┌' {
		node.Left = nextTree(node, wordsArray, rowIndex+1, rowsCount, colIndexes)
	}
	if lastChar == '┐' {
		node.Right = nextTree(node, wordsArray, rowIndex+1, rowsCount, colIndexes)
	}
	return node
}

func (sc *scanner) AddToRemoveQueue(key, value interface{}) {
	sc.filesForRemove.LoadOrStore(key, value)
}

// RemoveFiles removes all made files;
func (sc *scanner) RemoveFiles() {
	sc.filesForRemove.Range(func(key, value interface{}) bool {
		var fileName = key.(string)
		var _, err = os.Stat(fileName)
		if err != nil {
			return true
		}
		os.Remove(fileName)
		sc.filesForRemove.Delete(key)
		return true
	})
}
