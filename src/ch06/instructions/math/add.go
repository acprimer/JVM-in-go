package math

import (
	"ch06/instructions/base"
	"ch06/rtda"
)

type DADD struct{ base.NoOperandsInstruction }
type FADD struct{ base.NoOperandsInstruction }
type IADD struct{ base.NoOperandsInstruction }
type LADD struct{ base.NoOperandsInstruction }

func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}