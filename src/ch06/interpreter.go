package main

import (
	"fmt"
	"ch06/rtda"
	"ch06/instructions/base"
	"ch06/instructions"
)

func interpret(thread *rtda.Thread) {
	defer catchErr(thread)
	loop(thread)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		for !thread.IsStackEmpty() {
			frame := thread.PopFrame()
			method := frame.Method()
			className := method.Class().Name()
			fmt.Printf(">> pc:%4d %v.%v%v \n",
				frame.NextPC(), className, method.Name(), method.Descriptor())
		}
		panic(r)
	}
}

func loop(thread *rtda.Thread) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		// execute
		fmt.Printf("pc:%2d instruction: %T %v\n", pc, inst, inst)
		inst.Execute(frame)
		//frame.Print()
		if thread.IsStackEmpty() {
			break
		}
	}
}
