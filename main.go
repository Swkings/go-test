//go:build ignore

package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gosuri/uitable"
)

func main() {
	type pod struct {
		NAME, READY, STATUS, AGE string
		RESTARTS                 int
	}

	var pods = []pod{
		{"nginx1-0 ", "1/1", "Running", "11d", 0},
		{"nginx2-0", "1/1", "Running", "11d", 0},
	}

	table := uitable.New()
	table.MaxColWidth = 100
	table.RightAlign(10)

	table.AddRow("NAME", "READY", "STATUS", "RESTARTS", "AGE")
	for _, pod := range pods {
		table.AddRow(color.RedString(pod.NAME), color.WhiteString(pod.READY), color.BlueString(pod.STATUS), color.GreenString(pod.AGE), color.YellowString("%d", pod.RESTARTS))
	}
	fmt.Println(table)
}
