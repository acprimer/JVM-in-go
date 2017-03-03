package main

import "fmt"
import (
	"ch04/rtda"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.3")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
}

func testLocalVars(vars rtda.LocalVars)  {
	vars.SetInt(0, 100)
	vars.SetLong(1, 2997924580)
	vars.SetFloat(3, 3.14)
	vars.SetDouble(4, 2.71828)
	vars.SetRef(6, nil)
	fmt.Println(vars.GetInt(0))
	fmt.Println(vars.GetLong(1))
	fmt.Println(vars.GetFloat(3))
	fmt.Println(vars.GetDouble(4))
	fmt.Println(vars.GetRef(6))
}