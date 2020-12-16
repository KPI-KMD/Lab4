package main

import (
	"fmt"
	"github.com/KPI-KMD/Lab4/engine"
)

type printCommand struct {
	output string
}

func (print *printCommand) Execute(loop engine.Handler) {
	fmt.Println(print.output)
}
  
type catCommand struct {
	arg1, arg2 string
}
  
func (cat *catCommand) Execute(loop engine.Handler) {
	res := cat.arg1 + cat.arg2
	loop.Post(&printCommand {output: res})
}