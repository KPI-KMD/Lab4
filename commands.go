package main

import (
	"fmt"
)

type printCommand struct {
	output string
}

func (print *printCommand) Execute(loop Handler) {
	fmt.Println(print.output)
}
  
type catCommand struct {
	arg1, arg2 string
}
  
func (cat *catCommand) Execute(loop Handler) {
	res := cat.arg1 + cat.arg2
	loop.Post(&printCommand {output: res})
}