package math

import (
	"ch06/instructions/base"
	"ch06/rtda"
)

type DDIV struct{ base.NoOperandsInstruction }
type FDIV struct{ base.NoOperandsInstruction }
type IDIV struct{ base.NoOperandsInstruction }
type LDIV struct{ base.NoOperandsInstruction }

func (self *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

func (self *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

func (self *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 / v2
	stack.PushInt(result)
}

func (self *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 / v2
	stack.PushLong(result)
}