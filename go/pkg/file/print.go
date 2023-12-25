package file

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/shahzodshafizod/ptaskbook/go/pkg/mylog"
)

func PrintIntFile(fileName string) {
	var file, err = os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		mylog.Log("file", "printIntFile", "os.OpenFile error:", err)
		return
	}
	var data = make([]byte, 80)
	for {
		err = binary.Read(file, binary.LittleEndian, &data)
		if err == io.EOF {
			break
		} else if err != nil {
			mylog.Log("file", "printIntFile", "binary.Read error:", err)
			return
		}
		fmt.Printf("%s\t", string(data))
	}
	file.Close()
	fmt.Println()
}
