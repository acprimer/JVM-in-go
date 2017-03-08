package main

import (
	"ch06/rtda/heap"
	"ch06/rtda"
	"ch06/classpath"
	"ch06/instructions/base"
	"strings"
	"fmt"
)

type JVM struct {
	cmd *Cmd
	classLoader *heap.ClassLoader
	mainThread *rtda.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp)
	return &JVM{
		cmd:cmd,
		classLoader:classLoader,
		mainThread:rtda.NewThread(),
	}
}

func (self *JVM) start() {
	self.init()
	self.execMain()
}

func (self *JVM) init() {
	vmClass := self.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(self.mainThread, vmClass)
}

func (self *JVM) execMain() {
	className := strings.Replace(self.cmd.class, ".", "/", -1)
	mainClass := self.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", self.cmd.class)
	}
}