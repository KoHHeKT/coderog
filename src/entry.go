package main

import (
	"entry/cls"
	"fmt"
	"github.com/fatih/color"
	"time"
)

const refreshTimeout = 2 * time.Second

const initialText = `Think you're good keyboardist?
Repeat this text with your input %s
And we will see how good is your touch typing!`

func main() {
	cls.CallCls()
	fmt.Println(initialText)
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("Some text %s", red("RedTextInsertion"))
	time.Sleep(refreshTimeout)
	cls.CallCls()
	fmt.Println(initialText)
}
