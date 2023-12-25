package mylog

import (
	"fmt"
)

func Log(packageName string, functionName string, positionLabel string, v interface{}) {
	if v != nil {
		fmt.Printf("%s %s %s %+v\n", packageName, functionName, positionLabel, v)
	}
}
