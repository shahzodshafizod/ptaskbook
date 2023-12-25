package main

import (
	"fmt"

	"github.com/shahzodshafizod/ptaskbook/go/task"
)

// sudo chmod -R 0544 ./js-gov4/*
// sudo chmod -R 0544 ./kt-gov2/*

// js-gov4 - worked: [Begin-Text]; Param[1-55]
// kt-gov2 - worked: [Begin-Text,Dynamic]; Param[1-55]; Tree[1-8]

func main() {

	pathPrefix := "../data/kt-gov2"
	_ = pathPrefix

	solve(task.NewBeginTask(pathPrefix))
	// solve(task.NewIntegerTask(pathPrefix))
	// solve(task.NewBooleanTask(pathPrefix))
	// solve(task.NewIfTask(pathPrefix))
	// solve(task.NewCaseTask(pathPrefix))
	// solve(task.NewForTask(pathPrefix))
	// solve(task.NewWhileTask(pathPrefix))
	// solve(task.NewSeriesTask(pathPrefix))
	// solve(task.NewProcTask(pathPrefix))
	// solve(task.NewMinmaxTask(pathPrefix))
	// solve(task.NewArrayTask(pathPrefix))
	// solve(task.NewMatrixTask(pathPrefix))
	// solve(task.NewStringTask(pathPrefix))
	// solve(task.NewFileTask(pathPrefix, binary.LittleEndian))
	// solve(task.NewTextTask(pathPrefix, binary.LittleEndian))

	// solve(task.NewParamTask(pathPrefix, binary.LittleEndian))
	// solve(task.NewRecurTask(pathPrefix))
	// solve(task.NewDynamicTask(pathPrefix))
	// solve(task.NewTreeTask(pathPrefix))
}

func solve(task task.Task) {
	for taskNo := 1; taskNo <= task.GetCount(); taskNo++ {
		for testNo := 1; testNo <= 100; testNo++ {

			if err := task.SetData(taskNo, testNo); err != nil {
				fmt.Printf("%s%03d Test#%03d task.SetData error: %s\n", task.GetName(), taskNo, testNo, err.Error())
				return
			}

			var fn = task.GetFn(taskNo)
			if fn == nil {
				fmt.Printf("%s%03d Test#%03d task.GetFn error: function not found\n", task.GetName(), taskNo, testNo)
				return
			}
			var ok = fn()

			task.CleanData()

			if !ok {
				fmt.Printf("%s%03d Test#%03d has not been tested\n", task.GetName(), taskNo, testNo)
				return
			}

			if !task.ScannerEOF() {
				fmt.Printf("%s%03d Test#%03d input is not enough\n", task.GetName(), taskNo, testNo)
				return
			}

			if !task.CheckerEOF() {
				fmt.Printf("%s%03d Test#%03d output is not enough\n", task.GetName(), taskNo, testNo)
				return
			}
		}
		fmt.Printf("%s%03d has been tested\n", task.GetName(), taskNo)
	}
}
