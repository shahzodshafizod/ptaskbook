package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/shahzodshafizod/ptaskbook/tasks/task"
)

func main() {
	tasks := []task.Task{
		task.NewBeginTask(),
		task.NewIntegerTask(),
		task.NewBooleanTask(),
		task.NewIfTask(),
		task.NewCaseTask(),
		task.NewForTask(),
		task.NewWhileTask(),
		task.NewSeriesTask(),
		task.NewProcTask(),
		task.NewMinmaxTask(),
		task.NewArrayTask(),
		task.NewMatrixTask(),
		task.NewStringTask(),
		task.NewFileTask(),
		task.NewTextTask(),
		task.NewParamTask(),
		task.NewRecurTask(),
		task.NewDynamicTask(),
		task.NewTreeTask(),
	}

	for _, task := range tasks {
		en := task.En()
		tj := task.Tj()
		ru := task.Ru()
		taskName := task.Name()

		filename := "mds/" + strings.ToLower(taskName) + ".md"
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("os.Open", err)
			return
		}
		defer file.Close()

		count := task.Count()
		for i := 1; i < count; i++ {
			fmt.Fprintf(file, "## %s %d\n", taskName, i)

			fmt.Fprintf(file, "#### ![en](https://img.shields.io/badge/EN-blue) %s\n", en[i])

			if taskName != "Text" &&
				taskName != "Param" &&
				taskName != "Recur" &&
				taskName != "Dynamic" &&
				taskName != "Tree" {
				fmt.Fprintf(file, "#### ![tj](https://img.shields.io/badge/TJ-blue) %s\n", tj[i])
			}

			fmt.Fprintf(file, "#### ![ru](https://img.shields.io/badge/RU-blue) %s\n", ru[i])

			fmt.Fprint(file, "\n---\n\n")
		}
	}
}

/*
TASK	| TJ | EN | RU |
--------+----+----+----+
Begin	| +  | +  | +  |
Integer	| +  | +  | +  |
Boolean	| +  | +  | +  |
If		| +  | +  | +  |
Ternary	| +  | +  | +  |
Case	| +  | +  | +  |
For		| +  | +  | +  |
While	| +  | +  | +  |
Series	| +  | +  | +  |
Proc	| +  | +  | +  |
Minmax	| +  | +  | +  |
Array	| +  | +  | +  |
Matrix	| +  | +  | +  |
String	| +  | +  | +  |
File	| +  | +  | +  |
Text	|    | +  | +  |
Param	|    | +  | +  |
Recur	|    | +  | +  |
Dynamic	|    | +  | +  |
Tree	|    | +  | +  |
*/
